<template>
  <div class="app-container">

    <el-table v-loading="listLoading"
              :key="tableKey"
              :data="list"
              border="true"
              fit
              highlight-current-row
              style="width: 100%;"
              @sort-change="sortChange">
      <el-table-column :label="$t('application.table_id')"
                       prop="id"
                       sortable="custom"
                       align="center"
                       width="65">
        <template slot-scope="scope">
          <span> {{ scope.row.id }} </span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('application.table_name')"
                       prop="id"
                       align="center"
                       width="150px">
        <template slot-scope="scope">
          <span> {{ scope.row.name }} </span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('application.table_icon')"
                       prop="id"
                       align="center"
                       width="150px">
        <template slot-scope="scope">
          <span> {{ scope.row.icon }} </span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('application.table_desc')"
                       prop="id"
                       align="center"
                       width="150px">
        <template slot-scope="scope">
          <span> {{ scope.row.desc }} </span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('application.table_owner')"
                       prop="id"
                       align="center"
                       width="150px">
      <template slot-scope="scope">
        <span> {{ scope.row.owner }} </span>
      </template>
      </el-table-column>

      <el-table-column :label="$t('application.table_time')"
                       prop="id"
                       align="center"
                       width="150px">
        <template slot-scope="scope">
          <span> {{ scope.row.time }} </span>
        </template>
      </el-table-column>

    </el-table>
  </div>
</template>

<script>
export default {
  name: 'AppManager',
  data() {
    return {
      listLoading: true,
      tableKey: 0,
      list: null,
      total: 0,
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      setTimeout(() => {
        this.listLoading = false
      }, 1.5 * 1000)
    },
    sortChange(data) {
      const { prop, order } = data
      if (prop === 'id') {
        this.sortByID(order)
      }
    },
  }
}
</script>

<style scoped>

</style>
