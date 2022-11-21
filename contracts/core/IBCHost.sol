// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.9;

import "../proto/Client.sol";
import "../proto/Connection.sol";
import "../proto/Channel.sol";
import "./IClient.sol";

contract IBCHost {
    // Commitments
    mapping(bytes32 => bytes32) internal commitments;

    // Store
    mapping(string => address) internal clientRegistry; // clientType => clientImpl
    mapping(string => string) internal clientTypes; // clientID => clientType
    mapping(string => address) internal clientImpls; // clientID => clientImpl

    mapping(string => ConnectionEnd.Data) internal connections;
    mapping(string => mapping(string => Channel.Data)) public channels;
    mapping(string => mapping(string => uint64)) public nextSequenceSends;
    mapping(string => mapping(string => uint64)) internal nextSequenceRecvs;
    mapping(string => mapping(string => uint64)) internal nextSequenceAcks;
    mapping(string => mapping(string => mapping(uint64 => uint8))) internal packetReceipts;
    mapping(bytes => address[]) internal capabilities;
    uint64 internal nextClientSequence;
    uint64 internal nextConnectionSequence;
    uint64 internal nextChannelSequence;
    uint64 internal expectedTimePerBlock;

    // TODO make these immutable and move into each contract
    // module addresses
    address internal ibcChannelAddress;
    address internal ibcConnectionAddress;
    address internal ibcClientAddress;

    // Events
    event GeneratedClientIdentifier(string);
    event GeneratedConnectionIdentifier(string);
    event GeneratedChannelIdentifier(string);

    // host functions

    function validateSelfClient(bytes memory) internal view returns (bool) {
        this; // this is a trick that suppresses "Warning: Function state mutability can be restricted to pure"
        return true;
    }

    // capabilities

    function claimCapability(bytes memory name, address addr) internal {
        for (uint32 i = 0; i < capabilities[name].length; i++) {
            require(capabilities[name][i] != addr);
        }
        capabilities[name].push(addr);
    }

    function authenticateCapability(bytes memory name, address addr) internal view returns (bool) {
        for (uint32 i = 0; i < capabilities[name].length; i++) {
            if (capabilities[name][i] == addr) {
                return true;
            }
        }
        return false;
    }

    function getModuleOwner(bytes memory name) internal view returns (address, bool) {
        if (capabilities[name].length == 0) {
            return (address(0), false);
        }
        return (capabilities[name][0], true);
    }

    // Identifier generators

    function generateClientIdentifier(string calldata clientType) internal returns (string memory) {
        string memory identifier = string(abi.encodePacked(clientType, "-", uint2str(nextClientSequence)));
        nextClientSequence++;
        emit GeneratedClientIdentifier(identifier);
        return identifier;
    }

    function generateConnectionIdentifier() internal returns (string memory) {
        string memory identifier = string(abi.encodePacked("connection-", uint2str(nextConnectionSequence)));
        nextConnectionSequence++;
        emit GeneratedConnectionIdentifier(identifier);
        return identifier;
    }

    function generateChannelIdentifier() internal returns (string memory) {
        string memory identifier = string(abi.encodePacked("channel-", uint2str(nextChannelSequence)));
        nextChannelSequence++;
        emit GeneratedChannelIdentifier(identifier);
        return identifier;
    }

    // Storage accessors

    function getClient(string memory clientId) internal view returns (IClient) {
        address clientImpl = clientImpls[clientId];
        require(clientImpl != address(0));
        return IClient(clientImpl);
    }

    // Utility functions

    function uint2str(uint64 _i) internal pure returns (string memory _uintAsString) {
        if (_i == 0) {
            return "0";
        }
        uint64 j = _i;
        uint64 len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint64 k = len;
        while (_i != 0) {
            k = k - 1;
            uint8 temp = (48 + uint8(_i - _i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
        }
        return string(bstr);
    }
}
