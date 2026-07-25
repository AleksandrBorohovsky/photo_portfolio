/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./web/templates/**/*.html'],
    theme: {
        extend: {
            fontFamily: {
                sans: ['Inter', 'sans-serif'],
                heading: ['Playfair Display', 'serif'],
            },
        },
    },
    plugins: [],
};