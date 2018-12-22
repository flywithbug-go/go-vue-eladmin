
import { getToken, setToken, removeToken } from '../../utils/auth'


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
