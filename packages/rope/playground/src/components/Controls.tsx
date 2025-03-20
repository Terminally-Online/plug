import React, { useState } from 'react';
import { useRopeContext } from '../context/RopeContext';

interface ControlsProps {
  onAddNode: () => void;
}

const Controls: React.FC<ControlsProps> = ({ onAddNode }) => {
  const { knots, removeKnot, addKnot, setNodePosition } = useRopeContext();
  const [selectedKnotId, setSelectedKnotId] = useState<string | null>(null);
  const [newTemplate, setNewTemplate] = useState<string>('');
  
  // Function to extract input placeholders from a template
  const extractInputs = (template: string): string[] => {
    const regex = /\{(\d+)(?::([^}]+))?\}/g;
    const inputs: string[] = [];
    let match;
    
    while ((match = regex.exec(template)) !== null) {
      inputs.push(`{${match[1]}}`);
    }
    
    return inputs;
  };
  
  const handleAddNewKnot = () => {
    if (newTemplate.trim()) {
      // Create a custom ID
      const newKnotId = `knot-${Date.now()}`;
      
      // Place new knots in a reasonable starting position in the viewport
      const position = { x: 300, y: 200 };
      setNodePosition(newKnotId, position);
      
      // Add the custom template with the custom ID
      setTimeout(() => {
        addKnot(newTemplate, newKnotId);
        setNewTemplate('');
      }, 0);
    } else {
      // Use the default template from props
      onAddNode();
    }
  };
  
  return (
    <div className="w-64 bg-gray-100 p-4 overflow-y-auto flex flex-col">
      <h2 className="text-lg font-semibold mb-4">Rope Editor Controls</h2>
      
      <div className="mb-4">
        <input
          type="text"
          className="w-full px-3 py-2 border rounded mb-2"
          placeholder="Enter template (e.g. 'Swap {0} for {1}')"
          value={newTemplate}
          onChange={(e) => setNewTemplate(e.target.value)}
        />
        
        <button 
          className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded w-full"
          onClick={handleAddNewKnot}
        >
          {newTemplate.trim() ? 'Add Custom Knot' : 'Add New Knot'}
        </button>
      </div>
      
      <div className="border-t border-gray-300 my-4" />
      
      <h3 className="text-md font-medium mb-2">Knots</h3>
      
      <div className="space-y-2">
        {knots.map((knot) => {
          // Extract inputs to show as a preview
          const inputPlaceholders = extractInputs(knot.sentence);
          
          return (
            <div 
              key={knot.id} 
              className={`
                p-2 rounded cursor-pointer hover:bg-gray-200
                ${selectedKnotId === knot.id ? 'bg-gray-200' : 'bg-white'}
              `}
              onClick={() => setSelectedKnotId(knot.id)}
            >
              <div className="flex justify-between items-center">
                <span className="text-sm truncate flex-1">{knot.sentence}</span>
                <button 
                  className="text-red-500 hover:text-red-700 ml-2"
                  onClick={(e) => {
                    e.stopPropagation();
                    removeKnot(knot.id);
                    if (selectedKnotId === knot.id) {
                      setSelectedKnotId(null);
                    }
                  }}
                >
                  &times;
                </button>
              </div>
              
              {/* Show input placeholders */}
              {inputPlaceholders.length > 0 && (
                <div className="mt-1 flex flex-wrap gap-1">
                  {inputPlaceholders.map((input, idx) => (
                    <span 
                      key={`${knot.id}-input-${idx}`}
                      className="text-xs bg-gray-200 px-1 rounded"
                    >
                      {input}
                    </span>
                  ))}
                </div>
              )}
              
              {selectedKnotId === knot.id && (
                <div className="mt-2 text-xs text-gray-600">
                  <p>ID: {knot.id}</p>
                  <p>Protocol: {knot.protocol}</p>
                  <p>Action: {knot.action}</p>
                  <p>Complete: {knot.isComplete ? 'Yes' : 'No'}</p>
                  <p>Valid: {knot.isValid ? 'Yes' : 'No'}</p>
                </div>
              )}
            </div>
          );
        })}
      </div>
      
      {knots.length === 0 && (
        <p className="text-gray-500 text-sm">No knots added yet</p>
      )}
      
      <div className="border-t border-gray-300 my-4" />
      
      <div className="mt-auto text-xs text-gray-600">
        <h4 className="font-medium mb-1">Quick Tips:</h4>
        <ul className="list-disc pl-4 space-y-1">
          <li>Drag nodes to reposition them</li>
          <li>Drag from <span className="inline-block w-2 h-2 bg-green-500 rounded-full"></span> input handles to create value nodes</li>
          <li>Drag from <span className="inline-block w-2 h-2 bg-blue-500 rounded-full"></span> output handle to create a new sentence node</li>
          <li>Press 'Delete' key to remove nodes</li>
        </ul>
      </div>
    </div>
  );
};

export default Controls;