<script lang="ts" setup>
import { useUniAppSystemRectInfo } from '@tuniao/tnui-vue3-uniapp/hooks'
import TnButton from '@tuniao/tnui-vue3-uniapp/components/button/src/button.vue'
import { tnNavPage } from '@tuniao/tnui-vue3-uniapp/utils'
import type { CSSProperties } from 'vue'
import { ref,ComponentInternalInstance, getCurrentInstance,reactive,watch, toRefs,onUnmounted,onMounted  } from "vue";

import {  copyToClipBoard} from '@/utils'
const { navBarInfo } = useUniAppSystemRectInfo()

const {proxy} = getCurrentInstance() as ComponentInternalInstance;
const that:any = proxy;
that.$vcommon.checkExpire();
let meminfo = that.$vcache.vget("memberinfo");

const resDataCount = ref()

const RewordCount =async() => {
	const counturL ="/g/getrewordcount/"+meminfo.id
	resDataCount.value = await that.$request.post({url:counturL});	
}
RewordCount()
// 自定义按钮样式
const customButtonStyle: CSSProperties = {
  padding: '0rpx',
  borderRadius: '0rpx',
}

//用户信息
const userInfo = reactive({
  nickName: meminfo.cuname,
  avatarUrl: meminfo.avatar,
  desp:meminfo.desp,
  wallet_address:""
})
const edituserinfo = () => {
	console.log("asdf")
	tnNavPage("/info-pages/about/edit", 'navigateTo').catch(() => {})
}
const CopyWalletAddress = () => { 
  copyToClipBoard(
    meminfo.wallet_address
  ) 
  
}

const navPage = (path: string) => {
	//is_mech
	if(meminfo.is_mech ===1){
		tnNavPage(path, 'navigateTo').catch(() => {
		  uni.showToast({
		    title: '即将上线',
		    icon: 'none',
		  })
		})
	}else{
		uni.showToast({
		  title: '请您先成为vip合作伙伴',
		  icon: 'none',
		})
	}	  
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
  <view class="about-page">
    <!-- 顶部容器 -->
    <view class="about-page__top">
      <!-- 背景毛玻璃 -->
      <view class="bg-frosted-glass" />
      <!-- 信息容器 -->
	  <view>
		  <view
		    class="info-container tn-flex-center-between"
		    :style="{ top: `${navBarInfo.height}px` }"
		  >
		    <view class="left">
		      <view class="user-nickname disflex ">
					<text v-if="meminfo.nickname">
						{{ meminfo.nickname }}
					</text> 
					<text v-else>
						{{ meminfo.cuname }}
					</text>
					<view class="paddingLeft10">
					  <TnIcon  @click="()=>edituserinfo()" name="edit-write-fill" />
					</view>
		  		</view>
		      <view v-if="meminfo.desp" class="frame-version tn-gray-dark_text">
		         {{meminfo.desp}}
		      </view>  		  
		  		  
		    </view>
		    <view class="right">
		      <view class="user-avatar tn-type-primary_bg tn-flex-center">
		        <TnIcon name="panda" />
		      </view>
		    </view>
		  </view>
		  
		  <view @tap.stop="()=>CopyWalletAddress()" class="about_wallet_address mt15 tn-gray-dark_text font14">
		     钱包地址:{{meminfo.wallet_address}}
		  </view>
	  </view>   

      <!-- 图鸟组件数据 -->
      <view
        class="tuniao-github-data tn-flex"
        :style="{ top: `${navBarInfo.height + 145}px` }"
      >
        <view class="item tn-flex tn-flex-column">
          <view class="value">中奖的U</view>
          <view class="key tn-text-capitalize">{{ resDataCount }}</view>
        </view>
		<view class="item tn-flex tn-flex-column">
		  <view class="value">余额U</view>
		  <view class="key tn-text-capitalize">{{ meminfo.total_U }}</view>
		</view>
		
      </view>

      <!-- #ifndef MP-ALIPAY -->
      <!-- vip信息介绍 -->
      <view class="tuniao-vip tn-flex-center-between">
        <view class="vip-info">
          <view class="icon">
            <TnIcon
              transparent
              transparent-bg="gradient__about-page-vip-icon"
              name="vip-text"
            />
          </view>
          <view class="join-text tn-gray_text">成为合作伙伴</view>
        </view>
        <view
          class="operation-btn"
          hover-class="tn-u-btn-hover"
          :hover-stay-time="150"
		  @tap.stop="tnNavPage('/info-pages/merchant/index')"
        >
          立即加入
        </view>
      </view>
      <!-- #endif -->
    </view>

    <view class="about-page__top--placeholder" />

    <!-- #ifndef MP-ALIPAY -->
    <!-- 图鸟信息 -->
    <view class="tuniao-info tn-flex-center tn-white_bg"> 
      <view class="item">
        <view class="icon tn-bg-image tn-gradient-bg__cool-6">
          <TnIcon name="safe-fill" />
        </view>
        <view  
		class="text"		
		@tap.stop="tnNavPage('/info-pages/cooperate/agreement')"
		>合作协议</view>
      </view>
      <view class="item button">
        <TnButton
          text
          hover-class=""
          open-type="share"
          :custom-style="customButtonStyle"
        >
          <view class="item">
            <view class="icon tn-bg-image tn-gradient-bg__cool-8">
              <TnIcon name="share-triangle" />
            </view>
            <view class="text">分享好友</view>
          </view>
        </TnButton>
      </view>
    </view>
    <!-- #endif -->

    <!-- 项目信息 -->	
	
    <view class="project-info tn-white_bg">
		
		<view class="item-container">
		  <view class="item" @tap.stop="tnNavPage('/info-pages/wallet/index')">
		    <view class="left">
		      <view class="left-icon github tn-gradient-bg__cool-7">
		        <TnIcon name="lucky-money" />
		      </view>
		      <view class="left-text">钱包管理</view>
		    </view>
		   <view class="right">
		     <TnIcon name="right" />
		   </view>
		  </view>
		</view>
		
		<view class="item-container">
		  <view class="item" @tap.stop="navPage('/info-pages/publish/index')">
		    <view class="left">
		      <view class="left-icon github tn-gradient-bg__cool-1">
		        <TnIcon name="computer" />
		      </view>
		      <view class="left-text">发布游戏</view>
		    </view>
		   <view class="right">
		     <TnIcon name="right" />
		   </view>
		  </view>
		</view>
		     
		<view class="item-container">       
		    <view class="item"  @tap.stop="tnNavPage('/info-pages/publish/list')">
		      <view class="left">
		        <view class="left-icon wechat tn-gradient-bg__cool-12">
		          <TnIcon name="wechat-fill" />
		        </view>
		        <view class="left-text">发布记录</view>
		      </view>
		      <view class="right">
		        <TnIcon name="right" />
		      </view>
		    </view>
		</view>
		
      <view class="item-container">
        <view class="item" @tap.stop="tnNavPage('/info-pages/about/record')">
          <view class="left">
            <view class="left-icon tn-gradient-bg__cool-2">
              <TnIcon name="order-fill" />
            </view>
            <view class="left-text">中奖记录</view>
          </view>
          <view class="right">
            <TnIcon name="right" />
          </view>
        </view>
      </view>
	  <view class="item-container">
	    <view class="item" @tap.stop="tnNavPage('/info-pages/wallet/withdrawlist')">
	      <view class="left">
	        <view class="left-icon github tn-gradient-bg__cool-7">
	          <TnIcon name="pay-fill" />
	        </view>
	        <view class="left-text">提现记录</view>
	      </view>
	     <view class="right">
	       <TnIcon name="right" />
	     </view>
	    </view>
	  </view>
      
      <view class="item-container">
       
          <view class="item"		  
			@tap.stop="tnNavPage('/info-pages/kefu/index')"
		  >
            <view class="left">
              <view class="left-icon tn-gradient-bg__cool-16">
                <TnIcon name="reply-fill" />
              </view>
              <view class="left-text">客服反馈</view>
            </view>
            <view class="right">
              <TnIcon name="right" />
            </view>
          </view>
      </view> 
	  
	  <view class="item-container">	   
	      <view class="item"		  
	  			@tap.stop="tnNavPage('/info-pages/logout/index')"
	  		  >
	        <view class="left">
	          <view class="left-icon tn-gradient-bg__cool-6">
	            <TnIcon name="logout" />
	          </view>
	          <view class="left-text">退出登录</view>
	        </view>
	        <view class="right">
	          <TnIcon name="right" />
	        </view>
	      </view>
	  </view> 
	  
	
    </view>
  </view>
</template>

<style lang="scss" scoped>
@import './about.scss';
</style>
