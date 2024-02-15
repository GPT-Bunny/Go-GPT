<!-- Layout.vue -->

<template>
  <div>
    <Nav :options="options" />
    <router-view />

  </div>
</template>

<script setup>
import Nav from '@/views/Nav.vue';
import { ref } from 'vue';

// 在这里添加 options 属性
const options = ref([
  // 你的下拉选项，可能需要从外部传递或者根据需要定义
]);

// 其他的 import 和 setup 逻辑
</script>
