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
	// Fitness route
	r.POST("/fitness", func(c *gin.Context) {
		// Bind scenario and gene
		type body struct {
			Scenario ai.Scenario `json:"scenario"`
			Gene     string      `json:"gene"`
		}
		b := body{}
		err := c.BindJSON(&b)
		// Assert body
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}

		// Bind gene buffer
		geneBuffer, err := base64.StdEncoding.DecodeString(b.Gene)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad gene descriptor encoding (must be base64)"})
			return
		}

		// Decode into gene
		gene := ai.Gene{}
		gene.Decode(geneBuffer)
		fmt.Println(gene)

		// Show result
		c.JSON(http.StatusOK, gin.H{"fitness": ai.Stats(gene).Fitness(b.Scenario)})
	})

	// Process route
	r.POST("/process", func(c *gin.Context) {
		// Bind scenario
		type body struct {
			Scenario ai.Scenario `json:"scenario"`
		}
		b := body{}
		err := c.BindJSON(&b)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not read scenario from body"})
		}

		// Show result
		best, fitness := b.Scenario.Process(1000, 20000, 4, 16, .1)
		bestBuffer := best.Encode()
		bestBase64 := base64.StdEncoding.EncodeToString(bestBuffer[:])
		c.JSON(http.StatusOK, gin.H{"best": bestBase64, "fitness": fitness})
	})
}
