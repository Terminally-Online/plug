/**
 * Knot represents a data structure within the Rope
 */
export class Knot {
  private id: string;
  private data: Record<string, any>;

  /**
   * Create a new Knot
   * @param id Optional identifier for the knot
   * @param data Optional data to store in the knot
   */
  constructor(id?: string, data?: Record<string, any>) {
    this.id = id || this.generateId();
    this.data = data || {};
  }

  /**
   * Get the knot's ID
   * @returns The knot's ID
   */
  public getId(): string {
    return this.id;
  }

  /**
   * Set data in the knot
   * @param key Key to store the data under
   * @param value Value to store
   * @returns this instance for method chaining
   */
  public setData(key: string, value: any): this {
    this.data[key] = value;
    return this;
  }

  /**
   * Get data from the knot
   * @param key Key to retrieve
   * @returns The stored value or undefined if not found
   */
  public getData(key: string): any {
    return this.data[key];
  }

  /**
   * Get all data in the knot
   * @returns All data stored in the knot
   */
  public getAllData(): Record<string, any> {
    return { ...this.data };
  }

  /**
   * Generate a random ID
   * @returns A random string ID
   */
  private generateId(): string {
    return Math.random().toString(36).substring(2, 15);
  }
}
