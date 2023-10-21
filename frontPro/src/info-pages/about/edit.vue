<script lang="ts" setup>
import { ref,reactive,ComponentInternalInstance, getCurrentInstance  } from "vue";
import TnForm from '@tuniao/tnui-vue3-uniapp/components/form/src/form.vue'
import TnFormItem from '@tuniao/tnui-vue3-uniapp/components/form/src/form-item.vue'
import TnButton from '@tuniao/tnui-vue3-uniapp/components/button/src/button.vue'
import TnInput from '@tuniao/tnui-vue3-uniapp/components/input/src/input.vue'
import TnImageUpload from '@tuniao/tnui-vue3-uniapp/components/image-upload/src/image-upload.vue'
import type {FormRules,TnFormInstance} from '@tuniao/tnui-vue3-uniapp'
import CustomPage from '@/components/custom-page/src/custom-page.vue'
import { useDemoH5Page } from '@/hooks'
import type { TnImageUploadInstance} from '@tuniao/tnui-vue3-uniapp'
const {proxy} = getCurrentInstance() as ComponentInternalInstance;
const that:any = proxy;
that.$vcommon.checkExpire(); 
let meminfo = that.$vcache.vget("memberinfo");


const { isDemoH5Page } = useDemoH5Page()

const formRef = ref<TnFormInstance>()
const imglists =ref([])
// 表单数据
const formData = reactive({
  nickname: meminfo.nickname || "",
  avatar: "",  
})

const autoUpload = ref(true)
const uploadLimit = ref(1)
const base64img = ref()
const beforeUploadP = async(file) =>{	
	return new Promise((resolve, reject) => {
	   const url = (file as UniApp.ChooseImageSuccessCallbackResultFile).path
	   fetch(url)
	       .then(response => response.blob())
	       .then(blob => {
	           const reader = new FileReader();	   
	           reader.onload = function() {
	               const base64Data = reader.result;
					base64img.value = base64Data
	           };	   
	           reader.readAsDataURL(blob);
	       })
	       .catch(error => {
	           console.error("获取图片失败：", error);
	       });
	   
	   resolve(url)
	});
}
const beforeRemoveP = (file) => {
	return new Promise((resolve, reject) => {
		let returnsj = ""
		base64img.value = ""
		resolve(returnsj)
	})
}
const customUploadHandlerP = (file) =>{
console.log(file,"file")
return new Promise((resolve, reject) => {
    const url = (file as UniApp.ChooseImageSuccessCallbackResultFile).path  
	  resolve(url)
  })
}

const imageUploadRef = ref<TnImageUploadInstance>()

const formRules: FormRules = {
  nickname: [
    { required: true, message: '请输入用户名', trigger: ['change', 'blur'] },    
  ] 
}
/* 提交表单 */
const submitForm =async () => {
  formRef.value?.validate(async(valid: boolean) => {
    if (valid) {
		formData.avatar = base64img.value
		let senddata = {
			id:meminfo.id,
			nickname:formData.nickname,
			avatar:base64img.value,
		}
		const reqUrl = "/g/changememinfo"
		let resDataCount = await that.$request.post({url:reqUrl,data:senddata});
		if(resDataCount.id){
			uni.showToast({
			  title: '修改成功',
			})
			uni.navigateBack({
			  delta: 1, 
			})
		}
    
    } else {
      uni.showToast({
        title: '表单校验失败',
        icon: 'none',
      })
    }
  })
}
</script>

<template>
  <CustomPage title="表单" :is-h5-demo-page="isDemoH5Page">
    <TnForm
      ref="formRef"
      :model="formData"
      :rules="formRules"
      label-width="140"
      label-position="top"
    >
	<TnFormItem label="用户名">
	  <TnInput disabled="true" v-model="meminfo.cuname" clearable />
	</TnFormItem>
	
      <TnFormItem label="更改用户昵称" prop="nickname">
        <TnInput v-model="formData.nickname" clearable />
      </TnFormItem>
    
      <TnFormItem label="图像" prop="suggestionImages">
        <TnImageUpload
          v-model="imglists"          
		  :autoUpload="autoUpload"
		  :before-upload="beforeUploadP"
		  :before-remove="beforeRemoveP"		  
		  :custom-upload-handler="customUploadHandlerP"		
		  ref="imageUploadRef"
		  :limit = "uploadLimit"
        />
      </TnFormItem>     
    </TnForm>
    <view class="tn-mt tn-flex-center">
      <TnButton size="lg" @click="submitForm"> 提交 </TnButton>
    </view>
  </CustomPage>
</template>

<style lang="scss" scoped>
.form-container {
  position: relative;
  width: 100%;
}
</style>
