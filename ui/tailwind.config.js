/** @type {import('tailwindcss').Config} */
module.exports = {
  /*content: [
    `./src/pages/!**!/!*.{js,jsx,ts,tsx}`,
    `./src/components/!**!/!*.{js,jsx,ts,tsx}`,
    `./src/modules/!**!/!*.{js,jsx,ts,tsx}`,
  ],*/
  content: [
      "./gatsby-browser.{js,jsx,ts,tsx}",
      "./src/**/*.{js,jsx,ts,tsx}"
  ],
  theme: {
    extend: {
        fontFamily: {
            // redhat: ['Red Hat Display', 'sans-serif'],
        },
        colors: {
            'purple-blue': {
                100: 'rgb(234, 232, 255)',
                300: 'rgb(185, 176, 255)',
                500: 'rgb(89, 31, 249)',
                600: 'rgb(85, 28, 229)',
                700: 'rgb(72, 23, 192)',
            },
            grey: {
                50: '#f9fafb',
                100: 'rgb(250, 252, 254)',
                200: 'rgb(246, 248, 253)',
                300: 'rgb(244, 247, 254)',
                400: 'rgb(233, 237, 247)',
                500: 'rgb(224, 229, 242)',
                600: 'rgb(163, 174, 208)',
                700: 'rgb(112, 126, 174)',
                800: 'rgb(71, 84, 140)',
                900: 'rgb(43, 54, 116)',
                dark: {
                    100: 'rgb(239, 244, 251)',
                    300: 'rgb(201, 212, 234)',
                    500: 'rgb(143, 155, 186)',
                    600: 'rgb(104, 118, 159)',
                    700: 'rgb(72, 85, 133)',
                    800: 'rgb(45, 57, 107)',
                    900: 'rgb(27, 37, 89)',
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
