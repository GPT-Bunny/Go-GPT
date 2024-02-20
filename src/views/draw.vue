<template>
    <n-config-provider :theme="theme">
      <n-card>
        <n-space>
          <n-button @click="toggleTheme(darkTheme)">深色</n-button>
          <n-button @click="toggleTheme(null)">浅色</n-button>
        </n-space>
      </n-card>
    </n-config-provider>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { darkTheme } from 'naive-ui'
  
  const theme = ref(null)
  
  const toggleTheme = (newTheme) => {
    theme.value = newTheme
  }
  </script>
  
  <style scoped>
  .n-space {
    padding-top: 200px;
  }
  
  .n-button {
    padding-top: 200px;
  }
  </style>
  