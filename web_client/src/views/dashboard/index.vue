<template>
  <section class="content">
    <fixed-icon :bottom="3" @clickEvent="chat" class="fixed-container">
      <svg-icon icon-class="add" class="icon-add"></svg-icon>
    </fixed-icon>
  </section>
</template>

<script>
  import fixedIcon from '../../components/FixedButton';
  export default {
    name: 'Dashboard',
    components: {
      fixedIcon
    },
    data () {
      return {
      };
    },
    methods: {
      chat() { // 图标点击事件
        console.log('你好');
      },
    },
    mounted() {
      document.title = 'Dashboard';
    },
  };

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

