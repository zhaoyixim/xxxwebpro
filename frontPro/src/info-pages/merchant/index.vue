<script lang="ts" setup>
	import CustomPage from '@/components/custom-page/src/custom-page.vue'
	import Container from '@/components/subcontainer/src/container.vue'
	import TnForm from '@tuniao/tnui-vue3-uniapp/components/form/src/form.vue'
	import TnFormItem from '@tuniao/tnui-vue3-uniapp/components/form/src/form-item.vue'
	import TnInput from '@tuniao/tnui-vue3-uniapp/components/input/src/input.vue'
	import TnCollapseItem from '@tuniao/tnui-vue3-uniapp/components/collapse/src/collapse-item.vue'
	import TnCollapse from '@tuniao/tnui-vue3-uniapp/components/collapse/src/collapse.vue'
	
	import {  reactive,ref } from 'vue'
	import type {
	  FormRules,	
	} from '@tuniao/tnui-vue3-uniapp'
	
	import { useDemoH5Page } from '@/hooks'

	import { ComponentInternalInstance, getCurrentInstance } from "vue";
	const { isDemoH5Page } = useDemoH5Page()
	const memberinfo = ref()
	
	const {proxy} = getCurrentInstance() as ComponentInternalInstance;
	const that:any = proxy;
	that.$vcommon.checkExpire(); 
	memberinfo.value = that.$vcache.vget("memberinfo")
		
	// 表单数据
	const formData = reactive({
	   wallet_address: '',
	})
	
	const formRules: FormRules = {
		wallet_address:[
			 { required: true, message: '请输入钱包地址', trigger: ['change', 'blur'] }, 
		],
	}
	const formSubmit =async () => {
		const addressurl ="/g/editwalletaddress"
		let senddata = {
			cuid:memberinfo.value.id,
			wallet_address:formData.wallet_address
		}
		 let resData =await that.$request.post({url:addressurl,data:senddata});
		if(resData.id){
			uni.showToast({
			  title: '提交成功',
			  icon: 'none',
			})
			memberinfo.value =await that.$vcommon.updatememinfo()
		}
	}
	
	const currentCollapse = ref<number>()
	
	
	
</script>

<template>
<CustomPage title="商户" :is-h5-demo-page="isDemoH5Page">
	<Container class="mb40">
		
		<view v-if="memberinfo.is_mech==1" class="tn-gradient-bg__cool-13 mechTxt mb20 font20">
			您已经是一级VIP商户
		</view>	
		
		<view v-else class="tn-gradient-bg__cool-13 mechTxt mb20">
			成为商户将需要花费2000U加盟费
		</view>
		
		<TnCollapse v-model="currentCollapse">
		  <TnCollapseItem title="我要成为商户" >		   
				<view class="tn-gradient-bg__cool-16 mechTxt mb20">
					请用TG添加客服(telegram)tg:@sysCall2023)进行申请认证
				</view>
				
			   <TnForm
				 ref="formRef"
				 :model="formData"	
				 :rules="formRules"
				 label-width="140"
				 label-position="top"
			   >
			   
				<TnFormItem label="钱包地址" class="walletBorder" prop="wallet_address" >
				  <TnInput v-model="formData.wallet_address" clearable />
				</TnFormItem>
				<view class="pBtn-wrap mt40 mb40">
					<view @click="()=>formSubmit()" class="pBtn tn-gradient-bg__cool-7">确定</view>
				</view>		
			   </TnForm>
		   
		   
		  </TnCollapseItem>
		</TnCollapse>
		
				
	</Container>
</CustomPage>	
</template>


<style lang="scss">
	.tn-collapse-item{overflow:visible}
	.mechTxt{
		border-radius:8px;
		padding:20px; color:white;
	}
.pBtn-wrap{
	width:100%;	
	.pBtn{
		display:flex;
		align-items: center;
		justify-content: center;
		margin:auto;
		border-radius:8px;
		color:white;
		height:35px;		
		width:50%;padding:20px;
	}
}	
</style>