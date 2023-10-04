// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../../core/05-port/IIBCModule.sol";
import "../../core/04-channel/IIBCChannel.sol";
import "../commons/IBCAppBase.sol";

/// by default delegates all call to underlying app
abstract contract IIBCMiddleware is IBCAppBase {
    IBCAppBase _ibcModule;
    IICS04Wrapper _ics04Wrapper;

    constructor(address ibcModule, address ics04Wrapper) {
        require(ibcModule != address(0), "IIBCModule is the zero address");
        require(ics04Wrapper != address(0), "IICS04Wrapper is the zero address");
        _ibcModule = IBCAppBase(ibcModule);
        _ics04Wrapper = IICS04Wrapper(ics04Wrapper);
    }

    function ibcAddress() public view virtual override returns (address) {
        return _ibcModule.ibcAddress();
    }

    function onChanOpenInit(
        Channel.Order order,
        string[] calldata connectionHops,
        string calldata portId,
        string calldata channelId,
        ChannelCounterparty.Data calldata counterparty,
        string calldata version
    ) external virtual override onlyIBC {
        _ibcModule.onChanOpenInit(order, connectionHops, portId, channelId, counterparty, version);
    }

    /// ... here to fill all remaining default delegations
}
