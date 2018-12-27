import client from '../utils/fetch'

export function addApplication(bundleId, icon, name, desc) {
  const data = {
    bundleId,
    icon,
    name,
    desc
  }
  return client({
    url: '/app/add',
    method: 'post',
    data
  })
}
