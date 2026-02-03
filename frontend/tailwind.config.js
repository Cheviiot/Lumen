/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        // Премиум акцент - глубокий индиго
        'accent': {
          50: '#eef2ff',
          100: '#e0e7ff',
          200: '#c7d2fe',
          300: '#a5b4fc',
          400: '#818cf8',
          500: '#6366f1',
          600: '#4f46e5',
          700: '#4338ca',
          800: '#3730a3',
          900: '#312e81',
          950: '#1e1b4b',
        },
        // Вторичный - изумрудный
        'secondary': {
          50: '#ecfdf5',
          100: '#d1fae5',
          200: '#a7f3d0',
          300: '#6ee7b7',
          400: '#34d399',
          500: '#10b981',
          600: '#059669',
          700: '#047857',
          800: '#065f46',
          900: '#064e3b',
          950: '#022c22',
        },
        // Чистые акценты
        'brand': {
          blue: '#3b82f6',
          purple: '#8b5cf6',
          pink: '#ec4899',
          emerald: '#10b981',
          amber: '#f59e0b',
        },
        // Ультра-темная палитра для 4K четкости
        'dark': {
          50: '#fafafa',
          100: '#f5f5f5',
          200: '#e4e4e7',
          300: '#d4d4d8',
          400: '#a1a1aa',
          500: '#71717a',
          600: '#52525b',
          700: '#3f3f46',
          800: '#27272a',
          850: '#1f1f23',
          900: '#18181b',
          925: '#131316',
          950: '#09090b',
        },
        // Дополнительные акценты
        'success': {
          DEFAULT: '#10b981',
          light: '#34d399',
          dark: '#059669',
        },
        'warning': {
          DEFAULT: '#f59e0b',
          light: '#fbbf24',
          dark: '#d97706',
        },
        'error': {
          DEFAULT: '#ef4444',
          light: '#f87171',
          dark: '#dc2626',
        },
        'info': {
          DEFAULT: '#3b82f6',
          light: '#60a5fa',
          dark: '#2563eb',
        }
      },
      fontFamily: {
        'sans': ['Inter', 'system-ui', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'Roboto', 'sans-serif'],
        'display': ['Poppins', 'Inter', 'sans-serif'],
      },
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-conic': 'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
        'gradient-shine': 'linear-gradient(110deg, transparent 40%, rgba(255,255,255,0.15) 50%, transparent 60%)',
        'gradient-neon': 'linear-gradient(135deg, #667eea 0%, #764ba2 25%, #f093fb 50%, #4facfe 75%, #00f2fe 100%)',
        'gradient-rainbow': 'linear-gradient(90deg, #ff006e, #8338ec, #3a86ff, #06ffa5, #ffbe0b, #ff006e)',
        'gradient-cyber': 'linear-gradient(135deg, #7c3aed 0%, #06b6d4 50%, #8338ec 100%)',
        'mesh-gradient': 'radial-gradient(at 27% 37%, hsla(271, 87%, 56%, 0.4) 0, transparent 50%), radial-gradient(at 97% 21%, hsla(180, 87%, 56%, 0.4) 0, transparent 50%), radial-gradient(at 52% 99%, hsla(271, 87%, 56%, 0.4) 0, transparent 50%), radial-gradient(at 10% 29%, hsla(180, 87%, 56%, 0.4) 0, transparent 50%), radial-gradient(at 97% 96%, hsla(271, 87%, 56%, 0.4) 0, transparent 50%)',
      },
      boxShadow: {
        'glow': '0 0 30px rgba(124, 58, 237, 0.5), 0 0 60px rgba(124, 58, 237, 0.3)',
        'glow-lg': '0 0 50px rgba(124, 58, 237, 0.6), 0 0 100px rgba(124, 58, 237, 0.4)',
        'glow-accent': '0 0 30px rgba(6, 182, 212, 0.5), 0 0 60px rgba(6, 182, 212, 0.3)',
        'glow-pink': '0 0 30px rgba(255, 0, 110, 0.5), 0 0 60px rgba(255, 0, 110, 0.3)',
        'inner-glow': 'inset 0 0 30px rgba(124, 58, 237, 0.3)',
        'neon': '0 0 5px rgba(124, 58, 237, 0.8), 0 0 20px rgba(124, 58, 237, 0.6), 0 0 40px rgba(124, 58, 237, 0.4)',
        'neon-cyan': '0 0 5px rgba(6, 182, 212, 0.8), 0 0 20px rgba(6, 182, 212, 0.6), 0 0 40px rgba(6, 182, 212, 0.4)',
      },
      animation: {
        'fade-in': 'fadeIn 0.5s ease-out',
        'fade-in-up': 'fadeInUp 0.6s ease-out',
        'fade-in-down': 'fadeInDown 0.6s ease-out',
        'slide-in-left': 'slideInLeft 0.4s ease-out',
        'slide-in-right': 'slideInRight 0.4s ease-out',
        'slide-up': 'slideUp 0.5s ease-out',
        'scale-in': 'scaleIn 0.4s ease-out',
        'bounce-slow': 'bounce 3s ease-in-out infinite',
        'pulse-glow': 'pulseGlow 2s ease-in-out infinite',
        'shimmer': 'shimmer 2s linear infinite',
        'float': 'float 6s ease-in-out infinite',
        'spin-slow': 'spin 8s linear infinite',
        'wiggle': 'wiggle 1s ease-in-out infinite',
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
        fadeInUp: {
          '0%': { opacity: '0', transform: 'translateY(20px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        fadeInDown: {
          '0%': { opacity: '0', transform: 'translateY(-20px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        slideInLeft: {
          '0%': { opacity: '0', transform: 'translateX(-30px)' },
          '100%': { opacity: '1', transform: 'translateX(0)' },
        },
        slideInRight: {
          '0%': { opacity: '0', transform: 'translateX(30px)' },
          '100%': { opacity: '1', transform: 'translateX(0)' },
        },
        slideUp: {
          '0%': { opacity: '0', transform: 'translateY(30px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        scaleIn: {
          '0%': { opacity: '0', transform: 'scale(0.9)' },
          '100%': { opacity: '1', transform: 'scale(1)' },
        },
        pulseGlow: {
          '0%, 100%': { boxShadow: '0 0 30px rgba(124, 58, 237, 0.5), 0 0 60px rgba(124, 58, 237, 0.3)' },
          '50%': { boxShadow: '0 0 50px rgba(124, 58, 237, 0.8), 0 0 100px rgba(124, 58, 237, 0.5)' },
        },
        shimmer: {
          '0%': { backgroundPosition: '-200% 0' },
          '100%': { backgroundPosition: '200% 0' },
        },
        float: {
          '0%, 100%': { transform: 'translateY(0px)' },
          '50%': { transform: 'translateY(-20px)' },
        },
        wiggle: {
          '0%, 100%': { transform: 'rotate(-3deg)' },
          '50%': { transform: 'rotate(3deg)' },
        },
      },
      transitionTimingFunction: {
        'smooth': 'cubic-bezier(0.4, 0, 0.2, 1)',
        'bounce-in': 'cubic-bezier(0.68, -0.55, 0.265, 1.55)',
      },
      backdropBlur: {
        xs: '2px',
      },
    },
  },
  plugins: [],
}
