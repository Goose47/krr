// Package controllers provides transport layer logic for application endpoints.
package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type RecsController struct {
	log         *slog.Logger
	recommender Recommender
}

type Recommender interface {
	Recommend(ctx context.Context) (string, error)
}

// NewRecsController is a constructor for RecsController.
func NewRecsController(
	log *slog.Logger,
	recommender Recommender,
) *RecsController {
	return &RecsController{
		log:         log,
		recommender: recommender,
	}
}

func (con *RecsController) Recommend(c *gin.Context) {
	res, err := con.recommender.Recommend(context.TODO())

	if err != nil {
		c.JSON(500, fmt.Sprintf("failed to retrieve recommendations: %s. %v", err.Error(), err))
		return
	}

	c.JSON(200, res)
}
