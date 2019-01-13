import request from '../utils/fetch'

// 获取所有的菜单树
export function getMenusTree() {
  return request({
    url: '/menu/tree',
    method: 'get'
  })
}

export function buildMenus() {
  return request({
    url: '/menu/list',
    method: 'get'
  })
}

export function add(data) {
  console.log("menuAdd:", data)
  if (data.iframe === 'true'){
    data.iframe = true
  } else {
    data.iframe = false
  }
  return request({
    url: '/menu',
    method: 'post',
    data
  })
}

export function del(id) {
  const  data =  {
    id
  }
  return request({
    url: '/menu',
    method: 'delete',
    data
  })
}

export function edit(data) {
  if (data.iframe === 'true'){
    data.iframe = true
  } else {
    data.iframe = false
  }
  return request({
    url: '/menu',
    method: 'put',
    data
  })
}
