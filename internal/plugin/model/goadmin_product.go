package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetGoadminProductTable(ctx *context.Context) table.Table {

	goadminProduct := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := goadminProduct.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Desc", "desc", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("goadmin_product").SetTitle("商品管理").SetDescription("GoadminProduct")

	formList := goadminProduct.GetForm()
	//formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Desc", "desc", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("goadmin_product").SetTitle("商品管理").SetDescription("GoadminProduct")

	return goadminProduct
}
