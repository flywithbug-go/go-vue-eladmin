<template>
  <section class="content">
    <fixed-icon :bottom="3" class="fixed-container" @clickEvent="chat">
      <svg-icon icon-class="add" class="icon-add"/>
    </fixed-icon>
  </section>
</template>

<script>
import fixedIcon from '../../components/FixedButton'
export default {
  name: 'Dashboard',
  components: {
    fixedIcon
  },
  data() {
    return {
    }
  },
  mounted() {
    document.title = 'Dashboard'
  },
  methods: {
    chat() { // 图标点击事件
      console.log('你好')
    }
  }
}

</script>

<style scoped lang="scss">
  .content{
    height: 100vh;
  }
  .fixed-container{
    background-color: #eef1f6;
    .icon-add{
      width: 2rem;
      height: 1.9rem;
      background-size: 2rem 1.9rem;
    }
  }

</style>

