package database

import (
	"github.com/jinzhu/gorm"
)

//基础dao,根据绑定的table自动实现一些基础的增删改查功能
type IBasicDao interface {
	//会自动从DB中获取一个连接
	CreateCondition(*gorm.DB) *condition
	SelectOne(condition *condition) *interface{}
	SelectPage(condition *condition,page *Page) []*interface{}
	//返回这个表一个实例的方法
	newInstance() *interface{}

}

type BasicDao struct {

	Error error
}

func (b BasicDao)CreateCondition(db *gorm.DB) *condition{
	condition := condition{
		dB:db,
	}
	return &condition
}

func (b BasicDao)SelectOne(condition *condition) *interface{}{
	return nil
}
func (b BasicDao)SelectPage(condition *condition,page *Page)  []*interface{}{
	return nil
}

//从全局连接池获取表所在数据源连接的方法
type selectOne func(condition *condition) *interface{}
type selectPage func(condition *condition,page *Page) []*interface{}
type condition struct {
	dB *gorm.DB
}
//对Condition绑定可用的gorm条件选择器绑定
//条件where
func (c condition) Where(query interface{}, args ...interface{}) *condition{
	c.dB = c.dB.Where(query,args)
	return &c
}
//指定只返回某列
func (c condition) Select(query interface{}, args ...interface{}) *condition{
	c.dB = c.dB.Select(query,args)
	return &c
}

type Page struct {
	PageSize int
	PageNumber int
	Total int
}


type myTestDao struct {
	//由BasicDao来提供基本的操作方法,而myTestDao则实现getDB与newInstance方法
	BasicDao
}
func (T myTestDao)newInstance() *interface{}{
	return nil
}

func test()  {

}

