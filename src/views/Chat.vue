<template>
  <div class="conversation-container">
    <!-- 左侧面板 -->
    <div class="conversation-sidebar">
      <div class="conversation-sidebar-header">
        <span>新对话</span>
        <input type="text" v-model="searchKeyword" placeholder="搜索对话标题">
        <button @click="addConversation">新建</button>
      </div>
      <div class="conversation-sidebar-content">
        <div class="conversation-list">
          <div v-for="(conversation, index) in filteredConversations" :key="index" @click="selectConversation(index)" :class="{ 'active': index === selectedConversation }">
            对话 {{ index + 1 }}
          </div>
        </div>
        <div class="sidebar-footer">
          <button @click="addConversation">新建对话</button>
        </div>
      </div>
    </div>
    <!-- 右侧对话显示区域 -->
    <div v-if="selectedConversation !== null" class="conversation-content">
      <div class="conversation-messages">
        <div v-for="(message, index) in conversations[selectedConversation].messages" :key="index" :class="{ 'user-message': message.isUser, 'bot-message': !message.isUser }">
          <div v-html="renderMarkdown(message.text)"></div>
          <div class="message-metadata">
            <span>{{ message.timestamp }}</span>
          </div>
          <div class="copy-button" v-if="!message.isUser">
            <button @click="copyMessage(message.text)">复制</button>
          </div>
        </div>
        <textarea v-model="userInput" @keyup.enter="sendMessage" placeholder="输入消息..."></textarea>
      </div>
    </div>
    <div v-else class="select-conversation-message">
      请选择一个对话
      <button @click="addConversation">新建对话</button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import MarkdownIt from 'markdown-it';

const conversations = ref([]);
const selectedConversation = ref(null);
const userInput = ref('');
const searchKeyword = ref('');
const md = new MarkdownIt();

const addConversation = () => {
  conversations.value.push({ messages: [] });
  selectedConversation.value = conversations.value.length - 1;
};

const selectConversation = (index) => {
  selectedConversation.value = index;
};

const sendMessage = () => {
  if (userInput.value.trim() !== '') {
    if (selectedConversation.value === null) {
      addConversation();
    }
    conversations.value[selectedConversation.value].messages.push({ text: userInput.value, isUser: true, timestamp: new Date().toLocaleString() });
    // 模拟机器人的回复
    setTimeout(() => {
      const reply = '你好，世界';
      conversations.value[selectedConversation.value].messages.push({ text: reply, isUser: false, timestamp: new Date().toLocaleString() });
    }, 500);
    userInput.value = '';
  }
};

const renderMarkdown = (content) => {
  return md.render(content);
};

const copyMessage = (text) => {
  navigator.clipboard.writeText(text);
};

const filteredConversations = computed(() => {
  if (!searchKeyword.value) {
    return conversations.value;
  }
  return conversations.value.filter(conversation => conversation.title.includes(searchKeyword.value));
});
</script>

<style scoped>
/* 省略样式 */
</style>
