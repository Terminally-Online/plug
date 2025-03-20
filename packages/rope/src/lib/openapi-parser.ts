import { OpenAPIV3 } from 'openapi-types';
import * as fs from 'fs';
import * as path from 'path';

/**
 * Utility class for parsing and working with OpenAPI schemas
 */
export class OpenAPIParser {
  private schema: OpenAPIV3.Document;

  /**
   * Create a new OpenAPI parser instance
   * @param schemaPath Path to the OpenAPI schema file
   */
  constructor(schemaPath: string) {
    const schemaContent = fs.readFileSync(schemaPath, 'utf8');
    this.schema = JSON.parse(schemaContent) as OpenAPIV3.Document;
  }

  /**
   * Get the base schema object
   * @returns The parsed OpenAPI schema
   */
  public getSchema(): OpenAPIV3.Document {
    return this.schema;
  }

  /**
   * Get all paths defined in the schema
   * @returns Object containing all paths
   */
  public getPaths(): OpenAPIV3.PathsObject {
    return this.schema.paths || {};
  }

  /**
   * Get all components defined in the schema
   * @returns Object containing all components
   */
  public getComponents(): OpenAPIV3.ComponentsObject {
    return this.schema.components || {};
  }

  /**
   * Get all schema components defined in the schema
   * @returns Object containing all schema components
   */
  public getSchemas(): { [key: string]: OpenAPIV3.SchemaObject } {
    // Cast to the required type - we'll validate references when using the schemas
    return (this.schema.components?.schemas || {}) as { [key: string]: OpenAPIV3.SchemaObject };
  }

  /**
   * Get all path operations (endpoints)
   * @returns Array of path operations with metadata
   */
  public getOperations(): Array<{
    path: string;
    method: string;
    operation: OpenAPIV3.OperationObject;
    pathItem: OpenAPIV3.PathItemObject;
  }> {
    const paths = this.getPaths();
    const operations: Array<{
      path: string;
      method: string;
      operation: OpenAPIV3.OperationObject;
      pathItem: OpenAPIV3.PathItemObject;
    }> = [];

    Object.keys(paths).forEach((path) => {
      const pathItem = paths[path] as OpenAPIV3.PathItemObject;
      
      // Process each HTTP method in the path
      ['get', 'post', 'put', 'delete', 'patch'].forEach((method) => {
        const operation = pathItem[method as keyof OpenAPIV3.PathItemObject] as OpenAPIV3.OperationObject;
        
        if (operation) {
          operations.push({
            path,
            method,
            operation,
            pathItem,
          });
        }
      });
    });

    return operations;
  }

  /**
   * Check if the given reference is a schema reference
   * @param ref Reference string
   * @returns Whether the reference is a schema reference
   */
  public isSchemaRef(ref: string): boolean {
    return ref.startsWith('#/components/schemas/');
  }

  /**
   * Get schema name from a reference
   * @param ref Reference string
   * @returns Schema name
   */
  public getSchemaNameFromRef(ref: string): string {
    if (this.isSchemaRef(ref)) {
      return ref.replace('#/components/schemas/', '');
    }
    return '';
  }
}