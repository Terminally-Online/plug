export class API { 
    constructor(
        public readonly api = 'https://api.onplug.io/pool',
		public readonly apiKey = 'AAAAAAAAAAAAAAAAAAAA'
    ) {}

    async post(body: unknown): Promise<unknown> { 
        const response = await fetch(this.api, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
                'X-API-KEY': this.apiKey
			},
			body: JSON.stringify(body)
		})

		return response.json()
    }

    // Note: We will probably end up using TRPC for the real SDK.
    //       The current implementation is just a placeholder for raw testing.
}