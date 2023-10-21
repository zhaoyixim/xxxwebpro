import axios from 'axios'
// 导入axios实例的类型
import type { AxiosInstance } from 'axios'
import type { HYRequestInterceptors, HYRequestConfig } from './type'
// 引入loading组件
import { showLoadingToast  } from 'vant';
import vcache from '../vcache.js';


// 默认显示loading
const DEAFULT_LOADING = true

class HYRequest {
  // axios实例
  instance: AxiosInstance
  // 当前请求实例的拦截器
  interceptors?: HYRequestInterceptors
  // 是否显示loading
  showLoading: boolean
  // 保存的loading实例
  loading?: any

  constructor(config: HYRequestConfig) {
    // 创建axios实例
    this.instance = axios.create(config)
    // 保存基本信息
    this.interceptors = config.interceptors
    this.showLoading = config.showLoading ?? DEAFULT_LOADING

    // 使用拦截器
    // 1.从config中取出的拦截器是对应的实例的拦截器
    this.instance.interceptors.request.use(
      this.interceptors?.requestInterceptor,
      this.interceptors?.requestInterceptorCatch
    )
    this.instance.interceptors.response.use(
      this.interceptors?.responseInterceptor,
      this.interceptors?.responseInterceptorCatch
    )

    // 2.添加所有的实例都有的拦截器
    // 请求的时候,先添加的拦截器后执行
    // 响应的时候,先添加的拦截器先执行
    this.instance.interceptors.request.use(
      (config) => {        
        //console.log('所有的实例都有的拦截器: 请求成功拦截')
        // 所有的请求都添加loading
        if (this.showLoading) {
          // 添加loading
          this.loading = showLoadingToast({
            message: '加载中...',
            forbidClick: true,
          });
        }
        return config
      },
      (err) => {
        //console.log('所有的实例都有的拦截器: 请求失败拦截')
        return err
      }
    )

    this.instance.interceptors.response.use(
      (res) => {
        //console.log('所有的实例都有的拦截器: 响应成功拦截')
        // 所有的请求,将loading移除
        this.loading?.close()        
        // 因为我们需要的就是res.data,所以我们可以在所有请求实例的请求的响应拦截器里面,直接把res.data返回,这样我们就可以直接使用了
         if(res.response != undefined){
            //showToast('请求错误:'+res.response.data.message);
            console.log("请求错误:",res.response)
            
			if(res.response.status == 401){
				uni.showToast({
					title:"授权失效，请重新登录" ,
					icon: 'none',
				})
				
               //vcache.vset("vurl",window.location.hash)
               window.location.href  = "/#/info-pages/login/index"
           }
         }else{
			const data = res.data;		   
		   if ( undefined != data) {
			   if(data.status_code == 400){
			   	 return data
			   }else if(data.status_code && data.status_code == 201 || data.status_code == 200) {
								return data.data;
               }else {  
                console.log('请求失败~, 错误信息',data)
                return res;
               }       
          }else{
            console.log('请求错误:',res.message)
            //showToast('请求错误:'+res.message);
          }
         }
        // 判断当HttpErrorCode是200的时候,服务端和客户端一块自定义的错误信息
      },
      (err) => {
        //console.log('所有的实例都有的拦截器: 响应失败拦截')
        // 所有的请求,将loading移除
        this.loading?.close()

        // 判断不同的HttpErrorCode显示不同的错误信息
        if (err.response.statusCode === 404) {
          console.log('404的错误~')
        }
        return err
      }
    )
  }

  // 1.传入返回结果的类型T,这样在Promise中我们就知道返回值的类型是T了
  // 2.通过HYRequestConfig<T>,将返回值类型T告诉接口,从而在接口的返回响应拦截中指明返回值类型就是T
  request<T>(config: HYRequestConfig<T>): Promise<T> {
    // 返回一个Promise对象,好让使用者在外面拿到数据
    return new Promise((resolve, reject) => {
      // 1.单个请求对请求config的处理
      if (config.interceptors?.requestInterceptor) {
        // 如果有单个请求的拦截器,就执行一下这个函数,然后返回
        config = config.interceptors.requestInterceptor(config)
      }
      // 2.判断单个请求是否需要显示loading
      if (config.showLoading === false) {
        this.showLoading = config.showLoading
      }
      this.instance
        // request里面有两个泛型,第一个泛型默认是any,第二个泛型是AxiosResponse
        // 由于前面我们已经将res.data直接返回了,所以其实最后的数据就是T类型的,所以我们在第二个泛型中要指定返回值的类型T
        .request<any, T>(config)
        .then((res) => {
          // 1.单个请求对数据的处理
          if (config.interceptors?.responseInterceptor) {
            res = config.interceptors.responseInterceptor(res)
          }
          // 2.将showLoading设置true, 这样不会影响下一个请求
          this.showLoading = DEAFULT_LOADING

          // 3.将结果resolve返回出去
          resolve(res)
        })
        .catch((err) => {
          // 将showLoading设置true, 这样不会影响下一个请求
          this.showLoading = DEAFULT_LOADING
          reject(err)
          return err
        })
    })
  }

  get<T>(config: HYRequestConfig<T>): Promise<T> {
    return this.request<T>({ ...config, method: 'GET' })
  }

  post<T>(config: HYRequestConfig<T>): Promise<T> {
    return this.request<T>({ ...config, method: 'POST' })
  }

  delete<T>(config: HYRequestConfig<T>): Promise<T> {
    return this.request<T>({ ...config, method: 'DELETE' })
  }

  patch<T>(config: HYRequestConfig<T>): Promise<T> {
    return this.request<T>({ ...config, method: 'PATCH' })
  }
}

export default HYRequest