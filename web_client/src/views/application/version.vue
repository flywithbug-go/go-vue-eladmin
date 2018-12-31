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
