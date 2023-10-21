<script setup lang="ts">
	import TnNavbar from '@tuniao/tnui-vue3-uniapp/components/navbar/src/navbar.vue'
	import { ref,ComponentInternalInstance, getCurrentInstance,reactive,watch, toRefs,onUnmounted,onMounted  } from "vue";
	import Point from "./index";
	import TnButton from '@tuniao/tnui-vue3-uniapp/components/button/src/button.vue'
	import TnCountDown from '@tuniao/tnui-vue3-uniapp/components/count-down/src/count-down.vue'
	import TnPopup from '@tuniao/tnui-vue3-uniapp/components/popup/src/popup.vue'
	import { tnNavPage } from '@tuniao/tnui-vue3-uniapp/utils'
	
	const {proxy} = getCurrentInstance() as ComponentInternalInstance;
	const that:any = proxy;
	that.$vcache.vdelete("token")
	that.$vcache.vdelete("memberinfo")
	const backIcon = "left";
	const homeIcon = "home-capsule-fill";	
	const isH5DemoPage = false;
	const navbarFrosted = false;
	const navbarPlaceholder = false;
	const navbarBottomShadow = false;	

	const sysinfo = uni.getSystemInfoSync();
	const w = 400;
	const h = 1000;
	const   ctx =ref(null), 
			screenWidth=ref(sysinfo.screenWidth),
	        screenHeight=ref(sysinfo.screenHeight),
	        timer=ref(null), 
	        points=ref([]), 	        
	        // 当前选中的模式
	        currentModeIndex=ref(0), 
	        // 模式选中滑块
	        modeSliderStyle=ref({left:0}), 
	        // 是否显示密码
	        showPassword=ref(false), 
			code =  ref(null),
	        // 倒计时提示文字
	        tips=ref('获取验证码');
	 const from = ref('');
	 const gameloop=()=> {/**粒子进行*/
	        timer.value = setTimeout(gameloop, 100);	        
	         paint();
	    }
		
	 const paint =() => {/**清空画布*/
	        ctx.value.clearRect(0, 0, w, h);
	        for (var i = 0; i < points.value.length; i++) {
				points.value[i].move(w, h);
				points.value[i].draw(ctx.value);
			  for (var j = i + 1; j < points.value.length; j++) {
	            points.value[i].drawLine(ctx.value, points.value[j]);
	          }
	        }
	        ctx.value.draw();
	    }
	
	const formData = reactive({
		cuname:"",
		msgcode:"",
		passwd:"",
		is_login:1,
	})
	
	const modeSwitch=(index) =>{
		formData.cuname="",
		formData.msgcode="",
		formData.passwd="",		
	    currentModeIndex.value = index
	    showPassword.value = false
		let sliderWidth = uni.upx2px(476 / 2)	
		modeSliderStyle.value.left = `${sliderWidth * index}px`		
	}	     
	
	
	const formSubmit = async ()=>{
		
		formData.is_login = 1;
		if(formData.cuname == ""){
			uni.showToast({
					title:"账户名称不能为空"
			});
			return ;
		}
		if(currentModeIndex.value ==1){
			formData.is_login = 2;
			if(formData.msgcode == ""){
				uni.showToast({
						title:"验证码不能为空"
				});
				return ;
			}						
		}				
		if(formData.passwd == ""){
			uni.showToast({
					title:"密码不能为空"
			});
			return ;
		}		
		const loginUrl = '/g/login'		
		let resData = await that.$request.post({url:loginUrl,data:formData})
		if(resData.status_code == 400){
			uni.showToast({
					title:resData.message
			});
		}
		if(currentModeIndex.value ==1){
			if(resData.id > 0){
				uni.showToast({
						title:"注册成功,现在可以登录了"
				}).then(res=>{
					modeSwitch(0)
				});				
			}
		}else{
			
			if(resData.id > 0){
				uni.showToast({
						title:"登录成功"
				})
				
				that.$vcache.vset("memberinfo",resData)
				
				that.$vcommon.setToken(resData)				
				
				tnNavPage("/pages/index/index", 'navigateTo').catch(() => {
				  uni.showToast({
					title: '切换成功',
					icon: 'none',
				  })
				})
			
			}else{
				uni.showToast({
						title:"登录错误"
				})
			}
		}
		
		
	}
	
	
	
	const canGetCode = ref(true);
	// 开始倒计时
	const startCountDown = () => {
		canGetCode.value=false	
	   code.value?.start()
	}
	
	// 停止倒计时
	const stopCountDown = () => {
		canGetCode.value=true
	  code.value?.stop()
	}
	
	// 重置倒计时
	const resetCountDown = () => {
		canGetCode.value=true
	  code.value?.reset()
	}
	
	
	// 获取验证码
	const getCode=()=> {	
		console.log(code.value)
			if(canGetCode.value){
				
				showNormalPopup.value=true
				
				// uni.showLoading({
				// 	title: '正在获取验证码'
				// });				
				// //获取验证码
				// setTimeout(() => {
				// 	 uni.hideLoading();
				// 	 uni.showToast({
				// 		 title:"验证码已经发送"
				// 	 });
				// 	 startCountDown();				 
				// }, 1000)
			}else{
				// uni.showToast({
				// 		title:code.value.second+"秒后再重试"
				// });
			}
	    //    if (this.$refs.code.canGetCode) {
					
	    //      this.$tn.message.loading('正在获取验证码')
	    //      setTimeout(() => {
	    //        this.$tn.message.closeLoading()
	    //        this.$tn.message.toast('验证码已经发送')
	    //        // 通知组件开始计时
	    //        this.$refs.code.start()
	    //      }, 2000)
	    //    } else {
	    //      this.$tn.message.toast(this.$refs.code.secNum + '秒后再重试')
	    //    }
	}
	
	const initPoint = (options) => {		
		for (let i = 0; i < 80; i++) {
			points.value.push(new Point(Math.random() * w, Math.random() * h))
		}
		ctx.value = uni.createCanvasContext('star_canvas')
		gameloop()
	}
	initPoint();
	onUnmounted(()=>{
	    clearTimeout(timer.value)
	})

	const showNormalPopup = ref(false)
	const showNormalPopup2 = ref(false)
	const openNormalPopup = () => {
	  showNormalPopup.value = true
	}

	const openNormalPopup2 = () => {
	  showNormalPopup2.value = true
	}
	
	const popclose = ()=>{
		 showNormalPopup.value = false
		 showNormalPopup2.value = false
	}
</script>
<template>
	<!-- 弹框 -->
	<TnPopup v-model="showNormalPopup">
	  <view @click="()=>popclose()" class="popup-content center tn-gradient-bg__cool-8 whiteColor font14"> 
			请<a target="_blank" href="https://t.me/gKefu012_bot">点击此处</a>进去Telegram获取验证码			
	   </view>
	</TnPopup>
	
	<TnPopup v-model="showNormalPopup2">
	  <view @click="()=>popclose()" class="popup-content center tn-gradient-bg__cool-8 whiteColor font14"> 
			请<a target="_blank" href="https://t.me/gKefu012_bot">点击此处</a>联系客服		
	   </view>
	</TnPopup>
	
	 <view class="template-login2" >		 
		 <TnNavbar
		   fixed
		   bgColor="transparent"
		   :back-icon="backIcon"
		   :home-icon="homeIcon"
		   :frosted="navbarFrosted"
		   :bottom-shadow="navbarBottomShadow"
		   :placeholder="navbarPlaceholder"
		   :safe-area-inset-right="!isH5DemoPage"
		 >
		   
		 </TnNavbar>	
	    <canvas  canvas-id="star_canvas" class="mycanvas" :style="'width:' + screenWidth + 'px;height:' + screenHeight + 'px;'"></canvas>
    
	    <view class="login"> 
	      <view class="login__wrapper">
	        <!-- 登录/注册切换 -->
	        <view class="login__mode tn-flex tn-flex-direction-row tn-flex-nowrap tn-flex-col-center tn-flex-row-center">
	          <view  :class='["login__mode__item tn-flex-1",currentModeIndex === 0?"login__mode__item--active":""]'			   
			   @tap.stop="modeSwitch(0)">
	            登录
	          </view>
	          <view :class='["login__mode__item tn-flex-1",currentModeIndex === 1?"login__mode__item--active":""]' 
			  @tap.stop="modeSwitch(1)">
	            注册
	          </view>
	          
	          <!-- 滑块-->
	          <view class="login__mode__slider tn-cool-bg-color-15--reverse" :style="[modeSliderStyle]"></view>
	        </view>
	        
	        <!-- 输入框内容-->
	        <view class="login__info tn-flex tn-flex-direction-column tn-flex-col-center tn-flex-row-center">
	          <!-- 登录 -->
	          <block v-if="currentModeIndex === 0">
	            <view class="login__info__item__input tn-flex tn-flex-direction-row tn-flex-nowrap tn-flex-col-center tn-flex-row-left">
	              <view class="login__info__item__input__left-icon">
	                <view class="tn-icon-phone"></view>
	              </view>
	              <view class="login__info__item__input__content">
	                <input maxlength="20" v-model="formData.cuname" placeholder-class="input-placeholder" placeholder="请输入登录账号" />
	              </view>
	            </view>
	            
	            <view class="login__info__item__input tn-flex tn-flex-direction-row tn-flex-nowrap tn-flex-col-center tn-flex-row-left">
	              <view class="login__info__item__input__left-icon">
	                <view class="tn-icon-lock"></view>
	              </view>
	              <view class="login__info__item__input__content">
	                <input :password="!showPassword" v-model="formData.passwd" placeholder-class="input-placeholder" placeholder="请输入登录密码" />
	              </view>
	              <view class="login__info__item__input__right-icon" @click="showPassword = !showPassword">
	                <view :class="[showPassword ? 'tn-icon-eye' : 'tn-icon-eye-hide']"></view>
	              </view>
	            </view>
	          </block>
	          <!-- 注册 -->
	          <block v-if="currentModeIndex === 1">
	            <view class="login__info__item__input tn-flex tn-flex-direction-row tn-flex-nowrap tn-flex-col-center tn-flex-row-left">
	              <view class="login__info__item__input__left-icon">
	                <view class="tn-icon-phone"></view>
	              </view>
	              <view class="login__info__item__input__content">
	                <input maxlength="20" placeholder-class="input-placeholder" v-model="formData.cuname" placeholder="请输入注册账号" />
	              </view>
	            </view>
	            
	            <view class="login__info__item__input tn-flex tn-flex-direction-row tn-flex-nowrap tn-flex-col-center tn-flex-row-left">
	              <view class="login__info__item__input__left-icon">
	                <view class="tn-icon-code"></view>
	              </view>
	              <view class="login__info__item__input__content login__info__item__input__content--verify-code">
	                <input  v-model="formData.msgcode" type="number" placeholder-class="input-placeholder" placeholder="请输入验证码" />
	              </view>
	              <view class="login__info__item__input__right-verify-code" >
	                
					  <TnButton  @click="()=>getCode()"  :class='["tn-btn-class tn-btn tn_fixed tn-round"]' type="primary">
						<view v-if="canGetCode">
							{{ tips }}
						</view> 
						<view v-else>
							<TnCountDown ref="code"
							 :time="60" 
							  text-color="tn-gray-light"
							 :showDay="false"
							 :showHour="false"
							 :showMinute="false"
							 />秒后重新获取
							 
						</view>
					  </TnButton>
				  </view>
	            </view>
	            
	            <view class="login__info__item__input tn-flex tn-flex-direction-row tn-flex-nowrap tn-flex-col-center tn-flex-row-left">
	              <view class="login__info__item__input__left-icon">
	                <view class="tn-icon-lock"></view>
	              </view>
	              <view class="login__info__item__input__content">
	                <input :password="!showPassword" v-model="formData.passwd" placeholder-class="input-placeholder" placeholder="请输入登录密码" />
	              </view>
	              <view class="login__info__item__input__right-icon" @click="showPassword = !showPassword">
	                <view :class="[showPassword ? 'tn-icon-eye' : 'tn-icon-eye-hide']"></view>
	              </view>
	            </view>
	          </block>
	          
	          <view @click="()=>formSubmit()" class="login__info__item__button tn-cool-bg-color-7--reverse" hover-class="tn-hover" :hover-stay-time="150">{{ currentModeIndex === 0 ? '登录' : '注册'}}</view>
	          
	          <view @click="()=>openNormalPopup2()" v-if="currentModeIndex === 0" class="login__info__item__tips">忘记密码?</view>
			  
	        </view>	        
	        <!-- 其他登录方式 -->
	        <view v-if="false" class="login__way tn-flex tn-flex-col-center tn-flex-row-center">
	          <view class="tn-padding-sm tn-margin-xs">
	            <view class="login__way__item--icon tn-flex tn-flex-row-center tn-flex-col-center tn-shadow-blur tn-bg-green tn-color-white">
	              <view class="tn-icon-plane-fill"></view>
	            </view>
	          </view>	       
	        </view>
	      </view>
	    </view>
	  </view>
	</template>
	
<style lang="scss" scoped>
@import './base.css';
@import './base2.css';
@import '@/assets/css/custom_nav_bar.scss';	
@import './index.scss';
</style>