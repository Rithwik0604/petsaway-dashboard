package server

import (
	"database/sql"
	"net/http"
	"petsaway/internal/auth"
	"petsaway/internal/clients"

	"github.com/gin-gonic/gin"
)

func setupHandlers() {
	// login
	Router.POST("/api/auth/login", handleLogin)
	Router.POST("/api/auth/logout", handleLogout)

	Router.Use(authMiddleware())
	Router.GET("/api/clients", handleGetClients)
	Router.POST("/api/clients", handleAddClient)

	Router.GET("/api/clients/:id", handleGetClientId)
	Router.PATCH("/api/clients/:id", handlePatchClient)
}

func handleLogin(c *gin.Context) {
	var loginDetails LoginDetails

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := auth.VerifyPassword(loginDetails.Username, loginDetails.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	auth.SetCookie(c, token)

	c.Status(http.StatusOK)
}

func handleLogout(c *gin.Context) {
	auth.InvalidateToken(c)

	c.Status(http.StatusOK)
}

func handleGetClients(c *gin.Context) {
	// get clients
	clients, err := clients.GetAllClients()
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "No clients",
			})
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusOK, clients)
}

func handleAddClient(c *gin.Context) {
	// add client
	var clientDetails clients.ClientDTO

	if err := c.ShouldBindJSON(&clientDetails); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	client, err := clients.AddClient(clientDetails)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, client)
}

func handleGetClientId(c *gin.Context) {
	id := c.Param("id")

	// get client by ID
	client, err := clients.GetClientById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, client)
}

func handlePatchClient(c *gin.Context) {
	id := c.Param("id")

	var clientDetails clients.ClientDTO

	if err := c.ShouldBindJSON(&clientDetails); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	client, err := clients.UpdateClient(id, clientDetails)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, client)
}
