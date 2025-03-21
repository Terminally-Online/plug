/**
 * LRU cache implementation with max size and optional TTL
 * General-purpose cache utility for the Rope package
 */
export class LRUCache<K, V> {
  private cache = new Map<K, { value: V, expires?: number }>();
  private maxSize: number;
  private defaultTTL?: number;

  /**
   * Create a new LRU cache with size limits and optional TTL
   * @param options Configuration options for the cache
   */
  constructor(options: { maxSize: number, ttl?: number }) {
    this.maxSize = options.maxSize;
    this.defaultTTL = options.ttl;
  }

  /**
   * Get a value from the cache
   * @param key The cache key
   * @returns The cached value or undefined if not found/expired
   */
  get(key: K): V | undefined {
    const item = this.cache.get(key);
    
    // Handle cache miss
    if (!item) return undefined;
    
    // Check expiration if TTL was set
    if (item.expires && Date.now() > item.expires) {
      this.cache.delete(key);
      return undefined;
    }
    
    // Move to the end of the map for LRU behavior
    this.cache.delete(key);
    this.cache.set(key, item);
    
    return item.value;
  }

  /**
   * Set a value in the cache
   * @param key The cache key
   * @param value The value to cache
   * @param ttl Optional TTL in milliseconds (overrides default)
   */
  set(key: K, value: V, ttl?: number): void {
    // If at capacity, remove the oldest item (first in the Map)
    if (this.cache.size >= this.maxSize) {
      const iterator = this.cache.keys();
      const firstItem = iterator.next();
      if (!firstItem.done && firstItem.value !== undefined) {
        this.cache.delete(firstItem.value);
      }
    }
    
    // Calculate expiration if TTL is provided
    const expires = ttl || this.defaultTTL 
      ? Date.now() + (ttl || this.defaultTTL || 0)
      : undefined;
    
    // Add to cache
    this.cache.set(key, { value, expires });
  }

  /**
   * Check if a key exists in the cache and is not expired
   * @param key The cache key to check
   * @returns True if the key exists and is not expired
   */
  has(key: K): boolean {
    const item = this.cache.get(key);
    if (!item) return false;
    
    // Check expiration
    if (item.expires && Date.now() > item.expires) {
      this.cache.delete(key);
      return false;
    }
    
    return true;
  }

  /**
   * Remove a specific key from the cache
   * @param key The key to remove
   * @returns True if the key was found and removed
   */
  delete(key: K): boolean {
    return this.cache.delete(key);
  }

  /**
   * Clear the entire cache
   */
  clear(): void {
    this.cache.clear();
  }

  /**
   * Get the current number of items in the cache
   */
  get size(): number {
    return this.cache.size;
  }

  /**
   * Get all keys in the cache
   * @returns Array of all keys
   */
  keys(): K[] {
    return Array.from(this.cache.keys());
  }
  
  /**
   * Get an iterator over the keys in the cache
   * For compatibility with Map interface
   */
  [Symbol.iterator](): IterableIterator<[K, { value: V, expires?: number }]> {
    return this.cache[Symbol.iterator]();
  }

  /**
   * Remove all expired entries from the cache
   * @returns Number of entries removed
   */
  prune(): number {
    if (!this.defaultTTL) return 0;
    
    const now = Date.now();
    let count = 0;
    
    for (const [key, item] of this.cache.entries()) {
      if (item.expires && now > item.expires) {
        this.cache.delete(key);
        count++;
      }
    }
    
    return count;
  }
}