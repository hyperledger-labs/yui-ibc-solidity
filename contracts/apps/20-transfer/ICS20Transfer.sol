// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IBCAppBase} from "../commons/IBCAppBase.sol";
import {Packet} from "../../core/04-channel/IIBCChannel.sol";
import {IIBCModule} from "../../core/26-router/IIBCModule.sol";
import {Height} from "../../proto/Client.sol";
import {Channel} from "../../proto/Channel.sol";
import {ICS20Lib} from "./ICS20Lib.sol";
import {IICS20Errors} from "./IICS20Errors.sol";
import {IIBCHandler} from "../../core/25-handler/IIBCHandler.sol";

contract ICS20Transfer is IBCAppBase, IICS20Errors {
    string public constant ICS20_VERSION = "ics20-1";

    // mapping from denomination to account balances
    mapping(string denom => mapping(address account => uint256 balance)) internal _balances;

    /// @dev IIBCHandler instance
    IIBCHandler internal immutable ibcHandler;

    /// @param ibcHandler_ IIBCHandler instance
    constructor(IIBCHandler ibcHandler_) {
        ibcHandler = ibcHandler_;
    }

    // ------------------------------ Public Functions ------------------------------ //

    /**
     * @dev sendTransfer sends a transfer packet to the destination chain.
     * @param denom denomination of the token. It can assume the denom string is escaped or not required to be escaped.
     * @param amount amount of the token
     * @param receiver receiver address on the destination chain
     * @param sourcePort source port of the packet
     * @param sourceChannel source channel of the packet
     * @param timeoutHeight timeout height of the packet
     */
    function sendTransfer(
        string calldata denom,
        uint256 amount,
        string calldata receiver,
        string calldata sourcePort,
        string calldata sourceChannel,
        uint64 timeoutHeight
    ) external returns (uint64) {
        if (!ICS20Lib.isEscapedJSONString(receiver)) {
            revert ICS20InvalidReceiverAddress(receiver);
        }
        bytes memory denomPrefix = ICS20Lib.denomPrefix(sourcePort, sourceChannel);
        bytes memory denomBytes = bytes(denom);
        if (
            denomBytes.length < denomPrefix.length
                || !ICS20Lib.equal(ICS20Lib.slice(denomBytes, 0, denomPrefix.length), denomPrefix)
        ) {
            // src chain is the source of the token
            if (!_transferVoucherFrom(_msgSender(), getVoucherEscrow(sourceChannel), denom, amount)) {
                revert ICS20InsufficientBalance(_msgSender(), _balances[denom][_msgSender()], amount);
            }
        } else {
            // dst chain is the source of the token
            _burnVoucher(_msgSender(), denom, amount);
        }
        bytes memory packetData = ICS20Lib.marshalJSON(denom, amount, encodeAddress(_msgSender()), receiver);
        return ibcHandler.sendPacket(
            sourcePort, sourceChannel, Height.Data({revision_number: 0, revision_height: timeoutHeight}), 0, packetData
        );
    }

    /**
     * @dev depositSendTransfer sends a transfer packet to the destination chain after depositing the token.
     * @param tokenContract address of the token contract
     * @param amount amount of the token
     * @param receiver receiver address on the destination chain
     * @param sourcePort source port of the packet
     * @param sourceChannel source channel of the packet
     * @param timeoutHeight timeout height of the packet
     */
    function depositSendTransfer(
        address tokenContract,
        uint256 amount,
        string calldata receiver,
        string calldata sourcePort,
        string calldata sourceChannel,
        uint64 timeoutHeight
    ) external returns (uint64) {
        if (!ICS20Lib.isEscapedJSONString(receiver)) {
            revert ICS20InvalidReceiverAddress(receiver);
        }
        if (tokenContract == address(0)) {
            revert ICS20InvalidTokenContract(tokenContract);
        }
        address sender = _msgSender();
        if (!IERC20(tokenContract).transferFrom(sender, address(this), amount)) {
            revert ICS20FailedERC20Transfer(tokenContract, sender, address(this), amount);
        }
        _mintVoucher(getVoucherEscrow(sourceChannel), tokenContract, amount);
        bytes memory packetData =
            ICS20Lib.marshalJSON(ICS20Lib.addressToHexString(tokenContract), amount, encodeAddress(sender), receiver);
        return ibcHandler.sendPacket(
            sourcePort, sourceChannel, Height.Data({revision_number: 0, revision_height: timeoutHeight}), 0, packetData
        );
    }

    /**
     * @dev deposit deposits the ERC20 token to the contract.
     * @param tokenContract address of the token contract
     * @param amount amount of the token
     * @param to address to deposit the token
     */
    function deposit(address to, address tokenContract, uint256 amount) public {
        if (tokenContract == address(0)) {
            revert ICS20InvalidTokenContract(tokenContract);
        }
        address from = _msgSender();
        if (!IERC20(tokenContract).transferFrom(from, address(this), amount)) {
            revert ICS20FailedERC20Transfer(tokenContract, from, address(this), amount);
        }
        _mintVoucher(to, tokenContract, amount);
    }

    /**
     * @dev withdraw withdraws the ERC20 token from the contract.
     * @param tokenContract address of the token contract
     * @param amount amount of the token
     * @param to address to withdraw the token
     */
    function withdraw(address to, address tokenContract, uint256 amount) public {
        if (tokenContract == address(0)) {
            revert ICS20InvalidTokenContract(tokenContract);
        }
        address from = _msgSender();
        _burnVoucher(from, ICS20Lib.addressToHexString(tokenContract), amount);
        if (!IERC20(tokenContract).transfer(to, amount)) {
            revert ICS20FailedERC20TransferFrom(tokenContract, from, address(this), to, amount);
        }
    }

    /**
     * @dev transfer transfers the token to the given address.
     * @param to address to transfer the token
     * @param denom denomination of the token
     * @param amount amount of the token
     */
    function transfer(address to, string calldata denom, uint256 amount) public {
        if (to == address(0)) {
            revert ICS20InvalidReceiver(to);
        }
        address from = _msgSender();
        _burnVoucher(from, denom, amount);
        _mintVoucher(to, denom, amount);
    }

    /**
     * @dev balanceOf returns the balance of the account for the given denomination.
     * @param account account address
     * @param denom denomination of the token
     */
    function balanceOf(address account, string calldata denom) public view virtual returns (uint256) {
        return _balances[denom][account];
    }

    /**
     * @dev encodeAddress encodes an address to a hex string.
     *      The encoded address is used as `sender` field in the packet data.
     */
    function encodeAddress(address sender) public pure virtual returns (string memory) {
        return ICS20Lib.addressToHexString(sender);
    }

    /**
     * @dev getVoucherEscrow returns the voucher escrow address for the given channel.
     * @param channelId channel identifier
     */
    function getVoucherEscrow(string calldata channelId) public view virtual returns (address) {
        return address(uint160(uint256(keccak256(abi.encode(address(this), channelId)))));
    }

    /**
     * @dev ibcAddress returns the address of the IBC handler.
     */
    function ibcAddress() public view virtual override returns (address) {
        return address(ibcHandler);
    }

    // ------------------------------ IBC Module Callbacks ------------------------------ //

    /**
     * @dev See {IIBCModule-onRecvPacket}
     */
    function onRecvPacket(Packet calldata packet, address)
        public
        virtual
        override
        onlyIBC
        returns (bytes memory acknowledgement)
    {
        ICS20Lib.PacketData memory data = ICS20Lib.unmarshalJSON(packet.data);
        (address receiver, bool success) = _decodeReceiver(data.receiver);
        if (!success) {
            return ICS20Lib.FAILED_ACKNOWLEDGEMENT_JSON;
        }

        bytes memory denomPrefix = ICS20Lib.denomPrefix(packet.sourcePort, packet.sourceChannel);
        bytes memory denom = bytes(data.denom);
        if (
            denom.length >= denomPrefix.length
                && ICS20Lib.equal(ICS20Lib.slice(denom, 0, denomPrefix.length), denomPrefix)
        ) {
            // sender chain is not the source, unescrow tokens
            bytes memory unprefixedDenom = ICS20Lib.slice(denom, denomPrefix.length, denom.length - denomPrefix.length);
            if (
                _transferVoucherFrom(
                    getVoucherEscrow(packet.destinationChannel), receiver, string(unprefixedDenom), data.amount
                )
            ) {
                return ICS20Lib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON;
            } else {
                return ICS20Lib.FAILED_ACKNOWLEDGEMENT_JSON;
            }
        }

        // sender chain is the source, mint vouchers

        // ensure denom is not required to be escaped
        if (ICS20Lib.isEscapeNeededString(denom)) {
            return ICS20Lib.FAILED_ACKNOWLEDGEMENT_JSON;
        } else {
            _mintVoucher(
                receiver,
                string(abi.encodePacked(ICS20Lib.denomPrefix(packet.destinationPort, packet.destinationChannel), denom)),
                data.amount
            );
            return ICS20Lib.SUCCESSFUL_ACKNOWLEDGEMENT_JSON;
        }
    }

    /**
     * @dev See {IIBCModule-onAcknowledgementPacket}
     */
    function onAcknowledgementPacket(Packet calldata packet, bytes calldata acknowledgement, address)
        public
        virtual
        override
        onlyIBC
    {
        if (keccak256(acknowledgement) != ICS20Lib.KECCAK256_SUCCESSFUL_ACKNOWLEDGEMENT_JSON) {
            _refundTokens(ICS20Lib.unmarshalJSON(packet.data), packet.sourcePort, packet.sourceChannel);
        }
    }

    /**
     * @dev See {IIBCModule-onTimeoutPacket}
     */
    function onTimeoutPacket(Packet calldata packet, address) public virtual override onlyIBC {
        _refundTokens(ICS20Lib.unmarshalJSON(packet.data), packet.sourcePort, packet.sourceChannel);
    }

    /**
     * @dev See {IIBCModule-onChanOpenInit}
     */
    function onChanOpenInit(IIBCModule.MsgOnChanOpenInit calldata msg_)
        public
        virtual
        override
        onlyIBC
        returns (address, string memory)
    {
        if (msg_.order != Channel.Order.ORDER_UNORDERED) {
            revert IBCModuleChannelOrderNotAllowed(msg_.portId, msg_.channelId, msg_.order);
        }
        bytes memory versionBytes = bytes(msg_.version);
        if (versionBytes.length != 0 && keccak256(versionBytes) != keccak256(bytes(ICS20_VERSION))) {
            revert ICS20UnexpectedVersion(msg_.version);
        }
        return (address(this), ICS20_VERSION);
    }

    /**
     * @dev See {IIBCModule-onChanOpenTry}
     */
    function onChanOpenTry(IIBCModule.MsgOnChanOpenTry calldata msg_)
        public
        virtual
        override
        onlyIBC
        returns (address, string memory)
    {
        if (msg_.order != Channel.Order.ORDER_UNORDERED) {
            revert IBCModuleChannelOrderNotAllowed(msg_.portId, msg_.channelId, msg_.order);
        }
        if (keccak256(bytes(msg_.counterpartyVersion)) != keccak256(bytes(ICS20_VERSION))) {
            revert ICS20UnexpectedVersion(msg_.counterpartyVersion);
        }
        return (address(this), ICS20_VERSION);
    }

    /**
     * @dev See {IIBCModule-onChanOpenAck}
     */
    function onChanOpenAck(IIBCModule.MsgOnChanOpenAck calldata msg_) public virtual override onlyIBC {
        if (keccak256(bytes(msg_.counterpartyVersion)) != keccak256(bytes(ICS20_VERSION))) {
            revert ICS20UnexpectedVersion(msg_.counterpartyVersion);
        }
    }

    /**
     * @dev See {IIBCModule-onChanCloseInit}
     */
    function onChanCloseInit(IIBCModule.MsgOnChanCloseInit calldata msg_) public virtual override onlyIBC {
        revert IBCModuleChannelCloseNotAllowed(msg_.portId, msg_.channelId);
    }

    // ------------------------------ Internal Functions ------------------------------ //

    /**
     * @dev _mintVoucher mints the token to the given address.
     * @param to address to mint the token
     * @param denom denomination of the token
     * @param amount amount of the token
     */
    function _mintVoucher(address to, string memory denom, uint256 amount) internal {
        _balances[denom][to] += amount;
    }

    /**
     * @dev _mintVoucher mints the token to the given address.
     * @param to address to mint the token
     * @param tokenContract address of the token contract
     * @param amount amount of the token
     */
    function _mintVoucher(address to, address tokenContract, uint256 amount) internal {
        _balances[ICS20Lib.addressToHexString(tokenContract)][to] += amount;
    }

    /**
     * @dev _burnVoucher burns the token from the given address.
     * @param from address to burn the token
     * @param denom denomination of the token
     * @param amount amount of the token
     */
    function _burnVoucher(address from, string memory denom, uint256 amount) internal {
        uint256 accountBalance = _balances[denom][from];
        if (accountBalance < amount) {
            revert ICS20InsufficientBalance(from, accountBalance, amount);
        }
        unchecked {
            // SAFETY: balance is checked above
            _balances[denom][from] = accountBalance - amount;
        }
    }

    /**
     * @dev _transferVoucherFrom transfers the token from the sender to the receiver.
     * @param from address to transfer the token from
     * @param to address to transfer the token to
     * @param denom denomination of the token
     * @param amount amount of the token
     * @return true if the transfer is successful, false otherwise
     */
    function _transferVoucherFrom(address from, address to, string memory denom, uint256 amount)
        internal
        returns (bool)
    {
        if (from == address(0) || to == address(0)) {
            return false;
        }
        uint256 fromBalance = _balances[denom][from];
        if (fromBalance < amount) {
            return false;
        }
        unchecked {
            // SAFETY: balance is checked above
            _balances[denom][from] = fromBalance - amount;
        }
        _balances[denom][to] += amount;
        return true;
    }

    /**
     * @dev _refundTokens refunds the tokens to the sender.
     * @param data packet data
     * @param sourcePort source port of the packet
     * @param sourceChannel source channel of the packet
     */
    function _refundTokens(ICS20Lib.PacketData memory data, string calldata sourcePort, string calldata sourceChannel)
        internal
        virtual
    {
        bytes memory denomPrefix = ICS20Lib.denomPrefix(sourcePort, sourceChannel);
        bytes memory denom = bytes(data.denom);
        address sender = _decodeSender(data.sender);
        if (
            denom.length >= denomPrefix.length
                && ICS20Lib.equal(ICS20Lib.slice(denom, 0, denomPrefix.length), denomPrefix)
        ) {
            _mintVoucher(sender, data.denom, data.amount);
        } else {
            // sender was source chain
            address escrow = getVoucherEscrow(sourceChannel);
            if (!_transferVoucherFrom(escrow, sender, data.denom, data.amount)) {
                revert ICS20FailedRefund(escrow, sender, data.denom, data.amount);
            }
        }
    }

    /**
     * @dev _decodeSender decodes a hex string to an address.
     *      `sender` must be a valid address format.
     */
    function _decodeSender(string memory sender) internal pure virtual returns (address) {
        (address addr, bool ok) = ICS20Lib.hexStringToAddress(sender);
        if (!ok) {
            revert ICS20InvalidSenderAddress(sender);
        }
        return addr;
    }

    /**
     * @dev _decodeSender decodes a hex string to an address.
     *       `receiver` may be an invalid address format.
     */
    function _decodeReceiver(string memory receiver) internal pure virtual returns (address, bool) {
        return ICS20Lib.hexStringToAddress(receiver);
    }
}
