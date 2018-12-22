import axios from 'axios'
import { MessageBox,Message } from 'element-ui'

import store from '@/store'

let client = axios.create({
  baseURL: 'http://localhost:6201/api',
  timeout: 1000,
  headers:{'Authorization':localStorage.getItem("Authorization")}
})

client.interceptors.request.use(config => {
  if (store.getters.token !== '') {
    config.headers['Authorization'] = store.getters.token
  }
  return config
},error => {
  console.log(error) // for debug
  return Promise.reject(error)
})

client.interceptors.response.use(response => {
  const res = response.data
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
  return res.data
},error => {
  console.log('err' + error)
  Message({
    message: error.message,
    type:'error',
    duration: 5*1000
  })
  return Promise.reject(error)
})


export default client
