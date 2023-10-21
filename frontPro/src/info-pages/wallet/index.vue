<script lang="ts" setup>
	import CustomPage from '@/components/custom-page/src/custom-page.vue'
	import Container from '@/components/subcontainer/src/container.vue'
	import TnForm from '@tuniao/tnui-vue3-uniapp/components/form/src/form.vue'
	import TnFormItem from '@tuniao/tnui-vue3-uniapp/components/form/src/form-item.vue'
	import TnInput from '@tuniao/tnui-vue3-uniapp/components/input/src/input.vue'
	import TnCollapseItem from '@tuniao/tnui-vue3-uniapp/components/collapse/src/collapse-item.vue'
	import TnCollapse from '@tuniao/tnui-vue3-uniapp/components/collapse/src/collapse.vue'
	
	import { ref,reactive,ComponentInternalInstance, getCurrentInstance } from "vue";
	import type { FormRules	} from '@tuniao/tnui-vue3-uniapp'
	
	import { useDemoH5Page } from '@/hooks'
	const { isDemoH5Page } = useDemoH5Page()
	const memberinfo = ref()
	const {proxy} = getCurrentInstance() as ComponentInternalInstance;
	const that:any = proxy;
  
	that.$vcommon.checkExpire(); 
	memberinfo.value = that.$vcache.vget("memberinfo")
	//memberinfo =that.$vcommon.updatememinfo()
	
	
	// 表单数据
	const formData = reactive({
	   wallet_address: '',
	})
	
	const formWithDraw = reactive({
	   wallet_address: '',
	   withDrawNum:"",
	})
	const formRules: FormRules = {
		wallet_address:[
			 { required: true, message: '请输入钱包地址', trigger: ['change', 'blur'] }, 
		],
	}
	const formWithDrawRules: FormRules = {
		withDrawNum:[
			 { required: true, message: '请输入提现金额', trigger: ['change', 'blur'] }, 
		],
	}
	const WithdrawInputHandle = () =>{		
		if(parseInt(formWithDraw.withDrawNum) >parseInt(memberinfo.value.total_U)){	
			
			uni.showToast({
			  title: '您允许最大提现额度为'+ memberinfo.value.total_U,
			  icon: 'none',
			})
			
			formWithDraw.withDrawNum =memberinfo.value.total_U
		}
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
	const withdrawSubmit =async () => {
		const withUrl = "/g/withdraw"
		if(!memberinfo.value.id){
			uni.showToast({
			  title: '登录失效',
			  icon: 'none',
			})
			return ;
		}
		let senddata = {
			cuid:memberinfo.value.id,
			amount:parseInt(formWithDraw.withDrawNum)
		}
		 let resData =await that.$request.post({url:withUrl,data:senddata});
		if(resData.id){
			uni.showToast({
			  title: '提交成功',
			  icon: 'none',
			})
			
			setTimeout(()=>{
				uni.navigateBack({
				  delta: 1
				});
			},1000)
			
			
			return ;
		}
	}
	const currentCollapse = ref<number>()
	
	const withdrawWallet =ref<number>()
	const rateConfig = ref()
	const withdrawRate = async() => {
		const getConfigurl = "/g/getconfigbyname"
		let senddata = {
			ckey:"withdrawRate",		
		}		
		let resDataConfig =await that.$request.post({url:getConfigurl,data:senddata});
		rateConfig.value = resDataConfig
		//console.log(resDataConfig,"resDataConfig")
	}
	withdrawRate();
		
</script>

<template>
<CustomPage title="钱包地址" :is-h5-demo-page="isDemoH5Page">
	<Container class="mb40">
		<view class="tn-gradient-bg__cool-13 mechTxt mb20">
			您的钱包地址是：{{memberinfo.wallet_address}}<br><br>
			目前中奖金额提现手续费为提现的{{rateConfig.cvalue || 20}}%
		</view>		
		<TnCollapse v-model="currentCollapse" >
		  <TnCollapseItem title="更改钱包地址"  >		   
				<view class="tn-gradient-bg__cool-16 mechTxt mb20">
					为避免金额到账时间问题，请及时添加钱包地址进行更改，
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
				<view v-if="currentCollapse" class="pBtn-wrap mt40 mb40">
					<view @click.stop="()=>formSubmit()" class="pBtn tn-gradient-bg__cool-7">确定</view>
				</view>		
			   </TnForm>		   
		  </TnCollapseItem>
		</TnCollapse>
		
		<TnCollapse v-model="withdrawWallet">
		  <TnCollapseItem title="提现到钱包" >		   
				<view class="tn-gradient-bg__cool-16 mechTxt mb20">
					请输入提现金额（目前可提现：{{memberinfo.total_U || 0}}U ）
				</view>				
			   <TnForm
				 ref="formRef2"
				 :model="formWithDraw"	
				 :rules="formWithDrawRules"
				 label-width="140"
				 label-position="top"
			   >			   
				<TnFormItem label="提现金额" class="walletBorder" prop="wallet_address" >
				  <TnInput @input="()=>WithdrawInputHandle()"
				   v-model="formWithDraw.withDrawNum" 
				  type="number"
				  clearable />
				</TnFormItem>
				<view v-if="withdrawWallet" class="pBtn-wrap mt40 mb40">
					<view @click.stop="()=>withdrawSubmit()" class="pBtn tn-gradient-bg__cool-7">确定</view>
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