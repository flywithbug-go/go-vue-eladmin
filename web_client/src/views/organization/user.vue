<template>
  <div class="app-container">





  </div>
</template>

<script>
import waves from '@/directive/waves' // Waves directive
import fixedButton from '../../components/FixedButton'
import Pagination from '../../components/Pagination'

export default {
  name: 'AppManager',
  components: {
    fixedButton,
    Pagination
  },
  data() {
    return {
      listQuery: {
        page: 0,
        limit: 10,
        name: '',
        status: '',
        sort: '+_id'
      },

    }
  },
  directives: { waves },

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
