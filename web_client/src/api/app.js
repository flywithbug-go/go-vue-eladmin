import client from '../utils/fetch'

export function add(para) {
  const data = {
    bundle_id: para.bundle_id,
    icon: para.icon,
    name: para.name,
    desc: para.desc,
    manager_ids: para.manager_ids
  }
  return client({
    url: '/app',
    method: 'post',
    data
  })
}
export function edit(para) {
  const data = {
    icon: para.icon,
    name: para.name,
    desc: para.desc,
    id: para.id,
    manager_ids: para.manager_ids
  }
  return client({
    url: '/app',
    method: 'put',
    data
  })
}

export function del(data) {
  return client({
    url: '/app',
    method: 'delete',
    data
  })
}

export function list(query) {
  return client({
    url: '/app/list',
    method: 'get',
    params: query
  })
}

export function simpleList() {
  return client({
    url: '/app/list/simple',
    method: 'get'
  })
}

