package models

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//func (article *Article) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//
//	return nil
//}
//
//func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//
//	return nil
//}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func ExistArticleByName(name string) bool {
	var article Article
	db.Select("id").Where("name = ?", name).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

// 获取所有Article文章数量
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

// 获取所有文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

// 获取指定id文章
func GetArticle(id int) (article Article) {
	db.Where("id = ? AND deleted_on = ?", id, 0).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

// 新增文章
func AddArticle(data map[string]interface{}) bool {
	log.Printf("%#v\n", data)
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}


// 修改指定id文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)

	return true
}


// 删除指定id文章
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}