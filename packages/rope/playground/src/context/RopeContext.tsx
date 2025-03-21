import React, { createContext, useContext, useState, useCallback, useMemo, useEffect } from 'react';
import { useRope } from '../../../src/hooks/useRope';
import { PlugAPI } from '../../../src/models/api';
import * as Types from '../../../src/types/generated';
import axios from 'axios';

type KnotValue = {
  knotId: string;
  inputIndex: number;
  value: string;
};

// Define types for protocols and actions
export interface ProtocolAction {
  protocol: string;
  action: string;
  sentence: string;
  type?: string; // action, trigger, condition, etc.
  coils?: Record<string, string>;
  options?: Record<string, Types.ActionsOptions>;
}

type RopeContextType = {
  knots: any[];
  nodePositions: Record<string, { x: number, y: number }>;
  setNodePosition: (id: string, position: { x: number, y: number }) => void;
  addKnot: (protocolAction: ProtocolAction, customId?: string) => void;
  removeKnot: (id: string) => void;
  updateKnotValue: (knotId: string, inputIndex: number, value: string) => void;
  protocolActions: ProtocolAction[];
  isLoadingSchemas: boolean;
  protocolsWithMetadata: Record<string, Types.ActionsProtocolMetadata>;
  errorMessage: string | null;
  // Function to register a node ID as a persistent knot
  // This is used to prevent transformed action selectors from reverting
  registerPersistentKnot: (id: string) => void;
  // Get the list of persistent knot IDs
  getPersistentKnotIds: () => string[];
};

const RopeContext = createContext<RopeContextType | null>(null);

// Production API key 
const API_KEY = 'alphapapapapaalphapapaindia';

interface RopeProviderProps {
  children: React.ReactNode;
}

export const RopeProvider: React.FC<RopeProviderProps> = ({ children }) => {
  const [nodePositions, setNodePositions] = useState<Record<string, { x: number, y: number }>>({});
  const [protocolsWithMetadata, setProtocolsWithMetadata] = useState<Record<string, Types.ActionsProtocolMetadata>>({});
  const [protocolActions, setProtocolActions] = useState<ProtocolAction[]>([]);
  const [isLoadingSchemas, setIsLoadingSchemas] = useState(true);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  
  // Track node IDs that should always be treated as knots
  // This is used for action selectors that have been transformed
  const [persistentKnotIds, setPersistentKnotIds] = useState<string[]>([]);
  
  // Create a modified API client that uses the local proxy
  const apiClient = useMemo(() => {
    // Create standard API client with debug logging
    console.log('Creating API client with key:', API_KEY);
    
    // Apply header override to ensure X-Api-Key is used
    const originalCreateClient = axios.create;
    axios.create = (config) => {
      console.log('Intercepted axios.create with config:', config);
      // Override headers to use X-Api-Key
      config.headers = {
        ...config.headers,
        'X-Api-Key': API_KEY,
      };
      delete config.headers['Authorization']; // Remove any Authorization header
      
      console.log('Modified axios config headers:', config.headers);
      return originalCreateClient(config);
    };
    
    // Create client with the modified axios
    const client = new PlugAPI(API_KEY);
    
    // Restore original axios.create
    axios.create = originalCreateClient;
    
    // Override the getSolver method to use our local proxy
    const originalGetSolver = client.getSolver.bind(client);
    client.getSolver = async (queryParams?: Types.SchemaQueryParams) => {
      console.log('Intercepted getSolver call, using local proxy', queryParams);
      
      try {
        // Build the URL with query parameters, matching the Postman example
        let url = '/solver';
        
        // Add specific chainId parameter that's known to work
        url += '?chainId=8453';
        
        console.log('Using URL with query params:', url);
        
        // Use fetch with our local proxy instead of the original API
        const response = await fetch(url, {
          method: 'GET',
          headers: {
            'X-Api-Key': API_KEY
          }
        });
        
        // Log response status and headers for debugging
        console.log('Proxy response status:', response.status);
        console.log('Proxy response headers:', Object.fromEntries([...response.headers.entries()]));
        
        if (!response.ok) {
          // For 500 errors, try to get the detailed error message from the response
          const errorText = await response.text();
          console.error('Server error response:', errorText);
          
          let errorMsg = `HTTP error ${response.status}: ${response.statusText}`;
          if (errorText) {
            try {
              // Try to parse the error as JSON if possible
              const errorJson = JSON.parse(errorText);
              if (errorJson.message) {
                errorMsg += ` - ${errorJson.message}`;
              } else {
                errorMsg += ` - ${errorText.substring(0, 100)}`;
              }
            } catch (e) {
              // If it's not JSON, include the first part of the error text
              errorMsg += ` - ${errorText.substring(0, 100)}`;
            }
          }
          
          throw new Error(errorMsg);
        }
        
        // Parse the JSON response
        let data;
        try {
          const responseText = await response.text();
          console.log('Raw response text:', responseText.substring(0, 200) + '...');
          data = JSON.parse(responseText);
          console.log('Parsed response data:', data ? Object.keys(data).length : 'empty', 'protocols');
        } catch (error) {
          console.error('Error parsing response JSON:', error);
          throw new Error(`Failed to parse API response: ${error.message}`);
        }
        
        return data;
      } catch (error) {
        console.error('Proxy fetch error:', error);
        // Try a direct axios call matching the Postman example exactly
        console.log('Trying direct axios call matching Postman example');
        
        try {
          // Use the exact same configuration as the Postman example
          const axiosConfig = {
            method: 'get',
            maxBodyLength: Infinity,
            url: 'http://localhost:8080/solver?chainId=8453',
            headers: { 
              'X-Api-Key': 'alphapapapapaalphapapaindia'
            }
          };
          
          console.log('Making axios request with config:', axiosConfig);
          const axiosResponse = await axios.request(axiosConfig);
          console.log('Axios request succeeded!', Object.keys(axiosResponse.data).length, 'protocols');
          return axiosResponse.data;
        } catch (axiosError) {
          console.error('Axios direct request also failed:', axiosError);
          console.log('Falling back to original implementation as last resort');
          return originalGetSolver(queryParams);
        }
      }
    };
    
    return client;
  }, []);
  
  // Initialize with empty knots
  const { state, actions } = useRope({ 
    initialKnots: [] 
  });

  // Fetch schemas from API
  useEffect(() => {
    const fetchSchemas = async () => {
      try {
        setIsLoadingSchemas(true);
        setErrorMessage(null);
        console.log('Fetching schemas from API using proxy...');
        const schemas = await apiClient.getSolver();
        
        if (!schemas || Object.keys(schemas).length === 0) {
          throw new Error('Received empty schema data from API');
        }
        
        console.log('Schemas fetched successfully:', Object.keys(schemas).length, 'protocols');
        
        // Process schemas into a flattened list of protocol actions
        const actions: ProtocolAction[] = [];
        const metadata: Record<string, Types.ActionsProtocolMetadata> = {};
        
        // Extract all protocol/action combinations with their sentences
        Object.entries(schemas).forEach(([protocolName, protocolData]) => {
          // Store metadata for later use
          if (protocolData.metadata) {
            metadata[protocolName] = protocolData.metadata;
          }
          
          // Process each action in this protocol
          if (protocolData.schema) {
            Object.entries(protocolData.schema).forEach(([actionName, actionSchema]) => {
              if (actionSchema.sentence) {
                actions.push({
                  protocol: protocolName,
                  action: actionName,
                  sentence: actionSchema.sentence,
                  coils: actionSchema.coils,
                  options: actionSchema.options,
                  type: actionSchema.type
                });
              }
            });
          }
        });
        
        console.log('Processed actions:', actions.length);
        if (actions.length > 0) {
          console.log('Sample action:', actions[0]);
        }
        
        setProtocolActions(actions);
        setProtocolsWithMetadata(metadata);
      } catch (error: any) {
        console.error('Failed to fetch schemas:', error);
        
        // Set a detailed error message for display
        let errorMsg = 'Failed to fetch protocols from API.';
        
        if (error.message) {
          errorMsg += ` Error: ${error.message}`;
        }
        
        // Add more details for specific error types
        if (error.name === 'AxiosError' && error.code === 'ERR_NETWORK') {
          errorMsg += ' This may be due to CORS restrictions or the local server not running at http://localhost:8080.';
        }
        
        // For 500 errors, add specific guidance
        if (error.message && error.message.includes('500')) {
          errorMsg += `\n\nServer returned a 500 error. This usually indicates:
1. The API key may be invalid or misconfigured
2. The server is running but encountering an internal error
3. The endpoint format or parameters may be incorrect

Check the server logs for details.`;
        }
        
        setErrorMessage(errorMsg);
      } finally {
        setIsLoadingSchemas(false);
      }
    };
    
    fetchSchemas();
  }, [apiClient]);

  const setNodePosition = useCallback((id: string, position: { x: number, y: number }) => {
    setNodePositions(prev => ({
      ...prev,
      [id]: position
    }));
  }, []);

  const addKnot = useCallback((protocolAction: ProtocolAction, customId?: string) => {
    // Use provided ID or generate a new one
    const id = customId || `knot-${Date.now()}`;
    
    // Add the knot to the system
    actions.addKnot({
      id,
      sentence: protocolAction.sentence,
      protocol: protocolAction.protocol,
      action: protocolAction.action,
    });
  }, [actions]);

  const removeKnot = useCallback((id: string) => {
    actions.removeKnot(id);
    
    // Remove position data
    setNodePositions(prev => {
      const newPositions = { ...prev };
      delete newPositions[id];
      return newPositions;
    });
  }, [actions]);

  const updateKnotValue = useCallback((knotId: string, inputIndex: number, value: string) => {
    actions.setValue(knotId, inputIndex, value);
  }, [actions]);

  // Convert the knots object to an array for easier rendering
  const knots = useMemo(() => {
    return Object.entries(state.knots || {}).map(([id, knot]) => ({
      id,
      ...knot
    }));
  }, [state.knots]);

  // Function to register a node ID as a persistent knot
  const registerPersistentKnot = useCallback((id: string) => {
    console.log('Registering persistent knot:', id);
    setPersistentKnotIds(prev => {
      if (prev.includes(id)) return prev;
      return [...prev, id];
    });
  }, []);
  
  // Function to get the list of persistent knot IDs
  const getPersistentKnotIds = useCallback(() => {
    return persistentKnotIds;
  }, [persistentKnotIds]);
  
  // Provide the context values to children
  const contextValue = {
    knots,
    nodePositions,
    setNodePosition,
    addKnot,
    removeKnot,
    updateKnotValue,
    protocolActions,
    isLoadingSchemas,
    protocolsWithMetadata,
    errorMessage,
    registerPersistentKnot,
    getPersistentKnotIds
  };

  return (
    <RopeContext.Provider value={contextValue}>
      {children}
    </RopeContext.Provider>
  );
};

export const useRopeContext = () => {
  const context = useContext(RopeContext);
  if (!context) {
    throw new Error('useRopeContext must be used within a RopeProvider');
  }
  return context;
};
