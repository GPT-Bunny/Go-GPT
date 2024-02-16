<!-- Vue 3 template code -->
<template>
      <n-layout-sider
        collapse-mode="transform"
        :collapsed-width="120"
        :width="240"
        show-trigger="bar"
        content-style="padding: 24px;"
        bordered
      >
      </n-layout-sider>

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

<style>
.n-layout-sider{
  padding-top: 100px;
}
</style>