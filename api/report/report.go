package report

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikeblum/redispwned/api"
)

type Report struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}

func Routes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/servers-by-country", serversByCountry)
	v1.GET("/servers-by-version", serversByVersion)
}

func serversByCountry(c *gin.Context) {
	idx := api.NewSearchEngine()
	results, err := idx.ServersByCountry()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
	}
	report := &Report{}
	for _, result := range results {
		report.Labels = append(report.Labels, result.Value)
		report.Data = append(report.Data, result.Count)
	}
	c.JSON(http.StatusOK, report)
}

func serversByVersion(c *gin.Context) {
	idx := api.NewSearchEngine()
	results, err := idx.ServersByVersion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
	}
	report := &Report{}
	for _, result := range results {
		report.Labels = append(report.Labels, result.Value)
		report.Data = append(report.Data, result.Count)
	}
	c.JSON(http.StatusOK, report)
}
