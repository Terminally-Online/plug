import { 
  InputType, 
  InputState, 
  InputValues, 
  ComparisonOperator, 
  ComparisonValue,
  EvmType,
  ConstantType,
  SimpleUnionType
} from '../types/rope';

import { Result, LRUCache } from '../utils';

/**
 * Input definition extracted from a sentence
 */
export interface SentenceInput {
  index: number;
  name?: string;
  type?: InputType;
  defaultValue?: string;
  required?: boolean;
  dependentOn?: number;
  delimiter?: string;
  options?: string[];
  description?: string;
}

/**
 * Parsed sentence result
 */
export interface ParsedSentence {
  original: string;
  raw?: string; // For Cord compatibility
  parts: string[];
  inputs: SentenceInput[];
  template: string;
  values?: Map<number, InputState>;
}

/**
 * Use the Result type for validation results
 */
export type ValidationResult = Result<void>;

/**
 * Custom result type for setValue operations
 */
export interface SetValueResult {
  success: boolean;
  value: Map<number, InputState>;
  error?: string;
}

/**
 * Service for parsing and working with action sentences
 */
export class SentenceService {
  // Regular expression to extract input placeholders from sentences
  private static readonly INPUT_REGEX = /\{(\d+)(?:=>(\d+))?\}/g;
  
  // Cache for parsed sentences to avoid repeated parsing of the same sentence
  // Limit to 500 entries with a 1-hour TTL
  private static readonly parseCache = new LRUCache<string, ParsedSentence>({ 
    maxSize: 500, 
    ttl: 60 * 60 * 1000 // 1 hour
  });
  
  // Enhanced regex to support Cord's format with type information
  private static readonly ENHANCED_INPUT_REGEX = /\{(?:(\d+)=>)?(\d+)(?::([a-zA-Z0-9_]+))?(?:\|([\w\s().,\-+\/<>=!?:{}[\]"']+))?(?:\|([,;|]+))?\}/g;
  
  /**
   * Parse a type string into an InputType
   * @param typeString The type string to parse
   * @returns The parsed input type and optional default value
   */
  public parseTypeString(typeString: string): { type: InputType; defaultValue?: string } {
    // Handle constant types - "constant:value"
    if (typeString.startsWith('constant:')) {
      const constantValue = typeString.substring(9);
      return {
        type: { constant: constantValue },
        defaultValue: constantValue
      };
    }
    
    // Handle conditional types - "condition:1==ETH?address:string"
    if (typeString.startsWith('condition:')) {
      const conditionStr = typeString.substring(10);
      
      // Parse out the condition, trueType, and falseType
      const conditionMatch = conditionStr.match(/(.+?)([=<>!]+)(.+?)\?(.+?):(.+)/);
      if (conditionMatch) {
        const [, leftStr, operatorStr, rightStr, trueTypeStr, falseTypeStr] = conditionMatch;
        
        // Parse left side (could be value or reference)
        let left: ComparisonValue;
        if (leftStr.startsWith('@')) {
          // It's a reference to another input
          const refIndexStr = leftStr.substring(1);
          const refIndex = parseInt(refIndexStr);
          left = { reference: refIndex };
        } else {
          // It's a direct value
          left = leftStr;
        }
        
        // Parse right side (could be value or reference)
        let right: ComparisonValue;
        if (rightStr.startsWith('@')) {
          // It's a reference to another input
          const refIndexStr = rightStr.substring(1);
          const refIndex = parseInt(refIndexStr);
          right = { reference: refIndex };
        } else {
          // It's a direct value
          right = rightStr;
        }
        
        // Map the operator string to ComparisonOperator
        let operator: ComparisonOperator;
        switch (operatorStr) {
          case '==': operator = '=='; break;
          case '!=': operator = '!='; break;
          case '>': operator = '>'; break;
          case '<': operator = '<'; break;
          case '>=': operator = '>='; break;
          case '<=': operator = '<='; break;
          default: operator = '=='; // Default to equality
        }
        
        // Parse true and false types recursively
        const { type: trueType } = this.parseTypeString(trueTypeStr);
        const { type: falseType } = this.parseTypeString(falseTypeStr);
        
        return {
          type: {
            left,
            operator,
            right,
            trueType,
            falseType
          }
        };
      }
    }
    
    // Handle union types - "union:address|uint256"
    if (typeString.startsWith('union:')) {
      const unionStr = typeString.substring(6);
      const typeStrings = unionStr.split('|');
      
      const types: InputType[] = typeStrings.map(ts => {
        const { type } = this.parseTypeString(ts);
        return type;
      });
      
      return {
        type: { types }
      };
    }
    
    // Handle compound types - "compound:address|uint256,string"
    if (typeString.startsWith('compound:')) {
      const compoundStr = typeString.substring(9);
      const [baseTypeStr, ...metadataTypeStrs] = compoundStr.split(',');
      
      const { type: baseType } = this.parseTypeString(baseTypeStr);
      
      const metadata: (EvmType | ConstantType | SimpleUnionType)[] = metadataTypeStrs.map(mts => {
        const { type } = this.parseTypeString(mts);
        // Convert type to acceptable metadata type
        if (typeof type === 'string') {
          return type as EvmType;
        } else if ('constant' in type) {
          return type;
        } else if ('types' in type && !('trueType' in type)) {
          return type as SimpleUnionType;
        }
        // Fall back to string if not valid metadata type
        return 'string';
      });
      
      return {
        type: {
          baseType: baseType as EvmType | ConstantType | SimpleUnionType,
          metadata
        }
      };
    }
    
    // Handle default values - "string:default"
    const parts = typeString.split(':');
    const type = parts[0] as EvmType; // Base types are treated as EVM types
    const defaultValue = parts.length > 1 ? parts[1] : undefined;
    
    return { type, defaultValue };
  }
  
  /**
   * Parse a sentence into parts and extract input definitions
   * @param sentence The sentence to parse (e.g. "Swap {0} for {1} on {2}")
   * @returns The parsed sentence with extracted inputs
   */
  public parseSentence(sentence: string): ParsedSentence {
    // Return empty result for empty sentence
    if (!sentence) {
      return { original: '', raw: '', parts: [], inputs: [], template: '' };
    }
    
    // Check if this sentence is already in the cache
    const cachedResult = SentenceService.parseCache.get(sentence);
    if (cachedResult) {
      return cachedResult;
    }
    
    // Split the sentence into parts by input placeholders
    const parts: string[] = [];
    const inputs: SentenceInput[] = [];
    const values = new Map<number, InputState>();
    
    // First try the enhanced regex for Cord compatibility
    let usedEnhancedRegex = false;
    let lastIndex = 0;
    let match: RegExpExecArray | null;
    
    try {
      // Reset the regex to start from the beginning
      SentenceService.ENHANCED_INPUT_REGEX.lastIndex = 0;
      
      while ((match = SentenceService.ENHANCED_INPUT_REGEX.exec(sentence)) !== null) {
        // Add text before this match as a part
        if (match.index > lastIndex) {
          parts.push(sentence.substring(lastIndex, match.index));
        }
        
        // Extract match groups
        const [fullMatch, dependentOnStr, indexStr, typeStr, defaultValueStr, delimiterStr] = match;
        
        // Parse input index and dependency
        const inputIndex = parseInt(indexStr);
        const dependentOn = dependentOnStr ? parseInt(dependentOnStr) : undefined;
        
        // Add the placeholder as a standardized part
        parts.push(`{${inputIndex}}`);
        
        // Process type information if present
        let inputType: InputType = 'string';
        let defaultValue: string | undefined = defaultValueStr;
        let required = true;
        
        if (typeStr) {
          try {
            // Parse the type string
            const parseResult = this.parseTypeString(typeStr);
            inputType = parseResult.type;
            
            // Use explicit default value if provided
            if (defaultValueStr) {
              defaultValue = defaultValueStr;
            } 
            // Otherwise use any default from the type parsing
            else if (parseResult.defaultValue) {
              defaultValue = parseResult.defaultValue;
            }
            
            // If the type is a constant, set it in values
            if (typeof inputType === 'object' && 'constant' in inputType) {
              values.set(inputIndex, {
                value: inputType.constant,
                isDisabled: true
              });
              required = false;
            } 
            // Add default value to values if present
            else if (defaultValue) {
              values.set(inputIndex, { value: defaultValue });
            }
          } catch (error) {
            console.error(`Error parsing type: ${typeStr}`, error);
            // Default to string type on error
            inputType = 'string';
          }
        }
        
        // Add input if it doesn't exist yet
        if (!inputs.some(i => i.index === inputIndex)) {
          inputs.push({
            index: inputIndex,
            name: `input${inputIndex}`,
            type: inputType,
            defaultValue,
            required,
            dependentOn,
            delimiter: delimiterStr
          });
        }
        
        // Update last index to continue from end of match
        lastIndex = match.index + fullMatch.length;
        usedEnhancedRegex = true;
      }
      
      // Add any remaining text
      if (lastIndex < sentence.length) {
        parts.push(sentence.substring(lastIndex));
      }
    } catch (error) {
      console.error('Error in enhanced regex parsing', error);
      usedEnhancedRegex = false;
    }
    
    // Fall back to simple regex if enhanced didn't match or had errors
    if (!usedEnhancedRegex) {
      // Reset values from previous attempt
      parts.length = 0;
      inputs.length = 0;
      values.clear();
      lastIndex = 0;
      
      // Use the simpler regex
      SentenceService.INPUT_REGEX.lastIndex = 0;
      
      while ((match = SentenceService.INPUT_REGEX.exec(sentence)) !== null) {
        // Add text before this match as a part
        if (match.index > lastIndex) {
          parts.push(sentence.substring(lastIndex, match.index));
        }
        
        // Add the placeholder as a part
        parts.push(match[0]);
        
        // Extract input data
        const inputIndex = parseInt(match[1]);
        const dependentOn = match[2] ? parseInt(match[2]) : undefined;
        
        // Add input if it doesn't exist yet
        if (!inputs.some(i => i.index === inputIndex)) {
          inputs.push({
            index: inputIndex,
            name: `input${inputIndex}`,
            type: 'string', // Default to string type
            required: true,
            dependentOn
          });
        }
        
        // Update last index to continue from end of match
        lastIndex = match.index + match[0].length;
      }
      
      // Add any remaining text
      if (lastIndex < sentence.length) {
        parts.push(sentence.substring(lastIndex));
      }
    }
    
    // Handle conditional types in a second pass, now that all inputs are defined
    inputs.forEach(input => {
      if (input.type && typeof input.type === 'object' && 'left' in input.type) {
        const resolvedType = this.evaluateConditionalType(input.type, input.index, values);
        
        // If type resolves to a constant, set that value
        if (typeof resolvedType === 'object' && 'constant' in resolvedType) {
          values.set(input.index, {
            value: resolvedType.constant,
            isDisabled: true
          });
        }
      }
    });
    
    // Create result object
    const result: ParsedSentence = {
      original: sentence,
      raw: sentence,
      template: sentence,
      parts,
      inputs: inputs.sort((a, b) => a.index - b.index),
      values
    };
    
    // Store in cache for future use
    SentenceService.parseCache.set(sentence, result);
    
    return result;
  }
  
  /**
   * Clear the sentence parsing cache
   * This can be useful when freeing memory or after schema updates
   */
  public clearParseCache(): void {
    SentenceService.parseCache.clear();
  }
  
  /**
   * Get the current size of the parse cache
   * @returns Number of cached parsed sentences
   */
  public getParseCacheSize(): number {
    return SentenceService.parseCache.size;
  }
  
  /**
   * Generate a set of empty values for a parsed sentence
   * @param parsed The parsed sentence
   * @returns An object with empty values for each input
   */
  public generateEmptyValues(parsed: ParsedSentence): Record<string, string | undefined> {
    return parsed.inputs.reduce((values, input) => {
      values[input.index.toString()] = undefined;
      return values;
    }, {} as Record<string, string | undefined>);
  }
  
  /**
   * Format a sentence by replacing placeholders with values
   * @param parsed The parsed sentence
   * @param values Values for the inputs
   * @returns The formatted sentence with values substituted
   */
  public formatSentence(parsed: ParsedSentence, values: Record<string, any>): string {
    // First pass - resolve any conditional/constant inputs based on current values
    let resolvedValues = { ...values };
    
    // Convert object values to Map for processing
    const valuesMap = new Map<number, { value: string }>();
    Object.entries(values).forEach(([key, value]) => {
      const numKey = Number(key);
      if (!isNaN(numKey) && value !== undefined) {
        valuesMap.set(numKey, { value: String(value) });
      }
    });
    
    // Check for and resolve conditional types
    parsed.inputs.forEach(input => {
      if (input.type && typeof input.type === 'object' && 'left' in input.type) {
        // Get the resolved type based on current values
        const resolvedType = this.evaluateConditionalType(input.type, input.index, valuesMap);
        
        // If the resolved type is a constant, use its value
        if (typeof resolvedType === 'object' && 'constant' in resolvedType) {
          resolvedValues[input.index] = resolvedType.constant;
        }
      }
    });
    
    // Now handle parts
    return parsed.parts.map(part => {
      // Handle regular input placeholders
      const match = part.match(/\{(\d+)\}/);
      if (!match) {
        return part;
      }
      
      const inputIndex = parseInt(match[1]);
      const value = resolvedValues[inputIndex];
      
      // Format input value
      return value !== undefined ? String(value) : `[${inputIndex}]`;
    }).join('');
  }
  
  /**
   * Split template parts into fragments for rendering
   * Similar to Cord's template part processing
   * @param parsed The parsed sentence
   * @returns Array of string fragments for rendering
   */
  public getTemplateParts(parsed: ParsedSentence): string[] {
    const parts: string[] = [];
    
    if (!parsed) return parts;
    
    // Process each part
    parsed.parts.forEach(part => {
      if (part.match(/\{(\d+)\}/)) {
        // This is a placeholder - keep as is
        parts.push(part);
      } else {
        // Split non-placeholder text by whitespace, preserving whitespace
        const fragments = part.split(/(\s+)/g);
        parts.push(...fragments);
      }
    });
    
    return parts;
  }
  
  /**
   * Check if all required inputs have values
   * @param parsed The parsed sentence
   * @param values Values for the inputs
   * @returns Whether all required inputs have values
   */
  public isComplete(parsed: ParsedSentence, values: Record<string, any>): boolean {
    return parsed.inputs
      .filter(input => input.required)
      .every(input => values[input.index] !== undefined && values[input.index] !== '');
  }
  
  /**
   * Validate an input value based on its type
   * @param value The value to validate
   * @param type The expected type of the value
   * @returns A validation result with success flag and optional error message
   */
  public validateInput(value: string, type: InputType = 'string'): ValidationResult {
    // Handle empty values
    if (!value || value.trim() === '') {
      // Allow empty for boolean and null types
      const typeStr = String(type).toLowerCase();
      if (typeStr === 'bool' || typeStr === 'boolean' || typeStr === 'null') {
        return Result.success(undefined);
      }
      return Result.failure('Value is required');
    }
    
    // Check string type first
    if (typeof type === 'string') {
      // Handle primitive types
      const typeLower = type.toLowerCase();
      
      // Handle uint/int types
      if (typeLower.startsWith('uint') || typeLower.startsWith('int')) {
        const isUint = typeLower.startsWith('uint');
        const num = Number(value);
        
        if (isNaN(num)) {
          return Result.failure(`Value must be a${isUint ? ' positive' : 'n'} integer`);
        }
        
        if (!Number.isInteger(num)) {
          return Result.failure('Value must be an integer');
        }
        
        if (isUint && num < 0) {
          return Result.failure('Value must be a positive integer');
        }
        
        // Add bit size validation if needed
        // const bitSize = parseInt(typeLower.substring(isUint ? 4 : 3));
        // Could validate against max values by bit size
        
        return Result.success(undefined);
      }
      
      // Handle other primitive types
      switch (typeLower) {
        case 'float':
          if (isNaN(Number(value))) {
            return Result.failure('Value must be a number');
          }
          break;
          
        case 'address':
          // Simple Ethereum address validation
          if (!/^0x[a-fA-F0-9]{40}$/.test(value)) {
            return Result.failure('Invalid address format');
          }
          break;
          
        case 'bool':
        case 'boolean':
          if (value !== 'true' && value !== 'false' && value !== '1' && value !== '0') {
            return Result.failure('Value must be true or false');
          }
          break;
          
        case 'bytes':
          // Simple hex validation for bytes
          if (!/^(0x)?[a-fA-F0-9]*$/.test(value)) {
            return Result.failure('Invalid bytes format');
          }
          break;
      }
      
      return Result.success(undefined);
    }
    
    // Handle object types
    if (typeof type === 'object') {
      // Constant type
      if ('constant' in type) {
        if (value !== type.constant) {
          return Result.failure(`Value must be ${type.constant}`);
        }
        return Result.success(undefined);
      }
      
      // Union type
      if ('types' in type && Array.isArray(type.types)) {
        // For union types, validate against each possible type
        // If any validation passes, the whole validation passes
        for (const subType of type.types) {
          const result = this.validateInput(value, subType);
          if (result.success) {
            return Result.success(undefined);
          }
        }
        return Result.failure('Value does not match any accepted type');
      }
      
      // Compound type
      if ('baseType' in type) {
        // Validate against the base type first
        const baseResult = this.validateInput(value, type.baseType);
        if (!baseResult.success) {
          return baseResult;
        }
        // Additional metadata validation could be added here
        return Result.success(undefined);
      }
      
      // Comparison type - this should be handled elsewhere during value setting
      if ('left' in type && 'operator' in type && 'right' in type) {
        return Result.success(undefined);
      }
    }
    
    // Default success for unknown types
    return Result.success(undefined);
  }
  
  /**
   * Get placeholder text for an input type
   * @param type The input type
   * @returns A suitable placeholder text
   */
  public getPlaceholder(type: InputType = 'string'): string {
    // Handle string type identifiers
    if (typeof type === 'string') {
      const typeLower = type.toLowerCase();
      
      // Handle uint/int types
      if (typeLower.startsWith('uint') || typeLower.startsWith('int')) {
        return typeLower.startsWith('uint') 
          ? 'Enter a positive integer' 
          : 'Enter an integer';
      }
      
      // Handle primitive types
      switch (typeLower) {
        case 'number':
        case 'float':
          return 'Enter a number';
        case 'address':
          return 'Enter an address (0x...)';
        case 'token':
          return 'Select a token';
        case 'bool':
        case 'boolean':
          return 'Yes/No';
        case 'bytes':
          return 'Enter hex data';
        default:
          return 'Enter a value';
      }
    }
    
    // Handle object types
    if (typeof type === 'object') {
      // Constant type
      if ('constant' in type) {
        return type.constant;
      }
      
      // Union type
      if ('types' in type && Array.isArray(type.types)) {
        // Use the placeholder of the first type
        return this.getPlaceholder(type.types[0]);
      }
      
      // Compound type
      if ('baseType' in type) {
        return this.getPlaceholder(type.baseType);
      }
      
      // Comparison type
      if ('left' in type && 'operator' in type && 'right' in type) {
        // Try to indicate the conditional nature
        return 'Depends on other values';
      }
    }
    
    // Default placeholder
    return 'Enter a value';
  }
  
  /**
   * Determine if an input should be rendered based on dependencies and conditional types
   * @param input The input to check
   * @param allInputs All inputs in the sentence
   * @param getValueFn Function to get value for an input index
   * @returns Whether the input should be rendered
   */
  public shouldRenderInput(
    input: SentenceInput | InputType,
    allInputs: SentenceInput[],
    getValueFn: (index: number) => { value: string } | undefined
  ): boolean {
    // If input is a string type identifier
    if (typeof input === 'string') {
      // Find the input with the given type
      const matchingInput = allInputs.find(inp => 
        typeof inp.type === 'string' && inp.type === input
      );
      
      // Check dependency if found
      if (matchingInput?.dependentOn !== undefined) {
        const dependencyValue = getValueFn(matchingInput.dependentOn);
        if (!dependencyValue || !dependencyValue.value) {
          return false;
        }
      }
      
      return true;
    }
    
    // If input is an actual input object
    if ('index' in input) {
      // Early exit for null type
      if (input.type === 'null') {
        return false;
      }
      
      // Check if it has a direct dependency
      if (input.dependentOn !== undefined) {
        const dependencyValue = getValueFn(input.dependentOn);
        if (!dependencyValue || !dependencyValue.value) {
          return false;
        }
      }
      
      // Check for conditional types
      if (input.type && typeof input.type === 'object') {
        // Convert object values to Map for processing
        const valuesMap = new Map<number, { value: string }>();
        for (const otherInput of allInputs) {
          const value = getValueFn(otherInput.index);
          if (value) {
            valuesMap.set(otherInput.index, value);
          }
        }
        
        // If it's a conditional type, evaluate it
        if ('left' in input.type) {
          const resolvedType = this.evaluateConditionalType(input.type, input.index, valuesMap);
          
          // Don't render if it resolves to null
          if (resolvedType === 'null') {
            return false;
          }
          
          // For constant types, we typically don't render UI inputs
          if (typeof resolvedType === 'object' && 'constant' in resolvedType) {
            return false;
          }
        }
        
        // For constant types, don't render
        if ('constant' in input.type) {
          return false;
        }
      }
      
      return true;
    }
    
    // If input is a type object
    if (typeof input === 'object') {
      // Don't render constant types
      if ('constant' in input) {
        return false;
      }
      
      // For conditional types, we need to evaluate them
      if ('left' in input) {
        // Convert object values to Map for processing
        const valuesMap = new Map<number, { value: string }>();
        for (const otherInput of allInputs) {
          const value = getValueFn(otherInput.index);
          if (value) {
            valuesMap.set(otherInput.index, value);
          }
        }
        
        // For index, just use an arbitrary value since we're evaluating in isolation
        const dummyIndex = -1;
        const resolvedType = this.evaluateConditionalType(input, dummyIndex, valuesMap);
        
        // Don't render if it resolves to null
        if (resolvedType === 'null') {
          return false;
        }
        
        // For constant types, we typically don't render UI inputs
        if (typeof resolvedType === 'object' && 'constant' in resolvedType) {
          return false;
        }
      }
    }
    
    // Default to rendering
    return true;
  }
  
  /**
   * Compare two values using the specified operator
   * @param left First value for comparison
   * @param operator Comparison operator
   * @param right Second value for comparison
   * @param values Current value map for resolving references
   * @returns Result of the comparison
   */
  public compareValues(
    left: ComparisonValue,
    operator: ComparisonOperator,
    right: ComparisonValue,
    values: Map<number, { value: string }>
  ): boolean {
    // Resolve values (could be direct values or references)
    const leftValue = this.resolveComparisonValue(left, values);
    const rightValue = this.resolveComparisonValue(right, values);
    
    // Return false if either value couldn't be resolved
    if (leftValue === undefined || rightValue === undefined) {
      return false;
    }
    
    // Convert to numbers if possible for numeric comparisons
    const leftNum = !isNaN(Number(leftValue)) ? Number(leftValue) : leftValue;
    const rightNum = !isNaN(Number(rightValue)) ? Number(rightValue) : rightValue;
    
    // Compare values based on operator
    switch (operator) {
      case "==":
        return leftNum === rightNum;
      case "!=":
        return leftNum !== rightNum;
      case ">":
        return leftNum > rightNum;
      case "<":
        return leftNum < rightNum;
      case ">=":
        return leftNum >= rightNum;
      case "<=":
        return leftNum <= rightNum;
      default:
        return false;
    }
  }
  
  /**
   * Resolve a comparison value (direct value or reference)
   * @param value The comparison value to resolve
   * @param values Current value map for resolving references
   * @returns The resolved string value or undefined if can't be resolved
   */
  private resolveComparisonValue(
    value: ComparisonValue,
    values: Map<number, { value: string }>
  ): string | undefined {
    // Direct string value
    if (typeof value === 'string') {
      return value;
    }
    
    // Reference to another input
    if (typeof value === 'object' && 'reference' in value) {
      const refValue = values.get(value.reference);
      if (!refValue) return undefined;
      
      // If we need a specific part of the value and it contains a delimiter
      if (value.part !== undefined && refValue.value.includes(',')) {
        const parts = refValue.value.split(',');
        return parts[value.part] || undefined;
      }
      
      return refValue.value;
    }
    
    return undefined;
  }
  
  /**
   * Evaluate a conditional type based on current values
   * @param type The conditional type to evaluate
   * @param inputIndex Index of the input containing this conditional type
   * @param values Current value map
   * @returns The resolved type based on the condition
   */
  public evaluateConditionalType(
    type: InputType,
    inputIndex: number,
    values: Map<number, { value: string }>
  ): InputType {
    // Only evaluate comparison types
    if (typeof type === 'object' && 'left' in type) {
      const conditionMet = this.compareValues(
        type.left,
        type.operator,
        type.right,
        values
      );
      
      // Use the appropriate type based on the condition
      const resolvedType = conditionMet ? type.trueType : type.falseType;
      
      // If the resolved type is also a conditional, evaluate it recursively
      if (typeof resolvedType === 'object' && 'left' in resolvedType) {
        return this.evaluateConditionalType(resolvedType, inputIndex, values);
      }
      
      return resolvedType;
    }
    
    // Non-conditional types are returned as is
    return type;
  }
  
  /**
   * Set a value for an input and handle dependencies
   * @param parsedSentence The parsed sentence
   * @param currentValues Current values map
   * @param index Input index to set
   * @param value New value
   * @returns Result with success and updated values, or error
   */
  public setValue(
    parsedSentence: ParsedSentence,
    currentValues: Map<number, InputState>,
    index: number,
    value: string
  ): SetValueResult {
    // Find the input definition
    const input = parsedSentence.inputs.find(i => i.index === index);
    if (!input) {
      return {
        success: false,
        value: currentValues,
        error: `Input with index ${index} not found`
      };
    }
    
    // Create a new values map with the updated value
    const newValues = new Map(currentValues);
    
    // If the input has a conditional type, evaluate it with the new value temporarily in place
    let typeToValidate = input.type;
    if (input.type && typeof input.type === 'object' && 'left' in input.type) {
      // Create a temporary map with the new value for condition evaluation
      const tempValues = new Map(newValues);
      tempValues.set(index, { value });
      
      // Evaluate the condition with the new value in place
      typeToValidate = this.evaluateConditionalType(input.type, index, tempValues);
      
      // If the type evaluates to null, clear the value
      if (typeToValidate === 'null') {
        newValues.delete(index);
        return {
          success: true,
          value: newValues
        };
      }
      
      // If the type is a constant, set that value instead
      if (typeof typeToValidate === 'object' && 'constant' in typeToValidate) {
        newValues.set(index, {
          value: typeToValidate.constant,
          isDisabled: true
        });
        
        // Skip further validation as the value is predetermined
        return {
          success: true,
          value: newValues,
          error: "Constant value set automatically"
        };
      }
    }
    
    // Validate the input value based on its evaluated type
    const validation = this.validateInput(value, typeToValidate);
    if (!validation.success) {
      return {
        success: false,
        value: currentValues,
        error: validation.error
      };
    }
    
    // Set the new value
    newValues.set(index, { value });
    
    // Clear values of directly dependent inputs
    parsedSentence.inputs.forEach(otherInput => {
      if (otherInput.dependentOn === index) {
        newValues.delete(otherInput.index);
      }
    });
    
    // Update conditionally dependent inputs
    parsedSentence.inputs.forEach(otherInput => {
      if (otherInput.index === index) return;
      
      // Check if this input has a conditional type affected by our change
      if (otherInput.type && typeof otherInput.type === 'object' && 'left' in otherInput.type) {
        const resolvedType = this.evaluateConditionalType(
          otherInput.type,
          otherInput.index,
          newValues
        );
        
        // If the type resolves to null, clear the value
        if (resolvedType === 'null') {
          newValues.delete(otherInput.index);
        } 
        // If the type is a constant, set that value
        else if (typeof resolvedType === 'object' && 'constant' in resolvedType) {
          newValues.set(otherInput.index, {
            value: resolvedType.constant,
            isDisabled: true
          });
        }
        // If the value was previously a constant but no longer is, remove the disable flag
        else {
          const existingValue = newValues.get(otherInput.index);
          if (existingValue?.isDisabled) {
            newValues.set(otherInput.index, {
              value: existingValue.value
            });
          }
        }
      }
    });
    
    return {
      success: true,
      value: newValues
    };
  }
}