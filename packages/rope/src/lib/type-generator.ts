import { OpenAPIV3 } from 'openapi-types';
import { OpenAPIParser } from './openapi-parser';
import * as fs from 'fs';
import * as path from 'path';

/**
 * Generator for TypeScript interfaces and types from OpenAPI schema
 */
export class TypeGenerator {
  private parser: OpenAPIParser;
  
  /**
   * Create a new TypeGenerator instance
   * @param parser OpenAPIParser instance
   */
  constructor(parser: OpenAPIParser) {
    this.parser = parser;
  }

  /**
   * Generate TypeScript interfaces from schema components
   * @returns Generated TypeScript code as string
   */
  public generateInterfaces(): string {
    const schemas = this.parser.getSchemas();
    let typeCode = '// This file is auto-generated from the OpenAPI schema. Do not edit manually.\n\n';
    
    // Define references to schemas we've generated
    const generatedTypes = new Set<string>();
    
    // Process each schema in the components
    Object.entries(schemas).forEach(([schemaName, schema]) => {
      if (generatedTypes.has(schemaName)) {
        return;
      }
      
      typeCode += this.generateType(schemaName, schema as OpenAPIV3.SchemaObject, generatedTypes);
      generatedTypes.add(schemaName);
    });
    
    return typeCode;
  }
  
  /**
   * Generate a TypeScript type or interface for a schema
   * @param name Name of the schema
   * @param schema OpenAPI schema object
   * @param generatedTypes Set of already generated types
   * @returns Generated TypeScript code
   */
  private generateType(
    name: string, 
    schema: OpenAPIV3.SchemaObject, 
    generatedTypes: Set<string>
  ): string {
    // Process based on schema type
    if (schema.type === 'object' || schema.properties) {
      return this.generateObjectInterface(name, schema, generatedTypes);
    } else if (schema.type === 'array' && schema.items) {
      return this.generateArrayType(name, schema, generatedTypes);
    } else if (schema.enum) {
      return this.generateEnumType(name, schema);
    } else if (schema.oneOf || schema.anyOf || schema.allOf) {
      return this.generateUnionType(name, schema, generatedTypes);
    } else {
      // For primitives or unknown types
      const tsType = this.mapOpenAPITypeToTS(schema);
      return `export type ${name} = ${tsType};\n\n`;
    }
  }
  
  /**
   * Generate a TypeScript interface for an object schema
   * @param name Name of the interface
   * @param schema OpenAPI schema object
   * @param generatedTypes Set of already generated types
   * @returns Generated TypeScript interface code
   */
  private generateObjectInterface(
    name: string, 
    schema: OpenAPIV3.SchemaObject, 
    generatedTypes: Set<string>
  ): string {
    let interfaceCode = `export interface ${name} {\n`;
    
    // Process properties
    const properties = schema.properties || {};
    Object.entries(properties).forEach(([propName, propSchema]) => {
      const isRequired = schema.required?.includes(propName) || false;
      const propertySchema = propSchema as OpenAPIV3.SchemaObject;
      
      // Add JSDoc for property if description is available
      if (propertySchema.description) {
        interfaceCode += `  /** ${propertySchema.description} */\n`;
      }
      
      // Generate property with type
      interfaceCode += `  ${propName}${isRequired ? '' : '?'}: ${this.getTypeForProperty(propertySchema, generatedTypes)};\n`;
    });
    
    // Add additionalProperties support if specified
    if (schema.additionalProperties) {
      const additionalPropsSchema = 
        typeof schema.additionalProperties === 'boolean' 
          ? { type: 'any' } 
          : schema.additionalProperties;
        
      interfaceCode += `  [key: string]: ${this.getTypeForProperty(additionalPropsSchema as OpenAPIV3.SchemaObject, generatedTypes)};\n`;
    }
    
    interfaceCode += `}\n\n`;
    return interfaceCode;
  }
  
  /**
   * Generate a TypeScript type for an array schema
   * @param name Name of the type
   * @param schema OpenAPI schema object
   * @param generatedTypes Set of already generated types
   * @returns Generated TypeScript type code
   */
  private generateArrayType(
    name: string, 
    schema: OpenAPIV3.SchemaObject, 
    generatedTypes: Set<string>
  ): string {
    // Use a safe approach when accessing items property
    let itemType = 'any';
    
    // Check that we have a proper array schema with items
    if (schema.type === 'array' && 'items' in schema && schema.items) {
      // We know it's an array schema, so cast safely
      const itemsSchema = schema.items as OpenAPIV3.SchemaObject;
      itemType = this.getTypeForProperty(itemsSchema, generatedTypes);
    }
    
    return `export type ${name} = ${itemType}[];\n\n`;
  }
  
  /**
   * Generate a TypeScript enum for an enum schema
   * @param name Name of the enum
   * @param schema OpenAPI schema object
   * @returns Generated TypeScript enum code
   */
  private generateEnumType(name: string, schema: OpenAPIV3.SchemaObject): string {
    const values = schema.enum || [];
    
    if (schema.type === 'string') {
      // String enum
      let enumCode = `export enum ${name} {\n`;
      
      values.forEach((value) => {
        const safeValue = String(value).replace(/[^\w]/g, '_');
        enumCode += `  ${safeValue} = '${value}',\n`;
      });
      
      enumCode += `}\n\n`;
      return enumCode;
    } else {
      // Union type for non-string enums
      const valueType = this.mapOpenAPITypeToTS(schema);
      const valuesList = values.map(v => {
        if (typeof v === 'string') return `'${v}'`;
        return v;
      }).join(' | ');
      
      return `export type ${name} = ${valuesList};\n\n`;
    }
  }
  
  /**
   * Generate a TypeScript union type for a oneOf/anyOf/allOf schema
   * @param name Name of the type
   * @param schema OpenAPI schema object
   * @param generatedTypes Set of already generated types
   * @returns Generated TypeScript union type code
   */
  private generateUnionType(
    name: string, 
    schema: OpenAPIV3.SchemaObject, 
    generatedTypes: Set<string>
  ): string {
    const variants = schema.oneOf || schema.anyOf || schema.allOf || [];
    const typeNames: string[] = [];
    
    variants.forEach((variant, index) => {
      const variantSchema = variant as OpenAPIV3.SchemaObject;
      
      if ('$ref' in variantSchema && typeof variantSchema.$ref === 'string') {
        // Reference to another schema
        const refName = this.parser.getSchemaNameFromRef(variantSchema.$ref);
        typeNames.push(refName);
      } else {
        // Inline schema
        const inlineName = `${name}Variant${index + 1}`;
        const inlineType = this.generateType(inlineName, variantSchema, generatedTypes);
        
        // Add the inline type
        typeNames.push(inlineName);
      }
    });
    
    // Use intersection for allOf, union for oneOf/anyOf
    const operator = schema.allOf ? ' & ' : ' | ';
    return `export type ${name} = ${typeNames.join(operator)};\n\n`;
  }
  
  /**
   * Get TypeScript type name for a property schema
   * @param schema OpenAPI schema object
   * @param generatedTypes Set of already generated types
   * @returns TypeScript type name
   */
  private getTypeForProperty(
    schema: OpenAPIV3.SchemaObject,
    generatedTypes: Set<string>
  ): string {
    // Check for reference
    if ('$ref' in schema && typeof schema.$ref === 'string') {
      const refName = this.parser.getSchemaNameFromRef(schema.$ref);
      return refName;
    }
    
    // Array type
    if (schema.type === 'array') {
      // Handle case where items might not be defined properly
      if (!schema.items) {
        return 'any[]';
      }
      
      // Cast to proper schema object
      const itemsSchema = schema.items as OpenAPIV3.SchemaObject;
      const itemType = this.getTypeForProperty(itemsSchema, generatedTypes);
      return `${itemType}[]`;
    }
    
    // Object type
    if (schema.type === 'object' || schema.properties) {
      // Generate inline interface for objects
      const properties = schema.properties || {};
      const propStrings: string[] = [];
      
      Object.entries(properties).forEach(([propName, propSchema]) => {
        const isRequired = schema.required?.includes(propName) || false;
        const propType = this.getTypeForProperty(propSchema as OpenAPIV3.SchemaObject, generatedTypes);
        propStrings.push(`${propName}${isRequired ? '' : '?'}: ${propType}`);
      });
      
      // Handle additional properties
      if (schema.additionalProperties) {
        const additionalPropsSchema = 
          typeof schema.additionalProperties === 'boolean' 
            ? { type: 'any' } 
            : schema.additionalProperties;
          
        const indexType = this.getTypeForProperty(additionalPropsSchema as OpenAPIV3.SchemaObject, generatedTypes);
        propStrings.push(`[key: string]: ${indexType}`);
      }
      
      return `{ ${propStrings.join('; ')} }`;
    }
    
    // Union type
    if (schema.oneOf || schema.anyOf) {
      const types = (schema.oneOf || schema.anyOf || []).map(s => 
        this.getTypeForProperty(s as OpenAPIV3.SchemaObject, generatedTypes)
      );
      return types.join(' | ');
    }
    
    // Primitive types
    return this.mapOpenAPITypeToTS(schema);
  }
  
  /**
   * Map OpenAPI type to TypeScript type
   * @param schema OpenAPI schema object
   * @returns TypeScript type name
   */
  private mapOpenAPITypeToTS(schema: OpenAPIV3.SchemaObject): string {
    if (!schema.type) {
      return 'any';
    }
    
    switch (schema.type) {
      case 'integer':
      case 'number':
        return 'number';
      case 'string':
        if (schema.format === 'date-time' || schema.format === 'date') {
          return 'string'; // Or consider 'Date'
        }
        if (schema.format === 'binary') {
          return 'Blob';
        }
        return 'string';
      case 'boolean':
        return 'boolean';
      case 'array':
        return 'any[]';
      case 'object':
        return 'Record<string, any>';
      default:
        return 'any';
    }
  }
  
  /**
   * Save generated TypeScript interfaces to a file
   * @param outputPath Path to save the generated file
   */
  public saveToFile(outputPath: string): void {
    const code = this.generateInterfaces();
    const dir = path.dirname(outputPath);
    
    // Create directory if it doesn't exist
    if (!fs.existsSync(dir)) {
      fs.mkdirSync(dir, { recursive: true });
    }
    
    fs.writeFileSync(outputPath, code, 'utf8');
  }
}