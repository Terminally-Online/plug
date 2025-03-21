import { memo, useState, useMemo } from 'react';
import { Handle, Position, NodeProps } from 'reactflow';
import { useRopeContext } from '../../context/RopeContext';

const KnotNode = ({ data }: NodeProps) => {
  const { updateKnotValue, protocolsWithMetadata } = useRopeContext();
  const { knot } = data;
  
  // Get protocol metadata if available
  const protocolMetadata = useMemo(() => {
    if (knot?.protocol) {
      return protocolsWithMetadata[knot.protocol];
    }
    return null;
  }, [knot?.protocol, protocolsWithMetadata]);
  
  // Get protocol icon if available
  const protocolIcon = useMemo(() => {
    return protocolMetadata?.icon || null;
  }, [protocolMetadata]);
  
  // Extract information about inputs from the parsed sentence
  const inputs = knot.parsed?.inputs || [];
  
  // Create a map to track input values
  const [localValues, setLocalValues] = useState<Record<string, string>>({});
  
  // Function to handle input changes
  const handleInputChange = (index: number, value: string) => {
    setLocalValues(prev => ({ ...prev, [index]: value }));
    if (knot && knot.id) {
      updateKnotValue(knot.id, index, value);
    }
  };
  
  // Determine if the node is valid
  const isComplete = knot?.isComplete || false;
  const isValid = knot?.isValid || false;
  
  // Get the border and styling based on validation status
  const getBorderColor = () => {
    if (!isComplete) return 'border-yellow-300';
    return isValid ? 'border-green-300' : 'border-red-300';
  };
  
  if (!knot) {
    return (
      <div className="p-3 rounded-lg shadow-md bg-gradient-to-br from-white to-gray-50 border border-gray-300">
        <div className="text-gray-500 text-sm">Missing knot data</div>
        <Handle
          type="source"
          position={Position.Right}
          id="output"
          className="w-3 h-3 bg-blue-500 !right-[-5px] !top-[12.5%]"
        />
      </div>
    );
  }

  // Calculate the vertical positions for input handles
  const getInputHandlePosition = (index: number, total: number) => {
    // Reserve the first 25% for the continuity handle
    // And distribute the remaining 75% among the input handles
    const reservedTopPercentage = 0.25;
    
    // If there's only one input, place it in the middle of the remaining space
    if (total === 1) return reservedTopPercentage + ((1 - reservedTopPercentage) / 2);
    
    // Otherwise distribute evenly in the remaining space
    const remainingSpace = 1 - reservedTopPercentage;
    const step = remainingSpace / (total + 1);
    return reservedTopPercentage + (step * (index + 1));
  };
  
  // Prevent scroll events from propagating to the ReactFlow pane
  const handleWheel = (e: React.WheelEvent) => {
    e.stopPropagation();
  };

  return (
    <div 
      className={`p-3 rounded-lg shadow-md bg-gradient-to-br from-white to-blue-50 border ${getBorderColor()} max-w-md`}
      onWheel={handleWheel}>
      {/* Continuity handle on the left - for connecting from previous knots */}
      <Handle
        type="target"
        position={Position.Left}
        id="continuity"
        className="w-3 h-3 bg-blue-500 !left-[-5px] !top-[12.5%]"
        title="Connect from previous knot"
      />
      
      {/* Input handles on the left side - one for each input parameter */}
      {inputs.length > 0 && inputs.map((input, idx) => (
        <Handle
          key={`input-handle-${input.index}`}
          type="target"
          position={Position.Left}
          id={`input-${input.index}`}
          className="w-3 h-3 bg-green-500 !left-[-5px]"
          // Position the handles evenly along the left side
          style={{ top: `${getInputHandlePosition(idx, inputs.length) * 100}%` }}
          // Add a data attribute to identify this as an input handle
          data-input-index={input.index}
        />
      ))}
      
      {/* Output handle on the right side for the next node connection */}
      <Handle
        type="source"
        position={Position.Right}
        id="output"
        className="w-3 h-3 bg-blue-500 !right-[-5px] !top-[12.5%]"
        title="Drag to create a new knot"
      />
      
      <div className="mb-3">
        {/* Protocol and action header */}
        <div className="flex items-center mb-2">
          {protocolIcon && (
            <img 
              src={protocolIcon} 
              alt={knot.protocol} 
              className="w-5 h-5 rounded-full mr-2"
            />
          )}
          <div className="flex flex-col">
            <div className="flex items-center gap-1">
              <span className="font-medium text-sm text-gray-700">{knot.protocol}</span>
              <span className="text-gray-400 text-xs">/</span>
              <span className="font-medium text-sm text-gray-700">{knot.action}</span>
            </div>
            {/* Show available chains if metadata exists */}
            {protocolMetadata?.chains && protocolMetadata.chains.length > 0 && (
              <div className="flex items-center gap-1 mt-0.5">
                <span className="text-xs text-gray-500">Chains:</span>
                <div className="flex -space-x-1">
                  {protocolMetadata.chains.slice(0, 3).map((chain, idx) => (
                    <div 
                      key={idx} 
                      className="w-4 h-4 rounded-full border border-white bg-gray-100 flex items-center justify-center overflow-hidden"
                      title={chain.name}
                    >
                      {chain.icon?.default ? (
                        <img src={chain.icon.default} alt={chain.name || 'Chain'} className="w-full h-full object-cover" />
                      ) : (
                        <span className="text-[8px] text-gray-600">{chain.chainIds?.[0].toString().substring(0, 2) || '?'}</span>
                      )}
                    </div>
                  ))}
                  {protocolMetadata.chains.length > 3 && (
                    <div 
                      className="w-4 h-4 rounded-full border border-white bg-gray-200 flex items-center justify-center"
                      title={`${protocolMetadata.chains.length - 3} more chains`}
                    >
                      <span className="text-[8px] text-gray-600">+{protocolMetadata.chains.length - 3}</span>
                    </div>
                  )}
                </div>
              </div>
            )}
          </div>
        </div>
        
        {/* Sentence display */}
        <div className="font-medium text-gray-800 border-l-2 border-blue-300 pl-2 py-1">
          {knot.sentence || 'No sentence'}
        </div>
      </div>
      
      {inputs.length > 0 && (
        <div className="space-y-2 mt-3">
          {inputs.map((input) => (
            <div key={input.index} className="flex items-center group">
              <div 
                className="w-2 h-2 bg-green-500 rounded-full mr-2 group-hover:animate-pulse" 
                title={`Input ${input.index} - drag from left handle to create a value node`}
              />
              <span className="text-xs text-gray-600 mr-2 whitespace-nowrap">{input.index}:</span>
              <input
                type="text"
                className="flex-1 px-2 py-1 text-sm border border-gray-200 rounded-md focus:outline-none focus:ring-1 focus:ring-green-500 focus:border-green-500"
                placeholder="Enter value..."
                value={localValues[input.index] || (knot.values?.get(input.index)?.value || '')}
                onChange={(e) => handleInputChange(input.index, e.target.value)}
              />
            </div>
          ))}
        </div>
      )}
      
      {knot.resolvedSentence && (
        <div className="mt-3 p-2 bg-blue-50 border border-blue-100 rounded-md">
          <div className="text-xs font-semibold text-blue-700 mb-1 flex items-center">
            <svg className="w-3 h-3 mr-1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
              <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
              <path fillRule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clipRule="evenodd" />
            </svg>
            Resolved
          </div>
          <div className="text-sm font-medium text-blue-900">{knot.resolvedSentence}</div>
        </div>
      )}
      
      <div className="mt-2 flex space-x-2">
        <span className={`text-xs px-2 py-0.5 rounded-full flex items-center ${
          isComplete 
            ? 'bg-green-100 text-green-800 border border-green-200' 
            : 'bg-yellow-100 text-yellow-800 border border-yellow-200'
        }`}>
          <span className={`w-1.5 h-1.5 rounded-full mr-1 ${
            isComplete ? 'bg-green-500' : 'bg-yellow-500'
          }`}></span>
          {isComplete ? 'Complete' : 'Incomplete'}
        </span>
        {isComplete && (
          <span className={`text-xs px-2 py-0.5 rounded-full flex items-center ${
            isValid 
              ? 'bg-green-100 text-green-800 border border-green-200' 
              : 'bg-red-100 text-red-800 border border-red-200'
          }`}>
            <span className={`w-1.5 h-1.5 rounded-full mr-1 ${
              isValid ? 'bg-green-500' : 'bg-red-500'
            }`}></span>
            {isValid ? 'Valid' : 'Invalid'}
          </span>
        )}
      </div>
    </div>
  );
};

export default memo(KnotNode);