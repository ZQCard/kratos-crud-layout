package model

type BusinessModuleNameTable struct {
	Id int64
}

func (BusinessModuleNameTable) TableName() string {
	return "business_module_name_table"
}
