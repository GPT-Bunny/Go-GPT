<template>
  <!-- 使用 Naive UI 的卡片组件和表单组件 -->
  <n-card class="nav" title="登录">
    <n-form :model="form" label-position="top">
      <!-- 输入用户名的表单项 -->
      <n-form-item label="用户名" prop="username">
        <n-input v-model="form.username" placeholder="请输入用户名" />
      </n-form-item>

      <!-- 输入密码的表单项 -->
      <n-form-item label="密码" prop="password">
        <n-input v-model="form.password" type="password" placeholder="请输入密码" />
      </n-form-item>

      <!-- 登录按钮 -->
      <n-form-item>
        <n-button type="primary" @click="handleSubmit">登录</n-button>
      </n-form-item>
    </n-form>

    <!-- 显示错误消息的段落 -->
    <n-p v-if="errorMessage" style="color: red;">{{ errorMessage }}</n-p>
  </n-card>
</template>

<script setup>
import { ref } from 'vue';  // 引入 Vue 3 中的 ref 函数
import axios from 'axios';  // 引入 axios 库用于发送 HTTP 请求
import { useRouter } from 'vue-router';  // 引入 Vue Router 中的 useRouter 函数
import { NCard, NForm, NFormItem, NInput, NButton, NP } from 'naive-ui';  // 引入 Naive UI 组件

// 使用 useRouter 获取路由实例
const router = useRouter();

// 使用 ref 创建响应式数据
const form = ref({
  username: '',  // 初始化为空
  password: '',  // 初始化为空
});

const errorMessage = ref('');  // 错误消息

// 处理登录按钮点击事件
const handleSubmit = async () => {
  // 在这里发送登录请求
  try {
    // 使用 axios 发送 POST 请求
    const response = await axios.post('http://localhost:8080/login', {
      username: form.value.username,
      password: form.value.password,
    });

    // 输出请求和响应的信息
    console.log('Request:', {
      method: 'POST',
      url: 'http://localhost:8080/login',
      data: {
        username: form.value.username,
        password: form.value.password,
      },
    });

    console.log('Response:', response);

    // 处理登录成功
    if (response.data.success) {
      errorMessage.value = '';
      router.push({ name: 'Home' });  // 导航到 Home 页面
    } else {
      // 处理登录失败
      errorMessage.value = '登录失败，请检查用户名和密码';
    }
  } catch (error) {
    console.error('登录失败:', error);
    errorMessage.value = '登录失败，请检查用户名和密码';
  }
};

</script>
