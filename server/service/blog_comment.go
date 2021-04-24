package service

import (
	"blog/global"
	"blog/model"
	"blog/model/request"
	"strconv"
)

func AddComment(comment *model.BlogComment) error {
	err := global.GVA_DB.Create(comment).Error
	return err
}

func GetCommentTree(param request.TreeComment) (err error, count int64, comments []model.BlogComment) {
	limit := param.PageSize
	offset := (param.PageNo - 1) * param.PageSize
	err = global.GVA_DB.Model(&model.BlogComment{}).Where("parent_id = 0 AND article_id = ?", param.ArticleId).Count(&count).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Limit(limit).Offset(offset).Where("parent_id = 0 AND article_id = ?", param.ArticleId).Order("created_at desc").Find(&comments).Error
	if err != nil {
		return
	}
	var allChildrenComments []model.BlogComment
	err = global.GVA_DB.Where("parent_id != 0 AND article_id = ?", param.ArticleId).Find(&allChildrenComments).Error
	if err != nil {
		return
	}
	treeChildren := getChildrenTreeMap(allChildrenComments)
	for i := 0; i < len(comments); i++ {
		setCommentChildren(&comments[i], treeChildren)
	}
	return err, count, comments
}

func getChildrenTreeMap(childrenList []model.BlogComment) map[string][]model.BlogComment {
	var tree = make(map[string][]model.BlogComment)
	for i := 0; i < len(childrenList); i++ {
		key := strconv.Itoa(int(childrenList[i].ParentId))
		tree[key] = append(tree[key], childrenList[i])
	}
	return tree
}

func setCommentChildren(parentComment *model.BlogComment, allChildren map[string][]model.BlogComment) {
	parentId := strconv.Itoa(int(parentComment.ID))
	parentComment.Children = allChildren[parentId]
	for i := 0; i < len(parentComment.Children); i++ {
		setCommentChildren(&parentComment.Children[i], allChildren)
	}
}
