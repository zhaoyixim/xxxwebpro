<script lang="ts" setup>
import { onShareAppMessage, onShareTimeline } from '@dcloudio/uni-app'
import CustomPage from '@/components/custom-page/src/custom-page.vue'
import {  useWxShare } from '@/hooks'
import TnEmpty from '@tuniao/tnui-vue3-uniapp/components/empty/src/empty.vue'
import { onReachBottom} from "@dcloudio/uni-app";
import {computed, ref,ComponentInternalInstance, getCurrentInstance } from "vue";
import { tnNavPage } from '@tuniao/tnui-vue3-uniapp/utils'	
import TnAvatar from '@tuniao/tnui-vue3-uniapp/components/avatar/src/avatar.vue'
    import defaultIcon from '@/assets/images/default_avatar.jpg';

const {proxy} = getCurrentInstance() as ComponentInternalInstance;
const that:any = proxy;
// 微信分享
onShareAppMessage(() => ({}))
onShareTimeline(() => ({}))
useWxShare({
  path: '/plugin-demo/tn-dropdown/index',
})
const pageInfo = ref({
		current_page:1,
		page_size:10,
})
const listsurlunfinished = "/g/getallunfinishedlists"
const tapDataUnfinishedLists = ref([])

const UnFinishedDateInit = async(flag = false) => {
		const resBack = await that.$request.post({url:listsurlunfinished,data:pageInfo.value})
		if(flag) tapDataUnfinishedLists.value.push(resBack)
		else tapDataUnfinishedLists.value=resBack
}
UnFinishedDateInit()
let debounceCount = 0
const ButtonDebounce = (items:any) => {
		debounceCount++				 
		if(items.sta==2){
				uni.showToast({title: '该游戏已结束',icon: 'none',});
			return ;
		}
		if(items.sta==0){
				uni.showToast({title: '该游戏未开始',icon: 'none',});
			return ;
		}
		if(items.sta !=1){
				uni.showToast({title: '未知错误',icon: 'none',});
			return ;
		}
		tnNavPage('/info-pages/detail/dtl?id='+items.id, 'navigateTo').catch(() => {
		  uni.showToast({
			title: '即将上线',
			icon: 'none',
		  })
		})
		 
	  if (debounceCount > 100) {debounceCount = 0;}
	}
const toDtlPage = (items:any) => {
		debounceCount++			
		tnNavPage('/info-pages/detail/dtl?id='+items.id, 'navigateTo').catch(() => {
		  uni.showToast({
			title: '即将上线',
			icon: 'none',
		  })
		})				 
		if (debounceCount > 100) {debounceCount = 0;}
}	
onReachBottom(()=>{
		pageInfo.current_page ++ ;
		UnFinishedDateInit(true)
})
</script>

<template>
  <CustomPage title="游戏列表" padding="0">
    <view v-if="tapDataUnfinishedLists.length >0 " class="lists-container">      
	  <view class="page-container">
	  	<Container>
	  	    <view class="switch-tab-item shadow-container">	  	    
				<view v-for="(item,index) in tapDataUnfinishedLists" :key="index" class="shadow-item tn-shadow tn-type-primary_shadow">
					<view @click="()=>toDtlPage(item)" class="secondTxtColor width100">							
						<view class="font14 disflex width100 spaceBetween">
							<view class="disflex justaligncenter">
								<TnAvatar  size="sm" :url="defaultIcon" />								
								<view class="ml10">{{item.cuname}}</view>
							</view>
							<view>									
								<view v-if="item.sta == 1" class="tn-gradient-bg__cool-6 font12  border6 whiteColor staPadding">状态：进行中</view>																	
								<view v-if="item.sta == 2" class="tn-gradient-bg__cool-16 font12  border6 whiteColor staPadding">状态：已完成</view>
							</view>
						</view>					
						<view class="font14 width100">
							<view class="margintop10  width100 ">
								<view>游戏名称：{{item.ctitle}}</view>
							</view>
							<view class="margintop10 font14 mainTxtColor">筹款：{{item.total_U}}U(总筹)/<text class="fourTxtColor">{{item.charged_U}}u(已筹)</text></view>	
							
							<view class="font14 margintop10 disflex width100 spaceBetween thirdTxtColor">									
								<view>该单笔将有{{item.reword_count}}次中奖机会，</view>									
								<view v-if="item.payed_number>0" class="mainblue">您付款号为:{{item.payed_number}}
									 <text v-if="item.payed_number == item.reword_number">
										 (恭喜中奖)
									 </text>
								 </view>
								 <view v-else>
									 <TnButton debounce @click="()=>ButtonDebounce(item)">
									    我要参与游戏
									 </TnButton>
								 </view>
							</view>
						</view>		         	 	
					</view>	
				</view>				
			</view>
		  </Container>
	  </view>				
    </view>		
	<view v-else class="template-page">
	  <TnEmpty size="320rpx">
	    <template #icon>
	      <view class="empty-icon">
	        <TnIcon name="clip" />
	      </view>
	    </template>
	    <template #tips>
	      <view class="empty-tips"> 暂无游戏类目 </view>
	    </template>
	  </TnEmpty>
	</view>	
  </CustomPage>
</template>

<style lang="scss" scoped>
@import './list.scss';
</style>
