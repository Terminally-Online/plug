import { useEffect } from 'react';
import type { UseTRPCQueryResult } from '@trpc/react-query/shared';

/**
 * A hook that wraps tRPC query hooks to provide callback functionality
 * similar to the removed React Query v3 callbacks
 */
export function useResponse<TData, TError>(
  queryHook: () => UseTRPCQueryResult<TData, TError>,
  callbacks?: {
    onSuccess?: (data: TData) => void;
    onError?: (error: TError) => void;
    onLoading?: () => void;
    onSettled?: (data: TData | undefined, error: TError | undefined) => void;
  }
): UseTRPCQueryResult<TData, TError> {
  const query = queryHook();
  const { data, isLoading, isError, error, isSuccess } = query;

  useEffect(() => {
    if (isSuccess && callbacks?.onSuccess && data !== undefined) {
      // Force type assertion here to handle the async iterable issue
      callbacks.onSuccess(data as TData);
    }
  }, [isSuccess, data, callbacks]);

  useEffect(() => {
    if (isError && callbacks?.onError && error !== null) {
      callbacks.onError(error as TError);
    }
  }, [isError, error, callbacks]);

  useEffect(() => {
    if (isLoading && callbacks?.onLoading) {
      callbacks.onLoading();
    }
  }, [isLoading, callbacks]);

  useEffect(() => {
    if (!isLoading && callbacks?.onSettled) {
      callbacks.onSettled(data as TData | undefined, error as TError | undefined);
    }
  }, [isLoading, data, error, callbacks]);

  return query;
}
