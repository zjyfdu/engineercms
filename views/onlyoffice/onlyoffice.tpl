<!DOCTYPE html>
<html style="height: 100%;">
<head>
    <title>{{.Doc.FileName}}</title>
</head>
<body style="height: 100%; margin: 0;">

    <div id="placeholder" style="height: 100%"></div>
    <script type="text/javascript" src="http://192.168.1.104:9000/web-apps/apps/api/documents/api.js"></script>
    <script type="text/javascript">
        window.docEditor = new DocsAPI.DocEditor("placeholder",
            {
                "document": {
                    "fileType": "{{.fileType}}",
                    "key": "{{.Key}}",//"Khirz6zTPdfd7"
                    "title": "{{.Doc.FileName}}",
                    "url": "http://192.168.1.104:8088/attachment/onlyoffice/{{.Doc.FileName}}"
                },
                "documentType": "{{.documentType}}",
                "editorConfig": {
                    "callbackUrl": "http://192.168.1.104:8088/url-to-callback?id={{.Doc.Id}}",
                    "user": {
                        "id": "{{.Uid}}",
                        "name": "{{.Uname}}"
                    },
                    "lang": "zh-CN",//"en-US",
                    "customization": {
                        "autosave": true,
                        "forcesave": true,
                    },
                },
                "height": "100%",
                "width": "100%"
            });
    </script>
</body>
</html>