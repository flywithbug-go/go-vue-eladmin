import client from 'axios'
import { MessageBox,Message } from 'element-ui'
import store from '@/store'
import global_ from '../../config'


client.defaults.baseURL = global_.BaseURL;
client.defaults.headers.common['Authorization'] = localStorage.getItem("Authorization");
client.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
client.defaults.timeout = 60000 //60秒
//
// let client = axios.create({
//   baseURL: global_.BaseURL,
//   timeout: 60000, //60秒
//   headers:{'Authorization':localStorage.getItem("Authorization")}
// })

client.interceptors.request.use(config => {
  if (store.getters.token) {
    config.headers.common['Authorization'] = store.getters.token
  }else {
    config.headers.common['Authorization'] = ""
  }
  return config
},error => {
  console.log(error) // for debug
  return Promise.reject(error)
})

client.interceptors.response.use(response => {
  const res = response.data
  console.log("interceptorsRes:" ,res)
  if (res.code == 200){
    return res.data
  }
  if (res.code === 401) {
    MessageBox.confirm('已登出，可以取消继续留在该页面，或者重新登录','确定登出',{
      confirmButtonText: '重新登录',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      store.dispatch('FedLogOut').then(() => {
        location.reload() // 为了重新实例化vue-router对象 避免bug
      })
    })
    return Promise.reject('token 失效')
  }
  Message({
    message: res.msg,
    type:'error',
    duration: 5*1000
  })
  return Promise.reject(res.msg)
},error => {
  console.log('err:' + error)
  Message({
    message: error.message,
    type:'error',
    duration: 5*1000
  })
  return Promise.reject(error)
})

export default client
