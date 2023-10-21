const  vcache = {
	/**
	   *
	   * @param {缓存key} key
	   * @param {需要存储的缓存值} value
	   * @param {过期时间，默认0表示永久有效} expire
	   */
	  vset: (key, value, expire = 0) =>  {
	    let obj = {
	      data: value, //存储的数据
	      time: Date.now() / 1000, //记录存储的时间戳
	      expire: expire //记录过期时间，单位秒
	    }
			localStorage.setItem(key,JSON.stringify(obj))	   
	  },
	  /**
	   *
	   * @param {缓存key} key
	   */
	   vget:(key) => {
	    let val =  localStorage.getItem(key)
	    if (!val) {
	      return null
	    }
	    val = JSON.parse(val)
	    if (val.expire && Date.now() / 1000 - val.time > val.expire) {
				localStorage.removeItem(key);
	      return null
	    }
	    return val.data
	  },
	  vdelete:(key) => {
			 localStorage.removeItem(key);
		   return true
	  },
	  vclear:(key)=>{
			localStorage.clearItem(key)
		  return true
	  }
}
export default vcache
  
  