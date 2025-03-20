import { Knot } from './knot';
import { PlugAPI } from './api';
import * as Types from '../types/generated';
import { 
  NotInitializedError, 
  NotFoundError, 
  IntentError
} from '../services/errors';
import { CoilService } from '../services/coil';
import { SentenceService, ParsedSentence, SentenceInput } from '../services/sentence';

/**
 * Main SDK class for interacting with the Plug API
 * Represents a sequence of connected protocol actions (knots)
 * for a node-based editor experience
 */

export class Rope {
  private api: PlugAPI;
  private knots: Knot[];
  private availableSchemas: Record<string, Types.ActionsProtocolSchema>;
  private isInitialized: boolean = false;
  private coilService: CoilService;
  private sentenceService: SentenceService;

  /**
   * Create a new Rope instance
   * @param plugApiKey API key for authentication
   * @param config Optional configuration overrides
   */
  constructor(plugApiKey: string, config?: { 
    baseURL?: string;
    cacheTTL?: number;
    defaultChainId?: number;
  }) {
    this.api = new PlugAPI(plugApiKey, config);
    this.knots = [];
    this.availableSchemas = {};
    this.defaultChainId = config?.defaultChainId || 1; // Default to Ethereum mainnet
    this.coilService = new CoilService();
    this.sentenceService = new SentenceService();
  }
  
  /**
   * The default chain ID to use when not specified
   * @private
   */
  private defaultChainId: number;

  /**
   * Initialize the rope by fetching all available schemas
   * This loads the available protocol/action combinations that can be used as knots
   * @param chainId Optional chain ID to filter schemas (defaults to the defaultChainId)
   * @param options Optional initialization options
   * @returns Promise resolving when initialization is complete
   */
  public async initialize(
    chainId?: number,
    options?: {
      protocol?: string;
      action?: string;
      from?: string;
      search?: Types.IntentSearchQueryParam[];
      forceRefresh?: boolean;
    }
  ): Promise<void> {
    const shouldRefresh = !this.isInitialized || options?.forceRefresh;
    
    if (shouldRefresh) {
      // Determine which chain ID to use
      const effectiveChainId = chainId || this.defaultChainId;
      
      // Fetch schemas with optional filtering
      this.availableSchemas = await this.api.getSolver({
        chainId: effectiveChainId,
        protocol: options?.protocol,
        action: options?.action,
        from: options?.from,
        search: options?.search
      });
      
      this.isInitialized = true;
    }
  }
  
  /**
   * Initialize for a specific chain
   * Convenience method to initialize with a specific chain ID
   * @param chainId The chain ID to initialize with
   * @param options Optional initialization options
   * @returns Promise resolving when initialization is complete
   */
  public async initializeForChain(
    chainId: number,
    options?: Omit<Parameters<typeof this.initialize>[1], 'forceRefresh'>
  ): Promise<void> {
    return this.initialize(chainId, options);
  }

  /**
   * Check if the rope has been initialized
   * @returns Whether initialization has completed
   */
  public isReady(): boolean {
    return this.isInitialized;
  }

  /**
   * Get all available schemas that can be used to create knots
   * @returns Map of available protocol schemas
   */
  public getAvailableSchemas(): Record<string, Types.ActionsProtocolSchema> {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    return this.availableSchemas;
  }

  /**
   * Get schemas for a specific protocol
   * @param protocol Protocol name to filter by
   * @returns Schema for the specified protocol or undefined if not found
   */
  public getProtocolSchema(protocol: string): Types.ActionsProtocolSchema | undefined {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    return this.availableSchemas[protocol];
  }
  
  /**
   * Get supported chain IDs for a protocol
   * @param protocol Protocol name
   * @returns Array of supported chain IDs or empty array if protocol not found
   */
  public getProtocolChains(protocol: string): number[] {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    const protocolSchema = this.availableSchemas[protocol];
    if (!protocolSchema || !protocolSchema.metadata?.chains) {
      return [];
    }
    
    // Flatten array of chain IDs from all network objects
    return protocolSchema.metadata.chains.flatMap(network => network.chainIds || []);
  }
  
  /**
   * Check if a protocol supports a specific chain
   * @param protocol Protocol name
   * @param chainId Chain ID to check
   * @returns Whether the protocol supports the chain
   */
  public isProtocolSupportedOnChain(protocol: string, chainId: number): boolean {
    return this.getProtocolChains(protocol).includes(chainId);
  }
  
  /**
   * Get all protocols that support a specific chain
   * @param chainId Chain ID to filter by
   * @returns Array of protocol names that support the chain
   */
  public getProtocolsSupportingChain(chainId: number): string[] {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    return Object.entries(this.availableSchemas)
      .filter(([_, schema]) => {
        const chains = schema.metadata?.chains || [];
        return chains.some(network => (network.chainIds || []).includes(chainId));
      })
      .map(([protocol]) => protocol);
  }

  /**
   * Create a new knot for a specific protocol and action
   * @param protocol Protocol name
   * @param action Action name
   * @param data Initial data for the knot
   * @returns Newly created knot
   */
  public createKnot(protocol: string, action: string, data?: Record<string, any>): Knot {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    // Verify that the protocol and action exist in available schemas
    const protocolSchema = this.availableSchemas[protocol];
    if (!protocolSchema) {
      throw new NotFoundError(`Protocol "${protocol}" not found in available schemas`);
    }
    
    if (protocolSchema.schema && !protocolSchema.schema[action]) {
      throw new NotFoundError(`Action "${action}" not found in protocol "${protocol}"`);
    }
    
    // Create and add the knot
    const knot = new Knot(undefined, { ...data, protocol, action });
    return knot;
  }
  
  /**
   * Create a new knot with empty values for all inputs based on sentence parsing
   * @param protocol Protocol name
   * @param action Action name
   * @returns Newly created knot with empty values for all inputs
   */
  public createKnotWithEmptyValues(protocol: string, action: string): Knot {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    // Verify that the protocol and action exist in available schemas
    const protocolSchema = this.availableSchemas[protocol];
    if (!protocolSchema) {
      throw new NotFoundError(`Protocol "${protocol}" not found in available schemas`);
    }
    
    const actionSchema = protocolSchema.schema?.[action];
    if (!actionSchema) {
      throw new NotFoundError(`Action "${action}" not found in protocol "${protocol}"`);
    }
    
    // Parse the sentence to get input structure
    const sentence = actionSchema.sentence || '';
    const parsed = this.sentenceService.parseSentence(sentence);
    const emptyValues = this.sentenceService.generateEmptyValues(parsed);
    
    // Convert empty values to the format expected by Knot
    const data: Record<string, any> = {
      protocol,
      action,
      values: Object.entries(emptyValues).reduce((acc, [index, value]) => {
        acc[index] = {
          index,
          key: index,
          name: `input${index}`,
          value
        };
        return acc;
      }, {} as Record<string, any>)
    };
    
    return new Knot(undefined, data);
  }

  /**
   * Get the API client instance
   * @returns The API client instance
   */
  public getApi(): PlugAPI {
    return this.api;
  }
  
  /**
   * Search for protocols by name
   * @param query Search query string
   * @returns Array of matching protocol names
   */
  public searchProtocols(query: string): string[] {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    const lowercaseQuery = query.toLowerCase();
    return Object.keys(this.availableSchemas)
      .filter(protocol => protocol.toLowerCase().includes(lowercaseQuery));
  }
  
  /**
   * Search for actions within a protocol
   * @param protocol Protocol name to search within
   * @param query Search query string
   * @returns Array of matching action names
   */
  public searchActionsInProtocol(protocol: string, query: string): string[] {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    const protocolSchema = this.availableSchemas[protocol];
    if (!protocolSchema || !protocolSchema.schema) {
      return [];
    }
    
    const lowercaseQuery = query.toLowerCase();
    return Object.keys(protocolSchema.schema)
      .filter(action => action.toLowerCase().includes(lowercaseQuery));
  }
  
  /**
   * Search for actions across all protocols
   * @param query Search query string
   * @returns Array of {protocol, action} pairs that match the query
   */
  public searchAllActions(query: string): Array<{protocol: string; action: string}> {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    const results: Array<{protocol: string; action: string}> = [];
    const lowercaseQuery = query.toLowerCase();
    
    // First, search protocol names
    const matchingProtocols = this.searchProtocols(query);
    
    // For matching protocols, add all their actions
    for (const protocol of matchingProtocols) {
      const protocolSchema = this.availableSchemas[protocol];
      if (protocolSchema && protocolSchema.schema) {
        Object.keys(protocolSchema.schema).forEach(action => {
          results.push({ protocol, action });
        });
      }
    }
    
    // Then search all actions in all protocols
    for (const [protocol, schema] of Object.entries(this.availableSchemas)) {
      if (!matchingProtocols.includes(protocol) && schema.schema) {
        // Only search actions in protocols that weren't already fully included
        Object.keys(schema.schema)
          .filter(action => action.toLowerCase().includes(lowercaseQuery))
          .forEach(action => {
            results.push({ protocol, action });
          });
      }
    }
    
    return results;
  }
  
  /**
   * Filter protocols by tag
   * @param tag Tag to filter by
   * @returns Array of protocol names with the specified tag
   */
  public getProtocolsByTag(tag: string): string[] {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    const lowercaseTag = tag.toLowerCase();
    return Object.entries(this.availableSchemas)
      .filter(([_, schema]) => {
        const tags = schema.metadata?.tags || [];
        return tags.some(t => t.toLowerCase() === lowercaseTag);
      })
      .map(([protocol]) => protocol);
  }

  /**
   * Get the coil service instance
   * @returns The coil service instance
   */
  public getCoilService(): CoilService {
    return this.coilService;
  }
  
  /**
   * Get the sentence service instance
   * @returns The sentence service instance
   */
  public getSentenceService(): SentenceService {
    return this.sentenceService;
  }
  
  /**
   * Parse a sentence from an action schema
   * @param protocol Protocol name
   * @param action Action name
   * @returns Parsed sentence or null if not found
   */
  public parseSentence(protocol: string, action: string): ParsedSentence | null {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    const protocolSchema = this.availableSchemas[protocol];
    if (!protocolSchema || !protocolSchema.schema) {
      return null;
    }
    
    const actionSchema = protocolSchema.schema[action];
    if (!actionSchema || !actionSchema.sentence) {
      return null;
    }
    
    return this.sentenceService.parseSentence(actionSchema.sentence);
  }
  
  /**
   * Validate input values for a parsed sentence
   * @param parsedSentence The parsed sentence
   * @param values Input values to validate
   * @returns Map of input indices to validation errors, empty if all valid
   */
  public validateInputs(
    parsedSentence: ParsedSentence,
    values: Record<string, any>
  ): Map<number, string> {
    const errors = new Map<number, string>();
    
    parsedSentence.inputs.forEach((input: SentenceInput) => {
      const value = values[input.index];
      if (input.required && (value === undefined || value === '')) {
        errors.set(input.index, 'This field is required');
        return;
      }
      
      if (value !== undefined && value !== '') {
        const validation = this.sentenceService.validateInput(value, input.type);
        if (!validation.success) {
          errors.set(input.index, validation.error || 'Invalid value');
        }
      }
    });
    
    return errors;
  }
  
  /**
   * Find compatible coils for an action's inputs
   * @param protocol Protocol name
   * @param action Action name
   * @returns Map of input indices to arrays of compatible coil names
   */
  public findCompatibleCoils(protocol: string, action: string): Record<number, string[]> {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    // Parse the sentence for this action
    const parsed = this.parseSentence(protocol, action);
    if (!parsed) {
      return {};
    }
    
    // Get all available coils
    const availableCoils = this.getAvailableCoils();
    
    // Find compatible coils for each input
    return this.coilService.getCompatibleCoilsForParsedSentence(parsed, availableCoils);
  }
  
  /**
   * Resolve coil references in input values
   * @param values Input values with possible coil references
   * @param availableCoils Available coil values
   * @returns New values with resolved coil references
   */
  public resolveCoilReferences(
    values: Record<string, any>,
    availableCoils: Record<string, any>
  ): Record<string, any> {
    return this.coilService.resolveCoilReferences(values, availableCoils);
  }
  
  /**
   * Format a sentence with input values
   * @param protocol Protocol name
   * @param action Action name
   * @param values Input values to insert
   * @returns Formatted sentence or null if action not found
   */
  public formatActionSentence(
    protocol: string,
    action: string,
    values: Record<string, any>
  ): string | null {
    const parsed = this.parseSentence(protocol, action);
    if (!parsed) {
      return null;
    }
    
    return this.sentenceService.formatSentence(parsed, values);
  }
  
  /**
   * Check if an action has coils
   * @param protocol Protocol name
   * @param action Action name
   * @returns Whether the action has output coils
   */
  public hasCoils(protocol: string, action: string): boolean {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    const protocolSchema = this.availableSchemas[protocol];
    if (!protocolSchema || !protocolSchema.schema) {
      return false;
    }
    
    const actionSchema = protocolSchema.schema[action];
    return Object.keys(this.coilService.getCoilsFromSchema(actionSchema)).length > 0;
  }
  
  /**
   * Get all coils for an action
   * @param protocol Protocol name
   * @param action Action name
   * @returns Map of coil names to their types, or empty object if none
   */
  public getCoils(protocol: string, action: string): Record<string, string> {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    const protocolSchema = this.availableSchemas[protocol];
    if (!protocolSchema || !protocolSchema.schema) {
      return {};
    }
    
    const actionSchema = protocolSchema.schema[action];
    return this.coilService.getCoilsFromSchema(actionSchema);
  }
  
  /**
   * Get all available coils in the rope up to a specific knot index
   * @param upToIndex Only include coils from knots before this index
   * @returns Map of coil names to their types
   */
  public getAvailableCoils(upToIndex?: number): Record<string, string> {
    if (!this.isInitialized) {
      throw new NotInitializedError();
    }
    
    const result: Record<string, string> = {};
    const limit = upToIndex !== undefined ? Math.min(upToIndex, this.knots.length) : this.knots.length;
    
    // Collect coils from each knot
    for (let i = 0; i < limit; i++) {
      const knot = this.knots[i];
      const data = knot.getAllData();
      const protocol = data.protocol as string;
      const action = data.action as string;
      
      // Get coils for this action and add to result
      const coils = this.getCoils(protocol, action);
      Object.entries(coils).forEach(([name, type]) => {
        // Include the knot index in the coil reference
        result[`${name}@${i}`] = type;
        
        // Also include without index for backward compatibility
        if (!result[name]) {
          result[name] = type;
        }
      });
    }
    
    return result;
  }
  
  /**
   * Get compatible coils for a specific input type
   * @param inputType Type expected by the input
   * @param upToIndex Only include coils from knots before this index
   * @returns Map of compatible coil names to their types
   */
  public getCompatibleCoils(inputType: string, upToIndex?: number): Record<string, string> {
    const allCoils = this.getAvailableCoils(upToIndex);
    return this.coilService.getCompatibleCoils(allCoils, inputType);
  }

  /**
   * Add a new knot to the rope
   * @param knot Knot to add
   * @returns this instance for method chaining
   */
  public addKnot(knot: Knot): this {
    this.knots.push(knot);
    return this;
  }

  /**
   * Get all knots in the rope
   * @returns Array of knots
   */
  public getKnots(): Knot[] {
    return this.knots;
  }
  
  /**
   * Generate an intent from the current rope of knots
   * Creates a transaction intent that represents the sequence of actions
   * @returns Intent object that can be submitted to the API
   */
  public buildIntent(): Types.ModelsIntent {
    if (this.knots.length === 0) {
      throw new IntentError('Cannot build intent: rope has no knots');
    }
    
    // Extract input data from each knot
    const inputs = this.knots.map(knot => {
      const data = knot.getAllData();
      return {
        protocol: data.protocol,
        action: data.action,
        params: { ...data }
      };
    });
    
    // Create the intent structure
    const intent: Types.ModelsIntent = {
      chainId: this.defaultChainId, // Use the configured default chain ID
      from: '', // To be filled by user
      frequency: 0, // One-time execution
      inputs,
      locked: false,
      saved: false,
      status: 'active'
    };
    
    return intent;
  }
  
  /**
   * Submit the current rope as an intent to the solver
   * @param fromAddress Ethereum address to execute the intent from
   * @param chainId Optional chain ID (defaults to the instance's defaultChainId)
   * @returns Result from the solver
   */
  public async execute(fromAddress: string, chainId?: number): Promise<Record<string, any>> {
    const intent = this.buildIntent();
    intent.from = fromAddress;
    intent.chainId = chainId || this.defaultChainId;
    
    return this.api.postSolver(intent);
  }
}
