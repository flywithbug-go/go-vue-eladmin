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
      <el-table-column :label="$t('appVersion.versionN')" align="center" min-width="80px">
        <template slot-scope="scope">
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.parentVN')" align="center" min-width="80px">
        <template slot-scope="scope">
          <span>{{ scope.row.parent_version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.platform')" align="center" min-width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.platform }}</span>
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

      <el-table-column :label="$t('appVersion.createTime')" align="center" min-width="150px">
        <template slot-scope="scope">
          <span>{{ formatDate(scope.row.create_time) }}</span>
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

        <el-form-item
          :label="$t('appVersion.versionN')"
          prop="version">
          <el-input
            :disabled="dialogStatus==='update'"
            v-model="temp.version"
            :placeholder="$t('appVersion.versionN')"></el-input>
        </el-form-item>

        <el-form-item
          :label="$t('appVersion.parentVN')"
          prop="parent_version">
          <el-input
            :disabled="dialogStatus==='update'"
            v-model="temp.parent_version"
            :placeholder="$t('appVersion.parentVN')"></el-input>
        </el-form-item>

        <el-form-item :label="$t('appVersion.approvalTime')"
                      prop="timestamp">
          <el-date-picker v-model="temp.approval_time"
                          type="date"
                          :placeholder="$t('appVersion.approvalTime')">
          </el-date-picker>
        </el-form-item>

        <el-form-item :label="$t('appVersion.lockTime')"
                      prop="timestamp">
          <el-date-picker v-model="temp.lock_time"
                          type="date"
                          :placeholder="$t('appVersion.lockTime')">
          </el-date-picker>
        </el-form-item>

        <el-form-item :label="$t('appVersion.grayTime')"
                      prop="timestamp">
          <el-date-picker v-model="temp.gray_time"
                          type="date"
                          :placeholder="$t('appVersion.grayTime')">
          </el-date-picker>
        </el-form-item>


        <el-form-item :label="$t('appVersion.platform')" prop="platform">
          <el-select
            v-model="platformValues"
            multiple
            :placeholder="$t('selector.placeholder')">
            <el-option
              v-for="item in platformOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item :label="$t('appVersion.status')"
                      prop="status" v-show="dialogStatus==='update'">
          <el-select
            v-model="statusValues"
            :placeholder="$t('selector.placeholder')">
            <el-option
              v-for="item in statusOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>

      </el-form>

    </el-dialog>

  </div>
</template>

<script>
import fixedButton from '../../components/FixedButton'
import { getAppVersionListRequest } from '../../api/app'
import { formatDate } from '../../utils/date'

export default {
  name: 'MetaData',
  components: {
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
      statusValues: [],
      platformValues: [],
      rules: {

      },
      temp: {
        id: 0,
        version: '',
        parent_version: '',
        platform: '',
        approval_time: '',
        lock_time: '',
        gray_time: '',
        create_time: '',
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
    formatDate(time) {
      if (!time || time === 0) {
        return '-'
      }
      const date = new Date(time * 1000)
      return formatDate(date, 'yyyy-MM-dd hh:mm')
    },
    handleCreate() {
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleUpdate(data) {
      this.temp = Object.assign({}, data) // copy obj
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
