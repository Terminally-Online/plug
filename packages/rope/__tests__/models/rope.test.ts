import { Rope } from '../../src/models/rope';
import { Knot } from '../../src/models/knot';
import { PlugAPI } from '../../src/models/api';

// Mock the API class
jest.mock('../../src/models/api');

describe('Rope', () => {
  let rope: Rope;
  const mockApiKey = 'test-api-key';
  
  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();
    
    // Create a new Rope instance for each test
    rope = new Rope(mockApiKey);
  });
  
  describe('constructor', () => {
    it('should create a new instance with the provided API key', () => {
      expect(rope).toBeInstanceOf(Rope);
      expect(PlugAPI).toHaveBeenCalledWith(mockApiKey, undefined);
    });
    
    it('should accept configuration overrides', () => {
      const config = { baseURL: 'https://custom-api.example.com' };
      rope = new Rope(mockApiKey, config);
      
      expect(PlugAPI).toHaveBeenCalledWith(mockApiKey, config);
    });
  });
  
  describe('initialization', () => {
    it('should fetch available schemas during initialization', async () => {
      const mockSchemas = {
        'uniswap': {
          metadata: {
            icon: 'uniswap-icon',
            chains: [{ chainIds: [1], name: 'Ethereum' }],
            tags: ['dex', 'swap']
          },
          schema: {
            'swap': { type: 'action' },
            'addLiquidity': { type: 'action' }
          }
        }
      };
      
      const mockApi = {
        getSolver: jest.fn().mockResolvedValue(mockSchemas)
      };
      
      (PlugAPI as jest.Mock).mockImplementation(() => mockApi);
      
      rope = new Rope(mockApiKey);
      expect(rope.isReady()).toBe(false);
      
      await rope.initialize();
      
      expect(mockApi.getSolver).toHaveBeenCalledWith(expect.objectContaining({
        chainId: 1
      }));
      expect(rope.isReady()).toBe(true);
      expect(rope.getAvailableSchemas()).toEqual(mockSchemas);
    });
    
    it('should throw an error when accessing schemas before initialization', () => {
      expect(() => rope.getAvailableSchemas()).toThrow(/initialize/);
      expect(() => rope.getProtocolSchema('uniswap')).toThrow(/initialize/);
    });
    
    it('should return schemas for a specific protocol', async () => {
      const mockSchemas = {
        'uniswap': {
          metadata: {
            icon: 'uniswap-icon',
            chains: [{ chainIds: [1], name: 'Ethereum' }],
            tags: ['dex', 'swap']
          },
          schema: { 
            'swap': { type: 'action' },
            'addLiquidity': { type: 'action' }
          }
        },
        'aave': {
          metadata: {
            icon: 'aave-icon',
            chains: [{ chainIds: [1], name: 'Ethereum' }]
          },
          schema: { 'deposit': { type: 'action' } }
        }
      };
      
      const mockApi = {
        getSolver: jest.fn().mockResolvedValue(mockSchemas)
      };
      
      (PlugAPI as jest.Mock).mockImplementation(() => mockApi);
      
      rope = new Rope(mockApiKey);
      await rope.initialize();
      
      const uniswapSchema = rope.getProtocolSchema('uniswap');
      expect(uniswapSchema).toEqual(mockSchemas['uniswap']);
      
      const unknownSchema = rope.getProtocolSchema('unknown');
      expect(unknownSchema).toBeUndefined();
    });
  });
  
  describe('knot creation and management', () => {
    beforeEach(async () => {
      const mockSchemas = {
        'uniswap': {
          metadata: {
            icon: 'uniswap-icon',
            chains: [{ chainIds: [1], name: 'Ethereum' }]
          },
          schema: {
            'swap': { type: 'action' },
            'addLiquidity': { type: 'action' }
          }
        }
      };
      
      const mockApi = {
        getSolver: jest.fn().mockResolvedValue(mockSchemas)
      };
      
      (PlugAPI as jest.Mock).mockImplementation(() => mockApi);
      
      rope = new Rope(mockApiKey);
      await rope.initialize();
    });
    
    it('should create a knot for a valid protocol and action', () => {
      const knot = rope.createKnot('uniswap', 'swap', { amount: '1.0' });
      
      expect(knot).toBeInstanceOf(Knot);
      expect(knot.getData('protocol')).toBe('uniswap');
      expect(knot.getData('action')).toBe('swap');
      expect(knot.getData('amount')).toBe('1.0');
    });
    
    it('should throw an error for an invalid protocol', () => {
      expect(() => rope.createKnot('invalid', 'swap')).toThrow(/Protocol "invalid" not found/);
    });
    
    it('should throw an error for an invalid action', () => {
      expect(() => rope.createKnot('uniswap', 'invalid')).toThrow(/Action "invalid" not found/);
    });
    
    it('should add a knot to the rope', () => {
      const knot = rope.createKnot('uniswap', 'swap');
      rope.addKnot(knot);
      
      const knots = rope.getKnots();
      expect(knots).toHaveLength(1);
      expect(knots[0]).toBe(knot);
    });
    
    it('should return all knots in the rope', () => {
      const knot1 = rope.createKnot('uniswap', 'swap', { tokenIn: 'ETH' });
      const knot2 = rope.createKnot('uniswap', 'swap', { tokenIn: 'DAI' });
      
      rope.addKnot(knot1);
      rope.addKnot(knot2);
      
      const knots = rope.getKnots();
      expect(knots).toHaveLength(2);
      expect(knots[0]).toBe(knot1);
      expect(knots[1]).toBe(knot2);
    });
    
    it('should support method chaining for addKnot', () => {
      const knot1 = rope.createKnot('uniswap', 'swap', { tokenIn: 'ETH' });
      const knot2 = rope.createKnot('uniswap', 'swap', { tokenIn: 'DAI' });
      
      const result = rope
        .addKnot(knot1)
        .addKnot(knot2);
      
      expect(result).toBe(rope);
      const knots = rope.getKnots();
      expect(knots).toHaveLength(2);
      expect(knots[0]).toBe(knot1);
      expect(knots[1]).toBe(knot2);
    });
  });
  
  describe('intent building and execution', () => {
    beforeEach(async () => {
      const mockSchemas = {
        'uniswap': {
          metadata: {
            icon: 'uniswap-icon',
            chains: [{ chainIds: [1], name: 'Ethereum' }]
          },
          schema: {
            'swap': { type: 'action' },
            'addLiquidity': { type: 'action' }
          }
        }
      };
      
      const mockApi = {
        getSolver: jest.fn().mockResolvedValue(mockSchemas),
        postSolver: jest.fn().mockResolvedValue({ success: true, txHash: '0x123' })
      };
      
      (PlugAPI as jest.Mock).mockImplementation(() => mockApi);
      
      rope = new Rope(mockApiKey);
      await rope.initialize();
    });
    
    it('should build an intent from the rope knots', () => {
      const knot1 = rope.createKnot('uniswap', 'swap', { 
        tokenIn: 'ETH', 
        tokenOut: 'DAI', 
        amountIn: '1.0' 
      });
      
      const knot2 = rope.createKnot('uniswap', 'addLiquidity', { 
        tokenA: 'DAI', 
        tokenB: 'USDC', 
        amountA: '100', 
        amountB: '100'
      });
      
      rope.addKnot(knot1).addKnot(knot2);
      
      const intent = rope.buildIntent();
      
      expect(intent).toMatchObject({
        chainId: 1,
        from: '',
        frequency: 0,
        inputs: [
          {
            protocol: 'uniswap',
            action: 'swap',
            params: {
              protocol: 'uniswap',
              action: 'swap',
              tokenIn: 'ETH',
              tokenOut: 'DAI',
              amountIn: '1.0'
            }
          },
          {
            protocol: 'uniswap',
            action: 'addLiquidity',
            params: {
              protocol: 'uniswap',
              action: 'addLiquidity',
              tokenA: 'DAI',
              tokenB: 'USDC',
              amountA: '100',
              amountB: '100'
            }
          }
        ],
        locked: false,
        saved: false,
        status: 'active'
      });
    });
    
    it('should throw an error when building an intent with no knots', () => {
      expect(() => rope.buildIntent()).toThrow(/Cannot build intent: rope has no knots/);
    });
    
    it('should execute the rope as an intent', async () => {
      const knot = rope.createKnot('uniswap', 'swap', { 
        tokenIn: 'ETH', 
        tokenOut: 'DAI', 
        amountIn: '1.0' 
      });
      
      rope.addKnot(knot);
      
      const mockPostSolver = jest.fn().mockResolvedValue({ 
        success: true, 
        txHash: '0x123' 
      });
      
      const mockApi = {
        getSolver: jest.fn().mockResolvedValue({
          'uniswap': {
            metadata: { icon: 'uniswap-icon' },
            schema: { 
              'swap': { type: 'action' },
              'addLiquidity': { type: 'action' }
            }
          }
        }),
        postSolver: mockPostSolver
      };
      
      (PlugAPI as jest.Mock).mockImplementation(() => mockApi);
      
      // Create a new instance to use our updated mock
      rope = new Rope(mockApiKey);
      await rope.initialize();
      rope.addKnot(knot);
      
      const fromAddress = '0xabc123';
      const chainId = 42161; // Arbitrum
      const result = await rope.execute(fromAddress, chainId);
      
      // Check that the correct intent was submitted
      const expectedIntent = {
        chainId: 42161,
        from: fromAddress,
        frequency: 0,
        inputs: [
          {
            protocol: 'uniswap',
            action: 'swap',
            params: expect.objectContaining({
              tokenIn: 'ETH',
              tokenOut: 'DAI',
              amountIn: '1.0'
            })
          }
        ],
        locked: false,
        saved: false,
        status: 'active'
      };
      
      expect(mockPostSolver).toHaveBeenCalledWith(expect.objectContaining(expectedIntent));
      expect(result).toEqual({ success: true, txHash: '0x123' });
    });
  });
  
  describe('api access', () => {
    it('should return the API client instance', () => {
      const mockApiInstance = { mockMethod: jest.fn() };
      (PlugAPI as jest.Mock).mockImplementation(() => mockApiInstance);
      
      rope = new Rope(mockApiKey);
      const api = rope.getApi();
      
      expect(api).toBe(mockApiInstance);
    });
  });
});