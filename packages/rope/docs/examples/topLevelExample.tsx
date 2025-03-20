/**
 * This example demonstrates using the useRopeState hook to manage an entire Rope's state
 * from the top level of the application, with state flowing down to child components.
 */

import React, { useState, useEffect } from 'react';
import { useRopeState } from '../../src/hooks/useRopeState';

/**
 * Knot input component that receives state from parent
 */
type KnotInputProps = {
  knotId: string;
  inputIndex: number;
  inputType?: string;
  value: string;
  placeholder?: string;
  error?: { type: string, message: string };
  onChange: (knotId: string, inputIndex: number, value: string) => void;
};

const KnotInput: React.FC<KnotInputProps> = ({
  knotId,
  inputIndex,
  inputType = 'string',
  value,
  placeholder = 'Enter value',
  error,
  onChange,
}) => {
  return (
    <div className="knot-input">
      <input
        type={inputType === 'number' ? 'number' : 'text'}
        value={value || ''}
        onChange={(e) => onChange(knotId, inputIndex, e.target.value)}
        placeholder={placeholder}
        className={error ? 'error' : ''}
      />
      {error && <div className="error-message">{error.message}</div>}
    </div>
  );
};

/**
 * Knot component that receives state from parent
 */
type KnotComponentProps = {
  knot: any;
  availableCoils: Record<string, string>;
  onValueChange: (knotId: string, inputIndex: number, value: string) => void;
  onRemove: (knotId: string) => void;
  onMove: (knotId: string, direction: 'up' | 'down') => void;
};

const KnotComponent: React.FC<KnotComponentProps> = ({
  knot,
  availableCoils,
  onValueChange,
  onRemove,
  onMove,
}) => {
  return (
    <div className={`knot ${knot.isValid ? 'valid' : 'invalid'}`}>
      <div className="knot-header">
        <h3>{knot.protocol} / {knot.action}</h3>
        <div className="knot-actions">
          <button onClick={() => onMove(knot.id, 'up')}>↑</button>
          <button onClick={() => onMove(knot.id, 'down')}>↓</button>
          <button onClick={() => onRemove(knot.id)}>×</button>
        </div>
      </div>
      
      <div className="knot-sentence">
        {/* Render the sentence with inputs */}
        {knot.parsed && knot.parsed.parts.map((part: string, idx: number) => {
          const match = part.match(/\{(\d+)(?:=>(\d+))?\}/);
          
          if (!match) {
            // Regular text part
            return <span key={idx}>{part}</span>;
          }
          
          // This is an input placeholder
          const inputIndex = parseInt(match[1]);
          const input = knot.parsed.inputs.find((i: any) => i.index === inputIndex);
          
          if (!input) {
            return <span key={idx}>[Invalid input {inputIndex}]</span>;
          }
          
          const value = knot.values.get(inputIndex)?.value || '';
          const error = knot.validationErrors.get(inputIndex);
          
          return (
            <KnotInput
              key={idx}
              knotId={knot.id}
              inputIndex={inputIndex}
              inputType={input.type}
              value={value}
              placeholder={`Enter ${input.type || 'value'}`}
              error={error}
              onChange={onValueChange}
            />
          );
        })}
      </div>
      
      {knot.resolvedSentence && (
        <div className="resolved-sentence">
          <strong>Resolved:</strong> {knot.resolvedSentence}
        </div>
      )}
      
      {!knot.isValid && (
        <div className="knot-validation">
          <p>This knot has validation errors</p>
        </div>
      )}
      
      <div className="coil-selector">
        <h4>Available Coils:</h4>
        <div className="coils-list">
          {Object.entries(availableCoils).map(([name, value]) => (
            <div key={name} className="coil-item">
              <span>{name}: {value}</span>
              <button onClick={() => onValueChange(knot.id, 0, `<-{${name}}`)}>
                Use for first input
              </button>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

/**
 * Add Knot Form Component
 */
type AddKnotFormProps = {
  onAddKnot: (knot: { id: string, protocol: string, action: string, sentence: string }) => void;
};

const AddKnotForm: React.FC<AddKnotFormProps> = ({ onAddKnot }) => {
  const [protocol, setProtocol] = useState('');
  const [action, setAction] = useState('');
  const [sentence, setSentence] = useState('');
  
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!protocol || !action || !sentence) return;
    
    onAddKnot({
      id: `knot-${Date.now()}`, // Generate a unique ID
      protocol,
      action,
      sentence
    });
    
    // Reset form
    setProtocol('');
    setAction('');
    setSentence('');
  };
  
  return (
    <form onSubmit={handleSubmit} className="add-knot-form">
      <h3>Add New Knot</h3>
      <div>
        <label>
          Protocol:
          <input
            type="text"
            value={protocol}
            onChange={(e) => setProtocol(e.target.value)}
            placeholder="e.g. uniswap"
            required
          />
        </label>
      </div>
      <div>
        <label>
          Action:
          <input
            type="text"
            value={action}
            onChange={(e) => setAction(e.target.value)}
            placeholder="e.g. swap"
            required
          />
        </label>
      </div>
      <div>
        <label>
          Sentence:
          <input
            type="text"
            value={sentence}
            onChange={(e) => setSentence(e.target.value)}
            placeholder="e.g. Swap {0} for {1}"
            required
          />
        </label>
      </div>
      <button type="submit">Add Knot</button>
    </form>
  );
};

/**
 * Main application component using the useRopeState hook
 */
export const TopLevelRopeExample = () => {
  // Initialize the Rope State with some example knots
  const initialKnots = [
    {
      id: 'knot-1',
      protocol: 'uniswap',
      action: 'swap',
      sentence: 'Swap {0} for {1}',
    },
    {
      id: 'knot-2',
      protocol: 'aave',
      action: 'deposit',
      sentence: 'Deposit {0} into {1}',
    }
  ];
  
  // Use the Rope State hook
  const { state, actions, getters } = useRopeState({ initialKnots });
  
  // Handler for input value changes
  const handleValueChange = (knotId: string, inputIndex: number, value: string) => {
    actions.setValue(knotId, inputIndex, value);
  };
  
  // Submit the entire Rope
  const handleSubmitRope = () => {
    if (!state.isValid) {
      alert('Cannot submit - some knots have validation errors');
      return;
    }
    
    const exportedRope = getters.exportRope();
    alert(`Rope data: ${JSON.stringify(exportedRope, null, 2)}`);
  };
  
  return (
    <div className="rope-container">
      <h1>Rope Builder</h1>
      
      <div className="rope-status">
        <div className={`status-badge ${state.isValid ? 'valid' : 'invalid'}`}>
          {state.isValid ? 'Valid' : 'Invalid'}
        </div>
        <div className={`status-badge ${state.isComplete ? 'complete' : 'incomplete'}`}>
          {state.isComplete ? 'Complete' : 'Incomplete'}
        </div>
      </div>
      
      {state.error && (
        <div className="error-message">
          <strong>Error:</strong> {state.error.message}
        </div>
      )}
      
      <div className="knots-container">
        <h2>Knots</h2>
        {state.knots.map(knot => (
          <KnotComponent
            key={knot.id}
            knot={knot}
            availableCoils={state.allCoils}
            onValueChange={handleValueChange}
            onRemove={actions.removeKnot}
            onMove={actions.moveKnot}
          />
        ))}
      </div>
      
      <AddKnotForm onAddKnot={actions.addKnot} />
      
      <div className="rope-actions">
        <button 
          onClick={handleSubmitRope}
          disabled={!state.isValid || !state.isComplete}
        >
          Submit Rope
        </button>
        <button onClick={actions.clearRope}>Clear All</button>
      </div>
      
      <div className="rope-debug">
        <h3>Debug Information</h3>
        <details>
          <summary>All Coils</summary>
          <pre>{JSON.stringify(state.allCoils, null, 2)}</pre>
        </details>
        <details>
          <summary>All Validation Errors</summary>
          <pre>{JSON.stringify(getters.getAllValidationErrors(), null, 2)}</pre>
        </details>
      </div>
    </div>
  );
};