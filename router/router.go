package router

import (
	"blog/apps/api"
	"blog/middlewares"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.Conf.Mode)
	r := gin.New()
	r.Use(middlewares.Logger(), middlewares.Cors(), gin.Recovery())

	apiGroup := r.Group("/api")
	apiGroup.Use(middlewares.JWY())
	{
		cate := apiGroup.Group("/cate")
		{
			cate.PUT("", api.EditCateInfo)
			cate.DELETE("", api.DestroyCate)
			cate.POST("", api.CreateCate)
		}

		user := apiGroup.Group("/user")
		{
			user.PUT("", api.ChangeUserInfo)
			user.PUT("pwd", api.ChangePwd)
			user.DELETE("", api.DeleteUser)
		}

		post := apiGroup.Group("/post")
		{
			post.POST("", api.AddPost)
			post.DELETE(":id", api.DeletePost)
			post.PUT("", api.EditPostInfo)

		}
	}

	frontApiGroup := r.Group("/api")
	{
		frontCate := frontApiGroup.Group("/cate")
		{
			frontCate.GET("getList", api.GetCateList)
			frontCate.GET("get/:id", api.GetCateInfo)
		}
		frontPost := frontApiGroup.Group("/post")
		{
			frontPost.GET("getList", api.GetPostList)
			frontPost.GET("get/:id", api.GetPost)
			frontPost.GET("getPostsCate/:id", api.GetPostsCate)
		}
		fontLogin := frontApiGroup.Group("login")
		{
			fontLogin.POST("login", api.Login)
			fontLogin.POST("loginfront", api.FrontLogin)
		}
		frontUser := frontApiGroup.Group("/user")
		{
			frontUser.POST("", api.CreateUser)
			frontUser.GET(":id", api.GetUserInfo)
			frontUser.GET("userList", api.GetUserList)
		}
	}

	_ = r.Run(utils.Conf.HTTPPort)
}
