// Package controllers provides transport layer logic for application endpoints.
package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
)

// TestController deals with calculation endpoints.
type TestController struct {
	log        *slog.Logger
	testulator Testulator
}

type Testulator interface {
	Testulate(ctx context.Context) (string, error)
}

// NewTestController is a constructor for TestController.
func NewTestController(
	log *slog.Logger,
	testulator Testulator,
) *TestController {
	return &TestController{
		log:        log,
		testulator: testulator,
	}
}

// Testulate validates request params, calculates params and composes result message.
func (con *TestController) Testulate(c *gin.Context) {
	res, _ := con.testulator.Testulate(context.TODO())
	c.JSON(200, res)
}
