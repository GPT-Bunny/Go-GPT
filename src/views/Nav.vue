<template>
    <nav class="nav">
      <div class="logo-container">
        <img src="@/assets/32x32.png" alt="Logo" class="logo">
        <p class="logo-text">BUNNY</p>
      </div>
      <router-link to="/home" class="nav-link">首页</router-link>
      <router-link to="/Chat" class="nav-link"><i class="fa fa-comment-alt"></i><n-dropdown :options="options">聊天</n-dropdown></router-link>
      <router-link to="/draw" class="nav-link">绘画</router-link>
      <router-link to="/chat" class="nav-link login-button"> 登录</router-link>
    </nav>


  </template>
  
  <script setup>
  import { NIcon, NButton, NDropdown } from "naive-ui";
  
  // 在这里添加 options 属性
  const options = [
    // 你的下拉选项，可能需要从外部传递或者根据需要定义
  ];
  
  // 其他的 import 和 setup 逻辑
  </script>
<style>
@import url('Nav.css');
</style>
