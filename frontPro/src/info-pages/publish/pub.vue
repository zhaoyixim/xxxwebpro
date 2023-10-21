<script lang="ts" setup>
import CustomPage from '@/components/custom-page/src/custom-page.vue'
import Container from '@/components/subcontainer/src/container.vue'
import TnForm from '@tuniao/tnui-vue3-uniapp/components/form/src/form.vue'
import TnFormItem from '@tuniao/tnui-vue3-uniapp/components/form/src/form-item.vue'
import TnInput from '@tuniao/tnui-vue3-uniapp/components/input/src/input.vue'
import TnActionSheet from '@tuniao/tnui-vue3-uniapp/components/action-sheet/src/action-sheet.vue'
import { useDemoH5Page } from '@/hooks'
import type { TnActionSheetInstance } from '@tuniao/tnui-vue3-uniapp'
import { ref,watch,reactive,ComponentInternalInstance, getCurrentInstance } from "vue"
import { tnNavPage } from '@tuniao/tnui-vue3-uniapp/utils'
import { useRoute,useRouter } from "vue-router";
import type {  FormRules} from '@tuniao/tnui-vue3-uniapp'
const {proxy} = getCurrentInstance() as ComponentInternalInstance;
const that:any = proxy;
that.$vcommon.checkExpire(); 
const router=useRouter();
const route = useRoute(); 
const QueryId =parseInt(route.query.id)
const maxlength = ref(1)
const { isDemoH5Page } = useDemoH5Page()
let memberinfo = that.$vcache.vget("memberinfo")

// 表单数据
if(!QueryId){
	uni.showToast({
	  title: '无效参数',
	  icon: 'none',
	})
	router.back();
}
//检测是否有权限
const checkUrl ="/g/checkispub"
const checkLimit = ref({})
const checkAuth =async()=>{
	const backData = await	that.$request.post({url:checkUrl,data:{
		cuid:memberinfo.id,
		itemid:QueryId
	}});
	checkLimit.value = backData
	console.log(checkLimit.value,"backData")
	if(backData.id == 0){
		uni.showToast({
		  title: '无权操作',
		  icon: 'none',
		})
		router.back();
	}
}
checkAuth()

const formData = reactive({
  cuid:memberinfo.id,
  item_id:QueryId,
  ctitle: '',
  total_U:0,
  reword_count:1,
  pay_count:0,
  pay_unit:0,
  reword1:0,
  reword2:0,
  reword3:0,
  reword4:0,
  reword5:0,
  reword_num1:0,
  reword_num2:0,  
  reword_num3:0,
  reword_num4:0,
  reword_num5:0,
})

const formRules: FormRules = {
  ctitle: [
    { required: true, message: '请输入游戏标题', trigger: ['change', 'blur'] },
    {
      pattern: /^[\u4E00-\u9FA5\w-_]{3,16}$/,
      message: '请输入3-16位汉字、英文、数字、下划线、横线',
      trigger: ['change', 'blur'],
    },
  ],
  total_U:[
	 { required: true, message: '请输入游戏筹备总额', trigger: ['change', 'blur'] },	  
  ],
 
  pay_unit:[
	  { required: true, message: '输入单注需要的额度', trigger: ['change', 'blur'] }, 
  ]
};
const actionSheetRef = ref<TnActionSheetInstance>()

const setTimes = ref([
	{labelname:"第1次中奖金额",labelkey:"reword1",labelkey2:"reword_num1",labelNum:0,labelIndex:1,labelvalue:0},
	{labelname:"第2次中奖金额",labelkey:"reword2",labelkey2:"reword_num2",labelNum:0,labelIndex:2,labelvalue:0},
	{labelname:"第3次中奖金额",labelkey:"reword3",labelkey2:"reword_num3",labelNum:0,labelIndex:3,labelvalue:0},
	{labelname:"第4次中奖金额",labelkey:"reword4",labelkey2:"reword_num4",labelNum:0,labelIndex:4,labelvalue:0},
	{labelname:"第5次中奖金额",labelkey:"reword5",labelkey2:"reword_num5",labelNum:0,labelIndex:5,labelvalue:0},
])

const openActionSheet = () => {
	actionSheetRef.value?.show({
	  title: '设置中奖次数',
	  actions: [
	    { text: '中奖1次', value: 1},
	    { text: '中奖2次', value: 2},
	    { text: '中奖3次', value: 3},
		{ text: '中奖4次', value: 4},
		{ text: '中奖5次', value: 5},
	  ],
	  cancel: () => {	    
	    return true
	  },
	  select: (index: number, value?: string | number) => {
		
			if(checkLimit.value.is_limit_N ==0 && value >limitcount.value){			 			  
			 	uni.showToast({
			 		title: '中奖总数限制为' + limitcount.value,
			 		icon: 'error'
			 	})
			 	return ;			  		  
			} 
			
		  formData.reword_count = value	  
		  
	    return true
	  },
	})
}
const limitcount =ref();
watch(()=>formData.pay_unit, (newVal, oldVal) =>{
	if(checkLimit.value.is_limit_N ==0){
		limitcount.value = Math.floor(formData.total_U/newVal)		
	}	
})
const payUinitHandleInput = () => {	
	if(parseInt(formData.pay_unit)>parseInt(formData.total_U)){
		formData.pay_unit = formData.total_U
		//console.log(formData.pay_unit,formData.total_U,"formData.total_U")
	}
}

const labelNumHandleInput = (item:any) =>{
	if(item.labelNum > limitcount.value){
		uni.showToast({
			title: '中奖号码需要小于' + limitcount.value,
			icon: 'error'
		})
		item.labelNum = 0
		return ;
	}
	let findItems =setTimes.value.filter(it=>it.labelNum ==item.labelNum )
	if(item.labelNum  && findItems?.length>1){
		item.labelNum = 0
		uni.showToast({
			title: '中奖号已经使用',
			icon: 'error'
		})
		return ;
	} 
	
}
const labelvalueHandleInput = (item:any) =>{
	let totalU =0
	setTimes.value.forEach(its=>{
		totalU = totalU + its.labelvalue
	})
	if(totalU>formData.total_U){
		uni.showToast({
			title: '中奖值已经超过总值',
			icon: 'error'
		})
		item.labelvalue = 0
		return ;
	}	
}
const formDisabled = ref(false)
const formSubmit = async() => {
	if(formDisabled.value) {
		uni.showToast({
			title: '请勿重复操作',
			icon: 'none'
		})
		return ;
	}
	
	formDisabled.value=true
	formData.pay_unit = parseInt(formData.pay_unit)	
	formData.total_U = parseInt(formData.total_U)
	formData.sta = 1
	const computtotlaU = ref(0);
	setTimes.value.forEach(its=>{	
		computtotlaU.value = computtotlaU.value + parseInt(its.labelvalue)	
	})
	if(computtotlaU > formData.total_U){
		uni.showToast({
			title: '中奖值已经超过总值',
			icon: 'error'
		})
		return ;
	}
	setTimes.value.forEach(it=>{		
		formData[it.labelkey] = parseInt(it.labelvalue)
		formData[it.labelkey2] = parseInt(it.labelNum)	
	})
	
	const pubUrl = "/g/publish"
	let recData = await that.$request.post({url:pubUrl,data:formData})
	if(recData.item_id ==0){
		uni.showToast({
			title: '添加的类别出错',
			icon: 'none'
		})
		return ;
	}
	if(recData.cuid >0){		
		tnNavPage("/info-pages/detail/index?id="+recData.item_id, 'navigateTo').catch(() => {
		  uni.showToast({
		  	title: '添加成功',
		  	icon: 'none'
		  })
		})		
		
	}else{
		uni.showToast({
			title: '添加失败，请检查表单数据',
			icon: 'none'
		})
		formDisabled.value=false
	}
}

</script>

<template>
<CustomPage title="发布游戏" :is-h5-demo-page="isDemoH5Page">
	<Container>
	<TnForm
	  ref="formRef"
	  :model="formData"
	  :rules="formRules"
	  label-width="140"
	  label-position="top"
	>
	  <TnFormItem label="游戏名称" prop="username">
	    <TnInput v-model="formData.ctitle" clearable />
	  </TnFormItem>	  
	  <TnFormItem label="预筹总额度" prop="total_U">
	    <TnInput v-model="formData.total_U" type="number" placeholder="例如:1000" clearable />
	  </TnFormItem>	  
	  <TnFormItem label="单注额度(用户每笔付款最小额度)" prop="total_U">
	    <TnInput v-model="formData.pay_unit"  
		@input="payUinitHandleInput"
		type="number" placeholder="例如:100" clearable />
	  </TnFormItem>	
	  <view v-if="checkLimit.is_limit_N >0 && formData.pay_unit*formData.total_U>0" class="tn-gradient-bg__cool-4 mt10 mb10 unitItem">
		  该游戏总注数为:{{Math.floor(formData.total_U/formData.pay_unit)}}
	  </view>
	  
	  <TnFormItem v-if="checkLimit.is_limit_N ==3" label="设置付款总数" prop="reword_count">
	    <TnInput v-model="formData.pay_count"  type="number" clearable />
	  </TnFormItem> 
	  
	  <view v-if="checkLimit.is_limit_N ==0 " class="tn-gradient-bg__cool-6 mt6 mb6 unitItem">
	  		   模式1：付款数到达预设总数则游戏停止
	  </view>
	  <view v-if="checkLimit.is_limit_N ==1 " class="tn-gradient-bg__cool-6 mt6 mb6 unitItem">
	  		   模式2:不限制付款总数，手动设置完成游戏
	  </view>
	  <view v-if="checkLimit.is_limit_N ==1 " class="tn-gradient-bg__cool-6 mt6 mb6 unitItem">
	  		   模式3:设置付款总数,达到付款总数则完成游戏
	  </view>
	  
	  <TnFormItem label="设置中奖总数(默认只能设置5个中奖号)" prop="reword_count">
	    <TnInput v-model="formData.reword_count" @focus="()=>openActionSheet()" clearable />
	  </TnFormItem> 
	  
	  <TnFormItem v-show="formData.reword_count>=1"  label="设置中奖号和数额" class="itemform">	
		  <view v-if="checkLimit.is_limit_N ==0 " class="tn-gradient-bg__cool-6 mt6 mb6 unitItem">
		  		  中奖号码需要小于:{{Math.floor(formData.total_U/formData.pay_unit) || 1}}
		  </view>
		<view v-for="(item,index) in setTimes.filter(it=>it.labelIndex<=formData.reword_count) " :key="index" class="mt10">
			
			<TnInput v-model="item.labelNum" @input="()=>labelNumHandleInput(item)" type="number" clearable>
				<template v-slot:prefix>第{{item.labelIndex}}次中奖号</template>
			</TnInput>
			
			<TnInput v-model="item.labelvalue" @input="()=>labelvalueHandleInput(item)" class="mt10" type="number" clearable>
				<template v-slot:prefix>{{item.labelname}}</template>
				<template v-slot:suffix>U</template>					 
			</TnInput>
		</view>	   	
	  </TnFormItem>
	  
	  
		<view class="pBtn-wrap mt40 mb40">
			<view @click="()=>formSubmit()" class="pBtn tn-gradient-bg__cool-7">确定</view>
		</view>
	  </TnForm>
	</Container>
</CustomPage>
  <TnActionSheet ref="actionSheetRef" />
</template>
<style lang="scss" scoped>
	.unitItem{
		padding:20rpx 20rpx;color:white;
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
	
:deep .itemform{
	.tn-input__slot--left{width:48% !important;}
}
</style>