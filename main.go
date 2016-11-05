package main

import (
	"github.com/gin-gonic/gin"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/stevenmhernandez/the-group-cmsc355-api-server/database"
	"github.com/stevenmhernandez/the-group-cmsc355-api-server/models"
	"github.com/stevenmhernandez/the-group-cmsc355-api-server/utils"
	"os"
)

var dbmap gorp.DbMap = database.Init()

func main() {
	dbmap.AddTableWithName(models.Highscore{}, "highscores").SetKeys(true, "Id")

	router := gin.Default()

	router.GET("/highscores", HighScoreList)
	router.POST("/highscores", HighScorePost)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router.Run(":" + port)
}

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

func HighScoreList(c *gin.Context) {
	var highScores []models.Highscore
	_, err := dbmap.Select(&highScores, "select * from highscores order by score DESC limit 3")
	utils.LogError(err, "Select failed:")
	c.JSON(200, highScores)
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
	var highscore models.Highscore
	_ = dbmap.SelectOne(&highscore, "select * from highscores where username=?", username)

	if highscore.Id != 0 {
		if (score > highscore.Score) {
			highscore.Score = score;
			_, err := dbmap.Update(&highscore)
			utils.LogError(err, "Update failed:")
		}
	} else {
		model := models.Model{
			CreatedAt:    time.Now().UnixNano(),
		}

		highscore = models.Highscore{
			Model:           model,
			Username:        username,
			Score:           score,
		}
		err := dbmap.Insert(&highscore)
		utils.LogError(err, "Insert failed:")
	}

	return highscore
}
