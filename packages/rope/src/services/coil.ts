import * as Types from '../types/generated';
import { CoilCompatibilityError } from './errors';
import { ParsedSentence } from './sentence';
import { Result, LRUCache } from '../utils';
import { InputType } from '../types/rope';

/**
 * Coil reference format (used to link values between knots)
 */
export interface CoilReference {
  name: string;
  knotIndex?: number;
  type: string;
}

/**
 * Result of coil validation
 */
export type CoilValidationResult = Result<void, Record<string, string>>;

/**
 * Options for creating a coil proxy
 */
export interface CoilProxyOptions {
  /**
   * Whether to cache coil values after they're first computed
   * Default: true
   */
  cache?: boolean;
  
  /**
   * Whether to return a placeholder value for missing coils
   * Default: false
   */
  usePlaceholders?: boolean;
  
  /**
   * Custom placeholder format, will replace {name} with coil name
   * Default: "[Missing Coil: {name}]"
   */
  placeholderFormat?: string;
}

// LRUCache is imported from ../utils

/**
 * Service for handling coil operations and compatibility
 */
export class CoilService {
  // Regular expression to match coil references
  private static readonly COIL_REGEX = /^<-\{([^@}]+)(?:@(\d+))?\}$/;
  
  // Cache for parsed coil references with a max size of 300 and 1-hour TTL
  private static readonly referenceCache = new LRUCache<string, CoilReference | null>({
    maxSize: 300,
    ttl: 60 * 60 * 1000 // 1 hour
  });
  
  /**
   * Parse a coil reference from a string value
   * @param value The coil reference string (e.g. "<-{amount@2}" or "<-{amount}")
   * @returns Parsed coil reference or null if not a valid reference
   */
  public parseCoilReference(value: string): CoilReference | null {
    if (!this.isCoilReference(value)) {
      return null;
    }
    
    // Check cache first
    const cached = CoilService.referenceCache.get(value);
    if (cached !== undefined) {
      return cached;
    }
    
    // Extract the coil name and optional knot index
    const match = value.match(CoilService.COIL_REGEX);
    if (!match) {
      CoilService.referenceCache.set(value, null);
      return null;
    }
    
    const name = match[1];
    const knotIndex = match[2] ? parseInt(match[2]) : undefined;
    
    const result = { name, knotIndex, type: 'unknown' };
    CoilService.referenceCache.set(value, result);
    
    return result;
  }
  
  /**
   * Clear the coil reference parsing cache
   */
  public clearReferenceCache(): void {
    CoilService.referenceCache.clear();
  }
  
  /**
   * Get the current size of the reference cache
   * @returns Number of cached parsed coil references
   */
  public getReferenceCacheSize(): number {
    return CoilService.referenceCache.size;
  }

  /**
   * Create a lazy-loading proxy for coil values
   * This allows coil values to be computed only when they are actually accessed
   * 
   * @param availableCoils Base coil values or function to compute them
   * @param options Configuration options for the proxy
   * @returns A proxy object that computes coil values on demand
   */
  public createCoilProxy(
    availableCoils: Record<string, any> | (() => Record<string, any>),
    options: CoilProxyOptions = {}
  ): Record<string, any> {
    // Set default options
    const {
      cache = true,
      usePlaceholders = false,
      placeholderFormat = "[Missing Coil: {name}]"
    } = options;
    
    // Create cache for computed values if caching is enabled
    const valueCache = new Map<string, any>();
    
    return new Proxy({}, {
      get: (_, prop) => {
        const coilName = String(prop);
        
        // Return from cache if available
        if (cache && valueCache.has(coilName)) {
          return valueCache.get(coilName);
        }
        
        // Get coils object (might be a function)
        const coils = typeof availableCoils === 'function' 
          ? availableCoils() 
          : availableCoils;
        
        // Compute the value
        let value: any;
        
        if (coilName in coils) {
          value = coils[coilName];
        } else if (usePlaceholders) {
          value = placeholderFormat.replace('{name}', coilName);
        } else {
          // Return undefined for missing coils
          return undefined;
        }
        
        // Cache the value if caching is enabled
        if (cache) {
          valueCache.set(coilName, value);
        }
        
        return value;
      },
      
      has: (_, prop) => {
        const coilName = String(prop);
        
        // Check cache first
        if (cache && valueCache.has(coilName)) {
          return true;
        }
        
        // Check actual coils object
        const coils = typeof availableCoils === 'function'
          ? availableCoils()
          : availableCoils;
          
        return coilName in coils;
      }
    });
  }
  
  /**
   * Check if a value is a coil reference
   * @param value Value to check
   * @returns Whether the value is a coil reference
   */
  public isCoilReference(value: string | any): boolean {
    if (typeof value !== 'string') {
      return false;
    }
    
    return CoilService.COIL_REGEX.test(value);
  }
  
  /**
   * Format a value as a coil reference
   * @param name Coil name
   * @param knotIndex Optional knot index
   * @returns Formatted coil reference string
   */
  public formatCoilReference(name: string, knotIndex?: number): string {
    if (knotIndex !== undefined) {
      return `<-{${name}@${knotIndex}}`;
    }
    return `<-{${name}}`;
  }
  
  /**
   * Get all coils from an action schema
   * @param actionSchema The action schema containing coil definitions
   * @returns Map of coil names to their types
   */
  public getCoilsFromSchema(actionSchema?: Types.ActionsSchema): Record<string, string> {
    if (!actionSchema || !actionSchema.coils) {
      return {};
    }
    
    return actionSchema.coils;
  }
  
  /**
   * Check if a coil type is compatible with an input type
   * @param coilType Type of the coil
   * @param inputType Type expected by the input
   * @returns Whether the coil is compatible with the input
   */
  public isCoilCompatible(coilType: string, inputType: string): boolean {
    // Basic type compatibility check
    if (coilType === inputType) {
      return true;
    }
    
    // Handle numeric types
    if (
      (coilType === 'number' || coilType === 'integer' || coilType === 'float') &&
      (inputType === 'number' || inputType === 'integer' || inputType === 'float')
    ) {
      return true;
    }
    
    // Handle string types
    if (
      (coilType === 'string' || coilType === 'address' || coilType === 'hash') &&
      (inputType === 'string')
    ) {
      return true;
    }
    
    // Handle address types
    if (
      (coilType === 'address') &&
      (inputType === 'address')
    ) {
      return true;
    }
    
    // Handle token types
    if (
      (coilType === 'token' || coilType === 'address') &&
      (inputType === 'token')
    ) {
      return true;
    }
    
    return false;
  }
  
  /**
   * Validate coil compatibility
   * @param coilType Type of the coil
   * @param inputType Type expected by the input
   * @param coilName Name of the coil (for error reporting)
   * @returns Result indicating compatibility
   */
  public validateCoilCompatibility(coilType: string, inputType: string, coilName: string): Result<void> {
    if (!this.isCoilCompatible(coilType, inputType)) {
      return Result.failure(
        `Coil "${coilName}" of type "${coilType}" is not compatible with input of type "${inputType}"`
      );
    }
    return Result.success(undefined);
  }
  
  /**
   * Get compatible coils from a set of available coils
   * @param availableCoils Map of available coil names to their types
   * @param requiredType Type that the coil must be compatible with
   * @returns Map of compatible coil names to their types
   */
  public getCompatibleCoils(
    availableCoils: Record<string, string>,
    requiredType: InputType
  ): Record<string, string> {
    // For object types, we need to extract the base type
    const typeToCheck = this.extractBaseType(requiredType);
    
    return Object.entries(availableCoils)
      .filter(([_, type]) => this.isCoilCompatible(type, typeToCheck))
      .reduce((acc, [name, type]) => {
        acc[name] = type;
        return acc;
      }, {} as Record<string, string>);
  }
  
  /**
   * Extract the base type from an InputType for compatibility checking
   * @param type The InputType to extract from
   * @returns A string representation of the base type
   */
  private extractBaseType(type: InputType): string {
    // If it's a string, use it directly
    if (typeof type === 'string') {
      return type;
    }
    
    // If it's a constant type, we use the type "constant"
    if ('constant' in type) {
      return 'string';
    }
    
    // For compound types, use the base type
    if ('baseType' in type) {
      return this.extractBaseType(type.baseType);
    }
    
    // For union types, use the first type
    if ('types' in type && type.types.length > 0) {
      return this.extractBaseType(type.types[0]);
    }
    
    // For conditional types, we don't know which branch will be taken,
    // so default to string for maximum compatibility
    if ('left' in type) {
      return 'string';
    }
    
    // Default to string for unknown types
    return 'string';
  }
  
  /**
   * Extract all coil references from a values object
   * @param values Object containing input values
   * @returns Array of coil reference strings found
   */
  public extractCoilReferences(values: Record<string, any>): string[] {
    return Object.values(values)
      .filter(value => typeof value === 'string' && this.isCoilReference(value))
      .map(value => value as string);
  }
  
  /**
   * Resolve coil references in a values object
   * @param values Object containing input values with coil references
   * @param availableCoils Map of available coil names to their values
   * @param options Optional configuration for coil resolution
   * @returns New values object with resolved coil references
   */
  public resolveCoilReferences(
    values: Record<string, any>,
    availableCoils: Record<string, any>,
    options?: CoilProxyOptions
  ): Record<string, any> {
    // Create a lazy-loading coil proxy
    const coilProxy = this.createCoilProxy(availableCoils, {
      cache: true,
      usePlaceholders: true,
      ...options
    });
    
    const result = { ...values };
    
    // Process each value that might be a coil reference
    Object.entries(result).forEach(([key, value]) => {
      if (typeof value === 'string' && this.isCoilReference(value)) {
        const coilRef = this.parseCoilReference(value);
        if (coilRef) {
          // Use the proxy to lazily resolve the coil value
          const coilValue = coilProxy[coilRef.name];
          if (coilValue !== undefined) {
            // Replace coil reference with actual value
            result[key] = coilValue;
          }
        }
      }
    });
    
    return result;
  }
  
  /**
   * Resolve coil references in a sentence with more advanced handling
   * @param sentence The sentence or template containing possible coil references
   * @param values Input values with possible coil references
   * @param availableCoils Available coil values
   * @param options Optional configuration for coil resolution
   * @returns New values with resolved coil references
   */
  public resolveCoilReferencesInSentence(
    sentence: string, 
    values: Record<string, string>, 
    availableCoils: Record<string, string>,
    options?: CoilProxyOptions
  ): Record<string, string> {
    // Create a lazy-loading coil proxy with type-specific handling
    const coilProxy = this.createCoilProxy(() => {
      // This function will be called only when a coil value is actually requested
      const processedCoils: Record<string, string> = {};
      
      // Process each available coil to handle different types
      Object.entries(availableCoils).forEach(([name, value]) => {
        const coilType = typeof value;
        
        if (coilType === 'string' || coilType === 'number' || coilType === 'boolean') {
          // For primitive types, use the value directly
          processedCoils[name] = String(value);
        } else {
          // For complex types, stringify or provide a representation
          processedCoils[name] = `Coil: ${name}`;
        }
      });
      
      return processedCoils;
    }, {
      cache: true,
      usePlaceholders: true,
      placeholderFormat: "[Missing Coil: {name}]",
      ...options
    });
    
    const resolvedValues = { ...values };
    
    // Process each value that might be a coil reference
    Object.entries(resolvedValues).forEach(([key, value]) => {
      if (typeof value === 'string' && this.isCoilReference(value)) {
        const coilRef = this.parseCoilReference(value);
        if (coilRef) {
          // Use the proxy to lazily resolve the coil value
          resolvedValues[key] = coilProxy[coilRef.name];
        }
      }
    });
    
    return resolvedValues;
  }
  
  /**
   * Validate that all coil references in values are valid and compatible
   * @param values Values object potentially containing coil references
   * @param inputTypes Map of input names to their expected types
   * @param availableCoils Map of available coil names to their types
   * @returns Validation result with errors if any
   */
  public validateCoilUsage(
    values: Record<string, any>,
    inputTypes: Record<string, string>,
    availableCoils: Record<string, string>
  ): CoilValidationResult {
    const errors: Record<string, string> = {};
    
    // Check each value for coil references
    Object.entries(values).forEach(([key, value]) => {
      if (typeof value === 'string' && this.isCoilReference(value)) {
        const coilRef = this.parseCoilReference(value);
        if (coilRef) {
          // Verify coil exists
          if (!(coilRef.name in availableCoils)) {
            errors[key] = `Coil "${coilRef.name}" not found`;
            return;
          }
          
          // Verify type compatibility
          const coilType = availableCoils[coilRef.name];
          const inputType = inputTypes[key];
          
          if (!this.isCoilCompatible(coilType, inputType)) {
            errors[key] = `Coil "${coilRef.name}" of type "${coilType}" is not compatible with input "${key}" of type "${inputType}"`;
          }
        }
      }
    });
    
    return Object.keys(errors).length === 0
      ? Result.success(undefined)
      : Result.failure(errors);
  }
  
  /**
   * Validate coil references with specific type checking
   * @param coilReferences Array of coil reference strings to validate
   * @param requiredTypes Map of coil names to their required types
   * @param availableCoils Map of available coil names to their types
   * @returns Validation result with errors if any
   */
  public validateCoilReferences(
    coilReferences: string[],
    requiredTypes: Record<string, string>,
    availableCoils: Record<string, string>
  ): CoilValidationResult {
    const errors: Record<string, string> = {};
    
    coilReferences.forEach(ref => {
      const coilRef = this.parseCoilReference(ref);
      if (coilRef) {
        const coilName = coilRef.name;
        
        // Check if coil exists
        if (!(coilName in availableCoils)) {
          errors[coilName] = `Coil "${coilName}" not found`;
          return;
        }
        
        // Check type compatibility if type is specified
        const requiredType = requiredTypes[coilName];
        if (requiredType) {
          const availableType = availableCoils[coilName];
          if (!this.isCoilCompatible(availableType, requiredType)) {
            errors[coilName] = `Coil "${coilName}" of type "${availableType}" is not compatible with required type "${requiredType}"`;
          }
        }
      }
    });
    
    return Object.keys(errors).length === 0
      ? Result.success(undefined)
      : Result.failure(errors);
  }
  
  /**
   * Find all inputs in a parsed sentence that could accept coils
   * @param parsedSentence The parsed sentence
   * @param availableCoils Map of available coil names to their types
   * @returns Map of input indices to arrays of compatible coil names
   */
  public getCompatibleCoilsForParsedSentence(
    parsedSentence: ParsedSentence,
    availableCoils: Record<string, string>
  ): Record<number, string[]> {
    const result: Record<number, string[]> = {};
    
    // For each input, find compatible coils
    parsedSentence.inputs.forEach(input => {
      const inputType = input.type || 'string';
      const compatibleCoils = this.getCompatibleCoils(availableCoils, inputType);
      
      // Add to result if there are any compatible coils
      const coilNames = Object.keys(compatibleCoils);
      if (coilNames.length > 0) {
        result[input.index] = coilNames;
      }
    });
    
    return result;
  }
}