import client from '../utils/fetch'

export function addApplicationRequest(para) {
  const data = {
    bundle_id: para.bundle_id,
    icon: para.icon,
    name: para.name,
    desc: para.desc
  }
  return client({
    url: '/app/add',
    method: 'post',
    data
  })
}
export function deleteApplication(data) {
  return client({
    url: '/app',
    method: 'delete',
    data
  })
}

export function getApplicationlListRequest(query) {
  return client({
    url: '/app/list',
    method: 'get',
    params: query
  })
}

export function getSimpleApplicationListRequest() {
  return client({
    url: '/app/list/simple',
    method: 'get'
  })
}

export function updateApplicationRequest(para) {
  const data = {
    icon: para.icon,
    name: para.name,
    desc: para.desc,
    id: para.id
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
/*
para:
{
	"app_id":10001,
	"version":"1.0.2",
	"parent_version":"1.0.0",
	"platform":["iOS","Android"],
	"approval_time":1546244371,
	"lock_time":1546444371,
	"gray_time":1546744371
}
*/
export function addAppVersionRequest(app_id, version, parent_version, platform, approval_time, lock_time, gray_time) {
  const data = {
    app_id,
    version,
    parent_version,
    platform,
    approval_time,
    lock_time,
    gray_time
  }
  return client({
    url: '/app/version/add',
    method: 'post',
    data
  })
}
export function updateAppVersionRequest(id, app_id, version, parent_version, platform, approval_time, lock_time, gray_time, release_time) {
  const data = {
    id,
    app_id,
    version,
    parent_version,
    platform,
    approval_time,
    lock_time,
    gray_time,
    release_time
  }
  return client({
    url: '/app/version/update',
    method: 'post',
    data
  })
}

export function removeAppVersionRequest(id) {
  const data = {
    id
  }
  return client({
    url: '/app/version/remove',
    method: 'post',
    data
  })
}

export function updateStatusAppVersionRequest(id, status) {
  const data = {
    id,
    status
  }
  return client({
    url: '/app/version/update',
    method: 'post',
    data
  })
}

