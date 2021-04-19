package service

import (
	"blog/global"
	"blog/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func GetTags() ([]model.BlogTag, error) {
	var tags []model.BlogTag
	err := global.GVA_DB.Find(&tags).Error
	return tags, err
}

func CreateTag(tag *model.BlogTag) error {
	return global.GVA_DB.Create(tag).Error
}

func UpdateTag(tag *model.BlogTag) error {
	return global.GVA_DB.Model(tag).Updates(tag).Error
}

func DeleteTag(id int) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Exec("Delete FROM article_tags WHERE blog_tag_id = ?", id).Error
		if err != nil {
			return errors.New("删除与文章关联关系失败")
		}
		err = tx.Unscoped().Model(&model.BlogTag{}).Delete("id = ?", id).Error
		if err != nil {
			return errors.New("删除标签失败")
		}
		str := "UPDATE blog_articles a INNER JOIN (select id, SUBSTRING_INDEX(tag_id_group, '"+strconv.Itoa(id)+"', 1) as tag_id_group from blog_articles WHERE tag_id_group LIKE '%"+strconv.Itoa(id)+"%'\n) r ON a.id = r.id SET a.tag_id_group = r.tag_id_group"
		err = tx.Exec(str).Error
		if err != nil {
			return errors.New("删除标签失败")
		}
		return nil
	})
}
