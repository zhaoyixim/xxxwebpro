<script lang="ts" setup>
	import CustomPage from '@/components/custom-page/src/custom-page.vue'
	import Container from '@/components/subcontainer/src/container.vue'
	import TnCollapseItem from '@tuniao/tnui-vue3-uniapp/components/collapse/src/collapse-item.vue'
	import TnCollapse from '@tuniao/tnui-vue3-uniapp/components/collapse/src/collapse.vue'
	
	import {  reactive,ref } from 'vue'
	import type {
	  FormRules,	
	} from '@tuniao/tnui-vue3-uniapp'
	
	import { useDemoH5Page } from '@/hooks'
	import { ComponentInternalInstance, getCurrentInstance } from "vue";
	
	const {proxy} = getCurrentInstance() as ComponentInternalInstance;
	const that:any = proxy;
	that.$vcommon.checkExpire(); 
	const { isDemoH5Page } = useDemoH5Page()
	// 表单数据
	const formData = reactive({
	   wallet_address: '',
	})
	
	const formRules: FormRules = {
		wallet_address:[
			 { required: true, message: '请输入钱包地址', trigger: ['change', 'blur'] }, 
		],
	}
	const formSubmit = () => {
		
	}
	
	const currentCollapse = ref<number>()
	
	
	
</script>

<template>
<CustomPage title="商户" :is-h5-demo-page="isDemoH5Page">
	<Container class="mb40">
		
		<view class="tn-gradient-bg__cool-5 mechTxt mt20 mb20 font20">
			请用TG添加客服(@gameKefu)进行沟通合作
		</view>	
		
		<TnCollapse v-model="currentCollapse">
		  <TnCollapseItem title="我要合作" >		   
				<view class="tn-gradient-bg__cool-16 mechTxt mb20">
					请用TG添加客服(@gameKefu)进行沟通合作
				</view>	
				<view class="pBtn-wrap mt40 mb40">
					<view @click="()=>formSubmit()" class="pBtn tn-gradient-bg__cool-7">确定</view>
				</view>			   
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