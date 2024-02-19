import { defineStore } from 'pinia';
import axios from 'axios'; // 假设你使用axios来发送HTTP请求

export const useUserStore = defineStore('me', {
  state: () => ({
    isLoggedIn: false,
    user: null,
  }),
  getters: {
    // getters go here
  },
  actions: {
    async login(userCredentials) {
      try {
        // 发送登录请求到后端API
        const response = await axios.post('/api/login', userCredentials);
        
        // 假设后端返回的响应体包含用户信息和登录状态
        if (response.data.success) {
          this.isLoggedIn = true;
          this.user = response.data.user;
          return true; // 登录成功
        } else {
          // 登录失败的逻辑处理
          return false; // 登录失败
        }
      } catch (error) {
        console.error('登录请求失败:', error);
        return false; // 登录失败
      }
    },
  },
});
