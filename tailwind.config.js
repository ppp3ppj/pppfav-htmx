import { screens as _screens } from 'tailwindcss/defaultTheme';
import plugin from 'tailwindcss/plugin';

export const darkMode = 'class';
// example content path templ and *js
export const content = ["./views/**/*.templ", "./static/**/*.js"];
export const theme = {
    screens: {
      'xs': '320px',
      'sm': '640px',
      'md': '768px',
      'itn': '880px',
      'lg': '1024px',
      'itn2':'1150px',
      'xl': '1280px',
      '2xl': '1536px',
      '3xl': '1920px',
      '4xl': '2560px',
    },
    extend: {
        fontFamily: {
            Kanit: ['Kanit', 'sans-serif']
        },
    },
};
export const plugins = [
    plugin(function({ addBase }) {
        addBase({
            'html': {
                fontSize: "16px", // Default font size, which is typically 16px
                '@screen md': {
                    fontSize: "14px", // Font size on medium (md) breakpoint, which is typically 14px
                },
            },
        });
    }),
    require('@tailwindcss/typography'),
    require('daisyui')
];
export const daisyui = {
    themes: [
        "cupcake",
        "sunset",
        "cyberpunk",
        "nord",
        "synthwave",
        "valentine",
        "dracula",
        "dim",
    ],
    darkTheme: "sunset",
};
