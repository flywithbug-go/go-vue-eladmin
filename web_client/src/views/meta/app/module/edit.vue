<template>
  <div>
    <el-button :disabled="disabled" size="mini" type="success" @click="to">{{ $t('actions.edit') }}</el-button>
    <eForm ref="form" :sup_this="sup_this" :is-add="false"/>
  </div>
</template>
<script>
import eForm from './form'
export default {
  components: { eForm },
  props: {
    disabled: {
      type: Boolean,
      required: false,
      default: false
    },
    data: {
      type: Object,
      required: true
    },
    // index.vue 的this 可用于刷新数据
    sup_this: {
      type: Object,
      required: true
    }
  },
  methods: {
    to() {
      console.log('editTo: ', this.data)
      const _this = this.$refs.form
      _this.form = {
        id: this.data.id,
        name: this.data.name,
        owner: this.data.owner,
        desc: this.data.desc,
        icon: this.data.icon,
        bundle_id: this.data.bundle_id
      }
      _this.dialog = true
    }

  }
}
</script>

<style scoped>
  div{display: inline-block;margin-right: 3px;}
</style>
