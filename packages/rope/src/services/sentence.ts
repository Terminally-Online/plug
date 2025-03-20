/**
 * Input definition extracted from a sentence
 */
export interface SentenceInput {
  index: number;
  name: string;
  type?: string;
  defaultValue?: string;
  required?: boolean;
  dependentOn?: number;
  options?: string[];
  description?: string;
}

/**
 * Parsed sentence result
 */
export interface ParsedSentence {
  original: string;
  parts: string[];
  inputs: SentenceInput[];
  template?: string; // For compatibility with Cord
}

/**
 * Validation result interface
 */
export interface ValidationResult {
  success: boolean;
  error?: string;
}

/**
 * Service for parsing and working with action sentences
 */
export class SentenceService {
  // Regular expression to extract input placeholders from sentences
  private static readonly INPUT_REGEX = /\{(\d+)(?:=>(\d+))?\}/g;
  
  // Cache for parsed sentences to avoid repeated parsing of the same sentence
  private static readonly parseCache = new Map<string, ParsedSentence>();
  
  /**
   * Parse a sentence into parts and extract input definitions
   * @param sentence The sentence to parse (e.g. "Swap {0} for {1} on {2}")
   * @returns The parsed sentence with extracted inputs
   */
  public parseSentence(sentence: string): ParsedSentence {
    // Return empty result for empty sentence
    if (!sentence) {
      return { original: '', parts: [], inputs: [], template: '' };
    }
    
    // Check if this sentence is already in the cache
    const cachedResult = SentenceService.parseCache.get(sentence);
    if (cachedResult) {
      return cachedResult;
    }
    
    // Split the sentence into parts by input placeholders
    const parts: string[] = [];
    const inputs: SentenceInput[] = [];
    
    // Use regex to find all inputs and build parts array
    let lastIndex = 0;
    let match: RegExpExecArray | null;
    
    // Reset the regex to start from the beginning
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
    
    // Create result object
    const result: ParsedSentence = {
      original: sentence,
      template: sentence, // For compatibility with Cord
      parts,
      inputs: inputs.sort((a, b) => a.index - b.index)
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
    return parsed.parts.map(part => {
      const match = part.match(SentenceService.INPUT_REGEX);
      if (!match) {
        return part;
      }
      
      const inputIndex = match[1] ? parseInt(match[1]) : parseInt(match[0].replace(/[{}]/g, ''));
      const mappedIndex = match[2] ? parseInt(match[2]) : inputIndex;
      
      return values[mappedIndex] !== undefined ? values[mappedIndex] : `[${mappedIndex}]`;
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
      if (part.match(SentenceService.INPUT_REGEX)) {
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
  public validateInput(value: string, type: string = 'string'): ValidationResult {
    // Empty check (except for boolean which can be empty)
    if ((!value || value.trim() === '') && type !== 'boolean') {
      return { success: false, error: 'Value is required' };
    }
    
    switch (type.toLowerCase()) {
      case 'number':
      case 'integer':
      case 'float':
        if (isNaN(Number(value))) {
          return { success: false, error: 'Value must be a number' };
        }
        if (type === 'integer' && !Number.isInteger(Number(value))) {
          return { success: false, error: 'Value must be an integer' };
        }
        break;
        
      case 'address':
        // Simple Ethereum address validation
        if (!/^0x[a-fA-F0-9]{40}$/.test(value)) {
          return { success: false, error: 'Invalid address format' };
        }
        break;
        
      case 'boolean':
        if (value !== 'true' && value !== 'false' && value !== '' && value !== undefined) {
          return { success: false, error: 'Value must be true or false' };
        }
        break;
    }
    
    return { success: true };
  }
  
  /**
   * Get placeholder text for an input type
   * @param type The input type
   * @returns A suitable placeholder text
   */
  public getPlaceholder(type: string = 'string'): string {
    switch (type.toLowerCase()) {
      case 'number':
      case 'integer':
      case 'float':
        return 'Enter a number';
      case 'address':
        return 'Enter an address';
      case 'token':
        return 'Select a token';
      case 'boolean':
        return 'Yes/No';
      default:
        return 'Enter a value';
    }
  }
  
  /**
   * Determine if an input should be rendered based on dependencies
   * @param inputType The type of the input
   * @param allInputs All inputs in the sentence
   * @param getValueFn Function to get value for an input index
   * @returns Whether the input should be rendered
   */
  public shouldRenderInput(
    inputType: string,
    allInputs: SentenceInput[],
    getValueFn: (index: number) => { value: string } | undefined
  ): boolean {
    // Find the input with the given type
    const input = allInputs.find(input => input.type === inputType);
    
    // If input has a dependency, check if the dependency has a value
    if (input?.dependentOn !== undefined) {
      const dependencyValue = getValueFn(input.dependentOn);
      if (!dependencyValue || !dependencyValue.value) {
        return false;
      }
    }
    
    return true;
  }
  
  /**
   * Set a value for an input and handle dependencies
   * @param parsedSentence The parsed sentence
   * @param currentValues Current values map
   * @param index Input index to set
   * @param value New value
   * @returns Result with success, updated values, and optional error
   */
  public setValue(
    parsedSentence: ParsedSentence,
    currentValues: Map<number, { value: string }>,
    index: number,
    value: string
  ): { success: boolean; values: Map<number, { value: string }>; error?: string } {
    // Find the input definition
    const input = parsedSentence.inputs.find(i => i.index === index);
    if (!input) {
      return { 
        success: false, 
        error: `Input with index ${index} not found`, 
        values: currentValues 
      };
    }
    
    // Validate the input value based on its type
    const validation = this.validateInput(value, input.type);
    if (!validation.success) {
      return { 
        success: false, 
        error: validation.error, 
        values: currentValues 
      };
    }
    
    // Create a new values map with the updated value
    const newValues = new Map(currentValues);
    newValues.set(index, { value });
    
    // Clear values of dependent inputs
    parsedSentence.inputs.forEach(otherInput => {
      if (otherInput.dependentOn === index) {
        newValues.delete(otherInput.index);
      }
    });
    
    return { success: true, values: newValues };
  }
}
