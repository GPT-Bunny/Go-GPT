
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
import PrimeVue from 'primevue/config'; // 追加 PrimeVue import
import store from './store';
const app = createApp(App);
app.use(createPinia());
app.use(router);

library.add(faHome, faComment, faPaintBrush, faUsers);
app.component('font-awesome-icon', FontAwesomeIcon);
app.use(naive)
app.use(store);

const md = new MarkdownIt();
const markdownText = '# Hello, *world*!';
const htmlText = md.render(markdownText);
app.config.globalProperties.$markdown = htmlText;

app.use(PrimeVue); // 在唯一的 createApp 中添加 PrimeVue 插件
app.mount('#app');
