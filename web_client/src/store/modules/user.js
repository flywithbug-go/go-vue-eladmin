
import { getToken, setToken, removeToken } from '../../utils/auth'
import {getUserInfo, loginByAccount} from "../../api/user";

const user = {
  state: {
    token: getToken(),
    name: '',
    avatar: '',
    email: '',
    status: 0,
    role: 0,
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
    SET_ROLE: (state, role) => {
      state.role = role
    },
    SET_EMAIL: (state, email) => {
      state.email = email
    }
  },
  actions: {
    LoginByAccount(context,userInfo){
      return new Promise((resolve, reject) =>{
        loginByAccount(userInfo.account,userInfo.password).then(response => {
          const token = response.token
          context.commit('SET_TOKEN', token)
          setToken(token)
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },
    GetUserInfo(context) {
      return new  Promise((resolve ,reject)=> {
        getUserInfo().then(response => {
          const user = response.user
          if (user.role < 1){
            reject('getInfo: role must  > 0 !')
          }
          context.commit('SET_ROLE',user.role)
          context.commit('SET_NAME', user.name)
          context.commit('SET_AVATAR', user.avatar)
          context.commit('SET_EMAIL', user.email)
          context.commit('SET_STATUS', user.status)
          resolve(user)
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
