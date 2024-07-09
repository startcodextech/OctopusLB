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
            fontSize: {
                'title-xxl': ['44px', '55px'],
                'title-xl': ['36px', '45px'],
                'title-xl2': ['33px', '45px'],
                'title-lg': ['28px', '35px'],
                'title-md': ['24px', '30px'],
                'title-md2': ['26px', '30px'],
                'title-sm': ['20px', '26px'],
                'title-xsm': ['18px', '24px'],
            },
            spacing: {
                5.5: '1.375rem',
                7.5: '1.875rem',
                125: '31.25rem',
                62.5: '15.625rem',
            },
            maxHeight: {
                35: '8.75rem',
                70: '17.5rem',
                90: '22.5rem',
                550: '34.375rem',
                300: '18.75rem',
            },
            minWidth: {
                22.5: '5.625rem',
                42.5: '10.625rem',
                47.5: '11.875rem',
                75: '18.75rem',
            },
            zIndex: {
                999999: '999999',
                99999: '99999',
                9999: '9999',
                999: '999',
                99: '99',
                9: '9',
                1: '1',
            },
            opacity: {
                65: '.65',
            },
            transitionProperty: { width: 'width', stroke: 'stroke' },
            borderWidth: {
                6: '6px',
            },
            boxShadow: {
                default: '0px 8px 13px -3px rgba(0, 0, 0, 0.07)',
                card: '0px 1px 3px rgba(0, 0, 0, 0.12)',
                'card-2': '0px 1px 2px rgba(0, 0, 0, 0.05)',
                switcher:
                    '0px 2px 4px rgba(0, 0, 0, 0.2), inset 0px 2px 2px #FFFFFF, inset 0px -1px 1px rgba(0, 0, 0, 0.1)',
                'switch-1': '0px 0px 5px rgba(0, 0, 0, 0.15)',
                1: '0px 1px 3px rgba(0, 0, 0, 0.08)',
                2: '0px 1px 4px rgba(0, 0, 0, 0.12)',
                3: '0px 1px 5px rgba(0, 0, 0, 0.14)',
                4: '0px 4px 10px rgba(0, 0, 0, 0.12)',
                5: '0px 1px 1px rgba(0, 0, 0, 0.15)',
                6: '0px 3px 15px rgba(0, 0, 0, 0.1)',
                7: '-5px 0 0 #313D4A, 5px 0 0 #313D4A',
                8: '1px 0 0 #313D4A, -1px 0 0 #313D4A, 0 1px 0 #313D4A, 0 -1px 0 #313D4A, 0 3px 13px rgb(0 0 0 / 8%)',
            },
            dropShadow: {
                1: '0px 1px 0px #E2E8F0',
                2: '0px 1px 4px rgba(0, 0, 0, 0.12)',
            },
            colors: {
                primary: {
                    default: '#FA343E',
                    50: '#fef4f5',
                    100: '#fdced0',
                    200: '#fca7ac',
                    300: '#fb8187',
                    400: '#fa5a62',
                    500: '#FA343E',
                    600: '#c82931',
                    700: '#961f25',
                    800: '#641418',
                    900: '#310a0c',
                    light: {
                        50: '#fef7f7',
                        100: '#fed8db',
                        200: '#febabe',
                        300: '#fe9ba2',
                        400: '#fe7d86',
                        500: '#FE5F6A',
                        600: '#cb4c54',
                        700: '#98393f',
                        800: '#65262a',
                        900: '#321215',
                    },
                    dark: {
                        50: '#fdf4f4',
                        100: '#f6cbcc',
                        200: '#f0a2a4',
                        300: '#e97a7c',
                        400: '#e35154',
                        500: '#DD292D',
                        600: '#b02024',
                        700: '#84181b',
                        800: '#581012',
                        900: '#2c0808',
                    }
                },
                grey: {
                    50: '#fdfdfe',
                    100: '#f7f8fb',
                    200: '#f1f3f9',
                    300: '#ebeef6',
                    400: '#e5e9f4',
                    500: '#E0E5F2',
                    600: '#b3b7c1',
                    700: '#868991',
                    800: '#595b60',
                    900: '#2c2d30',
                    dark: {
                        50: '#f9fafb',
                        100: '#e4e7ee',
                        200: '#ced4e1',
                        300: '#b9c1d4',
                        400: '#a4aec7',
                        500: '#8F9BBA',
                        600: '#727c94',
                        700: '#555d6f',
                        800: '#393e4a',
                        900: '#1c1e25',
                    }
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
