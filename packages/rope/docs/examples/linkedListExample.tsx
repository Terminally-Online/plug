import React, { useState } from 'react';
import { useRope, useKnot, KnotId, KnotData } from '../../src';

/**
 * Example of a Knot Editor using the new linked list architecture
 */
const KnotEditor: React.FC<{ knotId: KnotId }> = ({ knotId }) => {
  const { knot, actions, wasUpdated, isActive } = useKnot(knotId);
  
  if (!knot) return <div>Knot not found</div>;
  
  return (
    <div className={`knot-editor ${isActive ? 'active' : ''} ${wasUpdated ? 'updated' : ''}`}>
      <div className="knot-header">
        <h3>{knot.protocol} - {knot.action}</h3>
        <div className="knot-actions">
          <button onClick={actions.moveUp}>⬆️</button>
          <button onClick={actions.moveDown}>⬇️</button>
          <button onClick={actions.remove}>🗑️</button>
          <button onClick={actions.reset}>🔄</button>
        </div>
      </div>
      
      <div className="knot-body">
        <div className="knot-sentence">
          {knot.resolvedSentence || knot.sentence}
        </div>
        
        {knot.parsed?.inputs.map(input => (
          <div key={input.index} className="knot-input">
            <label>{input.name || `Input ${input.index}`}</label>
            <input
              type="text"
              value={knot.values.get(input.index)?.value || ''}
              onChange={e => actions.setValue(input.index, e.target.value)}
              placeholder={input.type || 'string'}
            />
            {knot.validationErrors.has(input.index) && (
              <div className="error">
                {knot.validationErrors.get(input.index)?.message}
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
};

/**
 * Main Rope Editor Component with Linked List Architecture
 */
export const RopeEditor: React.FC = () => {
  const { state, actions, helpers } = useRope({
    initialKnots: [
      {
        id: 'knot-1',
        protocol: 'uniswap',
        action: 'swap',
        sentence: 'Swap {0} for {1} on {2}',
        values: { '0': 'ETH', '1': 'USDC' }
      }
    ]
  });
  
  const [newKnotCounter, setNewKnotCounter] = useState(2);
  
  const handleAddKnot = () => {
    const knotId = `knot-${newKnotCounter}`;
    actions.addKnot({
      id: knotId,
      protocol: 'compound',
      action: 'supply',
      sentence: 'Supply {0} tokens to {1}',
      values: {}
    });
    setNewKnotCounter(prev => prev + 1);
  };
  
  // Render knots in order using the linked list structure
  const renderKnots = () => {
    const knotElements: React.ReactNode[] = [];
    let currentId = state.head;
    
    while (currentId) {
      knotElements.push(<KnotEditor key={currentId} knotId={currentId} />);
      const currentKnot = state.knots[currentId];
      if (!currentKnot) break;
      currentId = currentKnot.next;
    }
    
    return knotElements;
  };
  
  return (
    <div className="rope-editor">
      <div className="rope-header">
        <h2>Rope Editor</h2>
        <div className="rope-status">
          <span>Status: {state.isComplete ? 'Complete' : 'Incomplete'}</span>
          <span>Valid: {state.isValid ? 'Yes' : 'No'}</span>
        </div>
        <div className="rope-actions">
          <button onClick={handleAddKnot}>Add Knot</button>
          <button onClick={actions.clearRope}>Clear All</button>
        </div>
      </div>
      
      <div className="rope-knots">
        {renderKnots()}
      </div>
      
      {state.error && (
        <div className="rope-error">
          <h3>Error: {state.error.type}</h3>
          <p>{state.error.message}</p>
        </div>
      )}
      
      <div className="rope-export">
        <h3>Exported Data:</h3>
        <pre>{JSON.stringify(helpers.exportRope(), null, 2)}</pre>
      </div>
    </div>
  );
};

/**
 * Example of using the single knot mode
 */
export const SingleKnotExample: React.FC = () => {
  const { state, actions, helpers } = useRope({
    initialSentence: 'I want to swap {0} for {1} on chain {2}',
    initialValues: {
      '0': 'ETH'
    }
  });
  
  const activeKnot = helpers.getActiveKnot();
  
  if (!activeKnot) return <div>No knot found</div>;
  
  return (
    <div className="single-knot-example">
      <h2>Single Knot Mode</h2>
      
      <div className="knot-sentence">
        {activeKnot.resolvedSentence || activeKnot.sentence}
      </div>
      
      {activeKnot.parsed?.inputs.map(input => (
        <div key={input.index} className="knot-input">
          <label>{input.name || `Input ${input.index}`}</label>
          <input
            type="text"
            value={activeKnot.values.get(input.index)?.value || ''}
            onChange={e => actions.setActiveKnotValue(input.index, e.target.value)}
            placeholder={input.type || 'string'}
          />
          {activeKnot.validationErrors.has(input.index) && (
            <div className="error">
              {activeKnot.validationErrors.get(input.index)?.message}
            </div>
          )}
        </div>
      ))}
      
      <div className="knot-status">
        <span>Complete: {activeKnot.isComplete ? 'Yes' : 'No'}</span>
        <span>Valid: {activeKnot.isValid ? 'Yes' : 'No'}</span>
      </div>
    </div>
  );
};

/**
 * Example demonstrating coil reference functionality
 */
export const CoilReferenceExample: React.FC = () => {
  const { state, actions, helpers } = useRope({
    initialKnots: [
      {
        id: 'source-knot',
        protocol: 'uniswap',
        action: 'swap',
        sentence: 'Swap {0} for {1}',
        values: { '0': 'ETH', '1': 'USDC' }
      },
      {
        id: 'dependent-knot',
        protocol: 'compound',
        action: 'supply',
        sentence: 'Supply {0} to Compound',
        values: { '0': '<-{token_swap}' } // Reference to the swap coil
      }
    ]
  });
  
  return (
    <div className="coil-example">
      <h2>Coil References Example</h2>
      
      <div className="knots">
        {helpers.getKnotsArray().map(knot => (
          <div key={knot.id} className="knot">
            <h3>{knot.protocol} - {knot.action}</h3>
            <div className="sentence">
              <strong>Original:</strong> {knot.sentence}
            </div>
            <div className="sentence">
              <strong>Resolved:</strong> {knot.resolvedSentence || 'Not resolved yet'}
            </div>
            
            <div className="inputs">
              {knot.parsed?.inputs.map(input => (
                <div key={input.index} className="input">
                  <label>{input.name || `Input ${input.index}`}</label>
                  <input
                    type="text"
                    value={knot.values.get(input.index)?.value || ''}
                    onChange={e => actions.setValue(knot.id, input.index, e.target.value)}
                  />
                </div>
              ))}
            </div>
          </div>
        ))}
      </div>
      
      <div className="coil-registry">
        <h3>Available Coils:</h3>
        <ul>
          {Object.entries(state.allCoils).map(([name, value]) => (
            <li key={name}>
              <strong>{name}:</strong> {value}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};