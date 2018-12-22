import Vue from 'vue'
import Vuex from 'vuex'

import app from './modules/app'
import user from './modules/user'
import errorLog from './modules/errorLog'
import getters from './getters'


Vue.use(Vuex)


const store = new Vuex.Store({
  modules: {
    app,
    user,
    errorLog,
  },
  getters:{
    getters
  }
})


export default store



