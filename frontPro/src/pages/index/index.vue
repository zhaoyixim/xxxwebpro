<script setup lang="ts">
import { computed, nextTick, provide, reactive, ref } from 'vue'
import { onLoad, onShareAppMessage, onShareTimeline } from '@dcloudio/uni-app'
import { useOrderedChildren } from '@tuniao/tnui-vue3-uniapp/hooks'
import { tnNavPage } from '@tuniao/tnui-vue3-uniapp/utils'

import TnTabbar from '@tuniao/tnui-vue3-uniapp/components/tabbar/src/tabbar.vue'
import TnTabbarItem from '@tuniao/tnui-vue3-uniapp/components/tabbar/src/tabbar-item.vue'
import BasicPage from './sub-page/components/basic/basic.vue'
import AboutPage from './sub-page/components/about/about.vue'
import List from './sub-page/components/lists/list.vue'
import { ComponentInternalInstance, getCurrentInstance } from "vue";
import type { CSSProperties } from 'vue'
import type { IndexSubPageContext } from '@/tokens'
import { indexPageContextKey } from '@/tokens'
const {proxy} = getCurrentInstance() as ComponentInternalInstance;
const that:any = proxy;
that.$vcommon.checkExpire(); 
let meminfo = that.$vcache.vget("memberinfo");
if(!meminfo) tnNavPage("/info-pages/login/index", 'navigateTo').catch(() => {});
const {
  children: items,
  addChild: addItem,
  removeChild: removeItem,
} = useOrderedChildren<IndexSubPageContext>()

// 底部导航栏数据
const tabbarData = [
  { text: '首页', icon: 'home-capsule-fill' },
  { text: '主页', icon: 'menu-match-fill' },
  { text: '我的', icon: 'my-job-fill' },
]
// 导航切换事件
const onTabbarChange = (index: string | number) => {
  if (!pageStatus.value?.[index as number]) {
    pageStatus.value[index as number] = true
    nextTick(() => {
      items.value?.[index as number]?.onLoad?.()
    })
  }
  items.value?.[index as number]?.onShow?.()
}

// 记录每个子页面的状态
const pageStatus = ref(Array.from({ length: tabbarData.length }, () => false))

// 当前选中的子页面的索引
const currentTabbarIndex = ref<number>(0)

// pageContainer的样式
const pageContainerStyle = computed<(index: number) => CSSProperties>(() => {
  return (index: number) => {
    const style: CSSProperties = {}
    if (index !== currentTabbarIndex.value) {
      style.display = 'none'
    }
    return style
  }
})

onLoad((options: any) => {
  // 获取当前进入的子页面的索引
  const index = Number(options?.index || 0)
  // 设置当前子页面的状态为true
  pageStatus.value[index] = true
  nextTick(() => {
    currentTabbarIndex.value = index
    setTimeout(() => {
      // 执行子页面的onLoad钩子
      items.value?.[index]?.onLoad?.()
    }, 50)
  })
})

provide(
  indexPageContextKey,
  reactive({
    items,
    addItem,
    removeItem,
  })
)
</script>

<template>
  <!-- 虚拟首页 -->
  <view class="page">
    <view v-if="pageStatus[0]" class="page__container" :style="pageContainerStyle(0)">
      <scroll-view class="scroll-view" scroll-y>
        <BasicPage />
      </scroll-view>
    </view>
    <view v-if="pageStatus[1]" class="page__container" :style="pageContainerStyle(1)">
      <scroll-view class="scroll-view" scroll-y>
        <List />
      </scroll-view>
    </view>
  
    <view v-if="pageStatus[2]" class="page__container" :style="pageContainerStyle(2)" >
      <scroll-view class="scroll-view" scroll-y>
        <AboutPage />
      </scroll-view>
    </view>
  </view>
  <TnTabbar
    v-model="currentTabbarIndex"
    fixed
    :placeholder="false"
    switch-animation
    frosted
    @change="onTabbarChange"
  >
    <TnTabbarItem
      v-for="(item, index) in tabbarData"
      :key="index"
      :name="index"
      :text="item.text"
      inactive-color="#c5cad5"
      :icon="item.icon"
      :active-icon="item.icon"
    />
  </TnTabbar>
</template>

<style lang="scss" scoped>
@import './style/index.scss';
</style>
