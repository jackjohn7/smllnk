const colors = require('tailwindcss/colors');

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './public/views/**/*.templ',
  ],
  theme: {
    extend: {
      colors: {
        'bistre': '#231006',
        'dutch-white': '#EFE4C5',
        'kombu-green': '#243010',
        'polished-pine': '#539987',
        'midnight-green': '#1A535C',
        'pink-accent': '#e15ad9',
        'primary-green': '#72e15a',
        'base-dark': '#121619',
        'base-up-one': '#1f272f',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ]
}
