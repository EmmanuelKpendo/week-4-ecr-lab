package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var africanCountries = []string{
	"Algeria", "Angola", "Benin", "Botswana", "Burkina Faso",
	"Burundi", "Cape Verde", "Cameroon", "Central African Republic",
	"Chad", "Comoros", "Democratic Republic of the Congo",
	"Republic of the Congo", "Djibouti", "Egypt", "Equatorial Guinea",
	"Eritrea", "Eswatini", "Ethiopia", "Gabon", "Gambia", "Ghana",
	"Guinea", "Guinea-Bissau", "Ivory Coast", "Kenya", "Lesotho",
	"Liberia", "Libya", "Madagascar", "Malawi", "Mali", "Mauritania",
	"Mauritius", "Morocco", "Mozambique", "Namibia", "Niger", "Nigeria",
	"Rwanda", "Sao Tome and Principe", "Senegal", "Seychelles",
	"Sierra Leone", "Somalia", "South Africa", "South Sudan", "Sudan",
	"Tanzania", "Togo", "Tunisia", "Uganda", "Zambia", "Zimbabwe",
}

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Home page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Random African Countries API
	router.GET("/random", func(c *gin.Context) {
		country := africanCountries[rand.Intn(len(africanCountries))]
		c.JSON(http.StatusOK, gin.H{"country": country})
	})

	//Health check
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	err := router.Run(":" + port)
	fmt.Println("Application running on port: " + port)
	if err != nil {
		return
	}
}
