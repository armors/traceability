package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
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
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	//info.AddActionButton("溯源", action.Jump("/admin/info/goadmin_traces?pid={{.Id}}"))
	info.AddColumnButtons("溯源", types.GetColumnButton("阶段", "", action.Jump("/admin/info/goadmin_traces?pid={{.Id}}")))
	info.SetTable("goadmin_company_products").SetTitle("商品管理").SetDescription("")

	formList := goadminCompanyProducts.GetForm()

	formList.AddField("商品名称", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("商品描述", "desc", db.Varchar, form.Text).FieldMust()

	formList.SetTable("goadmin_company_products").SetTitle("商品管理").SetDescription("")

	formList.SetInsertFn(func(values form2.Values) error{
		// 厂商的商品
		_, addProductErr := db.WithDriver(globalConnection).Table("goadmin_company_products").
			Insert(dialect.H{
				"company_id": companyId,
				"name": values.Get("name"),
				"desc": values.Get("desc"),
			})
		if addProductErr != nil {
			return addProductErr
		}
		return nil
	})

	return goadminCompanyProducts
}
