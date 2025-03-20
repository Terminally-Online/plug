// This file is auto-generated from the OpenAPI schema. Do not edit manually.

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import * as Types from '../types/generated';
import { SchemaCache } from '../services/cache';
import { ApiError, RopeError } from '../services/errors';

/**
 * Auto-generated API client for the Plug API
 */
export class PlugAPI {
  private apiKey: string;
  private client: AxiosInstance;
  private baseURL = 'https://api.onplug.io';
  private cache: SchemaCache;

  /**
   * Create a new PlugAPI instance
   * @param apiKey API key for authentication
   * @param config Optional configuration overrides
   */
  constructor(apiKey: string, config?: { baseURL?: string; cacheTTL?: number }) {
    this.apiKey = apiKey;
    this.baseURL = config?.baseURL || this.baseURL;
    this.cache = new SchemaCache(config?.cacheTTL);
    
    this.client = axios.create({
      baseURL: this.baseURL,
      headers: {
        'Authorization': `Bearer ${this.apiKey}`,
        'Content-Type': 'application/json',
      },
    });
  }

  /**
   * Create API Key
   * Creates a new API key with the specified parameters. Requires admin privileges.
   * @param requestBody Request body
   * @returns Promise resolving to the API response
   */
  public async postApiKey(requestBody?: Types.ApiKeyApiKeyCreateRequest): Promise<Types.ModelsApiKey> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = '/api-key';
      const response = await this.client.post<Types.ModelsApiKey>(endpoint, requestBody, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Get API Key
   * Retrieves an API key by ID. Requires admin privileges.
   * @param id Path parameter: id
   * @returns Promise resolving to the API response
   */
  public async getApiKeyId(id: string): Promise<Types.ModelsApiKey> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = `/api-key/${id}`;
      const response = await this.client.get<Types.ModelsApiKey>(endpoint, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Update API Key
   * Updates an existing API key with the provided parameters. Requires admin privileges.
   * @param id Path parameter: id
   * @param requestBody Request body
   * @returns Promise resolving to the API response
   */
  public async postApiKeyId(id: string, requestBody?: Types.ApiKeyApiKeyCreateRequest): Promise<any> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = `/api-key/${id}`;
      const response = await this.client.post<any>(endpoint, requestBody, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Delete API Key
   * Deletes an API key by ID. Requires admin privileges.
   * @param id Path parameter: id
   * @returns Promise resolving to the API response
   */
  public async deleteApiKeyId(id: string): Promise<Types.ModelsApiKey> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = `/api-key/${id}`;
      const response = await this.client.delete<Types.ModelsApiKey>(endpoint, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Health Check
   * Returns the health status of the API
   * @returns Promise resolving to the API response
   */
  public async getHealth(): Promise<Types.HealthHealthResponse> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = '/health';
      const response = await this.client.get<Types.HealthHealthResponse>(endpoint, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Get Schema
   * Retrieves available action schemas for protocols based on query parameters. Requires API key.
   * Uses caching to minimize API calls.
   * @param queryParams Query parameters
   * @returns Promise resolving to the API response
   */
  public async getSolver(queryParams?: Types.SchemaQueryParams): Promise<Record<string, Types.ActionsProtocolSchema>> {
    try {
      // Check if we have a cached response
      const cachedData = this.cache.get(queryParams);
      if (cachedData) {
        return cachedData;
      }

      // No cache hit, make the API request
      const config: AxiosRequestConfig = {};
      if (queryParams) {
        config.params = queryParams;
      }
      
      const endpoint = '/solver';
      const response = await this.client.get<Record<string, Types.ActionsProtocolSchema>>(endpoint, config);
      
      // Store the response in cache for future use
      this.cache.set(response.data, queryParams);
      
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }
  
  /**
   * Get chain-specific schemas
   * Retrieves schemas filtered by chain ID
   * @param chainId Chain ID to filter by
   * @param additionalParams Additional query parameters
   * @returns Promise resolving to filtered schemas
   */
  public async getSchemasByChain(
    chainId: number, 
    additionalParams?: Omit<Types.SchemaQueryParams, 'chainId'>
  ): Promise<Record<string, Types.ActionsProtocolSchema>> {
    const params: Types.SchemaQueryParams = {
      ...additionalParams,
      chainId
    };
    
    return this.getSolver(params);
  }

  /**
   * Solve Intent
   * Processes an intent and returns a solution (transaction data). Requires API key.
   * @param requestBody Request body
   * @returns Promise resolving to the API response
   */
  public async postSolver(requestBody?: Types.ModelsIntent): Promise<Record<string, any>> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = '/solver';
      const response = await this.client.post<Record<string, any>>(endpoint, requestBody, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Get Kill Switch Status
   * Returns the current status of the kill switch, which controls whether new solver operations are allowed
   * @returns Promise resolving to the API response
   */
  public async getSolverKill(): Promise<Types.KillKillResponse> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = '/solver/kill';
      const response = await this.client.get<Types.KillKillResponse>(endpoint, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Toggle Kill Switch
   * Toggles the state of the kill switch that controls whether new solver operations are allowed
   * @returns Promise resolving to the API response
   */
  public async postSolverKill(): Promise<Types.KillKillResponse> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = '/solver/kill';
      const response = await this.client.post<Types.KillKillResponse>(endpoint, {}, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Create Intent
   * Creates a new intent with the provided parameters
   * @param requestBody Request body
   * @returns Promise resolving to the API response
   */
  public async postSolverSave(requestBody?: Types.ModelsIntent): Promise<Types.ModelsIntent> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = '/solver/save';
      const response = await this.client.post<Types.ModelsIntent>(endpoint, requestBody, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Get Intent
   * Retrieves an intent or list of intents by ID or address
   * @param id Path parameter: id
   * @returns Promise resolving to the API response
   */
  public async getSolverSaveId(id: string): Promise<Types.SaveIntentListResponse> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = `/solver/save/${id}`;
      const response = await this.client.get<Types.SaveIntentListResponse>(endpoint, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Toggle Intent Saved
   * Toggles the saved status of an intent
   * @param id Path parameter: id
   * @returns Promise resolving to the API response
   */
  public async postSolverSaveId(id: string): Promise<Types.ModelsIntent> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = `/solver/save/${id}`;
      const response = await this.client.post<Types.ModelsIntent>(endpoint, {}, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Delete Intent
   * Deletes an intent by ID
   * @param id Path parameter: id
   * @returns Promise resolving to the API response
   */
  public async deleteSolverSaveId(id: string): Promise<any> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = `/solver/save/${id}`;
      const response = await this.client.delete<any>(endpoint, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }

  /**
   * Toggle Intent Status
   * Toggles the status of an intent between 'active' and 'paused'
   * @param id Path parameter: id
   * @returns Promise resolving to the API response
   */
  public async postSolverSaveIdStatus(id: string): Promise<Types.ModelsIntent> {
    try {
      const config: AxiosRequestConfig = {};
      const endpoint = `/solver/save/${id}/status`;
      const response = await this.client.post<Types.ModelsIntent>(endpoint, {}, config);
      return response.data;
    } catch (error) {
      this.handleError(error);
    }
  }


  /**
   * Handle API errors
   * @param error Error object from axios
   * @throws Rethrows error with additional context as ApiError
   */
  private handleError(error: any): never {
    if (axios.isAxiosError(error)) {
      const status = error.response?.status;
      const apiMessage = error.response?.data?.message || error.message;
      const message = `API Error (${status}): ${apiMessage}`;
      
      throw new ApiError(message, status, apiMessage);
    }
    
    // For non-Axios errors, wrap in a generic RopeError
    throw new RopeError(error.message || 'Unknown API error');
  }
}