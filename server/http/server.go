package http

import (
	"encoding/json"
	"net/http"
	"sokwva/acfun/dougaInfo/cache"
	"sokwva/acfun/dougaInfo/conf"
	"sokwva/acfun/dougaInfo/fetch"
	"sokwva/acfun/dougaInfo/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	RespOk = iota
	RespInvalidReq
	RespServerSideErr
)

func ginModeMapper(raw string) string {
	if raw == gin.DebugMode {
		return gin.DebugMode
	}
	if raw == gin.ReleaseMode {
		return gin.ReleaseMode
	}
	if raw == gin.TestMode {
		return gin.TestMode
	}
	return gin.ReleaseMode
}

func apiOkRtn(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}

func setCache(name string, result interface{}, expireTime uint) {
	x, redisErr := json.Marshal(result)
	if redisErr == nil {
		if expireTime == 0 {
			cache.Handle.Set(name, x, 0)
		}
		cache.Handle.Set(name, x, time.Duration(expireTime))
	}
}

func Start() {
	ginRoot := gin.Default()
	rootGrp := ginRoot.Group("/")
	{
		rootGrp.GET("/ping", func(ctx *gin.Context) {
			apiOkRtn(ctx, RespOk, "I'm ok!")
		})
	}
	videoGrp := ginRoot.Group("/video")
	{
		videoGrp.GET("/", func(ctx *gin.Context) {
			acid := ctx.Query("acid")
			if acid == "" {
				apiOkRtn(ctx, RespInvalidReq, "invalid param: acid")
				return
			}
			var err error
			var dr *structs.DougaInfo
			var result interface{}
			r, e := cache.Handle.Get(acid)
			if e != nil {
				dr, err = fetch.GetVideoInfo(acid)
				if err == nil {
					result = dr
					setCache(acid, dr, 0)
				}
			} else {
				rs := &structs.DougaInfo{}
				err = json.Unmarshal([]byte(r.(string)), rs)
				if err != nil {
					result, err = fetch.GetVideoInfo(acid)
				}
				result = rs
			}
			if err != nil {
				apiOkRtn(ctx, RespServerSideErr, err.Error())
				return
			}
			apiOkRtn(ctx, RespOk, result)
		})
	}
	bangumiGrp := ginRoot.Group("/bangumi")
	{
		bangumiGrp.GET("/", func(ctx *gin.Context) {
			aaid := ctx.Query("aaid")
			if aaid == "" {
				apiOkRtn(ctx, RespInvalidReq, "invalid param: aaid")
				return
			}
			var err error
			var dr *structs.BangumiInfo
			var result interface{}
			r, e := cache.Handle.Get(aaid)
			if e != nil {
				dr, err = fetch.GetBangumiInfo(aaid)
				if err == nil {
					result = dr
					setCache(aaid, dr, 3600)
				}
			} else {
				rs := &structs.BangumiInfo{}
				err = json.Unmarshal([]byte(r.(string)), rs)
				if err != nil {
					result, err = fetch.GetBangumiInfo(aaid)
				}
				result = rs
			}
			if err != nil {
				apiOkRtn(ctx, RespServerSideErr, err.Error())
				return
			}
			apiOkRtn(ctx, RespOk, result)
		})
	}
	userGrp := ginRoot.Group("/user")
	{
		userGrp.GET("/video", func(ctx *gin.Context) {
			userId := ctx.Query("userId")
			if userId == "" {
				apiOkRtn(ctx, RespInvalidReq, "invalid param: userId")
				return
			}
			page := ctx.Query("page")
			if page == "" {
				page = "1"
			}
			pageSize := ctx.Query("pageSize")
			if pageSize == "" {
				pageSize = "20"
			}
			pageInt, _ := strconv.Atoi(page)
			pageSizeInt, _ := strconv.Atoi(pageSize)
			var err error
			var dr *[]structs.DougaInfo
			var result interface{}
			r, e := cache.Handle.Get(userId + "-" + page + "-" + pageSize)
			if e != nil {
				dr, err = fetch.GetUserVideoDouga(userId, uint(pageInt), uint(pageSizeInt))
				if err == nil {
					result = dr
					setCache(userId+"-"+page+"-"+pageSize, dr, 600)
				}
			} else {
				rs := &[]structs.DougaInfo{}
				err = json.Unmarshal([]byte(r.(string)), rs)
				if err != nil {
					result, err = fetch.GetUserVideoDouga(userId, uint(pageInt), uint(pageSizeInt))
				}
				result = rs
			}
			if err != nil {
				apiOkRtn(ctx, RespServerSideErr, err.Error())
				return
			}
			apiOkRtn(ctx, RespOk, result)
		})
	}
	gin.SetMode(ginModeMapper(conf.Conf.Server.Mode))
	ginRoot.Run(conf.Conf.Server.Address + ":" + conf.Conf.Server.Port)
}
