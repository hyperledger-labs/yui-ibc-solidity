// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.12;

interface IICS20Bank {
    /**
     * @dev balanceOf returns the balance of the account.
     * @param account account address
     * @param denom denom of the token
     */
    function balanceOf(address account, string calldata denom) external view returns (uint256);

    /**
     * @dev transferFrom transfers tokens from the sender to the receiver.
     * @param from sender address
     * @param to receiver address
     * @param denom denom of the token
     * @param amount amount of the token
     */
    function transferFrom(address from, address to, string calldata denom, uint256 amount) external;

    /**
     * @dev mint mints tokens to the account.
     * @param account account address
     * @param denom denom of the token
     * @param amount amount of the token
     */
    function mint(address account, string calldata denom, uint256 amount) external;

    /**
     * @dev burn burns tokens from the account.
     * @param account account address
     * @param denom denom of the token
     * @param amount amount of the token
     */
    function burn(address account, string calldata denom, uint256 amount) external;

    /**
     * @dev addressToDenom returns the denom of the token corresponding to the contract address.
     *      The denom must be a json escaped string.
     */
    function addressToDenom(address tokenContract) external pure returns (string memory);

    /**
     * @dev deposit deposits tokens to the bank.
     * @param tokenContract token contract address
     * @param amount amount of the token
     * @param receiver receiver address on the bank
     */
    function deposit(address tokenContract, uint256 amount, address receiver) external;

    /**
     * @dev withdraw withdraws tokens from the bank.
     * @param tokenContract token contract address
     * @param amount amount of the token
     * @param receiver receiver address on the bank
     */
    function withdraw(address tokenContract, uint256 amount, address receiver) external;
}
