package public

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
	"webhook/internal/api/request"
	desc "webhook/pkg"
)

func RegisterRequestRequests(c *gin.RouterGroup, requestImpl *request.Implementation) {
	c.GET("/:webhookUuid", DumpRequest(requestImpl))
	c.POST("/:webhookUuid", DumpRequest(requestImpl))
	c.PUT("/:webhookUuid", DumpRequest(requestImpl))
	c.DELETE("/:webhookUuid", DumpRequest(requestImpl))
	c.OPTIONS("/:webhookUuid", DumpRequest(requestImpl))
	c.HEAD("/:webhookUuid", DumpRequest(requestImpl))
	c.PATCH("/:webhookUuid", DumpRequest(requestImpl))
}

func DumpRequest(requestImpl *request.Implementation) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		webhookUuid := ctx.Param("webhookUuid")
		dumpResponse, err := requestImpl.Dump(ctx, webhookUuid, dumpToRequest(ctx))
		if err != nil {
			ctx.JSON(500, gin.H{"message": err.Error()})
			return
		} else {
			ctx.JSON(200, dumpResponse)
		}
	}
	return fn
}

func dumpToRequest(c *gin.Context) *desc.Request {
	bodyData, err := c.GetRawData()
	if err != nil {
		return nil
	}

	return &desc.Request{
		Uuid:      uuid.New().String(),
		Host:      c.ClientIP(),
		Path:      c.Request.URL.Path,
		Method:    c.Request.Method,
		Body:      string(bodyData),
		Headers:   c.Request.Header,
		CreatedAt: time.Now().Unix(),
	}
}
