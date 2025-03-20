# Rope Playground

A visual node-based editor for testing and demonstrating the Rope package functionality.

## Features

- Interactive node-based editor using ReactFlow
- Visual representation of knots and their connections
- Live editing of template values
- Real-time validation feedback
- Drag-and-drop interface for arranging nodes

## Getting Started

### Installation

The playground is already set up as part of the Rope package. If you need to reinstall dependencies:

```bash
# From the rope package directory
cd playground
pnpm install
```

### Running the Playground

```bash
# From the rope package directory
pnpm playground

# Or directly from the playground directory
pnpm dev
```

## Usage

The playground provides a visual interface for working with Rope:

1. **Node Management**:
   - Add new nodes using the "Add New Knot" button
   - Remove nodes with the "×" button in the sidebar
   - Drag nodes to reposition them on the canvas

2. **Editing Knots**:
   - Each node represents a knot with a sentence template
   - Fill in the input fields to provide values for the template placeholders
   - See the resolved sentence update in real-time

3. **Validation**:
   - Nodes are color-coded based on their validation status:
     - Yellow: Incomplete (missing required inputs)
     - Green: Valid and complete
     - Red: Invalid (validation errors)

4. **Connection**:
   - Connect nodes by dragging from the output handle of one node to the input handle of another
   - Create sequences of actions that would form a Rope in production

## How It Works

The playground uses the core Rope package to handle the business logic:

- `useRope` hook manages the state of all knots
- Sentence parsing handles the template placeholders
- Validation logic ensures inputs meet requirements
- The visual interface visualizes the internal state

## Development

To modify the playground:

1. Edit React components in `src/components/`
2. Modify styles in `src/index.css`
3. Update the context provider in `src/context/RopeContext.tsx`

Changes to the core Rope package will be reflected in the playground when you rebuild the package.