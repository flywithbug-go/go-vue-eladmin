<template>
  <div>
    <el-button v-if="data.id != 10000"
               size="mini"
               type="success"
               @click="to">
      {{ $t('actions.edit') }}
    </el-button>
    <eForm ref="form"
           :permissions="permissions"
           :sup_this="sup_this"
           :is-add="false"/>
  </div>
</template>
<script>
import eForm from './form'
export default {
  components: { eForm },
  props: {
    data: {
      type: Object,
      required: true
    },
    sup_this: {
      type: Object,
      required: true
    },
    permissions: {
      type: Array,
      required: true
    }
  },
  methods: {
    to() {
      console.log("permissions:", this.permissions)
      const _this = this.$refs.form
      _this.permissionIds = []
      _this.form = { id: this.data.id, name: this.data.name, remark: this.data.alias, permissions: [] }
      this.data.permissions.forEach(function(data, index) {
        _this.permissionIds.push(data.id)
      })
      _this.dialog = true
    }
  }
}
</script>

<style scoped>
  div{
    display: inline-block;
    margin-right: 3px;
  }
</style>
