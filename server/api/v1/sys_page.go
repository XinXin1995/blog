package v1

import (
	"blog/global/response"
	"blog/model"
	"blog/model/request"
	resp "blog/model/response"
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddPage(c *gin.Context) {
	var page request.SysAddPage
	err := c.ShouldBindJSON(&page)
	if err != nil {
		response.FailWidthError("字段错误", err, c)
		return
	}
	P := &model.SysPage{
		Name:          page.Name,
		Path:          page.Path,
		RedirectPath:  page.RedirectPath,
		Exec:          page.Exec,
		ComponentPath: page.ComponentPath,
		Button:        page.Button,
		ParentId:      page.ParentId,
	}
	_, err = service.AddPage(P)
	if err != nil {
		response.FailWidthMessage(fmt.Sprintf("添加页面失败：%v", err.Error()), c)
	} else {
		response.OkWithNil(c)
	}
}

func UpdatePage(c *gin.Context) {
	var page request.SysUpdatePage
	err := c.ShouldBindJSON(&page)
	if err != nil {
		response.FailWidthError("字段错误", err, c)
		return
	}
	P := &model.SysPage{
		Id:            page.Id,
		Name:          page.Name,
		Path:          page.Path,
		RedirectPath:  page.RedirectPath,
		Exec:          page.Exec,
		ComponentPath: page.ComponentPath,
		Button:        page.Button,
		ParentId:      page.ParentId,
	}
	_, err = service.UpdatePage(P)
	if err != nil {
		response.FailWidthMessage(fmt.Sprintf("修改页面失败：%v", err.Error()), c)
	} else {
		response.OkWithNil(c)
	}
}

func DeletePage(c *gin.Context) {
	var param request.SysDelPage
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWidthError("参数错误：", err, c)
		return
	}
	if _, err := service.DeletePage(param.Ids); err != nil {
		response.FailWidthMessage(fmt.Sprintf("删除页面失败：%v", err.Error()), c)
	} else {
		response.OkWithNil(c)
	}
}

func PageTree(c *gin.Context) {
	pages, err := service.AllPages()
	if err != nil {
		response.FailWidthError("页面获取失败", err, c)
		return
	}
	formatPages(pages, c)
}

func formatPages(pages []model.SysPage, c *gin.Context) {
	var parentPage []*resp.PageTree
	for i := 0; i < len(pages); {
		page := pages[i]
		if page.ParentId == 0 {
			P := resp.PageTree{
				Id:            page.Id,
				Name:          page.Name,
				Path:          page.Path,
				RedirectPath:  page.RedirectPath,
				Exec:          page.Exec,
				ComponentPath: page.ComponentPath,
				ParentId:      page.ParentId,
				Children:      nil,
			}
			parentPage = append(parentPage, &P)
			pages = append(pages[:i], pages[i+1:]...)
		} else {
			i++
		}
	}
	for _, parent := range parentPage {
		if len(pages) > 0 {
			formatPageChildren(&pages, parent)
			fmt.Println("pages = ", pages)
			fmt.Println("parent.Children = ", parent.Children)
		}

	}
	response.OkWithData(parentPage, c)
}

func formatPageChildren(pages *[]model.SysPage, parent *resp.PageTree) {
	for i := 0; i < len(*pages); {
		page := (*pages)[i]
		if page.ParentId == parent.Id {
			P := resp.PageTree{
				Id:            page.Id,
				Name:          page.Name,
				Path:          page.Path,
				RedirectPath:  page.RedirectPath,
				Exec:          page.Exec,
				ComponentPath: page.ComponentPath,
				ParentId:      page.ParentId,
				Children:      nil,
			}
			parent.Children = append(parent.Children, &P)
			*pages = append((*pages)[:i], (*pages)[i+1:]...)
			if has(pages, P) {
				formatPageChildren(pages, &P)
			}
		} else {
			i++
		}
	}
}

func has(pages *[]model.SysPage, page resp.PageTree) bool {
	has := false
	for _, p := range *pages {
		if p.ParentId == page.Id {
			has = true
			break
		}
	}
	return has
}
