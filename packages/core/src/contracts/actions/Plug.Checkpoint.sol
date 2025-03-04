// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.23;

/**
 * @title Plug Checkpoint
 * @notice State checkpointing and context management for complex multi-step
 *         operations that may span across multiple transactions.
 * @author 🔌 Plug <hello@onplug.io> (https://onplug.io)
 * @author 🟠 CHANCE <chance@onplug.io> (https://onplug.io)
 */
contract PlugCheckpoint {
    mapping(address owner => mapping(bytes32 id => bytes32 checkpointData)) private checkpoints;
    mapping(address owner => mapping(bytes32 id => uint256 version )) private versions;
    mapping(address owner => mapping(bytes32 contextId => bytes contextData)) private contexts;

    /**
     * @notice Save a checkpoint with a specific identifier
     * @param id The identifier for this checkpoint
     * @param data The checkpoint data to store
     * @return version The version number (incremented each time)
     */
    function saveCheckpoint(bytes32 id, bytes32 data) external returns (uint256 version) {
        checkpoints[msg.sender][id] = data;
        return version = ++versions[msg.sender][id];
    }

    /**
     * @notice Load a checkpoint by its identifier
     * @param owner The address that owns the checkpoint
     * @param id The identifier for the checkpoint
     * @return data The checkpoint data
     * @return version The current version number
     */
    function loadCheckpoint(address owner, bytes32 id) external view returns (bytes32 data, uint256 version) {
        return (checkpoints[owner][id], versions[owner][id]);
    }

    /**
     * @notice Check if a checkpoint exists and has been initialized
     * @param owner The address that owns the checkpoint
     * @param id The identifier for the checkpoint
     * @return exists Whether the checkpoint exists (has a version number > 0)
     */
    function hasCheckpoint(address owner, bytes32 id) external view returns (bool exists) {
        return versions[owner][id] > 0;
    }

    /**
     * @notice Reset a checkpoint (clear data and reset version)
     * @param id The identifier for the checkpoint to reset
     */
    function resetCheckpoint(bytes32 id) external {
        delete checkpoints[msg.sender][id];
        delete versions[msg.sender][id];
    }

    /**
     * @notice Save an execution context (supports arbitrary binary data)
     * @param id The identifier for this context
     * @param data The context data to store
     */
    function saveContext(bytes32 id, bytes calldata data) external {
        contexts[msg.sender][id] = data;
    }

    /**
     * @notice Load an execution context by its identifier
     * @param owner The address that owns the context
     * @param id The identifier for the context
     * @return data The context data
     */
    function loadContext(address owner, bytes32 id) external view returns (bytes memory data) {
        return contexts[owner][id];
    }

    /**
     * @notice Clear an execution context
     * @param id The identifier for the context to clear
     */
    function clearContext(bytes32 id) external {
        delete contexts[msg.sender][id];
    }

    /**
     * @notice Execute with a checkpoint - saves state before and after execution
     * @param id The checkpoint identifier
     * @param preState The state before execution
     * @param target The contract to call
     * @param data The calldata to execute
     * @return success Whether the execution was successful
     * @return result The result of the call
     * @return preStateVersion The version of the pre-execution state
     * @return postStateVersion The version of the post-execution state
     */
    function executeWithCheckpoint(
        bytes32 id,
        bytes32 preState,
        address target,
        bytes calldata data
    ) external returns (bool success, bytes memory result, uint256 preStateVersion, uint256 postStateVersion) {
        // Save pre-execution state
        preStateVersion = ++versions[msg.sender][id];
        checkpoints[msg.sender][id] = preState;

        // Execute the call
        (success, result) = target.call(data);

        // Generate and save post-execution state (we're using success flag as part of the state)
        bytes32 postState = bytes32(abi.encodePacked(success, result.length > 0 ? keccak256(result) : bytes32(0)));
        checkpoints[msg.sender][id] = postState;
        postStateVersion = ++versions[msg.sender][id];

        return (success, result, preStateVersion, postStateVersion);
    }

    /**
     * @notice Resume execution if checkpoint matches expected state
     * @param id The checkpoint identifier
     * @param expectedState The expected current state
     * @param expectedVersion The expected current version
     * @param target The contract to call if state matches
     * @param data The calldata to execute if state matches
     * @return matched Whether the checkpoint matched the expected state
     * @return success Whether the execution was successful (if matched)
     * @return result The result of the call (if matched and successful)
     */
    function resumeIfMatch(
        bytes32 id,
        bytes32 expectedState,
        uint256 expectedVersion,
        address target,
        bytes calldata data
    ) external view returns (bool matched, bool success, bytes memory result) {
        // Check if the current state and version match expectations
        if (checkpoints[msg.sender][id] != expectedState || versions[msg.sender][id] != expectedVersion) {
            return (false, false, new bytes(0));
        }

        // Make a static call to prevent state changes (this is just for checking)
        matched = true;
        (success, result) = target.staticcall(data);

        return (matched, success, result);
    }
}
