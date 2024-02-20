<template>
  <div class="nav-1">
    <div class="text-container">
      <h1>
          BUNNY
      </h1>
      <h3>一款好用的AI助理</h3>
      <h3>以AI问答和AI绘图为核心，致力于推动AIGC的发展</h3>
      <RouterLink to="chat">
        <n-button strong secondary round type="warning">
          开始创作
        </n-button>
      </RouterLink>
    </div>
    <div class="image-container">
      <img src="@/assets/logo.png" alt="Logo">
    </div>
  </div>
</template>

<style scoped>
.nav-1 {
  width: auto;
  height: 500px;
  color: #ffffff;
  display: flex;
  align-items: center;
  background-color: #1d1d1d;
  background-image: url('路径/到/背景图.jpg');
  background-size: cover;
  background-position: center;
}

.text-container {
  padding-left: 50px;
  flex: 1;
  padding-right: 20px;
}

.image-container {
  margin-left: auto;
  margin-right: 200px;
}
</style>
