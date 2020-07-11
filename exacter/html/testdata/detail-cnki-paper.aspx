<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
  <head><base __dr2amgenerated="1" href="http://kns.cnki.net/KCMS/detail/detail.aspx?dbcode=CMFD&dbname=CMFDTEMP&filename=1020008317.nh&v=Mjk4MzZQSVI4ZVgxTHV4WVM3RGgxVDNxVHJXTTFGckNVUjdxZlkrWm9GQ3ZoVWJ2SlZGMjVIck80RnRMTnFKRWI=" __dr2am="latest">
				<script __cpp="1">
					window.__Cpn = window.__Cpn ? window.__Cpn : function() {this.mode = 'service_worker_reduce';this.origin = 'https://test.dr2am.cn' };
				</script>
				<script src="https://test.dr2am.cn/public/js/dr2am_product.js?c61cd4ff" ></script>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>开放式实验管理系统的设计与实现 - 中国知网
        </title>
    <link rel="stylesheet" type="text/css" href="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/resource/gb/css_min/Global.min.css?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1">
    <link rel="stylesheet" type="text/css" href="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/resource/gb/css_min/knetdetail.min.css?v=FBC16D09D6F9935E1123121&amp;__dp=http" __cpp="1" __dp="1">
    <link rel="stylesheet" type="text/css" href="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/resource/gb/css/ecplogin.min.css?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1">
    <link rel="stylesheet" type="text/css" href="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/resource/gb/css_min/picModule.min.css?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/kns/_/kcms/detail/js/getLink.aspx?__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/resource/gb/js/min/rs.min.js?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/jquery-1.4.2.min.js?__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/jquery.PrintArea.min.js?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/Common.min.js?v=FBC16D09D6F9935E0129&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/highcharts.min.js?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/json2.min.js?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/raphael.amd.js?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/CnkiFlashEmbed.min.js?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/jquery.mousewheel.min.js?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/NiceDCenter.min.js?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/Kajax.min.js?v=21FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/CatalogFun.min.js?v=FBC16D09D6F9935E102712&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/detail.min.js?v=FBC16D09D6F9935E12&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/recommendpic.min.js?v=FBC16D09D6F9935E&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/kns/script/jQuery-1.11.3.min.js?__dp=http" __cpp="1" __dp="1"></script><script src="https://test.dr2am.cn/--/net/cnki/a/_/a/quote/ad.js?__dp=http" __cpp="1" __dp="1"></script></head>
  <body><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/WideScreen.min.js?__dp=http" __cpp="1" __dp="1"></script><input id="loginuserid" type="hidden" value=""><input id="listv" type="hidden" value=""><input id="paramdbcode" type="hidden" value="CMFD"><input id="paramdbname" type="hidden" value="CMFDTEMP"><input id="paramfilename" type="hidden" value="1020008317.nh"><input id="paramcitingtimes" type="hidden" value="0"><input id="deliveryType" type="hidden" value=""><input id="deliveryUid" type="hidden" value=""><input id="deliveryUname" type="hidden" value=""><input id="deliveryCoutent" type="hidden" value="#"><input id="deliveryTable" type="hidden" value=""><input id="SingePointShow" type="hidden" value=""><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/kns/script/jQuery-1.11.3.min.js?__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/login/_/TopLogin/api/loginapi/get?type=top&amp;localCSS=&amp;returnurl=http%3a%2f%2fkns.cnki.net%2fKCMS%2fdetail%2fdetail.aspx%3fdbcode%3dCMFD%26dbname%3dCMFDTEMP%26filename%3d1020008317.nh%26v%3dMjk4MzZQSVI4ZVgxTHV4WVM3RGgxVDNxVHJXTTFGckNVUjdxZlkrWm9GQ3ZoVWJ2SlZGMjVIck80RnRMTnFKRWI%3d&amp;__dp=http" __cpp="1" __dp="1"></script><div id="headerBox"></div>
    <div class="line-box">
      <h1><span id="catalog_Ptitle"></span><span class="kcmslogo"></span></h1>
    </div><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/CnkiFlashEmbed.js?v=FBC16D09D6F9123&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/jquery.mousewheel.js?v=FBC16D09D6F99123&amp;__dp=http" __cpp="1" __dp="1"></script><div class="wxwrap">
      <div class="searchtop"><span class="showPlaceholder"><input id="searchinput" name="searchinput" type="text" onkeypress="setPlaceholder(this);" onpaste="return setPlaceholder(this);" onkeyup="setPlaceholder(this);"><label for="searchinput" class="placeholder" id="searchtext">请输入搜索内容</label></span><input id="searchbtn" type="button" value="检索" onclick="submitSearch()"></div>
    </div><script>      
      $("#searchinput").keypress(function (event) {
      if (event.keyCode == 13)
      submitSearch();
      });
    </script><script>$("#catalog_Ptitle").html("博硕论文");</script><div id="mainArea" class="wxwrap">
      <div class="wxside">
        <div id="CatalogSide" class="wxside kcmsCatalog">
          <div class="clHd">知识节点</div>
          <div class="clBd">
            <dl>
              <dt id="lcatalog_Ptitle" onclick="TurnToTitle(&#39;catalog_Ptitle&#39;)"><a><i></i>基本信息
            </a></dt>
            </dl>
            <dl>
              <dt id="lcatalog_ABSTRACT" onclick="TurnToTitle(&#39;catalog_ABSTRACT&#39;)"><a><i></i>摘要
              </a></dt>
            </dl>
            <dl>
              <dt id="lcatalog_KEYWORD" onclick="TurnToTitle(&#39;catalog_KEYWORD&#39;)"><a><i></i>关键词
              </a></dt>
            </dl>
            <dl>
              <dt id="lcatalog_TUTOR" onclick="TurnToTitle(&#39;catalog_TUTOR&#39;)"><a><i></i>导师
              </a></dt>
            </dl>
            <dl>
              <dt id="lcatalog_ZCDOI" onclick="TurnToTitle(&#39;catalog_ZCDOI&#39;)"><a><i></i>DOI
              </a></dt>
            </dl>
            <dl>
              <dt id="lcatalog_ZTCLS" onclick="TurnToTitle(&#39;catalog_ZTCLS&#39;)"><a><i></i>分类号
              </a></dt>
            </dl>
            <dl>
              <dt id="lcatalog_divimg" class="leftBar_Img" style="display:none;" onclick="TurnToTitle(&#39;catalog_divimg&#39;)"><a><i></i>文内图片
            </a></dt>
            </dl>
          </div>
          <div class="clHd">
      知识网络
      <input id="catalogIds" type="hidden" value=""></div>
          <div class="clBd">
            <dl>
              <dt id="lcatalog_ref" data-tit="引文网络" onclick="TurnToTitle(&#39;catalog_ref&#39;)"><a><i></i>引文网络
              </a></dt>
            </dl>
            <dl>
              <dt id="lcatalog_func601" data-tit="关联作者" onclick="GetAndShowFiles(this);"><a><i></i>关联作者
          </a></dt>
            </dl>
            <dl>
              <dt id="lcatalog_func604" data-tit="相似文献" onclick="GetAndShowFiles(this);"><a><i></i>相似文献
          </a></dt>
            </dl>
            <dl>
              <dt id="lcatalog_func605" data-tit="读者推荐" onclick="GetAndShowFiles(this);"><a><i></i>读者推荐
          </a></dt>
            </dl>
            <dl style="display:none;">
              <dt id="lcatalog_func602" data-tit="主题指数" onclick="GetAndShowFiles(this);"><a><i></i>主题指数
          </a></dt>
            </dl>
            <dl>
              <dt id="lcatalog_func603" data-tit="相关基金文献" onclick="GetAndShowFiles(this);"><a><i></i>相关基金文献
              </a></dt>
            </dl>
          </div>
        </div>
      </div>
      <div class="wxmain">
        <div class="wxTitle">
          <h2 class="title">开放式实验管理系统的设计与实现</h2><a class="btn-note" target="_blank" href="https://test.dr2am.cn/--/net/cnki/x/_/search/common/testlunbo?dbcode=CMFD&amp;tablename=CMFDTEMP&amp;filename=1020008317.nh&amp;filesourcetype=1&amp;__dp=http" __cpp="1" __dp="1"><img src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/kcms/detail/resource/gb/images/note-btn.gif?__dp=http" __cpp="1" __dp="1"></a><div class="author"><span><a target="_blank" onclick="
                    TurnPageToKnet(&#39;au&#39;,&#39;熊一鸣&#39;,&#39;44079026;&#39;)
                  ">熊一鸣</a></span></div>
          <div class="orgn"><span><a onclick="
                    TurnPageToKnet(&#39;in&#39;,&#39;江西财经大学&#39;,&#39;0188434&#39;)
                  ">江西财经大学</a></span></div>
          <div class="link"><a class="icon icon-output" href="javascript:void(0);" onclick="
        SubTurnExport(&#39;//kns.cnki.net/kns/ViewPage/viewsave.aspx&#39;,&#39;CMFDTEMP!1020008317.nh!1!0&#39;)
      "><i></i>导出/参考文献
    </a><span class="shareBoard" onmouseover="$(&#39;#sharedet&#39;).show();$(&#39;#this&#39;).addClass(&#39;shareBoardCUR&#39;)" onmouseout="$(&#39;#sharedet&#39;).hide();$(&#39;#this&#39;).removeClass(&#39;shareBoardCUR&#39;)"><a class="icon icon-share" href="https://test.dr2am.cn/--/net/cnki/kns/_/KCMS/detail/detail.aspx?__dp=http" __cpp="1" __dp="1"><i></i>分享<em></em></a><script> document.write(ShareAstr('shareHide'));</script></span><a class="icon icon-track" href="javascript:void(0);" id="RefTrack" onclick="RefTrack()"><i></i><label>创建引文跟踪</label></a><a class="icon icon-favor" href="javascript:void(0);" id="addfavtokpc" onclick="AddFavToKpc()"><i></i><label>收藏</label></a><a class="icon icon-print" href="javascript:void(0);" onclick="window.print();"><i></i>打印
    </a><a class="icon icon-comment" title="评论" href="https://test.dr2am.cn/--/net/cnki/kns/_/KCMS/detail/detail.aspx?__dp=http" style="display:none;" __cpp="1" __dp="1"><i></i>评论
    </a><span class="otherversions" style="display:none"></span></div>
        </div>
        <div class="wxInfo">
          <div class="wxBaseinfo">
            <p><label id="catalog_ABSTRACT">摘要：</label><span id="ChDivSummary" name="ChDivSummary">我国高校开放式实验管理普遍存在实验设备使用率较低、管理制度不完善,实验设备共享程度不高等诸多问题。要在更大范围推行开放式实验管理,就必须在开放式实验教学管理流程中,通过引入信息化管理,加大信息技术在其中的应用,才能真正发挥这种教学模式的开放性优势。本文基于B/S的三层体系结构,选择主流系统开发架构,设计开发了一个开放式实验室管理平台。为了便于平台的后期维护与系统功能扩展,整个平台的建立始终贯彻模块化的设计思想。开发平台的实现,依托微软的.NET,数据库管理采用ASP.NET技术,编程语言为C#。通过分析系统的用户需求,系统主要实现了系统平台管理、实验项目预约、共享信息分布等基本功能。测试与实际应用表明,本系统结构合理,数据库设计简单有效,能够满足普通高校开放式试验室管理的实际需求。</span><span><a id="ChDivSummaryMore" onclick="MoerSummary(&#39;ChDivSummary&#39;,&#39;ChDivSummaryMore&#39;,&#39;ChDivSummaryReset&#39;)" style="display:none">更多</a><a id="ChDivSummaryReset" onclick="ResetSummary(&#39;ChDivSummary&#39;,&#39;ChDivSummaryMore&#39;,&#39;ChDivSummaryReset&#39;)" style="display:none">还原</a></span><br></p><script type="text/javascript">
      AbstractFilter('ChDivSummary','ChDivSummaryMore','ChDivSummaryReset');
    </script><p><label id="catalog_KEYWORD">关键词：</label><a onclick="
                      TurnPageToKnet(&#39;kw&#39;,&#39;开放式实验管理&#39;,&#39;&#39;)
                    ">开放式实验管理;  </a><a onclick="
                      TurnPageToKnet(&#39;kw&#39;,&#39;用例模型&#39;,&#39;&#39;)
                    ">用例模型;  </a><a onclick="
                      TurnPageToKnet(&#39;kw&#39;,&#39;共享&#39;,&#39;&#39;)
                    ">共享;  </a><a onclick="
                      TurnPageToKnet(&#39;kw&#39;,&#39;模块化设计&#39;,&#39;&#39;)
                    ">模块化设计;  </a><a onclick="
                      TurnPageToKnet(&#39;kw&#39;,&#39;.NET框架&#39;,&#39;&#39;)
                    ">.NET框架;  </a></p>
            <p><label id="catalog_TUTOR">导师：</label>杨波;
                          </p>
            <p><label id="catalog_ZCDOI">DOI：</label>10.27175/d.cnki.gjxcu.2019.000139</p>
            <p><label id="catalog_ZTCLS">分类号：</label>TP311.52</p>
            <p><label id="catalog_divimg" style="display:none ;">文内图片：</label><div id="imgdiv" class="imgcont"></div>
            </p>
            <div class="dllink" id="DownLoadParts"><a target="_blank" onclick="WriteKrsDownLog()" class="icon icon-dlGreen" href="https://test.dr2am.cn/--/net/cnki/kns/_/kns/download.aspx?filename=ldxg0L5NzZShzYrpmTKRWMGREWKF0NZFlZPRFVyJ2MH1Ud41kc5NDazh1RCpnUjp3ZZ5mYUdWaP5mZ=0TQh1EZaRlW5ADZP9yURh2bvJWOXhVRzdmR19meSh3KxZzNvQGTDhWM0k3bwc1ZFlVbrIHTsxWb5U&amp;dflag=nhdown&amp;tablename=CMFDTEMP&amp;__dp=http" __cpp="1" __dp="1">整本下载
      </a><a target="_blank" onclick="WriteKrsDownLog()" class="icon icon-dlBlue" href="https://test.dr2am.cn/--/net/cnki/kns/_/kns/download.aspx?filename=ldxg0L5NzZShzYrpmTKRWMGREWKF0NZFlZPRFVyJ2MH1Ud41kc5NDazh1RCpnUjp3ZZ5mYUdWaP5mZ=0TQh1EZaRlW5ADZP9yURh2bvJWOXhVRzdmR19meSh3KxZzNvQGTDhWM0k3bwc1ZFlVbrIHTsxWb5U&amp;dflag=downpage&amp;tablename=CMFDTEMP&amp;__dp=http" __cpp="1" __dp="1">分页下载
      </a><a target="_blank" onclick="WriteKrsDownLog()" class="icon icon-dlBlue" href="https://test.dr2am.cn/--/net/cnki/kns/_/kns/download.aspx?filename=ldxg0L5NzZShzYrpmTKRWMGREWKF0NZFlZPRFVyJ2MH1Ud41kc5NDazh1RCpnUjp3ZZ5mYUdWaP5mZ=0TQh1EZaRlW5ADZP9yURh2bvJWOXhVRzdmR19meSh3KxZzNvQGTDhWM0k3bwc1ZFlVbrIHTsxWb5U&amp;dflag=catalog&amp;tablename=CMFDTEMP&amp;__dp=http" __cpp="1" __dp="1">分章下载
        </a><a target="_blank" onclick="WriteKrsDownLog()" class="icon icon-dlGreen" href="https://test.dr2am.cn/--/net/cnki/kreader/_//Kreader/RedriectPage.aspx?dbCode=cdmd&amp;filename=1020008317.nh&amp;tablename=CMFDTEMP&amp;__dp=http" __cpp="1" __dp="1">在线阅读
      </a></div>
            <div class="dllink-down">
              <div class="info">
                <div class="total"><span class="a"><label>下载：</label><b>29</b></span><span class="h"><label>页数：</label><b>59</b></span><span class="h"><label>大小：</label><b>2330K</b></span></div>
                <div class="hotspotCen" style="display:none;"><label>热点关注度：</label><div class="hotspot"><span value="" style="width:0%;" class="HotSpotPower"></span></div><span class="HotSpotValue" id="HotValue">0</span><b class="h">（注：最近下载、浏览的数量值）</b></div>
              </div>
              <div class="qr-code"><img alt="" src="https://test.dr2am.cn/--/net/cnki/app/_/Parts/QRCode/Get?source=KCMS&amp;text=http%3a%2f%2fm.cnki.net%2fcnkiday%2fappdownzwj.html%3ftype%3dCMFD%26id%3d1020008317.nh&amp;__dp=http" __cpp="1" __dp="1"><p class="text"><b>手机阅读本文</b><span>下载安装手机APP</span><span>扫码同步阅读本文</span></p><i class="icon-trangle"></i><div class="tip-pop">
                  <div class="inner">
                    <h6>即刻使用手机阅读</h6>
                    <div class="f-left first"><img alt="" src="https://test.dr2am.cn/--/net/cnki/kns/_/kcms/Detail/resource/gb/images/icon-qrcode-download.jpg?__dp=http" __cpp="1" __dp="1"><span>第一步</span><p>扫描二维码下载</p>
                      <p>"移动知网-全球学术快报"客户端</p>
                    </div>
                    <div class="f-left second"><img alt="" src="https://test.dr2am.cn/--/net/cnki/kns/_/kcms/Detail/resource/gb/images/icon-pop-sample.jpg?__dp=http" __cpp="1" __dp="1"><span>第二步</span><p>打开“全球学术快报”</p>
                      <p>点击首页右上角的扫描图标</p>
                    </div>
                    <div class="f-left third"><img alt="" src="https://test.dr2am.cn/--/net/cnki/app/_/Parts/QRCode/Get?source=KCMS&amp;text=http%3a%2f%2fm.cnki.net%2fcnkiday%2fappdownzwj.html%3ftype%3dCMFD%26id%3d1020008317.nh&amp;__dp=http" __cpp="1" __dp="1"><span>第三步</span><p>扫描二维码</p>
                      <p>手机同步阅读本篇文献</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div id="func608" class="wxsour"></div>
          <div class="wxsour sourlink"></div>
        </div>
        <div class="wxMod" id="catalog_ref" style="display:none">
          <h2 id="MapTitle" class="hd hdTab"><a class="active" href="javascript:void(0)" onclick="showRefChart(this)">引文网络</a><a href="javascript:void(0)" onclick="showFlash(this)">参考引证图谱</a></h2>
        </div>
        <div class="wxMod" id="ref_nodata" style="display:none">
          <h3 class="title2">
        引文网络
        <b class="titleTotle">未找到相关数据</b></h3>
        </div>
        <div id="MapArea" style="display:none">
          <div class="MapAreaLeft">
            <div class="map">
              <div class="jdwx" title="">节点文献</div>
              <div class="gywx"><a id="rl5" name="rl" title="（也称同引文献）与本文有相同参考文献的文献，与本文有共同研究背景或依据">共引文献</a><span id="rc5">(0)</span></div>
              <div class="tbywx"><a id="rl6" name="rl" title="与本文同时被作为参考文献引用的文献，与本文共同作为进一步研究的基础">同被引文献</a><span id="rc6">(0)</span></div>
              <div class="rjckwx"><a id="rl2" name="rl" title="本文参考文献的参考文献。进一步反映本文研究工作的背景和依据">二级参考文献</a><span id="rc2">(0)</span></div>
              <div class="ckwx"><a id="rl1" name="rl" title="反映本文研究工作的背景和依据">参考文献</a><span id="rc1">(0)</span></div>
              <div class="yzwx"><a id="rl3" name="rl" title="引用本文的文献。本文研究工作的继续、应用、发展或评价">引证文献</a><span id="rc3">(0)</span></div>
              <div class="rjyzwx"><a id="rl4" name="rl" title="本文引证文献的引证文献。更进一步反映本文研究工作的继续、发展或评价">二级引证文献</a><span id="rc4">(0)</span></div>
            </div>
            <div class="time"><span onmouseover="TimeAxisMoveRight(&#39;AxisFrameDivLeft&#39;);" onmouseout="TimeAxisStop();" onclick="TimeAxisMoveFaster();" class="ArrowLeftDisable" title="显示前面的年份" id="ArrowLeft_AxisFrameDivLeft"></span><div class="year" id="AxisFrameDivLeft"></div><span onmouseover="TimeAxisMoveLeft(&#39;AxisFrameDivLeft&#39;);" onmouseout="TimeAxisStop();" onclick="TimeAxisMoveFaster()" class="ArrowRightEnable" title="显示后面的年份" id="ArrowRight_AxisFrameDivLeft"></span><div class="TimeMiddle" id="AxisFrameDivCurrent" style="height: 55px;"></div><span onmouseover="TimeAxisMoveRight(&#39;AxisFrameDivRight&#39;);" onmouseout="TimeAxisStop();" onclick="TimeAxisMoveFaster()" class="ArrowLeftDisable" title="显示前面的年份" id="ArrowLeft_AxisFrameDivRight"></span><div class="year" id="AxisFrameDivRight"></div><span onmouseover="TimeAxisMoveLeft(&#39;AxisFrameDivRight&#39;);" onmouseout="TimeAxisStop();" onclick="TimeAxisMoveFaster()" class="ArrowRightEnable" title="显示后面的年份" id="ArrowRight_AxisFrameDivRight"></span><span style="display: none; visibility: hidden;" class="ArrowLeftEnable" alt=""></span><div class="clear"></div>
            </div>
            <div id="NodeValueDiv1" class="TimeHide" onmouseout="HideNodeValueDiv();" onmouseover="window.clearTimeout(window.timeout_NodeValueDiv);">
              <ul>
                <li><a id="NodeValueDiv1ReferType1Link" onclick="ChangeReferType(&#39;1&#39;);" title="反映本文研究工作的背景和依据">参考文献</a><span id="NodeValueDiv1ReferType1Level1"></span></li>
                <li><a id="NodeValueDiv1ReferType2Link" onclick="ChangeReferType(&#39;2&#39;);" title="本文参考文献的参考文献。进一步反映本文研究工作的背景和依据">二级参考文献</a><span id="NodeValueDiv1ReferType2Level2"></span></li>
              </ul>
            </div>
            <div id="NodeValueDiv2" class="TimeHide" onmouseout="HideNodeValueDiv();" onmouseover="window.clearTimeout(window.timeout_NodeValueDiv);">
              <ul>
                <li><a id="NodeValueDiv2ReferType4Link" onclick="ChangeReferType(&#39;3&#39;);" title="引用本文的文献。本文研究工作的继续、应用、发展或评价">引证文献</a><span id="NodeValueDiv2ReferType4Level1"></span></li>
                <li><a id="NodeValueDiv2ReferType16Link" onclick="ChangeReferType(&#39;4&#39;);" title="本文引证文献的引证文献。更进一步反映本文研究工作的继续、发展或评价">二级引证文献</a><span id="NodeValueDiv2ReferType16Level2"></span></li>
              </ul>
            </div>
            <div id="NodeValueDiv3" class="TimeHide" onmouseout="HideNodeValueDiv();" onmouseover="window.clearTimeout(window.timeout_NodeValueDiv);">
              <ul>
                <li><a id="NodeValueDiv3ReferType1Link" onclick="ChangeReferType(&#39;1&#39;);" title="反映本文研究工作的背景和依据">参考文献</a><span id="NodeValueDiv3ReferType1Level1"></span></li>
                <li><a id="NodeValueDiv3ReferType2Link" onclick="ChangeReferType(&#39;2&#39;);" title="本文参考文献的参考文献。进一步反映本文研究工作的背景和依据">二级参考文献</a><span id="NodeValueDiv3ReferType2Level2"></span></li>
                <li><a id="NodeValueDiv3ReferType4Link" onclick="ChangeReferType(&#39;3&#39;);" title="引用本文的文献。本文研究工作的继续、应用、发展或评价">引证文献</a><span id="NodeValueDiv3ReferType4Level1"></span></li>
                <li><a id="NodeValueDiv3ReferType16Link" onclick="ChangeReferType(&#39;4&#39;);" title="本文引证文献的引证文献。更进一步反映本文研究工作的继续、发展或评价">二级引证文献</a><span id="NodeValueDiv3ReferType16Level2"></span></li>
              </ul>
            </div>
          </div>
        </div><script type="text/javascript">       
        SetRefChartDataEx('CMFD','1020008317.nh','CMFDTEMP','2019');
      </script><iframe id="frame1" name="frame1" scrolling="no" height="0" frameborder="no" width="100%"></iframe>
        <div id="func607" class="wxMod"></div>
        <iframe id="framecatalog_CkFiles" name="framecatalog_CkFiles" width="100%" height="0" frameborder="no" scrolling="no" src=""></iframe>
        <iframe id="framecatalog_YzFiles" name="framecatalog_YzFiles" width="100%" height="0" frameborder="no" scrolling="no" src=""></iframe>
        <div id="func601" class="wxMod"></div>
        <div id="func604" class="wxMod"></div>
        <div id="func605" class="wxMod"></div>
        <div id="func602" class="wxMod"></div>
        <div id="func603" class="wxMod"></div>
      </div>
    </div><script type="text/javascript">
      try{ GetImgPath('1020008317.nh');}catch(err){};
      
      WriteToPage('1020008317.nh','CMFDTEMP','CMFD','CDMD','608');
      function LoadFilesFromId(oid)
      {
      if (!oid) return;
      switch (oid)
      {
      case "catalog_func607":  
      WriteToPage('1020008317.nh','CMFDTEMP','CMFD','CDMD','607');
      LoadFile('framecatalog_CkFiles','/kcms/detail/frame/list.aspx?filename=1020008317.nh&dbcode=CMFD&dbname=CMFDTEMP&reftype=1');
      LoadFile('framecatalog_YzFiles','/kcms/detail/frame/list.aspx?filename=1020008317.nh&dbcode=CMFD&dbname=CMFDTEMP&reftype=3');

      break;
      case "catalog_func601":  
      WriteToPage('1020008317.nh','CMFDTEMP','CMFD','CDMD','601');
      break;
      case "catalog_func602":
      WriteToPage('1020008317.nh','CMFDTEMP','CMFD','CDMD','602');
      break;
      case "catalog_func603":
      WriteToPage('1020008317.nh','CMFDTEMP','CMFD','CDMD','603');
      break;
      case "catalog_func604":
      WriteToPage('1020008317.nh','CMFDTEMP','CMFD','CDMD','604');
      break;
      case "catalog_func605":
      WriteToPage('1020008317.nh','CMFDTEMP','CMFD','CDMD','605');
      break;
      default : break;
      }
      }
    </script><div class="wxToolbar" id="wxDlToolbar">
      <div class="wxwrap">
        <div class="dllink"><a target="_blank" onclick="WriteKrsDownLog()" class="icon icon-dlGreen" href="https://test.dr2am.cn/--/net/cnki/kns/_/kns/download.aspx?filename=ldxg0L5NzZShzYrpmTKRWMGREWKF0NZFlZPRFVyJ2MH1Ud41kc5NDazh1RCpnUjp3ZZ5mYUdWaP5mZ=0TQh1EZaRlW5ADZP9yURh2bvJWOXhVRzdmR19meSh3KxZzNvQGTDhWM0k3bwc1ZFlVbrIHTsxWb5U&amp;dflag=nhdown&amp;tablename=CMFDTEMP&amp;__dp=http" __cpp="1" __dp="1">整本下载
          </a><a target="_blank" onclick="WriteKrsDownLog()" class="icon icon-dlBlue" href="https://test.dr2am.cn/--/net/cnki/kns/_/kns/download.aspx?filename=ldxg0L5NzZShzYrpmTKRWMGREWKF0NZFlZPRFVyJ2MH1Ud41kc5NDazh1RCpnUjp3ZZ5mYUdWaP5mZ=0TQh1EZaRlW5ADZP9yURh2bvJWOXhVRzdmR19meSh3KxZzNvQGTDhWM0k3bwc1ZFlVbrIHTsxWb5U&amp;dflag=downpage&amp;tablename=CMFDTEMP&amp;__dp=http" __cpp="1" __dp="1">分页下载
          </a><a target="_blank" onclick="WriteKrsDownLog()" class="icon icon-dlBlue" href="https://test.dr2am.cn/--/net/cnki/kns/_/kns/download.aspx?filename=ldxg0L5NzZShzYrpmTKRWMGREWKF0NZFlZPRFVyJ2MH1Ud41kc5NDazh1RCpnUjp3ZZ5mYUdWaP5mZ=0TQh1EZaRlW5ADZP9yURh2bvJWOXhVRzdmR19meSh3KxZzNvQGTDhWM0k3bwc1ZFlVbrIHTsxWb5U&amp;dflag=catalog&amp;tablename=CMFDTEMP&amp;__dp=http" __cpp="1" __dp="1">分章下载
            </a><a target="_blank" onclick="WriteKrsDownLog()" class="icon icon-dlGreen" href="https://test.dr2am.cn/--/net/cnki/kreader/_//Kreader/RedriectPage.aspx?dbCode=cdmd&amp;filename=1020008317.nh&amp;tablename=CMFDTEMP&amp;uid=&amp;__dp=http" __cpp="1" __dp="1">在线阅读
          </a></div>
        <div class="infotxt infotxtLar">
          <div class="total"><span class="a"><label>下载：</label><b>29</b></span><span class="h"><label>页数：</label><b>59</b></span><span class="h"><label>大小：</label><b>2330K</b></span></div>
          <div class="hotspotCen" style="display:none;"><label>热点关注度：</label><div class="hotspot"><span value="" style="width:0%;" class="HotSpotPower"></span></div><span class="HotSpotValue" id="HotValue">0</span><b class="h">（注：最近下载、浏览的数量值）</b></div>
        </div><script>FloatDownloadPartCantrol();</script></div>
    </div>
    <iframe scrolling="no" frameborder="0" id="headad" style="display:none;"></iframe>
    <div class="wait" id="waitDiv" style="visibility:hidden;">
          正在为您查找，请稍等...
        </div>
    <div class="dazhong-ad" style="padding:40px 0 0 0; width:1200px; margin:0 auto;margin-bottom:-23px;overflow:hidden;"><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/bianke/_/adfiles/ad/R018.js?sc=I138&amp;__dp=http" __cpp="1" __dp="1"></script></div>
    <div id="footerBox"></div>
    <div class="fixedbar">
      <div class="adviceside"><a class="fixCon" target="_blank" href="https://test.dr2am.cn/--/net/cnki/help/_/?__dp=http" __cpp="1" __dp="1">在线咨询</a><a target="_blank" class="fixFed" href="https://test.dr2am.cn/--/net/cnki/kns/_/knsvote/vote.aspx?__dp=http" __cpp="1" __dp="1">
          用户反馈
        </a></div>
      <div class="backtop hiddenV" id="backtop"><a id="backTopSide" href="javascript:scroll(0,0);" title="返回顶端"></a></div>
    </div><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/kns/_/KRS/Scripts/Recommend.js?__dp=http" __cpp="1" __dp="1"></script><script> toTop();</script><script>
      InsertCatalog();

      try{
        FlushLogin();
        modifyEcpHeader(true);
      }
      catch(es){}
      
      KLogin.getFooter();
    </script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/klib.min.js?v=20171013&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/KCMS/detail/js/min/Timeaxis.min.js?v=20171013&amp;__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/piccache/_/kdn/kcms/detail/js/piwikCommon70.js?__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript" src="https://test.dr2am.cn/--/net/cnki/ishufang/_/KRecord/krecord.min.js?__dp=http" __cpp="1" __dp="1"></script><script type="text/javascript">LoadScript('//piccache.cnki.net/kdn/KCMS/detail/js/min/cnkisug.min.js',function(){sugPara.IsTopK = false; sugPara.IsExp = false; sugPara.IsAttr = false;InitSug('http://acad3.cnki.net');});</script><script>
          setRecommendPic();
          isHasAddFav();
          isHasRefTrack();
        </script><input id="writebrowselog" type="hidden" value="5033"></body>
</html>