<!DOCTYPE html>
<!DOCTYPE html>
<html>
<head>
  <title>基于Golang的广告投放系统的设计与实现</title>
  <meta name="viewport" content="width=1310">
  <meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
  <meta http-equiv="description" content="this is my page">
  <meta http-equiv="content-type" content="text/html; charset=UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
  <script type="text/javascript"  src="http://cdn.wanfangdata.com.cn/js/min/jquery-1.9.1.min.js"></script>
  <link rel="shortcut icon" href="http://cdn.wanfangdata.com.cn/page/images/favicon.ico" >
  <link href="http://cdn.wanfangdata.com.cn/page/css/min/public.css?version=1.2020.236-20200506094020" rel="stylesheet" type="text/css">
  <link href="http://cdn.wanfangdata.com.cn/page/css/min/headerCom.css?version=1.2020.236-20200506094020" rel="stylesheet" type="text/css">

<body>
<!--公共头部-->
<div class="top me-top">
  <div class="container me-container">
    <!--
   logo
-->
<style>
   .anxs-top-88qwe-logo_sns{padding: 9px 0;float: left;height:84px;box-sizing: border-box;}.anxs-top-88qwe-logo_sns a{color: #333;cursor: pointer;text-decoration: none}.anxs-top-88qwe-logo_sns .anxs-top-88qwe-imgwrapper{margin-right:25px;display: block;float: left;}.anxs-top-index_sns{float: left;line-height: 60px;}.anxs-top-88qwe-logo_sns .nav_color_index{font-size: 16px;vertical-align: middle; margin: 0 10px;}.anxs-top-88qwe-logo_sns .nav_color_index:hover{color: #427dc9;}
</style>
<div class="anxs-top-88qwe-logo_sns">
    <a href="http://www.wanfangdata.com.cn/" class="anxs-top-88qwe-imgwrapper">
        <img src="http://cdn.login.wanfangdata.com.cn/Content/src/img/anxs-logo_sns.png">
    </a>
    <div class="anxs-top-index_sns">
        <a href="http://www.wanfangdata.com.cn/index.html?index=true" class="nav_color_index">首页</a>
        <a href="http://www.wanfangdata.com.cn/sns" class="nav_color_index">社区</a>
    </div>
</div>
<script type="text/javascript">
    $(function(){
        function anxsGetCookie(name){
            var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
            if(arr=document.cookie.match(reg))
            return unescape(arr[2]);
            else
            return null;
        }

        function anxsSetCookie(name,value){
            document.cookie = name + "="+ escape (value) + ";path=/;domain=wanfangdata.com.cn";
        }

        //查看当前url地址
        var backurl = anxsGetCookie("firstvisit_backurl");
        if(backurl == "" || backurl == null ||　backurl　== undefined){
            //首次进入记录访问url地址写入cookie
            backurl = "http://" + window.location.hostname;
            if(backurl.indexOf("www.") > -1 ||
                    backurl.indexOf("g.") > -1){
                anxsSetCookie("firstvisit_backurl",backurl);
            }
        }

        //获取当前的domain
        var currentDomain = "http://" + window.location.hostname;
        if(currentDomain.indexOf("www.wanfangdata.com.cn") > -1
                || currentDomain.indexOf("g.wanfangdata.com.cn") > -1){
            if(backurl != currentDomain){
                anxsSetCookie("firstvisit_backurl",currentDomain);
                backurl = currentDomain;
            }
        }

        $(".anxs-top-88qwe-imgwrapper").attr("href",backurl);
        if(backurl.indexOf("www.") > -1){
            $(".anxs-top-88qwe-returnold").attr("href","http://old.wanfangdata.com.cn/");
        }else if(backurl.indexOf("g.") > -1){
            $(".anxs-top-88qwe-returnold").attr("href","http://old.g.wanfangdata.com.cn/");
        }else{
            $(".anxs-top-88qwe-imgwrapper").attr("href","http://www.wanfangdata.com.cn/");
            $(".anxs-top-88qwe-returnold").attr("href","http://old.wanfangdata.com.cn/");
        }

        var currentURL = window.location.href;
        var advanceUrl = "http://librarian.wanfangdata.com.cn/?dbid=paper";
        if(currentURL.indexOf("/searchResult/getAdvancedSearch.do") > -1){
            $(".anxs-top-88qwe-imgwrapper").attr("href",backurl);
            $(".anxs-top-88qwe-returnold").attr("href",advanceUrl);
        }
    });
</script>
    <!--
    状态栏
-->
<link rel="stylesheet" href="http://cdn.login.wanfangdata.com.cn/Content/js/lib/skin/default/layer.css" type="text/css" />
<script type="text/javascript" src="http://cdn.login.wanfangdata.com.cn/Content/js/lib/layer-2.4.js"></script>
<style>
    .clear{zoom:1}.clear:after{content:"";display:block;clear:both;visibility:hidden;height:0}.anxs-8qwe-top-rt{float:right;padding-top:35px;font-family:'Microsoft YaHei'}.anxs-8qwe-top-rt .anxs-8qwe-list{float:left;border-left:1px solid #d4d9dc;color:#555;font-size:14px;font-family:'Microsoft YaHei'}.anxs-8qwe-reading{position:relative}.anxs-8qwe-readingDay{width:74px;height:20px;position:absolute;left:32px;top:-22px;display:none}.anxs-8qwe-top-rt .anxs-8qwe-list_bind{padding-right:20px;border-left:none;border-right:1px solid #d4d9dc;cursor:pointer;position:relative;z-index:999999;display:none;}.anxs-8qwe-list_bind .anxs-8qwe-list_bind-wx{position:absolute;top:36px;right:0px;width:320px;background-color:#fff;text-align:center;box-shadow:0 0 10px #888;border:1px solid #cccccc\9;cursor:default;display:none;}.anxs-8qwe-list_bind .anxs-8qwe-list_bind-wx span:first-child{position:absolute;top:-20px;right:-2px;display:inline-block;width:110px;height:36px;}.anxs-8qwe-list_bind .anxs-8qwe-list_bind-wx span:first-child i{display:inline-block;position:absolute;top:12px;left:54px;width:11px;height:8px;background:url(http://cdn.login.wanfangdata.com.cn/Content/src/img/anxs-tri.png) no-repeat;}.anxs-8qwe-list_bind .anxs-8qwe-list_bind-wx .anxs-8qwe-list_bind-img{padding-top:30px;width:120px;height:120px;}.anxs-8qwe-list_bind-wx  .anxs-8qwe-bind-layer{display:inline-block;position:absolute;top:30px;left:100px;width:120px;height:120px;background:rgba(0,0,0,0.6)!important;background:#000;filter:Alpha(opacity=60);display:none;}.anxs-8qwe-list_bind-wx  .anxs-8qwe-bind-layer span{display:block;text-align:center;color:#ffffff;font-size:12px;line-height:28px;}.anxs-8qwe-list_bind .anxs-8qwe-list_bind-wx .anxs-8qwe-bind-title{display:inline-block;padding:16px;font-size:12px;line-height:22px;text-align:left;}.anxs-8qwe-list_bind-wx .anxs-8qwe-bind-title i{color:#f00;font-style:normal;display:none;}.anxs-8qwe-list_bind-wx .anxs-8qwe-bind-title em{font-style:normal;color:#417dc9;text-decoration:underline;cursor:pointer;}.anxs-8qwe-top-rt .anxs-no-line{border:0}.anxs-8qwe-top-rt .anxs-8qwe-list-login .anxs-8qwe-login-gr,.anxs-8qwe-top-rt .anxs-8qwe-list-login .anxs-8qwe-login-jg{display:none}.anxs-8qwe-top-rt .anxs-8qwe-list>a{color:#555;padding:0 10px;text-decoration:none;cursor:pointer;}.anxs-8qwe-top-rt .anxs-8qwe-list.anxs-8qwe-nav:hover{color:#555}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-login{color:#ff6c00}.anxs-8qwe-top-rt a:hover{text-decoration:none}.anxs-8qwe-top-rt .anxs-8qwe-nav{position:relative;display:inline-block;cursor:pointer;padding-left:0px}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list{display:none;width:348px;cursor:default;position:absolute;top:33px;left:-325px;z-index:1;padding:16px 25px 12px;background:#fff;border:1px solid #f5f5f5;box-shadow:0 1px 12px 3px #e6e6e6}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list .anxs-8qwe-nav-list-filter{width:100%;height:20px;position:absolute;top:-20px;right:0;z-index:1}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list .anxs-8qwe-nav-list-filter .anxs-8qwe-icon{display:block;width:9px;height:7px;background:url(http://cdn.login.wanfangdata.com.cn/Content/src/img/anxs-tri.png) no-repeat;position:relative;top:12px;left:50%;margin-left:-5px}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list dl{float:left;color:#555;font-size:12px;font-family:'Microsoft YaHei'}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list dl dt{font-weight:700;margin-bottom:5px}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list dl dd{float:left;width:90px;line-height:26px;text-align:left;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list dl dd a{text-decoration:none;color:#555;}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list dl dd a:hover{color:#ff6c00;text-decoration:none;}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list .anxs-8qwe-list-resource{width:180px}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list .anxs-8qwe-list-service{width:100px}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list .anxs-8qwe-list-view{width:68px}.anxs-8qwe-top-rt .anxs-8qwe-list .anxs-8qwe-nav-list .anxs-8qwe-list-view dd{width:74px}.anxs-8qwe-top-rt .anxs-8qwe-nav:hover .anxs-8qwe-nav-list{display:block}.anxs-8qwe-top-rt .anxs-8qwe-list-jg{display:none;padding-right:20px;}.anxs-8qwe-top-rt .anxs-8qwe-list-jg .anxs-8qwe-jgName b{display:inline-block;font-weight:500;max-width:120px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;vertical-align:bottom;}.anxs-8qwe-top-rt .anxs-icon-jg{display:inline-block;width:6px;height:4px;background:url(http://cdn.login.wanfangdata.com.cn/Content/src/img/anxs-jg.png) no-repeat;margin-left:10px;position:relative;top:-3px;}.anxs-8qwe-top-rt .anxs-8qwe-list-jg .anxs-8qwe-list-jg-box{display:none;width:158px;left:50%;padding:10px 20px;margin-left:-94px}.anxs-8qwe-top-rt .anxs-8qwe-list-jg .anxs-8qwe-list-jg-box .anxs-8qwe-nav-list-filter{width:100%;text-align:center}.anxs-8qwe-top-rt .anxs-8qwe-list-jg .anxs-8qwe-list-jg-box .anxs-8qwe-nav-list-filter .anxs-8qwe-icon{left:50%;}.anxs-8qwe-top-rt .anxs-8qwe-list-jg .anxs-8qwe-list-jg-box .anxs-8qwe-nav-list-main h5{margin:0;margin-bottom:8px;}.anxs-8qwe-top-rt .anxs-8qwe-list-jg .anxs-8qwe-list-jg-box .anxs-8qwe-nav-list-main h5 .anxs-8qwe-nav-list-main-name{color:#333;font-weight:700;display:inline-block;max-width:116px;white-space:nowrap;text-overflow:ellipsis;overflow:hidden;}.anxs-8qwe-top-rt .anxs-8qwe-list-jg .anxs-8qwe-list-jg-box .anxs-8qwe-nav-list-main h5 .anxs-8qwe-nav-list-main-back{float:right;color:#666;font-size:12px;cursor:pointer;}.anxs-8qwe-top-rt .anxs-8qwe-list-jg .anxs-8qwe-list-jg-box .anxs-8qwe-nav-list-main p{margin:0;}.anxs-8qwe-top-rt .anxs-8qwe-list-jg .anxs-8qwe-list-jg-box .anxs-8qwe-nav-list-main p .anxs-8qwe-login-jg{font-size:12px;color:#333;font-weight:700;}.anxs-8qwe-top-rt .anxs-8qwe-list-common-box{padding-right:20px;display:none;}.anxs-8qwe-top-rt .anxs-8qwe-list-common-box .anxs-8qwe-nav-list{width:108px;left:50%;margin-left:-54px;padding:0;}.anxs-8qwe-top-rt .anxs-8qwe-list-common-box .anxs-8qwe-nav-list .anxs-8qwe-nav-list-main h5{line-height:30px;margin:0;}.anxs-8qwe-top-rt .anxs-8qwe-list-common-box .anxs-8qwe-nav-list .anxs-8qwe-nav-list-main h5:hover{background:#f1f1f1}.anxs-8qwe-top-rt .anxs-8qwe-list-common-box .anxs-8qwe-nav-list .anxs-8qwe-nav-list-main h5 a{color:#555;display:inline-block;width:100%;text-indent:20px;text-decoration:none;cursor:pointer;}.anxs-8qwe-top-rt .anxs-8qwe-list-gr .anxs-8qwe-grName{display:inline-block;max-width:94px;margin-right:20px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;padding:0;}.anxs-8qwe-top-rt .anxs-8qwe-list-gr .anxs-8qwe-certification-no{display:inline-block;width:18px;height:18px;background:url(http://cdn.login.wanfangdata.com.cn/Content/src/img/anxs-certification-no.png) no-repeat;position:absolute;right:21px;top:0;}.anxs-8qwe-top-rt .anxs-8qwe-list-gr .anxs-8qwe-certification{background:url(http://cdn.login.wanfangdata.com.cn/Content/src/img/anxs-certification.png) no-repeat}.anxs-8qwe-top-rt .anxs-8qwe-list-gr .anxs-icon-jg{margin-left:0;top:8px;position:absolute;right:12px;}.anxs-8qwe-top-rt .anxs-8qwe-list-gr .anxs-8qwe-list-gr-box .anxs-8qwe-nav-list-filter{width:100%;right:0;}.anxs-8qwe-top-rt .anxs-8qwe-list-gr .anxs-8qwe-list-gr-box .anxs-8qwe-nav-list-main h5{line-height:30px;}.anxs-8qwe-top-rt .anxs-8qwe-list-gr .anxs-8qwe-no-certification-wrapper{padding:0;}.anxs-8qwe-top-rt .anxs-8qwe-list-message .anxs-8qwe-jgName a{color:#555;text-decoration:none;}.anxs-8qwe-top-rt .anxs-8qwe-list-message .anxs-8qwe-num{background:#ff6c00;color:#fff;text-align:center;display:inline-block;line-height:8px;max-width:28px;border-radius:7px;font-size:12px;padding:3px 2px 4px;position:absolute;margin-top:8px;right:6px;font-weight:500;display:none;}.anxs-8qwe-top-rt .anxs-8qwe-list-message .anxs-8qwe-num-all{position:static;padding:0 4px;}.anxs-8qwe-top-rt .anxs-8qwe-list-message .anxs-8qwe-nav-list .anxs-8qwe-nav-list-main h5 a{text-indent:27px;}.anxs-8qwe-top-rt .anxs-8qwe-list-message .anxs-icon-jg{margin-left:0;}
 .anxs-8qwe-top-rt .new-semester-message-2018 { position:absolute;display:none; top: 15px; right: 110px;width: 96px; height: 18px;background:url(http://cdn.login.wanfangdata.com.cn/Content/src/img/new-semester-2018.png) no-repeat;} .anxs-8qwe-top-rt  .anxs-8qwe-reading{border-left:0px} .anxs-8qwe-top-rt .anxs-8qwe-nav-main .nav_color{color: #ff6c00;padding:0 4px;}.anxs-8qwe-nav-main a{color: #ff6c00}.anxs-8qwe-top-rt .anxs-8qwe-list_bind{border-right:0px}.anxs-8qwe-top-rt .anxs-8qwe-list{border-left:0px}.anxs-8qwe-top-rt .anxs-8qwe-nav-main{border-left: 1px solid #d4d9dc}.anxs-8qwe-top-rt .anxs-8qwe-list.anxs-8qwe-nav-main:hover { color: #f60;}.anxs-8qwe-top-rt .anxs-8qwe-return_old{display: inline-block;cursor: pointer; padding-left: 8px;}.anxs-8qwe-top-rt .anxs-8qwe-return_old a{color: #333;font-size: 14px;}
</style>
<div class="anxs-8qwe-top-rt">
<a class="new-semester-message-2018" href="http://my.wanfangdata.com.cn/user/personal/index" target="_blank"></a>
    <div class="anxs-8qwe-list_bind anxs-8qwe-list">绑定机构
        <div class="anxs-8qwe-list_bind-wx">
            <span><i></i></span>
            <img src="http://cdn.login.wanfangdata.com.cn/Content/src/img/anxs-loding-wx.gif" class="anxs-8qwe-list_bind-img">
            <div class="anxs-8qwe-bind-layer">
                <img src="http://cdn.login.wanfangdata.com.cn/Content/images/bind_right_icon.png" style="width:18px;height: 15px;margin-top:14px"/>
                <span>扫描成功</span>
                <span>请在APP上操作</span>
            </div>
            <span></span>
            <span class="anxs-8qwe-bind-title">
			打开万方数据APP，点击右上角"扫一扫"，扫描二维码即可将您登录的个人账号与机构账号绑定，绑定后您可在APP上享有机构权限，如需更换机构账号，可到个人中心解绑。
		</span>
        </div>
    </div>
    <div class="anxs-8qwe-list anxs-8qwe-nav anxs-8qwe-list-jg">
        <span class="anxs-8qwe-jgName">欢迎<b></b>的朋友</span><i class="anxs-icon-jg"></i>
        <div class="anxs-8qwe-nav-list anxs-8qwe-list-jg-box">
            <div class="anxs-8qwe-nav-list-filter">
                <i class="anxs-8qwe-icon"></i>
            </div>
            <div class="anxs-8qwe-nav-list-main">
                <p><a class="anxs-8qwe-login anxs-8qwe-login-jg">登录机构账号</a></p>
            </div>
        </div>
    </div>
    <div class="anxs-8qwe-list anxs-8qwe-list-login anxs-no-line">
        <a class="anxs-8qwe-login-all">登录 / 注册</a>
        <a class="anxs-8qwe-login-gr">登录 / 注册</a>
        <a class="anxs-8qwe-login-jg">机构登录</a>
    </div>
    <div class="anxs-8qwe-list anxs-8qwe-nav anxs-8qwe-list-common-box anxs-8qwe-list-gr">
        <a href="http://my.wanfangdata.com.cn/user/index" target="_blank" class="anxs-8qwe-person-center anxs-8qwe-grName"></a>
        <a href="http://my.wanfangdata.com.cn/auth/user/authenticationintro.do" target="_blank" class="anxs-8qwe-no-certification-wrapper"><i class="anxs-8qwe-certification-no"></i></a>
        <i class="anxs-icon-jg"></i>
        <div class="anxs-8qwe-nav-list anxs-8qwe-list-gr-box">
            <div class="anxs-8qwe-nav-list-filter">
                <i class="anxs-8qwe-icon"></i>
            </div>
            <div class="anxs-8qwe-nav-list-main">
                <h5><a class="anxs-8qwe-person-center" target="_blank" href="http://my.wanfangdata.com.cn/user/index">个人中心</a></h5>
                <h5><a class="anxs-8qwe-social-my" target="_blank" href="http://social.wanfangdata.com.cn/">我的学术圈</a></h5>
                <h5><a target="_blank" href="http://work.wanfangdata.com.cn/">我的书案</a></h5>
                <h5><a data-useid="" class="anxs-8qwe-nav-list-main-back anxs-8qwe-back">退出</a></h5>
            </div>
        </div>
    </div>
    <div class="anxs-8qwe-list anxs-8qwe-reading">
        <a target="_blank" href="http://login.wanfangdata.com.cn/activity/graduationSeason?#recharge" class="recharge-2019" style="padding: 0;position: absolute;display: none;width: 88px;height: 23px;top: -24px;right: -60px;background: url(http://cdn.login.wanfangdata.com.cn/Content/src/img/rechargIcon.png) no-repeat;"></a>
        <a class="anxs-8qwe-saveMoney" target="_blank" href="http://my.wanfangdata.com.cn/user/wallet/index">钱包</a>
    </div>
    <div class="anxs-8qwe-list anxs-8qwe-nav anxs-8qwe-list-common-box anxs-8qwe-list-message">
        <span class="anxs-8qwe-jgName"><a href="http://my.wanfangdata.com.cn/user/message/index" target="_blank">消息</a></span>
        <span class="anxs-8qwe-num anxs-8qwe-num-all"></span>
        <i class="anxs-icon-jg"></i>
        <div class="anxs-8qwe-nav-list anxs-8qwe-list-message-box">
            <div class="anxs-8qwe-nav-list-filter">
                <i class="anxs-8qwe-icon"></i>
            </div>
            <div class="anxs-8qwe-nav-list-main">
                <h5><a target="_blank" href="http://my.wanfangdata.com.cn/user/message/localMessages">系统通知</a><span class="anxs-8qwe-num anxs-8qwe-system"></span></h5>
                <!--<h5><a>成果动态</a><span class="anxs-8qwe-num"></span></h5>-->
                <h5><a target="_blank" href="http://work.wanfangdata.com.cn/subscribe/index.do">订阅消息</a><span class="anxs-8qwe-num anxs-8qwe-subscription"></span></h5>
                <h5><a class="anxs-8qwe-addPeople" target="_blank" href="http://social.wanfangdata.com.cn/follow/toFollowPage.do?removeFlag=0&followState=followMe">新增关注</a><span class="anxs-8qwe-num anxs-8qwe-attention"></span></h5>
                <!--<h5><a>私信</a><span class="anxs-8qwe-num"></span></h5>-->
            </div>
        </div>
    </div>
    <div class="anxs-8qwe-login  anxs-8qwe-list anxs-8qwe-nav anxs-8qwe-nav-main">
        <a href="http://www.wanfangdata.com.cn/resource_nav/index.do?resouccetype=resouccetype" class="nav_color" target="_blank">资源导航</a>
    </div>
    <div class="anxs-8qwe-return_old anxs-8qwe-list">
        <a href="http://old.wanfangdata.com.cn/" target="_blank" class="anxs-top-88qwe-returnold">返回旧版</a>
    </div>
</div>
<iframe id="anxs-8qwe-login" style="margin-left:-10px;margin-top:-100px;display:none" src="" height="750px;" width="710px;" scrolling="no" frameborder="0" class="layui-layer-wrap">
</iframe>
<script>
    var refer = document.referrer;
    if(refer!=null && (refer.indexOf("http://tongji.baidu.com/")!=-1 || refer.indexOf("https://tongji.baidu.com/")!=-1)){
    }else{
        if (window != parent) {
            top.location.href = location.href;
        }
    }
    document.domain="wanfangdata.com.cn";
    $(function(){
        var getProp = (function () {
            var me = $('#anxs-8qwe-login');
            var loginParent = $('.anxs-8qwe-list-login');
            var allLogin = loginParent.find('.anxs-8qwe-login-all'); //默认登陆按钮
            var grLogin = loginParent.find('.anxs-8qwe-login-gr');   //个人登陆按钮
            var grShow = $('.anxs-8qwe-list-gr');
            var grName = grShow.find('.anxs-8qwe-grName');           //个人名字
            var certificationImg = grShow.find('.anxs-8qwe-certification-no');  //认证图片
            var jgLogin = $('.anxs-8qwe-login-jg');   //机构登陆按钮
            var jgname = $('.anxs-8qwe-list-jg');
            var jgMain = jgname.find('.anxs-8qwe-nav-list-main');    //机构下拉列表
            var listJgName = jgname.find('.anxs-8qwe-login-jg');     //机构下拉列表中的机构登陆按钮
            var registerName = $('.anxs-8qwe-register');             //注册按钮
            var personBack = $('.anxs-8qwe-back');                   //个人下拉列表中的退出按钮
            var messageName = $('.anxs-8qwe-list-message');          //消息按钮
            var messageNum = messageName.find('.anxs-8qwe-num');     //消息按钮下拉列表下的消息数
            var castgc = '';
            var socialBtn = $('.anxs-8qwe-social-my');               //我的学术圈按钮
            var addPeople = $('.anxs-8qwe-addPeople');               //新增关注

            //Binding mecFhanism variable
            var bindMe = $('.anxs-8qwe-list_bind');
            var bindMeWx = $('.anxs-8qwe-list_bind-wx');//扫码
            var bindMeImg = $('.anxs-8qwe-list_bind-img');//图片
            var bindMeTitle = bindMeWx.find('.anxs-8qwe-bind-title');//文字提示
            var bindTime = '';
            var flagBindShow = $('#indexScanCodeBind');
              
            return {
                //绑定机构的出现
                bindMeShow:function () {
                    var that = this;
                    $.ajax({
                        url: 'http://login.wanfangdata.com.cn/showBindTip',
                        dataType: 'jsonp',
                        jsonp: 'callback',
                        cache:false,
                        success:function(data){
                            if((data == 'true') && (flagBindShow.length)) {
                                bindMe.show();
                                that.bindMechanism();
                            }else {
                                bindMe.hide();
                                $('.anxs-8qwe-bind-layer').hide();
                            }
                        }
                    });
                },
                //绑定机构
                bindMechanism:function () {
                    var that = this;
                    bindMe.hover(function (e) {
                        bindMeWx.show();
                        that.getBindqcInfo();
                        that.refreshQrCode();
                    },function () {
                        bindMeWx.hide();
                        clearInterval(that.bindTime);
                        clearInterval(that.sweepCode);
                    });
                },
                //拿到二维码信息
                getBindqcInfo:function () {
                    var that = this;
                    clearInterval(that.sweepCode);
                    $.ajax({
                        url:'http://login.wanfangdata.com.cn/getQRcodeMessage?time='+(+new Date()),
                        dataType:'jsonp',
                        jsonp:'callback',
                        success:function (data) {
                            if(data == 'isBind') {
                                window.location.reload();
                            }else if(data == 'error'){
                                window.location.reload();
                            }else {
                                var Ciphertext = data.ciphertext,
                                        CodeId = data.codeId;
                                bindMeImg.attr('src','http://login.wanfangdata.com.cn/createQRCode?ciphertext='+encodeURIComponent(Ciphertext))
                                that.sweepCodeSuccess(CodeId);
                                that.refreshsweepCodeSuccess(CodeId);
                            }
                        },
                        error:function() {
                            console.log('二维码请求失败!');
                        }
                    });
                },
                //1秒钟自动刷新扫描状态请求
                refreshsweepCodeSuccess:function (CodeId) {
                    var that = this;
                    clearInterval(that.sweepCode);
                    this.sweepCode = setInterval(function () {
                        that.sweepCodeSuccess(CodeId);
                    },1500)
                },
                //一分钟自动刷新
                refreshQrCode:function () {
                    var that = this;
                    clearInterval(that.bindTime);
                    this.bindTime = setInterval(function () {
                        that.getBindqcInfo();
                    },60000)
                },
                //扫描成功
                sweepCodeSuccess:function (CodeId) {
                    var that = this;
                    $.ajax({
                        url: 'http://login.wanfangdata.com.cn/getQRCodeStatus',
                        dataType:'jsonp',
                        jsonp:'callback',
                        data:{codeId:CodeId},
                        success:function(data){
                            if(data == 'scanned') {
                                clearInterval(that.sweepCode);
                                clearInterval(that.bindTime);
                                $('.anxs-8qwe-bind-layer').show();
                            }else if (data == 'available') {
                                $('.anxs-8qwe-bind-layer').hide();
                            }else{
                                $('.anxs-8qwe-bind-layer').hide();
                            }
                        },
                        error:function (err) {
                            console.log(err)
                        }
                    });
                },
                //弹出登录框
                layerAlert:function(){
                    var winHeight = $(window).height();
                    if(winHeight<800){
                        layer.open({
                            offset: '40px',
                            type: 1, //page层 1div，2页面
                            area: ['50%px'],
                            title: '',
                            shade: 0.6, //遮罩透明度
                            moveType: 1, //拖拽风格，0是默认，1是传统拖动
                            shift: 1, //0-6的动画形式，-1不开启
                            content: $("#anxs-8qwe-login"),
                            success:function(layero){
                                $(layero).css({'position':'absolute','top':'40px','left':'50%','marginLeft':'-350px','marginBottom':'40px'});
                                $(window).resize(function(){
                                    var winHeight = $(window).height();
                                    if(winHeight <800){
                                        $(layero).css({'position':'absolute','top':'40px','left':'50%','marginLeft':'-350px','marginBottom':'40px'});
                                    }else{
                                        $(layero).css({'position':'','top':'','left':'50%','marginLeft':'-350px','marginBottom':''});
                                    }
                                });
                            },
                            end: function(){

                            }
                        });
                    }else{
                        layer.open({
                            offset:'auto',
                            type: 1, //page层 1div，2页面
                            area: ['50%px'],
                            scrollbar: true,
                            title: '',
                            shade: 0.6, //遮罩透明度
                            moveType: 1, //拖拽风格，0是默认，1是传统拖动
                            shift: 1, //0-6的动画形式，-1不开启
                            content: $("#anxs-8qwe-login"),
                            success:function(layero){
                                $(window).resize(function(){
                                    var winHeight = $(window).height();
                                    if(winHeight <800){
                                        $(layero).css({'position':'absolute','top':'40px','left':'50%','marginLeft':'-350px','marginBottom':'40px'});
                                    }else{
                                        $(layero).css({'position':'','top':'','left':'50%','marginLeft':'-350px','marginBottom':''});
                                    }
                                });
                            },
                            end: function(){

                            }
                        });
                    }
                },
                //获取url参数
                getRequest: function () {
                    var url = location.search; //获取url中"?"符后的字串
                    if (!url) {
                        return '';
                    }
                    var theRequest = new Object();
                    if (url.indexOf("?") != -1) {
                        var str = url.substr(1);
                        var strs = str.split("&");
                        for (var i = 0; i < strs.length; i++) {
                            theRequest[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1]);
                        }
                    }
                    return theRequest;
                },
                //获取cookie名字
                getCookies: function (cookieName) {
                    var strCookie = document.cookie;
                    var arrCookie = strCookie.split("; ");
                    for (var i = 0; i < arrCookie.length; i++) {
                        var arr = arrCookie[i].split("=");
                        if (cookieName == arr[0]) {
                            return arr[1];
                        }
                    }
                    return "";
                },
                //赋值给iframe
                getSrc: function (loginModel) {
                    this.layerAlert();
                    me.attr('src', 'http://my.wanfangdata.com.cn/auth/user/'+loginModel+'?login_mode=AJAX&service=' + encodeURIComponent(window.location.href));
                },
                //消息发请求
                initMessage:function(){
                    $.ajax({
                        url: 'http://my.wanfangdata.com.cn/user/message/getUnreadAmountJSONP',
                        dataType: "jsonp",
                        timeout:3000,
                        jsonp: "callback",
                        success: function (data) {
                            if(data){
                                var unreadAmountNum = data.unreadAmount,
                                        unreadSystemAmountNum = data.unreadSystemAmount,
                                        unreadSubAmountNum = data.unreadSubAmount,
                                        unreadFollowAmountNum = data.unreadFollowAmount;
                                if(unreadAmountNum > 99){
                                    unreadAmountNum = '99+';
                                }
                                if(unreadSystemAmountNum > 99){
                                    unreadSystemAmountNum = '99+';
                                    messageName.find('.anxs-8qwe-nav-list-main h5 a').css('textIndent','17px');
                                }
                                if(unreadSubAmountNum > 99){
                                    unreadSubAmountNum = '99+';
                                    messageName.find('.anxs-8qwe-nav-list-main h5 a').css('textIndent','17px');
                                }
                                if(unreadFollowAmountNum > 99){
                                    unreadFollowAmountNum = '99+';
                                    messageName.find('.anxs-8qwe-nav-list-main h5 a').css('textIndent','17px');
                                }
                                if(unreadAmountNum){
                                    messageName.find('.anxs-8qwe-num-all').show().text(unreadAmountNum);//获取未读消息总量
                                }else{
                                    messageName.find('.anxs-8qwe-num-all').hide();
                                }
                                if(unreadSystemAmountNum){
                                    messageName.find('.anxs-8qwe-system').show().text(unreadSystemAmountNum);//获取未读系统消息
                                }else{
                                    messageName.find('.anxs-8qwe-system').hide();
                                }
                                if(unreadSubAmountNum){
                                    messageName.find('.anxs-8qwe-subscription').show().text(unreadSubAmountNum);//获取未读订阅消息
                                }else{
                                    messageName.find('.anxs-8qwe-subscription').hide();
                                }
                                if(unreadFollowAmountNum){
                                    messageName.find('.anxs-8qwe-attention').show().text(unreadFollowAmountNum);//获取未读关注消息
                                }else{
                                    messageName.find('.anxs-8qwe-attention').hide();
                                }

                            }
                        }
                    });
                },
                //获取ip的值
                isIp:function(nameData){
                    for(var prop in nameData){
                        if(prop == 'OnlyIPLogin'){
                            return nameData[prop];
                            break;
                        }
                    }
                    return false;
                },
                //判断是否存在机构
                isGroup:function(nameData){
                    for(var prop in nameData){
                        var propId = prop.split('.');
                        if(propId[0] == 'Group'){
                            return true;
                            break;
                        }
                    }
                    return false;
                },
                //判断是否是个人
                isPerson:function(nameData){
                    for(var prop in nameData){
                        var propId = prop.split('.');
                        if(propId[0] == 'Person'){
                            return true;
                        }
                    }
                    return false;
                },
                //获取cookie中的值，判断是否登录
                cookieAjax: function () {
                    var that = this;
                    $.ajax({
                        url: "http://login.wanfangdata.com.cn/getUserState",
                        dataType: "jsonp",
                        jsonp: "callback",
                        success: function (data) {
                            var dataParent = data.context;
                            var accountId = dataParent.accountIds;
                            var idLength = accountId.length;
                            if(!idLength){
                                return true;
                            }
                            var nameData = dataParent.data;
                            var nameArr = [];

                            // begin 判断是否有全局配置排序方法，如果有执行检测产品排序方法
                            if (window._accountOrderFun) {
                                //替换原有的accountId
                                try {
                                    accountId = _accountOrderFun(accountId)
                                } catch (error) {
                                    console.log(error)
                                }
                            }
                            // end

                            for(var i=0;i<idLength;i++){
                                var idArr = accountId[i].split('.');
                                if(idArr[0] == 'Group' ||  idArr[0] == 'GroupSub'){
                                    //显示样式
                                    allLogin.hide().parent().removeClass('anxs-no-line');
                                    jgLogin.hide();
                                    listJgName.show();
                                    grLogin.show();
                                    jgname.show().addClass('anxs-no-line');
                                    //取机构的最新值
                                    for(var prop in nameData){
                                        var propId = prop.split('.');
                                        if(propId[0] == 'Group' || propId[0] == 'GroupSub'){
                                            if(accountId[i]+'.'+propId[propId.length-1] == prop){
                                                nameArr.push(nameData[prop]);
                                            }
                                        }
                                    }
                                    jgname.find('.anxs-8qwe-jgName b').text(nameArr[0]);
                                }
                                if(idArr[0] == 'Person'){
                                    allLogin.hide();
                                    var hasGroup = that.isGroup(nameData);
                                    if(!hasGroup){
                                        jgLogin.show();
                                    }else{
                                        jgLogin.hide();
                                    }
                                    listJgName.show();
                                    that.backEvent(); //初始化退出
                                }
                                if(idArr[0] == 'Researcher'){
                                    certificationImg.addClass('anxs-8qwe-certification');
                                    certificationImg.parent().attr('href','http://my.wanfangdata.com.cn/user/personal/rights');
                                }
                            }

                             //给机构下拉框赋值
                            for(var i=idLength-1;i>=0;i--){
                                var idArr = accountId[i].split('.');
                                if(idArr[0] == 'Group' || idArr[0] == 'GroupSub'){
                                    var groupLength = idArr[0].length+1;
                                    jgMain.prepend('<h5><a href="http://www.wanfangdata.com.cn/institution/showAuth.do" target="_blank" class="anxs-8qwe-nav-list-main-name">'+accountId[i].slice(groupLength)+'</a><span data-useId="'+accountId[i].slice(groupLength)+'" class="anxs-8qwe-nav-list-main-back">退出</span></h5>');
                                    var onlyIpStr=nameData['OnlyIPLogin'];
                                    if(typeof onlyIpStr != "undefined" && onlyIpStr != null && onlyIpStr != ""){
                                        var onlyIpStrs = onlyIpStr.split(",");
                                        jgMain.find('h5').each(function(index){
                                            var nameValue = $(this).find('.anxs-8qwe-nav-list-main-name').html();
                                            for(var i=0;i<onlyIpStrs.length;i++){
                                                if(nameValue == onlyIpStrs[i]){
                                                    $(this).find('span').remove();
                                                }
                                            }
                                        });
                                    }
                                }
                            }

                            that.backEvent(); //初始化退出
                            var hasPerson = that.isPerson(nameData);
                            if(hasPerson){
                                that.initMessage();//初始化消息数量
                            }
                            for(var prop in nameData){
                                var propId = prop.split('.');
                                if(propId[0] == 'Person'){
                                    grLogin.hide();
                                    if($.trim(nameData[prop])){
                                        grName.text(nameData[prop]);
                                    }else{
                                        grName.text(propId[1]);
                                    }
                                    personBack.attr('data-useId',propId[1]);
                                    registerName.hide();
                                    messageName.show();
                                    grShow.show();
                                }
                            }
                            for(var prop in nameData){
                                var propId = prop.split('.');
                                if(propId[0] == 'Researcher'){
                                    if($.trim(nameData[prop])){
                                        grName.text(nameData[prop]);
                                    }else{
                                        grName.text(propId[1]);
                                    }

                                }
                            }
                        },
                        error: function () {
                            console.log('fail');
                        }
                    });
                },
                //登录按钮
                loginEvent:function(){
                    var that = this;
                    allLogin.click(function(){
                        that.getSrc('alllogin.do');
                    });
                    grLogin.click(function(){
                        that.getSrc('alllogin.do');
                    });
                    jgLogin.click(function(){
                        that.getSrc('jglogin.do');
                    });
                    registerName.click(function(){
                        var registerHref = "/auth/user/register.do";
                        var locaHref = window.location.pathname;
                        if(registerHref == locaHref) {
                            $(this).attr('href',window.location.href);
                        }else {
                            $(this).attr('href',"http://my.wanfangdata.com.cn/auth/user/register.do?service="+ encodeURIComponent(window.location.href));
                        }
                    });
                },
                //退出登录
                backEvent:function(){
                    var that = this;
                    $('.anxs-8qwe-nav-list-main-back').click(function(){
                        var backurl = window.location.href;
                        backurl = encodeURIComponent(backurl).replace("&","%26");//复制原先的代码，具体为什么replace不清楚
                        var username = encodeURIComponent($(this).data('useid'));
                        castgc = that.getCookies('CASTGC');
                        window.location.href="http://my.wanfangdata.com.cn/auth/rest/logout.do?user_id="+username+"&service="+backurl+"&token="+castgc;
                    });
                },
                //毕业季特惠活动自动上下线（下一次修改删除）
                newSemester:function () {
                    $.ajax({
                        url: 'http://login.wanfangdata.com.cn/getServerTime',
                        success:function(data){
                            var startDate = Date.parse("2019/04/23 00:00");
                            var endDate = Date.parse("2019/06/01 00:00");
                            if(startDate <= data && endDate > data) {
                                $(".recharge-2019").show();
                                $(".current_select_way").parent("li").siblings().children("div").children(".recharge-2019").hide()
                            }else{
                                $(".recharge-2019").remove();
                            }
                        }
                    });
                },
                init: function () {
                    this.loginEvent();
                    this.cookieAjax();
                    this.bindMeShow();
                    this.newSemester();
                }
            };
        })().init();
    })
</script>
    
  </div>
</div>
<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/common/js/exchangeSearchTypeValue.js?version=1.2020.236-20200506094020"></script>
<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/common/js/public_head.js?version=1.2020.236-20200506094020"></script>
<script type="text/javascript"  src="http://miner.wanfangdata.com.cn/wflog/Content/js/browers-log.js?version=1.2020.236-20200506094020"></script>
<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/common/min/jquery.cookie.js"></script>
<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/common/min/list.js"></script>
<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/common/min/lang.js"></script>
<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/js/headerCom.js?version=1.2020.236-20200506094020"></script>
<script src="http://cdn.wanfangdata.com.cn/js/share.js?version=1.2020.236-20200506094020"></script>
<!--第三方统计分析-->
<script>
  var _hmt = _hmt || [];
  (function() {
    var hm = document.createElement("script");
    hm.src = "https://hm.baidu.com/hm.js?838fbc4154ad87515435bf1e10023fab";
    var s = document.getElementsByTagName("script")[0];
    s.parentNode.insertBefore(hm, s);
    var winUrl =window.location.href;
    if(winUrl.indexOf("/tech/techindex.do") > -1){
      $('head').append('<link rel="stylesheet" href="http://cdn.wanfangdata.com.cn/page/css/ClassNav.css?version=1.2020.236-20200506094020" rel="stylesheet" type="text/css">');
    }
  })();
</script>

<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/common/js/checkHtmlVersion.js?version=1.2020.236-20200506094020"></script>
<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/common/js/exchangeSearchTypeValue.js?version=1.2020.236-20200506094020"></script>
<input type="hidden" class="navigation_type" value=""/>
<input type="hidden" id="operationPermiss" value="true"/>
<style>
    .safari {line-height:0px}
    .otherssss {line-height:40px}
    .indexIcon{background: url("http://cdn.wanfangdata.com.cn/page/images/serchIndexIcon.png") no-repeat center center;display: inline-block;}
    .wfzs_logo-img{width:138px;height: 32px; background-position: -90px -109px;}
</style>

<!--搜索条 开始-->
<div class="search_block sub_search">
    <div class="container">
        <div class="wfzs_logo"><a href="/" class="indexIcon wfzs_logo-img"></a></div>
        <div class="searchContainer">
            <input id="subtitle" type="hidden" value="学位">
            <form action="" method="get" id="subform" onsubmit="$('#loading').show(); all_showtype_keycode=13;   return sub();">
                <input type="hidden" value="" name="searchType" id="searchType" />
                <input type="hidden" value="" name="showType" id="showTypeTmp">
                <input type="hidden" value="" name="pageSize" id="pageSizeTmp" />
                <div class="resource_type clear">
                    <a href="javascript:void(0);"   >全部</a>                    <a href="javascript:void(0);"   >期刊</a>                    <a href="javascript:void(0);"    class="selected"   >学位</a>                    <a href="javascript:void(0);"   >会议</a>                    <a href="javascript:void(0);"   >专利</a>                    <a href="javascript:void(0);"   >科技报告</a>                    <a href="javascript:void(0);"   >成果</a>                    <a href="javascript:void(0);"   >标准</a>                    <a href="javascript:void(0);"   >法规</a>                    <a href="javascript:void(0);"   >地方志</a>                    <a href="javascript:void(0);"   >视频</a>                    <a href="javascript:void(0);" style="display: none;"  >新图书</a>                    <a href="/resource_nav/index.do" id="search_type_more" >更多 >></a>
                </div>
                <div class="searchsug" style="left: 180px;display: none;">                <ul>
                </ul>
        </div>
        <div class="searchBox">
            <div class="searchInput" style="padding-right:106px;">
                <input MeetName="keyWords" type="text" autocomplete="off" id="keyWords"  name="searchWord" value="" style="height: 40px; line-height: normal; line-height: 40px\9;" placeholder="海量资源，等你发现" id="headerBorder">
                <input MeetName="isFirst" type="hidden" value="true">
                <input id="triggerTag" type="hidden" name="isTriggerTag" >
            </div>

            <div class="searchBtn" style="display: none;">
                <input type="button" value="搜索"  />
            </div>
            <div class="pubsearchBtn">
                <input type="button" value="学位" />
            </div>
        </div>
        <div class="gjsearchBtn"><a href="/searchResult/getAdvancedSearch.do?searchType=degree" target="_blank" id='advancedSearch'>高级检索</a><br/>
            <a href="/history/getHistory.do?pageNum=1&pageSize=20" target="_blank">检索历史</a>
        </div>

        <div class="searchInput_bomb"  style="margin-top:10px;left: 180px;">        <i class="triangle"></i>
        <div class="bomb_list">
        </div>
    </div>
    </form>
</div>
</div>
</div>
<!--搜索条 结束-->

<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>检索详情页</title>
	      <link rel="stylesheet" href="http://cdn.wanfangdata.com.cn/page/css/SearchResult.css?version=1.2020.236-20200506094020" type="text/css"></link>
			<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/js/laypage-v1.3/laypage/laypage.js?version=default"></script>
			<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/search_detail/js/details.js?version=1.2020.236-20200506094020"></script>
			<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/search_detail/js/aysnDetail.js"></script>
			<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/search_detail/js/format.js"></script>
			<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/js/min/echarts.js"></script>
			<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/page/search_list/min/cookieUtil.js"></script>
			<link rel="stylesheet" href="http://cdn.wanfangdata.com.cn/page/css/min/laypage.css" type="text/css"></link>
			<script type="text/javascript" src="http://cdn.wanfangdata.com.cn/js/exportAll.js?version=1.2020.236-20200506094020"></script>
			<script src="http://cdn.wanfangdata.com.cn/page/search_list/js/permissionDe.js?version=1.2020.236-20200506094020"></script>
	<script>
		var perPath ='http://www.wanfangdata.com.cn/sns/third-web/per';
	</script>
</head>
<body>
<div id="div_a" style="background-color: #f5f5f5;">
<div class="container" style="background-color: #fff; padding: 0 20px;">
	<!-- 文献类型 -->
	<input hidden="hidden" id="IsticTeachURI" value='http://new.istic.wanfangdata.com.cn/'>
	<input type="hidden" id="aysn_keyword" value='[互联网广告, 投放系统, 软件设计, 功能模块]'>
	<input type="hidden" id="aysn_trans_key" value=''>
	<input type="hidden" id="aysn_trans_key" value=''>
	<input type="hidden" id="document_type" value="degree"/>
	<textarea style="display: none;" id="related_topics" >{"words":"$head_words:(互联网广告)+$head_words:(投放系统)+$head_words:(软件设计)+$head_words:(功能模块)","themeword":"$head_words","params":"$title:基于Golang的广告投放系统的设计与实现"}</textarea>
	<input type="hidden" id="jgUserId" value="">
		<input type="hidden" id="unit_name" value='湖南大学'>
	<input type="hidden" id="applicant_name" value=''>
	<input type="hidden" id="authors_name" value='肖霄'>
		<input type="hidden" id="article_id" value='Y3165623'>
	<input type="hidden" id="fund_info" value=''>
	<input type="hidden" id="imgUrl" value='http://video.wanfangdata.com.cn\wfresourse\pic\video\'>
	<input type="hidden" id="commonId" value="Y3165623"/>
	<input type="hidden" id="authorUrl" value="http://common.wanfangdata.com.cn"/>
	<input type="hidden" id="videoUrl" value="http://video.wanfangdata.com.cn\v\play\"/>
	<input type="hidden" id="ebookUrl" value="http://www.apabi.com/apabi/pub.mvc/Index2?pid=login&cult=CN" />
	<input type="hidden" id="userType" value="" />
	<input type="hidden" id="user" value="{}">
	<input type="hidden" id="doiUrl" value="http://dx.chinadoi.cn/">
	<input type="hidden" id="socialUrl" value="http://social.wanfangdata.com.cn/">
	<input type="hidden" id="userRoles" value="" />
	<input type="hidden" id="userId" value="" />
	<input type="hidden" id="searchUrl" value="www.wanfangdata.com.cn">
	<input type="hidden" id="nstlUrl" value="http://nstl.wanfangdata.com.cn/">
	<input type="hidden" id="myUrl" value="my.wanfangdata.com.cn">
	<input type="hidden" id="searchType" value="degree">
	<input type="hidden" id="first_publish" value="">
	<!--面包屑-->
	<div class="crumbs"><a href="#" onclick="toIndexPage()">
	<!--
		首页
	 -->
	 首页
	</a>&gt;
				<a href="" onclick="titleNavigation('学位')">
					<!--
						学位首页
					 -->
						学位首页
				</a>&nbsp;&gt;&nbsp;<font style="font-weight:bold;">基于Golang的广告投放系统的设计与实现</font>
	</div>
	<ul class="statistics-useful">
		<li class="statistics-useful-item" >
			<div class="statistics-use" data-flag="true"></div>
			<!--<div class="statistics-use-active statistics-use"></div>-->
			<div class="useful-num">有用 <em></em></div>
		</li>
		<li class="statistics-useful-item">
			<div class="statistics-useless" data-flag="true"></div>
			<!--<div class="statistics-useless statistics-use-active"></div>-->
			<div class="useful-num">没用<em></em></div>
		</li>
	</ul>
    <!--左主体-->
    <div class="left_con">
    	<!--文章介绍-->
    	<div class="left_con_top" style="padding: 15px 0;">
			 <!--  图书判断 -->
            <div class="add-label-container">
                <div class="add-label-title">添加标签</div>
                <i class="close-icon">×</i>
                <div class="add-label-content">
                    <div class="added-label-container">
                        <span class="added-count-container">已添加（<span class="added-label-count">0</span>/5）:</span>
                        <span class="added-labels-list"></span>
                    </div>
                    <div class="input-container">
                        <input type="text" placeholder="请输入标签（不超过10个字符）" maxlength="10">
                        <button>添加</button>
                    </div>
                    <div class="added-labels">
                        <p class="suggest-label-title">推荐标签：</p>
                        <p class="suggest-labels-list"></p>
                    </div>
                </div>
                <div class="operate-bannel">
                    <button class="submit">提交</button>
                    <button class="reset">取消</button>
                </div>
            </div>
        	<div class="title" style="display: inline-block;width: 790px;">
				<div id="chapterindex" style="display: none;float: left;">
					<input id="chapterstype" value="degree" type="hidden">
					<input id="chapters" value="Y3165623" type="hidden">
        			<a href="#" class="Catalog" onclick="chap('Y3165623')">目录</a>
				</div>
					基于Golang的广告投放系统的设计与实现
					<a title="WFMetrics"  href="javascript:void(0)" class="show-infor-btn">
						<i id="icon_Minerq" class="icon icon_Miner" style="margin-left: 0px;"></i>
<div class="icon_box" style="display: none;">
				            	<div class="horn"></div>
					            <div class="plate">
				                    <div class="use">文摘阅读	<b></b></div>
				                </div>
				                <div class="plate">
				                    <div class="use obtain">下载	<b></b></div>
				                </div>
				                <div class="plate">
				                    <div class="use quote">第三方链接	<b></b></div>
				                </div>
				                <div class="plate">
					                    <div class="use spread">被引	<b></b></div>
				                </div>
</div>					</a>
					<input type="hidden" value="Y3165623" id="articleId" />
					<input type="hidden" id="typed" value="degree" />
			</div>
				<input type="hidden" id="loginbyurlDetai" value="http://login.wanfangdata.com.cn/" />
				<p class="isShowPdfBtn"  style="float:right;display: none;font-size:13px;margin-right: 20px;color: #5b87ea;line-height: 35px;"><img src="/page/css/images/pdfToWordIcon.png" style="width: 20px;height: 22px;vertical-align: middle;margin-top: -3px;margin-right: 5px;"> <span style="text-decoration: underline;cursor: pointer" title="PDF转换功能由金山办公软件提供。">PDF&nbsp;转&nbsp;Word</span></p>

			<!--副标题-->

			<div class="Mbtn">
				<!-- 判断有全文 -->
							<a onclick="upLoadPermissions('63','Y3165623','chi','[WF, CNKI]','基于Golang的广告投放系统的设计与实现','0','degree')" id="ddownb" class="read detail_icon_down" ><i class="icon icon_down" title="下载"></i>下载</a>
							<a onclick="onlineReadingPermissions('63','Y3165623','chi','[WF, CNKI]','基于Golang的广告投放系统的设计与实现','0','degree');" target="_blank" class="read detail_icon_book" ><i class="icon icon_book" title="在线阅读"></i>在线阅读</a>
<!-- 判断无全文 -->             <!-- 导出 -->
				<a href="javascript:void(0)" onclick="exportOne('degree','degree_artical','XW_2165A362E3C056F65CF4E3951CFAB070','Y3165623')"  class="read detail_icon_export"><i class="icon icon_export"></i>导出</a>
				<form style="display: none" target="wf_export_data" action="/export/export.do" method="post" id="formExport">
					<input id="inputExport" type="hidden" name="export">
					<input type="hidden" name="exportType" value="degree">
				</form>
            <!-- 导出结束 -->
                <!--注：book资源暂时屏蔽“收藏”功能，待工作台加上相应资源可放开-->
                    <a href="javascript:void(0)" onclick="collect('Y3165623','degree','chi');" class="read detail_icon_collection" ><i class="icon icon_collection" title="收藏"></i>收藏</a>
				<a href="javascript:void(0)" class="read icon_share3 detail_icon_share"><i id="icon_share2q" class="icon icon_share"></i>分享</a>
				<!-- 更多来源 -->
            </div>

            <!--分享弹框-->
<div id="sharez00"
	style="position: absolute;z-index: 200;left: 411px;top: 299px;display: none;">
	<div class="shareto_div_01">
		<div class="shareto_div_02">
			<input type="hidden" class="share_summary"  value="随着科技的进步和人们生活水平的提高，互联网用户大量增加。据报告，中国网民数量将突破8亿。如此大的用户规模，使各类企业趋向于借助于网络广告推广自己的产品和服务。目前网络广告存在着无目标、爆炸式投放的缺点，无法满足企业的推广预期，关于广告精准投放的研究正在深入进行。<br>　　目前网络广告精准化投放主要有三种方式，基于物理位置的分类广告、基于用户行为网页内容的定向广告、基于手机终端的短信广告等。其中定向广告的覆盖面广、精准度高，被各大互联网公司和精准化广告服务商所重视。本文主要研究基于用户行为定向的精准化广告投放，并且进行了设计和实现。论文首先介绍一些背景知识和精准化投放技术的现状，之后介绍系统实现中的相关术语、技术、以及所使用的分类算法，然后将用户的行为分为短期行为和长期行为，通过网络广告追踪窗口对用户浏览的网页进行追踪。根据用户上网行为的分别对用户长期行为和短期行为进行特征分析，并据此进行广告投放。再重点介绍系统设计实现的过程，包括数据库设计以及各个功能模块的详细设计等。论文最后给出了全文的总结，提出系统中存在的不足情况以及下一步的工作重点，总结了本人在硕士研究生期间的工作和成果。<br>　　本系统目前已经在湖南某视频网站投入使用，根据运行数据显示，本系统广告投放精准度有一定的提高，并有效地提高了广告的投放成功率与ECPM，为企业带来了良好的口碑与很高的经济效益。"/>
				<!--注：book资源暂时屏蔽“万方学术圈”，待学术圈加上相应资源可放开-->
				<a href="javascript:void(0);"onclick="getWxShareTemplate(this,'基于Golang的广告投放系统的设计与实现','肖霄','details/detail.do?_type=degree&id=Y3165623','xuezhequan');" class="stitle"><span class="stico stico_xuezhequan">万方学术圈</span><input type="hidden" value="开心网kaixin-kx"></a>
				<a href="javascript:void(0);" onclick="getWxShareTemplate(this,'基于Golang的广告投放系统的设计与实现','肖霄','details/detail.do?_type=degree&id=Y3165623','sina');" class="stitle"><span class="stico stico_tsina">新浪微博</span><input type="hidden" value="新浪微博SINA微博wb-weibo"></a>
				<a href="javascript:void(0);" onclick="getWxShareTemplate(this,'基于Golang的广告投放系统的设计与实现','肖霄','details/detail.do?_type=degree&id=Y3165623','weixin');" class="stitle"><span class="stico stico_weixin">微信</span><input type="hidden" value="搜狐微博tsohu-wb-weibo"></a>
				<a href="javascript:void(0);" onclick="getWxShareTemplate(this,'基于Golang的广告投放系统的设计与实现','肖霄','details/detail.do?_type=degree&id=Y3165623','qqZone');" class="stitle"><span class="stico stico_qq">QQ空间</span><input type="hidden" value="豆瓣db-douban"></a>
				<a href="javascript:void(0);" onclick="getWxShareTemplate(this,'基于Golang的广告投放系统的设计与实现','肖霄','details/detail.do?_type=degree&id=Y3165623','renren');" class="stitle"><span class="stico stico_renren">人人网</span><input type="hidden" value="人人renren-rr"></a>
				<a href="javascript:void(0);" onclick="getWxShareTemplate(this,'基于Golang的广告投放系统的设计与实现','肖霄','details/detail.do?_type=degree&id=Y3165623','tieba');"  class="stitle"><span class="stico stico_baidu">百度贴吧</span><input type="hidden" value="搜狐微博tsohu-wb-weibo"></a>
			<div style="clear:both"></div>
		</div>
	</div>
</div>            <div class="abstract">
            			<textarea rows="" cols="" style="display: none;">随着科技的进步和人们生活水平的提高，互联网用户大量增加。据报告，中国网民数量将突破8亿。如此大的用户规模，使各类企业趋向于借助于网络广告推广自己的产品和服务。目前网络广告存在着无目标、爆炸式投放的缺点，无法满足企业的推广预期，关于广告精准投放的研究正在深入进行。<br>　　目前网络广告精准化投放主要有三种方式，基于物理位置的分类广告、基于用户行为网页内容的定向广告、基于手机终端的短信广告等。其中定向广告的覆盖面广、精准度高，被各大互联网公司和精准化广告服务商所重视。本文主要研究基于用户行为定向的精准化广告投放，并且进行了设计和实现。论文首先介绍一些背景知识和精准化投放技术的现状，之后介绍系统实现中的相关术语、技术、以及所使用的分类算法，然后将用户的行为分为短期行为和长期行为，通过网络广告追踪窗口对用户浏览的网页进行追踪。根据用户上网行为的分别对用户长期行为和短期行为进行特征分析，并据此进行广告投放。再重点介绍系统设计实现的过程，包括数据库设计以及各个功能模块的详细设计等。论文最后给出了全文的总结，提出系统中存在的不足情况以及下一步的工作重点，总结了本人在硕士研究生期间的工作和成果。<br>　　本系统目前已经在湖南某视频网站投入使用，根据运行数据显示，本系统广告投放精准度有一定的提高，并有效地提高了广告的投放成功率与ECPM，为企业带来了良好的口碑与很高的经济效益。</textarea>
            			<div id="see_alldiv" style="overflow:hidden;"><em>摘要</em>： 随着科技的进步和人们生活水平的提高，互联网用户大量增加。据报告，中国网民数量将突破8亿。如此大的用户规模，使各类企业趋向于借助于网络广告推广自己的产品和服务。目前网络广告存在着无目标、爆炸式投放的缺点，无法满足企业的推广预期，关于广告精准投放的研究正在深入进行。<br>　　目前网络广告精准化投放主要有三种方式，基于物理位置的分类广告、基于用户行为网页内容的定...&nbsp;&nbsp;
	            			<a href="#" id="see_all" onclick="see_all()">查看全部&gt;&gt;</a>
	            		</div>
	            		<p id="nosee_all" style="display: none"><a href="#" onclick="nosee_all()">收起<span style="font-size: 17px">∧</span></a></p>
            </div>

           <!--  图书判断 -->
            <!--人物名片 -->
            <!--为保障数据安全，避免我们已处理好的学者与机构的对应关系被别人抓取，故，在文献详情页的题录信息处及相关作者处，鼠标滑过作者姓名，不再显示学者名片。  -->
            <!--<div class="basic" style="display: none">
               <div class="photo"></div>
               <div>
	               <a href="javascript:void(0)" onclick="authorHome('degree','1','','')" class="photo_name"></a>
	               <span></span>
	                 <a href="javascript:void(0)" class="follow" onclick="follow(this,'Y3165623','')"></a>
               </div>
               <div></div>
             </div>-->

             <!--成果详情信息-->
<ul class="info">
					<li><div class="info_left">关键词：</div>
						<div class="info_right info_right_newline">
									<!--<a href="#" onclick="searchResult('degree','关键词:互联网广告')">互联网广告</a>-->
							<a title="知识脉络分析" href="#" onclick="wfAnalysis('degree','互联网广告','key_word')">互联网广告<i class="icon icon_key_word"></i></a>
									<!--<a href="#" onclick="searchResult('degree','关键词:投放系统')">投放系统</a>-->
							<a title="知识脉络分析" href="#" onclick="wfAnalysis('degree','投放系统','key_word')">投放系统<i class="icon icon_key_word"></i></a>
									<!--<a href="#" onclick="searchResult('degree','关键词:软件设计')">软件设计</a>-->
							<a title="知识脉络分析" href="#" onclick="wfAnalysis('degree','软件设计','key_word')">软件设计<i class="icon icon_key_word"></i></a>
									<!--<a href="#" onclick="searchResult('degree','关键词:功能模块')">功能模块</a>-->
							<a title="知识脉络分析" href="#" onclick="wfAnalysis('degree','功能模块','key_word')">功能模块<i class="icon icon_key_word"></i></a>
						</div>
					</li>
					<li>
						<div class="info_left">作者：</div>
						<div class="info_right">
									<a id="card01" href="#" class="info_right_name" onclick="authorHomeWfAnalysis('degree','1','5','肖霄','http://common.wanfangdata.com.cn','degreearticalY3165623','','湖南大学')">肖霄<i class="icon icon_key_word"></i></a>
									<input class="ly"  type="hidden" value="科技创新信息">
	                    			<input class="dw" type="hidden" value="湖南大学">
	                    			<input class="scholar_id" type="hidden" value="" />
	                    			<input class="cardId"  type="hidden" value="card01">
									<!-- <a title="分析" href="#" onclick="wfAnalysis('会议','肖霄','authors_name')"><i class="icon icon_name"></i></a> -->
						</div></li>
						
				<!--  -->					<li><div class="info_left">学位授予单位：</div>
						<div class="info_right info_right_newline" >
								<a title="知识脉络分析" href="javascript:void(0)" onclick="wfAnalysis('degree','湖南大学','unit_name')">湖南大学<i class="icon icon_key_word"></i></a>
						</div>
					</li>
					
					<li><div class="info_left">授予学位：</div>
						<div class="info_right author">硕士</div>
					</li>
					<li><div class="info_left">学科专业：</div>
						<div class="info_right"><a href="#" onclick="searchResult('degree','学科专业:软件工程')">软件工程</a></div>
					</li>
					
					
					
					
					
					
					<li>
					<div class="info_left">导师姓名：</div>
						<div class="info_right">
							<a id="card00" class="info_right_name" href="#" onclick="authorHome('degree','1','5','李丽娟','http://common.wanfangdata.com.cn','Y3165623','','湖南大学')">李丽娟</a>
							<!--<a id="card00" class="info_right_name" href="#" onclick="wfAnalysis('degree','李丽娟'+','+'湖南大学','authors_bootpage')">李丽娟<i class="icon icon_key_word"></i></a>-->
							<a id="card00" class="info_right_name" href="#" onclick="authorHome('degree','1','5','彭一江','http://common.wanfangdata.com.cn','Y3165623','','湖南大学')">彭一江</a>
							<!--<a id="card00" class="info_right_name" href="#" onclick="wfAnalysis('degree','彭一江'+','+'湖南大学','authors_bootpage')">彭一江<i class="icon icon_key_word"></i></a>-->
							
								<input class="ly"  type="hidden" value="">
                    			<input class="dw" type="hidden" value="">
                    			<input class="cardId"  type="hidden" value="card00">
						</div>
					</li>
					<li><div class="info_left">学位年度：</div>
						<div class="info_right author">2016</div>
					</li>
					<li><div class="info_left">语种：</div>
						<div class="info_right author">
						中文
						</div>
					</li>
                		<li>
								<div class="info_left">分类号：</div>
							<div class="info_right author" >
											F713.8
											TP311.52
							</div>
						</li>
					
                
                <li><div class="info_left">在线出版日期：</div><div class="info_right author">
                	2017-04-24<span>（万方平台首次上网日期，不代表论文的发表时间）</span>
                </div></li>
                
</ul>			<!--会议详情信息-->
        </div>
        <div class="left_con_middle" id="left_con_middle_">
        	<div class="title" id="showOrhidde">
            	<a href="javascript:void(0);" id="_ywwl"  onclick="xwwl()">引文网络</a>
				<a href="javascript:void(0);" id="_xgwx"  onclick="xgwx()">相关文献</a>
				<a href="javascript:void(0);" id="_mtzy"  onclick="mtzy()">媒体资源</a>
            </div>
             <div id="content_ywwl">
            <!--引文网络-->
           <!--参考文献-->
            <div class="reference" id="citation_paper" style="display: none;">
            	<div class="subject"><span></span>参考文献
            	<u id="ren_num" style="text-decoration: none;"></u>
            	<a href="javascript:void(0);" onclick="seeCitationMap(this,'Y3165623')">查看参考关系图</a>
            	</div>
            	<div id="CitationMap" style="height:400px;width:920px;display: none">
            	</div>

            	<!--期刊论文-->
                <div class="journal" id="perio_reference">
                </div>
		         <!--分页开始-->
				  	<div id="page" style="padding:25px 0px 30px 0px;border-bottom: 1px dotted #484848;">
					</div>
				<!--分页结束-->
            </div>

            <!--引证文献-->
            <div class="reference" id="reference_paper" style="display: none;">
            	<div class="subject">引证文献
            	<u id="cite_num" style="text-decoration: none;"></u>
				<a href="#" onclick="seeByCitationMap('Y3165623')">查看被引分布图</a>
            	</div>
            	<div id="byCitationMap" style="height:400px;width: 920px;display:none;">

            	</div>

                <div class="journal" id="conf_reference">
                </div>
				<!--分页结束-->
                <!--会议论文-->
		         <!--分页开始-->
			  	<div id="page1" style="padding:25px 0px 30px 0px;border-bottom: 1px dotted #484848;">
				</div>
               </div>

				</br>
          </div>
         <!--  相关文献 -->
          <div id="content_xgwx">
            <!--相似论文-->

            <div class="reference">
            	<div class="subject"><span></span>相关论文<span class="detailsBotWarn">(与本文研究主题相同或者相近的论文)</span></div>
                <div class="journal mt_10" id="sililar_paper_reference">
                	<div style="padding:20px 0 45px 0px;text-align:center"><img src="http://cdn.wanfangdata.com.cn/page/images/gif/loading.gif"></div>
                </div>
            </div>
            <!--同项目论文-->
            <div class="reference">
            	<div class="subject"><span></span>同项目论文<span class="detailsBotWarn">(和本文同属于一个基金项目成果的论文)</span></div>
                <div class="journal mt_10" id="same_project_reference">
                	<div style="padding:20px 0 45px 0px;text-align:center"><img src="http://cdn.wanfangdata.com.cn/page/images/gif/loading.gif"></div>
                </div>
            </div>
          </div>

          <div id="mediaResources">
	          <div class="reference" id="reference_blobs">
	    	 </div>
	    	 <div class="reference ResultCont no_border" id="reference_video">
	    	 	<div style="padding:20px 0 45px 0px;text-align:center"><img src="http://cdn.wanfangdata.com.cn/page/images/gif/loading.gif"></div>
    	 	 </div>
          </div>

        </div>
		<div style="margin: 0 auto;display: none;" class="frame-comment-list" data-snsurl="http://www.wanfangdata.com.cn/sns">
			<iframe src="" width="100%" frameborder="0" scrolling='no' id="frame-comment-list"></iframe>
		</div>
    </div>
    <div class="right">
    	<!--相关主题-->
        <div class="relevant clearfix" id="related_topic">
	       	<div style="padding:20px 0 45px 0px;text-align:center"><img src="http://cdn.wanfangdata.com.cn/page/images/gif/loading.gif"></div>
        </div>
        <!--相关机构-->
        <div class="relevant clearfix" id="unitInfo">
	       	<div style="padding:20px 0 45px 0px;text-align:center"><img src="http://cdn.wanfangdata.com.cn/page/images/gif/loading.gif"></div>
        </div>
        <!--相关学者-->
        <div class="relevant clearfix" id="authorInfo">
	       	<div style="padding:20px 0 45px 0px;text-align:center"><img src="http://cdn.wanfangdata.com.cn/page/images/gif/loading.gif"></div>
        </div>
        <!--我的标签-->
        <input type="hidden" value="" id="Theme_user">
        <div class="My relevant clearfix" style="display: block;">
        	<!--<div class="theme">我的标签<span style="font-size: 12px;color: #999;margin-left: 8px;">最多不超过5个</span></div>-->
			<div class="theme">我的标签<span class="rightWarn" data-c="true"></span>
				<div class="rightWarnText rightWarnItem4"><img src="http://cdn.wanfangdata.com.cn/page/css/images/guide/guide-up.png" alt="">您可以为文献添加知识标签，方便您在书案中进行分类、查找、关联
				</div>
			</div>
			<div id="theme_content">

            </div>
        	<div class="add">
		       <!-- <INPUT id="inputAddTag" style="display:block" placeholder="请输入添加的标签">
		        <i id="iconId" style="display:none" class="icon_add" onclick="add_tag()"></i> -->
		        <!-- <a id="addTag" href="javascript:void(0);" onclick="addTagshow()">添加标签
		        <i class="icon_add"></i></a> -->
				<style>
					.box-input {
						height: 34px;
						line-height: 34px;
						width: 210px;
						position: relative;
						margin: 0 auto;
						padding-left:10px;
					}

					.box-input input {
						height: 32px;
						line-height: 32px;
						width: 203px;
						padding-left: 5px;
						border: 1px solid #cccccc;
						color: #333;

					}
					.box-input input::-ms-clear { display: none; }
					.box-input .box-tips {
						position: absolute;
						top: 0;
						left: 8px;
						max-width:202;
						overflow:hidden;
						height:32px;
					}

					.box-input .box-tips .input-val {
						color: #999;
						visibility: hidden;
					}
					.box-input .box-tips span {
						 padding-right: 0;
					}

						.box-input .box-tips .input-val.on {
							display: inline-block;
							padding-right: 12px;
						}

					.box-input .box-tips .input-tips {
						color: #ccc;
						white-space:nowrap;
						padding-left: 10px;
					}
					.box-input .box-tips .input-tips img {
						vertical-align:middle;
						padding:0 3px;
					}
					/*请求原文弹出框*/
					.request-layer {
						position: fixed;
						top: 0;
						left: 0;
						right: 0;
						bottom: 0;
						z-index: 99;
					}

					.request-layer .wrap-layer {
						width: 100%;
						height: 100%;
						background-color: #000;
						opacity: 0.2;
						filter: Alpha(opacity=20)
					}

					.request-layer iframe {
						position: absolute;
						top: 190px;
						left: 50%;
						margin-left: -320px;
						z-index: 1
					}
				</style>

				<div id="box-input" class="box-input">
					<input type="text" id="inputAddTag" autocomplete="off" />
					<div class="box-tips"><span class="input-val"></span><span class="input-tips">请输入添加的标签</span></div>
				</div>
				<script type="text/javascript">
					var inputRelateTips = (function () {
						var defaultTips = '请输入添加的标签';
						var inputTips = '按回车' + '<img src="http://cdn.wanfangdata.com.cn/page/css/images/icon-huiche.png">' + '添加';
						function initTips(obj) {
							showTips(obj);
							var wrap = obj.parent();
							wrap.find('.input-val').text('').removeClass('on');
							wrap.find('.input-tips').html('请输入添加的标签');
						}
						function changeTips(obj) {
							showTips(obj);
							var wrap = obj.parent();
							wrap.find('.input-val').text(obj.val()).addClass('on');
							wrap.find('.input-tips').html(inputTips);
						}
						function hideTips(obj) {
							var wrap = obj.parent();
							wrap.find('.box-tips').hide();
						}
						function showTips(obj) {
							var wrap = obj.parent();
							wrap.find('.box-tips').show();
						}
						function addLabel(obj) {}
						return {
							init: (function () {
								$('#box-input').click(function () {
									$(this).find('#inputAddTag').focus();
								});
								$('#inputAddTag').focus(function () {
								var $this =$(this);
									 inputAddTagtime  = setInterval(function(){                          ;
									  $this.parent().find('.input-val').val($this.val());
									},100);
								});
								$('#inputAddTag').blur(function () {
									clearInterval(inputAddTagtime);
									inputAddTagtime = null;
								});
								function keyFunction(e,$this) {
									//解决浏览器
									var keyCode = e.keyCode ? e.keyCode : e.which ? e.which : e.charCode;
									if (keyCode == 13) {
										if ($this.val().length > 0) {
										//添加标签的函数在外部执行
										   // addLabel($this);
										}
										//$this.val('');
									   // initTips($this);
									} else {
										if ($this.val().length > 0) {
											changeTips($this);
											if ($this.val().length > 9) {
												hideTips($this);
											}
										} else {
											initTips($this);
										}
									}
								};
								$('#inputAddTag').keypress(function (e) {
									var $this = $(this);
									keyFunction(e,$this);
								});
								$('#inputAddTag').keyup(function (e) {
									var $this = $(this);
									keyFunction(e,$this);
								});
								$('#inputAddTag').keydown(function (e) {
									var $this = $(this);
									keyFunction(e, $this);
								});
							})()
						}
					})();
				</script>
	        </div>
        </div>
    </div>
    <!--  学位目录 -->
    <div class="mask" style="display:none;"></div>
	 <div class="popup cata_pop" style="display:none;">
	 	<div class="popup_close">
	 		<a href="javascript:void(0)" onclick="closeChap()"><i class="icon icon_close"></i></a>
	 	</div>
	 	<div class="pop_catalogue">
			<div class="box_right">
			  <a href="javascript:void(0)" class="news_down"  onclick="upLoadPermissions('63','Y3165623','chi','[WF, CNKI]','基于Golang的广告投放系统的设计与实现','0','degree')"><i></i>下载全文</a>
			  <a href="javascript:void(0)" class="news_full" onclick="fullChap('fullChap')" name="news_full" id="fullChap"><i></i>全屏查看</a>
			  <a href="javascript:void(0)" onclick="javascript:document.getElementById('catalogue_block').scrollTop = 0;" class="news_gototop"><i></i>回到顶部</a>
			</div>
   			<div class="title" id="chap">目录</div>
	   		<div class="catalogue_block" id="catalogue_block">
	     	</div>
		</div>
	</div>
	<!--  学位目录完 -->
</div>
</div>
</body>
<!--异步获取是学位否有目录-->
<script>
    $(function () {
		//获取url中的参数
		function getUrlParam(name) {
			var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
			var r = window.location.search.substr(1).match(reg);  //匹配目标参数
			if (r != null) return unescape(r[2]); return null; //返回参数值
		}

		var me = $('#anxs-8qwe-login');
		var id=getUrlParam('id');
		var articleId = $("#article_id").val();
		if (articleId != null && articleId != "" && articleId != undefined) {
			id = encodeURI(articleId);;
		}
		var urlParamdata = {resourceId: id, type: getUrlParam('_type')}
		var rePath = "http://oss.wanfangdata.com.cn/branch/word?type="+ urlParamdata.type +"&resourceId="+ urlParamdata.resourceId
		//弹出登录框
		function layerAlert(){
			var winHeight = $(window).height();
			if(winHeight<800){
				layer.open({
					offset: '40px',
					type: 1, //page层 1div，2页面
					area: ['50%px'],
					title: '',
					shade: 0.6, //遮罩透明度
					moveType: 1, //拖拽风格，0是默认，1是传统拖动
					shift: 1, //0-6的动画形式，-1不开启
					content: $("#anxs-8qwe-login"),
					success:function(layero){
						$(layero).css({'position':'absolute','top':'40px','left':'50%','marginLeft':'-350px','marginBottom':'40px'});
						$(window).resize(function(){
							var winHeight = $(window).height();
							if(winHeight <800){
								$(layero).css({'position':'absolute','top':'40px','left':'50%','marginLeft':'-350px','marginBottom':'40px'});
							}else{
								$(layero).css({'position':'','top':'','left':'50%','marginLeft':'-350px','marginBottom':''});
							}
						});
					},
					end: function(){

					}
				});
			}else{
				layer.open({
					offset:'auto',
					type: 1, //page层 1div，2页面
					area: ['50%px'],
					scrollbar: true,
					title: '',
					shade: 0.6, //遮罩透明度
					moveType: 1, //拖拽风格，0是默认，1是传统拖动
					shift: 1, //0-6的动画形式，-1不开启
					content: $("#anxs-8qwe-login"),
					success:function(layero){
						$(window).resize(function(){
							var winHeight = $(window).height();
							if(winHeight <800){
								$(layero).css({'position':'absolute','top':'40px','left':'50%','marginLeft':'-350px','marginBottom':'40px'});
							}else{
								$(layero).css({'position':'','top':'','left':'50%','marginLeft':'-350px','marginBottom':''});
							}
						});
					},
					end: function(){

					}
				});
			}
		}
		var accountId = ''
		var isJG = false
		var isGR = false
    	// 2、获取用户登录信息
		 $.ajax({
			// url: $("#loginbyurlDetai").val() + "/getUserState",
			url: "http://login.wanfangdata.com.cn/getUserState",
			dataType: "jsonp",
			jsonp: "callback",
		 	async: false,
			success: function (data) {
				var dataParent = data.context;
				accountId = dataParent.accountIds;
				for (var i=0; i<accountId.length;i++) {
					if (accountId[i].split(".")[0] === 'Person'){
						isGR = true
						continue
					}
					if (accountId[i].split(".")[0] === 'Group'){
						isJG = true
						continue
					}
				}
			},
			 complete:function () {
				 if (isGR && window.sessionStorage.getItem('isPerson') === "true"){
					 window.sessionStorage.setItem("isPerson", "false")
					 window.location.href = "http://oss.wanfangdata.com.cn/branch/word?type="+ urlParamdata.type +"&resourceId="+ urlParamdata.resourceId;
				 }
			 }

		});


		// 1、判断是否显示pdf转word按钮
		if (urlParamdata.type === 'perio' || urlParamdata.type === 'degree') {
			$.ajax({
				async: false,
				data: urlParamdata,
				url: "http://oss.wanfangdata.com.cn/branch/hasWord",
				dataType: "jsonp",
				jsonp: "callback",//传递给请求处理程序或页面的，用以获得jsonp回调函数名的参数名(一般默认为:callback)
				success: function(json){
					if (json) {
						$('.isShowPdfBtn').show()
					}
				}
			});
		}

		$('.isShowPdfBtn span').click(function() {
			if (!accountId.length) {
				window.sessionStorage.setItem("isPerson", "true")
				layerAlert();
				me.attr('src', 'http://my.wanfangdata.com.cn/auth/user/alllogin.do?login_mode=AJAX&service=' + rePath);
			}else {
				if (isJG && !isGR) {
					window.sessionStorage.setItem("isPerson", "true")
					layerAlert();
					me.attr('src', 'http://my.wanfangdata.com.cn/auth/user/login.do?login_mode=AJAX&service=' + rePath);
				}
				if (!isJG && !isGR) {
					window.sessionStorage.setItem("isPerson", "true")
					layerAlert();
					me.attr('src', 'http://my.wanfangdata.com.cn/auth/user/alllogin.do?login_mode=AJAX&service=' + rePath);
				}
				if (isGR){
					window.location.href = "http://oss.wanfangdata.com.cn/branch/word?type="+ urlParamdata.type +"&resourceId="+ urlParamdata.resourceId;
				}

			}
		})


        //有用无用
		var articleId = $('#article_id').val();
		var snsPerio = $('#snsPerio');
		if(snsPerio.length) {
            //是否显示有用没用
            showUserless(articleId);
		}
        //点击有用
        $(document).on('click','.statistics-use',function () {
            if($('.statistics-use').attr('data-flag') && $('.statistics-useless').attr('data-flag')) {
                var _this = $(this);
                var selectState;
                _this.attr('data-flag','');
                if(_this.hasClass('statistics-use-active')) {
                    selectState = 2;//撤销
                }else{
                    selectState = 1;
                }
                var data = {
                    articleId:articleId,
                    selectState:selectState
                }
                useClcikHandle('http://www.wanfangdata.com.cn/sns/third-web/per/article/addArticleUse',_this,data,selectState);
            }
        });

        //点击没用
        $(document).on('click','.statistics-useless',function () {
            if($('.statistics-useless').attr('data-flag') && $('.statistics-use').attr('data-flag')) {
                var _this = $(this);
                var selectState;
                _this.attr('data-flag','');
                if(_this.hasClass('statistics-use-active')) {
                    selectState = 2;
                }else{
                    selectState = 1;
                }
                var data = {
                    articleId:articleId,
                    selectState:selectState
                }
                uselessClcikHandle('http://www.wanfangdata.com.cn/sns/third-web/per/article/addArticleUnUse',_this,data,selectState);
			}

        });

        function showUserless(articleId) {
            var statisticsUseful = $('.statistics-useful');
            var use = statisticsUseful.find('.statistics-use');
            var useLess = statisticsUseful.find('.statistics-useless');
            $.ajax({
                url:'http://www.wanfangdata.com.cn/sns/third-web/per/article/articleUseInfo',
                data:{articleId:articleId},
				dataType:'json',
                success:function (res) {
                    if(res.code == 200) {
                        statisticsUseful.show();
                        var data = res.data;
                        if(data.useState == 'use') {
                            use.addClass('statistics-use-active');
                            useLess.removeClass('statistics-use-active');
                        }else if(data.useState == 'unuse'){
                            useLess.addClass('statistics-use-active');
                            use.removeClass('statistics-use-active');
                        }else{
                            use.removeClass('statistics-use-active');
                            useLess.removeClass('statistics-use-active');
                        }
                        thumbUpFor(data.useNum , use.siblings('.useful-num').find('em'));
                        thumbUpFor(data.unuseNum , useLess.siblings('.useful-num').find('em'));
                    }else{
                        statisticsUseful.hide();
                    }
                }
            })
        }

        function thumbUpFor(num , ele) {
            var num = Number(num);
            ele.attr('data-num',num);
            switch (true) {
                case 1 <= num && num < 1000: ele.text(num);
                    break;
                case 1000 <= num && num < 10000: ele.text((num/1000).toFixed(1) + 'k');
                    break;
                case 10000 <= num:  ele.text((num/10000).toFixed(1) + 'w');
                    break;
                default:
                    ele.text('');
                    break;
            }
        }

        function uselessClcikHandle(url,_this,data,selectState) {
            $.ajax({
                url:url,
                data:data,
                dataType:'json',
                success:function (res) {
                    if(res.code == 400) {
                        getloginurl();
                    }else if(res.code == 200) {
                        var eleUesNum = $('.statistics-use').siblings('.useful-num').find('em');
                        var eleUeslessNum = _this.siblings('.useful-num').find('em');
                        var uselessNum;
                        if(selectState == 2) {
                            _this.removeClass('statistics-use-active');
                            uselessNum = Number(eleUeslessNum.attr('data-num')) - 1;
                        }else{
                            _this.addClass('statistics-use-active');
                            if($('.statistics-use').hasClass('statistics-use-active')) {
                                $('.statistics-use').removeClass('statistics-use-active');
                                thumbUpFor(Number(eleUesNum.attr('data-num')) - 1 , eleUesNum);
                            }
                            uselessNum = Number(eleUeslessNum.attr('data-num')) + 1;
                        }
                        thumbUpFor(uselessNum , eleUeslessNum);
                    }else {
                        layer.msg("操作失败");
                    }
                },
                complete:function () {
                    _this.attr('data-flag',true)
                }

            })
        }

        function useClcikHandle(url,_this,data,selectState) {
           $.ajax({
               url:url,
               data:data,
               dataType:'json',
               success:function (res) {
                   if(res.code == 400) {
                       getloginurl();
                   }else if(res.code == 200) {
                       var eleUesNum = _this.siblings('.useful-num').find('em');
                       var eleUeslessNum = $('.statistics-useless').siblings('.useful-num').find('em');
                       var useNum;
                       if(selectState == 2) {
                           _this.removeClass('statistics-use-active');
                           useNum = Number(eleUesNum.attr('data-num')) -1;
                       }else{
                           _this.addClass('statistics-use-active');
                           if($('.statistics-useless').hasClass('statistics-use-active')) {
                               $('.statistics-useless').removeClass('statistics-use-active');
                               thumbUpFor(Number(eleUeslessNum.attr('data-num') - 1) , eleUeslessNum);
                           }
                           useNum = Number(eleUesNum.attr('data-num')) + 1;
                       }
                       thumbUpFor(useNum , eleUesNum);
                   }else {
                       layer.msg("操作失败");
                   }
               },
               complete:function () {
                   _this.attr('data-flag',true)
               }
           })
        }


        var chapterstype=$("#chapterstype").val();
        var result=$("#chapters").val();
        if(typeof chapterstype == "undefined"||chapterstype!="degree"){
            return;
		}
        if(result!=null&&result!=''&& typeof result != "undefined"){
            result="1:"+result;
            $.post("/search/ajaxGetChapter.do",{"chapters":result},function (data) {
                $.each(data,function(index,item){
                    if(item.status=="1"){
                        $("#chapterindex").show();
                    }
                })
            },"json");
        }
    });
</script>
</html>


<script language="javascript" src="http://cdn.login.wanfangdata.com.cn/Content/js/lang.js"></script>
<style>
body{overflow-x:hidden;}.anxs-8qwe-footRt ul{list-style:none;margin:0;padding:0;}.anxs-8qwe-footer .clear{zoom:1}.anxs-8qwe-footer .clear:after{content:"";display:block;clear:both;visibility:hidden;height:0}.anxs-8qwe-footer{font-family:'Microsoft YaHei';width:100%;background:#f3f3f3;color:#323232;clear:both;font-size:14px;padding-bottom:7px}.anxs-8qwe-footer a:hover{text-decoration:none;}.anxs-8qwe-footer .anxs-8qwe-footer-wrapper{max-width:1200px;width:1200px\9;margin:0 auto;cursor:default;}.anxs-8qwe-footer .anxs-8qwe-footLt{font-size:12px;position:relative;}.anxs-8qwe-footer  .anxs-8qwe-small{padding-top:30px;width:100%;padding-bottom:21px;}.anxs-8qwe-footer  .anxs-8qwe-small .anxs-8qwe-small-lt{float:left;}.anxs-8qwe-footer  .anxs-8qwe-small .anxs-8qwe-small-rt{float:right;}.anxs-8qwe-footer  .anxs-8qwe-small  a{white-space:nowrap;color:#323232;font-size:13px;font-weight:700;text-decoration:none;vertical-align:top;}.anxs-8qwe-footer  .anxs-8qwe-small  .anxs-8qwe-small-lt a{margin-right:72px;}.anxs-8qwe-footer  .anxs-8qwe-small  .anxs-8qwe-small-rt a:first-child+a{margin:0 4px;}.anxs-8qwe-footer .anxs-8qwe-footLt .anxs-8qwe-copy,.anxs-8qwe-footer .anxs-8qwe-footLt .anxs-8qwe-copy_middle{color:#666666;overflow:hidden;float:left;}.anxs-8qwe-footer .anxs-8qwe-footLt .anxs-8qwe-copy a,.anxs-8qwe-footer .anxs-8qwe-footLt .anxs-8qwe-copy_middle a{color:#666666;text-decoration:none;}.anxs-8qwe-footer .anxs-8qwe-footLt .anxs-8qwe-copy .anxs-lt,.anxs-8qwe-footer .anxs-8qwe-footLt .anxs-8qwe-copy_middle .anxs-rt{margin-bottom:10px;}.anxs-8qwe-footer .anxs-8qwe-footLt .anxs-8qwe-copy{width:600px;}.anxs-8qwe-footer .anxs-8qwe-footLt .anxs-8qwe-copy_middle{width:386px;}.anxs-8qwe-footer .anxs-8qwe-footRt{color:#65686f;font-size:14px;float:left;margin-top:-7px;}.anxs-8qwe-footer .anxs-8qwe-footRt .anxs-8qwe-footRt-top{background:0 0;margin-left:-4px;}.anxs-8qwe-footer .anxs-8qwe-footRt .anxs-8qwe-footRt-top .anxs-8qwe-icon{width:32px;height:34px;display:inline-block;background:url(http://cdn.login.wanfangdata.com.cn/Content/src/img/me-global.png) no-repeat;background-position:-256px -75px;vertical-align:middle}.anxs-8qwe-footer .anxs-8qwe-footRt .anxs-8qwe-footRt-contactus li{margin-bottom:2px}.anxs-8qwe-footer .anxs-8qwe-footRt .anxs-8qwe-footRt-contactus li a{cursor:pointer;color:#65686f;text-decoration:none}.anxs-8qwe-footer .anxs-8qwe-footRt .anxs-8qwe-footRt-contactus li i{background:url(http://cdn.login.wanfangdata.com.cn/Content/src/img/icon-contactus.png) no-repeat;display:inline-block;height:20px;width:20px;vertical-align:top;margin:2px 12px 0 0}.anxs-8qwe-footer .anxs-8qwe-footRt .anxs-8qwe-footRt-contactus li .anxs-8qwe-footRt-online{background-position:0 0}.anxs-8qwe-footer .anxs-8qwe-footRt .anxs-8qwe-footRt-contactus li .anxs-8qwe-footRt-tel{background-position:0 -20px;}@media(max-width:1200px){.anxs-8qwe-footer{width:1200px;}}
</style>
<div class="anxs-8qwe-footer">
    <div class="anxs-8qwe-footer-wrapper clear">
        <div class="anxs-8qwe-small clear">
            <div class="anxs-8qwe-small-lt">
                <a target="_blank" href="http://www.wanfangdata.com.cn/link/platformProducts.do">帮助</a>
                <a target="_blank" href="http://www.wanfangdata.com.cn/link/customerService.do">客户服务</a>
                <a target="_blank" href="https://www.wjx.top/jq/23244564.aspx">问卷调查</a>
                <a target="_blank" href="http://www.wanfangdata.com.cn/link/index.do">关于我们</a>
                <a target="_blank" href="http://www.wanfang.com.cn">公司首页</a>
                <a target="_blank" href="http://weibo.com/wanfangdata">平台微博</a>
                <a target="_blank" href="http://www.hotjob.cn/wt/wanfangdata/web/index">加入我们</a>
                <a target="_blank" href="http://www.wanfangdata.com.cn/link/siteMap.do">网站地图</a>
                <a target="_blank" href="http://login.wanfangdata.com.cn/notice/shop">官方店铺</a>
            </div>
            <div class="anxs-8qwe-small-rt">
                <a href="javascript:zh_tran('s');" class="zh_click" id="zh_click_s">简</a>
                <a href="javascript:zh_tran('t');" class="zh_click" id="zh_click_t">繁</a>
                <a target="_blank" href="http://www.wanfangdata.com/">ENG</a>
            </div>
        </div>
        <div class="anxs-8qwe-footLt clear">
            <div class="anxs-8qwe-copy">
                <div class="anxs-lt"><a target="_blank" href="http://ad.wanfangdata.com.cn/images/hlwcb.jpg">网络出版服务许可证：(总)网出证(京)字096号</a></div>
                <div class="anxs-lt"><a target="_blank" href="http://ad.wanfangdata.com.cn/images/zhengshu_2.jpg">互联网药品信息服务资格证书号：(京)-经营性-2016-0015</a></div>
                <div class="anxs-lt">万方数据知识服务平台--国家科技支撑计划资助项目（编号：2006BAH03B01）</div>
                <div class="anxs-lt"><a target="_blank" href="http://ad.wanfangdata.com.cn/images/wfdatazs.jpg">万方数据学术资源发现服务系统[简称：万方智搜]V1.0  证书号：软著登字第2255655号</a></div>
            </div>
            <div class="anxs-8qwe-copy_middle">
                <div class="anxs-rt"><a target="_blank" href="http://ad.wanfangdata.com.cn/images/icp.jpg">京ICP证：010071</a></div>
                <div class="anxs-rt"><a target="_blank" href="http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=11010802020237">京公网安备11010802020237号</a></div>
                <div class="anxs-rt"><a target="_blank" href="http://ad.wanfangdata.com.cn/images/stjmxkz2.jpg">信息网络传播视听节目许可证 许可证号：0108284</a></div>
                <div class="anxs-rt">&copy;北京万方数据股份有限公司 万方数据电子出版社</div>
            </div>
            <div class="anxs-8qwe-footRt">
                <div class="anxs-8qwe-footRt-top">
                    <i class="anxs-8qwe-icon"></i>
                    <span>联系客服</span>
                </div>
                <ul class="anxs-8qwe-footRt-contactus">
                    <li><i class="anxs-8qwe-footRt-online"></i><a class="anxs-8qwe-onlineService">在线客服</a></li>
                    <li><i class="anxs-8qwe-footRt-tel"></i>4000115888</li>
                    <li><a href="mailto:service@wanfangdata.com.cn">service@wanfangdata.com.cn</a></li>
                </ul>
            </div>
        </div>
    </div>
</div>
<script>
NTKF_PARAM = {
    siteid: 'kf_9358', // 企业ID，为固定值，必填
    settingid: 'kf_9358_1469523642099', // 接待组ID，为固定值，必填
    uid: "", // 用户ID，未登录可以为空，但不能给null，uid赋予的值显示到小能客户端上
    uname: "", // 用户名，未登录可以为空，但不能给null，uname赋予的值显示到小能客户端上
    isvip: '0', // 是否为vip用户，0代表非会员，1代表会员，取值显示到小能客户端上
    userlevel: '1', // 网站自定义会员级别，0-N，可根据选择判断，取值显示到小能客户端上
    erpparam: 'abc' // erpparam为erp功能的扩展字段，可选，购买erp功能后用于erp功能集成
};
 $('.anxs-8qwe-onlineService').click(function(){
    onlineService();
});
 function onlineService(){
    $.ajax({
        url:'http://login.wanfangdata.com.cn/user/getOnlineUser',
        dataType: "jsonp",
        jsonp: "callback",
        success: function (data) {
            window.NTKF_PARAM.uid = data.Id;
            window.NTKF_PARAM.uname = data.Phone;
            var protocol = "https:" == location.protocol ? "https://" : "http://";
            $.get(protocol + 'dl.ntalker.com/js/xn6/ntkfstat.js?siteid=kf_9358', function () {
                NTKF.im_openInPageChat('kf_9358_1469523642099');
            }, 'script');
        }
    });
 }
</script>
<!--建议-->
<style>
    .clear{zoom: 1}
    .clear:after{content: "";display: block;clear: both;visibility: hidden;height: 0}
    .anxs-left-bom{left:auto;width: 60px;background: #ffffff;color: #fff;font-size: 12px;position: fixed;bottom: 105px;right: 0;font-family: 'Microsoft YaHei';z-index: 19;box-shadow: 0 0 10px #CCC;border-radius: 0;}
    .anxs-left-bom-list{text-align: center;border-bottom: 1px solid #E9E9E9;height: 60px;border-top:0;cursor: pointer;position: relative;}
    .anxs-left-bom-list a,.announcementTamll_a{color: #fff;text-decoration: none;padding: 16px 0 15px 0;display: inline-block;margin-left: 0;height: 29px;width: 60px;}
    .announcementTamll_a{position: relative;}
    .announcementTamll_a>a{padding: 0;}
    .anxs-left-bom-list a em,.announcementTamll_a em{font-style:normal;color: #fff;font-weight: bold;-webkit-transition: left .3s ease-in-out .1s;-moz-transition: left .3s ease-in-out .1s;transition: left .3s ease-in-out .1s;display: none;line-height:28px;font-style:normal;}
    .anxs-left-bom-list a i,.announcementTamll_a i{background: url(http://cdn.login.wanfangdata.com.cn/Content/src/img/sideBox.png) no-repeat center;display: inline-block;position: relative;top: 0;left:2px;}
    .anxs-left-bom-list a .anxs-bom-icon-circle{width: 28px;height: 29px;background-position: 0 -3px;margin-left:auto;}
    .anxs-left-bom-list a .anxs-bom-icon-subscribe{width: 22px;height: 22px;background-position: 0 -46px;margin-left:auto;}
    .anxs-left-bom-list a .anxs-bom-icon-collection{width: 26px;height: 26px;background-position: 0 -82px;}
    .anxs-left-bom-list a .anxs-bom-icon-service{width: 23px;height: 27px;background-position: 0 -130px;}
    .anxs-left-bom-list a .anxs-bom-icon-top{width: 26px;height: 31px;background-position: 0 -171px;}
    .anxs-left-bom-list:hover{background:#417DC9;}
    .anxs-left-bom-list2{text-align: center;border-bottom: 1px solid #E9E9E9;height: 60px;border-top:0;cursor: pointer;width:60px;position: relative;}
    .anxs-left-bom-list2 a{color: #fff;text-decoration: none;padding: 16px 0 15px 0;display: inline-block;margin-left: 0;height: 29px;width: 60px;}
    .anxs-left-bom-list2 a em{font-style:normal;color: #fff;font-weight: bold;-webkit-transition: left .3s ease-in-out .1s;-moz-transition: left .3s ease-in-out .1s;transition: left .3s ease-in-out .1s;display: none;line-height:28px;font-style:normal;}
    .anxs-left-bom-list2 a i{background: url(http://cdn.login.wanfangdata.com.cn/Content/src/img/sideBox.png) no-repeat center;display: inline-block;position: relative;top: 0;left:2px;}
    .anxs-left-bom-list2 a .anxs-bom-icon-circle{width: 28px;height: 29px;background-position: 0 -3px;margin-left:auto;}
    .anxs-left-bom-list2 a .anxs-bom-icon-subscribe{width: 22px;height: 22px;background-position: 0 -46px;margin-left:auto;}
    .anxs-left-bom-list2 a .anxs-bom-icon-collection{width: 26px;height: 26px;background-position: 0 -82px;}
    .anxs-left-bom-list2 a .anxs-bom-icon-service{width: 23px;height: 27px;background-position: 0 -130px;}
    .anxs-left-bom-list2 a .anxs-bom-icon-top{width: 26px;height: 31px;background-position: 0 -171px;}
    .anxs-left-bom-list .announcementTamll_a i.anxs-bom-icon-tamll{display:inline-block;width: 28px;height: 27px;background: url("http://cdn.login.wanfangdata.com.cn/Content/images/tamllActive_iconC.png") center center no-repeat;}
    .announcementTamll{position: relative;}
    .announcementTamll:hover .announcementTamllText{color: rgb(255, 255, 255);}
    .announcementTamll .announcementTamllText{position: absolute;width: 60px;height: 60px;top: 0px;left: 0px;line-height: 60px; color: rgb(65, 125, 201);display:inline-block!important;}
    .announcementTamll-info{display: none;position: absolute;right: 70px;top:0;width:209px;height: 225px;background: url("http://cdn.login.wanfangdata.com.cn/Content/images/announcement_bg.png") center center no-repeat;}
    .announcementTamll-info .tamllHide{display: inline-block; position: absolute; height: 90%; width: 14px;right: -10px;top: 10%;}
    .announcementTamll-info  .tamll-info-title{font-size: 16px;color:#fe0000;position: relative;line-height: 36px;}
    .announcementTamll-info  .tamll-info-close{display: inline-block;background: url("http://cdn.login.wanfangdata.com.cn/Content/images/tamllActive_close.png") center center no-repeat;width:12px;height: 12px;position: absolute;right: 14px;  top: 28px;}
    .announcementTamll-info  .tamll-info-content{padding:24px 10px 0 10px;text-align:left;font-size:12px;}
    .announcementTamll-info  .tamll-info-contentA{ text-indent:2em;line-height: 20px;}
    .announcementTamll-info .tamll-info-contentB span{color:#fff;width:168px;padding:0;line-height: 27px;height: 27px;display: inline-block;cursor: auto;}
    .announcementTamll-info  .tamll-info-contentC{text-indent: 2em;line-height: 20px;}
    .announcementTamll-info  .tamll-info-contentC a{color:#f6ff00;text-decoration: underline;width:90%;padding:0;text-align: right;}

    .anxs-left-bom-list2:hover{background:#417DC9;}
    .anxs-left-bom-list2:hover .phone-p{color:#fff;}
    .anxs-left-bom .description2{display:none;position:absolute;right:65px;width:60px;height:62px;}
    .anxs-left-bom .description2 .message{position:absolute;right:8px;height:200px;width:240px;background-color: #FFF;color:#417DC9;font-size:14px;line-height:35px;border-radius: 5px;padding:10px;z-index:100;padding-bottom:12px;}
    .anxs-left-bom .description2 .arrow{position:absolute;top:12px;right:0;width:0;height:0;border-left:8px solid #FFF;border-top:5px solid transparent;border-bottom:5px solid transparent;}
    .code-box{overflow:auto;}
    .border-line{border-bottom:1px solid #ddd;}
    .code-box .code-img{float:left;width:80px;height:80px;margin:8px;}
    .code-box .code-info{float:left;width:140px;margin-top:5px;line-height:25px;text-align:left;}
    .code-box .code-info span{height:25px;line-height:25px;color:#333;font-weight:bold; }
    .code-box .code-info .code-span{color:#ff6c00;    font-weight: normal;}
    .code-box .code-img2{width:20px;height:22px;margin-right:5px;position:relative;top:5px;}
    .code-box .code-info .code-span2{color:#417dc9;display:block;text-align:left;height:25px;line-height:22px;    font-weight: normal;}
    .description2 .code-box a{width:240px;height:100px;display:block;padding-top:6px;}
    #anxs-top{display:none;}
    .anxs-left-bom .description{display:none;position:absolute;right:75px;width:180px;height:35px;}
    .anxs-left-bom .description.collection{top:15px;}
    .anxs-left-bom .description .message{position:absolute;right:8px;height:35px;background-color: #FFF;color:#417DC9;font-size:14px;line-height:35px;border-radius: 5px;padding:0 10px;z-index:100;}
    .anxs-left-bom .suggestionForm{position:absolute;right:80px;height:370px;bottom: -43px;}
    .anxs-left-bom .description .arrow{position:absolute;top:12px;right:0;width:0;height:0;border-left:8px solid #FFF;border-top:5px solid transparent;border-bottom:5px solid transparent;}
    .anxs-left-bom .phone-box{top:0px;right:60px;padding-right: 5px;}
    .anxs-left-bom .phone a p{ color: #417DC9;height: 25px; line-height: 25px; font-size: 14px;font-weight:bold;}
    .anxs-left-bom .circle{top:15px;}
    .anxs-left-bom .subscribe{top:15px;}
    .anxs-left-bom em{font-style:normal;}
    .new-message-tips{width:30px;height:17px;background:url(http://cdn.login.wanfangdata.com.cn/Content/src/img/new-message.png) no-repeat;position:absolute;top:0px;right:1px;display:none;}
    .new-collection-tips{width:22px;height:14px;background:url(http://cdn.login.wanfangdata.com.cn/Content/src/img/collect-numbers.png) no-repeat;position:absolute;top:0px;right:1px;display:none;line-height: 14px;text-align:center;}
    .app-setcards-box{ position: fixed; right: 0; bottom: 480px;width:60px;height: 44px;cursor: pointer;}
    .app-setcards{display: none;}
    .app-setcards-content{width:444px;height: 454px;background: url(http://cdn.login.wanfangdata.com.cn/Content/images/appbg.png) center center no-repeat;position: relative;}
    .app-setcards-content .app-setcards-close{display: inline-block;width:50px;height: 50px;position:absolute;right: 42px;top:9px;border-radius: 50%;cursor: pointer; -webkit-border-radius:50%; }
    .app-setcards-bg .layui-layer-content{background: none;}
    .layui-layer.app-setcards-bg{background: none;box-shadow: none;}
    .anxs-left-bom-list.topic .topictext{color: #417DC9;height: 25px;line-height: 25px;font-weight: bold;}
    .anxs-left-bom-list.topic:hover .topictext{color:#fff; }
    @media screen and (max-width:1200px){
           .anxs-left-bom {width: 50px;bottom:161px;}
           .anxs-left-bom-list{height: 50px;}
           .anxs-left-bom-list a,.announcementTamll_a{padding: 11px 0 15px 0;height: 24px;width: 50px;}
           .anxs-left-bom-list a em,.announcementTamll_a em{line-height:26px;}
           .anxs-left-bom-list a .anxs-bom-icon-top{width: 26px;height: 31px;background-position: 0 -171px;}
           .anxs-left-bom-list:hover{color:#fff;}
           .anxs-left-bom .description{right:65px;}
           .anxs-left-bom .description.collection{top:9px;}
           .anxs-left-bom .circle{top:9px;}
           .anxs-left-bom .subscribe{top:9px;}
           .app-setcards-box{bottom: 485px;}
       }
</style>

<!--app集卡活动-->
<div class="app-setcards">
    <div class="app-setcards-box" id="app-setcards-box">
        <img src="http://cdn.login.wanfangdata.com.cn/Content/images/appsamll.png"/>
    </div>
</div>
<div class="anxs-left-bom">
    <div class="anxs-left-bom-list announcementTamll">
        <div class="announcementTamll_a">
            <a href="http://login.wanfangdata.com.cn/notice/shop" target="_blank">
                <em class="announcementTamllText">店铺公告</em>
            </a>
        </div>
        <div class="announcementTamll-info">
            <span class="tamllHide"></span>
            <div class="tamll-info-title">
                公&nbsp;&nbsp;&nbsp;告
                <span class="tamll-info-close">  </span>
            </div>
            <div class="tamll-info-content">
                <p class="tamll-info-contentA">北京万方数据股份有限公司在天猫、京东开具唯一官方授权的直营店铺:</p>
                <p class="tamll-info-contentB">1、<span>天猫--万方数据教育专营店</span></p>
                <p class="tamll-info-contentB">2、<span>京东--万方数据官方旗舰店</span></p>
                <p class="tamll-info-contentC">敬请广大用户关注、支持!<a href="http://login.wanfangdata.com.cn/notice/shop" target="_blank">查看详情</a></p>
            </div>
        </div>
    </div>
    <div class="anxs-left-bom-list2 phone">
        <a href="javascript:">
            <p class="phone-p">手机版</p>
        </a>
        <div class="description2 phone-box">
            <div class="message">
                <div class="code-box border-line">
                    <img src="http://cdn.login.wanfangdata.com.cn/Content/src/img/code-img_03.png" class="code-img">
                    <p class="code-info">
                        <span>万方数据知识服务平台</span>
                        <span class="code-span">扫码关注微信公众号</span>
                    </p>
                </div>
                <div  class="code-box">
                    <a href="http://www.wanfangdata.com.cn/app/app.html" target="_blank">
                        <img src="http://cdn.login.wanfangdata.com.cn/Content/src/img/code-img_07.png" class="code-img">
                        <p class="code-info">
                            <span>万方数据APP</span>
                            <span class="code-span2"> <img src="http://cdn.login.wanfangdata.com.cn/Content/src/img/code-img_09.png" class="code-img2"> Android 版</span>
                            <span class="code-span2"> <img src="http://cdn.login.wanfangdata.com.cn/Content/src/img/code-img_13.png" class="code-img2"> Ios 版</span>
                        </p>
                    </a>
                </div>
            </div>
            <div class="arrow"></div>
        </div>
    </div>
    <div class="anxs-left-bom-list topic">
        <a href="http://topic.wanfangdata.com.cn/index.do" target="_blank">
            <p class="topictext">万方选题</p>
        </a>
        <div class="new-message-tips" style="display: block;"></div>
    </div>
    <div class="anxs-left-bom-list academic">
        <a href="javascript:void(0)" target="_blank">
            <i class="anxs-bom-icon-circle"></i>
            <em>学术圈</em>
        </a>
        <div class="description circle">
            <div class="message">实名学术社交</div>
            <div class="arrow"></div>
        </div>
    </div>

    <div class="anxs-left-bom-list sub">
        <a href="javascript:void(0)" target="_blank">
            <i class="anxs-bom-icon-subscribe"></i>
            <em>订阅</em>
        </a>
        <div class="description subscribe">
            <div class="message">个性化订阅推荐</div>
            <div class="arrow"></div>
        </div>
        <div class="new-message-tips"></div>
    </div>

    <div class="anxs-left-bom-list collection">
        <div class="new-collection-tips"><span></span></div>
        <a href="javascript:void(0)" target="_blank">
            <i class="anxs-bom-icon-collection"></i>
            <em>收藏</em>
        </a>
        <div class="description collection">
            <div class="message">快速查看收藏过的文献</div>
            <div class="arrow"></div>
        </div>
    </div>

    <div class="anxs-left-bom-list" id="anxs-btn-suggestionNew">
        <a href="javascript:void(0)">
            <i class="anxs-bom-icon-service"></i>
            <em style="line-height: 14px;"><span>客服</span><br/>
                <span>服务</span></em>
        </a>
    </div>
     <div class="suggestionForm" id="suggestionForm" style="display: none"></div>
    <div class="anxs-left-bom-list" id="anxs-top" style="border-bottom:0">
        <a href="javascript:void(0);">
            <i class="anxs-bom-icon-top"></i>
            <em style="line-height:14px;"><span>回到</span><br>
                <span>顶部</span></em>
        </a>
    </div>
</div>

<script>
 //cookie对象
      var CookieAppCards = {
         set : function (name, value, time, domain, path, secure) {
             var cookieText = "";
             cookieText += encodeURIComponent(name) + "=" + encodeURIComponent(value);
             if(time) {
                 var date=new Date();
                 var expiresDays=time;
                 date.setTime(date.getTime()+expiresDays*24*3600*1000);
                 cookieText += "; expires=" + date.toGMTString();
             }
             if (path) {
                 cookieText += "; path=" + path;
             }
             if (domain) {
                 cookieText += "; domain=" + domain;
             }
             if (secure) {
                 cookieText += "; secure";
             }
             document.cookie = cookieText;
         },
         get : function (name) {
             var cookieName = encodeURIComponent(name) + "=",
                     cookieStart = document.cookie.indexOf(cookieName),
                     cookieValue = "";
             if (cookieStart > -1) {
                 var cookieEnd = document.cookie.indexOf (";", cookieStart);
                 if (cookieEnd == -1) {
                     cookieEnd = document.cookie.length;
                 }
                 cookieValue = decodeURIComponent(document.cookie.substring(cookieStart + cookieName.length, cookieEnd));
             }
             return cookieValue;
         },
         unset : function (name, domain, path, secure) {
             this.set(name, "", Date(0), domain, path, secure);
         }
     };
    //announcementTamll-info的显示时间
      var announcementTamlltime = 0;

    //侧边栏列表鼠标悬浮效果
    $(".anxs-left-bom-list2").hover(function () {
        $(".description2").show();
    }, function () {
        $(".description2").hide();
    });
     //电商广告
     var isIndex = $.trim($('#indexScanCodeBind').val());
     var getCookieInfoShow =  CookieAppCards.get("infoShow");
    $(".anxs-left-bom-list").hover(function () {
                $(this).find("em").css("display", "inline-block");
                $(this).find("i").css("display", "none");
                $(this).children(".description").css("display","block");
            },
            function () {
                $(this).find("em").css("display", "none");
                $(this).find("i").css("display", "inline-block");
                $(this).children(".description").css("display","none");
            }
    );

    //app集成卡
        var getCookieAppsetcards =  CookieAppCards.get("appsetcards");
        if(isIndex) {
            jQuery.support.cors=true;
            $.ajax({
                url:"http://login.wanfangdata.com.cn/getServerTimeJSONP",
                dataType:"jsonp",
                jsonp:"callback",
                success:function (data) {
                    var startDate = Date.parse("2018/07/09 12:00");
                    var endDate = Date.parse("2018/07/18 12:00");
                    if(startDate <= data && endDate > data){
                        $(".app-setcards").show();
                        if(!getCookieAppsetcards) {
                            appsetcardLayer();
                            CookieAppCards.set("appsetcards", "appsetcardsActive",8);
                        }
                        $("#app-setcards-box").on("click",function () {
                            appsetcardLayer();
                        });
                    }else{
                        $(".app-setcards").remove();
                    }
                }
            });
        }else{
            $(".app-setcards").remove();
        }
         function appsetcardLayer() {
             layer.open({
                 type: 1,
                 title: "",
                 skin: 'app-setcards-bg', //样式类名
                 closeBtn: 0, //不显示关闭按钮
                 anim:2,
                 time: 20000,
                 area: ['444px', '454px'],
                 shadeClose: false, //开启遮罩关闭
                 content: "<div class='app-setcards-content'> <span class='app-setcards-close'></span> </div>",
                 success: function(layero,index){
                     $(".app-setcards-close").on("click",function () {
                         layer.close(index);
                     })
                 }
             });
         };

    /*点击客服服务*/
     $(function(){
         $.ajax({
             url:"http://my.wanfangdata.com.cn/user/suggestion/suggestionForm",
             success:function(result){
                 $('#suggestionForm').html(result);
             },
             error:function(data){
             }
         });
     });
     $('#anxs-btn-suggestionNew').click(function () {
         $('#suggestionForm').toggle();
     });
    function GetPageScroll() {
        var x, y; if (window.pageYOffset) {    // all except IE
            y = window.pageYOffset;
            x = window.pageXOffset;
        } else if (document.documentElement && document.documentElement.scrollTop) {    // IE 6 Strict
            y = document.documentElement.scrollTop;
            x = document.documentElement.scrollLeft;
        } else if (document.body) {    // all other IE
            y = document.body.scrollTop;
            x = document.body.scrollLeft;
        }
        /*return {X:x, Y:y};*/
        return y;
    }
    setInterval(function () {
        var scroTop = GetPageScroll();
        function mop(str) {
            return eval(str);
        }
        if ( scroTop > 1400) {
            $('#anxs-top').show();
        }else if(scroTop <1400){
            $('#anxs-top').hide();
        }
    }, 40);
    /*点击返回顶部*/
    $("#anxs-top").click(function(){
        $('body,html').animate({scrollTop:0},500);
    })
</script>

<script>
    $(function() {
        var param_academic,
                param_subscribe,
                param_collection;
        param_academic = "http://social.wanfangdata.com.cn/index/toIndex.do";
        param_subscribe = "http://work.wanfangdata.com.cn/toWorkbench/visitor.do";
        param_collection = "http://work.wanfangdata.com.cn/toWorkbench/visitor.do";
        $(".academic").find("a").attr("href", param_academic);
        $(".sub").find("a").attr("href", param_subscribe);
        $(".anxs-left-bom-list.collection").find("a").attr("href", param_collection);
        $.ajax(({
            //获取用户登录状态
            url: "http://login.wanfangdata.com.cn/getUserState",
            dataType: "jsonp",
            jsonp: "callback",
            success:function(accountsData) {
                var dataParent = accountsData.context;
                var accountId =dataParent.accountIds;
                //如果登录，获取用户是否有新的订阅、收藏
                if(accountId.length>0){
                    var arr = accountId[0].split('.');
                    //是否有新的订阅
                    $.ajax({
                        url:"http://work.wanfangdata.com.cn/subscribe/hasNew.do",
                        data:{"userId":arr[1]},
                        dataType:"jsonp",
                        jsonp:"callback",
                        success:function(data){
                            var num=data.result;
                            if(num>0) {
                                $(".new-message-tips").css("display","block");
                            }
                            else{
                                $(".new-message-tips").css("display","none");
                            }


                        }
                    })
                    $.ajax({
                        //是否有新的收藏，超过99显示99+，数量为0则不显示
                        url:"http://work.wanfangdata.com.cn/index/getNoReadCollection.do",
                        data:{"userId":arr[1]},
                        dataType:"jsonp",
                        jsonp:"callback",
                        success:function(data){
                            showSubscibeIcon();
                            if(data.key>0) {
                                var num;
                                $(".new-collection-tips").css("display","block");
                                if(data.key>99){
                                    num='99+';
                                }
                                else {
                                    num=data.key;
                                }
                                $(".new-collection-tips").find("span").text(num);
                            }
                            else{
                                $(".new-collection-tips").css("display","none");
                            }

                        }
                    })

                    param_academic="http://social.wanfangdata.com.cn/index/toTourIndex.do";
                    param_subscribe="http://work.wanfangdata.com.cn/subscribe/index.do";
                    param_collection="http://work.wanfangdata.com.cn/index/index.do?classifyId=1&dustbinState=1";
                    $(".academic").find("a").attr("href", param_academic);
                    $(".sub").find("a").attr("href", param_subscribe);
                    $(".anxs-left-bom-list.collection").find("a").attr("href", param_collection);
                }


            },
            error:function() {
                param_academic="http://social.wanfangdata.com.cn/index/toIndex.do";
                param_subscribe="http://work.wanfangdata.com.cn/toWorkbench/visitor.do";
                param_collection= "http://work.wanfangdata.com.cn/toWorkbench/visitor.do";
                $(".academic").find("a").attr("href",param_academic);
                $(".sub").find("a").attr("href",param_subscribe);
                $(".anxs-left-bom-list.collection").find("a").attr("href",param_collection);
            }
        }))
    });
    function showSubscibeIcon(){
        var subscibeCookieName = "noReadCollection";
        var subscibeCookieValue;
        var collectCookieName = "hasNewToken";
        var collectCookieValue;
        if (document.cookie && document.cookie != ''){
            var cookies = document.cookie.split(';');
            for(var i=0;i<cookies.length;i++){
                var cookie = cookies[i];
                if (cookie.substring(0, subscibeCookieName.length + 2).trim() == subscibeCookieName.trim() + "="){
                    subscibeCookieValue = cookie.substring(subscibeCookieName.length + 2, cookie.length);
                    if(subscibeCookieValue>0){
                        var count;
                        $(".new-collection-tips").css("display", "block");
                        if(subscibeCookieValue>99){
                            count = "99+";
                            $(".new-collection-tips span").text("99+");
                        }else{
                            count = subscibeCookieValue
                        }
                        $(".new-collection-tips span").text(count);
                    }else {
                        $(".new-collection-tips").css("display", "none");
                    }
                }
            }
        }
        if (document.cookie && document.cookie != ''){
            var cookies = document.cookie.split(';');
            var i=0;

            for( i=0;i<cookies.length;i++){
                var cookie = cookies[i];
                if (cookie.substring(0, collectCookieName.length + 1).trim() == collectCookieName.trim()){
                    $(".new-message-tips").css("display", "block");
                    break;
                }
            }
            if(i == cookies.length){
                $(".new-message-tips").css("display", "none");
            }
        }
        setTimeout('showSubscibeIcon()',1000);
    }

</script>
<style>
	.anxs-new-online-tips{
		background:url("http://cdn.wanfangdata.com.cn/page/images/index-tipsBG.png");
		opacity: 1;
	}
</style>
    <!-- 分享html -->
<div id ='jiathis_weixin_modal_test'></div>
<iframe id="login" style="display: none;margin-left: -10px;margin-top: -100px;" src="" height="750px;" width="710px;" scrolling="no" frameborder="0" ></iframe>
  </body>
</html>
