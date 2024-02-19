// store.js

import { createStore } from 'vuex';

export default createStore({
  state: {
    currentStyle: 'style1', // 默认样式
  },
  mutations: {
    setStyle(state, style) {
      state.currentStyle = style;
    },
  },
  // 其他配置...
});
