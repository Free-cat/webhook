package webhook_v1

import (
	"github.com/gin-gonic/gin"
	"webhook/internal/api/request"
)

func RegisterRequestRequests(c *gin.RouterGroup, requestImpl *request.Implementation) {
	c.GET("/:webhook_uuid", ListRequest(requestImpl))
}

// ListRequest godoc
// @Summary List requests
// @Description List requests
// @Tags request
// @Accept json
// @Produce json
// @Param webhook_uuid path string true "Webhook UUID"
// @Success 200 {object} []pkg.Request
// @Router /request/{webhook_uuid} [get]
func ListRequest(requestImpl *request.Implementation) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		webhookUuid := ctx.Param("webhook_uuid")
		listResponse, err := requestImpl.List(ctx, webhookUuid)
		if err != nil {
			ctx.JSON(500, gin.H{"message": err.Error()})
			return
		} else {
			ctx.JSON(200, listResponse)
		}
	}
	return fn
}
