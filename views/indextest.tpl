<style type="text/css">
	.snap-content {
  	position: absolute;
  	top: 0;
  	right: 0;
  	bottom: 0;
  	left: 0;
  	width: auto;
  	height: auto;
  	z-index: 2;
  	overflow: auto;
  	-webkit-overflow-scrolling: touch;
  	-webkit-transform: translate3d(0, 0, 0);
  	   -moz-transform: translate3d(0, 0, 0);
  	    -ms-transform: translate3d(0, 0, 0);
  	     -o-transform: translate3d(0, 0, 0);
  	        transform: translate3d(0, 0, 0);
	}

	/* Header*/
	.header{
		background-image:url(../ky_img/menu-bg.png);
	  background-size:100px 100px;
	  height:61px;
	}
	.header-logo{
	  margin-top:18px;
	  margin-left:25px;
	  background-image:url(../ky_img/logo.png);
	  width:80px;
	  height:24px;
	  background-size:80px 24px;
	  float:right;
	  margin-right:25px;
	}
	.header-clear{
	    height:30px;
	}
	.open-menu{
    color:#FF5722;/*bcbcbc*/
    position:absolute;
    width:55px;
    height:60px;
    left:0px;
    top:0px;
    transition:all 250ms ease;
	}
	.open-menu:hover{
	    color:#000;/*FFFFFF*/
	    background-color:rgba(255,255,255,0.02);
	    transition:all 250ms ease;
	}
	.open-menu i{
	    width:55px;
	    height:60px;
	    line-height:60px;
	    text-align:center;
	    font-size:14px;
	}
	
	/*Content Heading*/
	/*///////////////*/
	.content-heading{
		margin-bottom:30px;	
	}
	.content-heading h4{
		color:#FFFFFF;
		position:absolute;	
		z-index:9999;
		text-transform:uppercase;
		margin-top:27px;
		padding-left:30px;
	    pointer-events:none;
	    font-weight:800;
	}
	.content-heading{
		max-height:100px;	
	}
	.content-heading p{
		color:#FFFFFF;
		position:absolute;	
		z-index:999;
		margin-top:48px;
		padding-left:30px;
		opacity:0.5;
	    pointer-events:none;
	}
	.content-heading .overlay{
		z-index:99;
		background-color:rgba(233,233,233,0.4);/*rgba(0,0,0,0.8)*/
	}
	.content-heading i{
		font-size:32px;
		position:absolute;
		color:#FFFFFF;
		right:30px;	
		z-index:999;
		margin-top:36px;
	    pointer-events:none;
	}
	.content-heading img{
		width:100%;
		display:block;
		position:relative;
		z-index:2;
	    transition:all 300ms ease;
	}
	.content-heading img:hover{
	    filter: blur(3px);  
	    -webkit-filter:blur(3px);
	    transition:all 300ms ease;
	}
	@media (min-width:768px){
		.content-heading{
			max-height:140px;	
		}
		
		.content-heading h4{
			font-size:20px;	
			margin-top:45px;
			padding-left:50px;
		}
		
		.content-heading p{
			font-size:13px;	
			margin-top:75px;
			padding-left:50px;
		}
		
		.content-heading i{
			font-size:40px;
			margin-top:53px;
			right:50px;	
		}
	}
	
	.overlay{
		pointer-events:none;
		position:absolute;
		width:100%;
		height:100%;
		background-color:rgba(0,0,0,0.5);
		z-index:9999;
	}
	
	.content{
		clear:both;
		margin-left:30px;
		margin-right:30px;
	}
	@media (min-width:768px){
		.content{
			margin-left:70px;
			margin-right:70px;	
		}
	}
	
	.decoration{ 
		height:1px;
		background-color:rgba(0,0,0,0.1);
		margin-bottom:30px;
		display:block;
		clear:both;
	}
	
	.container{
		margin-bottom:30px;
	}
	.no-bottom{
		margin-bottom:0px;
		padding-bottom:0px;
	}
	
	@media (min-width:760px) { 
		.one-third-responsive{
			width:30%;
			float:left;
			margin-right:5%;	
		}
		
		.one-half-responsive{
			width:46%;
			float:left;
			margin-right:8%;
		}
		
		.sidebar-left-big{
			width:70%;
			float:left;
			margin-right:5%	
		}
		
		.sidebar-right-small{
			width:25%;
			float:right;	
		}
			
		.sidebar-right-big{
			width:70%;
			float:right;	
		}
		
		.sidebar-left-small{
			width:25%;
			float:left;
			margin-right:5%;	
		}
		
		.hide-if-responsive{
			display:none;
		}
	}
	
	/*Thumbnails Columns*/
	.thumb-clear{
		height:40px;
		display:block;
		width:100%;
	}
	.thumb-left{
		line-height:24px;
		display:block;
		padding-bottom:10px;
	}
	.thumb-left a{
		display:block;
		text-align:right;
	}
	.thumb-left img{
		width:100px;
		height:100px;
		border-radius:100px;
		float:left;
		margin-right:20px;
	    transition:all 500ms ease;
	}
	.thumb-left img:hover{
	    transform:scale(0.9, 0.9);
	    transition:all 500ms ease;
	}
	.thumb-left strong{
		color:#1a1a1a;
		display:inline-block;
		padding-bottom:5px;
		font-size:13px;
	}
	.thumb-left em{
		font-style:normal;
	}
	.thumb-right{
		line-height:24px;
		display:block;
		padding-bottom:10px;
	}
	.thumb-right img{
		width:100px;
		height:100px;
		border-radius:100px;
		float:right;
		margin-left:20px;
	    transition:all 500ms ease;
	}
	.thumb-right img:hover{
	    transform:scale(0.9, 0.9);
	    transition:all 500ms ease;
	}
	.thumb-right strong{
		color:#1a1a1a;
		display:inline-block;
		padding-top:5px;
		padding-bottom:5px;
		font-size:13px;
	}
	.thumb-right em{
		font-style:normal;
	}
	@media (min-width:600px){
		.thumb-left img{
			width:140px;
			height:140px;
			border-radius:150px;	
		}
		
		.thumb-left em{
			line-height:28px;	
		}
		
		.thumb-left strong{
			padding-top:10px;	
		}
		
		.thumb-right img{
			width:140px;
			height:140px;
			border-radius:150px;	
		}
		
		.thumb-right em{
			line-height:28px;	
		}
		
		.thumb-right strong{
			padding-top:10px;	
		}	
	}
	
	.last-column{
		margin-right:0%!important;
	}
</style>
<!-- Content Heding -->
<div class="content-heading">
  <h4>最新信息</h4>
  <p>从设计角度理解项目</p>
  <i class="fa fa-cog"></i>
  <div class="overlay"></div>
  <img src="images/5w.jpg" alt="img">
</div>
<!-- Page Content-->
<div class="content">
    <div class="decoration"></div>
      <div id="news" class="container no-bottom"> 
      <!-- <a href="#" class="next-staff"></a> -->
      <!-- <a href="#" class="prev-staff"></a> -->
      <!-- <div class="staff-slider" data-snap-ignore="true"> -->
      <!-- Page Columns-->
      <div>
        <div class="one-half-responsive">
          <p class="thumb-left no-bottom">
            <img id="img1" src="" alt="img">
            <strong id="title1"></strong>
            <em id="content1"><br></em> 
          </p>
          <div class="thumb-clear"></div>
        </div>
        <div class="decoration hide-if-responsive"></div>
        <div class="one-half-responsive last-column">
            <p class="thumb-right no-bottom">
                <img src="" alt="img">
                <strong></strong>
                <em><br></em> 
            </p>
            <div class="thumb-clear"></div>
        </div> 
      </div>
        <div>
          <div class="decoration hide-if-responsive"></div>
          <div class="one-half-responsive">
              <p class="thumb-left no-bottom">
                  <img src="images/3s.jpg" alt="img">
                  <strong></strong>
                  <em><br></em>
              </p>
              <div class="thumb-clear"></div>
          </div>
          <div class="decoration hide-if-responsive"></div>
          <div class="one-half-responsive last-column">
            <p class="thumb-right no-bottom">
                  <img src="images/4s.jpg" alt="img">
                  <strong></strong>
                  <em><br></em>
            </p>
              <div class="thumb-clear"></div>
          </div>
          </div>
          <div>
            <div class="one-half-responsive">
                <p class="thumb-left no-bottom">
                    <img src="images/1s.jpg" alt="img">
                    <strong></strong>
                    <em><br></em> 
                </p>
                <div class="thumb-clear"></div>
            </div>
            <div class="decoration hide-if-responsive"></div>
            <div class="one-half-responsive last-column">
                <p class="thumb-right no-bottom">
                    <img src="images/2s.jpg" alt="img">
                    <strong></strong>
                    <em><br></em> 
                </p>
                <div class="thumb-clear"></div>
            </div> 
          </div>
            <!-- <div> -->
                <!-- <div class="decoration hide-if-responsive"></div>
                <div class="one-half-responsive">
                    <p class="thumb-left no-bottom">
                        <img src="images/3s.jpg" alt="img">
                        <strong></strong>
                        <em><br></em>
                    </p>
                    <div class="thumb-clear"></div>
                </div>
                <div class="decoration hide-if-responsive"></div>
                <div class="one-half-responsive last-column">
                    <p class="thumb-right no-bottom">
                        <img src="images/4s.jpg" alt="img">
                        <strong></strong>
                        <em><br></em>
                    </p>
                    <div class="thumb-clear"></div>
                </div>
            </div> -->

            </div>
            <div class="decoration"></div>
        </div>

<!-- Content Heading -->
        <div class="content-heading">
            <h4>参与人员</h4>
            <p>建设者的睿智</p>
            <i class="fa fa-user"></i>
            <div class="overlay"></div>
            <img src="images/6w.jpg" alt="img">
        </div>
        <!-- Page Content-->
        <div class="content">
            <div class="decoration"></div>
            <!-- Staff Slider-->
            <div class="container">
                <a href="#" class="next-staff"></a>
                <a href="#" class="prev-staff"></a>
                <div class="staff-slider" data-snap-ignore="true">
                <div>
                  <div class="staff-item">
                    <img src="images/1s.jpg" alt="img">
                    <h4>John Doe</h4>
                    <em>总设计师</em>
                    <p style="text-align:center;">&nbsp;&nbsp;&nbsp;&nbsp;对于设计师来说，能够有机会参与这么有挑战的工作，<br/> 
                            是非常幸运和自豪的!</p>
                  </div>
                </div>
                    <div>
                        <div class="staff-item">
                            <img src="images/2s.jpg" alt="img">
                            <h4>Jane Hidden</h4>
                            <em>项目负责人</em>
                            <p style="text-align:center;">将设计任务完美地完成，业主满意，社会满意!</p> 
                            <!-- <a href="#" class="button button-green center-button">Text</a> -->
                        </div>
                    </div>
                    <div>
                        <div class="staff-item">
                            <img src="images/3s.jpg" alt="img">
                            <h4>Johanna Pear</h4>
                            <em>施工组织</em>
                            <p style="text-align:center;">
                                精确的施工和安装，做好每一步，安全摆在首位!
                            </p>
                            <!-- <a href="#" class="button button-blue center-button">Mail</a> -->
                        </div>
                    </div>
                    <div>
                        <div class="staff-item">
                            <img src="images/4s.jpg" alt="img">
                            <h4>Mike Grape</h4>
                            <em>总监</em>
                            <p style="text-align:center;">严格把关，为业主服务，发挥管理水平!
                            </p>
                            <!-- <a href="#" class="button button-dark center-button">Read More</a> -->
                        </div>
                    </div>
                    <div>
                        <div class="staff-item">
                            <img src="images/5s.jpg" alt="img">
                            <h4>Victor Leaf</h4>
                            <em>咨询工程师</em>
                            <!-- <strong> -->
                            <p style="text-align:center;">
                                互相合作，做最好的参谋!
                            </p>
                            <!-- </strong> -->
                            <!-- <a href="#" class="button button-orange center-button">Facebook</a> -->
                        </div>
                    </div>
                    <div>
                        <div class="staff-item">
                            <img src="images/6s.jpg" alt="img">
                            <h4>Snow White</h4>
                            <em>造价工程师</em>
                            <p style="text-align:center;">
                                控制好投资，节约成本!
                            </p>
                            <!-- <a href="#" class="button button-yellow center-button">Twitter</a> -->
                        </div>
                    </div>
              </div>
            </div>
            <div class="decoration"></div>
        </div>