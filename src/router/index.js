import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Layout',
      component: () => import('../views/Layout.vue'),
      children: [
        {
          path: 'Nav',
          name: 'Nav',
          component: () => import('../views/Nav.vue'),
        },
        {
          path: 'Chat',
          name: 'Chat',
          component: () => import('../views/Chat.vue'),
        },
        {
          path: 'draw',
          name: 'draw',
          component: () => import('../views/draw.vue'),
        },
        {
          path: 'home',
          name: 'home',
          component: () => import('../views/home.vue'),
        },
        {
          path: '/Login',
          name: 'Login',
          component: () => import('../views/Login.vue'),
        },
      ],
    },
  ],
});

export default router;
