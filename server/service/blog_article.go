package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

func AddArticle(article *model.BlogArticle, tags []uint) (err error) {
	db := global.GVA_DB
	var blogTags []model.BlogTag
	db.Where("id in ?", tags).Find(&blogTags)
	err = db.Create(article).Association("Tags").Append(blogTags)
	return err
}

func GetArticle(id int) (article model.BlogArticle, err error) {
	if errors.Is(global.GVA_DB.Where("id = ?", id).Preload("Tags").Preload("Category").First(&article).Error, gorm.ErrRecordNotFound) {
		return model.BlogArticle{}, errors.New("未查询到该文章")
	}
	return article, nil
}

func GetArticleList(param *request.ArticleList) (err error, count int64, list []model.BlogArticle) {
	offset := (param.PageNo - 1) * param.PageSize
	limit := param.PageSize
	db := global.GVA_DB
	if len(param.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+param.Keyword+"%")
	}
	if param.Category > 0 {
		db = db.Where("category_id = ?", param.Category)
	}
	if len(param.Tags) > 0 {
		for _, v := range param.Tags {
			s := strconv.Itoa(v)
			db = db.Where("tag_id_group LIKE ?", "%"+s+"%")
		}
	}
	if param.OrderType == 1 {
		db = db.Order("created_at desc")
	} else if param.OrderType == 2 {
		db = db.Order("likes desc")
	}
	err = db.Model(&model.BlogArticle{}).Count(&count).Error
	var columns = []string{
		"id", "title", "thumb", "likes", "view", "category_id", "created_at", "updated_at",
	}
	if param.WithContent {
		columns = append(columns, "content")
	}
	err = db.Select(columns).Limit(limit).Offset(offset).Preload("Category").Preload("Tags").Find(&list).Error
	return err, count, list
}

func UpdateArticle(param *request.UpdateArticleParam) error {
	var tags []model.BlogTag
	tagIdGroup := fmt.Sprintf("%d", param.Tags)
	db := global.GVA_DB
	db.Where("id in ?", param.Tags).Find(&tags)
	article := model.BlogArticle{}
	article.ID = param.ID
	article.Title = param.Title
	article.Content = param.Content
	article.CategoryId = param.Category
	article.TagIdGroup = tagIdGroup
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Updates(&article).Error
		if err != nil {
			return err
		}
		err = tx.Model(&article).Association("Tags").Replace(&tags)
		if err != nil {
			return err
		}
		return nil
	})
}

func LikeArticle(id int) error {
	var article model.BlogArticle
	db := global.GVA_DB
	err := db.Where("id = ?", id).First(&article).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("文章不存在")
	}
	article.Likes = article.Likes + 1
	return db.Model(&article).Select("likes").Updates(&article).Error
}

func DelArticle(id int) error {
	var article model.BlogArticle
	db := global.GVA_DB
	db = db.Where("id = ?", id).First(&article)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return errors.New("文章不存在")
	}
	return db.Where("id = ?", id).Delete(&model.BlogArticle{}).Error
}
