<script lang="ts" setup>
import TnDemoPageContainer from '../page-container/src/page-container.vue'
import { useBasic } from './composables'
import { ref,ComponentInternalInstance, getCurrentInstance } from "vue";
useBasic()
const {proxy} = getCurrentInstance() as ComponentInternalInstance;
const that:any = proxy;
const getItemList = ref({basic:{data:[]}});
const swiperPreUrl ="https://api.coincome.org"
//轮播图
const swiperData = ref([])
const initPageInfo = async()=>{
	const listsurl = "/g/getlists"
	getItemList.value.basic.data = await that.$request.get({url:listsurl});
	const swiperurl = '/g/getswiper'	
	let resSwiper =await that.$request.get({url:swiperurl})	
	if(resSwiper.length > 0){
		resSwiper.map((item:any)=>{
			swiperData.value.push(swiperPreUrl+item.imgurl)
		})
	}	
}
initPageInfo()
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
    :data="getItemList"
    :top-swiper-data="swiperData"   
    :show-title="false"
  />
</template>

<style lang="scss" scoped>
@import './styles/index.scss';
</style>
