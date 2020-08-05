package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetGoadminTracesInfoAccessTable(ctx *context.Context) table.Table {
	traceInfoId := ctx.Query("trace_info_id")

	goadminTracesInfoAccess := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := goadminTracesInfoAccess.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int).FieldFilterable()
	info.AddField("地址", "content", db.Text).FieldDisplay(func(value types.FieldModel) interface{} {
		return config.GetStore().URL(value.Value)
	}).FieldDownLoadable()
	info.AddField("溯源详情编号", "trace_info_id", db.Int).FieldFilterable()
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	info.AddField("类型", "acc_type", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "1" {
			return "图片"
		}
		if model.Value == "0" {
			return "视频"
		}
		return "unknown"
	})
	info.SetTable("goadmin_traces_info_access").SetTitle("溯源详情附件").SetDescription("")

	formList := goadminTracesInfoAccess.GetForm()
	//formList.AddField("ID", "id", db.Int, form.Default).FieldHide()
	formList.AddField("地址", "content", db.Text, form.File)
	formList.AddField("溯源详情编号", "trace_info_id", db.Int, form.Number).FieldDefault(traceInfoId).FieldHide()
	//formList.AddField("创建时间", "created_at", db.Timestamp, form.Datetime)
	//formList.AddField("更新时间", "updated_at", db.Timestamp, form.Datetime)
	//formList.AddField("Deleted_at", "deleted_at", db.Timestamp, form.Datetime)
	formList.AddField("类型", "acc_type", db.Varchar, form.SelectSingle).FieldOptions(
		types.FieldOptions{
			{Text: "图片",Value: "1"},
			{Text: "视频",Value: "0"},
		}).

		FieldDefault("1")
	formList.SetTable("goadmin_traces_info_access").SetTitle("溯源详情附件").SetDescription("")

	return goadminTracesInfoAccess
}
