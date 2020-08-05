package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetGoadminTracesTable(ctx *context.Context) table.Table {
	// 获取当前厂商ID
	user := auth.Auth(ctx)
	companyUser, err := db.WithDriver(globalConnection).Table("goadmin_company_users").Select("company_id").Where("user_id", "=", user.Id).First()
	if err != nil{
		panic("获取用户厂商编码失败.")
	}
	companyId := companyUser["company_id"].(int64)

	goadminTraces := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := goadminTraces.GetInfo().HideFilterArea()
	info.Where("cid", "=", companyId)

	info.AddField("ID", "id", db.Int).FieldFilterable()
	info.AddField("订单标识", "order_identifier", db.Varchar)
	info.AddField("截止日期", "ending_date", db.Varchar)
	info.AddField("企业编号", "cid", db.Int)
	info.AddField("公司名称", "company_name", db.Varchar)
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	//info.AddField("Deleted_at", "deleted_at", db.Timestamp)
	info.AddField("商品编号", "pid", db.Int).FieldFilterable()
	info.AddField("商品名称", "product_name", db.Varchar).FieldFilterable()
	info.AddActionButton("详细", action.Jump("/admin/info/goadmin_traces_infos?trace_id={{.Id}}"))
	info.SetTable("goadmin_traces").SetTitle("溯源").SetDescription("溯源事件")


	formList := goadminTraces.GetForm()

	pid := ctx.Query("pid")
	//formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("订单标识", "order_identifier", db.Varchar, form.Text).FieldMust()
	formList.AddField("截止日期", "ending_date", db.Varchar, form.Text).FieldMust()
	formList.AddField("Pid", "pid", db.Int, form.Number).FieldDefault(pid).FieldHide()
	//formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	//formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)
	//formList.AddField("Deleted_at", "deleted_at", db.Timestamp, form.Datetime)
	//formList.AddField("Cid", "cid", db.Int, form.Number).FieldHide().FieldDefault(string(companyId))
	//formList.AddField("Company_name", "company_name", db.Varchar, form.Text).FieldHide().FieldDefault(companyName)

	//formList.AddField("Product_name", "product_name", db.Varchar, form.Text).FieldHide().FieldDefault("header")

	formList.SetTable("goadmin_traces").SetTitle("溯源").SetDescription("溯源事件")

	formList.SetInsertFn(func(values form2.Values) error{
		company, err := db.WithDriver(globalConnection).Table("goadmin_companys").Select("name").Where("id", "=", companyId).First()
		if err != nil{
			panic("获取厂商失败.")
		}
		companyName := company["name"].(string)
		pid := values.Get("pid")

		product, err := db.WithDriver(globalConnection).Table("goadmin_company_products").Select("name").Where("id", "=", pid).Where("company_id", "=", companyId).First()
		if err != nil{
			panic("获取商品失败.")
		}
		productName := product["name"].(string)

		// 添加溯源
		_, addProductErr := db.WithDriver(globalConnection).Table("goadmin_traces").
			Insert(dialect.H{
				"cid": companyId,
				"company_name" : companyName,
				"pid" : pid,
				"product_name" : productName,
				"order_identifier": values.Get("order_identifier"),
				"ending_date": values.Get("ending_date"),
			})
		if addProductErr != nil {
			return addProductErr
		}
		return nil
	})
	return goadminTraces
}
