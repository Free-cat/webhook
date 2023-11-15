package webhook_v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"webhook/internal/api/webhook"
	serviceErrors "webhook/internal/service/errors"
	desc "webhook/pkg"
)

func RegisterWebhookRequests(c *gin.RouterGroup, webhookImpl *webhook.Implementation) {
	c.POST("/", CreateWebhook(webhookImpl))
	c.GET("/:uuid", GetWebhook(webhookImpl))
	c.DELETE("/:uuid", DeleteWebhook(webhookImpl))
}

// CreateWebhook godoc
// @Summary Create webhook
// @Description Create webhook
// @Tags webhook
// @Accept json
// @Produce json
// @Param webhook body pkg.CreateWebhookRequest true "Webhook"
// @Success 200 {object} pkg.CreateWebhookResponse
// @Router /webhook [post]
func CreateWebhook(webhookImpl *webhook.Implementation) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var request desc.CreateWebhookRequest
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			return
		} else {
			createResponse, err := webhookImpl.Create(ctx, &request)
			if err != nil {
				ctx.JSON(500, gin.H{"message": err.Error()})
				return
			} else {
				ctx.JSON(200, createResponse)
			}
		}
	}

	return fn
}

// GetWebhook godoc
// @Summary Get webhook
// @Description Get webhook
// @Tags webhook
// @Accept json
// @Produce json
// @Param uuid path string true "Webhook UUID"
// @Success 200 {object} pkg.CreateWebhookResponse
// @Router /webhook/{uuid} [get]
func GetWebhook(webhookImpl *webhook.Implementation) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")
		getResponse, err := webhookImpl.Get(ctx, uuid)
		if err != nil {
			if errors.Is(err, serviceErrors.NotFoundError) {
				ctx.JSON(404, gin.H{"message": err.Error()})
			} else {
				ctx.JSON(500, gin.H{"message": err.Error()})
			}
			return
		} else {
			ctx.JSON(200, getResponse)
		}
	}
	return fn
}

// DeleteWebhook godoc
// @Summary Delete webhook
// @Description Delete webhook
// @Tags webhook
// @Accept json
// @Produce json
// @Param uuid path string true "Webhook UUID"
// @Success 200 {object} string
// @Router /webhook/{uuid} [delete]
func DeleteWebhook(webhookImpl *webhook.Implementation) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		uuid := ctx.Param("uuid")
		deleteResponse, err := webhookImpl.Delete(ctx, uuid)
		if err != nil {
			ctx.JSON(500, gin.H{"message": err.Error()})
			return
		} else {
			ctx.JSON(200, deleteResponse)
		}
	}
	return fn
}
