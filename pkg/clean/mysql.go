package clean

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"

	"gin-blog/models"
)

const OneSecond = 1*time.Second

// 清理标记删除MySQL任务数据
func CleanMySQLDeleted() {
	c := cron.New(cron.WithSeconds())
	c.Start()
	defer c.Stop()

	// 每天清理标记删除的Tag
	c.AddFunc("*/5 * * * * ?", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})

	// 每天清理标记删除的Article
	c.AddFunc("*/5 * * * * ?", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	for {
		select {
		case <-time.After(OneSecond):
		}
	}
}