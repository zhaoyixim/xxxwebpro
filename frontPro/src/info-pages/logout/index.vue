<script setup lang="ts"> 
import { ref,ComponentInternalInstance, getCurrentInstance } from "vue"
import { tnNavPage } from '@tuniao/tnui-vue3-uniapp/utils'

const {proxy} = getCurrentInstance() as ComponentInternalInstance;
const that:any = proxy;

that.$vcache.vdelete("memberinfo")

tnNavPage("/info-pages/login/index", 'navigateTo').catch(() => {
				  uni.showToast({
					title: '退出登录成功',
					icon: 'none',
				  })
				})

</script>
<template>
	<view>
		
	</view>
</template>
<style>
	
</style>