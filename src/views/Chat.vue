<template>
  <div class="container">
    <div class="chat-container">

      <!-- 左侧边栏 -->
      <div v-if="!isSmallScreen && isSidebarVisible" class="sidebar">
        <n-button strong secondary round type="primary" @click="createChat" class="create-chat-btn">+ 新建聊天</n-button>
        <n-scrollbar style="max-height: 800px">
          <div v-for="(chat, index) in chats" :key="index">
            <div class="chat-header">
              <button @click="selectChat(index)" :class="{ 'active': index === selectedChatIndex }" class="chat-btn">{{
                chat.title }}</button>
              <span class="edit-icon" @click="editChatTitle(index)">&#9998;</span>
            </div>
          </div>
        </n-scrollbar>
      </div>

      <!-- 右侧区域 -->
      <!-- 右侧标题 -->
      <div class="chat-window">
        <div class="title-bar" style="border-bottom: 1px solid #ccc;">
          <h3 class="chat-title">{{ selectedChatTitle }}</h3>
          <span class="icon-wrapper" @click="toggleSidebar"><i class="fas fa-list"></i></span>
        </div>

        <!-- 聊天内容 -->
        <n-scrollbar class="messages" ref="messagesContainer" v-show="isChatListVisible">
          <!-- 机器人回复的消息 -->
          <div v-for="(message, index) in selectedChat.messages" :key="index" class="message">
            <div class="message-avatar">
              <img :src="message.avatar" class="avatar" alt="avatar">
            </div>
            <div :class="message.sender === 'user' ? 'user-message' : 'bot-message'">
              {{ message.content }}
              <div v-if="message.sender === 'bot'" class="message-info">
                <span>{{ message.timestamp }}</span>
              </div>
              
              <!-- 仅机器人回复的消息包含复制按钮 -->
              <div v-if="message.sender === 'bot'"  >
                <n-button strong secondary circle type="warning" size="mini"  i
                  class="fas fa-copy fa-sm  copy-icon " @click="success(message.content)" title="复制内容">
                </n-button>
              </div>
            </div>
          </div>
        </n-scrollbar>

        <!-- 输入框 -->
        <n-card size="small" class="input-area">
          <div class="clear-button-area">
            <n-button strong secondary round type="info" @click="clearChat">清空</n-button>
          </div>
          <div class="send-message-area">
            <textarea v-model="newMessage" @keydown="handleKeyPress" placeholder="回车发送,ctrl+shift换行" rows="4" style="resize: none;"></textarea>
          </div>
        </n-card>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { useMessage } from 'naive-ui';

const newMessage = ref('');
const chats = ref([
  { title: '新建对话', messages: [] },
]);
const selectedChatIndex = ref(0);
const isSmallScreen = ref(window.innerWidth < 900);
const isChatListVisible = ref(true);
const isSidebarVisible = ref(true); // 新增侧边栏显示/隐藏状态

onMounted(() => {
  window.addEventListener('resize', handleResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
});

const handleResize = () => {
  isSmallScreen.value = window.innerWidth < 900;
};

const createChat = () => {
  const newChatTitle = `新建对话`;
  chats.value.push({ title: newChatTitle, messages: [] });
  selectedChatIndex.value = chats.value.length - 1;
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

const selectedChat = computed(() => chats.value[selectedChatIndex.value]);

const selectedChatTitle = computed(() => selectedChat.value.title);

const sendMessage = () => {
  if (newMessage.value.trim() !== '') {
    selectedChat.value.messages.push({ sender: 'user', content: newMessage.value, avatar: 'src/assets/32x32.png' });
    // Simulate bot response (you can replace this with actual API call)
    setTimeout(() => {
      const botResponse = {
        sender: 'bot',
        content: '你好，有什么可以帮助你的吗？',
        avatar: 'src/assets/32x32.png',
        timestamp: new Date().toLocaleString(), // 添加时间戳
      };
      selectedChat.value.messages.push(botResponse);
      scrollToBottom(); // 确保滚动到底部
    }, 500);
    newMessage.value = ''; // 清空输入框
  }
};

const message = useMessage();

const success = (text) => {
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
  const messagesContainer = document.querySelector('.messages');
  messagesContainer.scrollTop = messagesContainer.scrollHeight;
};

const toggleChatList = () => {
  isChatListVisible.value = !isChatListVisible.value;
};

const toggleSidebar = () => {
  isSidebarVisible.value = !isSidebarVisible.value; // 切换侧边栏显示/隐藏状态
};
</script>


<style scoped>
/* 添加标题栏样式 */
.title-bar {
  padding: 10px;
  height: auto;
  min-height: 40px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}



/* 消息样式 */
.message {
  margin-top: 25px;
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.message-avatar {
  margin-right: 10px;
}

.avatar {
  width: 40px;
  /* 头像宽度 */
  height: 40px;
  /* 头像高度 */
  border-radius: 50%;
  /* 圆形头像 */
}



.container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: calc(100vh - 90px);
  /* 使用calc函数减去导航栏的高度 */
  padding-top: 90px;
  /* 用额外的上边距来弥补导航栏的高度 */
  background-color: #000;
  /* 背景色 */
}

.chat-container {
  display: flex;
  width: 100%;
  /* 修改为和页面一样大 */
  height: 100%;
  /* 修改为和页面一样大 */
  border: none;
  /* 或者 border: 1px solid transparent; */
}

.sidebar {
  flex: 1;
  padding: 20px;
  display: flex;
  flex-direction: column;
  background-color: #171717;
  /* 墨绿色背景 */
  transition: width 0.3s ease;
  /* 添加过渡效果 */
}

.sidebar.hidden {
  width: 0;
  /* 隐藏侧边栏 */
  padding: 0;
  /* 隐藏内边距 */
}

.create-chat-btn {
  width: 100%;
  padding: 20px;
  border: none;
  border-radius: 5px;
  background-color: #1b8a72;
  color: rgb(0, 255, 195);
  cursor: pointer;
  margin-bottom: 20px;

}

.icon-wrapper {
  background-color: #e8e8e8;
  /* 设置背景色 */
  padding: 5px;
  /* 设置内边距 */
  width: 30px;
  /* 设置宽度 */
  height: 30px;
  /* 设置高度 */
  border-radius: 50%;
  /* 设置圆角半径 */
  display: flex;
  /* 让内容居中 */
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
  cursor: pointer;
  color: #000000;
}

.icon-wrapper:hover {
  background-color: #dbdbdb;
  /* 设置鼠标悬停时的背景色 */
}


.chat-btn {
  width: calc(100% - 30px);
  /* 聊天记录按钮宽度减去修改按钮的宽度 */
  padding: 10px;
  margin-bottom: 5px;
  border: none;
  border-radius: 5px;
  background-color: #268e7c;
  color: black;
  cursor: pointer;
}

.chat-btn.active {
  background-color: #28d89b;
}

.edit-icon {
  cursor: pointer;
  margin-left: 10px;
  color: #ffffff;
  /* 修改图标颜色 */
}

/* 右侧内容背景色与侧边栏相同 */
.chat-window {
  flex: 3;
  display: flex;
  flex-direction: column;
  align-items: stretch;
  padding: 20px;
  background-color: #dddddd;

  /* 左侧侧边栏的背景色 */
}

.chat-title {
  margin-top: 0;
  margin-bottom: 10px;
  text-align: left;
  /* 标题左对齐 */
}



.input-area {
  background-color: #ffffff;
  display: flex;
  flex-direction: column;
}

.clear-button-area {
  margin-bottom: 10px;
  /* 清空按钮下边距 */
}

.send-message-area {
  display: flex;
  align-items: center;
}
.n-card{
  width: 100%;
}
textarea {
  width: 100%;
  flex: 1;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 15px;
  margin-top: 10px;
  height: 90px;
  /* 固定输入框高度为 100 像素 */
}
/* 
.n-button {
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
} */



.send-btn {

  margin-left: 10px;
  /* 发送按钮左边距 */
  background-color: #34b397;
  color: rgb(255, 255, 255);
}

.send-btn:hover {
  color: rgb(255, 255, 255);
  background-color: #159378;
}

/* 隐藏侧边栏样式 */
@media (max-width: 900px) {
  .sidebar {
    display: none;
  }

  .chat-container {
    justify-content: flex-end;
  }
}

.title-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.toggle-icon {
  cursor: pointer;
  font-size: 20px;
  color: #34b397;
  /* 图标颜色 */
}


.messages {
  flex: 1;
  overflow-y: auto;
  max-height: 800px;
  padding: 10px;
  /* 添加内边距 */
  background-color: #ececec;
  /* 设置背景色 */
}

.message {
  display: flex;
  align-items: flex-start;
  /* 在交叉轴上顶部对齐 */
  margin-bottom: 10px;
}

.message-avatar {
  margin-right: 10px;
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

.user-message,
.bot-message {
  display: flex;
  flex-direction: column;
  /* 将消息内容垂直排列 */
  background-color: #d6d6d6;
  padding: 10px;
  border-radius: 5px;
  font-size: 16px;
  margin-bottom: 10px;
}

.bot-message {
  background-color: #141847;
  color: #ffffff;
}

.message-info {
  display: flex;
  align-items: center;
  color: #999;
  font-size: 14px;
  margin-top: 5px;
}

.copy-icon {
  margin-left: 10px;
  margin-top: 20px; /* 调整提示框向下的距离 */
}

.message-success {
  background-color: #34b397;
  color: #fff;
  border-radius: 5px;
  padding: 10px 20px;
  position: fixed;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
}


</style>
