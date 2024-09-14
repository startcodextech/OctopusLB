import type { Config } from 'tailwindcss';
import defaultTheme from 'tailwindcss/defaultTheme';

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
    './src/modules/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  safelist: [],
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
          DEFAULT: '#8957FF',
          1: '#f7f0ff',
          2: '#e3d1ff',
          3: '#4a2bb3',
          dark: '#774ddc',
          'dark-1': '#f1ecfa',
          'dark-2': '#dfcef8',
          'dark-3': '#49327e',
          // Text color
          text: {
            DEFAULT: '#313131',
            dark: '#d6d7d7',
          },
        },
        // body
        base: {
          DEFAULT: '#F0F0F0',
          dark: '#232323',
        },
        // sidebar, navbar, bottom navigation, structural navigation elements
        frame: {
          DEFAULT: '#ffffff',
          dark: '#0A0F0F',
        },
        // Interactive elements such as buttons, inputs, selects and search bars
        action: {},
        // modal, alerts, floating elements that appear above the main content.
        overlay: {},
        // Cards or containers with high content, with shadows or borders to stand out from the background.
        layer: {},
        success: {
          DEFAULT: '#28CD41',
          1: '#f0fff0',
          2: '#d6ffd7',
          3: '#023310',
          dark: '#32D74B',
          'dark-1': '#f0fff0',
          'dark-2': '#deffde',
          'dark-3': '#043d14',
        },
        grey: {
          1: '#f3f3f3', // button background
          2: '#cacaca', // border
          3: '#bfbfbf', // text content card body
          4: '#afafaf', // button text
          5: '#929292', // text header
          6: '#313131', // text sidebar
          'dark-1': '#ffffff', // 1
          'dark-2': '#858585', // 2
          'dark-3': '#4a4b4b', // 3
          'dark-4': '#585a5a', // 4
          'dark-5': '#adadad', // 5
          'dark-6': '#616363', // 6
        },
        warning: {
          DEFAULT: '#ff9500',
          1: '#fff9e6',
          2: '#ffe5a3',
          3: '#662e00',
          dark: '#ff9f0a',
          'dark-1': '#fff9e6',
          'dark-2': '#ffe9ad',
          'dark-3': '#663000',
        },
        error: {
          DEFAULT: '#ff3b30',
          1: '#fff3f0',
          2: '#ffdcd4',
          3: '#66030b',
          dark: '#ff453a',
          'dark-1': '#fff3f0',
          'dark-2': '#ffe4de',
          'dark-3': '#66070f',
        },
        info: {
          DEFAULT: '#007aff',
          1: '#e6f6ff',
          2: '#a3dcff',
          3: '#002466',
          dark: '#0a84ff',
          'dark-1': '#e6f7ff',
          'dark-2': '#ade1ff',
          'dark-3': '#002566',
        },
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
