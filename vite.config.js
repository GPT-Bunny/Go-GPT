import path from 'path';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    proxy: {
      // 将后端的 /send_message 请求代理到 Go 服务器的 /send_message 端点
      '/send_message': 'http://localhost:8080',
    },
  },
});
