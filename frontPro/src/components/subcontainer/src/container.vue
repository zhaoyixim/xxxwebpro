<script lang="ts" setup>
import TnTitle from '@tuniao/tnui-vue3-uniapp/components/title/src/title.vue'

import { ContainerProps } from './container'

defineProps(ContainerProps)
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
  <view class="container">
    <view
      class="container__title"
      :class="[
        {
          'title-padding': titlePadding,
        },
      ]"
    >
      <TnTitle
        mode="hLine"
        :title="title"
        color="tn-type-primary"
        size="32rpx"
      />
    </view>
    <view class="container__content">
      <slot />
    </view>
  </view>
</template>

<style lang="scss" scoped>
@import '../styles/container.scss';
</style>
