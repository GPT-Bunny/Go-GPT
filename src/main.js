import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

import '@fortawesome/fontawesome-free/css/all.css'
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faHome, faComment, faPaintBrush, faUsers } from '@fortawesome/free-solid-svg-icons';
import MarkdownIt from 'markdown-it';

import naive from 'naive-ui'
const app = createApp(App);

app.use(createPinia());
app.use(router);

library.add(faHome, faComment, faPaintBrush, faUsers); // 将图标添加到库中

app.component('font-awesome-icon', FontAwesomeIcon); // 注册 font-awesome-icon 组件

app.mount('#app');

const md = new MarkdownIt();
const markdownText = '# Hello, *world*!';
const htmlText = md.render(markdownText);


app.use(naive)
// 将渲染后的 HTML 字符串传递给 Vue 组件
app.config.globalProperties.$markdown = htmlText;
