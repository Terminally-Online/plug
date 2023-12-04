import EventEmitter from "events"

const globalForEmitter = globalThis as unknown as {
	emitter: EventEmitter | undefined
}

export const emitter = globalForEmitter.emitter ?? new EventEmitter()

if (process.env.NODE_ENV !== "production") {
	globalForEmitter.emitter = emitter
}
