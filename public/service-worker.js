/// <reference lib="webworker" />
const CACHE_NAME = 'plug-cache-v1';
const urlsToCache = [
    '/',
    '/index.html',
    '/styles/main.css',
    '/script/main.js',
];
self.addEventListener('install', (event) => {
    event.waitUntil(caches.open(CACHE_NAME)
        .then((cache) => cache.addAll(urlsToCache)));
});
self.addEventListener('fetch', (event) => {
    event.respondWith(caches.match(event.request)
        .then((response) => response || fetch(event.request)));
});
export {};
