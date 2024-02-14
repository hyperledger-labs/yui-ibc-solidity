// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "./helpers/ICS03TestHelper.t.sol";

contract TestICS03Version is ICS03TestHelper {
    function testIsSupportedVersion() public {
        Version.Data[] memory versions = getConnectionVersions();
        assertTrue(IBCConnectionLib.isSupportedVersion(versions, versions[0]));
        Version.Data memory version = Version.Data({identifier: "", features: new string[](0)});
        assertFalse(IBCConnectionLib.isSupportedVersion(versions, version));
        version = Version.Data({identifier: "1", features: new string[](1)});
        version.features[0] = "ORDER_DAG";
        assertFalse(IBCConnectionLib.isSupportedVersion(versions, version));
    }

    function testFindSupportedVersion() public {
        // "valid supported version"
        {
            Version.Data[] memory versions = getConnectionVersions();
            (Version.Data memory v, bool found) =
                IBCConnectionLib.findSupportedVersion(IBCConnectionLib.defaultIBCVersion(), versions);
            assertTrue(found);
            assertEq(v.identifier, versions[0].identifier);
        }
        // "empty (invalid) version"
        {
            Version.Data[] memory versions = getConnectionVersions();
            (, bool found) = IBCConnectionLib.findSupportedVersion(
                Version.Data({identifier: "", features: new string[](0)}), versions
            );
            assertFalse(found);
        }
        // "empty supported versions"
        {
            Version.Data[] memory versions = new Version.Data[](0);
            (, bool found) = IBCConnectionLib.findSupportedVersion(IBCConnectionLib.defaultIBCVersion(), versions);
            assertFalse(found);
        }
        // "desired version is last"
        {
            Version.Data[] memory versions = new Version.Data[](4);
            versions[0] = Version.Data({identifier: "1.1", features: new string[](0)});
            versions[1] = Version.Data({identifier: "2", features: new string[](1)});
            versions[1].features[0] = "ORDER_UNORDERED";
            versions[2] = Version.Data({identifier: "3", features: new string[](0)});
            versions[3] = IBCConnectionLib.defaultIBCVersion();
            (Version.Data memory v, bool found) =
                IBCConnectionLib.findSupportedVersion(IBCConnectionLib.defaultIBCVersion(), versions);
            assertTrue(found);
            assertEq(v.identifier, versions[3].identifier);
        }
        // "desired version identifier with different feature set"
        {
            Version.Data[] memory versions = new Version.Data[](1);
            versions[0] =
                Version.Data({identifier: IBCConnectionLib.defaultIBCVersion().identifier, features: new string[](1)});
            versions[0].features[0] = "ORDER_DAG";
            (Version.Data memory v, bool found) =
                IBCConnectionLib.findSupportedVersion(IBCConnectionLib.defaultIBCVersion(), versions);
            assertTrue(found);
            assertEq(v.identifier, IBCConnectionLib.defaultIBCVersion().identifier);
        }
        // "version not supported"
        {
            Version.Data[] memory versions = new Version.Data[](1);
            versions[0] = Version.Data({identifier: "2", features: new string[](1)});
            versions[0].features[0] = "ORDER_DAG";
            (, bool found) = IBCConnectionLib.findSupportedVersion(IBCConnectionLib.defaultIBCVersion(), versions);
            assertFalse(found);
        }
    }

    function testPickVersion() public {
        // "valid default ibc version"
        {
            Version.Data[] memory supportedVersions = getConnectionVersions();
            Version.Data[] memory counterpartyVersions = getConnectionVersions();
            Version.Data memory v = IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
            assertEq(v.identifier, IBCConnectionLib.defaultIBCVersion().identifier);
        }
        // "valid version in counterparty versions"
        {
            Version.Data[] memory supportedVersions = getConnectionVersions();
            Version.Data[] memory counterpartyVersions = new Version.Data[](3);
            counterpartyVersions[0] = Version.Data({identifier: "version1", features: new string[](0)});
            counterpartyVersions[1] = Version.Data({identifier: "2.0.0", features: new string[](1)});
            counterpartyVersions[1].features[0] = "ORDER_UNORDERED-ZK";
            counterpartyVersions[2] = IBCConnectionLib.defaultIBCVersion();
            Version.Data memory v = IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
            assertEq(v.identifier, IBCConnectionLib.defaultIBCVersion().identifier);
        }
        // "valid identifier match but empty feature set not allowed"
        {
            Version.Data[] memory supportedVersions = getConnectionVersions();
            Version.Data[] memory counterpartyVersions = new Version.Data[](2);
            counterpartyVersions[0] =
                Version.Data({identifier: IBCConnectionLib.defaultIBCVersion().identifier, features: new string[](3)});
            counterpartyVersions[0].features[0] = "DAG";
            counterpartyVersions[0].features[1] = "ORDERED-ZK";
            counterpartyVersions[0].features[2] = "UNORDERED-zk";
            counterpartyVersions[1] =
                Version.Data({identifier: IBCConnectionLib.defaultIBCVersion().identifier, features: new string[](0)});
            vm.expectRevert();
            IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
        }
        // "empty counterparty versions"
        {
            Version.Data[] memory supportedVersions = getConnectionVersions();
            Version.Data[] memory counterpartyVersions = new Version.Data[](0);
            vm.expectRevert();
            IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
        }
        // "non-matching counterparty versions"
        {
            Version.Data[] memory supportedVersions = getConnectionVersions();
            Version.Data[] memory counterpartyVersions = new Version.Data[](2);
            counterpartyVersions[0] = Version.Data({identifier: "2.0.0", features: new string[](0)});
            counterpartyVersions[1] = Version.Data({identifier: "", features: new string[](0)});
            vm.expectRevert();
            IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
        }
        // "non-matching counterparty versions (uses ordered channels only) contained in supported versions (uses unordered channels only)"
        {
            Version.Data[] memory supportedVersions = new Version.Data[](1);
            supportedVersions[0] =
                Version.Data({identifier: IBCConnectionLib.defaultIBCVersion().identifier, features: new string[](1)});
            supportedVersions[0].features[0] = "ORDER_UNORDERED";
            Version.Data[] memory counterpartyVersions = new Version.Data[](1);
            counterpartyVersions[0] =
                Version.Data({identifier: IBCConnectionLib.defaultIBCVersion().identifier, features: new string[](1)});
            counterpartyVersions[0].features[0] = "ORDER_ORDERED";
            vm.expectRevert();
            IBCConnectionLib.pickVersion(supportedVersions, counterpartyVersions);
        }
    }

    function testVerifyProposedVersion() public {
        // "entire feature set supported"
        {
            Version.Data memory proposedVersion = IBCConnectionLib.defaultIBCVersion();
            Version.Data memory supportedVersion = Version.Data({identifier: "1", features: new string[](3)});
            supportedVersion.features[0] = "ORDER_ORDERED";
            supportedVersion.features[1] = "ORDER_UNORDERED";
            supportedVersion.features[2] = "ORDER_DAG";
            assertTrue(IBCConnectionLib.verifyProposedVersion(supportedVersion, proposedVersion));
        }
        // "empty feature sets not supported"
        {
            Version.Data memory proposedVersion = Version.Data({identifier: "1", features: new string[](0)});
            Version.Data memory supportedVersion = IBCConnectionLib.defaultIBCVersion();
            assertFalse(IBCConnectionLib.verifyProposedVersion(supportedVersion, proposedVersion));
        }
        // "one feature missing"
        {
            Version.Data memory proposedVersion = IBCConnectionLib.defaultIBCVersion();
            Version.Data memory supportedVersion = Version.Data({identifier: "1", features: new string[](2)});
            supportedVersion.features[0] = "ORDER_UNORDERED";
            supportedVersion.features[1] = "ORDER_DAG";
            assertFalse(IBCConnectionLib.verifyProposedVersion(supportedVersion, proposedVersion));
        }
        // "both features missing"
        {
            Version.Data memory proposedVersion = IBCConnectionLib.defaultIBCVersion();
            Version.Data memory supportedVersion = Version.Data({identifier: "1", features: new string[](1)});
            supportedVersion.features[0] = "ORDER_DAG";
            assertFalse(IBCConnectionLib.verifyProposedVersion(supportedVersion, proposedVersion));
        }
        // "identifiers do not match"
        {
            Version.Data memory proposedVersion = Version.Data({identifier: "2", features: new string[](2)});
            proposedVersion.features[0] = "ORDER_UNORDERED";
            proposedVersion.features[1] = "ORDER_ORDERED";
            Version.Data memory supportedVersion = IBCConnectionLib.defaultIBCVersion();
            assertFalse(IBCConnectionLib.verifyProposedVersion(supportedVersion, proposedVersion));
        }
    }

    function testVerifySupportedFeature() public {
        // "check ORDERED supported"
        {
            Version.Data memory version = IBCConnectionLib.defaultIBCVersion();
            assertTrue(IBCConnectionLib.verifySupportedFeature(version, "ORDER_ORDERED"));
        }
        // "check UNORDERED supported"
        {
            Version.Data memory version = IBCConnectionLib.defaultIBCVersion();
            assertTrue(IBCConnectionLib.verifySupportedFeature(version, "ORDER_UNORDERED"));
        }
        // "check DAG unsupported"
        {
            Version.Data memory version = IBCConnectionLib.defaultIBCVersion();
            assertFalse(IBCConnectionLib.verifySupportedFeature(version, "ORDER_DAG"));
        }
        // "check empty feature set returns false"
        {
            Version.Data memory version = Version.Data({identifier: "1", features: new string[](0)});
            assertFalse(IBCConnectionLib.verifySupportedFeature(version, "ORDER_ORDERED"));
        }
    }

    Version.Data[] internal testVersions;

    function testCopyVersions() public {
        {
            clearVersions();
            Version.Data[] memory vs = getConnectionVersions();
            IBCConnectionLib.copyVersions(vs, testVersions);
            matchVersions(vs);
        }
        {
            clearVersions();
            Version.Data[] memory vs = new Version.Data[](0);
            IBCConnectionLib.copyVersions(vs, testVersions);
            matchVersions(vs);
        }
        {
            clearVersions();
            Version.Data[] memory vs = new Version.Data[](2);
            vs[0] = IBCConnectionLib.defaultIBCVersion();
            vs[1] = Version.Data({identifier: "2", features: new string[](1)});
            vs[1].features[0] = "ORDER_DAG";
            IBCConnectionLib.copyVersions(vs, testVersions);
            matchVersions(vs);
        }
        {
            clearVersions();
            testVersions.push(IBCConnectionLib.defaultIBCVersion());
            Version.Data[] memory vs = new Version.Data[](1);
            vs[0] = Version.Data({identifier: "2", features: new string[](1)});
            vs[0].features[0] = "ORDER_DAG";
            IBCConnectionLib.copyVersions(vs, testVersions);
            matchVersions(vs);
        }
        {
            clearVersions();
            testVersions.push(IBCConnectionLib.defaultIBCVersion());
            Version.Data[] memory vs = new Version.Data[](0);
            IBCConnectionLib.copyVersions(vs, testVersions);
            matchVersions(vs);
        }
        {
            clearVersions();
            testVersions.push(IBCConnectionLib.defaultIBCVersion());
            testVersions.push(Version.Data({identifier: "2", features: new string[](1)}));
            testVersions[1].features[0] = "ORDER_DAG";
            Version.Data[] memory vs = new Version.Data[](1);
            vs[0] = Version.Data({identifier: "3", features: new string[](1)});
            vs[0].features[0] = "ORDERED-ZK";
            IBCConnectionLib.copyVersions(vs, testVersions);
            matchVersions(vs);
        }
    }

    function clearVersions() internal {
        uint256 versionsLength = testVersions.length;
        for (uint256 i = 0; i < versionsLength; i++) {
            testVersions.pop();
        }
        assert(testVersions.length == 0);
    }

    function matchVersions(Version.Data[] memory vs) internal {
        assertEq(testVersions.length, vs.length);
        for (uint256 i = 0; i < vs.length; i++) {
            assertEq(testVersions[i].identifier, vs[i].identifier);
            assertEq(testVersions[i].features.length, vs[i].features.length);
            for (uint256 j = 0; j < vs[i].features.length; j++) {
                assertEq(testVersions[i].features[j], vs[i].features[j]);
            }
        }
    }
}
