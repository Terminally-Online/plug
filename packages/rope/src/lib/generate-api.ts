import * as path from 'path';
import { OpenAPIParser } from './openapi-parser';
import { TypeGenerator } from './type-generator';
import { ApiGenerator } from './api-generator';

/**
 * Generate API client and types from OpenAPI schema
 */
export function generateApi() {
  // Define paths
  const schemaPath = path.resolve(__dirname, 'openapi.json');
  const typesOutputPath = path.resolve(__dirname, '../types/generated.ts');
  const apiOutputPath = path.resolve(__dirname, '../models/generated-api.ts');
  
  // Create parser and generators
  const parser = new OpenAPIParser(schemaPath);
  const typeGenerator = new TypeGenerator(parser);
  const apiGenerator = new ApiGenerator(parser);
  
  // Generate and save types and API client
  console.log('Generating types from OpenAPI schema...');
  typeGenerator.saveToFile(typesOutputPath);
  console.log(`Types saved to ${typesOutputPath}`);
  
  console.log('Generating API client from OpenAPI schema...');
  apiGenerator.saveToFile(apiOutputPath);
  console.log(`API client saved to ${apiOutputPath}`);
  
  console.log('API generation complete!');
}

// Run the generator if this file is executed directly
if (require.main === module) {
  generateApi();
}