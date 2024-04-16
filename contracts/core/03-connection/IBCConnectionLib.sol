// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import {Version} from "../../proto/Connection.sol";
import {IIBCConnectionErrors} from "./IIBCConnectionErrors.sol";

library IBCConnectionLib {
    string internal constant IBC_VERSION_IDENTIFIER = "1";
    string internal constant ORDER_ORDERED = "ORDER_ORDERED";
    string internal constant ORDER_UNORDERED = "ORDER_UNORDERED";

    /**
     * @dev defaultIBCVersion returns the latest supported version of IBC used in connection version negotiation
     */
    function defaultIBCVersion() internal pure returns (Version.Data memory) {
        Version.Data memory version = Version.Data({identifier: IBC_VERSION_IDENTIFIER, features: new string[](2)});
        version.features[0] = ORDER_ORDERED;
        version.features[1] = ORDER_UNORDERED;
        return version;
    }

    /**
     * @dev setSupportedVersions sets the supported versions to a given array.
     *
     * NOTE: `dst` must be an empty array
     */
    function setSupportedVersions(Version.Data[] memory supportedVersions, Version.Data[] storage dst) internal {
        if (dst.length != 0) {
            revert IIBCConnectionErrors.IBCConnectionVersionsAlreadySet();
        }
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            dst.push(supportedVersions[i]);
        }
    }

    /**
     * @dev isSupportedVersion returns true if the proposed version has a matching version
     * identifier and its entire feature set is supported or the version identifier
     * supports an empty feature set.
     */
    function isSupportedVersion(Version.Data[] memory supportedVersions, Version.Data memory version)
        internal
        pure
        returns (bool)
    {
        (Version.Data memory supportedVersion, bool found) = findSupportedVersion(version, supportedVersions);
        if (!found) {
            return false;
        }
        return verifyProposedVersion(supportedVersion, version);
    }

    function isSupported(Version.Data[] storage supportedVersions, string memory feature)
        internal
        view
        returns (bool)
    {
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            if (verifySupportedFeature(supportedVersions[i], feature)) {
                return true;
            }
        }
        return false;
    }

    /**
     * @dev verifyProposedVersion verifies that the entire feature set in the
     * proposed version is supported by this chain. If the feature set is
     * empty it verifies that this is allowed for the specified version
     * identifier.
     */
    function verifyProposedVersion(Version.Data memory supportedVersion, Version.Data memory proposedVersion)
        internal
        pure
        returns (bool)
    {
        if (
            keccak256(abi.encodePacked(proposedVersion.identifier))
                != keccak256(abi.encodePacked(supportedVersion.identifier))
        ) {
            return false;
        }
        if (proposedVersion.features.length == 0) {
            return false;
        }
        for (uint256 i = 0; i < proposedVersion.features.length; i++) {
            if (!contains(proposedVersion.features[i], supportedVersion.features)) {
                return false;
            }
        }
        return true;
    }

    /**
     * @dev findSupportedVersion returns the version with a matching version identifier
     * if it exists. The returned boolean is true if the version is found and
     * false otherwise.
     */
    function findSupportedVersion(Version.Data memory version, Version.Data[] memory supportedVersions)
        internal
        pure
        returns (Version.Data memory supportedVersion, bool found)
    {
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            supportedVersion = supportedVersions[i];
            if (
                keccak256(abi.encodePacked(supportedVersion.identifier))
                    == keccak256(abi.encodePacked(version.identifier))
            ) {
                return (supportedVersion, true);
            }
        }
        return (supportedVersion, false);
    }

    /**
     * @dev pickVersion iterates over the descending ordered set of compatible IBC
     * versions and selects the first version with a version identifier that is
     * supported by the counterparty. The returned version contains a feature
     * set with the intersection of the features supported by the source and
     * counterparty chains. If the feature set intersection is nil and this is
     * not allowed for the chosen version identifier then the search for a
     * compatible version continues. This function is called in the ConnOpenTry
     * handshake procedure.
     *
     * CONTRACT: pickVersion must only provide a version that is in the
     * intersection of the supported versions and the counterparty versions.
     */
    function pickVersion(Version.Data[] memory supportedVersions, Version.Data[] memory counterpartyVersions)
        internal
        pure
        returns (Version.Data memory)
    {
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            Version.Data memory supportedVersion = supportedVersions[i];
            (Version.Data memory counterpartyVersion, bool found) =
                findSupportedVersion(supportedVersion, counterpartyVersions);
            if (!found) {
                continue;
            }
            string[] memory featureSet =
                getFeatureSetIntersection(supportedVersion.features, counterpartyVersion.features);
            if (featureSet.length > 0) {
                return Version.Data({identifier: supportedVersion.identifier, features: featureSet});
            }
        }
        revert IIBCConnectionErrors.IBCConnectionNoMatchingVersionFound();
    }

    /**
     * @dev copyVersions copies `src` to `dst`
     */
    function copyVersions(Version.Data[] memory src, Version.Data[] storage dst) internal {
        uint256 srcLength = src.length;
        uint256 dstLength = dst.length;
        if (srcLength == dstLength) {
            for (uint256 i = 0; i < srcLength; i++) {
                copyVersion(src[i], dst[i]);
            }
        } else if (srcLength > dstLength) {
            for (uint256 i = 0; i < dstLength; i++) {
                copyVersion(src[i], dst[i]);
            }
            for (uint256 i = dstLength; i < srcLength; i++) {
                dst.push(src[i]);
            }
        } else {
            for (uint256 i = 0; i < srcLength; i++) {
                copyVersion(src[i], dst[i]);
            }
            for (uint256 i = srcLength; i < dstLength; i++) {
                dst.pop();
            }
        }
    }

    /**
     * @dev newVersions returns a new array with a given version
     */
    function newVersions(Version.Data memory version) internal pure returns (Version.Data[] memory ret) {
        ret = new Version.Data[](1);
        ret[0] = version;
    }

    /**
     * @dev verifySupportedFeature takes in a version and feature string and returns
     * true if the feature is supported by the version and false otherwise.
     */
    function verifySupportedFeature(Version.Data memory version, string memory feature) internal pure returns (bool) {
        bytes32 hashedFeature = keccak256(bytes(feature));
        for (uint256 i = 0; i < version.features.length; i++) {
            if (keccak256(bytes(version.features[i])) == hashedFeature) {
                return true;
            }
        }
        return false;
    }

    function getFeatureSetIntersection(string[] memory sourceFeatureSet, string[] memory counterpartyFeatureSet)
        private
        pure
        returns (string[] memory)
    {
        string[] memory featureSet = new string[](sourceFeatureSet.length);
        uint256 featureSetLength = 0;
        for (uint256 i = 0; i < sourceFeatureSet.length; i++) {
            if (contains(sourceFeatureSet[i], counterpartyFeatureSet)) {
                featureSet[featureSetLength] = sourceFeatureSet[i];
                featureSetLength++;
            }
        }
        string[] memory ret = new string[](featureSetLength);
        for (uint256 i = 0; i < featureSetLength; i++) {
            ret[i] = featureSet[i];
        }
        return ret;
    }

    function copyVersion(Version.Data memory src, Version.Data storage dst) private {
        dst.identifier = src.identifier;
        uint256 srcLength = src.features.length;
        uint256 dstLength = dst.features.length;

        if (srcLength == dstLength) {
            for (uint256 i = 0; i < srcLength; i++) {
                dst.features[i] = src.features[i];
            }
        } else if (srcLength > dstLength) {
            for (uint256 i = 0; i < dstLength; i++) {
                dst.features[i] = src.features[i];
            }
            for (uint256 i = dstLength; i < srcLength; i++) {
                dst.features.push(src.features[i]);
            }
        } else {
            for (uint256 i = 0; i < srcLength; i++) {
                dst.features[i] = src.features[i];
            }
            for (uint256 i = srcLength; i < dstLength; i++) {
                dst.features.pop();
            }
        }
    }

    function contains(string memory elem, string[] memory set) private pure returns (bool) {
        bytes32 hashedElem = keccak256(bytes(elem));
        for (uint256 i = 0; i < set.length; i++) {
            if (keccak256(bytes(set[i])) == hashedElem) {
                return true;
            }
        }
        return false;
    }
}
