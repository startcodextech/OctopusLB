const defaultTheme = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./gatsby-browser.{js,jsx,ts,tsx}",
        "./src/**/*.{js,jsx,ts,tsx}"
    ],
    theme: {
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
                primary: {
                    DEFAULT: '#511CCC',
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
            },
        },
    },
    variants: {
        extend: {
            animation: ['motion-safe'],
        },
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
        require('@tailwindcss/aspect-ratio'),
        require('@tailwindcss/line-clamp'),
    ],
}
