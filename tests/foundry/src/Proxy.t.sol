// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "../../../contracts/core/25-handler/IBCHandler.sol";
import "../../../contracts/core/02-client/IBCClient.sol";
import "../../../contracts/core/03-connection/IBCConnection.sol";
import "../../../contracts/core/04-channel/IBCChannelHandshake.sol";
import "../../../contracts/core/04-channel/IBCPacket.sol";
import "../../../contracts/core/24-host/IBCCommitment.sol";
import "../../../contracts/clients/MockClient.sol";
import "../../../contracts/proto/MockClient.sol";
import "../../../contracts/proto/Connection.sol";
import "../../../contracts/proto/Channel.sol";
import "forge-std/Test.sol";
import "./MockApp.t.sol";

contract UpgradableOwnableIBCHandler is IBCHandler, UUPSUpgradeable, Initializable, Ownable {
    constructor(address ibcClient, address ibcConnection, address ibcChannel, address ibcPacket)
        IBCHandler(ibcClient, ibcConnection, ibcChannel, ibcPacket)
    {}

    // NOTE: should we use reinitializer(1) instead as modifier?
    function initialize() public virtual initializer {
        _transferOwnership(_msgSender());
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}

    /**
     * @dev registerClient registers a new client type into the client registry
     */
    function registerClient(string calldata clientType, ILightClient client) public override onlyOwner {
        super.registerClient(clientType, client);
    }

    /**
     * @dev bindPort binds to an unallocated port, failing if the port has already been allocated.
     */
    function bindPort(string calldata portId, address moduleAddress) public override onlyOwner {
        super.bindPort(portId, moduleAddress);
    }

    /**
     * @dev setExpectedTimePerBlock sets expected time per block.
     */
    function setExpectedTimePerBlock(uint64 expectedTimePerBlock_) public override onlyOwner {
        super.setExpectedTimePerBlock(expectedTimePerBlock_);
    }

    function getVersion() public view virtual returns (uint256) {
        return _getInitializedVersion();
    }
}

contract NewUpgradableOwnableIBCHandler is UpgradableOwnableIBCHandler {
    constructor(address ibcClient, address ibcConnection, address ibcChannel, address ibcPacket)
        UpgradableOwnableIBCHandler(ibcClient, ibcConnection, ibcChannel, ibcPacket)
    {}

    function initialize() public virtual override reinitializer(2) {
        _transferOwnership(_msgSender());
    }

    function getVersion() public view virtual override returns (uint256) {
        return _getInitializedVersion();
    }
}

contract ProxyPattern is Test {
    ERC1967Proxy proxy;
    UpgradableOwnableIBCHandler logicHandler;
    UpgradableOwnableIBCHandler proxyHandler;

    string private constant MOCK_CLIENT_TYPE = "mock-client";
    string private constant MOCK_PORT_ID = "mock";

    function setUp() public {
        address ibcClient = address(new IBCClient());
        address ibcConnection = address(new IBCConnection());
        address ibcChannelHandshake = address(new IBCChannelHandshake());
        address ibcPacket = address(new IBCPacket());
        logicHandler = new UpgradableOwnableIBCHandler(ibcClient, ibcConnection, ibcChannelHandshake, ibcPacket);
        proxy = new ERC1967Proxy(address(logicHandler), abi.encodeWithSignature("initialize()"));
        proxyHandler = UpgradableOwnableIBCHandler(address(proxy));
        assertEq(proxyHandler.owner(), address(this));
    }

    function testProxy() public {
        proxyHandler.registerClient(MOCK_CLIENT_TYPE, new MockClient(address(proxy)));
        createMockClient(1);
        assertEq(1, proxyHandler.getVersion());

        address newLogicHandler = createNewUpgradableOwnableIBCHandler();
        proxyHandler.upgradeToAndCall(newLogicHandler, abi.encodeWithSignature("initialize()"));
        assertEq(2, proxyHandler.getVersion());
    }

    function createNewUpgradableOwnableIBCHandler() internal returns (address) {
        address ibcClient = address(new IBCClient());
        address ibcConnection = address(new IBCConnection());
        address ibcChannelHandshake = address(new IBCChannelHandshake());
        address ibcPacket = address(new IBCPacket());
        return address(new NewUpgradableOwnableIBCHandler(ibcClient, ibcConnection, ibcChannelHandshake, ibcPacket));
    }

    function createMockClient(uint64 revisionHeight) internal {
        proxyHandler.createClient(
            IBCMsgs.MsgCreateClient({
                clientType: MOCK_CLIENT_TYPE,
                clientStateBytes: wrapAnyMockClientState(
                    IbcLightclientsMockV1ClientState.Data({
                        latest_height: Height.Data({revision_number: 0, revision_height: revisionHeight})
                    })
                    ),
                consensusStateBytes: wrapAnyMockConsensusState(
                    IbcLightclientsMockV1ConsensusState.Data({timestamp: uint64(block.timestamp)})
                    )
            })
        );
    }

    function wrapAnyMockClientState(IbcLightclientsMockV1ClientState.Data memory clientState)
        internal
        pure
        returns (bytes memory)
    {
        Any.Data memory anyClientState;
        anyClientState.type_url = "/ibc.lightclients.mock.v1.ClientState";
        anyClientState.value = IbcLightclientsMockV1ClientState.encode(clientState);
        return Any.encode(anyClientState);
    }

    function wrapAnyMockConsensusState(IbcLightclientsMockV1ConsensusState.Data memory consensusState)
        internal
        pure
        returns (bytes memory)
    {
        Any.Data memory anyConsensusState;
        anyConsensusState.type_url = "/ibc.lightclients.mock.v1.ConsensusState";
        anyConsensusState.value = IbcLightclientsMockV1ConsensusState.encode(consensusState);
        return Any.encode(anyConsensusState);
    }
}
