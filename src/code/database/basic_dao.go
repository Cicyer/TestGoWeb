package database

import "reflect"

type BasicDao interface {
	selectOne(T interface{}) interface{}
	selectPage(page Page, conditionSql string) []interface{}
}
type Table struct {
	TableName string
}
type Page struct {
	start int
	pageSize int
}

func BindBasicDaoMethods(object *interface{}) BasicDao {
	//先判断传入类是否是Table的子类
	table := reflect.TypeOf(object).Elem()
	//抽出所有属性存为列名,顺便判断是否有表名,没有表名提示错误
	columns := []string
	for i := 0; i < table.NumField(); i++ {
		doc[table.Field(i).Tag.Get("json")] = table.Field(i).Tag.Get("doc")
	}
	tableName := table.FieldByName("TableName")

	//将BasicDao接口的方法实现绑定到该结构体上
}