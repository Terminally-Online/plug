import React, { useCallback, useEffect, useState } from 'react';
import ReactFlow, {
  MiniMap,
  Controls as FlowControls,
  Background,
  useNodesState,
  useEdgesState,
  addEdge,
  Connection,
  Edge,
  Node,
  OnConnectStartParams,
  useReactFlow,
  NodeMouseHandler,
  ReactFlowProvider
} from 'reactflow';
import 'reactflow/dist/style.css';

import { useRopeContext } from '../context/RopeContext';
import KnotNode from './nodes/KnotNode';
import ValueNode, { ValueNodeData } from './nodes/ValueNode';

// Define node types
const nodeTypes = {
  knotNode: KnotNode,
  valueNode: ValueNode,
};

const RopeEditor = () => {
  const { knots, nodePositions, setNodePosition, updateKnotValue, addKnot } = useRopeContext();
  const reactFlowInstance = useReactFlow();
  
  // Track the connection start parameters
  const [connectStart, setConnectStart] = useState<OnConnectStartParams | null>(null);
  
  // Track value nodes separately with position persistence
  const [valueNodes, setValueNodes] = useState<Node<ValueNodeData>[]>([]);
  // Track value node positions separately
  const [valueNodePositions, setValueNodePositions] = useState<Record<string, { x: number, y: number }>>({});
  // Track value edges separately
  const [valueEdges, setValueEdges] = useState<Edge[]>([]);
  
  // Initialize nodes from knots
  const initialKnotNodes: Node[] = knots.map((knot) => ({
    id: knot.id,
    type: 'knotNode',
    position: nodePositions[knot.id] || { x: 250, y: 100 },
    data: { knot },
  }));

  // Create edges connecting knots in sequence
  const initialSequenceEdges: Edge[] = knots.length > 1 ? knots.slice(0, -1).map((knot, index) => ({
    id: `e${knot.id}-${knots[index + 1].id}`,
    source: knot.id,
    sourceHandle: 'output',
    target: knots[index + 1].id,
    targetHandle: 'continuity', // Connect to the continuity handle
    type: 'smoothstep',
    animated: true,
  })) : [];

  // Initialize nodes and edges state
  const [nodes, setNodes, onNodesChange] = useNodesState([...initialKnotNodes, ...valueNodes]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([...initialSequenceEdges, ...valueEdges]);
  
  // Custom edge change handler to handle connection removals
  const handleEdgesChange = useCallback((changes) => {
    // Process the changes to detect edge removals
    changes.forEach(change => {
      if (change.type === 'remove') {
        const edgeId = change.id;
        
        // Check if it's a value edge by parsing the ID format
        // Format: e-{valueNodeId}-to-{knotId}-input-{inputIndex}
        if (edgeId.startsWith('e-') && edgeId.includes('-to-') && edgeId.includes('-input-')) {
          try {
            // Extract the parts from the edge ID
            const parts = edgeId.split('-to-');
            const sourceValueId = parts[0].substring(2); // Remove 'e-' prefix
            
            const targetParts = parts[1].split('-input-');
            const targetKnotId = targetParts[0];
            const inputIndex = parseInt(targetParts[1], 10);
            
            // Find the value node to update its connections
            setValueNodes(current => 
              current.map(vn => {
                if (vn.id === sourceValueId) {
                  // Remove this connection from the connections array
                  const updatedConnections = (vn.data.connections || [])
                    .filter(conn => 
                      !(conn.targetId === targetKnotId && conn.inputIndex === inputIndex)
                    );
                  
                  return {
                    ...vn,
                    data: {
                      ...vn.data,
                      connections: updatedConnections
                    }
                  };
                }
                return vn;
              })
            );
            
            // Clear the value in the target input
            updateKnotValue(targetKnotId, inputIndex, '');
          } catch (err) {
            console.error('Error parsing edge ID:', err);
          }
        }
      }
    });
    
    // Apply the changes
    onEdgesChange(changes);
  }, [onEdgesChange, setValueNodes, updateKnotValue]);
  
  // Custom node change handler to detect deletions
  const handleNodesChange = useCallback((changes) => {
    // Process the changes to detect node removals
    changes.forEach(change => {
      if (change.type === 'remove') {
        const nodeId = change.id;
        
        // Check if it's a value node being removed
        const isValueNode = valueNodes.some(vn => vn.id === nodeId);
        if (isValueNode) {
          // Find the value node to access its connections
          const valueNode = valueNodes.find(vn => vn.id === nodeId);
          
          // Clear any values that were set by this node
          if (valueNode && valueNode.data.connections) {
            valueNode.data.connections.forEach(conn => {
              // Clear the value for this connection
              updateKnotValue(conn.targetId, conn.inputIndex, '');
            });
          }
          
          // Remove from our value nodes tracking
          setValueNodes(current => current.filter(vn => vn.id !== nodeId));
          
          // Remove from positions tracking
          setValueNodePositions(current => {
            const updated = { ...current };
            delete updated[nodeId];
            return updated;
          });
          
          // Remove any associated edges
          setValueEdges(current => current.filter(edge => 
            edge.source !== nodeId && edge.target !== nodeId
          ));
        }
      }
    });
    
    // Apply the changes
    onNodesChange(changes);
  }, [valueNodes, onNodesChange, updateKnotValue]);

  // Update nodes when knots change
  useEffect(() => {
    const updatedKnotNodes = knots.map((knot) => ({
      id: knot.id,
      type: 'knotNode',
      position: nodePositions[knot.id] || { x: 250, y: 100 },
      data: { knot },
    }));
    
    // Create updated value nodes with preserved positions
    const updatedValueNodes = valueNodes.map(vNode => ({
      ...vNode,
      // Preserve position from our position tracking, fall back to node's current position
      position: valueNodePositions[vNode.id] || vNode.position
    }));
    
    // Combine with value nodes
    setNodes([...updatedKnotNodes, ...updatedValueNodes]);
    
    // Update sequence edges
    const updatedSequenceEdges = knots.length > 1 ? knots.slice(0, -1).map((knot, index) => ({
      id: `e${knot.id}-${knots[index + 1].id}`,
      source: knot.id,
      sourceHandle: 'output',
      target: knots[index + 1].id,
      targetHandle: 'continuity', // Connect to the continuity handle
      type: 'default', // Use default edge type for cleaner lines
      animated: true,
      style: { strokeWidth: 2, stroke: '#3b82f6' }, // Slightly thicker blue lines
    })) : [];
    
    // Combine with value edges
    setEdges([...updatedSequenceEdges, ...valueEdges]);
  }, [knots, nodePositions, valueNodes, valueEdges, valueNodePositions, setNodes, setEdges]);

  // Handle connection start
  const onConnectStart = useCallback((_, params) => {
    setConnectStart(params);
  }, []);

  // Handle connection end
  const onConnectEnd = useCallback(
    (event) => {
      if (!connectStart || !connectStart.handleId) {
        return;
      }
      
      // Get the exact mouse position and convert to flow coordinates
      const { clientX, clientY } = event;
      const position = reactFlowInstance.screenToFlowPosition({
        x: clientX,
        y: clientY,
      });
      
      // Get the source knot ID
      const sourceKnotId = connectStart.nodeId as string;
      const sourceKnot = knots.find(k => k.id === sourceKnotId);
      if (!sourceKnot) return;
      
      // Check if we're connecting TO a value node's output
      // Get the element at the current mouse position
      const targetElement = document.elementFromPoint(clientX, clientY);
      const targetNodeId = targetElement?.closest('.react-flow__node')?.getAttribute('data-id');
      const isTargetingValueNode = targetNodeId && nodes.some(
        n => n.id === targetNodeId && n.type === 'valueNode'
      );
      
      // Create different node types based on the handle type
      const isFromInputHandle = connectStart.handleId.startsWith('input-');
      const isFromOutputHandle = connectStart.handleId === 'output';
      
      if (isFromInputHandle) {
        // Extract the input index from the handle ID
        const inputIndex = parseInt(connectStart.handleId.replace('input-', ''), 10);
        
        // Check if we're connecting to an existing value node
        if (isTargetingValueNode) {
          // Find the value node we're connecting to
          const targetValueNode = valueNodes.find(vn => vn.id === targetNodeId);
          if (targetValueNode) {
            // Create a new connection in the target value node
            const newConnection = { targetId: sourceKnotId, inputIndex };
            
            // Check if there's already a connection to this input from any value node
            const existingEdge = edges.find(edge => 
              edge.target === sourceKnotId && 
              edge.targetHandle === `input-${inputIndex}`
            );
            
            // If there's an existing connection from a different value node, we need to:
            // 1. Remove the edge
            // 2. Update the source value node to remove this connection
            if (existingEdge && existingEdge.source !== targetNodeId) {
              // Remove the existing edge from our valueEdges state
              setValueEdges(vedges => vedges.filter(edge => edge.id !== existingEdge.id));
              
              // IMPORTANT: Also remove from the main edges state for proper rendering
              setEdges(eds => eds.filter(edge => edge.id !== existingEdge.id));
              
              // Find the source value node of the existing connection
              const oldSourceValueNodeId = existingEdge.source;
              
              // Update the source value node to remove this connection
              setValueNodes(vnodes => 
                vnodes.map(vn => {
                  if (vn.id === oldSourceValueNodeId) {
                    const updatedConnections = (vn.data.connections || [])
                      .filter(conn => 
                        !(conn.targetId === sourceKnotId && conn.inputIndex === inputIndex)
                      );
                    
                    return {
                      ...vn,
                      data: {
                        ...vn.data,
                        connections: updatedConnections
                      }
                    };
                  }
                  return vn;
                })
              );
            }
            
            // Check if this specific connection already exists
            const connectionExists = targetValueNode.data.connections?.some(
              c => c.targetId === sourceKnotId && c.inputIndex === inputIndex
            );
            
            // Only create a new connection if it doesn't exist already
            if (!connectionExists) {
              // Update the value node to add this connection
              setValueNodes(vnodes => 
                vnodes.map(vn => {
                  if (vn.id === targetNodeId) {
                    const connections = vn.data.connections || [];
                    return {
                      ...vn,
                      data: {
                        ...vn.data,
                        connections: [...connections, newConnection]
                      }
                    };
                  }
                  return vn;
                })
              );
              
              // Create a new edge for this connection
              const newValueEdge: Edge = {
                id: `e-${targetNodeId}-to-${sourceKnotId}-input-${inputIndex}`,
                source: targetNodeId,
                sourceHandle: 'value-output',
                target: sourceKnotId,
                targetHandle: `input-${inputIndex}`,
                type: 'default',
                animated: true,
                style: { strokeWidth: 2, stroke: '#22c55e' }, // Green lines for value connections
              };
              
              // Add the edge to our state
              setValueEdges((vedges) => [...vedges, newValueEdge]);
              
              // IMPORTANT: We also need to update the main edges state to ensure the edge is rendered
              // This is necessary because ReactFlow uses the edges state for rendering
              setEdges((eds) => [...eds, newValueEdge]);
              
              // Update the knot with the current value from the value node
              if (targetValueNode.data.value) {
                updateKnotValue(sourceKnotId, inputIndex, targetValueNode.data.value);
              }
            }
            
            // Stop here since we've handled the connection to an existing value node
            return;
          }
        }
        
        // If we're not connecting to an existing value node, create a new one
        const newValueNodeId = `value-${Date.now()}`;
        
        // Label for the value node based on the input
        const label = `Value for Input ${inputIndex}`;
        
        // Create the new value node with connection tracking
        const newValueNode: Node<ValueNodeData> = {
          id: newValueNodeId,
          type: 'valueNode',
          position,
          data: {
            label,
            value: '',
            // Keep these for backward compatibility
            inputIndex,
            targetKnotId: sourceKnotId,
            // New connection tracking system
            connections: [
              { targetId: sourceKnotId, inputIndex }
            ],
            onChange: (newValue: string) => {
              // Update all connected knot values when the input changes
              // We can use the connections array to update all connected knots
              if (newValueNode.data.connections) {
                newValueNode.data.connections.forEach(conn => {
                  updateKnotValue(conn.targetId, conn.inputIndex, newValue);
                });
              } else {
                // Fallback to the original pattern
                updateKnotValue(sourceKnotId, inputIndex, newValue);
              }
            }
          },
        };
        
        // Check if there's already a connection to this input
        const existingEdge = edges.find(edge => 
          edge.target === sourceKnotId && 
          edge.targetHandle === `input-${inputIndex}`
        );
        
        // If there's an existing connection, we need to:
        // 1. Remove the edge
        // 2. Update the source value node to remove this connection
        if (existingEdge) {
          // Remove the existing edge from our valueEdges state
          setValueEdges(vedges => vedges.filter(edge => edge.id !== existingEdge.id));
          
          // IMPORTANT: Also remove from the main edges state for proper rendering
          setEdges(eds => eds.filter(edge => edge.id !== existingEdge.id));
          
          // Find the source value node of the existing connection
          const sourceValueNodeId = existingEdge.source;
          
          // Update the source value node to remove this connection
          setValueNodes(vnodes => 
            vnodes.map(vn => {
              if (vn.id === sourceValueNodeId) {
                const updatedConnections = (vn.data.connections || [])
                  .filter(conn => 
                    !(conn.targetId === sourceKnotId && conn.inputIndex === inputIndex)
                  );
                
                return {
                  ...vn,
                  data: {
                    ...vn.data,
                    connections: updatedConnections
                  }
                };
              }
              return vn;
            })
          );
        }
        
        // Create a new edge from the value node to the input handle with a custom ID format
        const newValueEdge: Edge = {
          id: `e-${newValueNodeId}-to-${sourceKnotId}-input-${inputIndex}`,
          source: newValueNodeId,
          sourceHandle: 'value-output',
          target: sourceKnotId,
          targetHandle: `input-${inputIndex}`,
          type: 'default',
          animated: true,
          style: { strokeWidth: 2, stroke: '#22c55e' }, // Green lines for value connections
        };
        
        // Add the new value node and edge
        setValueNodes((vnodes) => [...vnodes, newValueNode]);
        setValueEdges((vedges) => [...vedges, newValueEdge]);
        
        // IMPORTANT: Also add to the main edges state for proper rendering
        setEdges((eds) => [...eds, newValueEdge]);
        
        // Store the position in our position tracker
        setValueNodePositions(prev => ({
          ...prev,
          [newValueNodeId]: position
        }));
      } 
      else if (isFromOutputHandle) {
        // Create a new knot node when dragging from the output handle
        
        // Create a new template with a timestamp ID to ensure uniqueness
        const timestamp = Date.now();
        const newKnotId = `knot-${timestamp}`;
        const defaultTemplate = 'New action with {0} and {1}';
        
        // First, set the position in nodePositions for the new knot 
        // This needs to happen before adding the knot
        setNodePosition(newKnotId, position);
        
        // Add the new knot to the system
        // We use setTimeout with 0ms delay to ensure the position is set in the state
        // before the knot is added, so that the useEffect picks up the position
        setTimeout(() => {
          addKnot(defaultTemplate, newKnotId);
        }, 0);
      }
      
      // Clear the connection start
      setConnectStart(null);
    },
    [connectStart, reactFlowInstance, knots, updateKnotValue, setValueNodePositions, setNodePosition, addKnot, nodes, valueNodes, setValueNodes, setValueEdges]
  );

  // Standard connection handler
  const onConnect = useCallback(
    (params: Edge | Connection) => {
      // Check if trying to connect to an input handle
      if (params.targetHandle && params.targetHandle.startsWith('input-')) {
        const inputIndex = parseInt(params.targetHandle.replace('input-', ''), 10);
        const targetKnotId = params.target as string;
        
        // Check if the source is a value node
        const sourceNode = [...nodes].find(n => n.id === params.source);
        if (sourceNode && sourceNode.type === 'valueNode') {
          // First, check if there's already an edge connected to this input
          const existingEdge = edges.find(edge => 
            edge.target === targetKnotId && 
            edge.targetHandle === params.targetHandle
          );
          
          // If there's an existing edge, remove it first
          if (existingEdge) {
            setEdges(edges => edges.filter(e => e.id !== existingEdge.id));
          }
          
          // Create a custom ID for the edge to make it identifiable
          const edgeId = `e-${params.source}-to-${targetKnotId}-${params.targetHandle}`;
          
          // Create the edge with customized properties
          const newEdge: Edge = {
            ...params,
            id: edgeId,
            type: 'default',
            animated: true,
            style: { strokeWidth: 2, stroke: '#22c55e' }, // Green lines for value connections
          };
          
          // Add the new edge
          setEdges((eds) => [...eds, newEdge]);
          
          // Update the knot with the value from the value node
          const valueNode = sourceNode as Node<ValueNodeData>;
          if (valueNode.data.value) {
            updateKnotValue(targetKnotId, inputIndex, valueNode.data.value);
          }
          
          // Update the value node to track which inputs it's connected to
          // This is useful for propagating value changes
          const valueNodeData = valueNode.data;
          
          // Update the value node's connections list if it doesn't already include this connection
          setValueNodes(vnodes => 
            vnodes.map(vn => {
              if (vn.id === valueNode.id) {
                // Create a new connections array if it doesn't exist
                const connections = vn.data.connections || [];
                
                // Add the new connection if it doesn't exist
                const newConnection = { targetId: targetKnotId, inputIndex };
                const connectionExists = connections.some(
                  c => c.targetId === targetKnotId && c.inputIndex === inputIndex
                );
                
                if (!connectionExists) {
                  return {
                    ...vn,
                    data: {
                      ...vn.data,
                      connections: [...connections, newConnection]
                    }
                  };
                }
              }
              return vn;
            })
          );
        }
      } else {
        // Standard edge (for continuity connections)
        setEdges((eds) => addEdge({
          ...params,
          type: 'default',
          animated: true,
          style: { strokeWidth: 2, stroke: '#3b82f6' }, // Blue lines for continuity
        }, eds));
      }
    },
    [nodes, edges, setEdges, setValueNodes, updateKnotValue]
  );

  const onNodeDragStop = useCallback(
    (_: any, node: Node) => {
      // Update positions based on node type
      if (node.type === 'knotNode') {
        // Update knot positions in the context
        setNodePosition(node.id, node.position);
      } else if (node.type === 'valueNode') {
        // Update value node positions in our local state
        setValueNodePositions(prev => ({
          ...prev,
          [node.id]: node.position
        }));
      }
    },
    [setNodePosition]
  );

  // Get the resolved sentence for the entire rope
  const getResolvedRope = useCallback(() => {
    // Get all knots in order and extract their resolved sentences
    const resolvedSentences = knots
      .filter(knot => knot.resolvedSentence)
      .map(knot => knot.resolvedSentence);
    
    return resolvedSentences.join(' → ');
  }, [knots]);
  
  // Track whether we should show the resolved rope
  const [showResolvedRope, setShowResolvedRope] = useState(false);

  return (
    <div className="flex-grow h-full relative">
      <ReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={handleNodesChange}
        onEdgesChange={handleEdgesChange}
        onConnect={onConnect}
        onConnectStart={onConnectStart}
        onConnectEnd={onConnectEnd}
        onNodeDragStop={onNodeDragStop}
        nodeTypes={nodeTypes}
        fitView
        deleteKeyCode="Delete"
        nodesFocusable={true}
        selectNodesOnDrag={false}
      >
        <Background variant="dots" gap={12} size={1} />
        <MiniMap nodeStrokeWidth={3} zoomable pannable />
        <FlowControls />
        
        {/* Floating control panel */}
        <div className="absolute top-4 left-4 z-10 bg-white bg-opacity-80 backdrop-blur-sm p-3 rounded-lg shadow-lg border border-gray-200">
          <h3 className="text-md font-semibold mb-2">Rope Editor</h3>
          <div className="flex space-x-2 mb-3">
            <button 
              className="bg-blue-500 hover:bg-blue-600 text-white text-sm px-3 py-1 rounded"
              onClick={() => {
                // Create a new knot at a random position
                const id = `knot-${Date.now()}`;
                const position = { 
                  x: 100 + Math.random() * 200, 
                  y: 100 + Math.random() * 100 
                };
                setNodePosition(id, position);
                setTimeout(() => {
                  addKnot('New sentence with {0} and {1}', id);
                }, 0);
              }}
            >
              Add Knot
            </button>
            <button 
              className="bg-gray-500 hover:bg-gray-600 text-white text-sm px-3 py-1 rounded"
              onClick={() => {
                // Fit the view to see all nodes
                reactFlowInstance.fitView({ padding: 0.2 });
              }}
            >
              Fit View
            </button>
          </div>
          
          <div className="flex items-center mb-2">
            <label className="inline-flex items-center cursor-pointer">
              <input 
                type="checkbox"
                className="sr-only peer"
                checked={showResolvedRope}
                onChange={() => setShowResolvedRope(!showResolvedRope)}
              />
              <div className="relative w-9 h-5 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-blue-600"></div>
              <span className="ms-2 text-xs font-medium text-gray-700">Show Resolved Rope</span>
            </label>
          </div>
          
          <div className="text-xs text-gray-600 space-y-1">
            <p><span className="inline-block w-2 h-2 bg-green-500 rounded-full mr-1"></span> Drag from left handles for value nodes</p>
            <p><span className="inline-block w-2 h-2 bg-blue-500 rounded-full mr-1"></span> Drag from top-right handle for new knots</p>
            <p>Press <kbd className="bg-gray-200 px-1 rounded">Delete</kbd> to remove nodes</p>
          </div>
        </div>
        
        {/* Resolved rope panel */}
        {showResolvedRope && (
          <div className="absolute bottom-4 left-4 right-4 z-10 bg-white bg-opacity-95 backdrop-blur-sm p-3 rounded-lg shadow-lg border border-blue-100 max-h-36 overflow-auto">
            <div className="flex items-center mb-2">
              <svg className="w-4 h-4 text-blue-500 mr-1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z" clipRule="evenodd" />
              </svg>
              <h4 className="text-sm font-semibold text-blue-900">Resolved Rope</h4>
            </div>
            <div className="text-sm font-mono p-2 bg-blue-50 rounded-md border border-blue-100 break-all">
              {getResolvedRope() || (
                <span className="text-gray-500 italic">Complete all required inputs to see the resolved rope</span>
              )}
            </div>
          </div>
        )}
      </ReactFlow>
    </div>
  );
};

// Wrapper component to provide ReactFlow context
const RopeEditorWithProvider = () => {
  return (
    <ReactFlowProvider>
      <RopeEditor />
    </ReactFlowProvider>
  );
};

export default RopeEditorWithProvider;