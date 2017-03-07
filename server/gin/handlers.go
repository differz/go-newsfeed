package gin

import (
	"strconv"

	"github.com/VitaliiHurin/go-newsfeed/api"
	"github.com/VitaliiHurin/go-newsfeed/entity"
	"github.com/gin-gonic/gin"
)

func auth(a *api.API, c *gin.Context) (*entity.User, error) {
	token := c.Request.Header.Get("session")
	if token == "" {
		return nil, api.ErrUnauthorized
	}
	return a.GetUser(token)
}

func handleGetServices(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := auth(a, c)
		if err != nil {
			c.Error(api.ErrUnauthorized)
			return
		}
		services, err := a.GetServices()
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, services)
	}
}

func handleGetRegistration(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := auth(a, c)
		token, err := a.Register(user)
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, gin.H{
			"token": token,
		})
	}
}

func handleGetLogout(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := auth(a, c)
		err := a.Logout(user)
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, nil)
	}
}

func handleGetUserTags(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := auth(a, c)
		if err != nil {
			c.Error(api.ErrUnauthorized)
			return
		}
		tags, err := a.GetUserTags(user)
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, tags)
	}
}

func handlePostUserTags(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := auth(a, c)
		if err != nil {
			c.Error(api.ErrUnauthorized)
			return
		}
		tag, ok := c.GetPostForm("tag")
		if !ok {
			c.Error(api.ErrInvalidArgument)
			return
		}
		err = a.AddUserTag(user, tag)
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, nil)
	}
}

func handleDeleteUserTags(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := auth(a, c)
		if err != nil {
			c.Error(api.ErrUnauthorized)
			return
		}
		tag, ok := c.GetQuery("tag")
		if !ok {
			c.Error(api.ErrInvalidArgument)
			return
		}
		err = a.DeleteUserTag(user, tag)
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, nil)
	}
}

func handleGetArticles(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := auth(a, c)
		if err != nil {
			c.Error(api.ErrUnauthorized)
			return
		}
		articles, err := a.GetArticles(user)
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, articles)
	}
}

func handleGetArticleMarkAsRead(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := auth(a, c)
		if err != nil {
			c.Error(api.ErrUnauthorized)
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.Error(api.ErrInvalidArgument)
			return
		}
		err = a.MarkArticleAsRead(user, entity.ArticleID(id))
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, nil)
	}
}


func handleGetArticleMarkAsUnread(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := auth(a, c)
		if err != nil {
			c.Error(api.ErrUnauthorized)
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.Error(api.ErrInvalidArgument)
			return
		}
		err = a.MarkArticleAsUnread(user, entity.ArticleID(id))
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, nil)
	}
}