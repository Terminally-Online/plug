/**
 * Result pattern for more consistent error handling
 * This approach avoids throwing exceptions and provides a more predictable API
 */

/**
 * Success result type
 */
export interface Success<T> {
  success: true;
  value: T;
}

/**
 * Error result type
 */
export interface Failure<E = string> {
  success: false;
  error: E;
}

/**
 * Union type representing either a success or failure result
 */
export type Result<T, E = string> = Success<T> | Failure<E>;

/**
 * Result utility functions
 */
export const Result = {
  /**
   * Create a success result
   * @param value The success value
   * @returns A success result object
   */
  success<T>(value: T): Success<T> {
    return { success: true, value };
  },

  /**
   * Create a failure result
   * @param error The error value
   * @returns A failure result object
   */
  failure<E = string>(error: E): Failure<E> {
    return { success: false, error };
  },

  /**
   * Wrap a function that might throw into a Result
   * @param fn The function to call
   * @returns A Result representing success or failure
   */
  try<T>(fn: () => T): Result<T> {
    try {
      return Result.success(fn());
    } catch (e) {
      return Result.failure(e instanceof Error ? e.message : String(e));
    }
  },

  /**
   * Wrap an async function that might reject into a Result
   * @param fn The async function to call
   * @returns A Promise that always resolves to a Result
   */
  async tryAsync<T>(fn: () => Promise<T>): Promise<Result<T>> {
    try {
      const value = await fn();
      return Result.success(value);
    } catch (e) {
      return Result.failure(e instanceof Error ? e.message : String(e));
    }
  },

  /**
   * Map a result to a new value if successful
   * @param result The input result
   * @param mapFn Function to transform the success value
   * @returns A new result with the mapped value or the original error
   */
  map<T, U, E = string>(result: Result<T, E>, mapFn: (value: T) => U): Result<U, E> {
    if (result.success) {
      return Result.success(mapFn(result.value));
    }
    return result;
  },

  /**
   * Apply a side effect to a result's value if successful
   * @param result The input result
   * @param fn Function to apply to the success value
   * @returns The original result
   */
  tap<T, E = string>(result: Result<T, E>, fn: (value: T) => void): Result<T, E> {
    if (result.success) {
      fn(result.value);
    }
    return result;
  },

  /**
   * Handle both success and error cases
   * @param result The input result
   * @param handlers Object with onSuccess and onFailure handlers
   * @returns The value returned by the appropriate handler
   */
  match<T, E = string, R = any>(
    result: Result<T, E>,
    handlers: {
      onSuccess: (value: T) => R;
      onFailure: (error: E) => R;
    }
  ): R {
    if (result.success) {
      return handlers.onSuccess(result.value);
    }
    return handlers.onFailure(result.error);
  },

  /**
   * Unwrap a result to get its value or throw an error
   * @param result The input result
   * @returns The success value
   * @throws The error value if the result is a failure
   */
  unwrap<T, E = string>(result: Result<T, E>): T {
    if (result.success) {
      return result.value;
    }
    throw typeof result.error === 'string'
      ? new Error(result.error)
      : result.error;
  },

  /**
   * Unwrap a result to get its value or return a default
   * @param result The input result
   * @param defaultValue Default value to return on failure
   * @returns The success value or the default
   */
  unwrapOr<T, E = string>(result: Result<T, E>, defaultValue: T): T {
    if (result.success) {
      return result.value;
    }
    return defaultValue;
  }
};