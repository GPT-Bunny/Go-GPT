<template>
  <n-card class="card">
    <n-carousel show-arrow="hover" autoplay>
      <n-carousel-item>
        <img src="@/assets/1.png" alt="Slide 1" class="carousel-image" />
      </n-carousel-item>
      <n-carousel-item>
        <img src="@/assets/1.png" alt="Slide 2" class="carousel-image" />
      </n-carousel-item>
      <n-carousel-item>
        <img src="@/assets/1.png" alt="Slide 3" class="carousel-image" />
      </n-carousel-item>
    </n-carousel>
    <n-image class="cover-img" src="@/assets/1.png" preview-src-list="@/assets/1.png" />
    <n-form ref="formRef" :model="formState" :rules="rules" label-placement="left" label-width="80px">
      <n-form-item label="账户" path="username">
        <n-input v-model:value="formState.username" placeholder="请输入用户名" />
      </n-form-item>
      <n-form-item label="密码" path="password">
        <n-input v-model:value="formState.password" type="password" placeholder="请输入密码" @keypress.enter="login" />
      </n-form-item>
      <n-form-item>
    <div class="button-group">
      <n-button class="login-button"  round type="warning" :loading="loading" @click="login">
        <i class="fas fa-sign-in-alt" > 登录</i>

      </n-button>
      <n-button class="register-button" round type="warning"   @click="register">
        <i class="fas fa-sign-in-alt" > 注册</i>

      </n-button>
    </div>
  </n-form-item>
    </n-form>
    <n-alert v-if="errorMessage" type="error" title="错误" closable :show-icon="false" class="error-message">
      {{ errorMessage }}
    </n-alert>
  </n-card>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import { useMessage } from 'naive-ui';

const router = useRouter();
const message = useMessage();

const formRef = ref(null);
const loading = ref(false);

const formState = ref({
  username: '',
  password: '',
});

const errorMessage = ref('');

const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' },
  ],
};
const register = () => {
  // 注册逻辑
  console.log('注册功能待实现');
};

const login = async () => {
  if (!formRef.value?.validate()) {
    return;
  }

  try {
    loading.value = true;
    const response = await axios.post('http://localhost:8080/login', formState.value);

    if (response.data.success) {
      errorMessage.value = '';
      message.success('登录成功');
      router.push('/home');
    } else {
      errorMessage.value = '登录失败，请检查用户名和密码';
    }
  } catch (error) {
    console.error('登录失败:', error);
    errorMessage.value = '登录失败，请检查用户名和密码';
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
/* 使用 CSS 变量管理颜色和间距，方便主题切换和维护 */
:root {
  --error-color: #ff4d4f;
  --card-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  --card-border-radius: 10px;
  --card-max-width: 800px;
  --error-border-left: 4px solid var(--error-color);
}

.card {
  z-index: 2;
  width: calc(100% - 2rem); /* 为卡片两侧留出空间 */
  max-width: var(--card-max-width);
  margin: 0 auto; /* 水平居中 */
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  overflow: hidden;
  border-radius: var(--card-border-radius);
  box-shadow: var(--card-shadow);

}

/* 图片样式 */
.carousel-image,
.cover-img {
  width: 100%;
  object-fit: cover;
}

/* 错误消息样式 */
.error-message {
  margin: 12px auto 0; /* 上边距保持为 12px，左右自动（居中） */
  padding: 8px;
  border-left: var(--error-border-left);
  max-width: min(100%, 600px); /* 保持错误消息的最大宽度 */
  text-align: center; /* 错误文本居中 */
}

/* 按钮组样式 */
.button-group {
  display: flex;
  justify-content: center; /* 按钮水平居中 */
  gap: 16px; /* 按钮之间的间隔，根据需要调整 */
  margin-top: 24px; /* 顶部间距，根据需要调整 */
  width: 100%; /* 确保按钮组占满整个容器宽度 */
}

.login-button,
.register-button {
  /* 移除宽度的计算，让按钮自然居中 */
  margin: 0; /* 移除外边距 */
}

/* 响应式设计 */
@media (max-width: 576px) {
  .button-group {
    flex-direction: column; /* 在小屏幕上堆叠按钮 */
    gap: 8px; /* 在小屏幕上按钮之间的间隔 */
  }
  .login-button,
  .register-button {
    width: 100%; /* 在小屏幕上按钮宽度为100% */
    margin: 8px 0; /* 在小屏幕上调整间隔 */
  }
}
/* 中等屏幕尺寸样式 */
@media (min-width: 577px) and (max-width: 992px) {
  .card {
    max-width: 90%;
  }
}

</style>
