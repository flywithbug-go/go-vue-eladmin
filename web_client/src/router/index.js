import Vue from 'vue'
import Router from 'vue-router'
import Login from "../views/auth/Login";
import HelloWorld from "../components/HelloWorld";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    }
  ],
  mode: 'history'
})
