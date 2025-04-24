import { TRPCError } from "@trpc/server"

import axios, { AxiosResponse } from "axios"
import { z } from "zod"

import { env } from "@/env"

export const zerion = axios.create({
	baseURL: "https://api.zerion.io/v1",
	headers: {
		accept: "application/json",
		authorization: `Basic ${env.ZERION_KEY}`
	}
})

/**
 * Wrapper for Zerion API calls that handles error checking
 * @param apiCall - Function that makes the actual API call
 * @returns The data from the API response
 */
export const zerionApi = async <T>(apiCall: () => Promise<AxiosResponse>): Promise<T> => {
	try {
		const response = await apiCall()
		return response.data as T
	} catch (error) {
		if (axios.isAxiosError(error)) {
			throw new TRPCError({
				code: "INTERNAL_SERVER_ERROR",
				message: `Zerion API error: ${error.response?.data?.message || error.message}`
			})
		}
		throw new TRPCError({
			code: "INTERNAL_SERVER_ERROR",
			message: "Failed to fetch data from Zerion"
		})
	}
}

// Zod helper to transform ISO date strings into Date objects
export const isoDateString = z.string().transform((isoString, ctx) => {
	try {
		return new Date(isoString)
	} catch (error) {
		ctx.addIssue({
			code: z.ZodIssueCode.custom,
			message: "Invalid ISO date format"
		})
		return z.NEVER
	}
})

/**
 * Builds query parameters for a Zerion API call
 * @param params Object containing parameters to be converted into query string
 * @returns Formatted query string (with leading '?' if not empty)
 */
export const buildQueryParams = (params: Record<string, any>): string => {
	const filteredParams = Object.entries(params).filter(
		([_, value]) => value !== undefined && value !== null && value !== ""
	)

	if (filteredParams.length === 0) return ""

	const queryParts = filteredParams.map(([key, value]) => {
		if (Array.isArray(value)) {
			if (key.includes("filter[")) {
				return `${key}=${value.join(",")}`
			}
			return value.map(item => `${key}[]=${item}`).join("&")
		}
		return `${key}=${encodeURIComponent(value)}`
	})

	return `?${queryParts.join("&")}`
}
