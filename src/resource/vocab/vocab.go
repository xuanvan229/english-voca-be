package vocab

import (
	"english-vocab/config"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gin-gonic/gin"
	"fmt"
)

type Vocab struct {
	ID uint `gorm:"primary_key" json:"id"`
	English string `json:"english"`
	Vietnamese string `json:"vietnamese"`
}

func CreateVocab(c *gin.Context) {
	if (c.Request.Method == "POST") {
		db, err := config.Connect()
		if err != nil {
			fmt.Println("Cant not connect db", err)
		}
		
		check := db.HasTable(&Vocab{});
		if !check {
			db.CreateTable(&Vocab{})
		}

		vocab := Vocab{}
		if err := c.ShouldBind(&vocab); err != nil {
			c.JSON(503, gin.H{"error": err.Error()})
			return
		}
		db.Create(&vocab)
		c.JSON(200, vocab)
		return
	} else {
		c.JSON(404, gin.H{"error": "error"})
	}

}


func GetVocabs(c *gin.Context) {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Cant not connect db", err)
	}
	vocabs := []Vocab{}
	db.Find(&vocabs)
	c.JSON(200, vocabs)
	return 
}

func DeleteVocab(c *gin.Context) {
	if (c.Request.Method == "DELETE") {
		db, err := config.Connect()
		if err != nil {
			fmt.Println("Cant not connect db", err)
		}
		id := c.Param("id")
		fmt.Println(id)
		db.Delete(&Vocab{}, id)
		vocabs := []Vocab{}
		db.Find(&vocabs)
		c.JSON(200, vocabs)
		return 
	} else {
		c.JSON(404, gin.H{"error": "error"})
	}
}