<template>
  <div class="app-container">
    <eHeader :query="query"/>
    <!--表格渲染-->
    <el-table v-loading="loading" :data="data" size="small" border style="width: 100%;">
      <el-table-column prop="user_id" label="用户ID"/>
      <el-table-column prop="client_ip" label="IP"/>
      <el-table-column :show-overflow-tooltip="true" prop="method" label="方法名称"/>
      <el-table-column :show-overflow-tooltip="true" prop="params" label="参数"/>
      <el-table-column :show-overflow-tooltip="true" prop="exceptionDetail" label="异常堆栈信息"/>
      <el-table-column prop="latency" label="请求耗时" align="center">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.latency/1000000 <= 300">{{ scope.row.latency/1000000 }}ms</el-tag>
          <el-tag v-else-if="scope.row.latency/1000000 <= 1000" type="warning">{{ scope.row.latency/1000000 }}ms</el-tag>
          <el-tag v-else type="danger">{{ scope.row.latency/1000000 }}ms</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="flag" label="日志类型" width="100px" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.flag === 'ERROR'" class="badge badge-bg-orange">{{ scope.row.flag }}</span>
          <span v-else class="badge">{{ scope.row.logType }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="createTime" label="创建日期" width="180px">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createTime) }}</span>
        </template>
      </el-table-column>
    </el-table>
    <!--分页组件-->
    <el-pagination
      :total="total"
      style="margin-top: 8px;"
      layout="total, prev, pager, next, sizes"
      @size-change="sizeChange"
      @current-change="pageChange"/>
  </div>
</template>

<script>
import initData from '../../../mixins/initData'
import { parseTime } from '@/utils/index'
import eHeader from './module/header'
export default {
  components: { eHeader },
  mixins: [initData],
  created() {
    this.$nextTick(() => {
      this.init()
    })
  },
  methods: {
    parseTime,
    beforeInit() {
      this.url = 'log/list'
      const sort = 'id,desc'
      const query = this.query
      const user_id = query.user_id
      const flag = query.flag
      this.params = { page: this.page, size: this.size, sort: sort }
      if (user_id && user_id) { this.params['user_id'] = user_id }
      if (flag !== '' && flag !== null) { this.params['flag'] = flag }
      return true
    }
  }
}
</script>

<style scoped>

</style>
