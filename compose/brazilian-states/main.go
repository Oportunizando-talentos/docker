package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type State struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Abbrev     string `json:"abbreviation"`
	Population int    `json:"population"`
}

func main() {
	dsn := os.Getenv("DSN")

	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/states/:id", func(c *gin.Context) {
		stateID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state ID"})
			return
		}

		var state State
		err = db.QueryRow(context.Background(), "SELECT id, name, abbreviation, population FROM states WHERE id = $1", stateID).Scan(&state.ID, &state.Name, &state.Abbrev, &state.Population)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "State not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve state"})
			}
			return
		}

		c.JSON(http.StatusOK, state)
	})

	fmt.Println("API server listening on port 8080")
	router.Run(":8080")
}
