const dynamicNames = 'dynamic-v1-beta-15';
const staticPageNames = 'static-page-v1-beta-15';
const frequentlyUpdatableNames = 'frequently-updatable-v1-beta-15';
const libNames = 'lib-v1-beta-15';

// assets
const staticPageAssets = ['/', '/fallback', '/about', '/contact'];

const frequentlyUpdatableAssets = [
    '/favicon.ico',
    '/css/styles.min.css',
    '/css/index.min.css',
    '/css/about.min.css',
    '/js/app.min.js',
    '/js/fallback.min.js',
    '/js/index.min.js',
    '/images/icons/icon-72x72.png',
    '/images/about/welcome.png',
    '/images/about/background.jpg',
    '/images/home/welcome.webp'
];

const libAssets = [
    '/css/aos/aos.min.css',
    '/css/jquery.mCustomScrollbar.min.css',
    '/css/owlcarousel/owl.carousel.min.css',
    '/css/owlcarousel/owl.theme.default.min.css',
    '/css/bootstrap/bootstrap.min.css',
    '/css/fontawesome/fontawesome.min.css',
    '/css/fontawesome/brands.min.css',
    '/css/fontawesome/solid.min.css',
    '/css/webfonts/fa-brands-400.eot',
    '/css/webfonts/fa-brands-400.svg',
    '/css/webfonts/fa-brands-400.ttf',
    '/css/webfonts/fa-brands-400.woff',
    '/css/webfonts/fa-brands-400.woff2',
    '/css/webfonts/fa-solid-900.eot',
    '/css/webfonts/fa-solid-900.svg',
    '/css/webfonts/fa-solid-900.ttf',
    '/css/webfonts/fa-solid-900.woff',
    '/css/webfonts/fa-solid-900.woff2',
    '/js/jquery/jquery-3.3.1.slim.min.js',
    'js/popper/popper.min.js',
    '/js/bootstrap/bootstrap.min.js',
    '/js/jquery.mCustomScrollbar.concat.min.js',
    '/js/owlcarousel/owl.carousel.min.js',
    '/js/aos/aos.min.js',
    'https://fonts.googleapis.com/css?family=Roboto|Spectral|Timmana',
    'https://fonts.gstatic.com/s/timmana/v4/6xKvdShfL9yK-rvpOmzRKQ.woff2',
    'https://fonts.gstatic.com/s/spectral/v6/rnCr-xNNww_2s0amA9M5kng.woff2',
    'https://fonts.gstatic.com/s/roboto/v20/KFOmCnqEu92Fr1Mu4mxK.woff2'
];

// install event
self.addEventListener('install', evt => {
    // Install Event processing
    console.log('service worker installed');

    // Skip waiting on install
    self.skipWaiting();

    // libNames
    evt.waitUntil(
        caches.open(libNames).then(cache => {
            console.log(`caching static ${libNames}`);
            cache.addAll(libAssets);
        })
    );

    // staticPageNames
    evt.waitUntil(
        caches.open(staticPageNames).then(cache => {
            console.log(`caching static ${staticPageNames}`);
            cache.addAll(staticPageAssets);
        })
    );

    // frequentlyUpdatable
    evt.waitUntil(
        caches.open(frequentlyUpdatableNames).then(cache => {
            console.log(`caching static ${frequentlyUpdatableNames}`);
            cache.addAll(frequentlyUpdatableAssets);
        })
    );
});

// activate event
self.addEventListener('activate', evt => {
    // Claiming clients for current page
    evt.waitUntil(self.clients.claim());
    console.log('service worker activated');

    // Skip waiting on activate
    self.skipWaiting();

    evt.waitUntil(
        caches
            .keys()
            .then(keys =>
                Promise.all(
                    keys
                        .filter(
                            key =>
                                key !== staticPageNames &&
                                key !== frequentlyUpdatableNames &&
                                key !== libNames
                        )
                        .map(key => caches.delete(key))
                )
            )
    );
});

// fetch event
self.addEventListener('fetch', evt => {
    // check if request is made by chrome extensions or web page
    // if request is made for web page url must contains http.
    if (!(evt.request.url.indexOf('http') === 0)) return; // skip the request. if request is not made with http protocol

    evt.respondWith(
        caches
            .match(evt.request)
            .then(
                cacheRes =>
                    cacheRes ||
                    fetch(evt.request).then(fetchRes =>
                        caches.open(dynamicNames).then(cache => {
                            cache.put(evt.request.url, fetchRes.clone());
                            // check cached items size
                            limitCacheSize(dynamicNames, 15);
                            return fetchRes;
                        })
                    )
            )
            .catch(() => caches.match('/fallback'))
    );
});

// cache size limit function
const limitCacheSize = (name, size) => {
    caches.open(name).then(cache => {
        cache.keys().then(keys => {
            if (keys.length > size) {
                cache.delete(keys[0]).then(limitCacheSize(name, size));
            }
        });
    });
};