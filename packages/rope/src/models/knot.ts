/**
 * Default data type for a Knot when none is specified
 */
export interface DefaultKnotData {
  [key: string]: unknown;
}

/**
 * Knot represents a data structure within the Rope
 * The generic parameter T allows for type-safe data storage
 */
export class Knot<T extends Record<string, unknown> = DefaultKnotData> {
  private id: string;
  private data: T;

  /**
   * Create a new Knot
   * @param id Optional identifier for the knot
   * @param data Optional data to store in the knot
   */
  constructor(id?: string, data?: T) {
    this.id = id || this.generateId();
    this.data = data || ({} as T);
  }

  /**
   * Get the knot's ID
   * @returns The knot's ID
   */
  public getId(): string {
    return this.id;
  }

  /**
   * Set data in the knot with type-safety
   * @param key Key to store the data under
   * @param value Value to store
   * @returns this instance for method chaining
   */
  public setData<K extends keyof T>(key: K, value: T[K]): this {
    this.data[key] = value;
    return this;
  }

  /**
   * Get data from the knot with type-safety
   * @param key Key to retrieve
   * @returns The stored value with correct type or undefined if not found
   */
  public getData<K extends keyof T>(key: K): T[K] | undefined {
    return this.data[key];
  }

  /**
   * Get all data in the knot
   * @returns A copy of all data stored in the knot
   */
  public getAllData(): T {
    return { ...this.data };
  }

  /**
   * Shallow merge multiple values into knot data
   * @param values Object with values to merge into knot data
   * @returns this instance for method chaining
   */
  public mergeData(values: Partial<T>): this {
    this.data = { ...this.data, ...values };
    return this;
  }

  /**
   * Generate a random ID
   * @returns A random string ID
   */
  private generateId(): string {
    return Math.random().toString(36).substring(2, 15) + 
           Math.random().toString(36).substring(2, 15);
  }
}
