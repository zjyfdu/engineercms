<!-- 这个是显示左侧栏，右边index_user显示用户的cms情况 -->
<!DOCTYPE html>
    <head>
        <title>珠三角水资源配置工程设代</title>

          <script src="/static/vue.js/vue.js"></script>
  <!-- 引入样式 -->
  <link rel="stylesheet" href="https://unpkg.com/element-ui@2.4.11/lib/theme-chalk/index.css">
  <!-- <link rel="stylesheet" href="/static/vue.js/index.css"> -->
  <!-- <script src="https://unpkg.com/element-ui@2.4.11/lib/index.js"></script> -->
  <script src="/static/vue.js/index.js"></script>
  <!-- 引入组件库 -->
  <!-- <script src="https://unpkg.com/element-ui/lib/index.js"></script> -->
  <!-- <script src="https://unpkg.com/axios/dist/axios.min.js"></script> -->
  <script src="/static/vue.js/axios.js"></script>

<style >
    .login-container {
      -webkit-border-radius: 5px;
      border-radius: 5px;
      -moz-border-radius: 5px;
      background-clip: padding-box;
      margin: 180px auto;
      width: 350px;
      padding: 35px 35px 15px 35px;
      background: #fff;
      border: 1px solid #eaeaea;
      box-shadow: 0 0 25px #cac6c6;
      .title {
        margin: 0px auto 40px auto;
        text-align: center;
        color: #505458;
      }
      .remember {
        margin: 0px 0px 35px 0px;
      }
    }
</style>
    </head>
<body>

<div id="app">
<!-- Table -->
<el-button type="text" @click="dialogTableVisible = true">打开嵌套表格的 Dialog</el-button>

<el-dialog title="收货地址" :visible.sync="dialogTableVisible">
  <el-table :data="gridData">
    <el-table-column property="date" label="日期" width="150"></el-table-column>
    <el-table-column property="name" label="姓名" width="200"></el-table-column>
    <el-table-column property="address" label="地址"></el-table-column>
  </el-table>
</el-dialog>

<!-- Form -->
<el-button type="text" @click="dialogFormVisible = true">打开嵌套表单的 Dialog</el-button>

<!-- <el-dialog title="收货地址" :visible.sync="dialogFormVisible" class="login-container">
  <el-form :model="form">
    <el-form-item label="活动名称" :label-width="formLabelWidth">
      <el-input v-model="form.name" autocomplete="off"></el-input>
    </el-form-item>
    <el-form-item label="活动区域" :label-width="formLabelWidth">
      <el-select v-model="form.region" placeholder="请选择活动区域">
        <el-option label="区域一" value="shanghai"></el-option>
        <el-option label="区域二" value="beijing"></el-option>
      </el-select>
    </el-form-item>
  </el-form>
  <div slot="footer" class="dialog-footer">
    <el-button @click="dialogFormVisible = false">取 消</el-button>
    <el-button type="primary" @click="dialogFormVisible = false">确 定</el-button>
  </div>
</el-dialog> -->

<template>
<!-- <el-dialog title="收货地址" :visible.sync="dialogFormVisible">   -->
  <el-form :model="ruleForm2" :rules="rules2" ref="ruleForm2" label-position="left" label-width="0px" class="demo-ruleForm login-container">
    <h3 class="title">系统登录</h3>
    <el-form-item prop="account">
      <el-input type="text" v-model="ruleForm2.account" auto-complete="off" placeholder="账号">
      </el-input>
    </el-form-item>
    <el-form-item prop="checkPass">
      <el-input type="password" v-model="ruleForm2.checkPass" auto-complete="off" placeholder="密码"></el-input>
    </el-form-item>
    <el-checkbox v-model="checked" checked class="remember">记住密码</el-checkbox>
    <el-form-item style="width:100%;">
      <el-button type="primary" style="width:100%;" @click.native.prevent="handleSubmit2" :loading="logining">登录</el-button>
    </el-form-item>
  </el-form>
<!-- </el-dialog> -->
</template>


<el-button round @click="dialogFormVisible = true">登录</el-button>
<el-dialog title="系统登录" :visible.sync="dialogFormVisible" center>
  <!-- 插入测试 -->
  <el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="100px" class="demo-ruleForm">
    <el-form-item label="账号" prop="num">
      <el-input v-model.number="ruleForm2.num" auto-complete="off" placeholder="账号"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="pass">
      <el-input type="password" v-model="ruleForm2.pass" auto-complete="off" placeholder="密码"></el-input>
    </el-form-item>
    <el-checkbox v-model="checked" checked class="remember">记住密码</el-checkbox>
    <el-form-item label="记住密码" prop="delivery">
      <el-switch v-model="ruleForm2.delivery"></el-switch>
    </el-form-item> 
    <span><a>忘记密码？</a></span>
  </el-form>
   <!-- 插入测试 -->
  <div slot="footer" class="dialog-footer">
    <el-button @click="dialogFormVisible = false; resetForm('ruleForm2')">取 消</el-button>
    <el-button type="primary" @click="submitForm('ruleForm2')">登 录</el-button>
  </div>
</el-dialog>

</div>

<script type="text/javascript">
  // export default {
  //   data() {
  //     return {
  //       logining: false,
  //       ruleForm2: {
  //         account: 'admin',
  //         checkPass: '123456'
  //       },
  //       rules2: {
  //         account: [
  //           { required: true, message: '请输入账号', trigger: 'blur' },
  //           //{ validator: validaePass }
  //         ],
  //         checkPass: [
  //           { required: true, message: '请输入密码', trigger: 'blur' },
  //           //{ validator: validaePass2 }
  //         ]
  //       },
  //       checked: true
  //     };
  //   },
  //   methods: {
  //     handleReset2() {
  //       this.$refs.ruleForm2.resetFields();
  //     },
  //     handleSubmit2(ev) {
  //       var _this = this;
  //       this.$refs.ruleForm2.validate((valid) => {

  //       });
  //     }
  //   }
  // }


var Main = {
    data() {
      var checkNum = (rule, value, callback) => {
        if (!value) {
          return callback(new Error('账号不能为空'));
        }
        setTimeout(() => {
          if (!Number.isInteger(value)) {
            callback(new Error('请输入数字值'));
          } else {
            var myreg=/^[1][3,4,5,7,8][0-9]{9}$/;
            if (!myreg.test(value) ) {
              callback(new Error('请输入正确的手机号码'));
            } else {
              callback();
            }

          }
        }, 1000);
      };
      var validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error('请输入密码'));
        } else {

          callback();
        }
      };
    
      // return {
      //   gridData: [{
      //     date: '2016-05-02',
      //     name: '王小虎',
      //     address: '上海市普陀区金沙江路 1518 弄'
      //   }, {
      //     date: '2016-05-04',
      //     name: '王小虎',
      //     address: '上海市普陀区金沙江路 1518 弄'
      //   }, {
      //     date: '2016-05-01',
      //     name: '王小虎',
      //     address: '上海市普陀区金沙江路 1518 弄'
      //   }, {
      //     date: '2016-05-03',
      //     name: '王小虎',
      //     address: '上海市普陀区金沙江路 1518 弄'
      //   }],
      //   dialogTableVisible: false,
      //   dialogFormVisible: false,
      //   form: {
      //     name: '',
      //     region: '',
      //     date1: '',
      //     date2: '',
      //     delivery: false,
      //     type: [],
      //     resource: '',
      //     desc: ''
      //   },
      //   formLabelWidth: '120px'
      // };
      return {
        
        loginPower:false,
        /*插入form方法*/
        /*设定规则指向*/
        ruleForm2: {
          pass: '',
          num: '',
           delivery: false,
        },
        rules2: {
          pass: [
            { validator: validatePass, trigger: 'blur' }
          ],

          num: [
            { validator: checkNum, trigger: 'blur' }
          ]
        },
        /*插入form方法*/
        dialogTableVisible: false,
        dialogFormVisible: false,
        form: {
          name: '',
          type: [],
          resource: '',
          desc: ''
        },
       formLabelWidth: '120px',


        gridData: [{
          date: '2016-05-02',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄'
        }, {
          date: '2016-05-04',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄'
        }, {
          date: '2016-05-01',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄'
        }, {
          date: '2016-05-03',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄'
        }],
        dialogTableVisible: false,
        dialogFormVisible: false,
        form: {
          name: '',
          region: '',
          date1: '',
          date2: '',
          delivery: false,
          type: [],
          resource: '',
          desc: ''
        },
        formLabelWidth: '120px',

        logining: false,
        // dialogFormVisible: false,
        ruleForm2: {
          account: 'admin',
          checkPass: '123456'
        },
        rules2: {
          account: [
            { required: true, message: '请输入账号', trigger: 'blur' },
            //{ validator: validaePass }
          ],
          checkPass: [
            { required: true, message: '请输入密码', trigger: 'blur' },
            //{ validator: validaePass2 }
          ]
        },
        checked: true,
      };
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
          //提交成功做的动作
          dialogFormVisible = false;
            /* alert('submit!') ; */
            this.$message({
              type: 'success',
              message: '提交成功' 
            });
          } else {
            console.log('error submit!!');
            return false;
          }
        }); 
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },

      handleReset2() {
        this.$refs.ruleForm2.resetFields();
      },
      handleSubmit2(ev) {
        var _this = this;
        this.$refs.ruleForm2.validate((valid) => {
        });
      }
    }
  };
var Ctor = Vue.extend(Main)
new Ctor().$mount('#app')
</script>
  </body>
