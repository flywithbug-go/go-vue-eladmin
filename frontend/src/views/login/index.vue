<template>
    <div class="login-container">
      <el-form ref="loginForm"
               :model="loginForm"
               class="login-form"
               auto-complete="on"
               label-position="left">
        <div class="title-container">
          <lang-select class="set-language"></lang-select>
          <h3 class="title" >{{ $t('login.title') }}</h3>
        </div>
        <el-form-item prop="account">
          <span class="svg-container">
            <svg-icon icon-class="user"></svg-icon>
          </span>
          <el-input v-model="loginForm.account"
                    :placeholder="$t('login.username')"
                    name="account"
                    type="text"
                    auto-complete="on" ></el-input>
        </el-form-item>

        <el-form-item prop="password">
          <span class="svg-container">
            <svg-icon icon-class="password"/>
          </span>
          <el-input :type="passwordType"
                    v-model="loginForm.password"
                    :placeholder="$t('login.password')"
                    auto-complete="on"
                    @keyup.enter.native="handleLogin"></el-input>
          <span class="show-pwd" @click="showPwd">
            <svg-icon :icon-class="pwdIcon"/>
          </span>
        </el-form-item>
        <el-button :loading="loading"
                   type="primary"
                   :disabled="loginBtnDisable"
                   style="width: 100%; margin-bottom: 30px;"
                   @click.native.prevent="handleLogin">
          {{ $t('login.logIn') }}
        </el-button>


      </el-form>

    </div>
</template>

<script>
  import LangSelect from '../../components/LangSelect/'
export default {
  name: "Login",
  components: {
    LangSelect,
  },
  data() {
    const validateUsername = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('Please enter the correct user name'))
      } else {
        callback()
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 6) {
        callback(new Error('The password can not be less than 6 digits'))
      } else {
        callback()
      }
    }
    return {
      loginForm: {
        account: '',
        password: ''
      },
      loginRules: {
        account: [{ required: true, trigger: 'blur', validator: validateUsername }],
        password: [{ required: true, trigger: 'blur', validator: validatePassword }]
      },
      loginDisable:true,
      passwordType: 'password',
      pwdIcon:'eye',
      loading: false,
      showDialog: false,
      redirect: undefined
    }
  },
  computed: {
    loginBtnDisable() {
      if (this.loginForm.password.length < 4 || this.loginForm.account.length < 4){
        return true
      }
      return false
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    },
  },
  methods: {
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
        this.pwdIcon = 'eye_open'
      } else {
        this.passwordType = 'password'
        this.pwdIcon = 'eye'
      }
    },
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid){
          this.loading = true
          this.$store.dispatch('LoginByAccount',this.loginForm).then(()=> {
            console.log('did login',this.redirect)
            this.loading = false
            this.$router.push({path: this.redirect || '/'})
          }) .catch(() => {
            this.loading = false
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>


<style rel="stylesheet/scss" lang="scss">
  /* 修复input 背景不协调 和光标变色 */
  /* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

  $bg:#283443;
  $light_gray:#eee;
  $cursor: #fff;

  @supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
    .login-container .el-input input{
      color: $cursor;
      &::first-line {
        color: $light_gray;
      }
    }
  }

  /* reset element-ui css */
  .login-container {
    .el-input {
      display: inline-block;
      height: 47px;
      width: 85%;
      input {
        background: transparent;
        border: 0px;
        -webkit-appearance: none;
        border-radius: 0px;
        padding: 12px 5px 12px 15px;
        color: $light_gray;
        height: 47px;
        caret-color: $cursor;
        &:-webkit-autofill {
          -webkit-box-shadow: 0 0 0px 1000px $bg inset !important;
          -webkit-text-fill-color: $cursor !important;
        }
      }
    }
    .el-form-item {
      border: 1px solid rgba(255, 255, 255, 0.1);
      background: rgba(0, 0, 0, 0.1);
      border-radius: 5px;
      color: #454545;
    }
  }
</style>

<style rel="stylesheet/scss" lang="scss" scoped>
  $bg:#2d3a4b;
  $dark_gray:#889aa4;
  $light_gray:#eee;

  .login-container {
    position: fixed;
    height: 100%;
    width: 100%;
    background-color: $bg;
    .login-form {
      position: absolute;
      left: 0;
      right: 0;
      width: 520px;
      max-width: 100%;
      padding: 35px 35px 15px 35px;
      margin: 120px auto;
    }
    .tips {
      font-size: 14px;
      color: #fff;
      margin-bottom: 10px;
      span {
        &:first-of-type {
          margin-right: 16px;
        }
      }
    }
    .svg-container {
      padding: 6px 5px 6px 15px;
      color: $dark_gray;
      vertical-align: middle;
      width: 30px;
      display: inline-block;
    }
    .title-container {
      position: relative;
      .title {
        font-size: 26px;
        color: $light_gray;
        margin: 0px auto 40px auto;
        text-align: center;
        font-weight: bold;
      }
      .set-language {
        color: #fff;
        position: absolute;
        top: 5px;
        right: 0px;
      }
    }
    .show-pwd {
      position: absolute;
      right: 10px;
      top: 7px;
      font-size: 16px;
      color: $dark_gray;
      cursor: pointer;
      user-select: none;
    }
    .thirdparty-button {
      position: absolute;
      right: 35px;
      bottom: 28px;
    }
  }
</style>
