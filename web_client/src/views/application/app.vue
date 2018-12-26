<template>
  <div class="app-container">
    <!--悬浮按钮-->
    <section class="content">
      <fixed-button :bottom="3" @clickEvent="addAction" class="fixed-container">
        <svg-icon icon-class="add" class="icon-add"></svg-icon>
      </fixed-button>
    </section>
<!--列表内容-->
    <el-table v-loading="listLoading"
              :key="tableKey"
              :data="list"
              border
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
                       width="150px" >
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

      <el-table-column :label="$t('application.table_desc')"
                       prop="id"
                       align="center"
                       min-width="150px">
        <template slot-scope="scope">
          <span> {{ scope.row.desc }} </span>
        </template>
      </el-table-column>
    </el-table>

    <!--创建弹窗-->
    <el-dialog :title="$t('application.table_createTitle')" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
        <el-form-item :label="$t('application.table_name')" prop="name">
          <el-input v-model="temp.name"/>
        </el-form-item>
        <el-form-item :label="$t('application.table_desc')" prop="desc">
          <el-input :autosize="{ minRows: 2, maxRows: 4}" type="textarea" v-model="temp.desc" :placeholder="$t('application.table_desc_placeholder')"/>
        </el-form-item>
      </el-form>

    </el-dialog>


  </div>
</template>

<script>
  import fixedButton from '../../components/FixedButton';

  export default {
  name: 'AppManager',
  components: {
    fixedButton
  },
  data() {
    return {
      listLoading: true,
      tableKey: 0,
      list: null,
      total: 0,
      dialogFormVisible:false,
      temp: {
        id: undefined,
        name: '',
        icon: '',
        owner:'',
        time: new Date(),
        desc: ''
      },
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
    addAction() {
      this.resetTemp()
      this.dialogFormVisible =  true

      console.log('added');
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        name: '',
        icon: '',
        owner:'',
        time: new Date(),
        desc: ''
      }
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
