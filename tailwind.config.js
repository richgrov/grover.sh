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
      0: "0",
      1: "1px",
    },
    extend: {
      keyframes: {
        "bg-shift": {
          "0%": { backgroundPosition: "0 0" },
          "100%": { backgroundPosition: "100% 0" },
        },
      },
      animation: {
        "bg-shift": "bg-shift 1.5s linear infinite",
      },
      backgroundImage: {
        rainbow: `linear-gradient(to right, darkorange, yellow, darkorange, yellow)`,
      },
    },
  },
  plugins: [],
};
