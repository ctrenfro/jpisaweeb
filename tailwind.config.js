/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "views/**/*.templ",
    "./views/**/*.templ",
    "./**/*.templ",
    "./views/home/*.{templ, html, go}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};
