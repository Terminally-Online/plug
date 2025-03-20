import { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import { RopeStateManager, RopeState, KnotId, KnotData, ProcessedKnot } from '../models/rope-state-manager';

/**
 * Configuration for initializing the useRope hook
 */
export interface RopeConfig {
  /** For single knot mode: the sentence template */
  initialSentence?: string;
  /** For single knot mode: initial values for inputs */
  initialValues?: Record<string, string | undefined>;
  /** For multi-knot mode: array of initial knots */
  initialKnots?: Array<KnotData>;
  /** Optional context identifier for using multiple Rope instances */
  contextId?: string;
}

// Map to store RopeStateManager instances for context sharing
const ropeManagerInstances = new Map<string, RopeStateManager>();

/**
 * Creates or retrieves a RopeStateManager instance
 * Allows sharing state across multiple hooks with the same contextId
 */
function getRopeManager(contextId?: string): RopeStateManager {
  if (!contextId) {
    return new RopeStateManager();
  }
  
  // Reuse existing manager if one exists for this context
  if (!ropeManagerInstances.has(contextId)) {
    ropeManagerInstances.set(contextId, new RopeStateManager());
  }
  
  return ropeManagerInstances.get(contextId)!;
}

/**
 * Custom equality function for memoizing objects
 */
function shallowEqual(objA: any, objB: any): boolean {
  if (objA === objB) {
    return true;
  }
  
  if (typeof objA !== 'object' || objA === null || 
      typeof objB !== 'object' || objB === null) {
    return false;
  }
  
  const keysA = Object.keys(objA);
  const keysB = Object.keys(objB);
  
  if (keysA.length !== keysB.length) {
    return false;
  }
  
  for (const key of keysA) {
    if (!objB.hasOwnProperty(key) || objA[key] !== objB[key]) {
      return false;
    }
  }
  
  return true;
}

/**
 * React hook for working with a Rope
 * 
 * This hook provides a complete interface for working with Rope state
 * including actions for manipulating knots and helper functions for
 * querying the state.
 * 
 * @param config - Configuration options for initializing the rope
 * @returns An object containing state, actions, and helper functions
 * 
 * @example
 * ```tsx
 * const { state, actions, helpers } = useRope({
 *   initialKnots: [
 *     { id: 'knot1', protocol: 'uniswap', action: 'swap', sentence: 'Swap {0} for {1}' }
 *   ]
 * });
 * 
 * // Add a new knot
 * actions.addKnot({
 *   id: 'knot2',
 *   protocol: 'aave',
 *   action: 'deposit',
 *   sentence: 'Deposit {0} into Aave'
 * });
 * 
 * // Set a value in a knot
 * actions.setValue('knot1', 0, 'ETH');
 * ```
 */
export function useRope(config: RopeConfig = {}) {
  // Get or create rope manager instance (for context sharing)
  const ropeManager = useMemo(
    () => getRopeManager(config.contextId),
    [config.contextId]
  );
  
  // State to track the manager's state for React
  const [state, setState] = useState<RopeState>(() => ropeManager.getState());
  
  // Keep track of previous config for initialization
  const initializedRef = useRef(false);
  
  // Subscribe to state changes and initialize
  useEffect(() => {
    // Only initialize once
    if (!initializedRef.current) {
      // Initialize with config
      if (config.initialKnots && config.initialKnots.length > 0) {
        // Multi-knot mode - add each knot in sequence
        config.initialKnots.forEach(knot => {
          ropeManager.addKnot(knot);
        });
      } else if (config.initialSentence) {
        // Single knot mode - normalize values
        const values: Record<string, string> = {};
        
        if (config.initialValues) {
          Object.entries(config.initialValues).forEach(([key, value]) => {
            if (value !== undefined) {
              values[key] = value;
            }
          });
        }
        
        // Add the single knot
        ropeManager.addKnot({
          id: 'single-knot',
          protocol: 'default',
          action: 'default',
          sentence: config.initialSentence,
          values
        });
      }
      
      initializedRef.current = true;
    }
    
    // Subscribe to changes in the rope state manager
    const unsubscribe = ropeManager.subscribe(setState);
    
    // Clean up subscription when unmounting
    return unsubscribe;
  }, [ropeManager, config.initialKnots, config.initialSentence, config.initialValues]);
  
  // Create stable action functions with proper dependency tracking
  const setValue = useCallback(
    (knotId: KnotId, inputIndex: number, value: string | undefined) => {
      ropeManager.setValue(knotId, inputIndex, value);
    },
    [ropeManager]
  );
  
  const setActiveKnotValue = useCallback(
    (inputIndex: number, value: string | undefined) => {
      const activeKnotId = ropeManager.getState().activeKnotId;
      if (activeKnotId) {
        ropeManager.setValue(activeKnotId, inputIndex, value);
      }
    },
    [ropeManager]
  );
  
  const addKnot = useCallback(
    (knot: KnotData) => {
      ropeManager.addKnot(knot);
    },
    [ropeManager]
  );
  
  const removeKnot = useCallback(
    (id: KnotId) => {
      ropeManager.removeKnot(id);
    },
    [ropeManager]
  );
  
  const updateKnotSentence = useCallback(
    (id: KnotId, sentence: string) => {
      ropeManager.updateKnotSentence(id, sentence);
    },
    [ropeManager]
  );
  
  const moveKnot = useCallback(
    (id: KnotId, direction: 'up' | 'down') => {
      ropeManager.moveKnot(id, direction);
    },
    [ropeManager]
  );
  
  const setActiveKnot = useCallback(
    (id: KnotId) => {
      ropeManager.setActiveKnot(id);
    },
    [ropeManager]
  );
  
  const clearRope = useCallback(
    () => {
      ropeManager.clearRope();
    },
    [ropeManager]
  );
  
  const resetKnot = useCallback(
    (id: KnotId) => {
      ropeManager.resetKnot(id);
    },
    [ropeManager]
  );
  
  // Memoize the actions object to prevent unnecessary re-renders
  const actions = useMemo(
    () => ({
      setValue,
      setActiveKnotValue,
      addKnot,
      removeKnot,
      updateKnotSentence,
      moveKnot,
      setActiveKnot,
      clearRope,
      resetKnot
    }),
    [
      setValue,
      setActiveKnotValue,
      addKnot,
      removeKnot,
      updateKnotSentence,
      moveKnot,
      setActiveKnot,
      clearRope,
      resetKnot
    ]
  );
  
  // Helper functions that derive data from state
  // These are memoized individually to avoid unnecessary recalculations
  
  const getActiveKnot = useCallback(
    () => {
      if (!state.activeKnotId) {
        return state.head ? ropeManager.getKnot(state.head) : null;
      }
      return ropeManager.getKnot(state.activeKnotId);
    },
    [ropeManager, state.activeKnotId, state.head]
  );
  
  const getKnot = useCallback(
    (id: KnotId) => ropeManager.getKnot(id),
    [ropeManager]
  );
  
  const getKnotsArray = useCallback(
    () => ropeManager.getKnotsArray(),
    [ropeManager]
  );
  
  const getKnotValues = useCallback(
    (id: KnotId) => {
      const knot = ropeManager.getKnot(id);
      if (!knot) return null;
      
      return Object.fromEntries(
        Array.from(knot.values.entries()).map(([k, v]) => [k, v.value])
      );
    },
    [ropeManager]
  );
  
  const getInputName = useCallback(
    (knotId: KnotId, index: number) => {
      const knot = ropeManager.getKnot(knotId);
      return knot?.parsed?.inputs.find(i => i.index === index)?.name;
    },
    [ropeManager]
  );
  
  const getInputValue = useCallback(
    (knotId: KnotId, index: number) => {
      const knot = ropeManager.getKnot(knotId);
      return knot?.values.get(index);
    },
    [ropeManager]
  );
  
  const getInputError = useCallback(
    (knotId: KnotId, index: number) => {
      const knot = ropeManager.getKnot(knotId);
      return knot?.validationErrors.get(index);
    },
    [ropeManager]
  );
  
  const getDependentInputs = useCallback(
    (knotId: KnotId, index: number) => {
      const knot = ropeManager.getKnot(knotId);
      if (!knot?.parsed) return [];
      return knot.parsed.inputs.filter(input => input.dependentOn === index);
    },
    [ropeManager]
  );
  
  const hasDependency = useCallback(
    (knotId: KnotId, index: number) => {
      const knot = ropeManager.getKnot(knotId);
      if (!knot?.parsed) return false;
      const input = knot.parsed.inputs.find(i => i.index === index);
      return input?.dependentOn !== undefined;
    },
    [ropeManager]
  );
  
  // For helpers that depend on the state, we use memoization
  // with proper dependency tracking
  
  const getAllValidationErrors = useMemo(
    () => {
      const errors = Object.values(state.knots).reduce((acc, knot) => {
        if (knot.validationErrors.size === 0) return acc;
        
        acc[knot.id] = Object.fromEntries(
          Array.from(knot.validationErrors.entries()).map(([k, v]) => [k, v])
        );
        return acc;
      }, {} as Record<string, Record<string, { type: string, message: string }>>);
      
      return errors;
    },
    [state.knots]
  );
  
  const exportRope = useMemo(
    () => {
      return {
        knots: Object.values(state.knots).map(knot => {
          const values = Object.fromEntries(
            Array.from(knot.values.entries()).map(([k, v]) => [k, v.value])
          );
          
          return {
            id: knot.id,
            protocol: knot.protocol,
            action: knot.action,
            values
          };
        })
      };
    },
    [state.knots]
  );
  
  const wasKnotUpdated = useCallback(
    (knotId: KnotId) => state.lastUpdated === knotId,
    [state.lastUpdated]
  );
  
  const isSingleKnot = useMemo(
    () => Object.keys(state.knots).length === 1,
    [state.knots]
  );
  
  // Memoize the entire helpers object to prevent unnecessary re-renders
  const helpers = useMemo(
    () => ({
      getActiveKnot,
      getKnot,
      getKnotsArray,
      getKnotValues,
      getInputName,
      getInputValue,
      getInputError,
      getDependentInputs,
      hasDependency,
      getAllValidationErrors,
      exportRope,
      wasKnotUpdated,
      isSingleKnot
    }),
    [
      getActiveKnot,
      getKnot,
      getKnotsArray,
      getKnotValues,
      getInputName,
      getInputValue,
      getInputError,
      getDependentInputs,
      hasDependency,
      getAllValidationErrors,
      exportRope,
      wasKnotUpdated,
      isSingleKnot
    ]
  );
  
  // Return stable references that won't change unless their contents change
  return useMemo(
    () => ({
      state,
      actions,
      helpers
    }),
    [state, actions, helpers]
  );
}

// We use ropeManagerInstances (defined at the top of the file)
// for sharing state between hook instances

/**
 * Hook for efficiently working with a specific knot from a rope
 * 
 * This hook is optimized to only re-render when the relevant knot or
 * its dependencies change. It's ideal for components that need to
 * render a single knot in a larger rope.
 * 
 * @param knotId - The ID of the knot to access
 * @param contextId - Optional context ID to connect to a specific rope
 * @returns The knot state and knot-specific methods
 * 
 * @example
 * ```tsx
 * function KnotComponent({ knotId }) {
 *   const { knot, actions } = useKnot(knotId);
 *   
 *   if (!knot) return null;
 *   
 *   return (
 *     <div>
 *       <p>{knot.resolvedSentence || knot.sentence}</p>
 *       <button onClick={() => actions.setValue(0, 'ETH')}>Set Token</button>
 *     </div>
 *   );
 * }
 * ```
 */
export function useKnot(knotId: KnotId, contextId?: string) {
  // Use the same rope manager instance as useRope to ensure state sharing
  const ropeManager = useMemo(
    () => getRopeManager(contextId),
    [contextId]
  );
  
  // Get the state directly from the manager
  const [state, setState] = useState<RopeState>(() => ropeManager.getState());
  
  // Subscribe to state changes
  useEffect(() => {
    const unsubscribe = ropeManager.subscribe(setState);
    return unsubscribe;
  }, [ropeManager]);
  
  // Create basic actions needed for this hook
  const actions = useMemo(() => ({
    setValue: (id: KnotId, index: number, value: string | undefined) => 
      ropeManager.setValue(id, index, value),
    removeKnot: (id: KnotId) => ropeManager.removeKnot(id),
    updateKnotSentence: (id: KnotId, sentence: string) => 
      ropeManager.updateKnotSentence(id, sentence),
    resetKnot: (id: KnotId) => ropeManager.resetKnot(id),
    moveKnot: (id: KnotId, direction: 'up' | 'down') => 
      ropeManager.moveKnot(id, direction)
  }), [ropeManager]);
  
  // Previous knot reference for equality comparison
  const prevKnotRef = useRef<ProcessedKnot | null>(null);
  
  // Extract just this knot's data with memoization
  const knot = useMemo(() => {
    const nextKnot = state.knots[knotId] || null;
    
    // Only update reference if the knot actually changed
    if (nextKnot !== null && prevKnotRef.current !== null) {
      // Check for deeper changes in the object
      if (shallowEqual(nextKnot, prevKnotRef.current)) {
        return prevKnotRef.current;
      }
    }
    
    prevKnotRef.current = nextKnot;
    return nextKnot;
  }, [state.knots, knotId]);
  
  // Check if this knot was the last one updated (for optimizations)
  const wasUpdated = state.lastUpdated === knotId;
  
  // Memoized flag for active status
  const isActive = useMemo(
    () => state.activeKnotId === knotId,
    [state.activeKnotId, knotId]
  );
  
  // Create stable action functions with proper dependency tracking
  const setValue = useCallback(
    (inputIndex: number, value: string | undefined) => {
      actions.setValue(knotId, inputIndex, value);
    },
    [actions, knotId]
  );
  
  const remove = useCallback(
    () => {
      actions.removeKnot(knotId);
    },
    [actions, knotId]
  );
  
  const updateSentence = useCallback(
    (sentence: string) => {
      actions.updateKnotSentence(knotId, sentence);
    },
    [actions, knotId]
  );
  
  const reset = useCallback(
    () => {
      actions.resetKnot(knotId);
    },
    [actions, knotId]
  );
  
  const moveUp = useCallback(
    () => {
      actions.moveKnot(knotId, 'up');
    },
    [actions, knotId]
  );
  
  const moveDown = useCallback(
    () => {
      actions.moveKnot(knotId, 'down');
    },
    [actions, knotId]
  );
  
  // Memoize the actions object
  const knotActions = useMemo(
    () => ({
      setValue,
      remove,
      updateSentence,
      reset,
      moveUp,
      moveDown
    }),
    [setValue, remove, updateSentence, reset, moveUp, moveDown]
  );
  
  // Return a memoized result to prevent unnecessary re-renders
  return useMemo(
    () => ({
      knot,
      actions: knotActions,
      wasUpdated,
      isActive
    }),
    [knot, knotActions, wasUpdated, isActive]
  );
}

/**
 * Hook for working with a single knot in a simpler interface
 * 
 * This hook provides a simplified API focused on the template sentence pattern
 * rather than the full rope structure. It's ideal for simple form-like interactions
 * where a single input/action pattern is sufficient.
 * 
 * @param params - Configuration options and initial values
 * @returns State and methods for working with the single knot
 * 
 * @example
 * Example usage:
 * 
 * ```
 * function SwapForm() {
 *   const { sentence, values, setValues } = useSingleKnot({
 *     initialSentence: "Swap {0} for {1} on {2}",
 *     initialValues: { "0": "ETH" }
 *   });
 *   
 *   return (
 *     <div>
 *       <p>{sentence}</p>
 *       <input 
 *         value={values["0"] || ""} 
 *         onChange={e => setValues({ "0": e.target.value })} 
 *       />
 *     </div>
 *   );
 * }
 * ```
 */
export function useSingleKnot(params: {
  initialSentence: string;
  initialValues?: Record<string, string>;
  protocol?: string;
  action?: string;
  contextId?: string;
}) {
  // Use the standard useRope hook with single-knot configuration
  const { state, actions, helpers } = useRope({
    initialSentence: params.initialSentence,
    initialValues: params.initialValues,
    contextId: params.contextId
  });
  
  // Get the single knot (or null if not found)
  const knot = useMemo(() => {
    if (Object.keys(state.knots).length === 0) {
      return null;
    }
    return Object.values(state.knots)[0];
  }, [state.knots]);
  
  // Generate a simplified API focused on the sentence pattern rather than the knot
  const sentence = knot?.sentence || params.initialSentence;
  const resolvedSentence = knot?.resolvedSentence || null;
  
  // Convert the values Map to a more developer-friendly object
  const values = useMemo(() => {
    if (!knot) return {};
    
    return Object.fromEntries(
      Array.from(knot.values.entries()).map(([k, v]) => [k, v.value])
    );
  }, [knot]);
  
  // Function to set values in a more convenient way
  const setValues = useCallback((newValues: Record<string, string | undefined>) => {
    if (!knot) return;
    
    // Process each value
    Object.entries(newValues).forEach(([indexStr, value]) => {
      const index = Number(indexStr);
      if (!isNaN(index)) {
        actions.setValue('single-knot', index, value);
      }
    });
  }, [knot, actions]);
  
  // Extract validation errors into a simpler structure
  const errors = useMemo(() => {
    if (!knot) return {};
    
    return Object.fromEntries(
      Array.from(knot.validationErrors.entries()).map(([k, v]) => [k, v.message])
    );
  }, [knot]);
  
  // Return a convenient, simplified API
  return {
    // Template and output
    sentence,
    resolvedSentence,
    
    // Values and manipulation
    values,
    setValues,
    
    // Status flags
    isComplete: knot?.isComplete || false,
    isValid: knot?.isValid || false,
    
    // Errors
    errors,
    
    // Access to the underlying knot if needed
    knot
  };
}

// Re-export types from rope-state-manager for better DX
export type { RopeState, KnotData, ProcessedKnot } from '../models/rope-state-manager';
