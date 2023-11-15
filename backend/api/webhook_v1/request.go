package webhook_v1

import (
	"github.com/gin-gonic/gin"
	"io"
	"time"
	"webhook/internal/api/request"
	"webhook/internal/converter"
)

func RegisterRequestRequests(c *gin.RouterGroup, requestImpl *request.Implementation) {
	c.GET("/:webhook_uuid", ListRequest(requestImpl))
	c.GET("/sse/:webhook_uuid", SubscribeUpdates(requestImpl))
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

// Server send events for updates based on redis pub sub
func SubscribeUpdates(requestImpl *request.Implementation) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		webhookUuid := ctx.Param("webhook_uuid")
		ch, err := requestImpl.Subscribe(ctx, webhookUuid)
		if err != nil {
			ctx.JSON(500, gin.H{"message": err.Error()})
			return
		}
		ctx.Stream(func(w io.Writer) bool {
			pingTicker := time.NewTicker(5 * time.Second) // Send a ping every 30 seconds
			defer pingTicker.Stop()

			for {
				select {
				case request, ok := <-ch:
					if !ok {
						return false
					}
					ctx.SSEvent("message", converter.ToRequestFromService(request))
				case <-pingTicker.C:
					ctx.SSEvent("ping", "keep-alive") // Send a ping event
				case <-ctx.Done():
					return false
				}

				// Flush the buffer to client
				ctx.Writer.Flush()

				// Continue looping
				return true
			}
		})
	}
	return fn
}
