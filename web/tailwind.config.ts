import type { Config } from "tailwindcss";
import defaultTheme from "tailwindcss/defaultTheme";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    fontFamily: {
      redhat: ['Red Hat Display', 'sans-serif'],
    },
    extend: {
      screens: {
        '2xsm': '375px',
        xsm: '425px',
        '3xl': '2000px',
        ...defaultTheme.screens,
      },
      fontFamily: {
        redhat: ['Red Hat Display', 'sans-serif'],
      },
      colors: {
        transparent: 'transparent',
        current: 'currentColor',
        primary: {
          50: '#f6f3fc',
          100: '#d5c8f2',
          200: '#b49de9',
          300: '#9372df',
          400: '#7247d5',
          500: '#511CCC',
          600: '#4016a3',
          700: '#30107a',
          800: '#200b51',
          900: '#100528',
        },
        sidebar: '#f6f7fa',
        'sidebar-text': '#A4A4AD',
        'sidebar-icon': '#767783',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
    require('@tailwindcss/aspect-ratio'),
    require('@tailwindcss/line-clamp'),
  ],
};
export default config;
