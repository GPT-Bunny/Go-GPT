<template>
  <div class="container">
    <div class="chat-container">
      <n-layout-sider collapse-mode="transform" :collapsed-width="0" :width="200" show-trigger="bar"
        content-style="padding: 24px;" borderedv-if="shouldShowSidebar">
        <n-button strong secondary round type="warning" @click="createChat" class="create-chat-btn">+ 新建对话</n-button>
        <n-scrollbar style="max-height: 800px">
          <div v-for="(chat, index) in chats" :key="index" class="chat-header">
            <n-button strong secondary round type="warning" @click="selectChat(index)"
              :class="{ 'active': index === selectedChatIndex }" class="chat-btn">{{ chat.title }}</n-button>
          </div>
        </n-scrollbar>
      </n-layout-sider>
      <div class="chat-window">
        <div v-show="isNewChatButtonVisible" class="new-chat-button">
          <n-button strong secondary round type="warning" @click="createChat">新建对话</n-button>
        </div>
        <div v-show="isChatAreaVisible">
          <n-card size="small" hoverable>
            <h3>{{ selectedChatTitle }}</h3>
          </n-card>

          <div style="height: 740px;  padding: 20px;">
            <n-scrollbar ref="messagesContainer" v-show="isChatListVisible" @scroll="handleScroll">
              <div v-for="(message, index) in selectedChat.messages" :key="index" class="message">
                <div class="message-avatar">
                  <img :src="message.avatar" class="avatar" alt="avatar">
                </div>
                <div :class="message.sender === 'user' ? 'user-message' : 'bot-message'">
                  {{ message.content }}
                  <div v-if="message.sender === 'bot'" class="message-info">
                    <span>{{ message.timestamp }}</span>
                  </div>
                  <div v-if="message.sender === 'bot'">
                    <n-button strong secondary circle type="warning" size="mini" class="fas fa-copy fa-sm copy-icon"
                      @click="copyToClipboard(message.content)" title="复制内容"></n-button>
                  </div>
                </div>
              </div>
            </n-scrollbar>
          </div>

          <n-layout embedded content-style="padding: 10px;">
            <n-button strong secondary round type="warning" @click="clearChat">
              <i class="fas fa-trash"></i>
            </n-button>
            <n-input type="textarea"  ound clearable :autosize="{ minRows: 3 }" :value="newMessage" @update:value="updateNewMessage"
              @keydown="handleKeyPress" placeholder="回车发送,ctrl+shift换行" rows="4"  />
          </n-layout>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, nextTick } from 'vue';
import { useMessage } from 'naive-ui';
import axios from 'axios';


const newMessage = ref('');
const chats = ref([]);
const selectedChatIndex = ref(-1);
const isSmallScreen = ref(window.innerWidth < 900);
const isChatListVisible = ref(true);
const isSidebarVisible = ref(true);
let isScrollLocked = ref(true);

onMounted(() => {
  window.addEventListener('resize', handleResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
});

const handleResize = () => {
  isSmallScreen.value = window.innerWidth < 900;
};

const generateChatId = () => {
  return `${botMessagePrefix}-${messageCounter++}`;
};

const createChat = () => {
  const newChatTitle = `新的聊天`;
  const welcomeMessage = {
    sender: 'bot',
    content: '欢迎使用聊天机器人！请开始你的聊天吧。',
    avatar: 'src/assets/32x32.png',
    timestamp: new Date().toLocaleString(),
    id: generateChatId(),
  };

  const newChat = { title: newChatTitle, messages: [welcomeMessage] };
  chats.value.push(newChat); // 添加新的聊天对象到数组中
  selectedChatIndex.value = chats.value.length - 1; // 选中新的聊天
  isNewChatButtonVisible.value = false;
  isChatAreaVisible.value = true;
};

const selectChat = (index) => {
  selectedChatIndex.value = index;
};

const editChatTitle = (index) => {
  const newTitle = prompt("请输入新的聊天记录名称：", chats.value[index].title);
  if (newTitle !== null && newTitle !== '') {
    chats.value[index].title = newTitle;
  }
};

const selectedChat = computed(() => selectedChatIndex.value === -1 ? {} : chats.value[selectedChatIndex.value]);

const selectedChatTitle = computed(() => selectedChat.value.title);

const userMessagePrefix = 'user';
const botMessagePrefix = 'bot';
let messageCounter = 0;

const sendMessage = async () => {
  if (newMessage.value.trim() !== '') {
    const isFirstUserMessage = selectedChat.value.messages.length === 1 && selectedChat.value.messages[0].sender === 'bot';

    if (isFirstUserMessage) {
      selectedChat.value.messages.shift();
    }

    const userMessage = { sender: 'user', content: newMessage.value, avatar: 'src/assets/32x32.png', id: `${userMessagePrefix}-${messageCounter++}` };
    selectedChat.value.messages.push(userMessage);
    newMessage.value = '';

    try {
      const response = await axios.post('http://localhost:8080/chat', {
        messages: [{ role: 'user', content: userMessage.content }],
        language: 'zh-CN',
      });

      // 模拟延迟，等待机器人回复
      await new Promise(resolve => setTimeout(resolve, 1000));

      const botResponse = {
        sender: 'bot',
        content: response.data.reply,
        avatar: 'src/assets/32x32.png',
        timestamp: new Date().toLocaleString(),
        id: `${botMessagePrefix}-${messageCounter++}`,
      };

      selectedChat.value.messages.push(botResponse);

      await nextTick();
      scrollToBottom();
    } catch (error) {
      console.error('Failed to communicate with the server', error);
    }
  }
};


const updateNewMessage = (value) => {
  newMessage.value = value;
};

const message = useMessage();

const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text)
    .then(() => {
      message.success('已复制到剪贴板');
    })
    .catch((error) => {
      console.error('复制文本时出错：', error);
      message.error('复制失败，请重试');
    });
};

const handleKeyPress = (event) => {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault();
    sendMessage();
  } else if (event.key === 'Enter' && event.ctrlKey && event.shiftKey) {
    event.preventDefault();
    newMessage.value += '\n';
  }
};

const clearChat = () => {
  selectedChat.value.messages = [];
};

const scrollToBottom = () => {
  if (isScrollLocked) {
    const messagesContainer = document.querySelector('.messages');
    if (messagesContainer) {
      messagesContainer.scrollTop = messagesContainer.scrollHeight;
    }
  }
};

const handleScroll = () => {
  const messagesContainer = document.querySelector('.messages');
  isScrollLocked.value = messagesContainer.scrollTop === messagesContainer.scrollHeight - messagesContainer.clientHeight;
};

const toggleSidebar = () => {
  isSidebarVisible.value = !isSidebarVisible.value;
};

const isNewChatButtonVisible = ref(true);
const isChatAreaVisible = ref(false);

// 计算属性
const shouldShowSidebar = computed(() => !isSmallScreen.value && isSidebarVisible.value);

</script>


<style scoped>
@import url(Chat.css);
</style>
