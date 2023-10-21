<script lang="ts" setup>
import { computed, ref, watch } from 'vue'
import { tnNavPage } from '@tuniao/tnui-vue3-uniapp/utils'

import { pageContainerProps } from './page-container'

const props = defineProps(pageContainerProps)

// 是否显示顶部轮播
const showTopSwiper = computed(() => props.topSwiperData?.length > 0)
// 列表数据
const listData = ref([])
watch(
  () => props.data,
  (val) => {
    if (val) {		
      listData.value = props.data
    }
  },
  {
    immediate: true,
  }
)
// 跳转到对应的演示页面
const navDemoPage = (item:any ) => {	
  tnNavPage(item.url+"?id="+item.id, 'navigateTo').catch(() => {
    uni.showToast({
      title: '参数错误',
      icon: 'none',
    })
  })
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
  <view class="page-container">
    <!-- 顶部轮播 -->
    <view v-if="showTopSwiper" class="top-swiper tn-animation-fade-in">
      <swiper class="swiper-container">
        <swiper-item
          v-for="(item, index) in topSwiperData"
          :key="index"
          class="swiper-item"
        >
          <image class="image" :src="item" mode="heightFix" />
        </swiper-item>
      </swiper>
    </view>

    <!-- 列表数据 -->
    <view class="list-container">
      <view v-for="(item, index) in listData" :key="index" class="list-item">    
        <view class="content-container">
          <view
            v-for="(dItem, dIndex) in item.data"
            :key="dIndex"
            class="content-item"
            @tap.stop="navDemoPage(dItem)"
          >		  
			<view class="item-img-cover">
				<image  class="image" :src="dItem.cover" mode="heightFix"></image>
			</view>	
				  
            <view class="bg tn-gradient-bg__blue-light" />
            <view class="data">
              <view class="title tn-text-ellipsis-1">{{ dItem.title }}</view>
              <view class="path tn-gray_text tn-text-ellipsis-1 ">              
                <text class="marginLeft0">{{ dItem.path }}</text>
              </view>
              <view class="icon tn-grey_text">
                <TnIcon :name="dItem.icon" />
              </view>
            </view>
          </view>
        </view>
      </view>
    </view>
  </view>
</template>

<style lang="scss" scoped>
@import '../styles/index.scss';
</style>
