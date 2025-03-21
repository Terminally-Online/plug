import React, { useState, useMemo, useEffect, useRef } from 'react';
import { useRopeContext, ProtocolAction } from '../context/RopeContext';

interface ActionSelectorProps {
  isOpen: boolean;
  onClose: () => void;
  onSelectAction: (action: ProtocolAction, position: { x: number, y: number }) => void;
  position: { x: number, y: number };
  isFromConnection?: boolean;
}

const ActionSelector: React.FC<ActionSelectorProps> = ({
  isOpen,
  onClose,
  onSelectAction,
  position,
  isFromConnection = false
}) => {
  const { 
    protocolActions, 
    isLoadingSchemas,
    protocolsWithMetadata
  } = useRopeContext();
  
  const [filterText, setFilterText] = useState('');
  const [selectedProtocol, setSelectedProtocol] = useState<string | null>(null);
  const [selectedType, setSelectedType] = useState<string | null>(null);
  const [selectedChain, setSelectedChain] = useState<number | null>(8453); // Default to Base chain
  const modalRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  
  // Focus input on open immediately
  useEffect(() => {
    if (isOpen && inputRef.current) {
      // Focus immediately without delay
      inputRef.current.focus();
    }
  }, [isOpen]);

  // Close on escape key
  useEffect(() => {
    const handleEscape = (e: KeyboardEvent) => {
      if (e.key === 'Escape') onClose();
    };
    
    if (isOpen) {
      document.addEventListener('keydown', handleEscape);
    }
    
    return () => {
      document.removeEventListener('keydown', handleEscape);
    };
  }, [isOpen, onClose]);
  
  // Click outside to close
  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (modalRef.current && !modalRef.current.contains(e.target as Node)) {
        onClose();
      }
    };
    
    if (isOpen) {
      document.addEventListener('mousedown', handleClickOutside);
    }
    
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [isOpen, onClose]);
  
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
    });
  }, [protocolActions, filterText, selectedProtocol, selectedType, selectedChain, protocolsWithMetadata]);
  
  // Get protocol icon if available
  const getProtocolIcon = (protocol: string) => {
    return protocolsWithMetadata[protocol]?.icon || null;
  };
  
  if (!isOpen) return null;
  
  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div 
        ref={modalRef}
        className="bg-white rounded-lg shadow-xl max-w-2xl w-full max-h-[80vh] flex flex-col"
        style={{ 
          position: 'absolute', 
          left: `${position.x}px`, 
          top: `${position.y}px`,
          transform: 'translate(-50%, -50%)'
        }}
      >
        <div className="p-4 border-b flex justify-between items-center">
          <h2 className="text-lg font-semibold text-gray-800">
            {isFromConnection ? 'Connect to Action' : 'Add New Action'}
          </h2>
          <button 
            onClick={onClose}
            className="text-gray-400 hover:text-gray-600"
          >
            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        
        <div className="p-4 overflow-hidden flex-1 flex flex-col">
          {isLoadingSchemas ? (
            <div className="flex-1 flex justify-center items-center p-4">
              <div className="animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-blue-500"></div>
            </div>
          ) : (
            <>
              <div className="space-y-2 mb-4">
                <div className="flex flex-col gap-2">
                  <input
                    ref={inputRef}
                    type="text"
                    placeholder="Search protocols and actions..."
                    className="w-full px-3 py-2 rounded border focus:outline-none focus:ring-2 focus:ring-blue-500"
                    value={filterText}
                    onChange={(e) => setFilterText(e.target.value)}
                  />
                  
                  <div className="grid grid-cols-3 gap-2">
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
              </div>

              <div className="overflow-y-auto flex-1 pr-1">
                <h3 className="text-sm text-gray-500 font-medium mb-2">Available Actions ({filteredActions.length})</h3>
                
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
                          onClick={() => onSelectAction(action, position)}
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
      </div>
    </div>
  );
};

export default ActionSelector;