package model

import (
	"database/sql"
	"errors"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"

	"github.com/GoAdminGroup/go-admin/modules/db"
)
func GetGoadminCompanysTable(ctx *context.Context) table.Table {

	goadminCompanys := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := goadminCompanys.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("企业名称", "name", db.Varchar)
	info.AddField("企业简介", "desc", db.Varchar)
	info.AddField("企业法人", "corporate", db.Varchar)
	info.AddField("企业地址", "address", db.Varchar)
	info.AddField("联系方式", "contact", db.Varchar)
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)

	info.SetTable("goadmin_companys").SetTitle("厂商管理").SetDescription("")

	formList := goadminCompanys.GetForm()

	isEdit := ctx.Query("__goadmin_edit_pk")
	if isEdit == "" {
		formList.AddField("登录账号", "username", db.Varchar, form.Text).FieldMust().FieldNotAllowEdit()
	}
	formList.AddField("企业名称", "name", db.Varchar, form.Text).FieldMust()
	formList.AddField("企业简介", "desc", db.Varchar, form.Text).FieldMust()
	formList.AddField("企业法人", "corporate", db.Varchar, form.Text).FieldMust()
	formList.AddField("企业地址", "address", db.Varchar, form.Text).FieldMust()
	formList.AddField("联系方式", "contact", db.Varchar, form.Text).FieldMust()

	formList.SetTable("goadmin_companys").SetTitle("厂商管理").SetDescription("")
	formList.SetInsertFn(func(values form2.Values) error {
		if values.IsEmpty("username", "name", "desc", "corporate", "address", "contact") {
			return errors.New("can not be empty")
		}
		_, txErr := db.WithDriver(globalConnection).WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {
			// 用户
			user, createUserErr := models.User().WithTx(tx).SetConn(globalConnection).New(values.Get("username"), "$2a$10$eNr29srZViBfZNLzMzeMSeXU8/SvJZKpg53s.00.MUva5MyoSCJbG", values.Get("username"), "")

			if db.CheckError(createUserErr, db.INSERT) {
				return createUserErr, nil
			}
			// 厂商
			companyId, addCompanyErr := db.WithDriver(globalConnection).WithTx(tx).Table("goadmin_companys").
				Insert(dialect.H{
						"name": values.Get("name"),
						"desc": values.Get("desc"),
						"corporate": values.Get("corporate"),
						"address": values.Get("address"),
						"contact": values.Get("contact"),
				})

			if db.CheckError(addCompanyErr, db.INSERT) {
				return addCompanyErr, nil
			}

			// 厂商 & 用户 关联
			_, addCompanyUserErr := db.WithDriver(globalConnection).WithTx(tx).Table("goadmin_company_users").
				Insert(dialect.H{
					"company_id": companyId,
					"user_id": user.Id,
				})

			if db.CheckError(addCompanyUserErr, db.INSERT) {
				return addCompanyUserErr, nil
			}
			// 角色
			values.Add("role_id[]", "3")
			for i := 0; i < len(values["role_id[]"]); i++ {
				_, addRoleErr := user.WithTx(tx).AddRole(values["role_id[]"][i])
				if db.CheckError(addRoleErr, db.INSERT) {
					return addRoleErr, nil
				}
			}
			return nil, nil
		})
		return txErr
	})

	return goadminCompanys
}
