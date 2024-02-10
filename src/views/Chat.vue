<template>
  <div  class="div">
  <n-flex justify="space-between" >
    <!-- 侧边栏 -->
    <n-layout has-sider>
      <n-layout-sider 
        collapse-mode="transform"
        :collapsed-width="20"
        :width="250"
        show-trigger="arrow-circle"
        content-style="padding: 24px;"
        bordered
      >
        <List />
      </n-layout-sider>

      <!-- 对话内容 -->
      <n-layout-content class="content">
        <Title />
        <n-divider />
        <Display />
        <n-divider />
        <Input />
      </n-layout-content>
    </n-layout>
  </n-flex>
</div>
</template>

<script setup>
import List from '@/components/Chat/List.vue'
import Display from '@/components/Chat/Display.vue'
import Input from '@/components/Chat/Input.vue'
import Title from '@/components/Chat/Title.vue'
</script>

<style  scoped>
@import url(Chat.css);
</style>