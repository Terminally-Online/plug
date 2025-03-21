/**
 * Base error class for Rope package
 */
export class RopeError extends Error {
  constructor(message: string) {
    super(message);
    this.name = 'RopeError';
  }
}

/**
 * Error thrown when Rope is not initialized
 */
export class NotInitializedError extends RopeError {
  constructor(message: string = 'Rope must be initialized before performing this operation. Call initialize() first.') {
    super(message);
    this.name = 'NotInitializedError';
  }
}

/**
 * Error thrown when an API request fails
 */
export class ApiError extends RopeError {
  public status?: number;
  public apiMessage?: string;
  
  constructor(message: string, status?: number, apiMessage?: string) {
    super(message);
    this.name = 'ApiError';
    this.status = status;
    this.apiMessage = apiMessage;
  }
}

/**
 * Error thrown when schema validation fails
 */
export class SchemaValidationError extends RopeError {
  public field?: string;
  
  constructor(message: string, field?: string) {
    super(message);
    this.name = 'SchemaValidationError';
    this.field = field;
  }
}

/**
 * Error thrown when a protocol or action is not found
 */
export class NotFoundError extends RopeError {
  constructor(message: string) {
    super(message);
    this.name = 'NotFoundError';
  }
}

/**
 * Error thrown when there's an issue with coil compatibility
 */
export class CoilCompatibilityError extends RopeError {
  public coilName?: string;
  public expectedType?: string;
  public actualType?: string;
  
  constructor(message: string, coilName?: string, expectedType?: string, actualType?: string) {
    super(message);
    this.name = 'CoilCompatibilityError';
    this.coilName = coilName;
    this.expectedType = expectedType;
    this.actualType = actualType;
  }
}

/**
 * Error thrown when there's an issue with intent building or execution
 */
export class IntentError extends RopeError {
  constructor(message: string) {
    super(message);
    this.name = 'IntentError';
  }
}

/**
 * Error thrown when a chain ID is not supported
 */
export class ChainUnsupportedError extends RopeError {
  public chainId?: number;
  public protocol?: string;
  
  constructor(message: string, chainId?: number, protocol?: string) {
    super(message);
    this.name = 'ChainUnsupportedError';
    this.chainId = chainId;
    this.protocol = protocol;
  }
}