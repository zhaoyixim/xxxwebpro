import md5 from 'js-md5'
import request from './request';
import vcache from './vcache.js';
import { useRouter } from "vue-router";
import { tnNavPage } from '@tuniao/tnui-vue3-uniapp/utils'

/*token 存储有效时间--单位秒*/
const expiretime = 3600
const router = useRouter()
const vcommon = {
	getRandomString:(code=10)=>{
	      let len = code
	      let $chars = 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz123456789' 
	      let maxLen = $chars.length
	      let pwd = ''
	      for (let i = 0; i < len; i++) {
	        pwd += $chars.charAt(Math.floor(Math.random() * maxLen ))
	      }
	      return pwd.toUpperCase()
	},
	
	checkExpire:async()=>{
		const Stime =	await vcache.vget("Stime");
		const Etime = Math.floor(Date.now() / 1000);
		const subPlustime  = Etime - Stime
		//15分钟动态
		const isExpire = subPlustime > 15*60
		let meminfo = await vcache.vget("memberinfo")		
		if(Stime && isExpire){			
			 //超过15分钟
			   vcache.vdelete("token")
			   vcache.vdelete("memberinfo")
			  //检测线上是否过期			  
			  vcommon.tokenCheck()
			  let setexpireurl = '/g/setexpire/'+meminfo.id
			  await request.get({url:setexpireurl});	
			  uni.showToast({
			  	title: '登录超时!',
			  	icon: 'error',
			  })
			tnNavPage("/info-pages/login/index", 'navigateTo').catch(() => {});
			  //uni.navigateTo({ url: '/info-pages/login/index' });	
		}else{				
			//换时间
			 vcache.vset("Stime",Math.floor(Date.now() / 1000));			 
			
		}
		if(!meminfo) {
			vcache.vdelete("token")
			vcache.vdelete("memberinfo")
			uni.navigateTo({ url: '/info-pages/login/index' });
		}		
	},
	setToken:async(authdata = {})=>{
		let timestamp=new Date().getTime()/*毫秒级*/
		let appid = vcommon.getRandomString(10);
		let secrect = vcommon.getRandomString(16);
		let mphone = vcommon.getRandomString(11);
		let senddata = {
			appId : appid.toUpperCase(),
			secrect : secrect,
			imel : md5(appid+secrect+mphone),
			timeId : timestamp,
			mphone : mphone,
			randStr : vcommon.getRandomString()
		}
		let json1 =  "";
		for(let key in senddata){
			json1 += senddata[key]
		}
		let json2 = md5(json1)
		senddata.authId = md5(json1+ json2)
		let authurl = '/g/gettoken'
		let postdata = {authcode:senddata,...authdata}
		let accesscode = await request.post({url:authurl,data:postdata});
		if(accesscode){			
			 vcache.vset("Stime",Math.floor(Date.now() / 1000));
			 vcache.vset("token",accesscode,expiretime);
			return true
		}else{			
			uni.showToast({
			  title: 'token获取失败!',
			  icon: 'none',
			})
			return false
		}
	},
	tokenCheck:async(flag = false)=>{
		/*flag == true 强制刷新token*/
		let gettoken = await vcache.vget("token")
		let meminfo =  await vcache.vget("memberinfo")
		if(!meminfo){
			uni.showToast({
			  title: '未登录!',
			  icon: 'none',
			})			
			
			setTimeout(function(){
				uni.navigateTo({ url: '/info-pages/login/index' });
			},1000)
			return ;
		}
		 if(flag || null == gettoken || undefined ==gettoken){
			 console.log("token过期")
			 let senddata = {}
			 senddata.username = meminfo.username
			 senddata.password = meminfo.password
			 return   vcommon.setToken(senddata);
		 }
		 return true
	},
	updatememinfo:async()=>{
		let getmeminfo = await vcache.vget("memberinfo")
		let memurl = '/g/memberinfo/'+getmeminfo.id	
		let meminfo = await  request.get({url:memurl})
	
		if(meminfo && meminfo.id>0) {
			   vcache.vdelete("memberinfo")
			  vcache.vset("memberinfo",meminfo)
			return meminfo;
		}else{
			uni.showToast({
			  title: '用户数据获取失败',
			  icon: 'none',
			})		
			
			setTimeout(function(){
				uni.navigateTo({ url: '/info-pages/login/index' });
			},1000)			
		}
	}
};
export default vcommon;