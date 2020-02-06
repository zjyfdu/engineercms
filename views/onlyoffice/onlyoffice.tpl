<!DOCTYPE html>
<html style="height: 100%;">
	<head>
	  <title>{{.Doc.FileName}}</title>

	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
	<meta name="renderer" content="webkit">
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
	
	</head>

	<body style="height: 100%; margin: 0;">
		<div id="placeholder" style="height: 100%"></div>
    <script type="text/javascript" src="http://192.168.100.21:9000/web-apps/apps/api/documents/api.js"></script>
    <script type="text/javascript">
      var onAppReady = function() {
          console.log("ONLYOFFICE Document Editor is ready");
      };
      var onCollaborativeChanges = function () {
          console.log("The document changed by collaborative user");
      };
      var onDocumentReady = function() {
          console.log("Document is loaded");
      };
      var onDocumentStateChange = function (event) {
          if (event.data) {
              console.log("The document changed");
          } else {
              console.log("Changes are collected on document editing service");
          }
      };
      var onDownloadAs = function (event) {
          console.log("ONLYOFFICE Document Editor create file: " + event.data);
      };
      var onError = function (event) {
          console.log("ONLYOFFICE Document Editor reports an error: code " + event.data.errorCode + ", description " + event.data.errorDescription);
      };
      var onOutdatedVersion = function () {
          location.reload(true);
      };
      var onRequestEditRights = function () {
          // console.log("ONLYOFFICE Document Editor requests editing rights");
          // document.location.reload();
          var he=location.href.replace("view","edit");
          location.href=he;
      };

    	//历史版本保留1个月。比如Unix时间戳（Unix timestamp）expires=1524547423
      var onRequestHistory = function() {
      	docEditor.refreshHistory({
        "currentVersion": {{.currentversion}},
        "history":{{.onlyhistory}},

  			});
			};

			var onRequestHistoryClose = function() {
  		  document.location.reload();
			};

			var onRequestHistoryData = function(event) {
    		var version = event.data;
        var history={{.onlyhistory}};
        var fileUrl="";
        var changeUrl="";
        var key="";
        var previousKey="";
        var previousUrl="";
    		for(var i=history.length-1;i>=0;i--){
    			if (version==history[i].version){
            changeUrl=history[i].changesurl;
            fileUrl=history[i].fileurl;
            key=history[i].key;
            if(i>0){
              previousKey=history[i-1].key;
              previousUrl=history[i-1].fileurl;
            }else{
              previousKey=key;
              previousUrl=fileUrl;
            }
            break;
					}
    		}
				var changeUrl2="http://192.168.100.21:8088"+changeUrl.replace(/\u0026/,"&");
        var previousurl="http://192.168.100.21:8088"+previousUrl
        var fileurl="http://192.168.100.21:8088"+fileUrl
    		docEditor.setHistoryData({
      		"changesUrl":changeUrl2,
      		"key": key,
      		"previous": {
      		  "key": previousKey,//这里不影响版本切换。与上个版本对比
      		  "url": previousurl//previousUrl//http://192.168.100.3700:9000/cache/files/1521953170330601700_4540/output.docx/output.docx?md5=eSwnrSSumTeMuh59IoXhCQ==&expires=1524547423&disposition=attachment&ooname=output.docx这里影响版本
      		},
      		"url": fileurl,//fileUrl,
      		"version": version
    		})
			};

    	window.docEditor = new DocsAPI.DocEditor("placeholder",
      	{
        "events": {
          "onAppReady": onAppReady,
          "onCollaborativeChanges": onCollaborativeChanges,
          "onDocumentReady": onDocumentReady,
          "onDocumentStateChange": onDocumentStateChange,
          "onDownloadAs": onDownloadAs,
          "onError": onError,
          "onRequestEditRights": onRequestEditRights,
          "onRequestHistory": onRequestHistory,
          "onRequestHistoryClose": onRequestHistoryClose,
          "onRequestHistoryData": onRequestHistoryData
        },

      	"document": {
          "fileType": "{{.fileType}}",
          "key": "{{.Key}}",//"Khirz6zTPdfd7"
          "title": "{{.Doc.FileName}}",
          "url": "http://192.168.100.21:8088/attachment/onlyoffice/{{.Doc.FileName}}?hotqinsessionid={{.Sessionid}}",
          "info": {
            "author": "John Smith",
            "created": "2010-07-07 3:46 PM",
            "folder": "Example Files",
            "sharingSettings": [
              {
                "permissions": "Full Access",
                "user": "John Smith"
              },
              {
                "permissions": "Read Only",
                "user": "Kate Cage"
              },
              {
                "permissions": "Deny Access",
                "user": "Albet Tlanp"
              }, 
            ]
          },
          "permissions": {
          	"comment": {{.Comment}},//true,
          	"download": {{.Download}},//true,
          	"edit": {{.Edit}},
          	"print": {{.Print}},//true,
          	"review": {{.Review}}//true
        	},
        },
        "documentType": "{{.documentType}}",
        "editorConfig": {
          "callbackUrl": "http://192.168.100.21:8088/url-to-callback?id={{.Doc.Id}}",
        	"createUrl": "https://example.com/url-to-create-document/",
          "user": {
            "id": {{.Uid}},
            "name": "{{.Username}}"
          },
					"customization": {
            "chat": true,
            "commentAuthorOnly": false,
            "compactToolbar": false,
            "customer": {
              "address": "116# Tianshou Road,Guangzhou China",
              "info": "QQ504284",
              "logo": "http://192.168.100.21:8088/static/img/user.jpg",//logo-big.png
              "mail": "xc-qin@163.com",
              "name": "Qin Xiao Chuan",
              "www": "github.com/3xxx"
            },
            "feedback": {
              "url": "http://192.168.100.21:8088/onlyoffice",
              "visible": true
            },
            "forcesave": false,
            "goback": {
              "text": "Go to Documents",
              "url": "http://192.168.100.21:8088/onlyoffice"
            },
            "logo": {
              "image": "http://192.168.100.21:8088/static/img/oo.png",//logo.png
              "imageEmbedded": "http://192.168.100.21:8088/static/img/oo.png",//logo_em.png
              "url": "http://192.168.100.21:8088/onlyoffice"
            },
            "showReviewChanges": false,
            "zoom": 100,
        	},
        	"embedded": {
            "embedUrl": "https://example.com/embedded?doc=exampledocument1.docx",
            "fullscreenUrl": "https://example.com/embedded?doc=exampledocument1.docx#fullscreen",
            "saveUrl": "https://example.com/download?doc=exampledocument1.docx",
            "shareUrl": "https://example.com/view?doc=exampledocument1.docx",
            "toolbarDocked": "top"
        	},
          "lang": "zh-CN",//"en-US",
					"mode": {{.Mode}},//"view",//edit
					"recent": [
            {
              "folder": "Example Files",
              "title": "exampledocument1.docx",
              "url": "https://example.com/exampledocument1.docx"
            },
            {
              "folder": "Example Files",
              "title": "exampledocument2.docx",
              "url": "https://example.com/exampledocument2.docx"
            }
        	]
        },
        "height": "100%",
        // "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.t-IDcSemACt8x4iTMCda8Yhe3iZaWbvV5XKSTbuAn0M",
    		"type": {{.Type}},//"desktop",embedded,mobile访问文档的平台类型 网页嵌入
        "width": "100%"
      });
   	</script>
	</body>
</html>