import client from '../utils/fetch'




export function add(data) {
  return client({
    url: '/permission',
    method: 'post',
    data
  })
}

export function del(id) {
  return client({
    url: '/permission/' + id,
    method: 'delete'
  })
}

export function edit(data) {
  return client({
    url: '/permission',
    method: 'put',
    data
  })
}
