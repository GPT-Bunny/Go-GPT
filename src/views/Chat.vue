<template>
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
      <n-layout-content >

           <Title />
           <Display   />

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



<style scoped>
/* 全局布局和颜色 */
.n-flex{
  flex: 1;
}

.n-layout-sider {
  flex: 1;
  color: #ffffff;
  background-color: #050a25;
  height: 100vh;
  padding-top: 90px;

}

.n-layout-content{
  padding: 30px;
  padding-top: 90px;
  background-color: #050a25;
  color: #ffffff;

}


</style>
