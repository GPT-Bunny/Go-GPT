<template>
  <div>
    <h2>历史聊天记录</h2>
    <ul>
      <li v-for="message in messages" :key="message.id">
        <strong>{{ message.sender }}</strong>: {{ message.content }}
      </li>
    </ul>
    <button @click="createNewChat">新建对话</button>
  </div>
</template>

<script>
export default {
  props: {
    messages: {
      type: Array,
      required: true
    }
  },
  methods: {
    createNewChat() {
      // 触发新建对话的操作，可以是打开一个新对话窗口、跳转到新对话页面等
      console.log('新建对话');
      // 这里可以添加具体的新建对话操作逻辑
    }
  }
}
</script>

<style>
/* 样式可以根据需求进行调整 */
</style>
