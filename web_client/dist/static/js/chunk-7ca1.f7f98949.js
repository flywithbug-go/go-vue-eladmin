(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-7ca1","chunk-7455","chunk-5f06"],{"5ZET":function(e,t,i){"use strict";i.r(t);var o=i("41Be"),n={components:{eForm:i("gAlZ").default},props:{query:{type:Object,required:!0},menus:{type:Array,required:!0},roles:{type:Array,required:!0}},data:function(){return{downloadLoading:!1}},methods:{checkPermission:o.a,toQuery:function(){console.log(this.query),this.$parent.page=0,this.$parent.init()}}},r=i("KHd+"),s=Object(r.a)(n,function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",{staticClass:"head-container"},[i("el-input",{staticClass:"filter-item",staticStyle:{width:"200px"},attrs:{clearable:"",placeholder:"输入名称搜索"},nativeOn:{keyup:function(t){return"button"in t||!e._k(t.keyCode,"enter",13,t.key,"Enter")?e.toQuery(t):null}},model:{value:e.query.value,callback:function(t){e.$set(e.query,"value",t)},expression:"query.value"}}),e._v(" "),i("el-button",{staticClass:"filter-item",attrs:{size:"mini",type:"primary",icon:"el-icon-search"},on:{click:e.toQuery}},[e._v("搜索")]),e._v(" "),i("div",{staticStyle:{display:"inline-block",margin:"0px 2px"}},[e.checkPermission(["ADMIN","MENU_ALL","MENU_CREATE"])?i("el-button",{staticClass:"filter-item",attrs:{size:"mini",type:"primary",icon:"el-icon-plus"},on:{click:function(t){e.$refs.form.dialog=!0}}},[e._v("新增")]):e._e(),e._v(" "),i("eForm",{ref:"form",attrs:{roles:e.roles,menus:e.menus,"is-add":!0}})],1)],1)},[],!1,null,null,null);s.options.__file="header.vue";t.default=s.exports},"8Uww":function(e,t,i){},"95TX":function(e,t,i){"use strict";i.r(t);var o=i("41Be"),n=i("zF5t"),r=i("itRl"),s=i("3ADX"),a=i("Hycs"),l=i("7Qib"),c=i("5ZET"),u=i("9Wvd"),d={components:{eHeader:c.default,edit:u.default,treeTable:r.a},mixins:[s.a],data:function(){return{columns:[{text:"名称",value:"name"}],delLoading:!1,sup_this:this,menus:[],roles:[]}},created:function(){var e=this;this.getRoles(),this.getMenus(),this.$nextTick(function(){e.init()})},methods:{parseTime:l.b,checkPermission:o.a,beforeInit:function(){this.url="/menu/list";var e=this.query.value;return this.params={page:this.page,size:this.size,sort:"id,desc"},e&&(this.params.name=e),!0},subDelete:function(e,t){var i=this;this.delLoading=!0,Object(a.c)(t.id).then(function(){i.delLoading=!1,t.delPopover=!1,i.init(),i.$notify({title:"删除成功",type:"success",duration:2500})}).catch(function(e){i.delLoading=!1,t.delPopover=!1,console.log(e.msg)})},getMenus:function(){var e=this;Object(a.e)().then(function(t){e.menus=[];var i={id:0,label:"顶级类目",children:[]};i.children=t.list,e.menus.push(i)})},getRoles:function(){var e=this;this.roles=[],Object(n.d)().then(function(t){e.roles=t.list})}}},f=(i("AFtx"),i("KHd+")),m=Object(f.a)(d,function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",{staticClass:"app-container"},[i("eHeader",{attrs:{roles:e.roles,menus:e.menus,query:e.query}}),e._v(" "),i("tree-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],attrs:{data:e.data,"expand-all":!0,columns:e.columns,border:"",size:"small"}},[i("el-table-column",{attrs:{prop:"icon",label:"图标",align:"center",width:"80px"},scopedSlots:e._u([{key:"default",fn:function(e){return[i("svg-icon",{attrs:{"icon-class":e.row.icon}})]}}])}),e._v(" "),i("el-table-column",{attrs:{prop:"sort",align:"center",width:"100px",label:"排序"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("el-tag",[e._v(e._s(t.row.sort))])]}}])}),e._v(" "),i("el-table-column",{attrs:{"show-overflow-tooltip":!0,prop:"path",label:"链接地址"}}),e._v(" "),i("el-table-column",{attrs:{"show-overflow-tooltip":!0,prop:"component",label:"组件路径"}}),e._v(" "),i("el-table-column",{attrs:{prop:"iframe",width:"100px",label:"内部菜单"},scopedSlots:e._u([{key:"default",fn:function(t){return[t.row.iframe?i("span",[e._v("否")]):i("span",[e._v("是")])]}}])}),e._v(" "),i("el-table-column",{attrs:{prop:"createTime",label:"创建日期"},scopedSlots:e._u([{key:"default",fn:function(t){return[i("span",[e._v(e._s(e.parseTime(t.row.createTime)))])]}}])}),e._v(" "),i("el-table-column",{attrs:{label:"操作",width:"150px",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e.checkPermission(["ADMIN","MENU_ALL","MENU_EDIT"])?i("edit",{attrs:{roles:e.roles,menus:e.menus,data:t.row,sup_this:e.sup_this}}):e._e(),e._v(" "),e.checkPermission(["ADMIN","MENU_ALL","MENU_DELETE"])?i("el-popover",{attrs:{placement:"top",width:"200"},model:{value:t.row.delPopover,callback:function(i){e.$set(t.row,"delPopover",i)},expression:"scope.row.delPopover"}},[i("p",[e._v("确定删除吗,如果存在下级节点则一并删除，此操作不能撤销！")]),e._v(" "),i("div",{staticStyle:{"text-align":"right",margin:"0"}},[i("el-button",{attrs:{size:"mini",type:"text"},on:{click:function(e){t.row.delPopover=!1}}},[e._v("取消")]),e._v(" "),i("el-button",{attrs:{loading:e.delLoading,type:"primary",size:"mini"},on:{click:function(i){e.subDelete(t.$index,t.row)}}},[e._v("确定")])],1),e._v(" "),i("el-button",{attrs:{slot:"reference",type:"danger",size:"mini"},on:{click:function(e){t.row.delPopover=!0}},slot:"reference"},[e._v("删除")])],1):e._e()]}}])})],1)],1)},[],!1,null,"4e8d3834",null);m.options.__file="index.vue";t.default=m.exports},"9Wvd":function(e,t,i){"use strict";i.r(t);var o={components:{eForm:i("gAlZ").default},props:{data:{type:Object,required:!0},sup_this:{type:Object,required:!0},menus:{type:Array,required:!0},roles:{type:Array,required:!0}},methods:{to:function(){var e=this.$refs.form;e.roleIds=[],e.form={id:this.data.id,component:this.data.component,name:this.data.name,sort:this.data.sort,pid:this.data.pid,path:this.data.path,iframe:this.data.iframe?"true":"false",roles:[],icon:this.data.icon},this.data.roles||(this.data.roles=[]),this.data.roles.forEach(function(t,i){e.roleIds.push(t.id)}),e.dialog=!0}}},n=(i("Du2c"),i("KHd+")),r=Object(n.a)(o,function(){var e=this.$createElement,t=this._self._c||e;return t("div",[t("el-button",{attrs:{size:"mini",type:"success"},on:{click:this.to}},[this._v("编辑")]),this._v(" "),t("eForm",{ref:"form",attrs:{roles:this.roles,menus:this.menus,sup_this:this.sup_this,"is-add":!1}})],1)},[],!1,null,"0bef9d48",null);r.options.__file="edit.vue";t.default=r.exports},AFtx:function(e,t,i){"use strict";var o=i("DitO");i.n(o).a},DitO:function(e,t,i){},Du2c:function(e,t,i){"use strict";var o=i("8Uww");i.n(o).a},HufR:function(e,t,i){"use strict";var o=i("oDta");i.n(o).a},gAlZ:function(e,t,i){"use strict";i.r(t);var o=i("Hycs"),n=i("cCY5"),r=i.n(n),s=/\.\/(.*)\.svg/,a=function(e){return e.keys()}(i("Uf/o")).map(function(e){return e.match(s)[1]}),l={name:"IconSelect",data:function(){return{name:"",iconList:a}},methods:{filterIcons:function(){var e=this;this.name?this.iconList=this.iconList.filter(function(t){return t.includes(e.name)}):this.iconList=a},selectedIcon:function(e){this.$emit("selected",e),document.body.click()},reset:function(){this.name="",this.iconList=a}}},c=(i("HufR"),i("KHd+")),u=Object(c.a)(l,function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",{staticClass:"icon-body"},[i("el-input",{staticStyle:{position:"relative"},attrs:{clearable:"",placeholder:"请输入图标名称"},on:{clear:e.filterIcons},nativeOn:{input:function(t){return e.filterIcons(t)}},model:{value:e.name,callback:function(t){e.name=t},expression:"name"}},[i("i",{staticClass:"el-icon-search el-input__icon",attrs:{slot:"suffix"},slot:"suffix"})]),e._v(" "),i("div",{staticClass:"icon-list"},e._l(e.iconList,function(t,o){return i("div",{key:o,on:{click:function(i){e.selectedIcon(t)}}},[i("svg-icon",{staticStyle:{height:"30px",width:"16px"},attrs:{"icon-class":t}}),e._v(" "),i("span",[e._v(e._s(t))])],1)}))],1)},[],!1,null,"1628183e",null);u.options.__file="index.vue";var d=u.exports,f=(i("VCwm"),{components:{Treeselect:r.a,IconSelect:d},props:{menus:{type:Array,required:!0},roles:{type:Array,required:!0},isAdd:{type:Boolean,required:!0},sup_this:{type:Object,default:null}},data:function(){return{loading:!1,dialog:!1,form:{name:"",sort:999,path:"",component:"",iframe:"false",roles:[],pid:0,icon:""},roleIds:[],rules:{name:[{required:!0,message:"请输入名称",trigger:"blur"}],sort:[{required:!0,message:"请输入序号",trigger:"blur",type:"number"}],iframe:[{required:!0,message:"请选择菜单类型",trigger:"blur"}]}}},methods:{cancel:function(){this.resetForm()},doSubmit:function(){var e=this;this.$refs.form.validate(function(t){if(!t)return!1;e.loading=!0,e.form.roles=[];var i=e;e.roleIds.forEach(function(e,t){var o={id:e};i.form.roles.push(o)}),e.isAdd?e.doAdd():e.doEdit()})},doAdd:function(){var e=this;Object(o.a)(this.form).then(function(t){e.resetForm(),e.$notify({title:"添加成功",type:"success",duration:2500}),e.loading=!1,setTimeout(function(){e.$parent.$parent.init(),e.$parent.$parent.getMenus()},200)}).catch(function(t){e.loading=!1,console.log(t.response.data.message)})},doEdit:function(){var e=this;Object(o.d)(this.form).then(function(t){e.resetForm(),e.$notify({title:"修改成功",type:"success",duration:2500}),e.loading=!1,setTimeout(function(){e.sup_this.init(),e.sup_this.getMenus()},200)}).catch(function(t){e.loading=!1,console.log(t.msg)})},resetForm:function(){this.dialog=!1,this.$refs.form.resetFields(),this.form={name:"",sort:999,path:"",component:"",iframe:"false",roles:[],pid:0,icon:""},this.roleIds=[]},selected:function(e){this.form.icon=e}}}),m=(i("j/Nk"),Object(c.a)(f,function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("el-dialog",{attrs:{"append-to-body":!0,visible:e.dialog,title:e.isAdd?"新增菜单":"编辑菜单",width:"600px"},on:{"update:visible":function(t){e.dialog=t}}},[i("el-form",{ref:"form",attrs:{model:e.form,rules:e.rules,size:"small","label-width":"80px"}},[i("el-form-item",{attrs:{label:"菜单图标"}},[i("el-popover",{attrs:{placement:"bottom-start",width:"460",trigger:"click"},on:{show:function(t){e.$refs.iconSelect.reset()}}},[i("IconSelect",{ref:"iconSelect",on:{selected:e.selected}}),e._v(" "),i("el-input",{staticStyle:{width:"460px"},attrs:{slot:"reference",placeholder:"点击选择图标",readonly:""},slot:"reference",model:{value:e.form.icon,callback:function(t){e.$set(e.form,"icon",t)},expression:"form.icon"}},[e.form.icon?i("svg-icon",{staticClass:"el-input__icon",staticStyle:{height:"32px",width:"16px"},attrs:{slot:"prefix","icon-class":e.form.icon},slot:"prefix"}):i("i",{staticClass:"el-icon-search el-input__icon",attrs:{slot:"prefix"},slot:"prefix"})],1)],1)],1),e._v(" "),i("el-form-item",{attrs:{label:"菜单名称",prop:"name"}},[i("el-input",{staticStyle:{width:"460px"},attrs:{placeholder:"名称"},model:{value:e.form.name,callback:function(t){e.$set(e.form,"name",t)},expression:"form.name"}})],1),e._v(" "),i("el-form-item",{attrs:{label:"菜单排序",prop:"sort"}},[i("el-input",{staticStyle:{width:"460px"},attrs:{placeholder:"序号越小越靠前"},model:{value:e.form.sort,callback:function(t){e.$set(e.form,"sort",e._n(t))},expression:"form.sort"}})],1),e._v(" "),i("el-form-item",{attrs:{label:"内部菜单",prop:"iframe"}},[i("el-radio",{attrs:{label:"false"},model:{value:e.form.iframe,callback:function(t){e.$set(e.form,"iframe",t)},expression:"form.iframe"}},[e._v("是")]),e._v(" "),i("el-radio",{attrs:{label:"true"},model:{value:e.form.iframe,callback:function(t){e.$set(e.form,"iframe",t)},expression:"form.iframe"}},[e._v("否")])],1),e._v(" "),i("el-form-item",{attrs:{label:"链接地址"}},[i("el-input",{staticStyle:{width:"460px"},attrs:{placeholder:"菜单路径"},model:{value:e.form.path,callback:function(t){e.$set(e.form,"path",t)},expression:"form.path"}})],1),e._v(" "),"false"===e.form.iframe?i("el-form-item",{attrs:{label:"组件路径"}},[i("el-input",{staticStyle:{width:"460px"},attrs:{placeholder:"菜单路径"},model:{value:e.form.component,callback:function(t){e.$set(e.form,"component",t)},expression:"form.component"}})],1):e._e(),e._v(" "),i("el-form-item",{attrs:{label:"上级类目"}},[i("treeselect",{staticStyle:{width:"460px"},attrs:{options:e.menus,placeholder:"选择上级类目"},model:{value:e.form.pid,callback:function(t){e.$set(e.form,"pid",t)},expression:"form.pid"}})],1),e._v(" "),i("el-form-item",{staticStyle:{"margin-top":"-10px","margin-bottom":"0px"},attrs:{label:"选择角色"}},[i("treeselect",{staticStyle:{width:"460px"},attrs:{multiple:!0,options:e.roles,placeholder:"请选择角色"},model:{value:e.roleIds,callback:function(t){e.roleIds=t},expression:"roleIds"}})],1)],1),e._v(" "),i("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[i("el-button",{attrs:{type:"text"},on:{click:e.cancel}},[e._v("取消")]),e._v(" "),i("el-button",{attrs:{loading:e.loading,type:"primary"},on:{click:e.doSubmit}},[e._v("确认")])],1)],1)},[],!1,null,"ed6b3d2c",null));m.options.__file="form.vue";t.default=m.exports},giYQ:function(e,t,i){},"j/Nk":function(e,t,i){"use strict";var o=i("giYQ");i.n(o).a},oDta:function(e,t,i){},zF5t:function(e,t,i){"use strict";i.d(t,"d",function(){return r}),i.d(t,"a",function(){return s}),i.d(t,"b",function(){return a}),i.d(t,"c",function(){return l});var o=i("bNJ/"),n=i("8SHQ");function r(){return Object(o.a)({url:n.a.PathRoleTree,method:"get"})}function s(e){return Object(o.a)({url:"/role",method:"post",data:e})}function a(e){var t={id:e};return Object(o.a)({url:"/role",method:"delete",data:t})}function l(e){return Object(o.a)({url:"/role",method:"put",data:e})}}}]);