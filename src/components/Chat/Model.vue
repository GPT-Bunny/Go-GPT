<template>
    <n-space>
      <n-dropdown trigger="click" :options="options" @select="handleSelect">
        <n-button>点击！</n-button>
      </n-dropdown>
    </n-space>
  </template>
  
  <script>
  import { ref } from "vue";
  import { useMessage, NButton, NDropdown } from 'naive-ui';
  
  export default {
    components: { NButton, NDropdown },
    setup() {
      const message = useMessage();
      const showDropdownRef = ref(false);
      const options = [
        {
          label: "布朗酒店，伦敦",
          key: "brown's hotel, london"
        },
        {
          label: "亚特兰蒂斯巴哈马，拿骚",
          key: "atlantis nahamas, nassau"
        },
        {
          label: "比佛利山庄酒店，洛杉矶",
          key: "the beverly hills hotel, los angeles"
        }
      ];
  
      function handleSelect(key) {
        message.info(String(key));
      }
  
      function handleClick() {
        showDropdownRef.value = !showDropdownRef.value;
      }
  
      return {
        options,
        showDropdown: showDropdownRef,
        handleSelect,
        handleClick
      };
    }
  };
  </script>
  