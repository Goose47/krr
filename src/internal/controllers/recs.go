// Package controllers provides transport layer logic for application endpoints.
package controllers

import (
	"context"
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
	res, _ := con.recommender.Recommend(context.TODO())
	c.JSON(200, res)
}
