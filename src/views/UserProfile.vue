<template>
    <div class="user-profile">
      <h1>个人资料</h1>
      <div>
        <img :src="user.avatar" alt="用户头像" class="profile-avatar" />
        <p>用户名: {{ user.name }}</p>
        <p>邮箱: {{ user.email }}</p>
        <!-- 其他用户信息 -->
      </div>
      <!-- 编辑资料的按钮或链接 -->
      <n-button @click="editProfile">编辑资料</n-button>
    </div>
  </template>
  
  <script setup>
  import { computed } from 'vue';
  import { useStore } from 'vuex';
  import { useRouter } from 'vue-router';
  
  const store = useStore();
  const user = computed(() => store.state.user);
  
  const editProfile = () => {
    // 实现编辑资料的逻辑，可能是导航到一个带有表单的编辑页面
  };
  </script>
  
  <style>
  .profile-avatar {
    width: 100px;
    height: 100px;
    border-radius: 50%;
  }
  /* 其他样式 */
  </style>
  