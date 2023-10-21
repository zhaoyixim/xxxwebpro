<script lang="ts" setup>
import { onShareAppMessage, onShareTimeline } from '@dcloudio/uni-app'
import TnSwipeAction from '@tuniao/tnui-vue3-uniapp/components/swipe-action/src/swipe-action.vue'
import TnSwipeActionItem from '@tuniao/tnui-vue3-uniapp/components/swipe-action/src/swipe-action-item.vue'

import SwipeContentItem from './components/swipe-content-item.vue'

import type { SwipeActionItemOption } from '@tuniao/tnui-vue3-uniapp'

import CustomPage from '@/components/custom-page/src/custom-page.vue'
import Container from '@/components/subcontainer/src/container.vue'

import { useDemoH5Page, useWxShare } from '@/hooks'
import { ComponentInternalInstance,ref, getCurrentInstance } from "vue";
import { onReachBottom} from "@dcloudio/uni-app";
	
const {proxy} = getCurrentInstance() as ComponentInternalInstance;
const that:any = proxy;
that.$vcommon.checkExpire(); 
let memberinfo = that.$vcache.vget("memberinfo")

// 微信分享
onShareAppMessage(() => ({}))
onShareTimeline(() => ({}))
useWxShare({
  path: '/demo-pages/component/swipe-action/index',
})
const { isDemoH5Page } = useDemoH5Page()
const pageInfo = ref({current_page:1,page_size:10,})
// 菜单选项
const swipeActionItemOption: SwipeActionItemOption = [
 
  { text: '删除', bgColor: 'tn-red' },
]
// { text: '置顶', bgColor: 'tn-orange' },
// 菜单选中事件
const swipeActionOptionSelectEvent =async (
  index: string | number,
  optionIndex: number
) => {
	let findId = rewordLists.value[0][index].RewordList.id;
	console.log(findId,"findId")
	const urlss="/g/delreword/"+findId 
	let backData = await that.$request.post({url:urlss})	
	rewordLists.value[0].splice(index,1)
}
const rewordLists = ref([])
const getRewordLists= async() => {
	const urls ="/g/getrewordlist/"+memberinfo.id
	let backData = await that.$request.post({url:urls,data:pageInfo.value})
	rewordLists.value.push(backData)
}
getRewordLists()
onReachBottom(()=>{
	pageInfo.value.current_page ++ ;	
	getRewordLists()
})
	
</script>

<template>
  <CustomPage title="游戏中奖记录" :is-h5-demo-page="isDemoH5Page">
    <Container>
      <view class="swipe-action-container">
        <view v-if="rewordLists[0].length>0" class="swipe-action-item">
          <TnSwipeAction @select="swipeActionOptionSelectEvent">
            <TnSwipeActionItem    
			  v-for="(item,index) in rewordLists[0]"
              :key="index" :options="swipeActionItemOption">
              <SwipeContentItem  :data="item" />							
            </TnSwipeActionItem>
          </TnSwipeAction>
        </view>
				<view v-else  class="template-page">
					 <TnEmpty size="320rpx">
					   <template #icon>
					     <view class="empty-icon">
					       <TnIcon name="clip" />
					     </view>
					   </template>
					   <template #tips>
					     <view class="empty-tips"> 暂无数据 </view>
					   </template>
					 </TnEmpty>
				</view>	
      </view>
    </Container>    
  </CustomPage>
</template>

<style lang="scss" scoped>
@import './styles/index.scss';
</style>
