<template>
  <nav class="nav">
    <router-link to="/home" class="logo-container">
      <img src="@/assets/32x32.png" alt="Logo" class="logo">
      <p class="logo-text">BUNNY</p>
    </router-link>

    <router-link to="/home" >
      <n-button strong secondary round type="warning">
      <i class="fas fa-home" > 首页</i>
    </n-button>
    </router-link>

    <router-link to="/Chat" >
      <n-button strong secondary round type="warning">
        <i class="fas fa-comment-alt" > 聊天</i>
      </n-button>
    </router-link>

    <router-link to="/draw" >
      <n-button strong secondary round type="warning">
        <i class="fas fa-palette" > 绘画</i>
      </n-button>
    </router-link>

    <router-link to="/admin" >
      <n-button strong secondary round type="warning">
        <i class="fas fa-wrench" > 管理</i>
      </n-button>
    </router-link>

      <router-link to="/login">
          <auth-button   />
      </router-link>
  </nav>
</template>
<script setup>
import AuthButton from '@/components/AuthButton.vue';
</script>
<style  scoped>
@import url('Nav.css');
</style>
