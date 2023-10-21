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
console.log(itemList.value,"itemList")
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
    <view v-if="itemList.avatar"  class="avatar">
		<image  :src="itemList.avatar" mode="aspectFill" />
    </view>
    <view class="content">
      <view class="top">
        <view class="nickname">钱包:{{itemList.wallet_address}}</view>
        <view class="date">{{itemList.create_at}}</view>
      </view>
      <view class="bottom">
        <view class="message">
          <text class="tn-text-ellipsis-1">
             用户ID:{{itemList.cuid}},提现{{itemList.amount}}U
			 <text v-if="itemList.sta == 2" class="font12 mainblue">提现成功</text>
			 <text v-if="itemList.sta == 1" class="font12 thirdTxtColor">审核中</text>
			 <text v-if="itemList.sta == 3" class="font12 fourTxtColor">失败</text>
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
