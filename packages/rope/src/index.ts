// Export main classes
export { Rope } from './models/rope';
export { Knot } from './models/knot';
export { PlugAPI } from './models/api';

// Export services
export { CoilService } from './services/coil';
export { SchemaCache } from './services/cache';
export { SentenceService } from './services/sentence';

// Export error types
export {
  RopeError,
  ApiError,
  NotInitializedError,
  SchemaValidationError,
  NotFoundError,
  ChainUnsupportedError,
  CoilCompatibilityError,
  IntentError
} from './services/errors';

// Export types
export type { Rope as RopeType } from './types/rope';
export type { CoilReference, CoilValidationResult } from './services/coil';
export type { SentenceInput, ParsedSentence, ValidationResult } from './services/sentence';
export * as PlugTypes from './types/generated';

// Export utility for generating API from OpenAPI schema
export { generateApi } from './lib/generate-api';

// Export React hooks and Rope types
export { useRope, useKnot, RopeConfig } from './hooks/useRope';
export { 
  RopeState, 
  ProcessedKnot as RopeKnot,
  KnotData, 
  KnotId,
  RopeStateManager 
} from './models/rope-state-manager';

// --------------------------------
// Cord Compatibility Layer
// --------------------------------

import { SentenceService, ParsedSentence, SentenceInput } from './services/sentence';
import { CoilService } from './services/coil';

// Create singleton instances for the compatibility layer
const _sentenceService = new SentenceService();
const _coilService = new CoilService();

/**
 * Parse a sentence string into parts and inputs (Cord-compatible)
 * @param sentence The sentence to parse
 * @returns Result object with success flag, parsed value or error
 */
export function parseCordSentence(sentence: string): { 
  success: boolean; 
  value?: ParsedSentence; 
  error?: string 
} {
  try {
    const result = _sentenceService.parseSentence(sentence);
    // Add template parts field for compatibility with Cord's expected structure
    result.template = sentence;
    return { success: true, value: result };
  } catch (error) {
    return { 
      success: false, 
      error: error instanceof Error ? error.message : 'Unknown error parsing sentence' 
    };
  }
}

/**
 * Maps from numeric indices to input values (Cord-compatible type)
 */
export type InputValues = Map<number, { value: string }>;

/**
 * Format a sentence by inserting values (Cord-compatible)
 * @param parsed The parsed sentence
 * @param values Map of input indices to values
 * @returns Result with resolved sentence or error
 */
export function resolveSentence(
  parsed: ParsedSentence,
  values: InputValues
): { success: boolean; value?: string; error?: string } {
  try {
    // Convert Map to plain object for formatSentence
    const valuesObj = Object.fromEntries(
      Array.from(values.entries()).map(([k, v]) => [k, v.value])
    );
    
    const result = _sentenceService.formatSentence(parsed, valuesObj);
    return { success: true, value: result };
  } catch (error) {
    return { 
      success: false, 
      error: error instanceof Error ? error.message : 'Unknown error resolving sentence' 
    };
  }
}

/**
 * Get a placeholder text for an input type (Cord-compatible)
 * @param type Input type
 * @returns Appropriate placeholder text
 */
export function getInputPlaceholder(type?: string): string {
  return _sentenceService.getPlaceholder(type);
}

/**
 * Determine if an input should be rendered based on dependencies (Cord-compatible)
 * @param inputType Input type
 * @param allInputs All inputs in the sentence
 * @param getValueFn Function to get values for inputs
 * @returns Whether the input should be rendered
 */
export function shouldRenderInput(
  inputType: string,
  allInputs: SentenceInput[],
  getValueFn: (index: number) => { value: string } | undefined
): boolean {
  return _sentenceService.shouldRenderInput(inputType, allInputs, getValueFn);
}

/**
 * Update a value in the input values map (Cord-compatible)
 * @param params Parameters for setting a value
 * @returns Result with updated values map or error
 */
export function setValue(params: {
  parsedSentence: ParsedSentence;
  currentValues: InputValues;
  index: number;
  value: string;
}): { success: boolean; value: InputValues; error?: string } {
  const result = _sentenceService.setValue(
    params.parsedSentence,
    params.currentValues,
    params.index,
    params.value
  );
  
  // Convert keys to match exported function's return type
  return {
    success: result.success,
    value: result.values,
    error: result.error
  };
}

/**
 * Create an empty state for input values (Cord-compatible)
 * @returns Empty Map for input values
 */
export function createInitialState(): InputValues {
  return new Map();
}

/**
 * Check if a string is a valid coil reference (Cord-compatible)
 * @param value String to check
 * @returns Whether it's a coil reference
 */
export function isCoilReference(value: string): boolean {
  return _coilService.isCoilReference(value);
}

/**
 * Parse a coil reference string (Cord-compatible)
 * @param value Coil reference string
 * @returns Parsed coil reference or null
 */
export function parseCoilReference(value: string) {
  return _coilService.parseCoilReference(value);
}

/**
 * Format a coil reference (Cord-compatible)
 * @param name Coil name
 * @param knotIndex Optional knot index
 * @returns Formatted coil reference string
 */
export function formatCoilReference(name: string, knotIndex?: number): string {
  return _coilService.formatCoilReference(name, knotIndex);
}

/**
 * Get template parts for UI rendering, compatible with Cord's part parsing
 * @param parsed Parsed sentence
 * @returns Array of string fragments for rendering
 */
export function getTemplateParts(parsed: ParsedSentence): string[] {
  if (!parsed) return [];
  return _sentenceService.getTemplateParts(parsed);
}

/**
 * Resolve coil references in a values object (Cord-compatible)
 * @param values Object containing input values with coil references
 * @param availableCoils Map of available coil names to their values
 * @returns New values object with resolved coil references
 */
export function resolveCoilReferences(
  values: Record<string, any>,
  availableCoils: Record<string, any>
): Record<string, any> {
  return _coilService.resolveCoilReferences(values, availableCoils);
}

/**
 * Resolve coil references in a sentence (Cord-compatible)
 * @param sentence The sentence or template containing possible coil references
 * @param values Input values with possible coil references
 * @param availableCoils Available coil values
 * @returns New values with resolved coil references
 */
export function resolveCoilReferencesInSentence(
  sentence: string,
  values: Record<string, string>,
  availableCoils: Record<string, string>
): Record<string, string> {
  return _coilService.resolveCoilReferencesInSentence(sentence, values, availableCoils);
}

/**
 * Validate coil references (Cord-compatible)
 * @param coilReferences Array of coil reference strings to validate
 * @param requiredTypes Map of coil names to their required types
 * @param availableCoils Map of available coil names to their types
 * @returns Validation result with errors if any
 */
export function validateCoilReferences(
  coilReferences: string[],
  requiredTypes: Record<string, string>,
  availableCoils: Record<string, string>
): { valid: boolean; errors: Record<string, string> } {
  return _coilService.validateCoilReferences(coilReferences, requiredTypes, availableCoils);
}

/**
 * Type aliases for Cord compatibility
 */
export type InputReference = SentenceInput;
export type CordState = {
  values: InputValues;
  parsed: ParsedSentence | null;
  resolvedSentence: string | null;
  error: { type: string; message: string } | null;
  validationErrors: Map<number, { type: string; message: string }>;
};