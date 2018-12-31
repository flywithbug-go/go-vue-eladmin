<template>
  <div class="app-container">
    <!--悬浮添加按钮-->
    <section class="content">
      <fixed-button :bottom="3" class="fixed-container" @clickEvent="handleCreate">
        <svg-icon icon-class="add" class="icon-add"/>
      </fixed-button>
    </section>

    <el-table
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      header-row-class-name="center"
      @sort-change="sortChange">
      <el-table-column :label="$t('table.id')" prop="id" sortable="custom" align="center" width="65">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('appVersion.versionN')" align="center" min-width="90px">
        <template slot-scope="scope">
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.parentVN')" align="center" min-width="90px">
        <template slot-scope="scope">
          <span>{{ scope.row.parent_version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.platform')" align="center" min-width="150px">
        <template slot-scope="scope">
          <span>{{ formatPlatform(scope.row.platform) }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.approvalTime')" align="center" min-width="150px">
        <template slot-scope="scope">
          <span>{{ formatDate(scope.row.approval_time) }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.lockTime')" align="center" min-width="150px">
        <template slot-scope="scope">
          <span>{{ formatDate(scope.row.lock_time) }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.grayTime')" align="center" min-width="150px">
        <template slot-scope="scope">
          <span>{{ formatDate(scope.row.gray_time) }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.status')" align="center" min-width="80px">
        <template slot-scope="scope">
          <span>{{ scope.row.app_status }}</span>
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('application.table_action')"
        align="center"
        min-width="150px"
        class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-show="scope.row.status === 1"
            type="primary"
            size="mini"
            @click="handleUpdate(scope.row)">
            {{ $t('appVersion.operate') }}
          </el-button>

          <el-button
            v-show="scope.row.status === 1"
            type="warning"
            size="mini"
            @click="handleUpdate(scope.row)">
            {{ $t('selector.develop') }}
          </el-button>

          <el-button
            v-show="scope.row.status === 2"
            type="warning"
            size="mini"
            @click="handleUpdate(scope.row)">
            {{ $t('selector.gray') }}
          </el-button>
          <el-button
            v-show="scope.row.status === 3"
            type="warning"
            size="mini"
            @click="handleUpdate(scope.row)">
            {{ $t('selector.releasing') }}
          </el-button>

        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.createTime')" align="center" min-width="150px">
        <template slot-scope="scope">
          <span>{{ formatDate(scope.row.create_time) }}</span>
        </template>
      </el-table-column>

    </el-table>

    <el-dialog
      :title="textMap[dialogStatus]"
      :visible.sync="dialogFormVisible">
      <el-form
        ref="dataForm"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="80px"
        style="width: 400px; margin-left:50px;">

        <el-form-item
          :label="$t('appVersion.versionN')"
          prop="version">
          <el-input
            :disabled="dialogStatus==='update' && temp.status > 1"
            v-model="temp.version"
            :placeholder="$t('appVersion.versionN')"/>
        </el-form-item>

        <el-form-item
          :label="$t('appVersion.parentVN')"
          prop="parent_version">
          <el-input
            :disabled="dialogStatus==='update' && temp.status > 1"
            v-model="temp.parent_version"
            :placeholder="$t('appVersion.parentVNPlaceholder')"/>
        </el-form-item>

        <el-form-item
          :label="$t('appVersion.approvalTime')"
          prop="approval_time">
          <el-date-picker
            :disabled="dialogStatus==='update' && temp.status > 1"
            v-model="temp.approval_time"
            :placeholder="$t('appVersion.approvalTime')"
            type="date"/>
        </el-form-item>

        <el-form-item
          :label="$t('appVersion.lockTime')"
          prop="lock_time">
          <el-date-picker
            :disabled="dialogStatus==='update' && temp.status > 1"
            v-model="temp.lock_time"
            format="yyyy-MM-dd"
            :placeholder="$t('appVersion.lockTime')"
            type="date"/>
        </el-form-item>

        <el-form-item
          :label="$t('appVersion.grayTime')"
          prop="gray_time">
          <el-date-picker
            :disabled="dialogStatus==='update' && temp.status > 1"
            v-model="temp.gray_time"
            format="yyyy-MM-dd"
            :placeholder="$t('appVersion.grayTime')"
            type="date"/>
        </el-form-item>

        <el-form-item :label="$t('appVersion.platform')" prop="platform">
          <el-select
            :disabled="dialogStatus==='update' && temp.status > 1"
            v-model="temp.platform"
            :placeholder="$t('selector.placeholder')"
            clearable
            multiple>
            <el-option
              v-for="item in platformOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"/>
          </el-select>
        </el-form-item>

        <!--<el-form-item-->
        <!--v-show="dialogStatus==='update'"-->
        <!--:label="$t('appVersion.status')"-->
        <!--prop="status">-->
        <!--<el-select-->
        <!--v-model="temp.app_status"-->
        <!--:placeholder="$t('selector.placeholder')">-->
        <!--<el-option-->
        <!--v-for="item in statusOptions"-->
        <!--:key="item.value"-->
        <!--:label="item.label"-->
        <!--:value="item.value"/>-->
        <!--</el-select>-->
        <!--</el-form-item>-->

      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{ $t('table.cancel') }}</el-button>
        <el-button type="primary" @click="dialogStatus==='create'? createData():updateDate()">{{ $t('table.confirm') }}</el-button>
      </div>

    </el-dialog>

  </div>
</template>

<script>
import fixedButton from '../../components/FixedButton'
import { getAppVersionListRequest } from '../../api/app'
import { formatDate } from '../../utils/date'
import ElTableFooter from 'element-ui'

export default {
  name: 'MetaData',
  components: {
    ElTableFooter,
    fixedButton
  },

  data() {
    return {
      listLoading: true,
      list: null,
      total: 10,
      dialogFormVisible: false,
      dialogStatus: 'create',
      textMap: {
        update: this.$t('application.table_edit'),
        create: this.$t('application.table_add')
      },
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
          },
        ],
        lock_time: [
          {
            required: true,
            message: '必选',
            trigger: 'change'
          },
        ],
        gray_time: [
          {
            required: true,
            message: '必选',
            trigger: 'change'
          },
        ],
        platform: [
          {
            required: true,
            message: '必选',
            trigger: 'change'
          },
        ]
      },
      temp: {
        id: 0,
        version: '',
        parent_version: '',
        platform: [],
        approval_time: new Date(),
        lock_time: new Date(),
        gray_time: new Date(),
        create_time: new Date(),
        status: 0,
        app_status: '',
        app_id: 0
      },
      listQuery: {
        page: 0,
        limit: 10,
        name: '',
        owner: '',
        sort: '-id'
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
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
        app_id: 0
      }
    },
    handleTempTime(data) {
      return {
        id: data.id,
        version: data.version,
        parent_version: data.parent_version,
        platform: data.platform ? data.platform : [],
        approval_time: data.approval_time?new Date(data.approval_time * 1000):new Date(),
        lock_time: data.lock_time? new Date(data.lock_time * 1000):new Date(),
        gray_time: data.gray_time? new Date(data.gray_time * 1000):new Date(),
        create_time: data.create_time? new Date(data.create_time * 1000):new Date(),
        status: data.status,
        app_status: data.app_status,
        app_id: data.app_id
      }
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          if (this.temp.approval_time > this.temp.lock_time) {
            this.$message.error("立项时间必须早于锁版时间")
            return
          }
          if (this.temp.lock_time > this.temp.gray_time) {
            this.$message.error("锁版时间必须早于灰度时间")
            return
          }

        }
      })
    },
    updateDate() {
      console.log('updateDate:', this.temp)
    },
    handleCreate() {
      this.resetTemp()
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
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    getList() {
      this.listLoading = true
      getAppVersionListRequest(this.listQuery).then(response => {
        this.list = response.list
        this.total = response.total
        this.listLoading = false
      }).catch(() => {
        this.listLoading = false
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
