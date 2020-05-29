
export default {
  mode: 'universal',
  /*
  ** Headers of the page
  */
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff' },
  /*
  ** Global CSS
  */
  css: ["@/assets/css/main.scss"],
  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
    { src: '~/plugins/nacl.js', mode: 'client' },
    { src: '~/plugins/address.js', mode: 'client' },
  ],
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
    // Doc: https://github.com/nuxt-community/nuxt-tailwindcss
    '@nuxtjs/tailwindcss',
    '@nuxtjs/google-analytics'
  ],
  // Google Analytics
  googleAnalytics: {
    id: 'UA-145357881-5'
  },
  /*
  ** Nuxt.js modules
  */
  modules: [
    "@nuxtjs/axios",
    "nuxt-socket-io",
    '@nuxtjs/device',
  ],
  /*
  ** Socket.io
  */
  io: {
    sockets: [
      {
        name: 'natricon',
        url: 'wss://natricon.com',
        default: true
      }
    ],
  },
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  axios: {
    baseURL:
      process.env.NODE_ENV === 'production' && process.env.PLATFORM_TYPE === 'docker' ? 'https://natricon.com' : 'http://localhost:8080'
  },
  /*
  ** Build configuration
  */
  build: {
    /*
    ** You can extend webpack config here
    */
    extend(config, ctx) {
    }
  }
}
