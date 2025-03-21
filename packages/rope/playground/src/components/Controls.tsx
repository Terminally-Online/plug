import React, { useState, useMemo } from 'react';
import { useRopeContext } from '../context/RopeContext';
import { ProtocolAction } from '../context/RopeContext';

interface ControlsProps {
  onAddNode: () => void;
}

const Controls: React.FC<ControlsProps> = ({ onAddNode }) => {
  const { 
    knots, 
    removeKnot, 
    addKnot, 
    setNodePosition, 
    protocolActions, 
    isLoadingSchemas,
    protocolsWithMetadata,
    errorMessage
  } = useRopeContext();
  
  const [selectedKnotId, setSelectedKnotId] = useState<string | null>(null);
  const [filterText, setFilterText] = useState('');
  const [selectedProtocol, setSelectedProtocol] = useState<string | null>(null);
  const [showActionSelector, setShowActionSelector] = useState(false);
  
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
  
  // State for chain filtering
  const [selectedChain, setSelectedChain] = useState<number | null>(null);
  const [selectedType, setSelectedType] = useState<string | null>(null);

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
    });
  }, [protocolActions, filterText, selectedProtocol, selectedType, selectedChain, protocolsWithMetadata]);
  
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
  
  const handleAddAction = (protocolAction: ProtocolAction) => {
    // Create a custom ID
    const newKnotId = `knot-${Date.now()}`;
    
    // Place new knots in a reasonable starting position in the viewport
    const position = { x: 300, y: 200 };
    setNodePosition(newKnotId, position);
    
    // Add the custom template with the custom ID
    setTimeout(() => {
      addKnot(protocolAction, newKnotId);
      setShowActionSelector(false);
    }, 0);
  };
  
  // Get protocol icon if available
  const getProtocolIcon = (protocol: string) => {
    return protocolsWithMetadata[protocol]?.icon || null;
  };
  
  return (
    <div className="w-80 bg-gray-100 p-4 overflow-y-auto flex flex-col">
      <h2 className="text-lg font-semibold mb-4">Rope Editor Controls</h2>
      
      <div className="mb-4">
        <button 
          className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded w-full"
          onClick={() => setShowActionSelector(!showActionSelector)}
        >
          {showActionSelector ? 'Hide Action Selector' : 'Show Available Actions'}
        </button>
      </div>
      
      {showActionSelector && (
        <div className="mb-4">
          {isLoadingSchemas ? (
            <div className="flex justify-center items-center p-4">
              <div className="animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-blue-500"></div>
            </div>
          ) : errorMessage ? (
            <div className="p-4 border border-red-200 bg-red-50 rounded-lg">
              <div className="flex items-start">
                <div className="flex-shrink-0">
                  <svg className="h-5 w-5 text-red-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                    <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                  </svg>
                </div>
                <div className="ml-3">
                  <h3 className="text-sm font-medium text-red-800">API Connection Error</h3>
                  <div className="mt-2 text-xs text-red-700">
                    <div className="bg-red-50 border border-red-200 rounded-md p-2 mb-2 overflow-auto max-h-32">
                      <code className="whitespace-pre-wrap break-all">{errorMessage}</code>
                    </div>
                    
                    <p className="mt-2 font-semibold">Troubleshooting Guide:</p>
                    
                    {errorMessage?.includes('500') && (
                      <div className="bg-yellow-50 border border-yellow-200 rounded-md p-2 mb-2">
                        <p className="font-medium text-yellow-800">⚠️ Server Error (500)</p>
                        <p className="mt-1">This indicates an error in the local API server. Check the server logs for more details.</p>
                      </div>
                    )}
                    
                    <ul className="list-disc pl-5 mt-1 space-y-1">
                      <li>Server is running at <code className="bg-red-50 px-1 rounded">http://localhost:8080</code></li>
                      <li>API key <code className="bg-red-50 px-1 rounded">alphapapapapaalphapapaindia</code> is being sent as <code className="bg-red-50 px-1 rounded">X-Api-Key</code> header</li>
                      <li>Server endpoint <code className="bg-red-50 px-1 rounded">/solver</code> exists and is properly configured</li>
                    </ul>
                    
                    <p className="mt-2 font-medium">Request Details:</p>
                    <div className="bg-gray-50 border border-gray-200 rounded-md p-2 mt-1 font-mono text-xs overflow-auto">
                      <div className="font-semibold mb-1">Working Postman Example:</div>
                      GET http://localhost:8080/solver?chainId=8453<br />
                      X-Api-Key: alphapapapapaalphapapaindia
                    </div>
                    
                    <p className="mt-2 text-sm italic">The playground uses a Vite proxy to avoid CORS issues. You may need to restart your dev server for changes to take effect.</p>
                  </div>
                  <div className="mt-3 flex gap-2">
                    <button
                      onClick={() => window.location.reload()}
                      className="text-xs font-medium text-red-800 hover:text-red-900 bg-red-100 hover:bg-red-200 px-2 py-1 rounded"
                    >
                      Retry
                    </button>
                    
                    <button
                      onClick={async () => {
                        // Show info to user
                        alert("Attempting direct fetch using Postman example configuration. This may fail due to CORS restrictions, but worth trying...");
                        
                        try {
                          const response = await fetch('http://localhost:8080/solver?chainId=8453', {
                            method: 'GET',
                            headers: {
                              'X-Api-Key': 'alphapapapapaalphapapaindia'
                            },
                            mode: 'cors' // Try with CORS mode
                          });
                          
                          if (response.ok) {
                            alert("Success! The direct request worked. Reloading the page...");
                            window.location.reload();
                          } else {
                            const text = await response.text();
                            alert(`Failed with status ${response.status}: ${text.substring(0, 100)}`);
                          }
                        } catch (error) {
                          alert(`Failed due to CORS restriction: ${error.message}`);
                          
                          // Offer the cors-anywhere option as fallback
                          if (confirm("Would you like to try with a CORS proxy?")) {
                            window.open('https://cors-anywhere.herokuapp.com/corsdemo', '_blank');
                          }
                        }
                      }}
                      className="text-xs font-medium text-blue-800 hover:text-blue-900 bg-blue-100 hover:bg-blue-200 px-2 py-1 rounded"
                    >
                      Try Direct Request
                    </button>
                  </div>
                </div>
              </div>
            </div>
          ) : (
            <>
              <div className="space-y-2 mb-4">
                <div className="flex flex-col gap-2">
                  <input
                    type="text"
                    placeholder="Search protocols and actions..."
                    className="w-full px-3 py-2 rounded border focus:outline-none focus:ring-2 focus:ring-blue-500"
                    value={filterText}
                    onChange={(e) => setFilterText(e.target.value)}
                  />
                  
                  <div className="grid grid-cols-2 gap-2">
                    <select
                      className="w-full px-3 py-2 rounded border focus:outline-none focus:ring-2 focus:ring-blue-500"
                      value={selectedProtocol || ''}
                      onChange={(e) => setSelectedProtocol(e.target.value || null)}
                    >
                      <option value="">All Protocols</option>
                      {protocols.map(protocol => (
                        <option key={protocol} value={protocol}>
                          {protocol}
                        </option>
                      ))}
                    </select>
                    
                    <select
                      className="w-full px-3 py-2 rounded border focus:outline-none focus:ring-2 focus:ring-blue-500"
                      value={selectedType || ''}
                      onChange={(e) => setSelectedType(e.target.value || null)}
                    >
                      <option value="">All Types</option>
                      <option value="action">Actions</option>
                      <option value="trigger">Triggers</option>
                      <option value="condition">Conditions</option>
                    </select>
                  </div>
                  
                  <select
                    className="w-full px-3 py-2 rounded border focus:outline-none focus:ring-2 focus:ring-blue-500"
                    value={selectedChain?.toString() || ''}
                    onChange={(e) => setSelectedChain(e.target.value ? parseInt(e.target.value) : null)}
                  >
                    <option value="">All Chains</option>
                    {chains.map(chainId => {
                      // Find the first protocol with this chain to get the name
                      let chainName = `Chain ID: ${chainId}`;
                      
                      // Try to find a readable name for this chain
                      for (const metadata of Object.values(protocolsWithMetadata)) {
                        if (metadata.chains) {
                          const chain = metadata.chains.find(c => c.chainIds?.includes(chainId));
                          if (chain && chain.name) {
                            chainName = chain.name;
                            break;
                          }
                        }
                      }
                      
                      return (
                        <option key={chainId} value={chainId.toString()}>
                          {chainName}
                        </option>
                      );
                    })}
                  </select>
                </div>
              </div>

              <div className="border rounded bg-white p-2 max-h-60 overflow-y-auto">
                <div className="text-xs text-gray-500 font-medium mb-2">Available Actions ({filteredActions.length})</div>
                <div className="grid grid-cols-1 gap-2">
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
                          className="bg-indigo-50 hover:bg-indigo-100 text-indigo-700 text-sm text-left px-3 py-2 rounded transition-colors w-full"
                          onClick={() => handleAddAction(action)}
                          title={`${action.protocol} / ${action.action}`}
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
                    <div className="text-gray-500 text-sm p-2 text-center">
                      No actions found matching your criteria
                    </div>
                  )}
                </div>
              </div>
            </>
          )}
        </div>
      )}
      
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
          <li>Drag from <span className="inline-block w-2 h-2 bg-blue-500 rounded-full"></span> output handle to create connections</li>
          <li>Press 'Delete' key to remove nodes</li>
        </ul>
      </div>
    </div>
  );
};

export default Controls;