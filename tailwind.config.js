/*
 ** TailwindCSS Configuration File
 **
 ** Docs: https://tailwindcss.com/docs/configuration
 ** Default: https://github.com/tailwindcss/tailwindcss/blob/master/stubs/defaultConfig.stub.js
 */
module.exports = {
  theme: {
    colors: {
      white: "#FFFFFF",
      black: "#000000",
      cyan: "#66FFFF",
      redPink: "#FE98CD",
      green: "#66FFB2",
      lightBlue: "#A3CDFF",
      lightPink: "#F199FF",
      yellow: "#FFEE52",
      lightOrange: "#FFC89F",
      aquaGreen: "#7BFFC8",
      brightPink: "#FFA4F6",
      lime: "#D4FF00",
      red: "#FF4646",
      lightGreen: "#B5FFD8",
      transparent: "transparent"
    },
    extend: {
      spacing: {
        '72': '18rem',
        '84': '21rem',
        '96': '24rem',
        '108': '27rem',
        '120': '30rem',
        '128': '33rem',
        '144': '39rem',
      },
      opacity: {
        '5': '.05',
        '8': '.08',
        '10': '.1',
        '15': '.15'
      },
      padding: {
        '0_5': '0.125rem',
        '1_5': '0.375rem'
      },
      borderRadius: {
        'xl': "1rem"
      },
      borderWidth: {
        '2': '2px',
        '3': '3px',
        '4': '4px'
      },
      fontSize: {
        'xxs': '0.5rem',
        'xxxs': '0.25rem',
      }
    }
  },
  variants: {},
  plugins: [],
  purge: {
    // Learn more on https://tailwindcss.com/docs/controlling-file-size/#removing-unused-css
    enabled: process.env.NODE_ENV === 'production',
    content: [
      'components/**/*.vue',
      'layouts/**/*.vue',
      'pages/**/*.vue',
      'plugins/**/*.js',
      'nuxt.config.js'
    ]
  }
};
