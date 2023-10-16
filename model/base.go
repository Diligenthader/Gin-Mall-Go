package model

type BasePage struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
	//表示一个映射到表单数据上
}
