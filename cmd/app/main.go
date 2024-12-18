package main

import (
	"github.com/gin-gonic/gin"

	v1 "ev-pooling-test-api/api/v1"
)

func main() {
	engine := gin.Default()

	v1.SetUp(engine)

	engine.Run(":8080")
}
