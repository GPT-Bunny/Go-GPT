<script setup>
import {  RouterView } from 'vue-router'
import { ref } from 'vue';

// 在这里设置您想要的默认主题
const theme = ref('light');
</script>

<template>
  
    <n-message-provider>
      <RouterView />
    </n-message-provider>
</template>



  