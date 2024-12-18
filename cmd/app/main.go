package main

import (
	v1 "ev-pooling-test-api/internal/infrastructure/api/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	v1.SetUp(engine)

	engine.Run(":8080")
}
