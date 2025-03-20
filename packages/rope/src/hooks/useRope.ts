import { useCallback, useEffect, useMemo, useState } from 'react';
import { RopeStateManager, RopeState, KnotId, KnotData, ProcessedKnot } from '../models/rope-state-manager';

/**
 * Configuration for initializing the useRope hook
 */
export interface RopeConfig {
  // For single knot mode
  initialSentence?: string;
  initialValues?: Record<string, string | undefined>;
  // For multi-knot mode
  initialKnots?: Array<KnotData>;
}

/**
 * React hook for working with a full Rope
 * Uses the RopeStateManager under the hood with linked list implementation
 * 
 * @param config Configuration options for the hook
 * @returns State, actions, and helper functions for working with the Rope
 */
export function useRope(config: RopeConfig = {}) {
  // Create a memoized rope state manager
  const ropeManager = useMemo(() => new RopeStateManager(), []);
  
  // State to track the manager's state for React
  const [state, setState] = useState<RopeState>(() => ropeManager.getState());
  
  // Subscribe to state changes from the manager
  useEffect(() => {
    // Subscribe to changes in the rope state manager
    const unsubscribe = ropeManager.subscribe(setState);
    
    // Initialize with config
    if (config.initialKnots && config.initialKnots.length > 0) {
      // Multi-knot mode
      config.initialKnots.forEach(knot => {
        ropeManager.addKnot(knot);
      });
    } else if (config.initialSentence) {
      // Single knot mode - convert initialValues to required format
      const values: Record<string, string> = {};
      if (config.initialValues) {
        Object.entries(config.initialValues).forEach(([key, value]) => {
          if (value !== undefined) {
            values[key] = value;
          }
        });
      }
      
      ropeManager.addKnot({
        id: 'single-knot',
        protocol: 'default',
        action: 'default',
        sentence: config.initialSentence,
        values
      });
    }
    
    // Clean up subscription when unmounting
    return unsubscribe;
  }, [ropeManager, config]);
  
  // Actions for manipulating the rope - direct proxy to ropeManager methods
  const actions = useMemo(() => ({
    setValue: useCallback((knotId: KnotId, inputIndex: number, value: string | undefined) => {
      ropeManager.setValue(knotId, inputIndex, value);
    }, [ropeManager]),
    
    setActiveKnotValue: useCallback((inputIndex: number, value: string | undefined) => {
      const activeKnotId = state.activeKnotId;
      if (activeKnotId) {
        ropeManager.setValue(activeKnotId, inputIndex, value);
      }
    }, [ropeManager, state.activeKnotId]),
    
    addKnot: useCallback((knot: KnotData) => {
      ropeManager.addKnot(knot);
    }, [ropeManager]),
    
    removeKnot: useCallback((id: KnotId) => {
      ropeManager.removeKnot(id);
    }, [ropeManager]),
    
    updateKnotSentence: useCallback((id: KnotId, sentence: string) => {
      ropeManager.updateKnotSentence(id, sentence);
    }, [ropeManager]),
    
    moveKnot: useCallback((id: KnotId, direction: 'up' | 'down') => {
      ropeManager.moveKnot(id, direction);
    }, [ropeManager]),
    
    setActiveKnot: useCallback((id: KnotId) => {
      ropeManager.setActiveKnot(id);
    }, [ropeManager]),
    
    clearRope: useCallback(() => {
      ropeManager.clearRope();
    }, [ropeManager]),
    
    resetKnot: useCallback((id: KnotId) => {
      ropeManager.resetKnot(id);
    }, [ropeManager])
  }), [ropeManager, state.activeKnotId]);
  
  // Helper functions for accessing state
  const helpers = useMemo(() => ({
    // Get the active knot
    getActiveKnot: useCallback(() => {
      if (!state.activeKnotId) {
        return state.head ? ropeManager.getKnot(state.head) : null;
      }
      return ropeManager.getKnot(state.activeKnotId);
    }, [ropeManager, state.activeKnotId, state.head]),
    
    // Get a specific knot by id - O(1) operation
    getKnot: useCallback((id: KnotId) => {
      return ropeManager.getKnot(id);
    }, [ropeManager]),
    
    // Get all knots as an array in order
    getKnotsArray: useCallback(() => {
      return ropeManager.getKnotsArray();
    }, [ropeManager]),
    
    // Get all values for a specific knot
    getKnotValues: useCallback((id: KnotId) => {
      const knot = ropeManager.getKnot(id);
      if (!knot) return null;
      
      return Object.fromEntries(
        Array.from(knot.values.entries()).map(([k, v]) => [k, v.value])
      );
    }, [ropeManager]),
    
    // Get an input's name for a knot
    getInputName: useCallback((knotId: KnotId, index: number) => {
      const knot = ropeManager.getKnot(knotId);
      return knot?.parsed?.inputs.find(i => i.index === index)?.name;
    }, [ropeManager]),
    
    // Get an input's value for a knot
    getInputValue: useCallback((knotId: KnotId, index: number) => {
      const knot = ropeManager.getKnot(knotId);
      return knot?.values.get(index);
    }, [ropeManager]),
    
    // Get an input's error for a knot
    getInputError: useCallback((knotId: KnotId, index: number) => {
      const knot = ropeManager.getKnot(knotId);
      return knot?.validationErrors.get(index);
    }, [ropeManager]),
    
    // Get dependent inputs for a specific input
    getDependentInputs: useCallback((knotId: KnotId, index: number) => {
      const knot = ropeManager.getKnot(knotId);
      if (!knot?.parsed) return [];
      return knot.parsed.inputs.filter(input => input.dependentOn === index);
    }, [ropeManager]),
    
    // Check if an input has a dependency
    hasDependency: useCallback((knotId: KnotId, index: number) => {
      const knot = ropeManager.getKnot(knotId);
      if (!knot?.parsed) return false;
      const input = knot.parsed.inputs.find(i => i.index === index);
      return input?.dependentOn !== undefined;
    }, [ropeManager]),
    
    // Get all validation errors across all knots
    getAllValidationErrors: useCallback(() => {
      return Object.values(state.knots).reduce((acc, knot) => {
        if (knot.validationErrors.size === 0) return acc;
        
        acc[knot.id] = Object.fromEntries(
          Array.from(knot.validationErrors.entries()).map(([k, v]) => [k, v])
        );
        return acc;
      }, {} as Record<string, Record<string, { type: string, message: string }>>);
    }, [state.knots]),
    
    // Export the entire rope as a structured object (for sending to API)
    exportRope: useCallback(() => {
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
    }, [state.knots]),
    
    // Check if a knot was recently updated (for optimization)
    wasKnotUpdated: useCallback((knotId: KnotId) => {
      return state.lastUpdated === knotId;
    }, [state.lastUpdated]),
    
    // Check if we're in single knot mode
    isSingleKnot: useMemo(() => 
      Object.keys(state.knots).length === 1, 
      [state.knots]
    )
  }), [ropeManager, state]);
  
  return {
    state,
    actions,
    helpers
  };
}

/**
 * Hook for efficiently working with a specific knot from a rope
 * Optimized to only re-render when the relevant knot changes
 * 
 * @param knotId The ID of the knot to access
 * @returns The knot state and knot-specific methods
 */
export function useKnot(knotId: KnotId) {
  const { state, actions } = useRope();
  
  // Extract just this knot's data
  const knot = useMemo(() => 
    state.knots[knotId] || null, 
    [state.knots, knotId]
  );
  
  // Check if this knot was the last one updated
  const wasUpdated = state.lastUpdated === knotId;
  
  // Knot-specific actions
  const knotActions = useMemo(() => ({
    setValue: useCallback((inputIndex: number, value: string | undefined) => {
      actions.setValue(knotId, inputIndex, value);
    }, [actions]),
    
    remove: useCallback(() => {
      actions.removeKnot(knotId);
    }, [actions]),
    
    updateSentence: useCallback((sentence: string) => {
      actions.updateKnotSentence(knotId, sentence);
    }, [actions]),
    
    reset: useCallback(() => {
      actions.resetKnot(knotId);
    }, [actions]),
    
    moveUp: useCallback(() => {
      actions.moveKnot(knotId, 'up');
    }, [actions]),
    
    moveDown: useCallback(() => {
      actions.moveKnot(knotId, 'down');
    }, [actions])
  }), [actions, knotId]);
  
  return {
    knot,
    actions: knotActions,
    wasUpdated,
    isActive: state.activeKnotId === knotId
  };
}

// Re-export types from rope-state-manager for better DX
export { RopeState, KnotData, ProcessedKnot } from '../models/rope-state-manager';
