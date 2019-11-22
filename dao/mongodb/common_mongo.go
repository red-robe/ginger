package mongodb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
通用新增数据到mongo collection

@param collectionName string - Collection name
@param dataMap map[string]interface{} - Insert data which is a map type

@return err error
*/
func Insert(collectionName string,dataMap map[string]interface{}) (err error){

	err = M(collectionName, func(c *mgo.Collection) error {
		return  c.Insert(dataMap)
	})
	return
}

/*
通用更新数据
@param collectionName string - Collection name
@param selector map[string]interface{} - update condition
@param updater map[string]interface{} - update data

@return err error
*/
func Update(collectionName string,selector,updater map[string]interface{})  (err error) {
	err = M(collectionName, func(c *mgo.Collection) error {
		return c.Update(selector, updater)
	})
	return
}

/*
通用更新数据
@param collectionName string - Collection name
@param id string - mongo id string
@param updater map[string]interface{} - update data

@return err error
*/
func UpdateId(collectionName,id string,updater map[string]interface{})  (err error) {
	err = M(collectionName, func(c *mgo.Collection) error {
		mongoId := bson.ObjectId(id)
		return c.UpdateId(mongoId, updater)
	})
	return
}


/*
通用更新数据
@param collectionName string - Collection name
@param selector map[string]interface{} - update condition
@param updater map[string]interface{} - update data

@return changeInfo *mgo.ChangeInfo
@return err error
*/
func UpdateAll(collectionName string,selector,updater map[string]interface{})  (changeInfo *mgo.ChangeInfo,err error) {
	err = M(collectionName, func(c *mgo.Collection) error {
		changeInfo, err = c.UpdateAll(selector, updater)
		return err
	})
	return
}

/*
更新插入数据
@param collectionName string - Collection name
@param selector map[string]interface{} - upsert condition
@param updater map[string]interface{} - upsert data

@return changeInfo *mgo.ChangeInfo
@return err error
*/
func Upsert(collectionName string,selector,updater map[string]interface{}) (changeInfo *mgo.ChangeInfo,err error) {

	err = M(collectionName, func(c *mgo.Collection) error {
		changeInfo, err = c.Upsert(selector, updater)
		return err
	})
	return
}

/*
更新插入数据
@param collectionName string - Collection name
@param id string - mongo id string
@param updater map[string]interface{} - upsert data

@return err error
*/
func UpsertId(collectionName,id string,updater map[string]interface{}) (changeInfo *mgo.ChangeInfo,err error) {
	err = M(collectionName, func(c *mgo.Collection) error {
		mongoId := bson.ObjectId(id)
		changeInfo ,err = c.UpsertId(mongoId, updater)
		return err
	})
	return
}



/*
通用删除数据 只删除符合条件的第一个
@param collectionName string - Collection name
@param selector map[string]interface{} - remove condition

@return error
*/
func Remove(collectionName string,selector map[string]interface{}) (err error) {
	err = M(collectionName, func(c *mgo.Collection) (err error) {
		return c.Remove(selector)
	})
	return
}

/*
通用删除数据 删除所有符合条件的数据
@param collectionName string - Collection name
@param selector map[string]interface{} - remove condition

@return changeInfo *mgo.ChangeInfo
@return error
*/
func RemoveAll(collectionName string,selector map[string]interface{}) (changeInfo *mgo.ChangeInfo,err error) {
	err = M(collectionName, func(c *mgo.Collection) (err error) {
		changeInfo,err = c.RemoveAll(selector)
		return
	})
	return
}

/*
通用删除数据 只删除符合条件的id
@param collectionName string - Collection name
@param id string - remove data with mongo id

@return error
*/
func RemoveId(collectionName string,id string) (err error) {
	err = M(collectionName, func(c *mgo.Collection) (err error) {
		mongoId := bson.ObjectId(id)
		return c.RemoveId(mongoId)
	})
	return
}



type FindResult map[string]interface{}
type FindResults []map[string]interface{}

/*
查找单个文档
@param collectionName string - Collection name
@param selector map[string]interface{} - find condition

@return rs FindResult
@return error
*/
func FindOne(collectionName string,selector map[string]interface{}) (rs FindResult,err error)  {
	err = M(collectionName, func(c *mgo.Collection) error {
		return c.Find(selector).One(&rs)
	})
	return
}

/*
根据id查找
@param collectionName string - Collection name
@param id string - find data with mongo id

@return error

*/
func FindById(collectionName,id string)  (rs FindResult,err error)  {
	err = M(collectionName, func(c *mgo.Collection) error {
		mongoId := bson.ObjectId(id)
		err := c.FindId(mongoId).One(&rs)
		return err
	})
	return
}

/*
查找多个文档
@param collectionName string - Collection name
@param selector map[string]interface{} - find condition

@return rs FindResults
@return error
*/
func FindAll(collectionName string,selector map[string]interface{}) (rs FindResults,err error)  {
	err = M(collectionName, func(c *mgo.Collection) error {
		return c.Find(selector).All(&rs)
	})
	return
}




/*
查询文档数
@param collectionName string - Collection name
@param selector map[string]interface{} - find condition

@return rs FindResults
@return error
*/
func FindCount(collectionName string,selector map[string]interface{}) (count int,err error){
	err = M(collectionName, func(c *mgo.Collection) error {
		 count,err = c.Find(selector).Count()
		 return err
	})
	return
}








