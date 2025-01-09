   /** @type {import('tailwindcss').Config} */
module.exports = {
   darkMode: 'class',

  content: [
      './assets/**/*.{css, scss}',
      './components/*.{vue,js}',
      './components/**/*.{vue,js}',
      './pages/*.vue',
      './pages/**/*.vue',
      './*.{vue,js,ts}',
      './nuxt.config.{js,ts}',
  ],
   
  theme: {
   extend: {
      colors: {
         
      },

      fontFamily: {
         
      },

      screens: {
         'sml': '480px',
         // => @media (min-width: 480px) { ... }

         'sm': '640px',
         // => @media (min-width: 640px) { ... }

         'hero': '966px',
         // => @media (min-width: 768px) { ... }

         'md': '768px',
         // => @media (min-width: 768px) { ... }

         'plg': '992px',
         // => @media (min-width: 768px) { ... }

         'lg': '1024px',
         // => @media (min-width: 1024px) { ... }

         'pxl': '1200px',
         // => @media (min-width: 1200px) { ... }

         'xl': '1280px',
         // => @media (min-width: 1280px) { ... }

         'p2xl': '1440px',
         // => @media (min-width: 1500px) { ... }

         '2xl': '1536px',
         // => @media (min-width: 1536px) { ... }

         '3xl': '1600px',
         // => @media (min-width: 1600px) { ... }

         'p4xl': '1600px',
         // => @media (min-width: 1600px) { ... }
      }
    }
  },
  plugins: [],
}