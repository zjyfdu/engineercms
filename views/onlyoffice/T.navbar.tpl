{{define "navbar"}}
<!-- navbar-inverse一个带有黑色背景白色文本的导航栏 
固定在页面的顶部，向 .navbar class 添加 class .navbar-fixed-top
为了防止导航栏与页面主体中的其他内容
的顶部相交错，需要向 <body> 标签添加内边距，内边距的值至少是导航栏的高度。
-->
<style type="text/css">
    a.navbar-brand{display: none;}  
    @media (max-width: 960px) { 
     a.navbar-brand{ display: inline-block;} 
    }
  </style>
<nav class="navbar navbar-default navbar-static-top" role = "navigation">
  <div class="collapse navbar-collapse" id = "target-menu"> 
    <ul class="nav navbar-nav">
      <li {{if .IsOnlyOffice}}class="active"{{end}}>
        <a href="/onlyoffice">OnlyOffice</a>
      </li>
    </ul>

    <!-- <div class="pull-right"> -->
    <ul class="nav navbar-nav navbar-right">
        {{if .IsLogin}}
          {{if .IsAdmin}}
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
              <ul class="dropdown-menu">
                <li><a href="/admin" title="管理">进入后台</a></li>
                <li><a href="javascript:void(0)" id="login">重新登录</a></li>
                <li><a href="javascript:void(0)" onclick="logout()">退出</a></li>
              </ul>
            </li>
          {{else}}
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
              <ul class="dropdown-menu">
                <li><a href="/user" title="用户资料">用户资料</a></li>
                <li><a href="javascript:void(0)" id="login">重新登录</a></li>
                <li><a href="javascript:void(0)" onclick="logout()">退出</a></li>
              </ul>
            </li>
          {{end}}
        {{else}}
          {{if .IsAdmin}}
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
              <ul class="dropdown-menu">
                <li><a href="/admin" title="管理">进入后台</a></li>
                <li><a href="javascript:void(0)" id="login">重新登录</a></li>
                <li><a href="javascript:void(0)" onclick="logout()">退出</a></li>
              </ul>
            </li>
          {{else}}
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown">{{.Username}} <b class="caret"></b></a>
              <ul class="dropdown-menu">
                <li><a href="javascript:void(0)" id="login">登陆</a></li>
              </ul>
            </li>
          {{end}}
        {{end}}
    </ul>
    <!-- </div> -->
  </div>  
</nav>

  <!-- 登录模态框 -->
  <div class="form-horizontal">
    <div class="modal fade" id="modalNav">
      <div class="modal-dialog" id="modalDialog">
        <div class="modal-content">
          <div class="modal-header" style="background-color: #8bc34a">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">登录</h3>
            <label id="status"></label>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group" style="width: 100%;">
                <label class="col-sm-3 control-label">用户名 或 邮箱</label>
                <div class="col-sm-7">
                  <input id="uname" name="uname" type="text" value="" class="form-control" placeholder="Enter account" list="cars" onkeypress="getKey()"></div>
              </div>
              <div class="form-group" style="width: 100%;">
                <label class="col-sm-3 control-label">密码</label>
                <div class="col-sm-7">
                  <input id="pwd" name="pwd" type="password" value="" class="form-control" placeholder="Password" onkeypress="getKey()"></div>
              </div>
              <div class="form-group" style="width: 100%;">
                <label class="col-sm-3 control-label"><input type="checkbox">自动登陆</label>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            <button type="button" class="btn btn-primary" onclick="login()">登录</button>
          </div>
        </div>
      </div>
    </div>
  </div>

<script type="text/javascript">
    // 弹出登录模态框
  $("#login").click(function() {
    $('#modalNav').modal({
    show:true,
    backdrop:'static'
    });
  })
    //登陆功能
    function login(){
        var uname=document.getElementById("uname");
        if (uname.value.length==0){
          alert("请输入账号");
          return
        }
        var pwd=document.getElementById("pwd");
        if (pwd.value.length==0){
          alert("请输入密码");
          return
        }

        $.ajax({
          type:'post',
          url:'/loginpost',
          data:{
            "uname":$("#uname").val(),
            "pwd":$("#pwd").val()
          },
          success:function(result){
            if(result.islogin==0){
              $("#status").html("登陆成功");
              $('#modalNav').modal('hide');
              window.location.reload();
            }else  if(result.islogin==1){
              $("#status").html("用户名或密码错误！") 
            } else if(result.islogin==2){
              $("#status").html("密码错误") 
            }
          }
        })
    }
    //登出功能
     function logout(){
        $.ajax({
            type:'get',
            url:'/logout',
            data:{},
            success:function(result){
              if(result.islogin){
                // $("#status").html("登出成功");
                alert("登出成功");
                window.location.reload();
              }else {
               // $("#status").html("登出失败");
               alert("登出失败")
             }
           }
        })
    }

  //监听输入框中回车键
  function getKey(){  
    if(event.keyCode==13){  
      login()
    }     
  }
</script>

{{end}}