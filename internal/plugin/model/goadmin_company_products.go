package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	_ "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"strconv"
	_ "strings"
)

func GetGoadminCompanyProductsTable(ctx *context.Context) table.Table {
	// 获取当前厂商ID
	user := auth.Auth(ctx)
	company, err := db.WithDriver(globalConnection).Table("goadmin_company_users").Select("company_id").Where("user_id", "=", user.Id).First()
	if err != nil{
		if err.Error() == "out of index" {
			panic("当前登录用户没有对应厂商")
		}else {
			panic(err)
		}
	}
	companyId := company["company_id"].(int64)

	// 模型
	goadminCompanyProducts := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))
	info := goadminCompanyProducts.GetInfo().HideFilterArea()
	info.Where("company_id", "=", companyId)
	info.AddField("编码", "id", db.Int).FieldFilterable()
	info.AddField("商品名称", "name", db.Varchar)
	info.AddField("商品描述", "desc", db.Varchar)
	//info.AddField("商品图", "picture",
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	//info.AddActionButton("溯源", action.Jump("/admin/info/goadmin_traces?pid={{.Id}}"))
	info.AddColumnButtons("溯源信息", types.GetColumnButton("溯源", "", action.Jump("/admin/info/goadmin_traces?pid={{.Id}}")))
	info.SetTable("goadmin_company_products").SetTitle("商品管理").SetDescription("")

	formList := goadminCompanyProducts.GetForm()
	formList.AddField("company_id", "company_id", db.Int, form.Default).FieldHide().FieldDefault(strconv.FormatInt(companyId,10))

	formList.AddField("商品名称", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("商品描述", "desc", db.Varchar, form.Text).FieldMust()
	formList.AddField("商品图片1", "pic1", db.Varchar, form.File)
	formList.AddField("商品图片2", "pic2", db.Varchar, form.File)
	formList.AddField("商品图片3", "pic3", db.Varchar, form.File)
	formList.AddField("商品图片4", "pic4", db.Varchar, form.File)
	formList.AddField("商品图片5", "pic5", db.Varchar, form.File)
	//formList.AddField("商品图片", "picture", db.Text, form.Multifile).FieldOptionExt(map[string]interface{}{
	//	"maxFileCount": 5,
	//})

	//pk := ctx.Query("__goadmin_edit_pk")

	//if pk != ""{
	//	product, err := db.WithDriver(globalConnection).Table("goadmin_company_products").Select("picture").Where("id", "=", pk).First()
	//	if err != nil{
	//		if err.Error() == "out of index" {
	//			formList.AddTable("商品图片", "setting", func(panel *types.FormPanel) {
	//				panel.AddField("上传", "picture", db.Text, form.File).FieldHideLabel()
	//			})
	//		}else{
	//			panic(err)
	//		}
	//	}
	//	pictures := product["picture"].(string)
	//	picturesArray := strings.Split(pictures, ",")
	//	for _, pic := range picturesArray {
	//		formList.AddTable("商品图片", "setting", func(panel *types.FormPanel) {
	//			panel.AddField("上传", "picture", db.Text, form.File).FieldHideLabel().FieldDefault(config.GetStore().URL(pic))
	//		})
	//	}
	//}else {
	//	formList.AddTable("商品图片", "setting", func(panel *types.FormPanel) {
	//		panel.AddField("上传", "picture", db.Text, form.File).FieldHideLabel()
	//	})
	//}
	//formList.AddTable("商品图片", "setting", func(panel *types.FormPanel) {
	//	panel.AddField("上传", "picture", db.Text, form.File).FieldHideLabel()
	//})

	formList.AddField("商品详情","content", db.Text, form.RichText)

	formList.SetTable("goadmin_company_products").SetTitle("商品管理").SetDescription("")

	//formList.SetInsertFn(func(values form2.Values) error{
	//	// 厂商的商品
	//	_, addProductErr := db.WithDriver(globalConnection).Table("goadmin_company_products").
	//		Insert(dialect.H{
	//			"company_id": companyId,
	//			"name": values.Get("name"),
	//			"desc": values.Get("desc"),
	//			"content": values.Get("content"),
	//		})
	//	if addProductErr != nil {
	//		return addProductErr
	//	}
	//	return nil
	//})

	return goadminCompanyProducts
}
