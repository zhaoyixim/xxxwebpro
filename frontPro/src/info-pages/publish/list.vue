<script lang="ts" setup>
	import CustomPage from '@/components/custom-page/src/custom-page.vue'
	import TnSwitchTab from '@tuniao/tnui-vue3-uniapp/components/switch-tab/src/switch-tab.vue'
	import TnAvatar from '@tuniao/tnui-vue3-uniapp/components/avatar/src/avatar.vue'
	import TnButton from '@tuniao/tnui-vue3-uniapp/components/button/src/button.vue'
	import { tnNavPage } from '@tuniao/tnui-vue3-uniapp/utils'
	
	import { useDemoH5Page, useWxShare } from '@/hooks'
	import { ref,ComponentInternalInstance, getCurrentInstance } from "vue";
	import { useRoute,useRouter } from "vue-router";
    import defaultIcon from '@/assets/images/default_avatar.jpg';
	import Container from '@/components/subcontainer/src/container.vue'
	import {onLoad, onPullDownRefresh, onReachBottom} from "@dcloudio/uni-app";

    const pageInfo = ref({
		current_page:1,
		page_size:10,
	})
	// 当前激活的选项卡
	const currentActiveTab = ref<number>(0)
	// 选项列表
	const tabsData: string[] = ['进行中', '已完成']
	const {proxy} = getCurrentInstance() as ComponentInternalInstance;
	const that:any = proxy;
		that.$vcommon.checkExpire(); 
	const route = useRoute(); 
	let meminfo = that.$vcache.vget("memberinfo");
	
	const listsurlfinished = "/g/publishfinishedlist/"+meminfo.id
	/*未完成--进行中*/
	const listsurlunfinished = "/g/publishunfinishedlist/"+meminfo.id
	const tapDataUnfinishedLists = ref([])
	const tapDatafinished = ref([])
	
	const UnFinishedDateInit = async(flag = false) => {
		const resBack = await that.$request.post({url:listsurlunfinished,data:pageInfo.value})
		if(flag) tapDataUnfinishedLists.value.push(resBack)
		else tapDataUnfinishedLists.value=resBack
	}
	const FinishedDateInit = async(flag = false) => {
		const resBack = await that.$request.post({url:listsurlfinished,data:pageInfo.value})
		if(flag) tapDatafinished.value.push(resBack)
		else tapDatafinished.value=resBack
	}
	
	UnFinishedDateInit()
	FinishedDateInit()
	// 跳转到对应的演示页面
	
	// 测试按钮点击防抖
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
		FinishedDateInit(true)	
	})
	
</script>

<template>
	<CustomPage
	  title="游戏发布记录"
	  page-bg-color="tn-gray-light"	 
	>
	<view class="page-container">		
		<Container>
		  <view class="switch-tab-container">
		    <view class="switch-tab-item shadow-container">
		      <TnSwitchTab v-model="currentActiveTab" :tabs="tabsData">
		        <view v-if="currentActiveTab === 0" class="paddingupdown10">					
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
									
								</view>
							</view>		         	 	
						</view>	
					</view>						
		        </view>
		        <view v-if="currentActiveTab === 1"  class="paddingupdown10">
		          <view v-for="(item,index) in tapDatafinished" :key="index" class="shadow-item tn-shadow tn-type-primary_shadow">
		          						
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
		          					 <TnButton  debounce @click="()=>ButtonDebounce(item)">
		          					   我要参与游戏
		          					 </TnButton>
		          				 </view>
		          			</view>
		          		</view>		         	 	
		          	</view>	
		          </view>	
				  
		        </view>
		      </TnSwitchTab>
		    </view>
		  </view>
		</Container>
		
		
	</view>
	</CustomPage>
</template>

<style>
.staPadding{padding:4rpx 12rpx}
</style>
