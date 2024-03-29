package model

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "product" => http://localhost:9033/admin/info/product
//
// "goadmin_company_products" => http://localhost:9033/admin/info/goadmin_company_products
// "goadmin_companys" => http://localhost:9033/admin/info/goadmin_companys
// "goadmin_company_users" => http://localhost:9033/admin/info/goadmin_company_users
//
// "goadmin_traces" => http://localhost:9033/admin/info/goadmin_traces
// "goadmin_traces_info_access" => http://localhost:9033/admin/info/goadmin_traces_info_access
// "goadmin_traces_infos" => http://localhost:9033/admin/info/goadmin_traces_infos
//
// example end
//

var Generators = map[string]table.Generator{
	"goadmin_company_products": 	GetGoadminCompanyProductsTable,
	"goadmin_companys":         	GetGoadminCompanysTable,
	"goadmin_company_users":    	GetGoadminCompanyUsersTable,
	"goadmin_traces":				GetGoadminTracesTable,
	"goadmin_traces_info_access" :  GetGoadminTracesInfoAccessTable,
	"goadmin_traces_infos" : 		GetGoadminTracesInfosTable,
}

var globalConnection db.Connection

func SetConnection(c db.Connection) {
	globalConnection = c
}
