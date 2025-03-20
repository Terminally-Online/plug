import { SentenceService, ParsedSentence } from '../services/sentence';
import { CoilService } from '../services/coil';

// Core types
export type KnotId = string;

export interface KnotData {
  id: KnotId;
  protocol: string;
  action: string;
  sentence: string;
  values?: Record<string, string>;
}

export interface ProcessedKnot {
  id: KnotId;
  protocol: string;
  action: string;
  sentence: string;
  values: Map<number, { value: string }>;
  parsed: ParsedSentence | null;
  resolvedSentence: string | null;
  validationErrors: Map<number, { type: string; message: string }>;
  isComplete: boolean;
  isValid: boolean;
  parts?: string[];
  // Linked list refs
  next: KnotId | null;
  prev: KnotId | null;
}

// State structure that will be used by the hooks
export interface RopeState {
  // Maintains the order of knots
  head: KnotId | null;
  tail: KnotId | null;
  // Stores all knot data for O(1) lookup
  knots: Record<KnotId, ProcessedKnot>;
  // For efficient iteration in specific contexts
  knotIds: KnotId[];
  // Top-level properties for quick access
  isValid: boolean;
  isComplete: boolean;
  activeKnotId: KnotId | null;
  // Coil system
  allCoils: Record<string, string>;
  // Error state
  error: { type: string; message: string } | null;
  // Metadata for optimizations
  lastUpdated: KnotId | null;
}

/**
 * Core Rope state management class using an immutable linked list approach
 * for efficient updates and processing
 */
export class RopeStateManager {
  private sentenceService = new SentenceService();
  private coilService = new CoilService();
  private state: RopeState;
  private listeners: Set<(state: RopeState) => void> = new Set();
  
  constructor(initialState?: Partial<RopeState>) {
    this.state = {
      head: null,
      tail: null,
      knots: {},
      knotIds: [],
      isValid: false,
      isComplete: false,
      activeKnotId: null,
      allCoils: {},
      error: null,
      lastUpdated: null,
      ...initialState
    };
  }
  
  // State subscription system
  subscribe(listener: (state: RopeState) => void): () => void {
    this.listeners.add(listener);
    return () => {
      this.listeners.delete(listener);
    };
  }
  
  private notify() {
    for (const listener of this.listeners) {
      listener(this.state);
    }
  }
  
  // Get current state
  getState(): RopeState {
    return this.state;
  }
  
  // Update state with immutability
  private setState(updater: (state: RopeState) => RopeState) {
    this.state = updater(this.state);
    this.notify();
  }
  
  // Add a knot to the rope
  addKnot(knotData: KnotData): void {
    this.setState(state => {
      // Check if a knot with this ID already exists
      if (state.knots[knotData.id]) {
        return {
          ...state,
          error: {
            type: 'duplicate_knot',
            message: `A knot with ID "${knotData.id}" already exists`
          }
        };
      }
      
      // Process the knot
      const processed = this.processKnot(knotData);
      
      // Link it in the list
      const next = null;
      const prev = state.tail;
      
      const newKnot: ProcessedKnot = {
        ...processed,
        next,
        prev
      };
      
      // Update the knotIds array
      const knotIds = [...state.knotIds, knotData.id];
      
      // Update the knots lookup
      const knots = { ...state.knots, [knotData.id]: newKnot };
      
      // Update head and tail
      let head = state.head;
      if (!head) head = knotData.id;
      
      const tail = knotData.id;
      
      // Update the previous tail's next pointer if it exists
      if (state.tail) {
        const prevTail = { ...knots[state.tail] };
        prevTail.next = knotData.id;
        knots[state.tail] = prevTail;
      }
      
      // Recalculate coils
      const allCoils = this.calculateCoils(knots, knotIds);
      
      // Recalculate overall state
      const isComplete = Object.values(knots).every(k => k.isComplete);
      const isValid = Object.values(knots).every(k => k.isValid);
      
      return {
        ...state,
        head,
        tail,
        knots,
        knotIds,
        allCoils,
        isComplete,
        isValid,
        activeKnotId: knotData.id,
        lastUpdated: knotData.id,
        error: null
      };
    });
  }
  
  // Remove a knot from the rope
  removeKnot(knotId: KnotId): void {
    this.setState(state => {
      if (!state.knots[knotId]) return state;
      
      const knot = state.knots[knotId];
      const knots = { ...state.knots };
      
      // Update the next and previous pointers
      if (knot.prev && knots[knot.prev]) {
        knots[knot.prev] = {
          ...knots[knot.prev],
          next: knot.next
        };
      }
      
      if (knot.next && knots[knot.next]) {
        knots[knot.next] = {
          ...knots[knot.next],
          prev: knot.prev
        };
      }
      
      // Update head and tail if needed
      let head = state.head;
      let tail = state.tail;
      
      if (state.head === knotId) {
        head = knot.next;
      }
      
      if (state.tail === knotId) {
        tail = knot.prev;
      }
      
      // Remove the knot
      delete knots[knotId];
      
      // Update the knotIds array
      const knotIds = state.knotIds.filter(id => id !== knotId);
      
      // Recalculate coils
      const allCoils = this.calculateCoils(knots, knotIds);
      
      // Recalculate overall state
      const isComplete = Object.values(knots).every(k => k.isComplete);
      const isValid = Object.values(knots).every(k => k.isValid);
      
      // Update active knot if needed
      let activeKnotId = state.activeKnotId;
      if (activeKnotId === knotId) {
        activeKnotId = head || tail || null;
      }
      
      return {
        ...state,
        head,
        tail,
        knots,
        knotIds,
        allCoils,
        isComplete,
        isValid,
        activeKnotId,
        lastUpdated: null,
        error: null
      };
    });
  }
  
  // Set a value for an input in a specific knot
  setValue(knotId: KnotId, inputIndex: number, value: string | undefined): void {
    this.setState(state => {
      const knot = state.knots[knotId];
      if (!knot || !knot.parsed) return state;
      
      // Create updated knot
      const knots = { ...state.knots };
      
      // Update the value
      const result = this.sentenceService.setValue(
        knot.parsed,
        knot.values,
        inputIndex,
        value ?? ''
      );
      
      // Process the updated knot
      const updatedKnot = {
        ...knot,
        values: result.values
      };
      
      // Apply the initial update
      knots[knotId] = updatedKnot;
      
      // Recalculate coils
      const allCoils = this.calculateCoils(knots, state.knotIds);
      
      // Reprocess all knots with the updated coils
      const processedKnots = this.processAllKnots(knots, state.knotIds, allCoils);
      
      // Recalculate overall state
      const isComplete = Object.values(processedKnots).every(k => k.isComplete);
      const isValid = Object.values(processedKnots).every(k => k.isValid);
      
      return {
        ...state,
        knots: processedKnots,
        allCoils,
        isComplete,
        isValid,
        lastUpdated: knotId,
        error: null
      };
    });
  }
  
  // Update a knot's sentence
  updateKnotSentence(knotId: KnotId, sentence: string): void {
    this.setState(state => {
      const knot = state.knots[knotId];
      if (!knot) return state;
      
      // Check if this update would actually change anything
      if (knot.sentence === sentence) {
        return state;
      }
      
      try {
        // Parse the new sentence
        const parsed = this.sentenceService.parseSentence(sentence);
        
        // Create updated knots
        const knots = { ...state.knots };
        
        // Update the knot with the new sentence and parsed result
        knots[knotId] = {
          ...knot,
          sentence,
          parsed,
          resolvedSentence: null,
          validationErrors: new Map()
        };
        
        // Recalculate coils
        const allCoils = this.calculateCoils(knots, state.knotIds);
        
        // Reprocess all knots with the updated coils
        const processedKnots = this.processAllKnots(knots, state.knotIds, allCoils);
        
        // Recalculate overall state
        const isComplete = Object.values(processedKnots).every(k => k.isComplete);
        const isValid = Object.values(processedKnots).every(k => k.isValid);
        
        return {
          ...state,
          knots: processedKnots,
          allCoils,
          isComplete,
          isValid,
          lastUpdated: knotId,
          error: null
        };
      } catch (error) {
        return {
          ...state,
          error: {
            type: 'parse',
            message: error instanceof Error ? error.message : 'Unknown parsing error'
          }
        };
      }
    });
  }
  
  // Set the active knot
  setActiveKnot(knotId: KnotId): void {
    this.setState(state => {
      if (!state.knots[knotId]) return state;
      
      return {
        ...state,
        activeKnotId: knotId,
        lastUpdated: knotId
      };
    });
  }
  
  // Move a knot up or down in the sequence
  moveKnot(knotId: KnotId, direction: 'up' | 'down'): void {
    this.setState(state => {
      const knot = state.knots[knotId];
      if (!knot) return state;
      
      // Can't move up if it's the first knot
      if (direction === 'up' && !knot.prev) return state;
      
      // Can't move down if it's the last knot
      if (direction === 'down' && !knot.next) return state;
      
      // Copy knots for immutable updates
      const knots = { ...state.knots };
      
      // Handle moving up
      if (direction === 'up' && knot.prev) {
        const prevKnot = knots[knot.prev];
        
        // Swap positions in the linked list
        const prevPrev = prevKnot.prev;
        
        // Update this knot
        knots[knotId] = {
          ...knot,
          prev: prevPrev,
          next: knot.prev
        };
        
        // Update previous knot
        knots[knot.prev] = {
          ...prevKnot,
          prev: knotId,
          next: knot.next
        };
        
        // Update knot before the previous knot
        if (prevPrev) {
          knots[prevPrev] = {
            ...knots[prevPrev],
            next: knotId
          };
        }
        
        // Update knot after this knot
        if (knot.next) {
          knots[knot.next] = {
            ...knots[knot.next],
            prev: knot.prev
          };
        }
        
        // Update head if needed
        let head = state.head;
        if (state.head === knot.prev) {
          head = knotId;
        }
        
        // Update tail if needed
        let tail = state.tail;
        if (state.tail === knotId) {
          tail = knot.prev;
        }
        
        // Update knotIds array to reflect the new order
        const knotIds = [...state.knotIds];
        const thisIndex = knotIds.indexOf(knotId);
        const prevIndex = knotIds.indexOf(knot.prev);
        if (thisIndex !== -1 && prevIndex !== -1) {
          knotIds[thisIndex] = knot.prev;
          knotIds[prevIndex] = knotId;
        }
        
        return {
          ...state,
          head,
          tail,
          knots,
          knotIds,
          lastUpdated: knotId
        };
      }
      
      // Handle moving down
      if (direction === 'down' && knot.next) {
        const nextKnot = knots[knot.next];
        
        // Swap positions in the linked list
        const nextNext = nextKnot.next;
        
        // Update this knot
        knots[knotId] = {
          ...knot,
          prev: knot.next,
          next: nextNext
        };
        
        // Update next knot
        knots[knot.next] = {
          ...nextKnot,
          prev: knot.prev,
          next: knotId
        };
        
        // Update knot before this knot
        if (knot.prev) {
          knots[knot.prev] = {
            ...knots[knot.prev],
            next: knot.next
          };
        }
        
        // Update knot after the next knot
        if (nextNext) {
          knots[nextNext] = {
            ...knots[nextNext],
            prev: knotId
          };
        }
        
        // Update head if needed
        let head = state.head;
        if (state.head === knotId) {
          head = knot.next;
        }
        
        // Update tail if needed
        let tail = state.tail;
        if (state.tail === knot.next) {
          tail = knotId;
        }
        
        // Update knotIds array to reflect the new order
        const knotIds = [...state.knotIds];
        const thisIndex = knotIds.indexOf(knotId);
        const nextIndex = knotIds.indexOf(knot.next);
        if (thisIndex !== -1 && nextIndex !== -1) {
          knotIds[thisIndex] = knot.next;
          knotIds[nextIndex] = knotId;
        }
        
        return {
          ...state,
          head,
          tail,
          knots,
          knotIds,
          lastUpdated: knotId
        };
      }
      
      return state;
    });
  }
  
  // Clear all knots
  clearRope(): void {
    this.setState(_ => ({
      head: null,
      tail: null,
      knots: {},
      knotIds: [],
      isValid: false,
      isComplete: false,
      activeKnotId: null,
      allCoils: {},
      error: null,
      lastUpdated: null
    }));
  }
  
  // Reset a knot to its initial state
  resetKnot(knotId: KnotId): void {
    this.setState(state => {
      const knot = state.knots[knotId];
      if (!knot) return state;
      
      // Create updated knots
      const knots = { ...state.knots };
      
      // Reset the knot values and errors
      knots[knotId] = {
        ...knot,
        values: new Map(),
        validationErrors: new Map(),
        resolvedSentence: null,
        isComplete: false,
        isValid: false
      };
      
      // Recalculate coils
      const allCoils = this.calculateCoils(knots, state.knotIds);
      
      // Reprocess all knots with the updated coils
      const processedKnots = this.processAllKnots(knots, state.knotIds, allCoils);
      
      // Recalculate overall state
      const isComplete = Object.values(processedKnots).every(k => k.isComplete);
      const isValid = Object.values(processedKnots).every(k => k.isValid);
      
      return {
        ...state,
        knots: processedKnots,
        allCoils,
        isComplete,
        isValid,
        lastUpdated: knotId,
        error: null
      };
    });
  }
  
  // Process a knot from scratch
  private processKnot(knotData: KnotData): Omit<ProcessedKnot, 'next' | 'prev'> {
    try {
      // Parse the sentence
      const parsed = this.sentenceService.parseSentence(knotData.sentence);
      
      // Create values map
      const values = new Map<number, { value: string }>();
      if (knotData.values) {
        Object.entries(knotData.values).forEach(([key, value]) => {
          if (value !== undefined) {
            values.set(Number(key), { value });
          }
        });
      }
      
      // Filter inputs based on dependencies
      const filteredInputs = parsed.inputs.filter(input =>
        this.sentenceService.shouldRenderInput(
          input.type || 'string',
          parsed.inputs,
          index => values.get(index)
        )
      );
      
      // Create parsed with filtered inputs
      const parsedWithFilteredInputs = {
        ...parsed,
        inputs: filteredInputs
      };
      
      // Split template into parts for rendering
      const parts = parsedWithFilteredInputs.parts.map(part => {
        if (part.match(/\{[^}]+\}/)) return [part];
        return part.split(/(\s+)/g);
      }).flat();
      
      // Check if all required inputs have values
      const isComplete = filteredInputs
        .filter(input => input.required)
        .every(input => {
          const value = values.get(input.index);
          return value !== undefined && value.value !== '';
        });
      
      // Validate all inputs
      const validationErrors = new Map<number, { type: string, message: string }>();
      filteredInputs.forEach(input => {
        const value = values.get(input.index)?.value;
        if (value !== undefined) {
          const result = this.sentenceService.validateInput(value, input.type);
          if (!result.success) {
            validationErrors.set(input.index, {
              type: 'validation',
              message: result.error || 'Invalid value'
            });
          }
        } else if (input.required) {
          validationErrors.set(input.index, {
            type: 'required',
            message: 'This field is required'
          });
        }
      });
      
      return {
        id: knotData.id,
        protocol: knotData.protocol,
        action: knotData.action,
        sentence: knotData.sentence,
        values,
        parsed: parsedWithFilteredInputs,
        parts,
        resolvedSentence: null, // Will be resolved later with coils
        validationErrors,
        isComplete,
        isValid: isComplete && validationErrors.size === 0
      };
    } catch (e) {
      // Return minimal valid knot on error
      return {
        id: knotData.id,
        protocol: knotData.protocol,
        action: knotData.action,
        sentence: knotData.sentence,
        values: new Map(),
        parsed: null,
        resolvedSentence: null,
        validationErrors: new Map(),
        isComplete: false,
        isValid: false
      };
    }
  }
  
  // Process all knots with updated coils
  private processAllKnots(
    knots: Record<KnotId, ProcessedKnot>,
    knotIds: KnotId[],
    coils: Record<string, string>
  ): Record<KnotId, ProcessedKnot> {
    const result = { ...knots };
    
    // Process each knot in order (important for coil dependencies)
    let current = this.state.head;
    while (current) {
      const knot = result[current];
      if (!knot) break;
      
      // Process the knot with the updated coils
      result[current] = this.processExistingKnot(knot, coils);
      
      // Move to the next knot
      current = knot.next;
    }
    
    return result;
  }
  
  // Process an existing knot with coils
  private processExistingKnot(knot: ProcessedKnot, coils: Record<string, string>): ProcessedKnot {
    if (!knot.parsed) return knot;
    
    // Filter inputs based on dependencies
    const filteredInputs = knot.parsed.inputs.filter(input =>
      this.sentenceService.shouldRenderInput(
        input.type || 'string',
        knot.parsed?.inputs || [],
        index => knot.values.get(index)
      )
    );
    
    // Create parsed with filtered inputs
    const parsedWithFilteredInputs = {
      ...knot.parsed,
      inputs: filteredInputs
    };
    
    // Split template into parts for rendering
    const parts = parsedWithFilteredInputs.parts.map(part => {
      if (part.match(/\{[^}]+\}/)) return [part];
      return part.split(/(\s+)/g);
    }).flat();
    
    // Check if all required inputs have values
    const isComplete = filteredInputs
      .filter(input => input.required)
      .every(input => {
        const value = knot.values.get(input.index);
        return value !== undefined && value.value !== '';
      });
    
    // Validate all inputs
    const validationErrors = new Map<number, { type: string, message: string }>();
    filteredInputs.forEach(input => {
      const value = knot.values.get(input.index)?.value;
      if (value !== undefined) {
        const result = this.sentenceService.validateInput(value, input.type);
        if (!result.success) {
          validationErrors.set(input.index, {
            type: 'validation',
            message: result.error || 'Invalid value'
          });
        }
      } else if (input.required) {
        validationErrors.set(input.index, {
          type: 'required',
          message: 'This field is required'
        });
      }
    });
    
    // Resolve sentence with coils if complete
    let resolvedSentence: string | null = null;
    try {
      if (isComplete) {
        const valuesObj = Object.fromEntries(
          Array.from(knot.values.entries()).map(([k, v]) => [k, v.value])
        );
        
        // Use lazy coil resolution
        const resolvedValues = this.coilService.resolveCoilReferences(
          valuesObj,
          coils,
          { cache: true }
        );
        
        resolvedSentence = this.sentenceService.formatSentence(knot.parsed, resolvedValues);
      }
    } catch (e) {
      // Leave resolvedSentence as null on error
    }
    
    return {
      ...knot,
      parsed: parsedWithFilteredInputs,
      parts,
      resolvedSentence,
      validationErrors,
      isComplete,
      isValid: isComplete && validationErrors.size === 0
    };
  }
  
  // Recalculate coils for the entire rope
  private calculateCoils(knots: Record<KnotId, ProcessedKnot>, knotIds: KnotId[]): Record<string, string> {
    const coils: Record<string, string> = {};
    
    // Process knots in order (follow the linked list)
    let current = this.state.head;
    let index = 0;
    
    while (current) {
      const knot = knots[current];
      if (!knot) break;
      
      // Add standard coils
      const potentialCoils = [
        `amount_${knot.protocol}`,
        `token_${knot.action}`,
        `address_${knot.id}`,
      ];
      
      potentialCoils.forEach(coilName => {
        // With index
        coils[`${coilName}@${index}`] = `Value for ${coilName}`;
        // Without index (first one wins)
        if (!coils[coilName]) {
          coils[coilName] = `Value for ${coilName}`;
        }
      });
      
      // Add resolved sentence as a coil if available
      if (knot.resolvedSentence) {
        coils[`knot_${index}`] = knot.resolvedSentence;
        coils[`knot_${knot.id}`] = knot.resolvedSentence;
      }
      
      // Move to the next knot
      current = knot.next;
      index++;
    }
    
    return coils;
  }
  
  // Get all knots as an array in order
  getKnotsArray(): ProcessedKnot[] {
    const result: ProcessedKnot[] = [];
    let current = this.state.head;
    
    while (current) {
      const knot = this.state.knots[current];
      if (!knot) break;
      
      result.push(knot);
      current = knot.next;
    }
    
    return result;
  }
  
  // Get a knot by ID with O(1) lookup
  getKnot(id: KnotId): ProcessedKnot | null {
    return this.state.knots[id] || null;
  }
  
  // Iterate through knots
  *iterateKnots(): Generator<ProcessedKnot> {
    let current = this.state.head;
    
    while (current) {
      const knot = this.state.knots[current];
      if (!knot) break;
      
      yield knot;
      current = knot.next;
    }
  }
}