import React, { useState, useMemo, useEffect, useRef } from 'react';
import { Handle, Position, NodeProps, useReactFlow } from 'reactflow';
import { useRopeContext, ProtocolAction } from '../../context/RopeContext';

interface ActionSelectorNodeData {
  id: string;
  onSelectAction?: (id: string, action: ProtocolAction) => void;
  isFromConnection?: boolean;
}

const ActionSelectorNode: React.FC<NodeProps<ActionSelectorNodeData>> = ({ data, id }) => {
  const { 
    protocolActions, 
    isLoadingSchemas,
    protocolsWithMetadata,
    addKnot,
    setNodePosition,
    registerPersistentKnot
  } = useRopeContext();
  
  // Get access to the React Flow instance
  const reactFlowInstance = useReactFlow();
  
  // Add a ref to track if this node is being processed for transformation
  // This prevents duplicate nodes being created when selecting an action
  const isBeingTransformedRef = useRef(false);
  
  const [filterText, setFilterText] = useState('');
  const [selectedProtocol, setSelectedProtocol] = useState<string | null>(null);
  const [selectedType, setSelectedType] = useState<string | null>(null);
  const [selectedChain, setSelectedChain] = useState<number | null>(8453); // Default to Base chain
  // Always expanded - no collapsed state
  const inputRef = useRef<HTMLInputElement>(null);
  
  // Focus the search input on mount immediately
  useEffect(() => {
    if (inputRef.current) {
      inputRef.current.focus();
    }
  }, []);
  
  // Group protocols for dropdown
  const protocols = useMemo(() => {
    const uniqueProtocols = [...new Set(protocolActions.map(pa => pa.protocol))];
    return uniqueProtocols.sort();
  }, [protocolActions]);

  // Get unique chains from all protocols
  const chains = useMemo(() => {
    const uniqueChains = new Set<number>();
    
    Object.values(protocolsWithMetadata).forEach(metadata => {
      if (metadata.chains) {
        metadata.chains.forEach(chain => {
          if (chain.chainIds) {
            chain.chainIds.forEach(id => uniqueChains.add(id));
          }
        });
      }
    });
    
    return Array.from(uniqueChains).sort((a, b) => a - b);
  }, [protocolsWithMetadata]);
  
  // Filter actions by search text, protocol, chain, and type
  const filteredActions = useMemo(() => {
    return protocolActions.filter(action => {
      // Text filter
      const matchesFilter = 
        action.sentence.toLowerCase().includes(filterText.toLowerCase()) ||
        action.protocol.toLowerCase().includes(filterText.toLowerCase()) ||
        action.action.toLowerCase().includes(filterText.toLowerCase());
      
      // Protocol filter
      const matchesProtocol = selectedProtocol ? action.protocol === selectedProtocol : true;
      
      // Type filter
      const matchesType = selectedType ? action.type === selectedType : true;
      
      // Chain filter
      let matchesChain = true;
      if (selectedChain !== null) {
        // Check if this protocol supports the selected chain
        const metadata = protocolsWithMetadata[action.protocol];
        if (metadata?.chains) {
          // Check if any chain in the metadata has the selected chainId
          matchesChain = metadata.chains.some(chain => 
            chain.chainIds?.includes(selectedChain)
          );
        } else {
          // If no chain metadata, we can't filter
          matchesChain = false;
        }
      }
      
      return matchesFilter && matchesProtocol && matchesType && matchesChain;
    }).slice(0, 20); // Limit to 20 results to avoid performance issues
  }, [protocolActions, filterText, selectedProtocol, selectedType, selectedChain, protocolsWithMetadata]);
  
  // Get protocol icon if available
  const getProtocolIcon = (protocol: string) => {
    return protocolsWithMetadata[protocol]?.icon || null;
  };
  
  // Show the correct connection handles based on the connection type
  const isFromConnection = data.isFromConnection;
  
  // Handle wheel events to prevent them from propagating to the ReactFlow canvas
  const handleWheel = (e: React.WheelEvent) => {
    e.stopPropagation();
  };

  return (
    <div 
      className="rounded-lg shadow-lg bg-white border border-indigo-300 flex flex-col p-0 max-h-[500px]"
      style={{ 
        overflow: 'visible', // Make the container allow overflow for the handles
        width: 'fit-content',
        minWidth: '250px',
        position: 'relative'
      }}
      onWheel={handleWheel}
    >
      {/* Connection handle for input */}
      {isFromConnection && (
        <Handle
          type="target"
          position={Position.Left}
          id="continuity"
          className="w-3 h-3 bg-blue-500"
          style={{ 
            zIndex: 10,
            left: '-6px',
            top: '25px'
          }}
        />
      )}
      
      {/* Output handle for connecting to the next node */}
      <Handle
        type="source"
        position={Position.Right}
        id="output"
        className="w-3 h-3 bg-blue-500"
        style={{ 
          zIndex: 10,
          right: '-6px',
          top: '25px'
        }}
      />
      
      {/* Header section */}
      <div className="bg-white rounded-t-lg overflow-hidden">
        <div className="p-3 border-b bg-gradient-to-r from-indigo-50 to-blue-50 flex justify-between items-center">
          <h2 className="text-lg font-semibold text-indigo-800">
            {isFromConnection ? 'Select Next Action' : 'Select Action'}
          </h2>
        </div>
      </div>
      
      {/* Content section with overflow */}
      <div className="flex-1 flex flex-col overflow-hidden rounded-b-lg w-auto">
        <div className="px-3 py-2 flex-1 flex flex-col w-auto">
          {isLoadingSchemas ? (
            <div className="flex-1 flex justify-center items-center p-4">
              <div className="animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-indigo-500"></div>
            </div>
          ) : (
            <>
              <div 
                className="space-y-2 mb-3" 
                onWheel={(e) => e.stopPropagation()}
              >
                <div className="flex flex-col gap-2 w-full">
                  <input
                    ref={inputRef}
                    type="text"
                    placeholder="Search protocols and actions..."
                    className="w-full px-3 py-2 rounded border focus:outline-none focus:ring-2 focus:ring-indigo-500"
                    value={filterText}
                    onChange={(e) => setFilterText(e.target.value)}
                    onWheel={(e) => e.stopPropagation()}
                    style={{ boxSizing: 'border-box', width: '100%' }}
                  />
                </div>
              </div>

              <div 
                className="overflow-y-auto flex-1 space-y-2 max-h-[400px] pr-1 pb-1"
                style={{ width: '100%' }}
                onWheel={(e) => {
                  // Stop wheel events from propagating to the ReactFlow pane
                  e.stopPropagation();
                }}
              >                
                {filteredActions.length > 0 ? (
                  filteredActions.map((action, index) => {
                    // Calculate number of inputs
                    const inputRegex = /\{(\d+)(?:[^}]*)\}/g;
                    const inputMatches = action.sentence.match(inputRegex) || [];
                    const inputCount = inputMatches.length;
                    
                    // Get chain info from metadata
                    const metadata = protocolsWithMetadata[action.protocol];
                    const chainCount = metadata?.chains?.length || 0;
                    
                    // Check if action has coils
                    const hasCoils = action.coils && Object.keys(action.coils).length > 0;
                    
                    return (
                      <button
                        key={`${action.protocol}-${action.action}-${index}`}
                        className="bg-indigo-50 hover:bg-indigo-100 text-indigo-700 text-sm text-left px-2 py-2 rounded transition-colors w-full"
                        style={{ width: '100%', overflowWrap: 'break-word', whiteSpace: 'normal' }}
                        onClick={(e) => {
                          e.preventDefault(); // Prevent default
                          e.stopPropagation(); // Stop event propagation
                          
                          // EXTENSIVE DEBUGGING
                          console.log('------------- ACTION SELECTION START -------------');
                          console.log('Action selected:', action);
                          console.log('Action selector Node ID:', id);
                          console.log('Current event:', e);
                          
                          // Log all current nodes and edges
                          const allCurrentNodes = reactFlowInstance.getNodes();
                          const allCurrentEdges = reactFlowInstance.getEdges();
                          console.log('Current nodes before transform:', allCurrentNodes);
                          console.log('Current edges before transform:', allCurrentEdges);
                          
                          // Get this node's position from ReactFlow
                          const nodePosition = reactFlowInstance.getNode(id)?.position || { x: 0, y: 0 };
                          
                          // Create a new knot ID with timestamp for uniqueness
                          const newKnotId = `knot-${Date.now()}`;
                          
                          console.log('Creating new knot:', newKnotId, 'at position:', nodePosition);
                          
                          // Set position for the new knot
                          setNodePosition(newKnotId, nodePosition);
                          
                          console.log('*** RADICAL APPROACH - DIRECTLY MODIFYING NODE ***');
                          
                          // RADICAL APPROACH: Instead of creating a new node, let's just modify this one!
                          // This should prevent any weird race conditions or double-creation issues
                          
                          console.log('First saving the current node position:', nodePosition);
                          
                          // Save the position explicitly before making any changes
                          // This ensures the position is preserved when the node transforms
                          setNodePosition(id, nodePosition);
                          
                          console.log('Creating a new knot with the original node ID:', id);
                          
                          // Add the knot to the rope context with the ORIGINAL node ID
                          // This ensures the node data exists in the Rope state manager
                          addKnot(action, id);
                          
                          // Register this ID as a persistent knot so it stays as a knot in React Flow
                          registerPersistentKnot(id);
                          
                          // Add a small delay to allow the knot to be properly added to the state
                          // and processed by the Rope system
                          setTimeout(() => {
                            console.log('After delay, getting fresh node data');
                            
                            // DIRECTLY modify this node to be a knot instead of an action selector
                            // We keep the same ID to avoid any edge connection issues
                            const existingNode = reactFlowInstance.getNode(id);
                            
                            if (existingNode) {
                              console.log('Found existing node, transforming it directly:', existingNode);
                              
                              // Create a modified version of the existing node
                              const modifiedNode = {
                                ...existingNode,
                                type: 'knotNode',
                                // Ensure position is preserved from the original node
                                position: existingNode.position,
                                data: { 
                                  knot: {
                                    id: id,
                                    sentence: action.sentence,
                                    protocol: action.protocol,
                                    action: action.action,
                                    // Add other necessary properties that KnotNode might expect
                                    parsed: { inputs: [] }, // Add empty inputs array as placeholder
                                    isComplete: false,
                                    isValid: false
                                  }
                                }
                              };
                              
                              // Log position to debug
                              console.log('Preserving position:', existingNode.position);
                              
                              console.log('About to replace node with modified version:', modifiedNode);
                              
                              // Replace just this one node, keeping all others
                              reactFlowInstance.setNodes(nodes => {
                                console.log('Current nodes before replacement:', nodes.map(n => n.id));
                                return nodes.map(node => node.id === id ? modifiedNode : node);
                              });
                              
                              console.log('Node directly transformed to knot type');
                              
                              // We don't need to force a re-render now, and it was causing position changes
                              // Removed the fitView call as it may have been repositioning nodes
                            }
                          }, 10);
                          
                          // Stop processing here since we're doing the transformation asynchronously
                          return;
                          
                          // If we couldn't find the node (unlikely), fall back to the old approach
                          console.log('Could not find existing node, falling back to replacement approach');
                          
                          // Add the knot to rope context with the new ID
                          addKnot(action, newKnotId);
                          
                          // Create the knot node data
                          const newKnotNode = {
                            id: newKnotId,
                            type: 'knotNode',
                            position: nodePosition,
                            data: { 
                              knot: {
                                id: newKnotId,
                                sentence: action.sentence,
                                protocol: action.protocol,
                                action: action.action,
                              }
                            },
                          };
                          
                          // This is the fallback code path only used if direct transformation fails
                          // It should never be reached in normal operation
                          
                          const isFromConnection = data.isFromConnection || false;
                          console.log('Fallback code path - using old node replacement logic');
                          
                          // Find all current edges
                          const allEdges = reactFlowInstance.getEdges();
                          
                          // Handle incoming connections 
                          const edgesToUpdate = allEdges.filter(edge => edge.target === id);
                          if (edgesToUpdate.length > 0) {
                            const newEdges = edgesToUpdate.map(edge => ({
                              ...edge,
                              id: `${edge.id}-updated`,
                              target: newKnotId
                            }));
                            reactFlowInstance.addEdges(newEdges);
                          }
                          
                          // Handle outgoing connections
                          const outgoingEdges = allEdges.filter(edge => edge.source === id);
                          if (outgoingEdges.length > 0) {
                            const newOutEdges = outgoingEdges.map(edge => ({
                              ...edge,
                              id: `${edge.id}-out-updated`,
                              source: newKnotId
                            }));
                            reactFlowInstance.addEdges(newOutEdges);
                          }
                          
                          // Replace old node with new one
                          reactFlowInstance.setNodes(nodes => 
                            [...nodes.filter(n => n.id !== id), newKnotNode]
                          );
                          
                          // Clean up edges
                          reactFlowInstance.setEdges(edges => 
                            edges.filter(e => e.source !== id && e.target !== id)
                          );
                          
                          console.log('Fallback approach complete');
                          
                          // Set the transformation flag to prevent new nodes being created
                          isBeingTransformedRef.current = true;
                          
                          // Log final state after transformation
                          setTimeout(() => {
                            console.log('------------- ACTION SELECTION COMPLETE -------------');
                            const nodesAfter = reactFlowInstance.getNodes();
                            const edgesAfter = reactFlowInstance.getEdges();
                            console.log('Nodes after transform:', nodesAfter);
                            console.log('Edges after transform:', edgesAfter);
                            console.log('Is this node still present?', nodesAfter.some(n => n.id === id));
                            console.log('New knot node present?', nodesAfter.some(n => n.id === newKnotId));
                            console.log('--------------------------------------------------');
                          }, 100);
                          
                          // Don't call the parent callback - we're handling everything internally now
                          // data.onSelectAction(id, action);
                        }}
                      >
                        <div className="flex items-center gap-2 mb-1">
                          {getProtocolIcon(action.protocol) && (
                            <img 
                              src={getProtocolIcon(action.protocol)} 
                              alt={action.protocol} 
                              className="w-5 h-5 rounded-full"
                            />
                          )}
                          <div className="flex flex-col">
                            <div className="flex items-center gap-1">
                              <span className="font-medium">{action.protocol} / {action.action}</span>
                              {action.type && (
                                <span className="text-xs bg-indigo-100 text-indigo-800 px-1.5 py-0.5 rounded">
                                  {action.type}
                                </span>
                              )}
                            </div>
                          </div>
                        </div>
                        
                        <div className="text-xs text-gray-600">
                          {action.sentence}
                        </div>
                        
                        <div className="flex items-center gap-2 mt-2 text-xs">
                          <span className="bg-gray-100 text-gray-700 px-1.5 py-0.5 rounded-full flex items-center">
                            <svg className="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                            </svg>
                            {inputCount} input{inputCount !== 1 ? 's' : ''}
                          </span>
                          
                          {hasCoils && (
                            <span className="bg-blue-100 text-blue-700 px-1.5 py-0.5 rounded-full flex items-center">
                              <svg className="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                              </svg>
                              Has coils
                            </span>
                          )}
                          
                          {chainCount > 0 && (
                            <span className="bg-green-100 text-green-700 px-1.5 py-0.5 rounded-full flex items-center">
                              <svg className="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9" />
                              </svg>
                              {chainCount} chain{chainCount !== 1 ? 's' : ''}
                            </span>
                          )}
                        </div>
                      </button>
                    );
                  })
                ) : (
                  <div className="text-gray-500 text-sm p-4 text-center bg-gray-50 rounded-lg">
                    No actions found matching your criteria.
                    <div className="mt-2">
                      <button 
                        className="text-indigo-600 hover:text-indigo-800 underline text-xs"
                        onClick={() => {
                          setFilterText('');
                          setSelectedProtocol(null);
                          setSelectedType(null);
                        }}
                      >
                        Clear filters
                      </button>
                    </div>
                  </div>
                )}
                
                {filteredActions.length > 0 && filteredActions.length === 20 && (
                  <div className="text-xs text-gray-500 text-center py-1">
                    Showing first 20 results. Refine your search to see more.
                  </div>
                )}
              </div>
            </>
          )}
        </div>
      </div>
    </div>
  );
};

export default ActionSelectorNode;