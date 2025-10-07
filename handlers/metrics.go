package handlers

import (
	"educabot.com/bookshop/models"
	"educabot.com/bookshop/services"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type MetricController struct {
	metricService *services.MetricService
}

func NewMetricController(metricService *services.MetricService) *MetricController {
	return &MetricController{metricService: metricService}
}

func (mc *MetricController) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var query models.GetMetricsRequest
		err := ctx.ShouldBindQuery(&query)
		if err != nil {
			slog.ErrorContext(ctx, "GetMetrics: failed to bind query", "err", err)
			return
		}

		res := mc.metricService.GetMetrics(ctx, query)

		ctx.JSON(http.StatusOK, res)
	}
}
