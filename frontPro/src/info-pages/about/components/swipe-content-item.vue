<script lang="ts" setup>	
import { defineProps,watch,ref } from 'vue'
const props = defineProps({
	data:{
		type: Array,
		required: true,		
	}
})

const itemList = ref()
watch(
  () => props.data,
  (val) => {
    if (val) {		
     itemList.value = val
    }
  },
  {
    immediate: true,
  }
)
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
  <view class="swipe-item-content mt15">
    <view class="avatar">
		<image :src="itemList.TitleData.cover" mode="aspectFill" />
    </view>
    <view class="content">
      <view class="top">
        <view class="nickname">{{itemList.TitleData.ctitle}}游戏</view>
        <view class="date">{{itemList.RewordList.create_at}}</view>
      </view>
      <view class="bottom">
        <view class="message">
          <text class="tn-text-ellipsis-1">
             {{itemList.TitleData.ctitle}}游戏用户中奖{{itemList.RewordList.reword_U}}U
          </text>
        </view>
      </view>
    </view>
  </view>
</template>

<style lang="scss" scoped>
.swipe-item-content {
  position: relative;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .avatar {
    flex-shrink: 0;
    flex-grow: 0;
    width: 110rpx;
    height: 110rpx;
    border-radius: 50%;
    background-color: var(--tn-color-gray-light);
  }

  .content {
    flex: 1;
    margin-left: 26rpx;
    line-height: 1;

    .top {
      width: 100%;
      display: flex;
      align-items: flex-start;
      justify-content: space-between;
      .nickname {
        font-size: 34rpx;
      }
      .date {
        font-size: 22rpx;
        color: var(--tn-color-gray);
        margin-top: 8rpx;
      }
    }

    .bottom {
      width: 100%;
      display: flex;
      align-items: flex-end;
      justify-content: space-between;

      .message {
        height: 1em;
        overflow: hidden;
        color: var(--tn-color-gray);
        margin-top: 14rpx;
        margin-right: 80rpx;
      }
    }
  }
}
</style>
