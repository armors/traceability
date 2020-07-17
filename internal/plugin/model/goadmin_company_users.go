package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetGoadminCompanyUsersTable(ctx *context.Context) table.Table {

	goadminCompanyUsers := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := goadminCompanyUsers.GetInfo().HideFilterArea()

	info.AddField("Company_id", "company_id", db.Int)
	info.AddField("User_id", "user_id", db.Int)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("goadmin_company_users").SetTitle("GoadminCompanyUsers").SetDescription("GoadminCompanyUsers")

	formList := goadminCompanyUsers.GetForm()
	formList.AddField("Company_id", "company_id", db.Int, form.Number)
	formList.AddField("User_id", "user_id", db.Int, form.Number)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("goadmin_company_users").SetTitle("GoadminCompanyUsers").SetDescription("GoadminCompanyUsers")

	return goadminCompanyUsers
}
