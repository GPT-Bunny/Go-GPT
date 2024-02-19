<template>
  <div class="container">
    <div class="chat-container">
      <n-layout-sider collapse-mode="transform" :collapsed-width="0" :width="300" show-trigger="bar"
        content-style="padding: 24px;"  v-if="shouldShowSidebar">
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
            <h3> {{ selectedChatTitle }}</h3>
          </n-card>
          <div style="height: 730px;  padding: 20px;">
            <n-scrollbar ref="messagesContainerRef" v-show="isChatListVisible" @scroll="handleScroll">
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

            <select v-model="selectedModel" @change="handleChangeModel" class="select">
              <option value="gpt-3.5-turbo">GPT-3.5 Turbo</option>
              <option value="gpt-3.5-turbo-1106">gpt-3.5-turbo-1106</option>
              <option value="gpt-4.0">GPT-4.0</option>
            </select>

            <n-input type="textarea" ound clearable :autosize="{ minRows: 3 }" :value="newMessage"
              @update:value="updateNewMessage" @keydown="handleKeyPress" placeholder="回车发送,ctrl+shift换行" rows="4" />
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


const ws = ref(null);

onMounted(() => {
  // 建立 WebSocket 连接
  ws.value = new WebSocket('ws://localhost:8080/ws');

  // 监听从服务器接收到的消息
  ws.value.onmessage = (event) => {
    const message = JSON.parse(event.data);
    // 假设服务器发送的消息格式是 { sender: 'bot', content: '...' }
    // 这里可以根据实际情况调整
    if (message.sender === 'bot') {
      // 将消息添加到聊天中
      selectedChat.value.messages.push({
        sender: 'bot',
        content: message.content,
        avatar: 'src/assets/32x32.png',
        timestamp: new Date().toLocaleString(),
        id: generateChatId(),
      });
    }
  };

  // 监听 WebSocket 错误事件
  ws.value.onerror = (error) => {
    console.error('WebSocket Error', error);
  };

  // 监听 WebSocket 连接关闭事件
  ws.value.onclose = (event) => {
    console.log('WebSocket Connection Closed', event);
  };
});

onUnmounted(() => {
  // 组件卸载时关闭 WebSocket 连接
  if (ws.value) {
    ws.value.close();
  }
});

// 发送消息到服务器的函数
const sendWebSocketMessage = (messageContent) => {
  if (ws.value && ws.value.readyState === WebSocket.OPEN) {
    ws.value.send(JSON.stringify({ content: messageContent }));
  }
};

const messagesContainerRef = ref(null);



const newMessage = ref('');
const chats = ref([]);
const selectedChatIndex = ref(-1);
const isSmallScreen = ref(window.innerWidth < 900);
const isChatListVisible = ref(true);
const isSidebarVisible = ref(true);
let isScrollLocked = ref(true);

const selectedModel = ref('gpt-3.5-turbo');

const handleChangeModel = async () => {
  try {
    // 确保在发送请求之前获取 selectedModel 的值
    const modelValue = selectedModel.value;

    const response = await axios.post('http://localhost:8080/set-model', {
      model: modelValue,
    });

    // 处理后端返回的响应（如果有需要的话）
    console.log(response.data);
  } catch (error) {
    console.error('Failed to communicate with the server', error);
  }
};

const messagesContainer = document.querySelector('.messages');
if (messagesContainer) {
  // 如果找到了元素，执行接下来的代码
  isScrollLocked.value = messagesContainer.scrollTop === messagesContainer.scrollHeight - messagesContainer.clientHeight;
}




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
    sendWebSocketMessage(userMessage.content);

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
  if (isScrollLocked.value && messagesContainerRef.value) {
    messagesContainerRef.value.scrollTop = messagesContainerRef.value.scrollHeight;
  }
};


const handleScroll = () => {
  const messagesContainer = document.querySelector('.messages');

  // 添加安全检查，确保 messagesContainer 存在
  if (messagesContainer) {
    isScrollLocked.value = messagesContainer.scrollTop === messagesContainer.scrollHeight - messagesContainer.clientHeight;
  }
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
