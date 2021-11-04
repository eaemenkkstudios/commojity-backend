package api

import (
	"net/http"

	"github.com/eaemenkkstudios/commojity-backend/ai"
	"github.com/eaemenkkstudios/commojity-backend/data"
	"github.com/gin-gonic/gin"
)

// Register API routes
func RegisterRoutes(r *gin.Engine) {
	r.GET("/gene", func(c *gin.Context) {
		encodedGene := c.Request.URL.Query()["g"]
		if len(encodedGene) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "gene emoji code not present"})
		}
		geneData := data.GeneData{}
		if err := geneData.Decode(encodedGene[0]); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			gene := ai.Gene{}
			gene.Decode(geneData)
			c.JSON(http.StatusOK, gin.H{"buffer": geneData, "gene": gene.Stats()})
		}
	})

	r.POST("/start", func(c *gin.Context) {

	})
}
