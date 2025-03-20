import React, { createContext, useContext, useState, useCallback, useMemo, useEffect } from 'react';
import { useRope } from '../../../src/hooks/useRope';

type KnotValue = {
  knotId: string;
  inputIndex: number;
  value: string;
};

type RopeContextType = {
  knots: any[];
  sentenceTemplates: string[];
  nodePositions: Record<string, { x: number, y: number }>;
  setNodePosition: (id: string, position: { x: number, y: number }) => void;
  addKnot: (template: string) => void;
  removeKnot: (id: string) => void;
  updateKnotValue: (knotId: string, inputIndex: number, value: string) => void;
};

const RopeContext = createContext<RopeContextType | null>(null);

interface RopeProviderProps {
  children: React.ReactNode;
  initialSentences: string[];
}

export const RopeProvider: React.FC<RopeProviderProps> = ({ children, initialSentences }) => {
  const [sentenceTemplates] = useState<string[]>(initialSentences);
  const [nodePositions, setNodePositions] = useState<Record<string, { x: number, y: number }>>({});
  
  // Initialize with empty knots, we'll create them from templates
  const { state, actions } = useRope({ 
    initialKnots: [] 
  });

  // Initialize knots from templates if needed
  useEffect(() => {
    if (Object.keys(state.knots || {}).length === 0 && sentenceTemplates.length > 0) {
      // Create a knot for each template with a staggered layout
      sentenceTemplates.forEach((template, index) => {
        const id = `knot-${index}`;
        
        // Set initial position with staggered layout
        // Initial knots are arranged in a diagonal pattern
        setNodePositions(prev => ({
          ...prev,
          [id]: { 
            x: 100 + index * 80, 
            y: 100 + index * 100 
          }
        }));
        
        // Add the knot with the template
        actions.addKnot({
          id,
          sentence: template,
          protocol: "default",
          action: "default",
        });
      });
    }
  }, [state.knots, sentenceTemplates, actions]);

  const setNodePosition = useCallback((id: string, position: { x: number, y: number }) => {
    setNodePositions(prev => ({
      ...prev,
      [id]: position
    }));
  }, []);

  const addKnot = useCallback((template: string, customId?: string) => {
    // Use provided ID or generate a new one
    const id = customId || `knot-${Date.now()}`;
    
    // Add the knot to the system
    actions.addKnot({
      id,
      sentence: template,
      protocol: "default",
      action: "default",
    });
    
    // Position is managed separately through setNodePosition
    // We let the caller set the position before calling this function
    // This ensures that nodes appear exactly where users expect them
  }, [actions]);

  const removeKnot = useCallback((id: string) => {
    actions.removeKnot(id);
    
    // Remove position data
    setNodePositions(prev => {
      const newPositions = { ...prev };
      delete newPositions[id];
      return newPositions;
    });
  }, [actions]);

  const updateKnotValue = useCallback((knotId: string, inputIndex: number, value: string) => {
    actions.setValue(knotId, inputIndex, value);
  }, [actions]);

  // Convert the knots object to an array for easier rendering
  const knots = useMemo(() => {
    return Object.entries(state.knots || {}).map(([id, knot]) => ({
      id,
      ...knot
    }));
  }, [state.knots]);

  // Provide the context values to children
  const contextValue = {
    knots,
    sentenceTemplates,
    nodePositions,
    setNodePosition,
    addKnot,
    removeKnot,
    updateKnotValue
  };

  return (
    <RopeContext.Provider value={contextValue}>
      {children}
    </RopeContext.Provider>
  );
};

export const useRopeContext = () => {
  const context = useContext(RopeContext);
  if (!context) {
    throw new Error('useRopeContext must be used within a RopeProvider');
  }
  return context;
};