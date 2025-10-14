package handlers

import (
	"log/slog"
	"net/http"

	"educabot.com/bookshop/models"
	"educabot.com/bookshop/providers"
	"github.com/gin-gonic/gin"
)

type MetricController struct {
	metricService providers.MetricService
}

func NewMetricController(metricService providers.MetricService) *MetricController {
	return &MetricController{metricService: metricService}
}

func (mc *MetricController) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var query models.GetMetricsRequest
		err := ctx.ShouldBindQuery(&query)
		if err != nil {
			slog.ErrorContext(ctx, "GetMetrics: failed to bind query", "err", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
			return
		}

		res, err := mc.metricService.GetMetrics(ctx, query)
		if err != nil {
			slog.ErrorContext(ctx, "GetMetrics: service error", "err", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
