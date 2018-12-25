


const getters = {
  sidebar: state => state.app.sidebar,
  language: state => state.app.language,
  size: state => state.app.size,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  status: state => state.user.status,
  // role: state => state.user.role,
  roles: state => state.user.roles,
  email: state => state.user.email,
  errorLogs: state => state.errorLog.logs,
  permission_routers: state => state.permission.routers,
  addRouters: state => state.permission.addRouters

}

export default getters





