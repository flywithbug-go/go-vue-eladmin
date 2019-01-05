export default {
  route: {
    versionManager:'版本管理',
    index: '首页',
    appManager:'APP管理',
    metadata: '元数据',
    userManager: '人员管理',
    organizationStruct:'组织架构',
    permissionManager: '权限管理',
    roleManager:'角色管理',
    user: '人员',
    dashboard: '首页',
    guide: '引导页',
    icons: '图标',
  },
  navbar: {
    logOut: '退出登录',
    dashboard: '首页',
    github: '项目地址',
    screenfull: '全屏',
    theme: '换肤',
    size: '布局大小'
  },
  login: {
    title: '系统登录',
    logIn: '登录',
    username: '账号',
    password: '密码',
    any: '随便填',
    thirdparty: '第三方登录',
    thirdpartyTips: '本地不能模拟，请结合自己业务进行模拟！！！'
  },
  documentation: {
    documentation: '文档',
    github: 'Github 地址'
  },
  table: {
    title: '标题',
    type: '类型',
    search: '搜索',
    add: '添加',
    id: '序号',
    date: '时间',
    status: '状态',
    actions: '操作',
    edit: '编辑',
    delete: '删除',
    cancel: '取 消',
    confirm: '确 定'
  },
  errorLog: {
    tips: '请点击右上角bug小图标',
    description: '现在的管理后台基本都是spa的形式了，它增强了用户体验，但同时也会增加页面出问题的可能性，可能一个小小的疏忽就导致整个页面的死锁。好在 Vue 官网提供了一个方法来捕获处理异常，你可以在其中进行错误处理或者异常上报。',
    documentation: '文档介绍'
  },
  theme: {
    change: '换肤',
    documentation: '换肤文档',
    tips: 'Tips: 它区别于 navbar 上的 theme-pick, 是两种不同的换肤方法，各自有不同的应用场景，具体请参考文档。'
  },
  tagsView: {
    refresh: '刷新',
    close: '关闭',
    closeOthers: '关闭其它',
    closeAll: '关闭所有'
  },
  application: {
    table_action: "操作",
    table_add: "添加",
    table_edit: "编辑",
    table_search: '搜索',
    table_show: "",
    table_name: "名称",
    table_name_bundleId: "名称(包名)",
    table_name_placeHolder: "请输入应用名字",
    table_bundleId: "包名",
    table_bundleId_warning: "请输入应用包名",
    table_bundleId_placeHolder: "格式: com.xxx.xx",
    table_icon: "图标",
    table_desc: "描述",
    table_desc_placeholder: "请输入描述内容",
    table_owner: "负责人",
    table_time: "时间",
    table_create_time: "创建时间",
    table_createTitle: "创建应用",
    table_app_icon: "应用图标",
    table_app_icon_warning: "请上传应用图标",
  },
  appVersion: {
    versionN: "版本号",
    parentVN: "父版本号",
    parentVNPlaceholder:"父版本号(必须已存在）",
    platform: "平台",
    approvalTime: "立项时间",
    lockTime: "锁版时间",
    releaseTime:"发布时间",
    grayTime: "灰度时间",
    status: "状态",
    createTime: "添加时间",
    operate: "编辑"
  },
  selector: {
    placeholder: "请选择",
    preparing: "准备中",
    developing: "开发中",
    gray:"灰度",
    release: "已发布",
    workDone: "Work done",
    changeStatus:"状态",
    develop: "开发",
    releasing: "发布",
    cancel: "取消",
    confirm: "确定",
    confirmChange:"切换状态到",
    confirmDelete:"确定要删除当前版本？"
  },
  organization:{
    name: "名字",
    phone:"电话",
    role: "职位",
    status:"状态",
    note:"备注"
  }

}
