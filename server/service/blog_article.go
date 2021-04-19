package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
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
	err = db.Select("id", "title", "thumb", "likes", "view", "category_id", "created_at", "updated_at").Limit(limit).Offset(offset).Preload("Category").Preload("Tags").Order("created_at desc").Find(&list).Error
	return err, count, list
}
func articleCount(param *request.ArticleList) int64 {
	var strBuilder strings.Builder
	strBuilder.WriteString("select COUNT(1) as total FROM ( ")
	strBuilder.WriteString("SELECT ")
	strBuilder.WriteString("a.id")
	strBuilder.WriteString("group_concat(t.id) as group_tag_id, ")
	strBuilder.WriteString("FROM ")
	strBuilder.WriteString("blog_articles a ")
	strBuilder.WriteString("LEFT JOIN blog_categories c on a.category_id = c.id ")
	strBuilder.WriteString("LEFT JOIN article_tags r ON r.blog_article_id = a.id ")
	strBuilder.WriteString("LEFT JOIN blog_tags t ON r.blog_tag_id = t.id ")
	strBuilder.WriteString("GROUP BY a.id ")
	strBuilder.WriteString(") a ")
	strBuilder.WriteString("WHERE a.deleted_at IS NULL  ")
	if len(param.Keyword) > 0 {
		strBuilder.WriteString(fmt.Sprintf("AND a.title LIKE '%%%s%%' ", param.Keyword))
	}
	if param.Category > 0 {
		strBuilder.WriteString(fmt.Sprintf("AND a.category_id = %d ", param.Category))
	}
	if len(param.Tags) > 0 {
		for _, tagId := range param.Tags {
			strBuilder.WriteString(fmt.Sprintf(" AND group_tag_id LIKE '%%%d%%' ", tagId))
		}
	}
	strSql := strBuilder.String()
	row := global.GVA_DB.Raw(strSql).Row()
	var total int64
	row.Scan(&total)
	return total
}

func formatArticleList(articles []model.Articles) []model.Articles {
	var tags []model.BlogTag
	global.GVA_DB.Find(&tags)
	var tagsTree = make(map[uint]model.BlogTag)
	if len(tags) > 0 {
		for _, tag := range tags {
			tagsTree[tag.ID] = tag
		}
	}
	for i, v := range articles {
		groupTagIds := strings.Split(v.GroupTagId, ",")
		for _, tagId := range groupTagIds {
			id, _ := strconv.Atoi(tagId)
			articles[i].Tags = append(articles[i].Tags, tagsTree[uint(id)])
		}
	}
	return articles
	//if len(*articles) == 0 {
	//	return nil
	//}
	//resultMap := make(map[uint]response.RespArticle)
	//for _, v := range *articles {
	//	if _, ok := resultMap[v.ID]; ok {
	//		article := resultMap[v.ID]
	//		article.Tags = append(article.Tags, response.RespArticleTag{v.TagId, v.TagName, v.TagColor})
	//	} else {
	//		article := response.RespArticle{}
	//		article.ID = v.ID
	//		article.Title = v.Title
	//		article.Content = v.Content
	//		article.Thumb = v.Thumb
	//		article.CreatedAt = v.CreatedAt
	//		article.UpdatedAt = v.UpdatedAt
	//		article.Category = response.RespArticleCategory{}
	//		article.Category.ID = v.CategoryId
	//		article.Category.Name = v.CategoryName
	//		article.Tags = []response.RespArticleTag{}
	//		article.Tags = append(article.Tags, response.RespArticleTag{v.TagId, v.TagName, v.TagColor})
	//		resultMap[v.ID] = article
	//	}
	//}
	//var respArticleList []response.RespArticle
	//for _, v := range resultMap {
	//	respArticleList = append(respArticleList, v)
	//}
	//return respArticleList
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
