# Rope Package Documentation

This directory contains examples and documentation for the Rope package.

## Examples

The `examples` directory contains example React components that demonstrate how to use the Rope package in a real-world application.

### Components

- `example.tsx`: Demonstrates using the individual `useRope` hook for working with sentences and inputs, including coil references.
- `topLevelExample.tsx`: Shows the recommended top-level approach using the `useRopeState` hook to manage all knots in a Rope from a central location.

## Usage Instructions

These examples are meant to be references and are not included in the main build. They demonstrate best practices for using the Rope package in a React application.

### React Integration

There are two primary hooks for integrating Rope with React applications:

#### 1. useRope Hook (Individual Sentence)

The `useRope` hook provides functionality for a single sentence. It includes:

1. Parsing a sentence template into parts and inputs
2. Managing input values and validation
3. Tracking dependencies between inputs
4. Resolving coil references
5. Formatting the final resolved sentence

#### 2. useRopeState Hook (Top-level Management)

The `useRopeState` hook provides centralized management for an entire Rope with multiple knots. This approach is recommended for most applications, as it:

1. Centralizes state management at the application level
2. Allows monitoring and validation of the entire Rope
3. Enables easy access to all available coils across knots
4. Provides actions for manipulating knots (add, remove, move)
5. Makes it simpler to build UI that reflects the entire Rope's state

See the examples in the `examples` directory for practical usage patterns:
- `example.tsx`: Shows using the individual `useRope` hook
- `topLevelExample.tsx`: Demonstrates the centralized `useRopeState` approach

## Cord Independence

The Rope package is designed to function independently without requiring Cord. It implements all the functionality that was previously handled by Cord, including:

1. Sentence parsing and rendering
2. Input state management and validation
3. Coil reference handling
4. UI component support

For migration from Cord to Rope, use the compatibility layer functions exported from the main package.