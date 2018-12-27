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

      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">

        <el-form-item align="center">
          <label>上传图标</label>
        </el-form-item>
        <el-form-item  align="center" >
          <el-upload
            class="avatar-uploader"
            :action="actionURL"
            :show-file-list="false"
            :headers="headers"
            :on-success="handleAvatarSuccess"
            :before-upload="beforeAvatarUpload">
            <img v-if="temp.icon" :src="temp.icon" width="100%" class="avatar">
            <img v-else="temp.icon" :src="imagePlaceHolder" width="100%" class="avatar">
          </el-upload>
        </el-form-item>

        <el-form-item :label="$t('application.table_name')" prop="name">
          <el-input v-model="temp.name"/>
        </el-form-item>
        <el-form-item :label="$t('application.table_bundleId')"  prop="bundleId">
          <el-input v-model="temp.bundleId" :placeholder="$t('application.table_bundleId_placeHolder')"/>
        </el-form-item>
        <el-form-item :label="$t('application.table_desc')" prop="desc">
          <el-input :autosize="{ minRows: 2, maxRows: 4}" type="textarea" v-model="temp.desc" :placeholder="$t('application.table_desc_placeholder')"/>
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
  import fixedButton from '../../components/FixedButton';
  import global_ from '../../config'
  import store from '@/store'
  import { generateAppTitle } from '../../utils/i18n'
  export default {
    name: 'AppManager',
    components: {
      fixedButton
    },
    data() {
      return {
        listLoading: true,
        tableKey: 0,
        headers: {'Authorization': store.getters.token},
        actionURL:global_.UploadImageURL,
        list: null,
        total: 0,
        dialogFormVisible:false,
        dialogStatus:'create',
        temp: {
          id: undefined,
          name: '',
          owner:'',
          desc: '',
          icon:'',
          bundleId:''
        },
        rules: {
          name: [
            {
              required: true,
              message: this.$t('application.table_name_placeHolder'),
            },
            {
              max: 10,
              message: '请输入不多于10个字符',
            },
            {
              min: 2,
              message: '请输入不少于2个字符',
            },
            {
              pattern: /^[A-Za-z\u4e00-\u9fa5]+$/,
              message: '只允许输入汉字或者英文字母'
            }
          ],
          bundleId: [
            {
              required: true,
              message: this.$t('application.table_bundleId_warning'),
            },
            {
              pattern: /^(com).[A-Za-z0-9.-]+(.)[A-Za-z0-9.-]+$/,
              message: '(格式：com.xxx.xxx)'
            },
            {
              min: 10,
              message: '请输入不少于10个字符',
            },
            {
              max: 30,
              message: '请输入不多于30个字符',
            },
          ],
          desc: [
            {
              required: true,
              message: this.$t('application.table_desc_placeHolder'),
            },
            {
              min: 10,
              message: '请输入不少于10个字符',
            },
            {
              max: 200,
              message: '请输入不多于200个字符',
            },
          ],
        },
        imagePlaceHolder:require('../../assets/image_placeholder.png'),
      }
    },
    computed: {
      rulesName() {
        return "请输入应用名称"
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
        this.dialogStatus = 'create'
        this.dialogFormVisible =  true
        this.$nextTick(() => {
          this.$refs['dataForm'].clearValidate()
        })
      },
      createData() {
        if (this.temp.icon === ''){
          this.$message.error('icon should upload first')
          return
        }
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {



          }
        })
      },
      updateDate() {

      },
      resetTemp() {
        this.temp = {
          id: undefined,
          name: '',
          owner:'',
          time: new Date(),
          desc: '',
          bundleId:''
        }
      },
      handleAvatarSuccess(res) {
        if (res.code === 200) {
          this.temp.icon=global_.downloadImageURL + res.data["imagePath"]
        } else {
          this.$message.error(res.msg);
        }
      },
      beforeAvatarUpload(file) {
        const isLt10M = file.size / 1024 / 1024 < 10;
        if (!isLt10M) {
          this.$message.error('上传头像图片大小不能超过 10MB!');
          return false
        }
        const isSize = new Promise(function (resolve, reject) {
          let _URL = window.URL || window.webkitURL;
          let img = new Image();
          img.onload = function () {
            let valid = img.width === img.height
            valid ? resolve() : reject()
          }
          img.src = _URL.createObjectURL(file)
        }).then(() => {
          return file;
        }, () => {
          this.$message.error('图片格式宽高比必须是1:1');
          return Promise.reject();
        });
        return  isLt10M && isSize;
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
</style>
