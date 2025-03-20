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
