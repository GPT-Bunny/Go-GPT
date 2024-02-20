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


.n-input{
  height: auto;
  min-height: 100px;
}
.n-button{
  height: 50px;
}
</style>