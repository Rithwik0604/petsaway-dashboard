package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupHandlers() {
	Router.GET("/api/clients", handleGetClients)
	Router.POST("/api/clients", handleAddClient)

	Router.GET("/api/clients/:id", handleGetClientId)
	Router.PATCH("/api/clients/:id")
}

func handleGetClients(c *gin.Context) {
	// get clients
	// ...

	c.JSON(http.StatusOK, nil)
}

func handleAddClient(c *gin.Context) {
	// add client
	// ...

	c.JSON(http.StatusCreated, nil)
}

func handleGetClientId(c *gin.Context) {
	id := c.Param("id")
	_ = id

	// get client by ID
	// ...

	c.JSON(http.StatusOK, nil)
}

func handlePatchClient(c *gin.Context) {
	id := c.Param("id")
	_ = id

	// get client and patch
	// ...

	c.JSON(http.StatusOK, nil)
}
