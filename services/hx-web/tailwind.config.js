/** @type {import('tailwindcss').Config} */

const { fontSize } = require('./src/app/theme/fontStyles')

module.exports = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    fontFamily: {
      display: ['Inter', 'system-ui', 'sans-serif'],
      body: ['Inter', 'system-ui', 'sans-serif'],
    },
    fontSize,
    extend: {
      boxShadow: { 
        'card': 'inset 0 1px 0 0 #ffffff0d'
      },
      backgroundImage: {
        hero: "url('/league-art.png')",
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      colors: {
        primary: "#0D1321",
      },
    },
  },
  plugins: [],
};
