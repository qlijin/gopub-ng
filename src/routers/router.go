package routers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/linclin/gopub/src/controllers"
	apicontrollers "github.com/linclin/gopub/src/controllers/api"
	confcontrollers "github.com/linclin/gopub/src/controllers/conf"
	othercontrollers "github.com/linclin/gopub/src/controllers/other"
	p2pcontrollers "github.com/linclin/gopub/src/controllers/p2p"
	recordcontrollers "github.com/linclin/gopub/src/controllers/record"
	taskcontrollers "github.com/linclin/gopub/src/controllers/task"
	usercontrollers "github.com/linclin/gopub/src/controllers/user"
	wallecontrollers "github.com/linclin/gopub/src/controllers/walle"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowOrigins:    []string{"*"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "UserToken", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		MaxAge:          5 * time.Minute,
	}))

	beego.Router("/login", &controllers.LoginController{}, "post:Post")
	beego.Router("/logout", &controllers.LogoutController{}, "post:Post")

	beego.Router("/loginbydocke", &controllers.LoginByDockerController{}, "get:Get")
	beego.Router("/changePasswd", &controllers.ChangePasswdController{}, "post:Post")
	beego.Router("/register", &controllers.RegisterController{}, "post:Post")

	ns_api := beego.NewNamespace("/api",
		beego.NSRouter("/get/walle/detection", &wallecontrollers.DetectionController{}, "get:Get"),
		beego.NSRouter("/get/walle/detectionssh", &wallecontrollers.DetectionsshController{}, "get:Get"),
		beego.NSRouter("/get/walle/release", &wallecontrollers.ReleaseController{}, "get:Get"), //TODO
		beego.NSRouter("/get/walle/md5", &wallecontrollers.GetMd5Controller{}, "get:Get"),
		beego.NSRouter("/get/walle/flush", &wallecontrollers.FlushController{}, "get:Get"),
		beego.NSRouter("/get/conf/list", &confcontrollers.ListController{}, "get:Get"),
		beego.NSRouter("/get/conf/mylist", &confcontrollers.MyListController{}, "get:Get"),
		beego.NSRouter("/get/conf/get", &confcontrollers.ConfController{}, "get:Get"), //TODO
		beego.NSRouter("/post/conf/save", &confcontrollers.SaveController{}, "get:Get"),
		beego.NSRouter("/get/conf/del", &confcontrollers.DelController{}, "get:Get"),
		beego.NSRouter("/get/conf/copy", &confcontrollers.CopyController{}, "get:Get"),
		beego.NSRouter("/get/conf/tags", &confcontrollers.TagsController{}, "get:Get"),
		beego.NSRouter("/get/conf/lock", &confcontrollers.LockController{}, "get:Get"),
		beego.NSRouter("/get/conf/server_groups", &confcontrollers.ServerGroupsController{}, "get:Get"),
		beego.NSRouter("/get/conf/groupinfo", &confcontrollers.GroupInfoController{}, "get:Get"),

		beego.NSRouter("/get/git/branch", &wallecontrollers.BranchController{}, "get:Get"),
		beego.NSRouter("/get/git/commit", &wallecontrollers.CommitController{}, "get:Get"),
		beego.NSRouter("/get/git/gitpull", &wallecontrollers.GitpullController{}, "get:Get"),
		beego.NSRouter("/get/git/gitlog", &wallecontrollers.GitlogController{}, "get:Get"),
		beego.NSRouter("/get/git/tag", &wallecontrollers.TagController{}, "get:Get"),

		beego.NSRouter("/get/jenkins/commit", &wallecontrollers.JenkinsController{}, "get:Get"),

		beego.NSRouter("/get/task/list", &taskcontrollers.ListController{}, "get:Get"),
		beego.NSRouter("/get/task/chart", &taskcontrollers.TaskChartController{}, "get:Get"),
		beego.NSRouter("/post/task/save", &taskcontrollers.SaveController{}, "get:Get"),
		beego.NSRouter("/get/task/get", &taskcontrollers.TaskController{}, "get:Get"),
		beego.NSRouter("/get/task/changes", &taskcontrollers.ChangesController{}, "get:Get"),
		beego.NSRouter("/get/task/last", &taskcontrollers.LastTaskController{}, "get:Get"),
		beego.NSRouter("/get/task/rollback", &taskcontrollers.RollBackController{}, "get:Get"),
		beego.NSRouter("/get/task/del", &taskcontrollers.DelController{}, "get:Get"),

		beego.NSRouter("/get/p2p/task", &p2pcontrollers.TaskController{}, "get:Get"),
		beego.NSRouter("/get/p2p/check", &p2pcontrollers.CheckController{}, "get:Get"),
		beego.NSRouter("/post/p2p/agent", &p2pcontrollers.AgentController{}, "get:Get"),
		beego.NSRouter("/get/p2p/send", &p2pcontrollers.SendAgentController{}, "get:Get"),

		beego.NSRouter("/get/record/list", &recordcontrollers.ListController{}, "get:Get"),

		beego.NSRouter("/get/other/noauto", &othercontrollers.NoAutoController{}, "get:Get"),
		beego.NSRouter("/get/test/api", &controllers.TestApiController{}, "get:Get"),
		beego.NSRouter("/get/user/project", &usercontrollers.UserProjectController{}, "get:Get"),
		beego.NSRouter("/get/user", &usercontrollers.UserController{}, "get:Get"),
		// beego.NSInclude(
		// 	&controllers.ClusterController{},
		// ),
	)
	beego.AddNamespace(ns_api)

	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/token",
			beego.NSInclude(
				&apicontrollers.TokenController{},
			),
		),
		beego.NSNamespace("/task",
			beego.NSInclude(
				&apicontrollers.TaskController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
