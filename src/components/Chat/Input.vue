<template>
  <div>
    <!-- 消息输入框 -->
    <input v-model="userInput" @keyup.enter="sendMessage" placeholder="输入消息..." />
    <!-- 发送按钮 -->
    <button @click="sendMessage">发送</button>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const userInput = ref(''); // 用户输入的消息

const sendMessage = () => {
  // 发送消息的逻辑
};
</script>
