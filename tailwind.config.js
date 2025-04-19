/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.{html,js}"],
    theme: {
        extend: {
            colors: {
                'blueGray': {
                    50: '#f8fafc',
                    100: '#f1f5f9',
                    200: '#e2e8f0',
                    300: '#cbd5e1',
                    400: '#94a3b8',
                    500: '#64748b',
                    600: '#475569',
                    700: '#334155',
                    800: '#1e293b',
                    900: '#0f172a',
                },
                'main': {
                    1: '#c0737a',
                    2: '#ffffff',
                },
                'side': {
                    1: '#3f3d56',
                    2: '#f0f0f0',
                },
            },
        },
    },
    plugins: [],
}