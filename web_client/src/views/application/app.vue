<template>
  <div class="app-container">

    <!--悬浮添加按钮-->
    <section class="content">
      <fixed-button :bottom="3" class="fixed-container" @clickEvent="handleCreate">
        <svg-icon icon-class="add" class="icon-add"/>
      </fixed-button>
    </section>

    <div class="filter-container">
      <el-input
        :placeholder="$t('application.table_name')"
        v-model="listQuery.name"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"/>
      <el-input
        :placeholder="$t('application.table_owner')"
        v-model="listQuery.owner"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"/>
      <el-button
        v-waves
        class="filter-item"
        type="primary"
        icon="el-icon-search"
        @click="handleFilter">
        {{ $t('application.table_search') }}
      </el-button>
      <el-button
        class="filter-item"
        style="margin-left: 10px;"
        type="primary"
        icon="el-icon-edit"
        @click="handleCreate">{{ $t('table.add') }}</el-button>
    </div>

    <!--列表内容-->
    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      header-row-class-name="center"
      @sort-change="sortChange">
      <el-table-column
        :label="$t('application.table_id')"
        prop="id"
        sortable="custom"
        align="center"
        width="65">
        <template slot-scope="scope">
          <span> {{ scope.row.id }} </span>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('application.table_name_bundleId')"
        prop="id"
        align="center"
        width="160px">
        <template slot-scope="scope">
          <span style="color: #4a9ff9; font-weight: bolder;font-size: 18px;">
            {{ scope.row.name }}
          </span>
          <br>
          <span style="color: #2d2d2d; font-weight: bolder;font-size: 10px;">
            {{ scope.row.bundle_id }}
          </span>
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('application.table_icon')"
        prop="id"
        align="center"
        width="100px">
        <template slot-scope="scope" >
          <img :src="scope.row.icon" class="app-icon" width="auto" align="center">
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('application.table_owner')"
        prop="id"
        align="center"
        width="150px">
        <template slot-scope="scope">
          <span> {{ scope.row.owner }} </span>
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('application.table_create_time')"
        prop="id"
        align="center"
        width="150px">
        <template slot-scope="scope">
          <span> {{ formatDate(scope.row.time) }} </span>
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('application.table_desc')"
        prop="id"
        align="center"
        min-width="150px">
        <template slot-scope="scope">
          <span> {{ scope.row.desc }} </span>
        </template>
      </el-table-column>
      <el-table-column
        :label="$t('application.table_action')"
        align="center"
        width="100px"
        class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            type="primary"
            size="mini"
            @click="handleUpdate(scope.row)">
            {{ $t('application.table_edit') }}
          </el-button>
        </template>
      </el-table-column>

    </el-table>

    <!--分页-->
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getList"/>

    <!--创建弹窗添加修改-->
    <el-dialog
      :title="textMap[dialogStatus]"
      :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="70px"
        style="width: 400px; margin-left:50px;">
        <el-form-item align="center">
          <label>{{ $t('application.table_app_icon') }}</label>
        </el-form-item>
        <el-form-item align="center" >
          <el-upload
            :action="actionURL"
            :show-file-list="false"
            :headers="headers"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload"
            class="avatar-uploader">
            <img v-if="temp.icon" :src="temp.icon" width="100%" class="avatar">
            <img v-if="!temp.icon" :src="imagePlaceHolder" width="100%" class="avatar">
          </el-upload>
        </el-form-item>

        <el-form-item :label="$t('application.table_name')" prop="name">
          <el-input v-model="temp.name"/>
        </el-form-item>
        <el-form-item
          :label="$t('application.table_bundleId')"
          prop="bundle_id">
          <el-input
            :disabled="dialogStatus==='update'"
            v-model="temp.bundle_id"
            :placeholder="$t('application.table_bundleId_placeHolder')"/>
        </el-form-item>

        <el-form-item
          :label="$t('application.table_desc')"
          prop="desc">
          <el-input
            :autosize="{ minRows: 2, maxRows: 4}"
            :placeholder="$t('application.table_desc_placeholder')"
            v-model="temp.desc"
            type="textarea"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
        <el-button type="primary" @click="dialogStatus==='create'? createData():updateDate()">{{ $t('table.confirm') }}</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import waves from '@/directive/waves' // Waves directive
import fixedButton from '../../components/FixedButton'
import global_ from '../../config'
import store from '../../store'
import Pagination from '../../components/Pagination'
import { addApplicationRequest, getApplicationlListRequest, updateApplicationRequest } from '../../api/app'
import { formatDate } from '../../utils/date'

export default {
  name: 'AppManager',
  components: {
    fixedButton,
    Pagination
  },
  directives: { waves },
  data() {
    return {
      listLoading: true,
      headers: { 'Authorization': store.getters.token },
      actionURL: global_.UploadImageURL,
      imagePlaceHolder: require('../../assets/image_placeholder.png'),
      list: null,
      total: 10,
      dialogFormVisible: false,
      dialogStatus: 'create',
      dialogEditCount: 0,
      textMap: {
        update: this.$t('application.table_edit'),
        create: this.$t('application.table_add')
      },
      listQuery: {
        page: 0,
        limit: 2,
        name: '',
        owner: '',
        sort: '+_id'
      },
      temp: {
        id: 0,
        name: '',
        owner: '',
        desc: '',
        icon: '',
        time: '',
        bundle_id: '',
        app_id: ''
      },
      rules: {
        name: [
          {
            required: true,
            message: this.$t('application.table_name_placeHolder'),
            trigger: 'blur'
          },
          {
            max: 10,
            message: '请输入不多于10个字符',
            trigger: 'blur'
          },
          {
            min: 2,
            message: '请输入不少于2个字符',
            trigger: 'blur'
          },
          {
            pattern: /^[A-Za-z\u4e00-\u9fa5]+$/,
            message: '只允许输入汉字或者英文字母',
            trigger: 'blur'
          }
        ],
        bundle_id: [
          {
            required: true,
            message: this.$t('application.table_bundleId_warning'),
            trigger: 'blur'
          },
          {
            pattern: /^(com).[A-Za-z0-9.-]+(.)[A-Za-z0-9.-]+$/,
            message: '(格式：com.xxx.xxx)',
            trigger: 'blur'
          },
          {
            min: 10,
            message: '请输入不少于10个字符',
            trigger: 'blur'
          },
          {
            max: 30,
            message: '请输入不多于30个字符',
            trigger: 'blur'
          }
        ],
        desc: [
          {
            required: true,
            message: this.$t('application.table_desc_placeholder'),
            trigger: 'blur'
          },
          {
            min: 10,
            message: '请输入不少于10个字符',
            trigger: 'blur'
          },
          {
            max: 200,
            message: '请输入不多于200个字符',
            trigger: 'blur'
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
    this.getList()
  },
  methods: {
    formatDate(time) {
      const date = new Date(time * 1000)
      return formatDate(date, 'yyyy-MM-dd hh:mm')
    },
    getList() {
      this.listLoading = true
      getApplicationlListRequest(this.listQuery).then(response => {
        this.list = response.list
        this.total = response.total
        this.listLoading = false
      }).catch(() => {
        this.listLoading = false
      })
    },
    handleUpdate(row) {
      this.temp = Object.assign({}, row) // copy obj
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
      this.listQuery.page = 1
      this.getList()
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
    createData() {
      if (this.temp.icon === '') {
        this.$message.error(this.$t('application.table_app_icon_warning'))
        return
      }
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          addApplicationRequest(
            this.temp.bundle_id,
            this.temp.icon,
            this.temp.name,
            this.temp.desc).then(() => {
            this.list.unshift(this.temp)
            this.dialogFormVisible = false
            this.resetTemp()
            // this.getList()
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
    updateDate() {
      if (this.dialogEditCount < 1) {
        this.dialogFormVisible = false
        return
      }
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          updateApplicationRequest(this.temp.icon, this.temp.name, this.temp.desc, this.temp.id).then(() => {
            for (const v of this.list) {
              if (v.id === this.temp.id) {
                const index = this.list.indexOf(v)
                this.list.splice(index, 1, this.temp)
                break
              }
            }
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
    resetTemp() {
      this.temp = {
        id: undefined,
        name: '',
        owner: '',
        time: undefined,
        desc: '',
        bundle_id: ''
      }
    },
    handleAvatarSuccess(res) {
      if (res.code === 200) {
        this.temp.icon = global_.downloadImageURL + res.data['imagePath']
      } else {
        this.$message.error(res.msg)
      }
    },
    beforeAvatarUpload(file) {
      const isLt10M = file.size / 1024 / 1024 < 2
      if (!isLt10M) {
        this.$message.error('图片大小不能超过 2MB!')
        return false
      }
      // 限定宽高比
      const isSize = new Promise(function(resolve, reject) {
        const _URL = window.URL || window.webkitURL
        const img = new Image()
        img.onload = function() {
          const valid = img.width === img.height
          valid ? resolve() : reject()
        }
        img.src = _URL.createObjectURL(file)
      }).then(() => {
        return file
      }, () => {
        this.$message.error('图片宽高比需为1:1')
        return Promise.reject()
      })
      return isLt10M && isSize
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

  .avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;

  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 100px;
    height: 100px;
    line-height: 110px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
  .app-icon {
    display: inline-block;
    vertical-align: middle;
    width: 60px;
    height: 60px;
  }
</style>
