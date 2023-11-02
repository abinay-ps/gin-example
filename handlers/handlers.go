package handlers

import (
	"database/sql"
	"net/http"

	"github.com/abinay-ps/gin-example/models"
	"github.com/gin-gonic/gin"
)

func Iniatilize_routes(router *gin.Engine, db *sql.DB) {
	router.GET("/movies", getMovieDetails(db))
	router.POST("/movies", createMovie(db))
	router.GET("/movies/:id", getMovieById(db))
	router.PUT("/movies/:id", updateMovieById(db))
	router.DELETE("/movies/:id", deleteMovieById(db))
}

func getMovieDetails(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query_Param := c.Request.URL.Query()

		if len(query_Param) == 0 {
			query := "Select * from movies"

			rows, err := db.Query(query)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			var movielist []models.Movie
			for rows.Next() {
				var movie models.Movie
				if err := rows.Scan(&movie.ID, &movie.Title, &movie.Director); err != nil {
					c.JSON(http.StatusInternalServerError, err.Error())
				} else {
					movielist = append(movielist, movie)
				}
			}
			c.JSON(http.StatusOK, movielist)
		} else {
			for key, values := range query_Param {
				for _, value := range values {
					switch key {
					case "id":
						query := "Select * from movies where id=$1"
						rows, err := db.Query(query, value)
						if err != nil {
							c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
							return
						}
						var movielist []models.Movie
						for rows.Next() {
							var movie models.Movie
							if err := rows.Scan(&movie.ID, &movie.Title, &movie.Director); err != nil {
								c.JSON(http.StatusInternalServerError, err.Error())
							} else {
								movielist = append(movielist, movie)
							}
						}
						c.JSON(http.StatusOK, movielist)
					case "title":
						query := "Select * from movies where title=$1"
						rows, err := db.Query(query, value)
						if err != nil {
							c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
							return
						}
						var movielist []models.Movie
						for rows.Next() {
							var movie models.Movie
							if err := rows.Scan(&movie.ID, &movie.Title, &movie.Director); err != nil {
								c.JSON(http.StatusInternalServerError, err.Error())
							} else {
								movielist = append(movielist, movie)
							}
						}
						c.JSON(http.StatusOK, movielist)
					case "director":
						query := "Select * from movies where director=$1"
						rows, err := db.Query(query, value)
						if err != nil {
							c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
							return
						}
						var movielist []models.Movie
						for rows.Next() {
							var movie models.Movie
							if err := rows.Scan(&movie.ID, &movie.Title, &movie.Director); err != nil {
								c.JSON(http.StatusInternalServerError, err.Error())
							} else {
								movielist = append(movielist, movie)
							}
						}
						c.JSON(http.StatusOK, movielist)
					default:
						c.JSON(400, gin.H{"message": "No query parameters"})
					}
				}
			}
		}
	}
}

func createMovie(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie models.Movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		query := "Insert into movies (id,title,director) values($1,$2,$3)"

		res, err := db.Exec(query, &movie.ID, &movie.Title, &movie.Director)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else if rowsaffected, _ := res.RowsAffected(); rowsaffected < 1 {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "No row inserted"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Row is inserted successfully"})
	}

}

func getMovieById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		query := "Select * from movies where id=$1"
		var movie models.Movie
		err := db.QueryRow(query, id).Scan(&movie.ID, &movie.Title, &movie.Director)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "No rows returned"})
			return
		}
		c.JSON(http.StatusOK, movie)
	}
}

func updateMovieById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var movie models.Movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		query := "Update movies set title=$1,director=$2 where id=$3"

		res, err := db.Exec(query, &movie.Title, &movie.Director, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else if rowsaffected, _ := res.RowsAffected(); rowsaffected < 1 {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "No row updated"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Row is updated successfully"})
	}
}

func deleteMovieById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		query := "Delete from movies where id=$1"

		res, err := db.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else if rowsaffected, _ := res.RowsAffected(); rowsaffected < 1 {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "No row deleted"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Row is deleted successfully"})
	}
}
