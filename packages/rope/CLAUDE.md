# CLAUDE.md - Configuration for the Rope Package

## Build and Test Commands
```bash
# Install dependencies
pnpm install

# Build the package
pnpm build

# Watch mode for development
pnpm dev

# Run all tests
pnpm test

# Run a single test file
pnpm test -- __tests__/path/to/file.test.ts

# Run tests in watch mode
pnpm test:watch

# Lint code
pnpm lint

# Fix linting issues
pnpm lint:fix
```

## Code Style Guidelines
- **TypeScript**: Use proper typing for all variables, parameters, and return types
- **Classes**: Use PascalCase for class names (e.g., `Rope`, `Knot`, `PlugAPI`)
- **Variables/Properties**: Use camelCase for variables and properties
- **Imports**: Group imports by external libraries first, then internal modules
- **Quotes**: Use single quotes for strings
- **Semicolons**: Always use semicolons
- **Error Handling**: Use typed error objects and proper try/catch blocks
- **Documentation**: Document public methods with JSDoc comments
- **Line Length**: Keep lines under 100 characters
- **Indentation**: Use 2 spaces for indentation
- **Accessibility**: Use explicit access modifiers (private, public, protected)
- **Arrow Functions**: Avoid parentheses for single parameters

## Rope Package Development Todo

### Schema Management
- Implement schema caching with TTL to minimize API calls
- Support chain-specific schema filtering
- Add schema versioning for API changes
- Improve schema type definitions for better type safety

### Coil System
- Formalize coil type system and validation
- Add helper methods for working with coils
- Implement validation for coil references between actions

### Intent Building
- Enhance intent building with validation
- Support batched intent creation
- Add transaction simulation before execution
- Improve error handling with specific error types

### Protocol & Action Discovery
- Add search and filter utilities for actions
- Implement protocol metadata helpers
- Track popular/recent actions
- Support protocol documentation

### Testing & Documentation
- Expand test coverage for all components
- Add integration examples
- Generate documentation from schemas
- Improve TypeScript type definitions

## Cord Independence Requirements

To ensure Rope functions independently without requiring Cord, the following must be implemented:

### Sentence Parsing and Rendering
- Rope's `SentenceService` must fully handle all sentence template parsing that Cord currently manages
- Enhance `parseSentence` to support all input types and dependent inputs patterns
- Implement a UI wrapper that uses Rope's sentence parsing instead of Cord's `parseCordSentence`

### Coil System Integration
- `CoilService` must handle all aspects of coil references that Cord manages
- Implement UI components that work with Rope's coil system
- Ensure complete compatibility with coil references like `<-{coilName@index}`

### Input Management
- Replace Cord's input state management with Rope's equivalent
- Implement client-side validation using Rope's built-in validation system
- Create React hooks to replace `useCord` with Rope-compatible versions
  - This may include a `useRope` hook for managing sentence parsing and input state

### UI Components
- Create replacements for any UI components that depend on Cord
- Particularly focus on the sentence part rendering (currently in `part.tsx`)
- Maintain the same UX patterns but with Rope's data structures

### Required Conversions
1. Cord's `CordState` ⟷ Rope's own state management
2. Cord's `parseCordSentence` ⟷ Rope's `SentenceService.parseSentence`
3. Cord's `resolveSentence` ⟷ Rope's string formatting equivalents
4. Cord's `shouldRenderInput` ⟷ Equivalent logic in Rope

### Migration Strategy
- Create drop-in replacements for Cord functions
- Implement a compatibility layer if needed during transition
- Gradually migrate components from Cord to Rope dependencies
- Write comprehensive tests to ensure Rope matches Cord's functionality exactly

## Implementation Plan for Cord Independence

### 1. Enhance SentenceService
```typescript
// Extend SentenceInput to match Cord's InputReference
export interface SentenceInput {
  index: number;
  name: string;
  type?: string;
  defaultValue?: string;
  required?: boolean;
  dependentOn?: number;
  // Add any other properties from Cord's InputReference
}

// Add shouldRenderInput function similar to Cord
public shouldRenderInput(
  inputType: string, 
  allInputs: SentenceInput[], 
  getValueFn: (index: number) => { value: string } | undefined
): boolean {
  // Implement logic to determine if an input should be rendered
  // based on its type and dependencies
  
  // If input has a dependency, check if the dependency has a value
  const dependentInput = allInputs.find(input => input.type === inputType);
  if (dependentInput?.dependentOn !== undefined) {
    const dependencyValue = getValueFn(dependentInput.dependentOn);
    if (!dependencyValue || !dependencyValue.value) {
      return false;
    }
  }
  
  return true;
}

// Enhance the parseSentence method to extract more metadata
public parseSentence(sentence: string): ParsedSentence {
  // Enhanced implementation to support all Cord features
  // Specifically, we need to better handle dependency patterns:
  // {0} - Regular input
  // {0=>1} - Input 0 depends on input 1
  
  if (!sentence) {
    return { original: '', parts: [], inputs: [] };
  }
  
  const parts: string[] = [];
  const inputs: SentenceInput[] = [];
  const dependencyRegex = /\{(\d+)(?:=>(\d+))?\}/g;
  
  let lastIndex = 0;
  let match: RegExpExecArray | null;
  
  // Reset the regex to start from the beginning
  dependencyRegex.lastIndex = 0;
  
  while ((match = dependencyRegex.exec(sentence)) !== null) {
    // Add text before this match as a part
    if (match.index > lastIndex) {
      parts.push(sentence.substring(lastIndex, match.index));
    }
    
    // Add the placeholder as a part
    parts.push(match[0]);
    
    // Extract input data
    const inputIndex = parseInt(match[1]);
    const dependentOn = match[2] ? parseInt(match[2]) : undefined;
    
    // Add input if it doesn't exist yet
    if (!inputs.some(i => i.index === inputIndex)) {
      inputs.push({
        index: inputIndex,
        name: `input${inputIndex}`,
        type: 'string', // Default to string type
        required: true,
        dependentOn
      });
    }
    
    // Update last index to continue from end of match
    lastIndex = match.index + match[0].length;
  }
  
  // Add any remaining text
  if (lastIndex < sentence.length) {
    parts.push(sentence.substring(lastIndex));
  }
  
  return {
    original: sentence,
    parts,
    inputs: inputs.sort((a, b) => a.index - b.index)
  };
}

// Add input validation similar to Cord
public validateInput(value: string, type: string): { success: boolean, error?: string } {
  // Implement type-specific validation
  if (!value && type !== 'boolean') {
    return { success: false, error: 'Value is required' };
  }
  
  switch (type) {
    case 'number':
    case 'integer':
    case 'float':
      if (isNaN(Number(value))) {
        return { success: false, error: 'Value must be a number' };
      }
      if (type === 'integer' && !Number.isInteger(Number(value))) {
        return { success: false, error: 'Value must be an integer' };
      }
      break;
      
    case 'address':
      // Simple Ethereum address validation
      if (!/^0x[a-fA-F0-9]{40}$/.test(value)) {
        return { success: false, error: 'Invalid address format' };
      }
      break;
      
    case 'boolean':
      if (value !== 'true' && value !== 'false' && value !== '' && value !== undefined) {
        return { success: false, error: 'Value must be true or false' };
      }
      break;
  }
  
  return { success: true };
}
```

### 2. Create RopeState Interface and Management
```typescript
export interface RopeState {
  values: Map<number, { value: string }>;
  parsed: ParsedSentence | null;
  resolvedSentence: string | null;
  error: { type: string, message: string } | null;
  validationErrors: Map<number, { type: string, message: string }>;
}

// Function to handle value setting with dependency management
export function setValue(params: {
  parsedSentence: ParsedSentence;
  currentValues: Map<number, { value: string }>;
  index: number;
  value: string;
}): { success: boolean; value: Map<number, { value: string }>; error?: string } {
  const { parsedSentence, currentValues, index, value } = params;
  
  // Find the input definition
  const input = parsedSentence.inputs.find(i => i.index === index);
  if (!input) {
    return { success: false, error: `Input with index ${index} not found`, value: currentValues };
  }
  
  // Validate the input value based on its type
  const sentenceService = new SentenceService();
  const validation = sentenceService.validateInput(value, input.type || 'string');
  if (!validation.success) {
    return { success: false, error: validation.error, value: currentValues };
  }
  
  // Create a new values map with the updated value
  const newValues = new Map(currentValues);
  newValues.set(index, { value });
  
  // Clear values of dependent inputs
  // e.g., if input3 depends on input1 and we changed input1, clear input3
  parsedSentence.inputs.forEach(otherInput => {
    if (otherInput.dependentOn === index) {
      newValues.delete(otherInput.index);
    }
  });
  
  return { success: true, value: newValues };
}
```

### 3. Implement useRope React Hook
```typescript
import { useCallback, useEffect, useMemo, useState } from 'react';

// Helper function to convert record to Map
const createStateFromValues = (values: Record<string, string | undefined>) => {
  const state = new Map<number, { value: string }>();
  Object.entries(values).forEach(([key, value]) => {
    if (value !== undefined) {
      state.set(Number(key), { value });
    }
  });
  return state;
};

export function useRope(sentence: string, values: Record<string, string | undefined>) {
  const [state, setState] = useState<RopeState>(() => ({
    values: createStateFromValues(values),
    parsed: null,
    resolvedSentence: null,
    error: null,
    validationErrors: new Map()
  }));

  // Create a memoized sentenceService instance
  const sentenceService = useMemo(() => new SentenceService(), []);

  // Parse sentence on change
  const parsed = useMemo(() => {
    try {
      const result = sentenceService.parseSentence(sentence);
      setState(prev => ({ ...prev, error: null }));
      return result;
    } catch (error) {
      setState(prev => ({
        ...prev,
        error: { type: 'parse', message: error.message },
        parsed: null
      }));
      return null;
    }
  }, [sentence, sentenceService]);

  // Filter inputs that should be rendered
  const filteredInputs = useMemo(() => {
    if (!parsed) return [];
    return parsed.inputs.filter(input =>
      sentenceService.shouldRenderInput(
        input.type || 'string', 
        parsed.inputs, 
        index => state.values.get(index)
      )
    );
  }, [parsed, state.values, sentenceService]);

  // Create a version of parsed with filtered inputs
  const parsedWithFilteredInputs = useMemo(() => {
    if (!parsed) return null;
    return {
      ...parsed,
      inputs: filteredInputs
    };
  }, [parsed, filteredInputs]);

  // Split template into parts for rendering
  const parts = useMemo(() => parsedWithFilteredInputs
    ? parsedWithFilteredInputs.parts.map(part => {
        if (part.match(/\{[^}]+\}/)) return [part];
        return part.split(/(\s+)/g);
      }).flat()
    : [], [parsedWithFilteredInputs]);

  // Format sentence with values
  const resolvedSentence = useMemo(() => {
    if (!parsed) return null;

    const allInputsHaveValues = parsed.inputs.every(input => state.values.has(input.index));
    if (!allInputsHaveValues) return null;

    try {
      // Convert Map to plain object for formatSentence
      const valuesObj = Object.fromEntries(
        Array.from(state.values.entries()).map(([k, v]) => [k, v.value])
      );
      const result = sentenceService.formatSentence(parsed, valuesObj);
      return result;
    } catch (error) {
      setState(prev => ({
        ...prev,
        error: { type: 'resolution', message: error.message }
      }));
      return null;
    }
  }, [parsed, state.values, sentenceService]);

  // Actions to update values
  const actions = {
    setValue: useCallback((index: number, value: string | undefined) => {
      if (!parsed) return;

      const result = setValue({
        parsedSentence: parsed,
        currentValues: state.values,
        index,
        value: value ?? ''
      });

      setState(prev => ({
        ...prev,
        values: result.value,
        validationErrors: result.error
          ? new Map(prev.validationErrors).set(index, {
              type: 'validation',
              message: result.error
            })
          : new Map([...prev.validationErrors].filter(([k]) => k !== index))
      }));
    }, [parsed, state.values])
  };

  // Helper functions
  const helpers = {
    getInputName: useCallback((index: number) => 
      parsed?.inputs.find(i => i.index === index)?.name, 
      [parsed]
    ),
    getInputValue: useCallback((index: number) => 
      state.values.get(index), 
      [state.values]
    ),
    getInputError: useCallback((index: number) => 
      state.validationErrors.get(index), 
      [state.validationErrors]
    ),
    getDependentInputs: useCallback((index: number) => {
      if (!parsed) return [];
      return parsed.inputs.filter(input => input.dependentOn === index);
    }, [parsed]),
    hasDependency: useCallback((index: number) => {
      if (!parsed) return false;
      return parsed.inputs.some(input => input.dependentOn === index);
    }, [parsed]),
    isComplete: useMemo(
      () => parsedWithFilteredInputs?.inputs.every(input => {
        const value = state.values.get(input.index);
        return value !== undefined;
      }) ?? false,
      [parsedWithFilteredInputs, state.values]
    ),
    isValid: useMemo(() => {
      if (!parsedWithFilteredInputs) return false;
      // Check for validation errors
      if (state.validationErrors.size > 0) return false;
      // Check that all values that exist are non-empty strings
      return !Array.from(state.values.values()).some(value => value?.value === '');
    }, [parsedWithFilteredInputs, state.validationErrors, state.values])
  };

  return {
    state: {
      ...state,
      parsed: parsedWithFilteredInputs,
      parts,
      resolvedSentence
    },
    actions,
    helpers
  };
}
```

### 4. Create Compatibility Layer
```typescript
// For easy migration, create functions that match Cord's API
export function parseCordSentence(sentence: string): { success: boolean, value?: ParsedSentence, error?: string } {
  try {
    const sentenceService = new SentenceService();
    const result = sentenceService.parseSentence(sentence);
    return { success: true, value: result };
  } catch (error) {
    return { success: false, error: error.message };
  }
}

export function resolveSentence(
  parsed: ParsedSentence, 
  values: Map<number, { value: string }>
): { success: boolean, value?: string, error?: string } {
  try {
    const sentenceService = new SentenceService();
    const valueMap = Object.fromEntries(
      Array.from(values.entries()).map(([k, v]) => [k, v.value])
    );
    const result = sentenceService.formatSentence(parsed, valueMap);
    return { success: true, value: result };
  } catch (error) {
    return { success: false, error: error.message };
  }
}

export function getInputPlaceholder(type?: string): string {
  const sentenceService = new SentenceService();
  return sentenceService.getPlaceholder(type);
}

export function shouldRenderInput(
  inputType: string,
  allInputs: SentenceInput[],
  getValueFn: (index: number) => { value: string } | undefined
): boolean {
  const sentenceService = new SentenceService();
  return sentenceService.shouldRenderInput(inputType, allInputs, getValueFn);
}

export function createInitialState(): Map<number, { value: string }> {
  return new Map();
}

// Export type aliases to match Cord
export type InputReference = SentenceInput;
export type CordState = RopeState;
```

### 5. Enhance Coil Integration
```typescript
// Extend CoilService to better integrate with SentenceService
public resolveCoilReferences(
  sentence: string, 
  values: Record<string, string>, 
  availableCoils: Record<string, string>
): Record<string, string> {
  // Create a new values object with resolved coil references
  const resolvedValues = { ...values };
  
  // For each value, check if it's a coil reference
  Object.entries(resolvedValues).forEach(([key, value]) => {
    if (typeof value === 'string' && this.isCoilReference(value)) {
      const coilRef = this.parseCoilReference(value);
      if (coilRef && availableCoils[coilRef.name]) {
        // For now, we just mark it as a valid reference
        // In a real implementation, we would resolve to actual values
        resolvedValues[key] = `Coil: ${coilRef.name}`;
      }
    }
  });
  
  return resolvedValues;
}

// Add coil validation with specific type checking
public validateCoilUsage(
  coilReferences: string[], 
  requiredTypes: Record<string, string>,
  availableCoils: Record<string, string>
): { valid: boolean, errors: Record<string, string> } {
  const errors: Record<string, string> = {};
  
  coilReferences.forEach(ref => {
    const coilRef = this.parseCoilReference(ref);
    if (coilRef) {
      const coilName = coilRef.name;
      const requiredType = requiredTypes[coilName];
      const availableType = availableCoils[coilName];
      
      if (!availableType) {
        errors[coilName] = `Coil "${coilName}" not found`;
      } else if (requiredType && !this.isCoilCompatible(availableType, requiredType)) {
        errors[coilName] = `Coil "${coilName}" of type "${availableType}" is not compatible with required type "${requiredType}"`;
      }
    }
  });
  
  return { 
    valid: Object.keys(errors).length === 0,
    errors
  };
}
```

## Performance Optimization Strategy

### 1. Better Memoization

```typescript
// Instead of recalculating everything in each useEffect, use more granular memoization:
const filteredInputs = useMemo(() => {
  if (!knot.parsed) return [];
  return knot.parsed.inputs.filter(input => 
    sentenceService.shouldRenderInput(
      input.type || 'string',
      knot.parsed?.inputs || [],
      index => knot.values.get(index)
    )
  );
}, [knot.parsed, knot.values]);

// Memoize at the individual knot level
const processedKnots = useMemo(() => {
  return state.knots.map(knot => processKnot(knot, state.allCoils));
}, [state.knots, state.allCoils]);

// Memoize overall state calculations
const { isComplete, isValid } = useMemo(() => ({
  isComplete: processedKnots.every(knot => knot.isComplete),
  isValid: processedKnots.every(knot => knot.isValid)
}), [processedKnots]);
```

### 2. Optimized State Updates

```typescript
// Use immer or a similar library for immutable updates
import produce from 'immer';

// Update only the specific knot that changed
const setValue = useCallback((knotId: string, inputIndex: number, value: string | undefined) => {
  setState(produce(draft => {
    const knotIndex = draft.knots.findIndex(k => k.id === knotId);
    if (knotIndex === -1) return;

    const knot = draft.knots[knotIndex];
    if (!knot.parsed) return;

    const result = sentenceService.setValue(
      knot.parsed,
      knot.values,
      inputIndex,
      value ?? ''
    );

    // Only update the specific knot's values
    knot.values = result.values;
    
    // Update validation errors efficiently
    if (result.error) {
      knot.validationErrors.set(inputIndex, {
        type: 'validation',
        message: result.error
      });
    } else {
      knot.validationErrors.delete(inputIndex);
    }
  }));
}, [sentenceService]);
```

### 3. Lazy Coil Evaluation

```typescript
// Instead of computing all coils immediately, use a proxy-based approach
const createCoilProxy = (knots) => {
  const cache = new Map();
  
  return new Proxy({}, {
    get: (target, prop) => {
      if (cache.has(prop)) {
        return cache.get(prop);
      }
      
      // Only compute the coil value when it's actually requested
      const value = computeCoilValue(prop.toString(), knots);
      cache.set(prop, value);
      return value;
    }
  });
};

// Then use it in sentence resolution
const resolvedSentence = formatSentenceWithCoils(knot.parsed, values, createCoilProxy(state.knots));
```

### 4. Sentence Parsing Optimization

```typescript
// Use a cache for parsed sentences
const parseSentenceCache = new WeakMap();

const getParsedSentence = (sentence: string) => {
  if (parseSentenceCache.has(sentence)) {
    return parseSentenceCache.get(sentence);
  }
  
  const parsed = sentenceService.parseSentence(sentence);
  parseSentenceCache.set(sentence, parsed);
  return parsed;
};

// Use this in updateKnotSentence
const updateKnotSentence = useCallback((id: string, sentence: string) => {
  try {
    const parsed = getParsedSentence(sentence);
    
    setState(prev => ({
      ...prev,
      knots: prev.knots.map(knot => 
        knot.id === id ? { ...knot, sentence, parsed } : knot
      )
    }));
  } catch (error) {
    // Handle error
  }
}, []);
```

### 5. Architectural Changes for Bundle Size

```typescript
// Split services into individual files for better tree-shaking
// src/services/sentence/index.ts
export { parseSentence } from './parseSentence';
export { formatSentence } from './formatSentence';
export { validateInput } from './validateInput';
// etc.

// Create specialized hooks for specific use cases
export const useSingleRope = (config) => {
  // Optimized version just for single knot usage
};

export const useMultiRope = (config) => {
  // Optimized version just for multi-knot usage
};

// Main hook uses either based on config
export const useRope = (config) => {
  if (config.initialKnots) {
    return useMultiRope(config);
  }
  return useSingleRope(config);
};
```

### 6. Performance-Critical Hooks

```typescript
// Add these hooks for performance-critical operations

// Hook to track if any knot has changed
const useKnotChanged = (knots) => {
  const prevKnotsRef = useRef(knots);
  
  const changed = useMemo(() => {
    const hasChanged = !isEqual(prevKnotsRef.current, knots);
    prevKnotsRef.current = knots;
    return hasChanged;
  }, [knots]);
  
  return changed;
};

// Use react-tracked or similar for performance optimization
import { createContainer } from 'react-tracked';

const { Provider, useTracked } = createContainer(() => useState(initialRopeState));

// This allows components to only re-render when the parts they use change
```

### 7. Streamlined Type System

```typescript
// Make types more concise to improve developer experience

// Instead of having separate interfaces, use composition
type KnotBase = {
  id: string;
  protocol: string;
  action: string;
  sentence: string;
};

type KnotInput = KnotBase & {
  values?: Record<string, string>;
};

type ProcessedKnot = KnotBase & {
  values: Map<number, { value: string }>;
  parsed: ParsedSentence | null;
  resolvedSentence: string | null;
  validationErrors: Map<number, { type: string; message: string }>;
  isComplete: boolean;
  isValid: boolean;
  parts?: string[];
};

// This makes the relationships clearer and reduces duplication
```

## Rope Architecture: Linked List Implementation

The Rope package has been re-architectured using a linked list approach for managing knots. This provides significant performance benefits, especially for ropes with multiple knots.

### Core Components

1. **RopeStateManager**: Central state management class
   - Maintains a linked list of knots
   - Provides O(1) lookups for knots by ID
   - Handles immutable state updates
   - Manages coil resolution between knots

2. **useRope Hook**: Main React integration
   - Subscribes to RopeStateManager updates
   - Exposes a clean API for manipulating ropes
   - Optimized to minimize unnecessary renders

3. **useKnot Hook**: Fine-grained knot access
   - Component-level hook that only updates when specific knot changes
   - Provides knot-specific actions
   - Tracks if knot was most recently updated

### Performance Features

- **Immutable Updates**: Only re-renders affected parts of the UI
- **Linked List Structure**: Efficient traversal and updates
- **O(1) Lookups**: Fast access to any knot by ID
- **Lazy Coil Resolution**: Only computes values when needed
- **Memoization**: Prevents unnecessary recalculations
- **Update Tracking**: Tracks which knot was last updated

### Usage Example

```tsx
// Top-level component using useRope
function RopeEditor() {
  const { state, actions, helpers } = useRope({
    initialKnots: [
      {
        id: 'knot-1',
        protocol: 'uniswap',
        action: 'swap',
        sentence: 'Swap {0} for {1}',
        values: { '0': 'ETH' }
      }
    ]
  });
  
  return (
    <div>
      {/* Render knots in order using the linked list */}
      {(() => {
        const elements = [];
        let currentId = state.head;
        
        while (currentId) {
          elements.push(
            <KnotComponent key={currentId} knotId={currentId} />
          );
          const knot = state.knots[currentId];
          currentId = knot?.next || null;
        }
        
        return elements;
      })()}
      
      <button onClick={() => actions.addKnot({
        id: `knot-${Date.now()}`,
        protocol: 'compound',
        action: 'supply',
        sentence: 'Supply {0} to {1}'
      })}>Add Knot</button>
    </div>
  );
}

// Knot component using useKnot (only re-renders when this specific knot changes)
function KnotComponent({ knotId }) {
  const { knot, actions, wasUpdated } = useKnot(knotId);
  
  if (!knot) return null;
  
  return (
    <div className={wasUpdated ? 'highlight' : ''}>
      <div>{knot.resolvedSentence || knot.sentence}</div>
      
      {knot.parsed?.inputs.map(input => (
        <input
          key={input.index}
          value={knot.values.get(input.index)?.value || ''}
          onChange={e => actions.setValue(input.index, e.target.value)}
        />
      ))}
      
      <button onClick={actions.moveUp}>Up</button>
      <button onClick={actions.moveDown}>Down</button>
      <button onClick={actions.remove}>Remove</button>
    </div>
  );
}
```

## Implementation Plan Status

✅ **Phase 1: Base Services**
- ✅ Added sentence parsing cache in SentenceService
- ✅ Implemented lazy coil evaluation in CoilService
- ✅ Refined validation for better performance

✅ **Phase 2: Enhanced State Management**
- ✅ Implemented linked list architecture
- ✅ Added efficient immutable state updates
- ✅ Added update tracking for performance

✅ **Phase 3: React Integration**
- ✅ Created useRope hook for managing complete ropes
- ✅ Added specialized useKnot hook for component-level integration
- ✅ Added helpers for traversing linked list structure

🔲 **Phase 4: Testing & Validation**
- 🔲 Add performance benchmarking tests
- 🔲 Compare with previous implementation
- 🔲 Add comprehensive test suite

The implementation can handle complex ropes with many knots while maintaining excellent performance characteristics. The linked list approach allows efficient traversal, manipulation, and coil resolution with minimal overhead.