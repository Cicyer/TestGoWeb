package database

import (
	//"reflect"
	"strconv"
	//"fmt"
	"time"
)
//基础dao,根据绑定的table自动实现一些基础的增删改查功能
type BasicDao struct {
	table Table
	Error error
}
type Table struct {
	TableName string
	columns []string
	//结构体对应表列的属性名
	columnProperties []string
}

type Page struct {
	start int
	pageSize int
}
//sql条件生成类
type Condition struct {
	whereConditions []string
	values []string
	Error error
}

//绑定sql条件生成的可选操作,子类也会有
func (condition *Condition) eq(columnName string, value interface{}) *Condition{
	if condition.whereConditions[len(condition.whereConditions)] != "" {
		condition.whereConditions[len(condition.whereConditions)] += " AND "
	}
	condition.whereConditions[len(condition.whereConditions)] += *castToSqlValue(columnName,"=",value)
	return condition
}

func castToSqlValue(columnName string ,connector string, value interface{}) *string {
	var result string = columnName+connector
	switch value.(type) {
	case string:
		result += "'" + value.(string) + "'"
	case int:
		result += strconv.Itoa(value.(int))
	case int64:
		result += strconv.FormatInt(value.(int64), 10)
	case bool:
		if value.(bool) {
			result += "1"
		} else {
			result += "0"
		}
	case float32:
		result += strconv.FormatFloat(value.(float64),'E',-1,32)
	case float64:
		result += strconv.FormatFloat(value.(float64),'E',-1,64)
	case time.Time:

	default:
		//无法解析的类型 记录日志,不做操作
		result = ""
	}
	return &result
}



//func BindBasicDaoMethods(object *interface{}) BasicDao {
//	//抽出所有属性存为列名,顺便判断是否有表名,没有表名提示错误
//	table := reflect.TypeOf(object).Elem()
//	columns := []string
//	for i := 0; i < table.NumField(); i++ {
//		doc[table.Field(i).Tag.Get("json")] = table.Field(i).Tag.Get("doc")
//	}
//	tableName := table.FieldByName("TableName")
//
//	//将BasicDao接口的方法实现绑定到该结构体上
//}

