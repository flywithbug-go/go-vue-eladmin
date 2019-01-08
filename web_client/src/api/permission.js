import client from '../utils/fetch'




export function addPermissionRequest(para) {
  const data = {
    bundle_id:para.bundle_id,
    icon:para.icon,
    name:para.name,
    desc:para.desc
  }
  return client({
    url: '/permission/add',
    method: 'post',
    data
  })
}


export function getPermissionListRequest(query) {
  return client({
    url: '/permission/list',
    method: 'get',
    params: query
  })
}
