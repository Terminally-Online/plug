export function transformIPFSUrl(url: string | undefined): string | undefined {
    if (!url) return undefined
    
    if (url.startsWith('ipfs://')) {
        const hash = url.replace('ipfs://', '')
        return `https://dweb.link/ipfs/${hash}`
    }
    
    return url
}