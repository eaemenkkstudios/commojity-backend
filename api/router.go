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
	r.GET("/:gene", func(c *gin.Context) {
		// Assert param
		geneString := c.Param("gene")
		if geneString == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "gene descriptor not present"})
		}

		// Bind gene buffer
		geneBuffer, err := base64.StdEncoding.DecodeString(geneString)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad gene descriptor encoding (must be base64)"})
		}

		// Bind scenario
		scenario := ai.Scenario{}
		err = c.BindJSON(&scenario)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not read scenario from body"})
		}

		// Decode into gene
		gene := ai.Gene{}
		gene.Decode(geneBuffer)

		// Show result
		c.JSON(http.StatusOK, gin.H{"profit": ai.Stats(gene).Fitness(scenario)})
	})

	// Process route
	r.POST("/", func(c *gin.Context) {
		// Bind scenario
		scenario := ai.Scenario{}
		err := c.BindJSON(&scenario)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not read scenario from body"})
		}

		// Show result
		best, fitness := scenario.Process(1000, 20000, 4, 16, .1)
		bestBuffer := best.Encode()
		bestBase64 := base64.StdEncoding.EncodeToString(bestBuffer[:])
		c.JSON(http.StatusOK, gin.H{"best": bestBase64, "fitness": fitness})
	})
}
