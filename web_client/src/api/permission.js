import client from '../utils/fetch'

// 获取所有的权限树
export function getPermissionTree() {
  return client({
    url: '/permission/tree',
    method: 'get'
  })
}



export function add(data) {
  return client({
    url: '/permission',
    method: 'post',
    data
  })
}

export function del(id) {
  const data = {
    id
  }
  return client({
    url: '/permission',
    method: 'delete',
    data
  })
}

export function edit(data) {
  return client({
    url: '/permission',
    method: 'put',
    data
  })
}
