<template>
  <div class="app-container">
    <!--悬浮添加按钮-->
    <section class="content">
      <fixed-button :bottom="3" @clickEvent="handleCreate" class="fixed-container">
        <svg-icon icon-class="add" class="icon-add"></svg-icon>
      </fixed-button>
    </section>






    </div>
</template>

<script>
  import fixedButton from '../../components/FixedButton';
  export default {
    name: "MetaData",
    components: {
      fixedButton
    },
    methods: {
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




</style>
