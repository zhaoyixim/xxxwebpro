<script lang="ts" setup>
	import CustomPage from '@/components/custom-page/src/custom-page.vue'
	import { tnNavPage } from '@tuniao/tnui-vue3-uniapp/utils'
	import TnListItem from '@tuniao/tnui-vue3-uniapp/components/list/src/list-item.vue'
	
	import { ref,ComponentInternalInstance, getCurrentInstance } from "vue";
	import { useRoute,useRouter } from "vue-router";
	const {proxy} = getCurrentInstance() as ComponentInternalInstance;
	const that:any = proxy;
		that.$vcommon.checkExpire(); 
	let meminfo = that.$vcache.vget("memberinfo");
	const router=useRouter();

	const getItemList = ref(null);
	const listsurl = "/g/getlist/"+meminfo.id
	that.$request.get({url:listsurl}).then(res=>{
		getItemList.value = res
	})
	// 当前激活的选项卡


	const route = useRoute(); 
	//const QueryId = route.query.id
	const checkUrl ="/g/checkispub"
	
	// 跳转到对应的演示页面
	const navPage = async(items) => {	
		
		const backData = await	that.$request.post({url:checkUrl,data:{
			cuid:meminfo.id,
			itemid:parseInt(items.id)
		}});
		if(backData.remain_N == 0){
			uni.showToast({
			  title: '剩余发布量不够了',
			  icon: 'none',
			})
			return ;
		}
		if(backData.id == 0){
			uni.showToast({
			  title: '无权操作',
			  icon: 'none',
			})
			router.back();
			return ;
		}else{
			tnNavPage("/info-pages/publish/pub?id="+items.id, 'navigateTo').catch(() => {
			  uni.showToast({
			    title: '切换失败',
			    icon: 'none',
			  })
			})
		}
	 
	}
</script>

<template>
	<CustomPage
	  title="选择需要创建的游戏"
	  page-bg-color="tn-gray-light"
	  padding="0"
	>
	<view class="page-container ">		
		<view class="list-container">
		  <view class="list-item">
		    <TnListItem>请选择您要创建的游戏</TnListItem>
		  </view>
		</view>
		
		<view class="list-container">
		  <view  class="list-item">	   
			<view class="content-container">
			  <view
				v-for="(dItem, dIndex) in getItemList"
				:key="dIndex"
				class="content-item"
				@tap.stop="navPage(dItem)"
			  >		  
				<view class="item-img-cover">
					<image  class="image" :src="dItem.cover" mode="heightFix"></image>
				</view>		  
				<view class="bg tn-gradient-bg__blue-light" />
				<view class="data">
				  <view class="title tn-text-ellipsis-1">{{ dItem.title }}</view>
				  <view class="path tn-gray_text tn-text-ellipsis-1 ">              
					<text class="marginLeft0 ">
						剩余发布量
							<text class="fourTxtColor font16"> {{ dItem.remain_N }} </text>
						个
						</text>
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
	</CustomPage>
</template>

<style lang="scss">
.staPadding{padding:4rpx 12rpx}
@import "./index.scss";
</style>
