import { usePathname } from "next/navigation"
import { useCallback, useEffect, useRef, useState } from "react"

const DEFAULT_DELAY = 250

export const useDebounce = <T,>(
	initial: T,
	delay = DEFAULT_DELAY,
	callback = (_: T) => { }
): [T, T, (value: T) => void, React.ForwardedRef<T>] => {
	const pathname = usePathname()

	const valueRef = useRef<T>(initial)

	const [value, setValue] = useState<T>(initial)
	const [debounced, setDebounced] = useState<typeof value>(value)

	const debouncedCallback = useCallback(() => {
		setDebounced(value)
		if (callback) callback(value)
	}, [value, callback])

	useEffect(() => {
		const timeout = setTimeout(debouncedCallback, delay)

		return () => clearTimeout(timeout)
	}, [value, delay, debouncedCallback])

	useEffect(() => {
		if (!pathname) return

		setValue(initial)
		setDebounced(initial)
	}, [pathname, initial])

	return [value, debounced, setValue, valueRef]
}

export const useDebounceInline = <T,>(value: T, delay = DEFAULT_DELAY, callback?: (value: T) => void): T => {
  const [debouncedValue, setDebouncedValue] = useState<T>(value);
  const callbackRef = useRef(callback);
  const previousValueRef = useRef<T>(value);
  
  useEffect(() => {
    callbackRef.current = callback;
  }, [callback]);
  
  useEffect(() => {
    // Skip if the value is the same by reference
    // This works for primitives and prevents unnecessary stringification for objects
    if (value === previousValueRef.current) {
      return;
    }
    
    // For objects, we update the reference regardless
    // This avoids deep equality checks which can be problematic
    previousValueRef.current = value;
    
    const timeout = setTimeout(() => {
      setDebouncedValue(value);
      callbackRef.current?.(value);
    }, delay);
    
    return () => clearTimeout(timeout);
  }, [value, delay]);
  
  return debouncedValue;
};
