const { addIconSelectors } = require("@iconify/tailwind");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["luxury"],
  },
  plugins: [require("daisyui"), addIconSelectors(["mdi"])],
};
