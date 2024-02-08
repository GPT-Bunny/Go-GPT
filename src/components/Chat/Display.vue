<template>
    <div>
      <!-- 展示选中对话的消息 -->
      <div v-for="(message, index) in messages" :key="index">
        {{ message.text }}
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const messages = ref([]); // 当前对话的消息列表
  </script>
  