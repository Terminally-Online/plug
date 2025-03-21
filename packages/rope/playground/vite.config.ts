import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@plug/rope': path.resolve(__dirname, '../src'),
    },
  },
  optimizeDeps: {
    include: ['react', 'react-dom', 'reactflow'],
    exclude: ['@plug/rope']
  },
  server: {
    port: 5173,
    open: true,
    proxy: {
      '/solver': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false
      }
    }
  },
})