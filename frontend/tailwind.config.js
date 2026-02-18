/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        brand: {
          yellow: '#FDFD96',
          'yellow-light': '#FEFECB',
          'yellow-pale': '#FFFDE7',
          blue: '#87CEEB',
          'blue-dark': '#5BADE6',
          gray: '#E0E0E0',
          'gray-light': '#F5F5F5',
          black: '#1A1A1A',
        },
      },
      fontFamily: {
        sans: ['Inter', 'system-ui', 'sans-serif'],
        cursive: ['"Dancing Script"', 'cursive'],
        handwritten: ['"Caveat"', 'cursive'],
      },
      borderRadius: {
        '2xl': '1rem',
        '3xl': '1.5rem',
        '4xl': '2rem',
      },
    },
  },
  plugins: [],
}
