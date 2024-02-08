<template>
    <div>
      <!-- 对话列表 -->
      <div v-for="(conversation, index) in conversations" :key="index" @click="selectConversation(index)">
        Conversation {{ index + 1 }}
      </div>
      <!-- 新建对话按钮 -->
      <button @click="addConversation">新建对话</button>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const conversations = ref([]); // 对话列表数据
  
  const addConversation = () => {
    // 添加新对话的逻辑
  };
  
  const selectConversation = (index) => {
    // 选中对话的逻辑
  };
  </script>
  