package service

import (
	"blog/global"
	"blog/model"
	"blog/model/response"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func GetCategories() ([]model.BlogCategory, error) {
	var categories []model.BlogCategory
	err := global.GVA_DB.Find(&categories).Error
	return categories, err
}
func GetCategoriesStatistic() (categories []response.StatisticCategory, err error) {
	sqlStr := `SELECT c.id, c.name, c.type, COUNT(a.id) AS count FROM blog_categories c LEFT JOIN blog_articles a on  a.category_id = c.id WHERE c.type=2 GROUP BY c.id`
	fmt.Println("111111111111111111111111")
	rows, err := global.GVA_DB.Debug().Model(&model.BlogCategory{}).Raw(sqlStr).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var cate response.StatisticCategory
		err := global.GVA_DB.ScanRows(rows, &cate)
		if err != nil {
			break
		}
		categories = append(categories, cate)
	}
	return
}

func CreateCategory(category *model.BlogCategory) error {
	return global.GVA_DB.Create(category).Error
}

func UpdateCategory(category *model.BlogCategory) error {
	return global.GVA_DB.Model(category).Updates(category).Error
}

func DeleteCategory(id int) error {
	if !errors.Is(global.GVA_DB.Where("category_id = ?", id).First(&model.BlogArticle{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("有文章关联此分类，不能删除")
	}
	return global.GVA_DB.Unscoped().Model(&model.BlogCategory{}).Delete("id = ?", id).Error
}
