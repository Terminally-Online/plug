import { OpenAPIV3 } from 'openapi-types';
import { OpenAPIParser } from './openapi-parser';
import * as fs from 'fs';
import * as path from 'path';

/**
 * Generator for TypeScript API client from OpenAPI schema
 */
export class ApiGenerator {
  private parser: OpenAPIParser;
  
  /**
   * Create a new ApiGenerator instance
   * @param parser OpenAPIParser instance
   */
  constructor(parser: OpenAPIParser) {
    this.parser = parser;
  }

  /**
   * Generate TypeScript API client from OpenAPI schema
   * @returns Generated TypeScript code as string
   */
  public generateApiClient(): string {
    const operations = this.parser.getOperations();
    let apiCode = '// This file is auto-generated from the OpenAPI schema. Do not edit manually.\n\n';
    
    // Import dependencies
    apiCode += 'import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from \'axios\';\n';
    apiCode += 'import * as Types from \'../types/generated\';\n\n';
    
    // Define API class
    apiCode += `/**
 * Auto-generated API client for the Plug API
 */
export class PlugAPI {
  private apiKey: string;
  private client: AxiosInstance;
  private baseURL = 'https://api.onplug.io';

  /**
   * Create a new PlugAPI instance
   * @param apiKey API key for authentication
   * @param config Optional configuration overrides
   */
  constructor(apiKey: string, config?: { baseURL?: string }) {
    this.apiKey = apiKey;
    this.baseURL = config?.baseURL || this.baseURL;
    
    this.client = axios.create({
      baseURL: this.baseURL,
      headers: {
        'Authorization': \`Bearer \${this.apiKey}\`,
        'Content-Type': 'application/json',
      },
    });
  }

`;
    
    // Generate methods for each operation
    operations.forEach(op => {
      apiCode += this.generateMethod(op);
    });
    
    // Add utility methods and error handling
    apiCode += `
  /**
   * Handle API errors
   * @param error Error object from axios
   * @throws Rethrows error with additional context
   */
  private handleError(error: any): never {
    if (axios.isAxiosError(error)) {
      const status = error.response?.status;
      const message = error.response?.data?.message || error.message;
      
      throw new Error(\`API Error (\${status}): \${message}\`);
    }
    
    throw error;
  }
}`;
    
    return apiCode;
  }
  
  /**
   * Generate a TypeScript method for an API operation
   * @param operation Operation details
   * @returns Generated TypeScript method code
   */
  private generateMethod(operation: {
    path: string;
    method: string;
    operation: OpenAPIV3.OperationObject;
    pathItem: OpenAPIV3.PathItemObject;
  }): string {
    const { path, method, operation: op } = operation;
    
    // Generate method name based on operation ID or path
    let methodName = op.operationId || this.getMethodNameFromPath(path, method);
    methodName = this.camelCase(methodName);
    
    // Extract params from path
    const pathParams = this.extractPathParams(path);
    
    // Extract query params from operation
    const queryParams = this.extractQueryParams(op);
    
    // Extract request body from operation
    const requestBody = this.extractRequestBody(op);
    
    // Extract response types from operation
    const responseType = this.extractResponseType(op);
    
    // Generate method signature
    let methodCode = `  /**\n`;
    if (op.summary) {
      methodCode += `   * ${op.summary}\n`;
    }
    if (op.description) {
      methodCode += `   * ${op.description}\n`;
    }
    
    // Add param docs
    pathParams.forEach(param => {
      methodCode += `   * @param ${param.name} ${param.description || ''}\n`;
    });
    
    if (queryParams.length > 0) {
      methodCode += `   * @param queryParams Query parameters\n`;
    }
    
    if (requestBody) {
      methodCode += `   * @param requestBody Request body\n`;
    }
    
    methodCode += `   * @returns Promise resolving to the API response\n`;
    methodCode += `   */\n`;
    
    // Method signature
    methodCode += `  public async ${methodName}(`;
    
    // Method parameters
    const parameters: string[] = [];
    pathParams.forEach(param => {
      parameters.push(`${param.name}: ${param.type}`);
    });
    
    if (queryParams.length > 0) {
      parameters.push(`queryParams?: ${this.getQueryParamsType(queryParams)}`);
    }
    
    if (requestBody) {
      parameters.push(`requestBody${requestBody.required ? '' : '?'}: ${requestBody.type}`);
    }
    
    methodCode += parameters.join(', ');
    methodCode += `): Promise<${responseType}> {\n`;
    
    // Method implementation
    methodCode += `    try {\n`;
    
    // Create config and interpolate path params
    methodCode += `      const config: AxiosRequestConfig = {};\n`;
    
    if (queryParams.length > 0) {
      methodCode += `      if (queryParams) {\n`;
      methodCode += `        config.params = queryParams;\n`;
      methodCode += `      }\n`;
    }
    
    // Create endpoint with path parameters
    let endpoint = path;
    if (pathParams.length > 0) {
      methodCode += `      const endpoint = \`${this.pathToTemplate(path)}\`;\n`;
    } else {
      methodCode += `      const endpoint = '${path}';\n`;
    }
    
    // Make the request based on method type
    switch (method.toLowerCase()) {
      case 'get':
        methodCode += `      const response = await this.client.get<${responseType}>(endpoint, config);\n`;
        break;
      case 'post':
        methodCode += `      const response = await this.client.post<${responseType}>(endpoint, ${requestBody ? 'requestBody' : '{}'}, config);\n`;
        break;
      case 'put':
        methodCode += `      const response = await this.client.put<${responseType}>(endpoint, ${requestBody ? 'requestBody' : '{}'}, config);\n`;
        break;
      case 'delete':
        methodCode += `      const response = await this.client.delete<${responseType}>(endpoint, config);\n`;
        break;
      case 'patch':
        methodCode += `      const response = await this.client.patch<${responseType}>(endpoint, ${requestBody ? 'requestBody' : '{}'}, config);\n`;
        break;
      default:
        methodCode += `      const response = await this.client.request<${responseType}>({ method: '${method}', url: endpoint, data: ${requestBody ? 'requestBody' : 'undefined'}, ...config });\n`;
    }
    
    methodCode += `      return response.data;\n`;
    methodCode += `    } catch (error) {\n`;
    methodCode += `      this.handleError(error);\n`;
    methodCode += `    }\n`;
    methodCode += `  }\n\n`;
    
    return methodCode;
  }
  
  /**
   * Extract path parameters from a path string
   * @param path Path string
   * @returns Array of path parameters
   */
  private extractPathParams(path: string): Array<{
    name: string;
    type: string;
    description?: string;
  }> {
    const params: Array<{ name: string; type: string; description?: string }> = [];
    const paramRegex = /{([^}]+)}/g;
    let match;
    
    while ((match = paramRegex.exec(path)) !== null) {
      params.push({
        name: match[1],
        type: 'string',
        description: `Path parameter: ${match[1]}`,
      });
    }
    
    return params;
  }
  
  /**
   * Extract query parameters from an operation
   * @param operation OpenAPI operation object
   * @returns Array of query parameters
   */
  private extractQueryParams(operation: OpenAPIV3.OperationObject): Array<{
    name: string;
    type: string;
    required: boolean;
    description?: string;
  }> {
    const params = operation.parameters || [];
    const queryParams: Array<{
      name: string;
      type: string;
      required: boolean;
      description?: string;
    }> = [];
    
    params.forEach(paramRef => {
      const param = paramRef as OpenAPIV3.ParameterObject;
      
      if (param.in === 'query') {
        let type = 'any';
        
        if (param.schema) {
          const schema = param.schema as OpenAPIV3.SchemaObject;
          type = this.schemaToTypeName(schema);
        }
        
        queryParams.push({
          name: param.name,
          type,
          required: param.required || false,
          description: param.description,
        });
      }
    });
    
    return queryParams;
  }
  
  /**
   * Extract request body from an operation
   * @param operation OpenAPI operation object
   * @returns Request body info or null if none
   */
  private extractRequestBody(operation: OpenAPIV3.OperationObject): {
    type: string;
    required: boolean;
    description?: string;
  } | null {
    if (!operation.requestBody) {
      return null;
    }
    
    const requestBody = operation.requestBody as OpenAPIV3.RequestBodyObject;
    const content = requestBody.content || {};
    
    // Try to use application/json content type
    const jsonContent = content['application/json'];
    
    if (jsonContent && jsonContent.schema) {
      const schema = jsonContent.schema as OpenAPIV3.SchemaObject;
      return {
        type: this.schemaToTypeName(schema),
        required: requestBody.required || false,
        description: requestBody.description,
      };
    }
    
    // Fallback to any content type if application/json not found
    for (const contentType in content) {
      const mediaType = content[contentType];
      if (mediaType && mediaType.schema) {
        const schema = mediaType.schema as OpenAPIV3.SchemaObject;
        return {
          type: this.schemaToTypeName(schema),
          required: requestBody.required || false,
          description: requestBody.description,
        };
      }
    }
    
    return {
      type: 'any',
      required: requestBody.required || false,
      description: requestBody.description,
    };
  }
  
  /**
   * Extract response type from an operation
   * @param operation OpenAPI operation object
   * @returns TypeScript type name for the response
   */
  private extractResponseType(operation: OpenAPIV3.OperationObject): string {
    const responses = operation.responses || {};
    
    // Look for success responses (2xx)
    for (const status in responses) {
      if (status.startsWith('2')) {
        const response = responses[status] as OpenAPIV3.ResponseObject;
        const content = response.content || {};
        
        // Try application/json first
        const jsonContent = content['application/json'];
        
        if (jsonContent && jsonContent.schema) {
          const schema = jsonContent.schema as OpenAPIV3.SchemaObject;
          return this.schemaToTypeName(schema);
        }
        
        // Then try any other content type
        for (const contentType in content) {
          const mediaType = content[contentType];
          if (mediaType && mediaType.schema) {
            const schema = mediaType.schema as OpenAPIV3.SchemaObject;
            return this.schemaToTypeName(schema);
          }
        }
      }
    }
    
    return 'any';
  }
  
  /**
   * Convert schema to TypeScript type name
   * @param schema OpenAPI schema object
   * @returns TypeScript type name
   */
  private schemaToTypeName(schema: OpenAPIV3.SchemaObject): string {
    if ('$ref' in schema && typeof schema.$ref === 'string') {
      const refName = this.parser.getSchemaNameFromRef(schema.$ref);
      return `Types.${refName}`;
    }
    
    if (schema.type === 'array' && schema.items) {
      const itemsSchema = schema.items as OpenAPIV3.SchemaObject;
      const itemType = this.schemaToTypeName(itemsSchema);
      return `${itemType}[]`;
    }
    
    if (schema.type === 'object') {
      if (schema.additionalProperties) {
        const valueType = typeof schema.additionalProperties === 'boolean' 
          ? 'any'
          : this.schemaToTypeName(schema.additionalProperties as OpenAPIV3.SchemaObject);
        return `Record<string, ${valueType}>`;
      }
      
      return 'Record<string, any>';
    }
    
    switch (schema.type) {
      case 'string':
        return 'string';
      case 'number':
      case 'integer':
        return 'number';
      case 'boolean':
        return 'boolean';
      default:
        return 'any';
    }
  }
  
  /**
   * Get query parameters interface type
   * @param queryParams Array of query parameters
   * @returns TypeScript interface for query parameters
   */
  private getQueryParamsType(queryParams: Array<{
    name: string;
    type: string;
    required: boolean;
    description?: string;
  }>): string {
    if (queryParams.length === 0) {
      return 'Record<string, any>';
    }
    
    const props: string[] = [];
    queryParams.forEach(param => {
      props.push(`${param.name}${param.required ? '' : '?'}: ${param.type}`);
    });
    
    return `{ ${props.join('; ')} }`;
  }
  
  /**
   * Convert path to template literal for interpolation
   * @param path Path string
   * @returns Template string
   */
  private pathToTemplate(path: string): string {
    return path.replace(/{([^}]+)}/g, '${$1}');
  }
  
  /**
   * Generate a method name from path and HTTP method
   * @param path Path string
   * @param httpMethod HTTP method
   * @returns Method name
   */
  private getMethodNameFromPath(path: string, httpMethod: string): string {
    let name = httpMethod.toLowerCase();
    
    const cleanPath = path
      .replace(/^\/|\/$/g, '') // Remove leading/trailing slashes
      .replace(/{([^}]+)}/g, '$1'); // Replace {param} with param
    
    const parts = cleanPath.split('/');
    parts.forEach((part, index) => {
      if (index === 0 && name.length > 0) {
        name += this.capitalize(part);
      } else {
        name += this.capitalize(part);
      }
    });
    
    return name;
  }
  
  /**
   * Convert a string to camelCase
   * @param str Input string
   * @returns camelCase string
   */
  private camelCase(str: string): string {
    return str
      .replace(/[^a-zA-Z0-9]+(.)/g, (_, char) => char.toUpperCase())
      .replace(/^[A-Z]/, c => c.toLowerCase());
  }
  
  /**
   * Capitalize a string
   * @param str Input string
   * @returns Capitalized string
   */
  private capitalize(str: string): string {
    if (!str) return '';
    return str.charAt(0).toUpperCase() + str.slice(1);
  }
  
  /**
   * Save generated API client to a file
   * @param outputPath Path to save the generated file
   */
  public saveToFile(outputPath: string): void {
    const code = this.generateApiClient();
    const dir = path.dirname(outputPath);
    
    // Create directory if it doesn't exist
    if (!fs.existsSync(dir)) {
      fs.mkdirSync(dir, { recursive: true });
    }
    
    fs.writeFileSync(outputPath, code, 'utf8');
  }
}