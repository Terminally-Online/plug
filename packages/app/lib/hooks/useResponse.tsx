import { useEffect, useRef } from "react"

import type { UseTRPCQueryResult } from "@trpc/react-query/shared"

type Callbacks<TData, TError> = {
	onSuccess?: (data: TData) => void
	onError?: (error: TError) => void
	onLoading?: () => void
	onSettled?: (data: TData | undefined, error: TError | undefined) => void
}

/**
 * A hook that wraps tRPC query hooks to provide callback functionality
 * similar to the removed React Query v3 callbacks
 */
export function useResponse<TData, TError>(
	queryHook: () => UseTRPCQueryResult<TData, TError>,
	callbacks?: Callbacks<TData, TError>
): UseTRPCQueryResult<TData, TError> {
	const query = queryHook()
	const { data, isLoading, isError, error, isSuccess } = query

	const callbackRefs = useRef(callbacks)

	useEffect(() => {
		callbackRefs.current = callbacks
	})

	useEffect(() => {
		if (isSuccess && data !== undefined) {
			callbackRefs.current?.onSuccess?.(data as TData)
		}
	}, [isSuccess, data])

	useEffect(() => {
		if (isError && error !== null) {
			callbackRefs.current?.onError?.(error as TError)
		}
	}, [isError, error])

	useEffect(() => {
		if (isLoading) {
			callbackRefs.current?.onLoading?.()
		}
	}, [isLoading])

	useEffect(() => {
		if (!isLoading) {
			callbackRefs.current?.onSettled?.(data as TData | undefined, error as TError | undefined)
		}
	}, [isLoading, data, error])

	return query
}
