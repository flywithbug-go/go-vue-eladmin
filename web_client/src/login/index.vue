<template>
  <div class="login-container"
       style="background-color: #141a48;margin: 0px;overflow:hidden;">

  </div>
</template>

<script>
  import { isWscnEmail } from 'utils/validate';
  export default {
    name: 'login',
    data() {
      const validateEmail = (rule, value, callback) => {
        if (!isWscnEmail(value)) {
          callback(new Error('请输入正确的邮箱地址'));
        } else {
          callback()
        }
      }
      const validatePass = (rule, value, callback) => {
        if (value.length < 6) {
          callback(new Error('密码不能小于6位'));
        } else {
          callback()
        }
      }
      return {
        loginForm: {
          email: '',
          password: ''
        },
        loginRules: {
          email: [
            {required: true, trigger: 'blur', validator: validateEmail}
          ],
          password: [
            {required: true, trigger: 'blur', validator: validatePass}
          ]
        },
        loading: false,
        showDialog: false
      }
    },
  }
</script>

<style scoped>



</style>
