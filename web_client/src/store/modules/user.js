
import { getToken, setToken, removeToken } from '../../utils/auth'
import {loginByAccount} from "../../api/user";

const user = {
  state: {
    token: getToken(),
    name: '',
    avatar: '',
    status: '',
  },
  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
      if (token !== ''){
        setToken(token)
      }
    },
    SET_NAME: (state, name) => {
      state.name = name
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_STATUS: (state, status) => {
      state.status = status
    },
  },
  actions: {
    LoginByAccount(context,userInfo){
      return new Promise((resolve, reject) =>{
        loginByAccount(userInfo.account,userInfo.password).then(response => {
          if (response.code === 200) {
            const token = response.data.token
            context.commit('SET_TOKEN', token)
            setToken(token)
            resolve()
          }else {
            reject(response.msg)
          }
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 前端 登出
    FedLogOut(context) {
      return new Promise(resolve => {
        context.commit('SET_TOKEN', '')
        removeToken()
        resolve()
      })
    },
  }
}

export default user
