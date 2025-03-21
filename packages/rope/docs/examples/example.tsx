/**
 * NOTE: This file is an example of how to use the useRope hook with React.
 * It is not meant to be used directly in the codebase, but serves as documentation.
 */

import React, { useState, useEffect } from 'react';
import { useRope } from '../../src/hooks/useRope';

/**
 * Example input component for rendering the parsed sentence
 * with inputs inline in the text.
 */
type SentencePartProps = {
  part: string;
  parsedSentence: any;
  values: Map<number, { value: string }>;
  onInputChange: (index: number, value: string) => void;
  errors: Map<number, { type: string, message: string }>;
};

/**
 * Renders a single part of a sentence, which could be text or an input field
 */
const SentencePart: React.FC<SentencePartProps> = ({
  part,
  parsedSentence,
  values,
  onInputChange,
  errors,
}) => {
  const match = part.match(/\{(\d+)(?:=>(\d+))?\}/);
  
  if (!match) {
    // Regular text part
    return <span>{part}</span>;
  }

  // This is an input placeholder
  const inputIndex = parseInt(match[1]);
  const input = parsedSentence?.inputs?.find((i: any) => i.index === inputIndex);
  
  if (!input) {
    return <span>[Invalid input {inputIndex}]</span>;
  }

  const value = values.get(inputIndex)?.value || '';
  const error = errors.get(inputIndex);
  
  return (
    <span className="input-wrapper">
      <input
        type={input.type === 'number' ? 'number' : 'text'}
        value={value}
        onChange={(e) => onInputChange(inputIndex, e.target.value)}
        placeholder={`Enter ${input.type || 'value'}`}
        className={error ? 'error' : ''}
      />
      {error && <div className="error-message">{error.message}</div>}
    </span>
  );
};

/**
 * Example component demonstrating the useRope hook
 */
export const RopeExample = () => {
  // Example sentence with inputs
  const [sentence, setSentence] = useState('Swap {0} for {1} on {2}');
  
  // Use the Rope hook
  const { 
    state, 
    actions, 
    helpers 
  } = useRope(sentence, { /* Initial values */ });

  // Handle value changes
  const handleInputChange = (index: number, value: string) => {
    actions.setValue(index, value);
  };

  return (
    <div className="rope-example">
      <h2>Sentence Preview</h2>
      
      <div className="sentence">
        {state.parts.map((part, idx) => (
          <SentencePart
            key={idx}
            part={part}
            parsedSentence={state.parsed}
            values={state.values}
            onInputChange={handleInputChange}
            errors={state.validationErrors}
          />
        ))}
      </div>
      
      {state.resolvedSentence && (
        <div className="resolved-sentence">
          <h3>Resolved Sentence</h3>
          <p>{state.resolvedSentence}</p>
        </div>
      )}
      
      <div className="status">
        <p>Status: {helpers.isComplete ? 'Complete' : 'Incomplete'}</p>
        <p>Valid: {helpers.isValid ? 'Yes' : 'No'}</p>
      </div>
      
      {state.error && (
        <div className="error-box">
          <h3>Error</h3>
          <p>{state.error.type}: {state.error.message}</p>
        </div>
      )}
    </div>
  );
};

/**
 * Example usage with coil references
 */
export const RopeWithCoilsExample = () => {
  // Example sentence with inputs that will use coil references
  const [sentence, setSentence] = useState('Send {0} to {1}');
  
  // Available coils (would typically come from a Rope instance)
  const [availableCoils, setAvailableCoils] = useState({
    'amount': '100',
    'token': 'ETH',
    'recipient': '0x1234...'
  });
  
  // Use the Rope hook
  const { state, actions, helpers } = useRope(sentence);

  // Create a coil reference
  const useCoilReference = (coilName: string) => {
    // This would create a reference like "<-{amount}"
    return `<-{${coilName}}`;
  };
  
  // Set a coil reference as a value
  const handleUseCoil = (inputIndex: number, coilName: string) => {
    actions.setValue(inputIndex, useCoilReference(coilName));
  };

  return (
    <div className="rope-with-coils-example">
      <h2>Sentence with Coils</h2>
      
      <div className="sentence">
        {state.parts.map((part, idx) => (
          <SentencePart
            key={idx}
            part={part}
            parsedSentence={state.parsed}
            values={state.values}
            onInputChange={actions.setValue}
            errors={state.validationErrors}
          />
        ))}
      </div>
      
      <div className="coil-selector">
        <h3>Available Coils</h3>
        {Object.entries(availableCoils).map(([name, value]) => (
          <div key={name} className="coil-item">
            <span>{name}: {value}</span>
            <button onClick={() => handleUseCoil(0, name)}>
              Use for input 0
            </button>
            <button onClick={() => handleUseCoil(1, name)}>
              Use for input 1
            </button>
          </div>
        ))}
      </div>
      
      {state.resolvedSentence && (
        <div className="resolved-sentence">
          <h3>Resolved Sentence</h3>
          <p>{state.resolvedSentence}</p>
        </div>
      )}
    </div>
  );
};