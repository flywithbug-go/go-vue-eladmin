import client from '../utils/fetch'

export function addApplicationRequest(bundle_id, icon, name, desc) {
  const data = {
    bundle_id,
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

export function getApplicationlistRequest(query) {
  return client({
    url: '/app/list',
    method: 'get',
    params: query
  })
}

export function updateApplicationRequest(icon, name, desc, id) {
  const data = {
    icon,
    name,
    desc,
    id
  }
  return client({
    url: '/app/update',
    method: 'post',
    data
  })
}

export function getAppVersionListRequest(query) {
  return client({
    url: '/app/version/list',
    method: 'get',
    params: query
  })
}

