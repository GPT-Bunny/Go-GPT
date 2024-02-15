<!-- Vue 3 template code -->
<template>
  <div>
    <!-- Your existing Vue code here -->

    <div  style="padding-top: 100px;">
      <input v-model="message" placeholder="Type your message">
      <button @click="sendMessage">Send</button>
      <div v-if="chatResponse">{{ chatResponse }}</div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';

const message = ref('');
const chatResponse = ref('');

const sendMessage = async () => {
  try {
    // 请替换为你的后端API地址
    const response = await axios.post('http://localhost:8080/chat', {
      messages: [
        {
          role: 'user',
          content: message.value,
        },
      ],
    });

    chatResponse.value = response.data.reply;
  } catch (error) {
    console.error('Failed to communicate with the server', error);
  }
};
</script>

