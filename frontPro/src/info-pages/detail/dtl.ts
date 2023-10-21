import md5 from 'js-md5';
import request from '@/utils/requests/request';
import vcache from '@/utils/requests/vcache';

const vpay = {
	setToken:(length = 10)=>{		
		const currentTime = Date.now().toString();		
		const hash = md5(currentTime)
		const token = hash.substr(15, length); 
		return token;
	},
	setReqpost: async (itemId="0", amount = 100,notify_url,redirect_url)=>{		
		// let notify_url = "https://your.domain.com";
		// let redirect_url = "https://your.domain.com";
		let orderId = vpay.setToken(10)
		let payUniqueKey = "DXXomQx49sUAk8VHcUlDmQRuDxm4";
		let md5Str = 'amount='+amount+'&notify_url='+notify_url+'&order_id='+orderId+'&redirect_url='+redirect_url+payUniqueKey;
		let meminfo = vcache.vget("memberinfo");
		
		let signature = md5(md5Str);
		let sendData = {
			order_id:orderId,
			amount:amount,
			notify_url:notify_url,
			redirect_url:redirect_url,
			signature:signature,
			itemId:itemId,
			cuid:meminfo.id,			
		}
		let pourl ="/api/v1/order/create-transaction"
		let returndata = await request.post({url:pourl,showLoading:false,data:sendData})
		return returndata		
	},
}
export default vpay;