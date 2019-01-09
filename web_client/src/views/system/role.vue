<template>
  <div class="app-container">
    <!--悬浮添加按钮-->
    <section class="content">
      <fixed-button :bottom="3" class="fixed-container" @clickEvent="handleCreate">
        <svg-icon icon-class="add" class="icon-add"/>
      </fixed-button>
    </section>

  </div>
</template>

<script>
import fixedButton from '../../components/FixedButton'
import Pagination from '../../components/Pagination'

import {
  getSimpleApplicationListRequest,
  getAppVersionListRequest,
  addAppVersionRequest,
  updateAppVersionRequest,
  updateStatusAppVersionRequest, removeAppVersionRequest } from '../../api/app'
import { formatDate } from '../../utils/date'
import ElTableFooter from 'element-ui'

export default {
  name: 'MetaData',
  components: {
    ElTableFooter,
    fixedButton,
    Pagination
  },

  data() {
    return {
      listLoading: true,
      list: null,
      total: 10,
      dialogFormVisible: false,
      dialogStatus: 'create',
      dialogEditCount: 0,
      imagePlaceHolder: require('../../assets/image_placeholder.png'),
      textMap: {
        update: this.$t('application.table_edit'),
        create: this.$t('application.table_add')
      },
      temp: {
        id: 0,
        version: '',
        parent_version: '',
        platform: [],
        approval_time: undefined,
        lock_time: undefined,
        gray_time: undefined,
        create_time: undefined,
        release_time: undefined,
        status: 0,
        app_status: '',
        app_id: 0
      },
      listQuery: {
        page: 0,
        limit: 10,
        name: '',
        owner: '',
        sort: '-id',
        app_id: 0
      },
      simpleAppList: null,
      currentSimpleApp: null,
      platformOptions: [{
        value: 'iOS',
        label: 'iOS'
      }, {
        value: 'Android',
        label: 'Android'
      }, {
        value: 'H5',
        label: 'H5'
      }, {
        value: 'Server',
        label: 'Server'
      }],
      statusOptions: [{
        value: '1',
        label: this.$t('selector.preparing')
      }, {
        value: '2',
        label: this.$t('selector.developing')
      }, {
        value: '3',
        label: this.$t('selector.gray')
      }, {
        value: '4',
        label: this.$t('selector.release')
      }],
      rules: {
        version: [
          {
            required: true,
            message: '必填,格式:1.0.0',
            trigger: 'blur'
          },
          {
            pattern: /^\d+(.)\d+(.)\d+$/,
            message: '输入格式1.0.0,只能是`数字`和 `.`',
            trigger: 'blur'
          }
        ],
        approval_time: [
          {
            required: true,
            message: '必选',
            trigger: 'change'
          }
        ],
        lock_time: [
          {
            required: true,
            message: '必选',
            trigger: 'change'
          }
        ],
        gray_time: [
          {
            required: true,
            message: '必选',
            trigger: 'change'
          }
        ],
        platform: [
          {
            required: true,
            message: '必选',
            trigger: 'change'
          }
        ]
      }

    }
  },

  watch: {
    temp: {
      handler: function() {
        if (this.dialogStatus === 'update') {
          this.dialogEditCount++
        }
      },
      deep: true
    }
  },
  created() {
    this.getSimpleAppList()
  },
  methods: {
    getSimpleAppList() {
      this.listLoading = true
      getSimpleApplicationListRequest().then(response => {
        this.simpleAppList = response.list
        this.currentSimpleApp = this.simpleAppList[0]
        this.listQuery.app_id = this.currentSimpleApp.id
        this.getList()
      })
    },
    getList() {
      getAppVersionListRequest(this.listQuery).then(response => {
        this.list = response.list
        this.total = response.total
        this.listLoading = false
      }).catch(() => {
        this.listLoading = false
      })
    },
    formatStatusString(status) {
      switch (status + 1) {
        case 1:
          return this.$t('selector.preparing')
        case 2:
          return this.$t('selector.developing')
        case 3:
          return this.$t('selector.gray')
        case 4:
          return this.$t('selector.release')
        case 5:
          return this.$t('selector.workDone')
        default:
          return 'title undefined'
      }
    },
    formatStatusButtonConfirmString(status) {
      // switch (status) {
      //   case 1, 2, 3:
      //     return this.$t('selector.confirmChange')
      //   case 4:
      //     return this.$t('table.delete')
      // }
      return this.$t('selector.confirmChange')
    },
    formatTagString(status) {
      switch (status) {
        case 1:
          return ''
        case 2:
          return 'success'
        case 3:
          return 'warning'
        case 4:
          return 'danger'
        case 5:
          return 'info'
      }
      return ''
    },
    deleteVersionPopover(data) {
      data.pop_de_status = false
      this.removeAppVersion(data)
    },
    cancelPopover(data) {
      data.pop_status = false
      data.pop_de_status = false
    },
    confirmPopover(data) {
      data.pop_status = false
      this.updateStatus(data)
    },
    formatPlatform(list) {
      if (list) {
        return list.join(',')
      }
      return '-'
    },
    formatDate(time) {
      if (!time || time === 0) {
        return '-'
      }
      const date = new Date(time * 1000)
      return formatDate(date, 'yyyy-MM-dd')
    },

    resetTemp() {
      this.temp = {
        id: 0,
        version: '',
        parent_version: '',
        platform: [],
        approval_time: undefined,
        lock_time: undefined,
        gray_time: undefined,
        create_time: undefined,
        status: 0,
        app_status: '',
        app_id: 10000
      }
    },
    handleTempTime(data) {
      return {
        id: data.id,
        version: data.version,
        parent_version: data.parent_version,
        platform: data.platform ? data.platform : [],
        approval_time: data.approval_time ? new Date(data.approval_time * 1000) : new Date(),
        lock_time: data.lock_time ? new Date(data.lock_time * 1000) : new Date(),
        gray_time: data.gray_time ? new Date(data.gray_time * 1000) : new Date(),
        create_time: data.create_time ? new Date(data.create_time * 1000) : new Date(),
        release_time: data.release_time ? new Date(data.release_time * 1000) : 0,
        status: data.status,
        app_status: data.app_status,
        app_id: data.app_id
      }
    },
    removeAppVersion(data) {
      removeAppVersionRequest(data.id).then(() => {
        this.getList()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          if (this.temp.approval_time > this.temp.lock_time) {
            this.$message.error('立项时间必须早于锁版时间')
            return
          }
          if (this.temp.lock_time > this.temp.gray_time) {
            this.$message.error('锁版时间必须早于灰度时间')
            return
          }

          addAppVersionRequest(
            this.listQuery.app_id,
            this.temp.version,
            this.temp.parent_version,
            this.temp.platform,
            this.temp.approval_time.valueOf() / 1000,
            this.temp.lock_time.valueOf() / 1000,
            this.temp.gray_time.valueOf() / 1000).then((response) => {
            this.list.unshift(response.app)
            this.dialogFormVisible = false
            this.resetTemp()
            this.$notify({
              title: '成功',
              message: '创建成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    updateStatus(data) {
      updateStatusAppVersionRequest(data.id, data.status + 1).then(() => {
        this.getList()
        this.$notify({
          title: '成功',
          message: '创建成功',
          type: 'success',
          duration: 2000
        })
      })
    },
    updateDate() {
      if (this.dialogEditCount < 1) {
        this.dialogFormVisible = false
        return
      }
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          let gray_time = parseInt('0')
          let lock_time = parseInt('0')
          let approval_time = parseInt('0')
          let release_time = parseInt('0')
          let parent_version = this.temp.parent_version

          this.temp.lock_time = this.temp.lock_time.getTime() / 1000
          this.temp.gray_time = this.temp.gray_time.getTime() / 1000
          this.temp.approval_time = this.temp.approval_time.getTime() / 1000
          if (this.temp.status < 2) {
            approval_time = this.temp.approval_time
          }
          if (this.temp.status < 3) {
            lock_time = this.temp.lock_time
            if (this.temp.approval_time > this.temp.lock_time) {
              this.$message.error('立项时间必须早于锁版时间')
              return
            }
          }
          if (this.temp.status < 4) {
            gray_time = this.temp.gray_time
            if (this.temp.lock_time > this.temp.gray_time) {
              this.$message.error('锁版时间必须早于灰度时间')
              return
            }
          }
          if (this.temp.status == 5) {
            if (this.temp.gray_time > this.temp.release_time && this.temp.status >= 4) {
              this.$message.error('灰度时间必须早于发布时间')
              return
            }
            this.temp.release_time = this.temp.release_time.valueOf() / 1000
            release_time = this.temp.release_time
          }

          if (this.temp.parent_version === '-') {
            parent_version = ''
          }

          updateAppVersionRequest(
            this.temp.id,
            this.temp.app_id,
            this.temp.version,
            parent_version,
            this.temp.platform,
            approval_time,
            lock_time,
            gray_time,
            release_time).then(() => {
            for (const v of this.list) {
              if (v.id === this.temp.id) {
                const index = this.list.indexOf(v)
                this.list.splice(index, 1, this.temp)
                break
              }
            }
            this.resetTemp()
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '修改成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handleCreate() {
      if (this.dialogStatus === 'update') {
        this.resetTemp()
      }
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleUpdate(data) {
      this.temp = this.handleTempTime(data)
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.dialogEditCount = -1
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    sortChange(data) {
      const { prop, order } = data
      if (prop === 'id') {
        this.sortByID(order)
      }
    },
    sortByID(order) {
      if (order === 'ascending') {
        this.listQuery.sort = '+_id'
      } else {
        this.listQuery.sort = '-_id'
      }
      this.handleFilter()
    },
    handleFilter() {
      this.currentSimpleApp = this.simpleAppList.filter(item => item.id == this.listQuery.app_id)[0]
      this.listQuery.page = 1
      this.getList()
    }
  }
}
</script>

<style lang="scss" scoped>
  .fixed-container{
    background-color: #eef1f6;
    .icon-add{
      width: 2rem;
      height: 1.9rem;
      background-size: 2rem 1.9rem;
    }
  }

</style>
