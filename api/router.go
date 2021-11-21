package api

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/eaemenkkstudios/commojity-backend/ai"
	"github.com/gin-gonic/gin"
)

// Register API routes
func RegisterRoutes(r *gin.Engine) {
	r.GET("/:gene", func(c *gin.Context) {
		// Assert param
		geneString := c.Param("gene")
		if geneString == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "gene descriptor not present"})
		}

		// Decode into gene
		gene := ai.Gene{}
		geneBuffer, err := base64.StdEncoding.DecodeString(geneString)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad gene descriptor encoding (must be base64)"})
		} else {
			gene.Decode(geneBuffer)
			// Show result
			c.JSON(http.StatusOK, gin.H{"gene": ai.Stats(gene)})
		}
	})
}
