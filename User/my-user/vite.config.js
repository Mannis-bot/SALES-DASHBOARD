import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'  // Add this line to import path module

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    hmr: true,  // Enable Hot Module Replacement
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'), // Set '@' to resolve to 'src' directory
    },
  },
})

