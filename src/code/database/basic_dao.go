package database

import (
	"github.com/jinzhu/gorm"
)

//基础dao,根据绑定的table自动实现一些基础的增删改查功能
type BasicDao struct {
	db    *gorm.DB
	table Table
	Error error
}
type Table struct {
	TableName string
	columns   []string
	//结构体对应表列的属性名
	columnProperties []string
}

//sql条件生成类,对gorm封装
type Condition struct {
	db    *gorm.DB
	Error error
}

//从全局的连接池获取一个数据连接
func getDB() *gorm.DB {
	return nil
}

//绑定BasicDao的基础操作方法,如果是自定义sql则不借助基础的Table,而是直接使用gorm自带的model获取返回

//绑定Condition的条件创建方法

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
