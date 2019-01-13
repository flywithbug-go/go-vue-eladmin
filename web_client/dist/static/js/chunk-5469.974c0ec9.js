(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-5469"],{"99or":function(t,e,a){"use strict";a.r(e);var s=a("QbLZ"),i=a.n(s),o=a("L2JU"),r=a("7Qib"),n=a("Q2AE");var l=a("wk8/"),c={data:function(){var t=this;return{loading:!1,dialog:!1,title:"修改密码",form:{oldPass:"",newPass:"",confirmPass:""},rules:{oldPass:[{required:!0,validator:function(t,e,a){validPass(e).then(function(t){200===t.status?a():a(new Error("旧密码错误，请检查"))})},trigger:"blur"}],newPass:[{required:!0,message:"请输入新密码",trigger:"blur"},{min:6,max:20,message:"长度在 6 到 20 个字符",trigger:"blur"}],confirmPass:[{required:!0,validator:function(e,a,s){t.form.newPass!==a?s(new Error("两次输入的密码不一致")):s()},trigger:"blur"}]}}},methods:{cancel:function(){this.resetForm()},doSubmit:function(){var t=this;this.$refs.form.validate(function(e){if(!e)return!1;t.loading=!0,Object(l.i)(t.form.confirmPass,t.form.oldPass).then(function(e){t.resetForm(),t.$notify({title:"密码修改成功，请重新登录",type:"success",duration:1500}),setTimeout(function(){n.a.dispatch("LogOut").then(function(){location.reload()})},1500)}).catch(function(e){t.loading=!1,console.log(e.response.data.message)})})},resetForm:function(){this.dialog=!1,this.$refs.form.resetFields(),this.form={oldPass:"",newPass:"",confirmPass:""}}}},d=(a("c2zk"),a("KHd+")),u=Object(d.a)(c,function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticStyle:{display:"inline-block"}},[a("el-button",{staticClass:"button",attrs:{size:"mini",type:"info"},on:{click:function(e){t.dialog=!0}}},[t._v("修改")]),t._v(" "),a("el-dialog",{attrs:{visible:t.dialog,title:t.title,width:"500px"},on:{"update:visible":function(e){t.dialog=e},close:t.cancel}},[a("el-form",{ref:"form",attrs:{model:t.form,rules:t.rules,size:"small","label-width":"88px"}},[a("el-form-item",{attrs:{label:"旧密码",prop:"oldPass"}},[a("el-input",{staticStyle:{width:"370px"},attrs:{type:"password","auto-complete":"on"},model:{value:t.form.oldPass,callback:function(e){t.$set(t.form,"oldPass",e)},expression:"form.oldPass"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"新密码",prop:"newPass"}},[a("el-input",{staticStyle:{width:"370px"},attrs:{type:"password","auto-complete":"on"},model:{value:t.form.newPass,callback:function(e){t.$set(t.form,"newPass",e)},expression:"form.newPass"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"确认密码",prop:"confirmPass"}},[a("el-input",{staticStyle:{width:"370px"},attrs:{type:"password","auto-complete":"on"},model:{value:t.form.confirmPass,callback:function(e){t.$set(t.form,"confirmPass",e)},expression:"form.confirmPass"}})],1)],1),t._v(" "),a("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{attrs:{type:"text"},on:{click:t.cancel}},[t._v("取消")]),t._v(" "),a("el-button",{attrs:{loading:t.loading,type:"primary"},on:{click:t.doSubmit}},[t._v("确认")])],1)],1)],1)},[],!1,null,"65649c0f",null);u.options.__file="updatePass.vue";var m=u.exports;var f={props:{email:{type:String,required:!0}},data:function(){var t=this;return{loading:!1,dialog:!1,title:"修改邮箱",form:{pass:"",email:"",code:""},user:{email:"",password:""},codeLoading:!1,codeData:{type:"email",value:""},buttonName:"获取验证码",isDisabled:!1,time:60,rules:{pass:[{required:!0,validator:function(t,e,a){Object(l.j)(e).then(function(t){200===t.status?a():a(new Error("密码错误，请重新输入"))})},trigger:"blur"}],email:[{required:!0,validator:function(e,a,s){""===a?s(new Error("新邮箱不能为空")):a===t.email?s(new Error("新邮箱不能与旧邮箱相同")):function(t){return/^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/.test(t)}(a)?s():s(new Error("邮箱格式错误"))},trigger:"blur"}],code:[{required:!0,message:"验证码不能为空",trigger:"blur"}]}}},methods:{cancel:function(){this.resetForm()},sendCode:function(){var t=this;if(this.form.email&&this.form.email!==this.email){this.codeLoading=!0,this.buttonName="验证码发送中",this.codeData.value=this.form.email;var e=this;Object(l.g)(this.form.email).then(function(a){t.$message({showClose:!0,message:"发送成功，验证码有效期5分钟",type:"success"}),t.codeLoading=!1,t.isDisabled=!0,t.buttonName=t.time--+"秒后重新发送",t.timer=window.setInterval(function(){e.buttonName=e.time+"秒后重新发送",--e.time,e.time<0&&(e.buttonName="重新发送",e.time=60,e.isDisabled=!1,window.clearInterval(e.timer))},1e3)}).catch(function(e){t.codeLoading=!1,console.log(e.msg)})}},doSubmit:function(){var t=this;this.$refs.form.validate(function(e){if(!e)return!1;t.loading=!0,t.user={email:t.form.email,password:t.form.pass},Object(l.h)(t.form.code,t.user,t.form.pass).then(function(e){t.loading=!1,t.resetForm(),t.$notify({title:"邮箱修改成功",type:"success",duration:1500}),n.a.dispatch("GetInfo").then(function(){})}).catch(function(e){t.loading=!1,console.log(e.response.data.message)})})},resetForm:function(){this.dialog=!1,this.$refs.form.resetFields(),window.clearInterval(this.timer),this.time=60,this.buttonName="获取验证码",this.isDisabled=!1,this.form={pass:"",email:"",code:""}}}},p=(a("RpK0"),Object(d.a)(f,function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticStyle:{display:"inline-block"}},[a("el-button",{staticClass:"button",attrs:{size:"mini",type:"info"},on:{click:function(e){t.dialog=!0}}},[t._v("修改")]),t._v(" "),a("el-dialog",{attrs:{visible:t.dialog,title:t.title,width:"475px"},on:{"update:visible":function(e){t.dialog=e},close:t.cancel}},[a("el-form",{ref:"form",attrs:{model:t.form,rules:t.rules,size:"small","label-width":"88px"}},[a("el-form-item",{attrs:{label:"新邮箱",prop:"email"}},[a("el-input",{staticStyle:{width:"200px"},attrs:{"auto-complete":"on"},model:{value:t.form.email,callback:function(e){t.$set(t.form,"email",e)},expression:"form.email"}}),t._v(" "),a("el-button",{attrs:{loading:t.codeLoading,disabled:t.isDisabled,size:"small"},on:{click:t.sendCode}},[t._v(t._s(t.buttonName))])],1),t._v(" "),a("el-form-item",{attrs:{label:"验证码",prop:"code"}},[a("el-input",{staticStyle:{width:"320px"},model:{value:t.form.code,callback:function(e){t.$set(t.form,"code",e)},expression:"form.code"}})],1),t._v(" "),a("el-form-item",{attrs:{label:"当前密码",prop:"pass"}},[a("el-input",{staticStyle:{width:"320px"},attrs:{type:"password"},model:{value:t.form.pass,callback:function(e){t.$set(t.form,"pass",e)},expression:"form.pass"}})],1)],1),t._v(" "),a("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{attrs:{type:"text"},on:{click:t.cancel}},[t._v("取消")]),t._v(" "),a("el-button",{attrs:{loading:t.loading,type:"primary"},on:{click:t.doSubmit}},[t._v("确认")])],1)],1)],1)},[],!1,null,"609ab6fa",null));p.options.__file="updateEmail.vue";var v=p.exports,h=a("8SHQ"),b={name:"Center",components:{updatePass:m,updateEmail:v},data:function(){return{updateAvatarApi:h.a.UploadImageURL,headers:{Authorization:n.a.getters.token}}},computed:i()({},Object(o.b)(["avatar","name","email","createTime"])),methods:{formatEmail:function(t){return Object(r.c)(t)},handleSuccess:function(t,e,a){this.$notify({title:"头像修改成功",type:"success",duration:2500}),n.a.dispatch("GetInfo").then(function(){})},handleError:function(t){var e=JSON.parse(t.msg);this.$notify({title:e.message,type:"error",duration:2500})}}},g=(a("fnfW"),Object(d.a)(b,function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"app-container"},[a("div",{staticStyle:{width:"600px"},attrs:{id:"content-main"}},[a("el-card",{staticClass:"box-card user-info",attrs:{shadow:"never"}},[a("div",{staticClass:"avatar-content"},[a("el-upload",{staticClass:"avatar-uploader",attrs:{"show-file-list":!1,"on-success":t.handleSuccess,"on-error":t.handleError,headers:t.headers,action:t.updateAvatarApi}},[t.avatar?a("img",{staticClass:"avatar",attrs:{src:t.avatar,title:"点击上传头像"}}):a("i",{staticClass:"el-icon-plus avatar-uploader-icon"})])],1),t._v(" "),a("div",{staticClass:"user-info-content"},[a("div",[t._v("登录账号："+t._s(t.name))]),t._v(" "),a("div",[t._v("注册时间："+t._s(t.createTime))]),t._v(" "),a("div",[t._v("账号状态："),a("span",{staticStyle:{color:"#909399"}},[t._v("正常")])])])]),t._v(" "),a("el-card",{staticClass:"box-card reset-pass",attrs:{shadow:"never"}},[a("h4",{staticClass:"account-label-icon"},[t._v("登录密码：")]),t._v(" "),a("updatePass"),t._v(" "),a("p",[t._v("安全性高的密码可使账号更安全，建议设置同时包含字母，数字，符号的密码。")])],1),t._v(" "),a("el-card",{staticClass:"box-card reset-email",attrs:{shadow:"never"}},[a("h4",{staticClass:"account-label-icon"},[t._v("邮箱验证：")]),t._v(" "),a("updateEmail",{attrs:{email:t.email}}),t._v(" "),a("p",[t._v("你的邮箱："+t._s(t.formatEmail(t.email))+" ")]),t._v(" "),a("h4",[t._v("绑定邮箱可用于")]),t._v(" "),a("ul",[a("li",[t._v("安全管理，密码重置与修改")]),t._v(" "),a("li",[t._v("账号使用，使用邮箱登录系统")])])],1)],1)])},[],!1,null,null,null));g.options.__file="center.vue";e.default=g.exports},HZIk:function(t,e,a){},NDxK:function(t,e,a){},RpK0:function(t,e,a){"use strict";var s=a("oq/e");a.n(s).a},c2zk:function(t,e,a){"use strict";var s=a("NDxK");a.n(s).a},fnfW:function(t,e,a){"use strict";var s=a("HZIk");a.n(s).a},"oq/e":function(t,e,a){}}]);