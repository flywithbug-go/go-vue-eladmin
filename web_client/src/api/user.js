import client from '../utils/fetch'

export function login(username, password) {
  const data = {
    username,
    password
  }
  return client({
    url: '/login',
    method: 'post',
    data
  })
}

export function getUserInfo() {
  return client({
    url: '/user/info',
    method: 'get'
  })
}

export function logout() {
  return client({
    url: '/logout',
    method: 'post'
  })
}

export function getUserListInfoRequest(query) {
  return client({
    url: '/user/list',
    method: 'get',
    params: query
  })
}


export function add(data) {
  if (data.enabled === 'true'){
    data.enabled = true
  } else {
    data.enabled = false
  }
  return client({
    url: 'user',
    method: 'post',
    data
  })
}

export function del(id) {
  const data = {
    id
  }
  return client({
    url: '/user' ,
    method: 'delete',
    data
  })
}

export function edit(data) {
  if (data.enabled === 'true'){
    data.enabled = true
  } else {
    data.enabled = false
  }
  return client({
    url: '/user',
    method: 'put',
    data
  })
}


export function validPassword(password) {
  const data = {
    password
  }
  return client({
    url: '/user/password/valid',
    method: 'get',
    data
  })
}


export function updatePassword(password,old_password) {
  const data = {
    password,
    old_password
  }
  return client({
    url: '/user/password',
    method: 'put',
    data
  })

}

export function updateEmail(code,mail,password) {
  const data = {
    mail,
    code,
    password
  }
  return client({
    url: '/user/mail',
    method: 'put',
    data
  })
}

export function sendMailVerifyCode(mail) {
  const data = {
    mail
  }
  return client({
    url: '/mail/verify',
    method: 'post',
    data
  })
}
