/** @type {import('tailwindcss').Config} */
const { fontFamily } = require('tailwindcss/defaultTheme')

module.exports = {
  content: [
    "./internal/views/**/*.templ",
    "./internal/components/**/*.templ",
],
  theme: {
    extend: {
      fontFamily: {
        lato: ['"Lato"', ...fontFamily.sans]
      },
      colors: {
        'codera-blue': {
          '50': '#f0f5fe',
          '100': '#dde8fc',
          '200': '#c3d8fa',
          '300': '#99bff7',
          '400': '#699df1',
          '500': '#467beb',
          '600': '#315cdf',
          '700': '#2849cd',
          '800': '#273da6',
          '900': '#273b8d',
          '950': '#1b2450',
        },
        'codera-red': {
          '50': '#fef2f3',
          '100': '#fee2e5',
          '200': '#fecacf',
          '300': '#fba6ae',
          '400': '#f7727e',
          '500': '#ee4555',
          '600': '#db2738',
          '700': '#be1e2d',
          '800': '#981c28',
          '900': '#7e1e27',
          '950': '#450a10',
        },
      }
    },
  },
  plugins: [],
};
