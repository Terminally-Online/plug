import React, { useCallback, useEffect, useState, useRef } from 'react';
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
import Controls from './Controls';
import ActionSelectorNode from './nodes/ActionSelectorNode';

// Define node types
const nodeTypes = {
  knotNode: KnotNode,
  valueNode: ValueNode,
  actionSelectorNode: ActionSelectorNode,
};

const RopeEditor = () => {
  const { knots, nodePositions, setNodePosition, updateKnotValue, addKnot, getPersistentKnotIds } = useRopeContext();
  const reactFlowInstance = useReactFlow();
  
  // Track the connection start parameters
  const [connectStart, setConnectStart] = useState<OnConnectStartParams | null>(null);
  
  // Track value nodes separately with position persistence
  const [valueNodes, setValueNodes] = useState<Node<ValueNodeData>[]>([]);
  // Track value node positions separately
  const [valueNodePositions, setValueNodePositions] = useState<Record<string, { x: number, y: number }>>({});
  // Track value edges separately
  const [valueEdges, setValueEdges] = useState<Edge[]>([]);
  // Track action selector nodes
  const [actionSelectorNodes, setActionSelectorNodes] = useState<Node[]>([]);
  
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

  // Track action selector nodes that have been transformed to knots
  const [transformedActionSelectors, setTransformedActionSelectors] = useState<string[]>([]);
  
  // Track nodes that should persist as knots
  // This is used to ensure transformed selectors stay as knots
  const persistentKnotIds = useRef<Set<string>>(new Set());
  
  // Initialize nodes and edges state
  const [nodes, setNodes, onNodesChange] = useNodesState([...initialKnotNodes, ...valueNodes, ...actionSelectorNodes]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([...initialSequenceEdges, ...valueEdges]);
  
  // Function to check for and remove any invalid edges
  const cleanupInvalidEdges = useCallback(() => {
    console.log('Cleaning up invalid edges...');

    // Get current nodes and edges
    const currentNodes = reactFlowInstance.getNodes();
    const currentEdges = reactFlowInstance.getEdges();

    // Get all valid node IDs
    const nodeIds = new Set(currentNodes.map(node => node.id));
    
    // Filter edges to keep only valid ones (both source and target exist)
    const validEdges = currentEdges.filter(edge => {
      const sourceExists = nodeIds.has(edge.source);
      const targetExists = nodeIds.has(edge.target);
      
      // Keep the edge only if both source and target exist
      const isValid = sourceExists && targetExists;
      
      if (!isValid) {
        console.log('Found invalid edge to remove:', edge);
        return false;
      }
      
      return true;
    });
    
    // If we found invalid edges, update the edges
    if (validEdges.length < currentEdges.length) {
      console.log(`Removing ${currentEdges.length - validEdges.length} invalid edges`);
      reactFlowInstance.setEdges(validEdges);
      
      // Also update our tracked value edges
      setValueEdges(prev => 
        prev.filter(edge => 
          nodeIds.has(edge.source) && nodeIds.has(edge.target)
        )
      );
    }
  }, [reactFlowInstance, setValueEdges]);

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
        console.log('Node removal detected:', nodeId);
        
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
        
        // Clean up any edges connected to this node, regardless of node type
        setEdges(eds => eds.filter(edge => 
          edge.source !== nodeId && edge.target !== nodeId
        ));
      }
    });
    
    // Apply the changes
    onNodesChange(changes);
    
    // Immediately clean up any invalid edges
    cleanupInvalidEdges();
  }, [valueNodes, onNodesChange, updateKnotValue, setEdges, cleanupInvalidEdges]);

  // Update nodes when knots or action selectors change
  useEffect(() => {
    console.log("Effect triggered - updating nodes from state:", {
      knots: knots.length,
      valueNodes: valueNodes.length,
      actionSelectors: actionSelectorNodes.length
    });
    
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
    
    // Get the list of persistent knot IDs
    const persistentKnots = getPersistentKnotIds();
    console.log('Persistent knot IDs to preserve:', persistentKnots);
    
    // Filter out action selectors that have been transformed to knots or are registered as persistent
    const filteredActionSelectors = actionSelectorNodes.filter(node => {
      // Keep the node only if it's not transformed and not registered as persistent
      const shouldKeepAsSelector = 
        !transformedActionSelectors.includes(node.id) && 
        !persistentKnots.includes(node.id);
      
      if (!shouldKeepAsSelector) {
        console.log(`Node ${node.id} is registered as a knot, filtering out from action selectors`);
      }
      
      return shouldKeepAsSelector;
    });
    
    // Combine with value nodes and action selector nodes
    console.log("Updating ReactFlow nodes:", {
      knots: updatedKnotNodes.length,
      valueNodes: updatedValueNodes.length,
      actionSelectors: filteredActionSelectors.length,
      actionSelectorIds: filteredActionSelectors.map(n => n.id)
    });
    
    const allNodes = [...updatedKnotNodes, ...updatedValueNodes, ...filteredActionSelectors];
    console.log("Total nodes to render:", allNodes.length);
    
    // Set the nodes with all three types
    setNodes(allNodes);
    
    // Clean up any invalid edges before we create new ones
    cleanupInvalidEdges();
    
    // We no longer automatically create sequence edges
    // Each connection should be created manually by the user
    const updatedSequenceEdges = [];
    
    // Create edges involving action selector nodes
    // Get existing edges from ReactFlow
    const currentEdges = reactFlowInstance.getEdges();
    
    // Filter out edges involving action selectors
    const actionSelectorEdges = currentEdges.filter(edge => {
      const isActionSelectorSource = actionSelectorNodes.some(node => node.id === edge.source);
      const isActionSelectorTarget = actionSelectorNodes.some(node => node.id === edge.target);
      return isActionSelectorSource || isActionSelectorTarget;
    });
    
    console.log("Action selector edges:", actionSelectorEdges);
    
    // Combine all edge types
    const allEdges = [...updatedSequenceEdges, ...valueEdges, ...actionSelectorEdges];
    console.log("Total edges to render:", allEdges.length);
    
    // Preserve existing edges and add our new ones
    setEdges(eds => {
      // Filter out any edges we're explicitly managing
      const otherEdges = eds.filter(edge => {
        // Check if this is a sequence edge (has format e{knotId}-{nextKnotId})
        const isSequenceEdge = edge.id && edge.id.match(/^e(knot-\d+)-(knot-\d+)$/);
        
        // Check if this is a value edge (from value node to knot)
        const isValueEdge = edge.id && edge.id.startsWith('e-value-') && edge.id.includes('-to-');
        
        // Keep edges that are not sequence edges or value edges
        return !isSequenceEdge && !isValueEdge;
      });
      
      // Create a new set of deduplicated edges by using a Map with edge IDs as keys
      const edgeMap = new Map();
      
      // Add other edges to the map
      otherEdges.forEach(edge => {
        edgeMap.set(edge.id, edge);
      });
      
      // Add our managed edges to the map (will overwrite any duplicates by ID)
      allEdges.forEach(edge => {
        edgeMap.set(edge.id, edge);
      });
      
      // Convert the map back to an array
      return Array.from(edgeMap.values());
    });
  }, [knots, nodePositions, valueNodes, valueEdges, valueNodePositions, actionSelectorNodes, transformedActionSelectors, setNodes, setEdges, reactFlowInstance, getPersistentKnotIds, cleanupInvalidEdges]);

  // Handle connection start
  const onConnectStart = useCallback((_, params) => {
    setConnectStart(params);
  }, []);

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

  // Placeholder comment - we removed the duplicate cleanupInvalidEdges function
  // The first declaration is at the top of the file

  const onNodeDragStop = useCallback(
    (_: any, node: Node) => {
      // Get persistent knot IDs
      const persistentKnots = getPersistentKnotIds();
      
      // Update positions based on node type
      if (node.type === 'knotNode' || persistentKnots.includes(node.id)) {
        // Update knot positions in the context - include persistent knots even if they're a different type
        console.log('Saving knot position for:', node.id);
        setNodePosition(node.id, node.position);
        
        // We no longer need to refresh sequence edges, as they're no longer automatically created
        // Just make sure to clean up any invalid edges
        cleanupInvalidEdges();
      } else if (node.type === 'valueNode') {
        // Update value node positions in our local state
        setValueNodePositions(prev => ({
          ...prev,
          [node.id]: node.position
        }));
      } else {
        // For other node types, at least log the position update
        console.log('Node drag stop for:', node.id, node.type, node.position);
      }
      
      // Clean up any invalid edges after dragging
      cleanupInvalidEdges();
    },
    [setNodePosition, getPersistentKnotIds, cleanupInvalidEdges, knots, setEdges]
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

  // Function to add a new knot node for the onAddNode prop
  const handleAddDefaultNode = useCallback(() => {
    // Create a new knot at a random position
    const id = `knot-${Date.now()}`;
    const position = { 
      x: 100 + Math.random() * 200, 
      y: 100 + Math.random() * 100 
    };
    setNodePosition(id, position);
    
    // Use a default protocol action when adding a new knot
    const defaultAction = {
      protocol: 'default',
      action: 'default',
      sentence: 'New sentence with {0} and {1}'
    };
    addKnot(defaultAction, id);
  }, [setNodePosition, addKnot]);

  // Track action selector nodes that are currently being transformed
  // This prevents automatic creation of new action selectors during transformation
  const [transformingSelectors, setTransformingSelectors] = useState<string[]>([]);

  // Handle selecting an action from an action selector node
  const handleSelectAction = useCallback((actionSelectorId: string, action: ProtocolAction) => {
    console.log('handleSelectAction called with ID:', actionSelectorId);
    console.log('Action to create:', action);
    
    // Mark this selector as being transformed to prevent auto-creation of new selectors
    setTransformingSelectors(prev => [...prev, actionSelectorId]);
    
    // First, search in the latest ReactFlow nodes since they're the actual rendered elements
    const currentNodes = reactFlowInstance.getNodes();
    console.log('Current ReactFlow nodes:', currentNodes.map(n => ({ id: n.id, type: n.type })));
    
    // Find the action selector node from the current ReactFlow nodes
    let actionSelectorNode = currentNodes.find(node => node.id === actionSelectorId);
    
    if (!actionSelectorNode) {
      console.error('Action selector node not found in ReactFlow nodes for ID:', actionSelectorId);
      
      // As a fallback, try using our tracked state
      const fallbackNode = actionSelectorNodes.find(node => node.id === actionSelectorId);
      if (!fallbackNode) {
        console.error('Action selector node also not found in tracked state');
        return;
      }
      console.log('Using fallback node from tracked state:', fallbackNode);
      
      // Continue with the fallback node
      actionSelectorNode = fallbackNode as any;
    }
    
    console.log('Found action selector node:', actionSelectorNode);
    
    const position = actionSelectorNode.position;
    const isFromConnection = actionSelectorNode.data?.isFromConnection || false;
    
    // Create a new knot ID with the current timestamp to ensure uniqueness
    const newKnotId = `knot-${Date.now()}`;
    
    // Set the position for the new knot in our tracker
    setNodePosition(newKnotId, position);
    
    // Add the knot to the rope context
    addKnot(action, newKnotId);
    
    // Create the knot node for ReactFlow
    const newKnotNode = {
      id: newKnotId,
      type: 'knotNode',
      position,
      data: { 
        knot: {
          id: newKnotId,
          sentence: action.sentence,
          protocol: action.protocol,
          action: action.action,
        }
      },
    };
    
    // Update connections if needed
    if (isFromConnection) {
      setEdges(eds => {
        // Find edges targeting the action selector
        const edgesToUpdate = eds.filter(edge => edge.target === actionSelectorId);
        
        // Create new edges pointing to the new knot
        const newEdges = edgesToUpdate.map(edge => ({
          ...edge,
          id: `${edge.id}-updated`,
          target: newKnotId,
          targetHandle: edge.targetHandle === 'continuity' ? 'continuity' : edge.targetHandle
        }));
        
        // Return updated edges
        return [...eds.filter(edge => edge.target !== actionSelectorId), ...newEdges];
      });
    }
    
    // Directly update ReactFlow nodes for immediate UI update
    reactFlowInstance.setNodes(nodes => {
      console.log('Directly updating ReactFlow nodes, removing:', actionSelectorId, 'adding:', newKnotId);
      return [
        ...nodes.filter(node => node.id !== actionSelectorId),
        newKnotNode
      ];
    });
    
    // Mark this action selector as transformed to prevent it from reappearing
    setTransformedActionSelectors(prev => [...prev, actionSelectorId]);
    
    // Update our tracked action selector nodes
    setActionSelectorNodes(nodes => 
      nodes.filter(node => node.id !== actionSelectorId)
    );
    
  }, [reactFlowInstance, actionSelectorNodes, setNodePosition, addKnot, setEdges, setActionSelectorNodes, setTransformedActionSelectors]);
  
  // Handle double-click on the pane to add a new action selector node
  const onPaneDoubleClick = useCallback((event: React.MouseEvent) => {
    // Get the click position in flow coordinates
    const position = reactFlowInstance.screenToFlowPosition({
      x: event.clientX,
      y: event.clientY,
    });
    
    // Create a unique ID for the action selector node
    const id = `action-selector-${Date.now()}`;
    
    // Create the new action selector node
    const newNode: Node = {
      id,
      type: 'actionSelectorNode',
      position,
      data: { 
        id,
        onSelectAction: handleSelectAction,
        isFromConnection: false
      },
    };
    
    // First update the ReactFlow nodes directly
    setNodes(currentNodes => {
      const updatedNodes = [...currentNodes, newNode];
      console.log('Updating ReactFlow nodes with new action selector:', updatedNodes);
      return updatedNodes;
    });
    
    // Then update our action selector tracking
    setActionSelectorNodes(prev => {
      const updatedSelectorNodes = [...prev, newNode];
      console.log('Adding to actionSelectorNodes:', updatedSelectorNodes);
      return updatedSelectorNodes;
    });
  }, [reactFlowInstance, handleSelectAction, setNodes, setActionSelectorNodes]);
  
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
        // Check if the source node is currently being transformed
        // If it is, we don't want to create a new action selector
        const sourceNodeId = connectStart.nodeId as string;
        const isSourceNodeBeingTransformed = transformingSelectors.includes(sourceNodeId);
        
        // If the source node is being transformed, bail out early
        if (isSourceNodeBeingTransformed) {
          console.log('Source node is being transformed, skipping auto-creation of action selector');
          return;
        }
        
        // Create a unique ID for the action selector node
        const id = `action-selector-${Date.now()}`;
        
        // Create the new action selector node with connection info
        const newNode: Node = {
          id,
          type: 'actionSelectorNode',
          position,
          data: { 
            id,
            onSelectAction: handleSelectAction,
            isFromConnection: true
          },
        };
        
        // First update the ReactFlow nodes directly
        setNodes(currentNodes => {
          const updatedNodes = [...currentNodes, newNode];
          console.log('Drag-create: Updating ReactFlow nodes with new action selector:', updatedNodes);
          return updatedNodes;
        });
        
        // Then update our action selector tracking
        setActionSelectorNodes(prev => {
          const updatedSelectorNodes = [...prev, newNode];
          console.log('Drag-create: Adding to actionSelectorNodes:', updatedSelectorNodes);
          return updatedSelectorNodes;
        });
        
        // Create an edge from the source node to this action selector
        if (connectStart.nodeId) {
          const newEdge: Edge = {
            id: `e${connectStart.nodeId}-${id}`,
            source: connectStart.nodeId as string,
            sourceHandle: 'output',
            target: id,
            targetHandle: 'continuity',
            type: 'default',
            animated: true,
            style: { strokeWidth: 2, stroke: '#3b82f6' },
          };
          
          // Add the edge with both ReactFlow and our state
          console.log('Creating connection edge from source:', connectStart.nodeId, 'to action selector:', id);
          
          // First directly update ReactFlow for immediate visual feedback
          reactFlowInstance.addEdges([newEdge]);
          
          // Then update our tracked state
          setEdges(eds => {
            // Remove any duplicate edges between the same nodes
            const filteredEdges = eds.filter(e => 
              !(e.source === connectStart.nodeId && e.target === id)
            );
            return [...filteredEdges, newEdge];
          });
          
          // No need to update the view - we want to maintain the user's current view
        }
      }
      
      // Clear the connection start
      setConnectStart(null);
    },
    [connectStart, reactFlowInstance, knots, updateKnotValue, setValueNodePositions, nodes, valueNodes, setValueNodes, setValueEdges, edges, setEdges, transformingSelectors]
  );

  return (
    <div className="flex-grow h-full relative flex">
      <div className="flex-grow h-full">
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={handleNodesChange}
          onEdgesChange={handleEdgesChange}
          onConnect={onConnect}
          onConnectStart={onConnectStart}
          onConnectEnd={onConnectEnd}
          onNodeDragStop={onNodeDragStop}
          onPaneDoubleClick={onPaneDoubleClick}
          nodeTypes={nodeTypes}
          // Removed fitView to prevent automatic zooming when nodes are added
          deleteKeyCode="Delete"
          nodesFocusable={true}
          selectNodesOnDrag={false}
          proOptions={{ hideAttribution: true }}
          zoomOnScroll={false}
          zoomOnPinch={true}
          panOnScroll={true}
          zoomOnDoubleClick={false}
          zoomActivationKeyCode="Meta"
        >
          <Background variant="dots" gap={12} size={1} />
          
          {/* Floating control panel */}
          <div className="absolute top-4 left-4 z-10 bg-white bg-opacity-80 backdrop-blur-sm p-3 rounded-lg shadow-lg border border-gray-200">
            <h3 className="text-md font-semibold mb-2">Rope Editor</h3>
            <div className="flex space-x-2 mb-3">
              <button 
                className="bg-blue-500 hover:bg-blue-600 text-white text-sm px-3 py-1 rounded"
                onClick={() => {
                  // Create an action selector in the center of the current view
                  const { x, y } = reactFlowInstance.getViewport();
                  const centerPosition = reactFlowInstance.project({
                    x: window.innerWidth / 2 - x,
                    y: window.innerHeight / 2 - y,
                  });
                  
                  // Create a unique ID for the action selector node
                  const id = `action-selector-${Date.now()}`;
                  
                  // Create the new action selector node
                  const newNode: Node = {
                    id,
                    type: 'actionSelectorNode',
                    position: centerPosition,
                    data: { 
                      id,
                      onSelectAction: handleSelectAction,
                      isFromConnection: false
                    },
                  };
                  
                  // First update the ReactFlow nodes directly
                  setNodes(currentNodes => {
                    const updatedNodes = [...currentNodes, newNode];
                    console.log('Button-click: Updating ReactFlow nodes with new action selector:', updatedNodes);
                    return updatedNodes;
                  });
                  
                  // Then update our action selector tracking
                  setActionSelectorNodes(prev => {
                    const updatedSelectorNodes = [...prev, newNode];
                    console.log('Button-click: Adding to actionSelectorNodes:', updatedSelectorNodes);
                    return updatedSelectorNodes;
                  });
                }}
              >
                Add Action
              </button>
            </div>
          </div>
        
        </ReactFlow>
      </div>
    </div>
  );
};

// Wrapper component to provide ReactFlow context
const RopeEditorWithProvider = () => (
  <ReactFlowProvider>
    <RopeEditor />
  </ReactFlowProvider>
);

export default RopeEditorWithProvider;