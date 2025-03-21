import { Knot } from '../../src/models/knot';

describe('Knot', () => {
  describe('constructor', () => {
    it('should create a new instance with provided id and data', () => {
      const id = 'test-id';
      const data = { key: 'value' };
      const knot = new Knot(id, data);
      
      expect(knot.getId()).toBe(id);
      expect(knot.getAllData()).toEqual(data);
    });
    
    it('should create a new instance with generated id when not provided', () => {
      const knot = new Knot();
      
      expect(knot.getId()).toBeTruthy();
      expect(typeof knot.getId()).toBe('string');
    });
    
    it('should create a new instance with empty data when not provided', () => {
      const knot = new Knot('test-id');
      
      expect(knot.getAllData()).toEqual({});
    });
  });
  
  describe('data management', () => {
    let knot: Knot;
    
    beforeEach(() => {
      knot = new Knot('test-id');
    });
    
    it('should set and get data by key', () => {
      const key = 'testKey';
      const value = 'testValue';
      
      knot.setData(key, value);
      expect(knot.getData(key)).toBe(value);
    });
    
    it('should return undefined for non-existent keys', () => {
      expect(knot.getData('nonExistentKey')).toBeUndefined();
    });
    
    it('should get all data', () => {
      const data = {
        key1: 'value1',
        key2: 'value2',
      };
      
      Object.entries(data).forEach(([key, value]) => {
        knot.setData(key, value);
      });
      
      expect(knot.getAllData()).toEqual(data);
    });
    
    it('should return a copy of the data, not a reference', () => {
      knot.setData('key', 'value');
      
      const data = knot.getAllData();
      data.key = 'modified';
      
      expect(knot.getData('key')).toBe('value');
    });
    
    it('should support method chaining for setData', () => {
      const result = knot
        .setData('key1', 'value1')
        .setData('key2', 'value2');
      
      expect(result).toBe(knot);
      expect(knot.getData('key1')).toBe('value1');
      expect(knot.getData('key2')).toBe('value2');
    });
  });
});