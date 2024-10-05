"use strict";
/// <reference lib="webworker" />
Object.defineProperty(exports, "__esModule", { value: true });
var CACHE_NAME = 'plug-cache-v1';
var urlsToCache = [
    '/',
    '/index.html',
    '/styles/main.css',
    '/script/main.js',
];
self.addEventListener('install', function (event) {
    event.waitUntil(caches.open(CACHE_NAME)
        .then(function (cache) { return cache.addAll(urlsToCache); }));
});
self.addEventListener('fetch', function (event) {
    event.respondWith(caches.match(event.request)
        .then(function (response) { return response || fetch(event.request); }));
});
