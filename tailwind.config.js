const { fontFamily } = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["content/*.html", "layouts/**/*.html"],
  theme: {
    fontFamily: {
      sans: ["Inter", ...fontFamily.sans],
      pixel: ["Unispace", ...fontFamily.mono],
    },
    borderWidth: {
      1: "1px",
    },
    extend: {},
  },
  plugins: [],
};
