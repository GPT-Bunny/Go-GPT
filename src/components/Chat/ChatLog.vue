<template>
  <div class="conversation-container">
    <!-- 左侧面板 -->
    <transition name="sidebar-slide">
      <div :class="['conversation-sidebar', { 'collapsed': sidebarCollapsed }]">
        <div class="conversation-sidebar-header" @click="toggleSidebar">
          <span>New Conversation</span>
          <i :class="['arrow', sidebarCollapsed ? 'right' : 'left']"></i>
          <button @click="addConversation">New</button>
        </div>
        <div v-show="!sidebarCollapsed" class="conversation-sidebar-content">
          <div class="conversation-list">
            <div v-for="(conversation, index) in conversations" :key="index" @click="selectConversation(index)" :class="{ 'active': index === selectedConversation }">
              Conversation {{ index + 1 }}
            </div>
          </div>
          <div class="sidebar-footer">
            <!-- Add styles for the footer content here -->
          </div>
        </div>
      </div>
    </transition>
    <!-- 右侧对话显示区域 -->
    <div class="conversation-content">
      <div v-if="selectedConversation !== null" class="conversation-messages">
        <div v-for="(message, index) in conversations[selectedConversation].messages" :key="index" :class="{ 'user-message': message.isUser, 'bot-message': !message.isUser }">
          {{ message.text }}
        </div>
        <input v-model="userInput" @keyup.enter="sendMessage" placeholder="Type a message..." />
      </div>
      <div v-else class="select-conversation-message">
        Please select a conversation
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

// 定义 conversations 数据
const conversations = ref([
  { messages: [{ text: 'Welcome to Chat Room!', isUser: false }] }
]);

// 定义选中的对话索引
const selectedConversation = ref(null);

// 定义用户输入的消息
const userInput = ref('');

// 定义侧边栏是否折叠的状态
const sidebarCollapsed = ref(false);

// 新增对话
const addConversation = () => {
  conversations.value.push({ messages: [] });
};

// 选择对话
const selectConversation = (index) => {
  selectedConversation.value = index;
};

// 发送消息
const sendMessage = () => {
  if (userInput.value.trim() !== '') {
    conversations.value[selectedConversation.value].messages.push({ text: userInput.value, isUser: true });
    // 模拟 bot 的回复，实际情况中可能需要通过 API 请求获取真实的回复
    setTimeout(() => {
      conversations.value[selectedConversation.value].messages.push({ text: 'Bot is typing...', isUser: false });
      setTimeout(() => {
        conversations.value[selectedConversation.value].messages.push({ text: 'Bot: Hello, I am a bot.', isUser: false });
      }, 1000);
    }, 1000);
    userInput.value = '';
  }
}

// 切换侧边栏的折叠状态
const toggleSidebar = () => {
  sidebarCollapsed.value = !sidebarCollapsed.value;
}
</script>

<style>
.conversation-container {
  display: flex;
}

.conversation-sidebar {
  width: 200px;
  background-color: #353535;
  padding: 20px;
  transition: width 0.3s ease;
}

.conversation-sidebar.collapsed {
  width: 0;
}

.conversation-sidebar-header {
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.arrow {
  border: solid #000;
  border-width: 0 2px 2px 0;
  display: inline-block;
  padding: 3px;
}

.right {
  transform: rotate(45deg);
}

.left {
  transform: rotate(-135deg);
}

.conversation-sidebar-content {
  overflow: hidden;
}

.conversation-list {
  overflow: hidden;
}

.conversation-list div {
  cursor: pointer;
}

.conversation-list div.active {
  font-weight: bold;
}

.sidebar-footer {
  /* Add styles for the footer content here */
}

.conversation-content {
  flex: 1;
  padding: 20px;
}

.conversation-messages {
  margin-bottom: 20px;
}

.user-message {
  background-color: #e5e5ea;
  padding: 8px;
  border-radius: 8px;
  margin-bottom: 4px;
}

.bot-message {
  background-color: #007bff;
  color: #fff;
  padding: 8px;
  border-radius: 8px;
  margin-bottom: 4px;
}

.input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 8px;
}
</style>
