<script lang="ts" setup>
import TnDemoPageContainer from '../page-container/src/page-container.vue'
import { useComponent } from './composables'

import type { PageContainerData } from '../page-container'

useComponent()

const topSwiperData = [
  'https://resource.tuniaokj.com/images/vue3/banner/vue3-2-min.jpg',
]

const pageData: PageContainerData = {
  布局组件: {
    data: [
      {
        title: '列表',
        icon: 'rocket',
        path: 'list',
        url: '/demo-pages/component/list/index',
      },
      {
        title: '吸顶',
        icon: 'rocket',
        path: 'sticky',
        url: '/demo-pages/component/sticky/index',
      },
      {
        title: '折叠面板',
        icon: 'rocket',
        path: 'collapse',
        url: '/demo-pages/component/collapse/index',
      },
      {
        title: '瀑布流',
        icon: 'rocket',
        path: 'water-fall',
        url: '/demo-pages/component/water-fall/index',
      },
      {
        title: '分段器',
        icon: 'rocket',
        path: 'subsection',
        url: '/demo-pages/component/subsection/index',
      },
      {
        title: '搜索框',
        icon: 'rocket',
        path: 'search-box',
        url: '/demo-pages/component/search-box/index',
      },
      {
        title: '选项卡切换',
        icon: 'rocket',
        path: 'switch-tabs',
        url: '/demo-pages/component/switch-tabs/index',
      },
      {
        title: '横向滚动',
        icon: 'rocket',
        path: 'scroll-list',
        url: '/demo-pages/component/scroll-list/index',
      },
    ],
  },  
}
</script>

// #ifdef MP-WEIXIN
<script lang="ts">
export default {
  options: {
    // 在微信小程序中将组件节点渲染为虚拟节点，更加接近Vue组件的表现(不会出现shadow节点下再去创建元素)
    virtualHost: true,
  },
}
</script>
// #endif

<template>
  <TnDemoPageContainer
   :top-swiper-data="topSwiperData" :data="pageData" />
</template>

<style lang="scss" scoped></style>
