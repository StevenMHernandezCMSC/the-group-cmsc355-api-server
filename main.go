package main

import (
	"github.com/gin-gonic/gin"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
	"github.com/stevenmhernandez/the-group-cmsc355-api-server/database"
	"github.com/stevenmhernandez/the-group-cmsc355-api-server/models"
	"github.com/stevenmhernandez/the-group-cmsc355-api-server/utils"
)

var dbmap gorp.DbMap = database.Init()


func main() {
	dbmap.AddTableWithName(models.Highscore{}, "highscores").SetKeys(true, "Id")

	router := gin.Default()

	router.GET("/highscores", HighScoreList)
	router.POST("/highscores", HighScorePost)

	router.Run(":3000")
}

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

func HighScoreList(c *gin.Context) {
	var highScores []models.Highscore
	_, err := dbmap.Select(&highScores, "select * from highscores order by score DESC limit 3")
	utils.LogError(err, "Select failed:")
	content := gin.H{}
	for k, v := range highScores {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)
}

func HighScorePost(c *gin.Context) {
	var json models.Highscore

	c.Bind(&json)
	highScore := addHighScore(json.Username, json.Score)
	if highScore.Username == json.Username {
		c.JSON(200, gin.H{"status": "success"})
	} else {
		c.JSON(500, gin.H{"status": "error"})
	}
}

func addHighScore(username string, score int64) models.Highscore {
	model := models.Model{
		CreatedAt:    time.Now().UnixNano(),
	}

	highscore := models.Highscore{
		Model:           model,
		Username:        username,
		Score:           score,
	}

	err := dbmap.Insert(&highscore)
	utils.LogError(err, "Insert failed:")
	return highscore
}