// @Author: 2014BDuck
// @Date: 2020/9/26

package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")

	ErrorGetArticleFail    = NewError(2002001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(2002002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(2002003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(2002004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(2002005, "删除文章失败")
)
