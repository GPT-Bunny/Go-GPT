<template>
    <n-space vertical>
      <n-layout>
        <n-layout-header :inverted="inverted" bordered>
          <n-menu mode="horizontal" :inverted="inverted" :options="navigator" />
        </n-layout-header>
  
        <n-layout has-sider>
          <!-- 列表 -->
          <n-layout-sider
            bordered
            show-trigger
            collapse-mode="width"
            :collapsed-width="64"
            :width="240"
            :native-scrollbar="false"
            :inverted="inverted"
            style="max-height: 320px"
          >
            <n-menu
              :inverted="inverted"
              :collapsed-width="64"
              :collapsed-icon-size="22"
              :options="menuOptions"
            />
          </n-layout-sider>
        </n-layout>
      </n-layout>
    </n-space>
  </template>
  
  <script setup>
  import { ref,h } from 'vue';
  import { NIcon } from 'naive-ui';
  import {
    BookOutline as BookIcon,
    PersonOutline as PersonIcon,
    WineOutline as WineIcon
  } from '@vicons/ionicons5';
  
  // Helper function to render icons
  const renderIcon = (icon) => {
    return () => h(NIcon, null, { default: () => h(icon) });
  };
  
  // Setup refs and reactive data
  const inverted = ref(true);
  const menuOptions = ref([
    {
      label: '且听风吟',
      key: 'hear-the-wind-sing',
      icon: renderIcon(BookIcon)
    },
    {
      label: '1973年的弹珠玩具',
      key: 'pinball-1973',
      icon: renderIcon(BookIcon),
      disabled: false,
    },
    {
      label: '寻羊冒险记',
      key: 'a-wild-sheep-chase',
      disabled: false,
      icon: renderIcon(BookIcon)
    },
  ]);

  
  const navigator = ref([
    {
      label: '聊天',
      key: 'hear-the-wind-sing',
      icon: renderIcon(BookIcon)
    },
    {
      label: 'gpts',
      key: 'pinball-1973',
      icon: renderIcon(BookIcon),
      disabled: false,
    },
    {
      label: '思维导图',
      key: 'a-wild-sheep-chase',
      disabled: false,
      icon: renderIcon(BookIcon)
    },
    {
      label: 'PDF解析',
      key: 'a-wild-sheep-chase',
      disabled: false,
      icon: renderIcon(BookIcon)
    },
  ]);
  </script>
  