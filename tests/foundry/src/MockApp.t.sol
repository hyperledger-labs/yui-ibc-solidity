// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../../../contracts/proto/Channel.sol";
import "../../../contracts/core/25-handler/IBCHandler.sol";
import "../../../contracts/core/24-host/IBCHost.sol";
import "../../../contracts/apps/commons/IBCAppBase.sol";
import "@openzeppelin/contracts/utils/Context.sol";

contract MockApp is IBCAppBase {
    event MockRecv(bool ok);

    address immutable ibcAddr;

    constructor(address ibcAddr_) {
        ibcAddr = ibcAddr_;
    }

    function ibcAddress() public view virtual override returns (address) {
        return ibcAddr;
    }

    /// Module callbacks ///

    function onRecvPacket(Packet.Data calldata, address) external virtual override onlyIBC returns (bytes memory) {
        emit MockRecv(true);
        return bytes("1");
    }
}
