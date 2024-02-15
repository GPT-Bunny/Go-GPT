<template>
    <div>
      <div class="form-group">
        <label for="title">标题</label>
        <input type="text" id="title" v-model="title" />
      </div>
      <div class="form-group">
        <label for="subtitle">副标题</label>
        <input type="text" id="subtitle" v-model="subtitle" />
      </div>
      <div class="form-group">
        <label for="image">图片</label>
        <input type="file" id="image" @change="onFileChange" />
        <img :src="imageUrl" alt="Logo Preview" />
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const title = ref('BUNNY');
  const subtitle = ref('一款好用的AI助理');
  const imageUrl = ref(require('@/assets/32x32.png'));
  
  const onFileChange = (event) => {
    const file = event.target.files[0];
    imageUrl.value = URL.createObjectURL(file);
  };
  </script>
  
  <style>
  .form-group {
    margin-bottom: 20px;
  }
  </style>
  