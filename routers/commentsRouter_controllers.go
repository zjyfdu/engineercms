package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Category",
			Router: `/category/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "AddCategory",
			Router: `/category/addcategory`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "CategoryTitle",
			Router: `/categorytitle`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"],
		beego.ControllerComments{
			Method: "ErrLog",
			Router: `/errlog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"],
		beego.ControllerComments{
			Method: "InfoLog",
			Router: `/infolog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "AddWxArticle",
			Router: `/addwxarticle`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "AddWxArticles",
			Router: `/addwxarticles/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetListArticles",
			Router: `/getlistarticles`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetWxArticle",
			Router: `/getwxarticle/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetWxArticles",
			Router: `/getwxarticles`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetWxArticless",
			Router: `/getwxarticless/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowAccessContext",
			Router: `/flowaccesscontext`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowAccessContextList",
			Router: `/flowaccesscontextlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowAction",
			Router: `/flowaction`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowActionList",
			Router: `/flowactionlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowDoc",
			Router: `/flowdoc`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowEvent",
			Router: `/flowdocevent`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowEventList",
			Router: `/flowdoceventlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowDocList",
			Router: `/flowdoclist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowDocumentDetail",
			Router: `/flowdocumentdetail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowDocumentList",
			Router: `/flowdocumentlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowGroup",
			Router: `/flowgroup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowGroupList",
			Router: `/flowgrouplist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowGroupMailbox",
			Router: `/flowgroupmailbox`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowGroupRole",
			Router: `/flowgrouprole`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowGroupRoleList",
			Router: `/flowgrouprolelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowGroupUsersList",
			Router: `/flowgroupuserslist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowNext",
			Router: `/flownext`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowNode",
			Router: `/flownode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowNodeList",
			Router: `/flownodelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowPermission",
			Router: `/flowpermission`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowRole",
			Router: `/flowrole`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowRoleList",
			Router: `/flowrolelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowRolePermissionList",
			Router: `/flowrolepermissionlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowState",
			Router: `/flowstate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowStateList",
			Router: `/flowstatelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowTransition",
			Router: `/flowtransition`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowTransitionList",
			Router: `/flowtransitionlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowType",
			Router: `/flowtype`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowTypeDelete",
			Router: `/flowtypedelete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowTypeList",
			Router: `/flowtypelist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowTypeUpdate",
			Router: `/flowtypeupdate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowUser",
			Router: `/flowuser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowUserGroup",
			Router: `/flowusergroup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowUserList",
			Router: `/flowuserlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowUserMailbox",
			Router: `/flowusermailbox`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowWorkflow",
			Router: `/flowworkflow`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "FlowWorkflowList",
			Router: `/flowworkflowlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FlowController"],
		beego.ControllerComments{
			Method: "WorkFlow",
			Router: `/workflow`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
		beego.ControllerComments{
			Method: "UploadWxImg",
			Router: `/uploadwximg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
		beego.ControllerComments{
			Method: "UploadWxImgs",
			Router: `/uploadwximgs/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"],
		beego.ControllerComments{
			Method: "WxLogin",
			Router: `/wxlogin/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"],
		beego.ControllerComments{
			Method: "WxPdf",
			Router: `/wxpdf/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"],
		beego.ControllerComments{
			Method: "WxStandardPdf",
			Router: `/wxstandardpdf/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:RegistController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:RegistController"],
		beego.ControllerComments{
			Method: "WxRegist",
			Router: `/wxregist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "AddWxLike",
			Router: `/addwxlike/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "AddWxRelease",
			Router: `/addwxrelease/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "DeleteWxRelease",
			Router: `/deletewxrelease/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"],
		beego.ControllerComments{
			Method: "SearchWxDrawings",
			Router: `/searchwxdrawings`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
		beego.ControllerComments{
			Method: "SearchWxStandards",
			Router: `/searchwxstandards`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
