
        export const PAGE_SIZE = 20;

        export interface Post {
            filename: string;
            slug: string;
            title: string;
            description: string;
            image: string;
            content: string;
            attributes: {
                created: string;
            } & Partial<{
                updated: string;
                tags: string[];
                related: string[];
                inbound: string[];
                author: string;
                // display settings
                imagePosition: 'top' | 'bottom';
                imagePadded: 'true' | 'false';
                className: string;
                variant: string;
                unlisted: 'true' | 'false';
                sidebar: 'show' | 'hide' | undefined;
            }>
        }

        export type Posts = Record<string, Post>;

        export const faviconUrls = {"docs.onplug.io":"data:image/vnd.microsoft.icon;base64,AAABAAMAMDAAAAEAIACoJQAANgAAADAwAAABACAAqBAAAN4lAAAwMAAAAQAgAGgEAACGNgAAKAAAADAAAABgAAAAAQAgAAAAAAAoCQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlKGPYoKQe7xrfWTrTGFD/kJYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0xhQ/5rfWTrgpB7vJShj2IAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJmklUZ/j3nEV2xP+0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1dsT/t/j3nEmaSVRgAAAAAAAAAAAAAAAAAAAAAAAAAAmaaTUHGBauNCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/cYFq45mmk1AAAAAAAAAAAAAAAACZpJVGcYFq40JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/3GBauOZpJVGAAAAAAAAAAB/j3nEQlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9/j3nEAAAAAJShj2JXbE/7Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9XbE/7lKGPYoKQe7xCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/gpB7vGt9ZOtCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/a31k60xhQ/5CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/TGFD/kJYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1R/X/9Uf1//VH9f/1R/X/9Uf1//VH9f/1R/X/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9mpoX/ivPS/4rz0v+K89L/ivPS/4rz0v+K89L/ivPS/4rz0v+K89L/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/2amhf+K89L/ivPS/2amhf9mpoX/ZqaF/2amhf9mpoX/ZqaF/3jMq/+K89L/ivPS/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/3jMq/+K89L/VH9f/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9mpoX/ivPS/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9Uf1//ivPS/3jMq/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9Uf1//ivPS/3jMq/9mpoX/ZqaF/2amhf9Uf1//VH9f/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9Uf1//ZqaF/3jMq/9mpoX/ivPS/3jMq/94zKv/ivPS/4rz0v+K89L/ivPS/4rz0v94zKv/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/2amhf+K89L/ivPS/4rz0v9mpoX/ivPS/3jMq/9Uf1//ZqaF/2amhf9mpoX/eMyr/3jMq/+K89L/ivPS/3jMq/9Uf1//Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/eMyr/4rz0v+K89L/ZqaF/0JYOP9Uf1//ivPS/3jMq/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ZqaF/4rz0v+K89L/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9Uf1//ivPS/3jMq/9CWDj/Qlg4/0JYOP9Uf1//ivPS/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1R/X/94zKv/ivPS/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9mpoX/ivPS/1R/X/9CWDj/Qlg4/0JYOP9mpoX/ivPS/1R/X/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9Uf1//ivPS/4rz0v9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9mpoX/ivPS/2amhf9CWDj/Qlg4/0JYOP94zKv/ivPS/1R/X/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ZqaF/4rz0v9mpoX/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9mpoX/ivPS/3jMq/9Uf1//ZqaF/3jMq/+K89L/eMyr/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/ZqaF/4rz0v+K89L/ivPS/4rz0v9mpoX/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/2amhf+K89L/VH9f/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/1R/X/9mpoX/ZqaF/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/2amhf+K89L/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1R/X/+K89L/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1R/X/+K89L/eMyr/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1R/X/+K89L/VH9f/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/3jMq/+K89L/VH9f/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/3jMq/+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ZqaF/2amhf9Uf1//Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ZqaF/4rz0v9mpoX/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9mpoX/ivPS/4rz0v+K89L/eMyr/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9Uf1//ivPS/3jMq/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/1R/X/+K89L/eMyr/2amhf94zKv/ivPS/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/2amhf+K89L/ivPS/1R/X/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP+K89L/ivPS/0JYOP9mpoX/eMyr/1R/X/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9mpoX/eMyr/4rz0v+K89L/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9mpoX/ivPS/3jMq/94zKv/ZqaF/2amhf9mpoX/ZqaF/2amhf9mpoX/eMyr/4rz0v+K89L/ivPS/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/eMyr/4rz0v+K89L/ivPS/4rz0v+K89L/ivPS/4rz0v+K89L/ivPS/3jMq/9mpoX/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/3jMq/+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9Uf1//ZqaF/2amhf9Uf1//VH9f/1R/X/9Uf1//Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/2amhf+K89L/eMyr/2amhf9mpoX/ZqaF/2amhf9mpoX/ZqaF/2amhf+K89L/ivPS/1R/X/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP94zKv/ivPS/4rz0v+K89L/ivPS/4rz0v+K89L/ivPS/4rz0v+K89L/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/VH9f/1R/X/9mpoX/ZqaF/2amhf9mpoX/ZqaF/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0xhQ/5CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/TGFD/mt9ZOtCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/a31k64KQe7xCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/gpB7vJShj2JXbE/7Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9XbE/7lKGPYgAAAAB/j3nEQlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9/j3nEAAAAAAAAAACZpJVGcYFq40JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/3GBauOZpJVGAAAAAAAAAAAAAAAAmaaTUHGBauNCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/cYFq45mmk1AAAAAAAAAAAAAAAAAAAAAAAAAAAJmklUZ/j3nEV2xP+0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1dsT/t/j3nEmaSVRgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlKGPYoKQe7xrfWTrTGFD/kJYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0xhQ/5rfWTrgpB7vJShj2IAAAAAAAAAAAAAAAAAAAAAAAAAAPgAAAAAHwAA4AAAAAAHAADAAAAAAAMAAIAAAAAAAQAAgAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAABAACAAAAAAAEAAMAAAAAAAwAA4AAAAAAHAAD4AAAAAB8AACgAAAAgAAAAQAAAAAEAIAAAAAAAKAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACLmIaab4Bn5lRoS/xCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1RoS/xvgGfmi5iGmgAAAAAAAAAAAAAAAAAAAACXpJFReIhx1kVaO/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9FWjv/eIhx1pekkVEAAAAAAAAAAHiIcdZCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/eIhx1gAAAACLmIaaRVo7/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9FWjv/i5iGmm+AZ+ZCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9vgGfmVGhL/EJYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1RoS/xCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9Uf1//ZqaF/2amhf9mpoX/ZqaF/1R/X/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/VH9f/4rz0v94zKv/eMyr/3jMq/94zKv/eMyr/4rz0v9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP94zKv/VH9f/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ivPS/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v9Uf1//Qlg4/0JYOP9CWDj/Qlg4/0JYOP94zKv/ZqaF/1R/X/9Uf1//VH9f/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ivPS/1R/X/9CWDj/Qlg4/1R/X/9mpoX/eMyr/3jMq/94zKv/ivPS/4rz0v+K89L/ivPS/3jMq/9Uf1//Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP+K89L/VH9f/0JYOP9Uf1//ivPS/2amhf9mpoX/eMyr/2amhf9CWDj/Qlg4/0JYOP9Uf1//ZqaF/4rz0v9mpoX/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v9Uf1//Qlg4/3jMq/9mpoX/Qlg4/0JYOP94zKv/VH9f/0JYOP9CWDj/Qlg4/0JYOP9CWDj/VH9f/4rz0v9mpoX/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ivPS/1R/X/9CWDj/ivPS/1R/X/9CWDj/VH9f/4rz0v9Uf1//Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ZqaF/4rz0v9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP+K89L/VH9f/0JYOP9mpoX/eMyr/2amhf+K89L/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/eMyr/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v9Uf1//Qlg4/0JYOP94zKv/eMyr/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9mpoX/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ivPS/1R/X/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/2amhf94zKv/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP+K89L/VH9f/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/eMyr/3jMq/9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v9Uf1//Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP94zKv/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ivPS/1R/X/9CWDj/Qlg4/0JYOP9CWDj/VH9f/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/VH9f/4rz0v9Uf1//Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP+K89L/VH9f/0JYOP9CWDj/VH9f/4rz0v+K89L/eMyr/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP+K89L/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/4rz0v9Uf1//Qlg4/0JYOP9mpoX/ivPS/0JYOP94zKv/VH9f/0JYOP9CWDj/Qlg4/0JYOP9mpoX/ivPS/2amhf9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/ivPS/1R/X/9CWDj/Qlg4/0JYOP+K89L/eMyr/3jMq/94zKv/eMyr/3jMq/+K89L/ivPS/4rz0v9mpoX/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP+K89L/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9Uf1//ZqaF/2amhf9mpoX/ZqaF/2amhf9Uf1//Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/2amhf+K89L/eMyr/2amhf9mpoX/ZqaF/2amhf+K89L/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1R/X/9mpoX/ZqaF/2amhf9mpoX/ZqaF/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/VGhL/EJYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/1RoS/xvgGfmQlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/b4Bn5ouYhppFWjv/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0VaO/+LmIaaAAAAAHiIcdZCWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JYOP9CWDj/Qlg4/0JY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"} as const;

        export const posts: Posts = {helloworld:{filename:"say-hello",slug:"hello-world",title:"Hello World.",description:"In a market that never sleeps, Plug is your heartbeat that keeps your onchain presence constant....",image:"/cdn/papers/hello-world.png",content:"\n\nTLDR: After a year of building in stealth, Plug is ready for you to try. We're making onchain automation accessible to everyone, not just trading firms. Use my access code [abcd-1234-efgh-5678] to skip the waitlist and start building your first automated strategy today.\n\nHey -- I've been quiet about what we're building at [Plug](/) for the past year. Today, that changes.\n\nIf you're like me, you've probably missed countless opportunities because you were sleeping, in a meeting, or simply away from your computer. The reality is that crypto never sleeps, but we're not machines – we need rest, we have jobs, and we have lives outside of watching charts and Discord announcements.\n\nThat's why we built Plug. It's your always-on presence in the digital world, executing your strategies exactly as you'd do them yourself. No more missed airdrops, no more liquidations because you couldn't react in time, and no more manual compounding of your DeFi positions.\n\n## Where We Are Today\n\nToday marks an important milestone as we open Plug's beta to everyone. From day one, you'll be able to automate your onchain presence across major networks like Ethereum, Optimism, Base, and Arbitrum. Connect with the protocols you already use through our simple drag-and-drop interface, set up custom triggers based on price movements or time schedules, and monitor everything in real-time.\n\nWhat excites me most is how we've made complex blockchain interactions feel natural. You don't need to understand the underlying mechanics – just tell Plug what you want to achieve, and it figures out the best way to make it happen.\n\n## The Road Ahead\n\nThis is just the beginning. We're already working with major protocols to expand our integration suite. Soon, you'll be able to create even more sophisticated automations, share your successful strategies with others (and yes, earn from them), and manage everything from your phone.\n\nThe vision is clear: we want to make onchain automation accessible to everyone, not just trading firms with million-dollar budgets. Whether you're looking to DCA into ETH, auto-compound your staking rewards, or bid on NFTs with specific traits, Plug will handle it all.\n\n## Join Us\n\nWe're building Plug in public because we believe the best products are shaped by their users. Start small – maybe set up a simple DCA strategy or automate your reward compounding. Tell us what works, what doesn't, and what you wish you could do.\n\nI've created a special access code for early readers: [abcd-1234-efgh-5678]. Use it at [app.onplug.io](https://app.onplug.io) to skip the waitlist and start building your first automation.\n\nFollow our journey:\n\n- Twitter: [[@onplug_io]](https://twitter.com/onplug_io)\n- Documentation: [docs.onplug.io](https://docs.onplug.io)\n\nSee you onchain,\nChance\n",attributes:{created:"2025-01-10T05:00:00.000Z",className:"",tags:["perspective"],author:"nftchance"}}};

        // * Get all the Posts for a given page.
        export const getPosts = (
            page = 1,
            pageSize = PAGE_SIZE,
            filter?:
                | Partial<Record<'date' | 'tag' | 'search', Partial<string>>>
                | undefined
        ): {
            posts: Post[]
            count: number
            hasNext: boolean
            random: Post
        } => {
            // ! Filter the posts before paginating so that we can get a final count.
            const filteredPosts = Object.values(
                Object.values(posts ?? {}).reduce<Record<string, Post>>(
                    (acc, article) => {
                        if(article.attributes.unlisted) return acc

                        if (!filter) {
                            acc[article.slug] = article

                            return acc
                        }

                        const { tag, date, search } = filter

                        // * Check the lowercase tags against the lowercase tag parameter.
                        const matchesTag =
                            !tag ||
                            (article.attributes.tags?.some(
                                articleTag =>
                                    articleTag.toLowerCase() === tag.toLowerCase()
                            ) ??
                                false)

                        // ! Require the exact date match.
                        const matchesDate =
                            !date ||
                            ((article.attributes.created?.includes(date) ||
                                article.attributes.updated?.includes(date)) ??
                                false)

                        const matchesSearch =
                            !search ||
                            article.title
                                ?.toLowerCase()
                                .includes(search.toLowerCase()) ||
                            article.description
                                ?.toLowerCase()
                                .includes(search.toLowerCase()) ||
                            article.content
                                ?.toLowerCase()
                                .includes(search.toLowerCase()) ||
                            (article.attributes.tags?.some(articleTag =>
                                articleTag.toLowerCase().includes(search.toLowerCase())
                            ) ??
                                false)

                        if (matchesTag && matchesDate && matchesSearch) {
                            acc[article.slug] = article
                        }

                        return acc
                    },
                    {}
                )
            )

            const count = filteredPosts.length
            const keys = Object.keys(posts)

            return {
                posts: filteredPosts.slice(
                    (page - 1) * pageSize,
                    page * pageSize
                ),
                count,
                hasNext: count > page * pageSize,
                random: posts[keys[Math.floor(Math.random() * keys.length)]]
            }
        }

        // * Get a specific Post by the value of the slug parameter in each Post.
        // ! The parameter can be the dictionary key or the slug which is a slug on the Post.
        export type PostLookupKey = keyof typeof posts
        export type PostLookup = (typeof posts)[PostLookupKey]['slug']

        export const getPost = (lookup: PostLookup) => {
            const article = posts[lookup.replaceAll('-', '') as PostLookupKey]

            if (!article) throw new Error('Post not found')

            return article
        }

        // * Get the favicon for a given URL.
        export const getFavicon = (url?: string) => {
            if (!url) return null

            const faviconUrl =
                faviconUrls[
                    url
                        .replace('https://', '')
                        .replace('http://', '') as keyof typeof faviconUrls
                ]

            if (!faviconUrl) return null

            return faviconUrl
        }
        