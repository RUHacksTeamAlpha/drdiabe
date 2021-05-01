import colors from 'vuetify/es5/util/colors'

module.exports = {
  // Global page headers: https://go.nuxtjs.dev/config-head

  // srcDir: '.',
  // buildDir: 'functions/.nuxt',
  head: {
    titleTemplate: '%s - drdiabetech',
    title: 'drdiabe.tech',
    htmlAttrs: {
      lang: 'en',
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    { src: '~/plugins/vue-scroll-reveal.js', ssr: false },
    { src: '~/plugins/vue-apexcharts.js', ssr: false },
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    '@nuxtjs/eslint-module',
    // https://go.nuxtjs.dev/vuetify
    '@nuxtjs/vuetify',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    [
      '@nuxtjs/firebase',
      {
        config: {
          apiKey: 'AIzaSyAUzqt2lVNOULpyLdVsF-73f5ozXhkCc3o',
          authDomain: 'diabetech-515ed.firebaseapp.com',
          projectId: 'diabetech-515ed',
          storageBucket: 'diabetech-515ed.appspot.com',
          messagingSenderId: '209540750751',
          appId: '1:209540750751:web:5b2ff265751d32a3db31aa',
          measurementId: 'G-YHVVFGH5RC',
        },
        services: {
          auth: true, // Just as example. Can be any other service.
        },
      },
    ],
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {},

  // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: true,
      themes: {
        dark: {
          primary: '#345995',
          accent: '#60AFFF',
          secondary: '#121212',
          info: '#A29587',
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3,
        },
      },
    },
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    // extractCSS: true,
  },

  //Custom by me, Allen!
  // middleware: ["assests/auth"],
  auth: {
    persistence: 'local', // default
    initialize: {
      onAuthStateChangedMutation: 'ON_AUTH_STATE_CHANGED_MUTATION',
      onAuthStateChangedAction: 'onAuthStateChangedAction',
      subscribeManually: false
    },
    ssr: false, // default
    emulatorPort: 9099,
    emulatorHost: 'http://localhost',
  }
}
