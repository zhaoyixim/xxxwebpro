const validate = {
	notEmpty:(val)=>{
		if(undefined == val || "" == val || null == val) return false
		else return true
	},
	telphone:(_val)=>{
		let reg = /^[1][3-9][\d]{9}/;
		return reg.test(_val)
	},
	email:(_val)=>{
		let reg = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;
		return reg.test(_val);
	},
	checkusernamelength:(_val)=>{
		if(_val.length>=5 && _val.length<=16) return true;
		return false;
	},
	checkpasswordlength:(_val)=>{
		if(_val.length>5 && _val.length<=16) return true;
		return false;
	},
	checkpasswordagain:(_val,oldval)=>{
		return _val == oldval;
	}
}
export default validate