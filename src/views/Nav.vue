<template>
  <nav class="nav">
    <router-link to="/home" class="logo-container">
      <img src="@/assets/32x32.png" alt="Logo" class="logo">
      <p class="logo-text">BUNNY</p>
    </router-link>

    <router-link to="/home" >
      <n-button strong secondary round type="warning">
      <i class="fas fa-home" > 首页</i>
    </n-button>
    </router-link>

    <router-link to="/Chat" >
      <n-button strong secondary round type="warning">
        <i class="fas fa-comment-alt" > 聊天</i>
      </n-button>
    </router-link>

    <router-link to="/draw" >
      <n-button strong secondary round type="warning">
        <i class="fas fa-palette" > 绘画</i>
      </n-button>
    </router-link>

      <router-link to="/login">
        <n-button strong secondary round type="warning"  class="login">
          <i class="fas fa-sign-in-alt" > 登录</i>
        </n-button>
      </router-link>

  </nav>
</template>
<style  scoped>
@import url('Nav.css');
</style>
