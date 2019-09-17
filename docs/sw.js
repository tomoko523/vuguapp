// Cache name
const CACHE_NAME = 'pwa-sample-caches-v1';
// Cache targets
const urlsToCache = [
    '.',
    'https://cdn.jsdelivr.net/npm/text-encoding@0.7.0/lib/encoding.min.js',
    'https://cdnjs.cloudflare.com/ajax/libs/galleriffic/2.0.1/css/loader.gif'
];

self.addEventListener('install', (event) => {
  event.waitUntil(
    caches
      .open(CACHE_NAME)
      .then((cache) => {
        return cache.addAll(urlsToCache);
      })
  );
});

self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches
      .match(event.request)
      .then((response) => {
        return response ? response : fetch(event.request);
      })
  );
});