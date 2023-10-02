// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "./ICS20Transfer.sol";
import "./IICS20Bank.sol";
import "../../core/25-handler/IBCHandler.sol";

contract ICS20TransferBank is ICS20Transfer {
    using strings for *;

    IBCHandler private immutable ibcHandler;
    IICS20Bank private immutable bank;

    constructor(IBCHandler ibcHandler_, IICS20Bank bank_) {
        ibcHandler = ibcHandler_;
        bank = bank_;
    }

    function sendTransfer(
        string calldata denom,
        uint256 amount,
        string calldata receiver,
        string calldata sourcePort,
        string calldata sourceChannel,
        uint64 timeoutHeight
    ) external {
        if (!denom.toSlice().startsWith(_makeDenomPrefix(sourcePort, sourceChannel))) {
            // sender is source chain
            require(_transferFrom(_msgSender(), _getEscrowAddress(sourceChannel), denom, amount));
        } else {
            require(_burn(_msgSender(), denom, amount));
        }

        ICS20Lib.PacketData memory data = ICS20Lib.PacketData({
            denom: denom,
            amount: amount,
            sender: _encodeSender(_msgSender()),
            receiver: receiver,
            memo: ""
        });
        IBCHandler(ibcAddress()).sendPacket(
            sourcePort,
            sourceChannel,
            Height.Data({revision_number: 0, revision_height: timeoutHeight}),
            0,
            ICS20Lib.marshalUnsafeJSON(data)
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

    function _encodeSender(address sender) internal pure override returns (string memory) {
        return ICS20Lib.addressToHexString(sender);
    }

    function _decodeSender(string memory sender) internal pure override returns (address) {
        (address addr, bool ok) = ICS20Lib.hexStringToAddress(sender);
        require(ok);
        return addr;
    }

    function _decodeReceiver(string memory receiver) internal pure override returns (address, bool) {
        return ICS20Lib.hexStringToAddress(receiver);
    }

    function ibcAddress() public view virtual override returns (address) {
        return address(ibcHandler);
    }
}
