<!-- ChatList.vue -->
<template>
  <div class="list">
    <n-button strong secondary type="warning" style="width: 100%;">
      新建对话
    </n-button>
    <ul>
      <n-divider />
      <li v-for="(chat, index) in chatList" :key="index" @click="selectChat(chat)">
        {{ chat.title }}
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const chatList = ref([
  { id: 1, title: 'Chat 1' },
  { id: 2, title: 'Chat 2' },
  { id: 3, title: 'Chat 3' },
]);

const selectChat = (chat) => {
  console.log('Selected chat:', chat);
  // Handle the selection logic here
};
</script>

<style scoped>
/* Your CSS styles here */
</style>
