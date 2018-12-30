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
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.platform')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.approvalTime')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.lockTime')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.grayTime')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.status')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('appVersion.createTime')" align="center" width="150px">
        <template slot-scope="scope">
          <span>{{ scope.row.version }}</span>
        </template>
      </el-table-column>


    </el-table>




    </div>
</template>

<script>
  import fixedButton from '../../components/FixedButton';
  export default {
    name: "MetaData",
    data() {
      return {
        listLoading: true,
        list: null,
        total: 10,
      }
    },
    components: {
      fixedButton
    },
    methods: {
      handleCreate() {

      },
      sortChange() {

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
