import { memo, useState, useEffect, useRef } from 'react';
import { Handle, Position, NodeProps } from 'reactflow';

// Node data for value nodes
// Connection between a value node and an input
export interface ValueConnection {
  targetId: string;   // The target knot's ID
  inputIndex: number; // The input index on the target knot
}

export interface ValueNodeData {
  label: string;
  value: string;
  inputIndex?: number;
  targetKnotId?: string;  // The original knot ID this value is connected to (for backward compatibility)
  connections?: ValueConnection[]; // All connections from this value node
  onChange?: (value: string) => void;
}

const ValueNode = ({ data, id }: NodeProps<ValueNodeData>) => {
  // Use state with callback to ensure we always have the latest value
  const [value, setValue] = useState(data.value || '');
  const inputRef = useRef<HTMLInputElement>(null);
  
  // Update the internal state if the external data changes
  useEffect(() => {
    if (data.value !== undefined && data.value !== value) {
      setValue(data.value);
    }
  }, [data.value]);
  
  // Auto-focus the input when the node is created
  useEffect(() => {
    // Focus the input field after creation
    // Short delay to ensure the node is fully rendered
    const timer = setTimeout(() => {
      if (inputRef.current) {
        inputRef.current.focus();
      }
    }, 100);
    
    return () => clearTimeout(timer);
  }, []);
  
  // Prevent focus loss when updating state
  const handleValueChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newValue = e.target.value;
    setValue(newValue);
    
    // Update all connected inputs with the new value
    // Use requestAnimationFrame to batch the update with rendering
    requestAnimationFrame(() => {
      // Simply call onChange once - it will handle updating all connections
      if (data.onChange) {
        data.onChange(newValue);
      }
    });
  };
  
  return (
    <div className="p-3 rounded-lg shadow-md bg-white border border-green-400 max-w-[200px] bg-gradient-to-br from-white to-green-50">
      {/* Output handle on the right */}
      <Handle
        type="source"
        position={Position.Right}
        id="value-output"
        className="w-3 h-3 bg-green-500 !right-[-5px]"
        title="Connect to input"
      />
      
      <div className="flex items-center mb-2">
        <div className="w-2 h-2 bg-green-500 rounded-full mr-2" />
        <div className="font-medium text-gray-700 text-xs">
          {data.connections && data.connections.length > 1 
            ? `Shared Value (${data.connections.length} connections)` 
            : (data.label || 'Value')}
        </div>
      </div>
      
      <input
        ref={inputRef}
        type="text"
        className="w-full px-2 py-1 text-sm border border-gray-200 rounded-md focus:outline-none focus:ring-1 focus:ring-green-500 focus:border-green-500"
        placeholder="Enter value..."
        value={value}
        onChange={handleValueChange}
        // Prevent event propagation to stop ReactFlow from stealing focus
        onKeyDown={(e) => e.stopPropagation()}
        onClick={(e) => e.stopPropagation()}
      />
      
      {/* Show connections information */}
      <div className="mt-1 text-xs text-gray-500 flex items-center space-x-1">
        {/* Original input reference if available */}
        {data.inputIndex !== undefined && (
          <span className="bg-green-100 text-green-800 text-xs font-medium px-2 py-0.5 rounded">
            Input {data.inputIndex}
          </span>
        )}
        
        {/* Show number of connections */}
        {data.connections && data.connections.length > 0 && (
          <span className="bg-blue-100 text-blue-800 text-xs font-medium px-2 py-0.5 rounded flex items-center">
            <svg className="w-3 h-3 mr-0.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
              <path d="M5 3a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2V5a2 2 0 00-2-2H5zM5 11a2 2 0 00-2 2v2a2 2 0 002 2h2a2 2 0 002-2v-2a2 2 0 00-2-2H5zM11 5a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V5zM14 11a1 1 0 011 1v1h1a1 1 0 110 2h-1v1a1 1 0 11-2 0v-1h-1a1 1 0 110-2h1v-1a1 1 0 011-1z" />
            </svg>
            {data.connections.length} connection{data.connections.length !== 1 ? 's' : ''}
          </span>
        )}
      </div>
    </div>
  );
};

export default memo(ValueNode);