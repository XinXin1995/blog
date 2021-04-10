package service

import (
	"blog/global"
	"blog/model"
	"errors"
	"gorm.io/gorm"
)

func GetCategories() ([]model.BlogCategory, error) {
	var categories []model.BlogCategory
	err := global.GVA_DB.Find(&categories).Error
	return categories, err
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
