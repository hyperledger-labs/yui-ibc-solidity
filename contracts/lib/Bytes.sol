pragma solidity ^0.6.8;

library Bytes {
    function toAddress(bytes memory bz) internal pure returns (address addr) {
        require(bz.length == 20, "cannot convert into address");
        assembly {
            addr := mload(add(bz, 20))
        }
    }

    function toBytes32(bytes memory bz) internal pure returns (bytes32 ret) {
        require(bz.length == 32, "cannot convert into address");
        assembly {
            ret := mload(add(bz, 32))
        }
    }
}
