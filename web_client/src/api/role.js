import client from '../utils/fetch'


// 获取所有的Role
export function getRoleTree() {
  return client({
    url: '/role/list',
    method: 'get'
  })
}

export function add(data) {
  return client({
    url: '/role',
    method: 'post',
    data
  })
}

export function del(id) {
  let data = {
    id
  }
  return client({
    url: '/role' ,
    method: 'delete',
    data
  })
}

export function edit(data) {
  return client({
    url: '/role',
    method: 'put',
    data
  })
}
