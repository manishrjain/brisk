/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx,svelte}",
  ],
  theme: {
    extend: {
      colors: {
        monokai: {
          pink: '#FF6188',
          orange: '#FC9867',
          cyan: '#78DCE8',
          green: '#A9DC76',
          yellow: '#FFD866',
          purple: '#AB9DF2',
          bg: '#000000',
          'bg-light': '#1a1a1a',
          text: '#FCFCFA',
          'text-muted': '#939293',
          'text-dim': '#5c5c5c',
          border: '#2d2d2d',
        },
      },
    },
  },
  plugins: [],
  darkMode: 'class',
}
