import type { Config } from "tailwindcss";
import defaultTheme from "tailwindcss/defaultTheme";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/modules/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  safelist: [
    {
      pattern: /bg-chip-(success|warning|error)-(background|dot|text)/,
    },
    {
      pattern: /text-chip-(success|warning|error)-(background|dot|text)/
    },
    {
      pattern: /bg-primary-(50|100|200|300|400|500|600|700|800|900)/,
    },
    {
      pattern: /text-primary-(50|100|200|300|400|500|600|700|800|900)/
    }
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
        sidebar: {
          background: '#F6F7FA',
          icon: '#767783',
          text: '#A4A4AD',
        },
        neutral: {
          background: '#F0F0F6',
        },
        text: {
          primary: '#080D30',
          secondary: '#5C5C74',
        },
        chip: {
          success: {
            background: '#CEEFDFFF',
            text: '#458850FF',
            dot: '#34C759FF',
          },
          warning: {
            background: '#FFF4B3FF', // Fondo warning
            text: '#886D1CFF',        // Texto warning
            dot: '#FFC734FF',         // Punto warning
          },
          error: {
            background: '#FFDDDDFF', // Fondo error
            text: '#9F2F2FFF',         // Texto error
            dot: '#FF4545FF',          // Punto error
          },
        }
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
