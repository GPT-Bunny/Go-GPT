<template>
  <n-flex justify="space-between" style="height: 100%;padding-top: 50px;">
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
      <n-layout-content style="min-height: 1000px;min-width: 700px;width: auto;">

           <Title />
           <Display style="min-height: 800px;"/>

           <Input  />
      </n-layout-content>
    </n-layout>
  </n-flex>
</template>


<script setup>
import  List  from  '@/components/Chat/List.vue'
import  Display  from  '@/components/Chat/Display.vue'
import  Input  from  '@/components/Chat/Input.vue'
import  Title  from  '@/components/Chat/Title.vue'
</script>



