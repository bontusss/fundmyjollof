// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
   compatibilityDate: '2024-11-01',
  
   devtools: { enabled: true },

   experimental: {
      cookieStore: false
   },

   app: {
      head: {
         title: "Fund My Jollof",

         htmlAttrs: {
            lang: "en",
         },
         meta: [
            { charset: 'utf-8' },
            { name: 'viewport', content: 'width=device-width, initial-scale=1' },
            { hid: 'description', name: 'description', content: '' },
            { hid: 'keywords', name: 'keywords', content: '' },
            { hid: 'theme-color', name: 'theme-color', content: '' },

            { hid: 'og:locale', property: 'og:locale', content: 'en_US' },
            { hid: 'og:type', property: 'og:type', content: 'website' },
            { hid: 'og:title', property: 'og:title', content: 'Fund My Jollof' },
            { hid: 'og:description', property: 'og:description', content: '' },
            { hid: 'og:url', property: 'og:url', content: 'https://fundmyjollof.com' },
            { hid: 'og:image', property: 'og:image', content: 'https://fundmyjollof.com/og-image.webp' },
            { hid: 'og:image:width', property: 'og:image:width', content: '1200' },
            { hid: 'og:image:height', property: 'og:image:height', content: '600' },
            { hid: 'og:image:alt', property: 'og:image:alt', content: 'Website Image Snippet' },
            { hid: 'og:site_name', property: 'og:site_name', content: 'Fund My Jollof' },
            { hid: 'robots', name: 'robots', content: 'follow, index' },

            { hid: 'twitter:card', name: 'twitter:card', content: 'summary_large_image' },
            { hid: 'twitter:title', name: 'twitter:title', content: 'Fund My Jollof' },
            { hid: 'twitter:description', name: 'twitter:description', content: '' },
            { hid: 'twitter:url', name: 'twitter:url', content: 'https://fundmyjollof.com' },
            { hid: 'twitter:image', name: 'twitter:image', content: 'https://fundmyjollof.com/og-image.webp' },
            { hid: 'twitter:site', name: 'twitter:site', content: '@fundmyjollof' },
            { hid: 'twitter:creator', name: 'twitter:creator', content: '@fundmyjollof' },
         ],

         link: [
            { rel: "icon", type: "image/svg+xml", href: "/favicon.svg" },
            { rel: "apple-touch-icon", type: "image/png", href: "/webclip.png" },
         ]
      },
   },

    components: [
      {
         path: "~/components",
         pathPrefix: false,
      },
   ],

   phosphor: {
      prefix: "Phos",
   },

   gtag: {
      id: process.env.NODE_ENV === 'development' ? '' : process.env.GTAG_ID
   },

   css: [
      "@/style/main.scss"
   ],

   imports: {
      dirs: [
         'composables/**',
         'utils/**'
      ]
   },

   pinia: {
      autoImports: [
         'defineStore', 
         ['defineStore', 'definePiniaStore']
      ],

      storesDirs: ['./store/**']
   },

   modules: [
      "@nuxtjs/tailwindcss",
      "nuxt-phosphor-icons",
      "nuxt-gtag",
      "@pinia/nuxt",
      "@nuxt/image",
      '@nuxt/content'
   ]

})
