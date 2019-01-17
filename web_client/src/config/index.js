const BaseURL = process.env.BASE_API

const UploadImageURL = process.env.BASE_API + '/upload/image'
const downloadImageURL = process.env.BASE_API + '/image'

const PathPermissionTree = '/permission/tree'
const PathPermissionList = '/permission/list'

const PathRoleList = '/role/list'
const PathRoleTree = '/role/tree'

export default {
  BaseURL,
  UploadImageURL,
  downloadImageURL,

  PathPermissionTree,
  PathPermissionList,
  PathRoleList,
  PathRoleTree
}
