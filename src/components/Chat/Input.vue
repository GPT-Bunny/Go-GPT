<!-- InputBox.vue -->
<template>
  <div class="input-container"  style="">
    <n-input type="text" placeholder="可以清除" clearable />
    <n-button strong secondary type="warning" >
      发送
    </n-button>
  </div>
</template>


<script setup>
import { ref } from 'vue';

const inputValue = ref('');

const submit = () => {
  emit('submit', inputValue.value);
};
</script>

<style scoped>
.input-container {
  display: flex;
  align-items: center;
}

.input-container n-input {
  margin-right: 20px; /* 调整输入框与发送按钮之间的间距 */
}
</style>