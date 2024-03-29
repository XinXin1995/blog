package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

	if len(param.Tags) > 0 {
		var articleIds []uint
		rows, err := db.Raw(" select DISTINCT(blog_article_id) from article_tags where blog_tag_id in ?", param.Tags).Rows()
		if err != nil {
			global.GVA_LOG.Error("get article data column(tags) err", zap.Any("err", err))
		}
		for rows.Next() {
			var id uint
			_ = rows.Scan(&id)
			articleIds = append(articleIds, id)
		}
		db = db.Where("id in ?", articleIds)
	}
	if len(param.Keyword) > 0 {
		db = db.Where("title LIKE ?", "%"+param.Keyword+"%")
	}
	if param.Category > 0 {
		db = db.Where("category_id = ?", param.Category)
	}
	if param.OrderType == 1 {
		db = db.Order("created_at desc")
	} else if param.OrderType == 2 {
		db = db.Order("likes desc")
	} else if param.OrderType == 3 {
		db = db.Order("top desc")
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

func SetTopArticle(id int) error {
	var article model.BlogArticle
	if errors.Is(global.GVA_DB.Where("id = ?", id).First(&article).Error, gorm.ErrRecordNotFound) {
		return errors.New("为找到相关文章")
	}
	var maxTop int
	row := global.GVA_DB.Raw("SELECT MAX(top) FROM blog_articles").Row()
	if err := row.Scan(&maxTop);err != nil {
		return err
	}
	article.Top = maxTop + 1
	return global.GVA_DB.Select("Top").Updates(&article).Error

}
