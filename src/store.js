// store.js
import { createStore } from 'vuex';

export default createStore({
  state: {
    userAvatar: '',
    botAvatar: '',
    botWelcomeMessage: '',
  },
  mutations: {
    setUserAvatar(state, avatar) {
      state.userAvatar = avatar;
    },
    setBotAvatar(state, avatar) {
      state.botAvatar = avatar;
    },
    setBotWelcomeMessage(state, message) {
      state.botWelcomeMessage = message;
    },
  },
  actions: {
    saveSettings({ commit, state }) {
      // 在这里可以将设置保存到其他地方，例如 localStorage
      localStorage.setItem('userAvatar', state.userAvatar);
      localStorage.setItem('botAvatar', state.botAvatar);
      localStorage.setItem('botWelcomeMessage', state.botWelcomeMessage);

      alert('设置已保存！');
    },
  },
});
