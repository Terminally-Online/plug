import * as Types from '../types/generated';

/**
 * Cache entry with timestamp for TTL validation
 */
interface CacheEntry<T> {
  timestamp: number;
  data: T;
}

/**
 * The schema cache key format
 */
type SchemaCacheKey = string;

/**
 * Schema cache service for caching API responses to minimize network requests
 */
export class SchemaCache {
  private cache: Map<SchemaCacheKey, CacheEntry<Record<string, Types.ActionsProtocolSchema>>>;
  private ttl: number;

  /**
   * Create a new schema cache instance
   * @param ttl Time to live in milliseconds (default: 5 minutes)
   */
  constructor(ttl: number = 5 * 60 * 1000) {
    this.cache = new Map();
    this.ttl = ttl;
  }

  /**
   * Generate a cache key from query parameters
   * @param queryParams Query parameters for the request
   * @returns A string key for the cache
   */
  private generateKey(queryParams?: Types.SchemaQueryParams): SchemaCacheKey {
    if (!queryParams) return 'all';
    
    // Sort the parameters for consistent key generation
    const params = Object.entries(queryParams)
      .filter(([_, value]) => value !== undefined)
      .sort(([keyA], [keyB]) => keyA.localeCompare(keyB));
      
    return params.length === 0 
      ? 'all'
      : params.map(([key, value]) => {
          if (key === 'search' && Array.isArray(value)) {
            return `${key}=${value.sort().join(',')}`;
          }
          return `${key}=${value}`;
        }).join('&');
  }

  /**
   * Check if a cache entry is still valid based on TTL
   * @param entry The cache entry to check
   * @returns Whether the entry is still valid
   */
  private isValid(entry: CacheEntry<any>): boolean {
    return Date.now() - entry.timestamp < this.ttl;
  }

  /**
   * Get a schema from the cache
   * @param queryParams Query parameters to generate the cache key
   * @returns The cached schema data or undefined if not found or expired
   */
  public get(queryParams?: Types.SchemaQueryParams): Record<string, Types.ActionsProtocolSchema> | undefined {
    const key = this.generateKey(queryParams);
    const entry = this.cache.get(key);
    
    if (entry && this.isValid(entry)) {
      return entry.data;
    }
    
    // Remove expired entries
    if (entry) {
      this.cache.delete(key);
    }
    
    return undefined;
  }

  /**
   * Store schema data in the cache
   * @param data Schema data to cache
   * @param queryParams Query parameters to generate the cache key
   */
  public set(data: Record<string, Types.ActionsProtocolSchema>, queryParams?: Types.SchemaQueryParams): void {
    const key = this.generateKey(queryParams);
    this.cache.set(key, {
      timestamp: Date.now(),
      data
    });
  }

  /**
   * Clear all entries from the cache
   */
  public clear(): void {
    this.cache.clear();
  }

  /**
   * Remove all expired entries from the cache
   * @returns Number of entries removed
   */
  public purgeExpired(): number {
    let count = 0;
    for (const [key, entry] of this.cache.entries()) {
      if (!this.isValid(entry)) {
        this.cache.delete(key);
        count++;
      }
    }
    return count;
  }
}