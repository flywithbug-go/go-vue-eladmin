import client from '../utils/fetch'

export function loginByAccount(account, password) {
  const data = {
    account,
    password
  }
  return client({
    url: '/login',
    method: 'post',
    data
  })
}

export  function getUserInfo() {
  return client({
    url: '/user/info',
    method: 'get',
  })
}

export function logout() {
  return client({
    url: '/logout',
    method: 'post'
  })
}
