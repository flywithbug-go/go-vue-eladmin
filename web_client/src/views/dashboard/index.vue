<template>
  <section class="content">
    <fixed-icon :bottom="3" @clickEvent="chat" class="fixed-container">
      <svg-icon icon-class="add" class="icon-chat"></svg-icon>
    </fixed-icon>
  </section>
</template>


<script>
  import fixedIcon from '../../components/FixedIcon';

  export default {
    name: 'test',
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
      document.title = 'Vue制作固定定位图标滑动隐藏效果';
    },
  };

</script>


<style scoped lang="scss">
  .content{
    height: 100vh;

  }
  .fixed-container{
    background-color: #eef1f6;
  }
  .icon-chat{
    width: 2rem;
    height: 1.9rem;
    background-size: 2rem 1.9rem;
  }
</style>

