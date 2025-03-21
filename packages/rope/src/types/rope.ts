/**
 * Type definition for Rope API responses
 */
export type Rope = Array<{
  id: string;
  name?: string;
  type?: string;
  data?: any;
  metadata?: Record<string, any>;
  createdAt?: string;
  updatedAt?: string;
}>;

/**
 * Interface for API request options
 */
export interface RequestOptions {
  headers?: Record<string, string>;
  timeout?: number;
  retries?: number;
}

/**
 * Interface for API error responses
 */
export interface ApiError {
  status: number;
  code: string;
  message: string;
  details?: any;
}

/**
 * EVM Types
 */
type UintSizes =
  | 8 | 16 | 24 | 32 | 40 | 48 | 56 | 64 | 72 | 80 | 88 | 96
  | 104 | 112 | 120 | 128 | 136 | 144 | 152 | 160 | 168 | 176
  | 184 | 192 | 200 | 208 | 216 | 224 | 232 | 240 | 248 | 256;

type UintType = `uint${UintSizes}`;
type IntType = `int${UintSizes}`;
type BytesType = `bytes${
  | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16
  | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 | 26 | 27 | 28 | 29 | 30 | 31 | 32}`;

/**
 * Ethereum Virtual Machine types for validator
 */
export type EvmType =
  | UintType
  | IntType
  | BytesType
  | "address"
  | "bool"
  | "string"
  | "bytes"
  | "null"
  | "float";

/**
 * Enhanced Type System for Rope
 */

/**
 * Constant type - used for fixed values
 */
export type ConstantType = { constant: string };

/**
 * Simple union type to avoid circular references
 */
export type SimpleUnionType = {
  types: (EvmType | ConstantType)[];
};

/**
 * Compound type - combines a base type with metadata
 */
export type CompoundType = {
  baseType: EvmType | ConstantType | SimpleUnionType;
  metadata: (EvmType | ConstantType | SimpleUnionType)[];
};

/**
 * Comparison operator for conditional types
 */
export type ComparisonOperator = "==" | ">" | "<" | ">=" | "<=" | "!=";

/**
 * Value to compare - can be a string or a reference to another input
 */
export type ComparisonValue = string | { reference: number; part?: number };

/**
 * Conditional type based on comparing values
 */
export type ComparisonType = {
  left: ComparisonValue;
  operator: ComparisonOperator;
  right: ComparisonValue;
  trueType: InputType;
  falseType: InputType;
};

/**
 * Union type - one of multiple possible types
 */
export type UnionType = {
  types: InputType[];
};

/**
 * All possible input types
 */
export type InputType = 
  | EvmType 
  | ConstantType 
  | CompoundType 
  | ComparisonType 
  | UnionType;

/**
 * Input error types
 */
export type InputError = { 
  type: "validation" | "resolution"; 
  message: string 
};

/**
 * Input state
 */
export interface InputState {
  value: string;
  error?: InputError;
  isDisabled?: boolean;
}

/**
 * Map of input values
 */
export type InputValues = Map<number, InputState>;

/**
 * Result from setting a value
 */
export type SetValueResult = {
  success: true;
  value: InputValues;
  error?: string;
};
