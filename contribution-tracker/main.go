package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type Contribution struct {
	ID          uint      `json:"id"`
	Repository  string    `json:"respository"`
	Contributor string    `json:"contributor"`
	commits     int       `json:"commits"`
	Date        time.Time `json:"date"`
}

func main() {
	db, err := gorm.Open("sqlite3", "contribution.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Contribution{})

	r := gin.Default()
	r.GET("/contributions", getContributions)
	r.POST("/contributions", createContribution)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func getContributions(c *gin.Context) {
	var contributions []Contribution
	if err := db.Find(&contributions).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, &contributions)
	}
}

func createContribution(c *gin.Context) {
	var contribution Contribution
	c.BindJSON(&contribution)
	contribution.Date = time.Now()

	db.Create(contribution)
	c.JSON(200, contribution)

}
