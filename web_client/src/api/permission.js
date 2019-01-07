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
