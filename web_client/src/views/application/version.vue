<template>
  <div class="app-container">
    <!--悬浮添加按钮-->
    <section class="content">
      <fixed-button :bottom="3" @clickEvent="handleCreate" class="fixed-container">
        <svg-icon icon-class="add" class="icon-add"></svg-icon>
      </fixed-button>
    </section>

    <el-table :data="list"
              border
              fit
              highlight-current-row
              style="width: 100%;"
              @sort-change="sortChange"
              header-row-class-name="center">
      <el-table-column :label="$t('table.id')" prop="id" sortable="custom" align="center" width="65">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('appVersion.versionN')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.parentVN')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.parent_version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.platform')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.platform }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.approvalTime')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.approval_time }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.lockTime')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.lock_time }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.grayTime')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.gray_time }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.status')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.app_status }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.createTime')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.create_time }}</span>
        </template>
      </el-table-column>

    </el-table>


    </div>
</template>

<script>
  import fixedButton from '../../components/FixedButton';
  import {getAppVersionlistRequest} from '../../api/app';
  import { formatDate } from '../../utils/date';

  export default {
    name: "MetaData",
    data() {
      return {
        listLoading: true,
        list: null,
        total: 10,
        listQuery: {
          page: 0,
          limit: 2,
          name: '',
          owner: '',
          sort: '+_id'
        },
      }
    },
    components: {
      fixedButton
    },
    created() {
      this.getList()
    },
    methods: {
      formatDate(time) {
        let date = new Date(time*1000);
        return formatDate(date, 'yyyy-MM-dd hh:mm');
      },
      handleCreate() {

      },
      getList() {
        this.listLoading = true
        getAppVersionlistRequest(this.listLoading).then(responst => {

        })
      },
      sortChange() {
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
