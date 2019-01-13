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
  return request({
    url: '/menu?id=?' + id,
    method: 'delete'
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
