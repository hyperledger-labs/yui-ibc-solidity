// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

import {BytesLib} from "solidity-bytes-utils/contracts/BytesLib.sol";
import {Height} from "../../proto/Client.sol";
import {IIBCHandler} from "../../core/25-handler/IIBCHandler.sol";
import {ICS20Transfer} from "./ICS20Transfer.sol";
import {IICS20Bank} from "./IICS20Bank.sol";
import {ICS20Lib} from "./ICS20Lib.sol";

contract ICS20TransferBank is ICS20Transfer {
    using BytesLib for bytes;

    IIBCHandler private immutable ibcHandler;
    IICS20Bank private immutable bank;

    constructor(IIBCHandler ibcHandler_, IICS20Bank bank_) {
        ibcHandler = ibcHandler_;
        bank = bank_;
    }

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
    ) external {
        require(ICS20Lib.isEscapedJSONString(receiver), "unescaped receiver");
        bytes memory denomPrefix = _getDenomPrefix(sourcePort, sourceChannel);
        bytes memory denomBytes = bytes(denom);
        if (denomBytes.length < denomPrefix.length || !denomBytes.slice(0, denomPrefix.length).equal(denomPrefix)) {
            // sender is source chain
            require(_transferFrom(_msgSender(), _getEscrowAddress(sourceChannel), denom, amount));
        } else {
            require(_burn(_msgSender(), denom, amount));
        }
        bytes memory packetData = ICS20Lib.marshalJSON(denom, amount, _encodeSender(_msgSender()), receiver);
        IIBCHandler(ibcAddress()).sendPacket(
            sourcePort, sourceChannel, Height.Data({revision_number: 0, revision_height: timeoutHeight}), 0, packetData
        );
    }

    function _transferFrom(address sender, address receiver, string memory denom, uint256 amount)
        internal
        override
        returns (bool)
    {
        try bank.transferFrom(sender, receiver, denom, amount) {
            return true;
        } catch (bytes memory) {
            return false;
        }
    }

    function _mint(address account, string memory denom, uint256 amount) internal override returns (bool) {
        try bank.mint(account, denom, amount) {
            return true;
        } catch (bytes memory) {
            return false;
        }
    }

    function _burn(address account, string memory denom, uint256 amount) internal override returns (bool) {
        try bank.burn(account, denom, amount) {
            return true;
        } catch (bytes memory) {
            return false;
        }
    }

    function ibcAddress() public view virtual override returns (address) {
        return address(ibcHandler);
    }
}
