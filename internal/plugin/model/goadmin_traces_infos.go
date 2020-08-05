package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template/types/action"

	//"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	//"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetGoadminTracesInfosTable(ctx *context.Context) table.Table {
	traceId := ctx.Query("trace_id")

	goadminTracesInfos := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := goadminTracesInfos.GetInfo().HideFilterArea()
	info.AddField("ID", "id", db.Int).FieldFilterable()
	info.AddField("编辑时间", "modify_date", db.Timestamp)
	info.AddField("内容", "content", db.Text)
	info.AddField("创建时间", "created_at", db.Timestamp)
	info.AddField("更新时间", "updated_at", db.Timestamp)
	info.AddField("溯源编号", "trace_id", db.Int).FieldFilterable().FieldHide()
	info.SetTable("goadmin_traces_infos").SetTitle("溯源详情").SetDescription("溯源详细内容")

	formList := goadminTracesInfos.GetForm()
	formList.SetTable("goadmin_traces_infos").SetTitle("溯源详情").SetDescription("溯源详细内容")
	formList.AddField("编辑时间", "modify_date", db.Timestamp, form.Datetime).FieldMust()
	formList.AddField("内容", "content", db.Text, form.RichText).FieldMust()
	formList.AddField("溯源编号", "trace_id", db.Int, form.Number).FieldDefault(traceId).FieldHide()
	info.AddActionButton("附件", action.Jump("/admin/info/goadmin_traces_info_access?trace_info_id={{.Id}}"))
	//formList.AddTable("图片", "setting", func(panel *types.FormPanel) {
	//	panel.AddField("上传", "image", db.Varchar, form.File).FieldHideLabel()
	//})
	//formList.AddField("上传图片", "image", db.Varchar, form.Multifile).FieldOptionExt(map[string]interface{}{
	//	"maxFileCount": 10,
	//}).FieldDisplay(func(value types.FieldModel) interface{} {
	//	images, err := db.WithDriver(globalConnection).Table("goadmin_traces_info_access").Select("content").Where("trace_info_id", "=", value.ID).Where("acc_type", "=", 1).All()
	//	if err != nil{
	//		panic("获取商品失败.")
	//	}
	//	res := "["
	//	for _, item := range images {
	//		res += "'" + config.GetStore().URL(item["content"].(string)) + "',"
	//	}
	//	res += "]"
	//	return res
	//})

	// 创建
	formList.SetInsertFn(func(values form2.Values) error{

		traceId := values.Get("trace_id")
		// 添加溯源详情
		_, addTraceInfoErr := db.WithDriver(globalConnection).Table("goadmin_traces_infos").
			Insert(dialect.H{
				"trace_id": traceId,
				"modify_date" : values.Get("modify_date"),
				"content": values.Get("content"),
			})
		if addTraceInfoErr != nil {
			return addTraceInfoErr
		}


		// 添加溯源详情附件
		//for k, v := range values {
		//	if k == "image" {
		//		for _, imageUrl := range v {
		//			_, addTraceInfoImageErr := db.WithDriver(globalConnection).Table("goadmin_traces_info_access").
		//				Insert(dialect.H{
		//					"trace_info_id" : traceInfoId,
		//					"content": imageUrl,
		//					"acc_type": 1,
		//				})
		//			if addTraceInfoImageErr != nil {
		//				return addTraceInfoImageErr
		//			}
		//		}
		//	}
		//}

		return nil
	})

	// 更新
	formList.SetUpdateFn(func(values form2.Values) error{
		//traceId := values.Get("trace_id")
		// 编辑溯源详情
		_, addTraceInfoErr := db.WithDriver(globalConnection).Table("goadmin_traces_infos").Where("id","=", values.Get("id")).
			Update(dialect.H{
				"modify_date" : values.Get("modify_date"),
				"content": values.Get("content"),
			})
		if addTraceInfoErr != nil {
			return addTraceInfoErr
		}

		return nil
	})
	return goadminTracesInfos
}
