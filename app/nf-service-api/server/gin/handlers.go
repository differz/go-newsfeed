package gin

import (
	"strconv"

	"github.com/VitaliiHurin/go-newsfeed/app/nf-service-api/api"
	"github.com/VitaliiHurin/go-newsfeed/app/nf-service-api/security"
	"github.com/VitaliiHurin/go-newsfeed/entity"
	"github.com/gin-gonic/gin"
)

func auth(a *api.API, c *gin.Context) (*entity.User, error) {
	wsse := c.Request.Header.Get("X-WSSE")
	token, err := security.ParseToken(wsse)
	if err != nil {
		return nil, err
	}
	user, err := a.GetUser(token.Username)
	if err != nil {
		return nil, err
	}
	err = a.SecurityManager.ValidateWSSEToken(token, string(user.Password))
	if err != nil {
		c.Error(err)
		return user, err
	}
	c.Header("X-WSSE", token.ToString())
	return user, nil
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

func handlePostRegistration(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, ok := c.GetPostForm("email")
		if !ok {
			c.Error(api.ErrInvalidArgument)
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			c.Error(api.ErrInvalidArgument)
		}

		user, err := a.Register(email, password)
		if err != nil {
			c.Error(err)
			return
		}
		c.Header("X-WSSE", a.SecurityManager.CreateWSSEToken(string(user.Email), string(user.Password)).ToString())
		responseSuccess(c, gin.H{
			"token": user.Token,
		})
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
		_, err := auth(a, c)
		if err != nil {
			c.Error(api.ErrUnauthorized)
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.Error(api.ErrInvalidArgument)
			return
		}
		err = a.MarkArticleAsRead(entity.ArticleID(id))
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, nil)
	}
}


func handleGetArticleMarkAsUnread(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := auth(a, c)
		if err != nil {
			c.Error(api.ErrUnauthorized)
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.Error(api.ErrInvalidArgument)
			return
		}
		err = a.MarkArticleAsUnread(entity.ArticleID(id))
		if err != nil {
			c.Error(err)
			return
		}
		responseSuccess(c, nil)
	}
}

func handlePostRestoreToken(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := auth(a, c)
		if err != nil && err != security.ErrWSSETokenExpired{
			c.Error(api.ErrUnauthorized)
			return
		}
		token, _ := c.GetPostForm("token")
		err = a.RestoreToken(user, token)
		if err != nil {
			c.Error(err)
			return
		}
		c.Header("X-WSSE", a.SecurityManager.CreateWSSEToken(string(user.Email), string(user.Password)).ToString())
		responseSuccess(c, gin.H{
			"token": user.Token,
		})
	}
}

func handlePostLogin(a *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, ok := c.GetPostForm("email")
		if !ok {
			c.Error(api.ErrInvalidArgument)
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			c.Error(api.ErrInvalidArgument)
		}
		user, err := a.GetUser(email)
		if err != nil {
			c.Error(err)
			return
		}
		err = a.SecurityManager.ValidatePassword(string(user.Password), string(user.Salt), password)
		if err != nil {
			c.Error(api.ErrUnauthorized)
			return
		}
		c.Header("X-WSSE", a.SecurityManager.CreateWSSEToken(string(user.Email), string(user.Password)).ToString())
		responseSuccess(c, gin.H{
			"token": user.Token,
		})
	}
}