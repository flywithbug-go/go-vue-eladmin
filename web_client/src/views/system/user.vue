<template>
  <div class="app-container">

    <!--<div class="filter-container">-->
      <!--<el-input-->
        <!--:placeholder="$t('application.table_name')"-->
        <!--v-model="listQuery.name"-->
        <!--style="width: 200px;"-->
        <!--class="filter-item"-->
        <!--@keyup.enter.native="handleFilter"/>-->

      <!--<el-button-->
        <!--v-waves-->
        <!--class="filter-item"-->
        <!--type="primary"-->
        <!--icon="el-icon-search"-->
        <!--@click="handleFilter">-->
        <!--{{ $t('application.table_search') }}-->
      <!--</el-button>-->
      <!--<el-button-->
        <!--class="filter-item"-->
        <!--style="margin-left: 10px;"-->
        <!--type="primary"-->
        <!--icon="el-icon-edit"-->
        <!--@click="handleCreate">{{ $t('table.add') }}</el-button>-->
    <!--</div>-->

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
        :label="$t('table.id')"
        prop="id"
        sortable="custom"
        align="center"
        width="80px">
        <template slot-scope="scope">
          <span> {{ scope.row.id }} </span>
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('system.name')"
        prop="id"
        align="center"
        width="160px">
        <template slot-scope="scope">
          <span style="color: #4a9ff9; font-weight: bolder;font-size: 18px;">
            {{ formatUndefine(scope.row.name) }}
          </span>
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('system.avatar')"
        prop="id"
        align="center"
        width="100px">
        <template slot-scope="scope" >
          <img :src="scope.row.avatar" class="app-icon" width="auto" align="center">
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('system.title')"
        prop="id"
        align="center"
        width="160px">
        <template slot-scope="scope">
          <span >
            {{ formatUndefine(scope.row.title) }}
          </span>
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('system.phone')"
        prop="id"
        align="center"
        width="160px">
        <template slot-scope="scope">
          <span >
            {{ formatUndefine(scope.row.phone) }}
          </span>
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('system.email')"
        prop="id"
        align="center"
        width="200px">
        <template slot-scope="scope">
          <span >
            {{ formatUndefine(scope.row.email) }}
          </span>
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('system.status')"
        prop="id"
        align="center"
        width="160px">
        <template slot-scope="scope">
          <span style="color: #4a9ff9; font-weight: bolder;font-size: 18px;">
            {{ formatUndefine(scope.row.status) }}
          </span>
        </template>
      </el-table-column>

      <el-table-column
        :label="$t('system.note')"
        prop="id"
        align="center"
        min-width="160px">
        <template slot-scope="scope">
          <span style="color: #4a9ff9; font-weight: bolder;font-size: 18px;">
            {{ formatUndefine(scope.row.note) }}
          </span>
        </template>
      </el-table-column>

    </el-table>

  </div>
</template>

<script>
import waves from '@/directive/waves' // Waves directive
import fixedButton from '../../components/FixedButton'
import Pagination from '../../components/Pagination'
import ElTableFooter from 'element-ui'

import { getUserListInfoRequest } from '../../api/user'

export default {
  name: 'AppManager',
  components: {
    ElTableFooter,
    fixedButton,
    Pagination
  },
  directives: { waves },
  data() {
    return {
      listLoading: true,
      list: null,
      total: 0,
      listQuery: {
        page: 0,
        limit: 10,
        name: '',
        status: '',
        sort: '+_id'
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    formatUndefine(obj) {
      if (obj) {
        return obj
      }
      return '-'
    },
    getList() {
      getUserListInfoRequest(this.listQuery).then(response => {
        this.list = response.list
        this.total = response.total
        this.listLoading = false
      }).catch(() => {
        this.listLoading = false
      })
    },
    sortChange() {

    },
    handleFilter() {

    },
    handleCreate() {

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
