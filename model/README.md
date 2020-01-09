## 中文说明
- 该目录存放mysql相关的业务存储逻辑
- 一般按照数据库表建立相关go文件，每个go文件编写该表相关的业务存储逻辑
- 如有个别模型需查询多个表的数据也可另行创建模块包


### 关于SQL Builder构建语法：

#### 一、构建select语句

方法签名：
> BuildSelect(table string, where map[string]interface{}, field []string) (string,[]interface{},error)

##### 1.常用操作符:

- = 
- \>
- <
- =
- <=
- \>=
- !=
- <>
- in
- not in
- like
- not like
- between
- not between

```
where := map[string]interface{}{
	"foo <>": "aha",
	"bar <=": 45,
	"sex in": []interface{}{"girl", "boy"},
	"name like": "%James",
}
```

##### 2.其他支持的操作符：

- _orderby
- _groupby
- _having
- _limit

```
where := map[string]interface{}{
	"age >": 100,
	"_orderby": "fieldName asc",
	"_groupby": "fieldName",
	"_having": map[string]interface{}{"foo":"bar",},
	"_limit": []uint{offset, row_count},
}
```

#### 注意：

- _having  _groupby 操作符必须同时使用

- _limit 使用方式是无符号整型数组

```
"_limit": []uint{a,b} => LIMIT a,b
"_limit": []uint{a} => LIMIT 0,a
```

如果你想清除where map中的零值可以使用 builder.OmitEmpty
```
where := map[string]interface{}{
		"score": 0,
		"age": 35,
	}
finalWhere := builder.OmitEmpty(where, []string{"score", "age"})
// finalWhere = map[string]interface{}{"age": 35}

// support: Bool, Array, String, Float32, Float64, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Ui

```


#### 3.聚合查询
方法签名：
> AggregateQuery(ctx context.Context, db *sql.DB, table string, where map[string]interface{}, aggregate AggregateSymbleBuilder) (ResultResolver, error)

支持的聚合方法：
- AggregateSum()
- AggregateAvg
- AggregateMax()
- AggregateMin()
- AggregateCount()


举个例子:
```
where := map[string]interface{}{
       "score > ": 100,
       "city in": []interface{}{"Beijing", "Shijiazhuang",}
   }
   
// 聚合数字类型的和
result, err := AggregateQuery(ctx, db, "tableName", where, AggregateSum("age"))
sumAge := result.Int64()

// 聚合查询行数
result,err = AggregateQuery(ctx, db, "tableName", where, AggregateCount("*")) 
numberOfRecords := result.Int64()

// 聚合查询平均数
result,err = AggregateQuery(ctx, db, "tableName", where, AggregateAvg("score"))
averageScore := result.Float64()

```

#### 4.复杂查询
方法签名：
> func NamedQuery(sql string, data map[string]interface{}) (string, []interface{}, error)

对于比较复杂的查询, NamedQuery将会派上用场:
```
condition, values, err := builder.NamedQuery("select * from tb where name={{name}} and id in (select uid from anothertable where score in {{m_score}})", map[string]interface{}{
	"name": "caibirdme",
	"m_score": []float64{3.0, 5.8, 7.9},
})

// 检验构建的条件和值是否和预想sql语句一致
assert.Equal("select * from tb where name=? and id in (select uid from anothertable where score in (?,?,?))", cond)
assert.Equal([]interface{}{"caibirdme", 3.0, 5.8, 7.9}, values)

// 执行
db.Query(condition,values)

```

slice类型的值会根据slice的长度自动展开
这种方式基本上就是手写sql，非常便于DBA review同时也方便开发者进行复杂sql的调优
对于关键系统，推荐使用这种方式



### 二、构建Update语句
方法签名： 

> BuildUpdate(table string, where map[string]interface{}, update map[string]interface{}) (string, []interface{}, error)


BuildUpdate 与 BuildSelect 的条件语句类似，但更新条件不支持以下操作符： 

- _orderby
- _groupby
- _limit
- _having

只支持常规的条件操作符，举个例子：
```
where := map[string]interface{}{
	"foo <>": "aha",
	"bar <=": 45,
	"sex in": []interface{}{"girl", "boy"},
}
update := map[string]interface{}{
	"role": "primaryschoolstudent",
	"rank": 5,
}
cond,vals,err := qb.BuildUpdate("table_name", where, update)

db.Exec(cond, vals...)
```

### 三、构建Insert语句
对应构建插入语句三种方式有三个签名方法：

##### 1.BuildInsert  [insert into ...]
> 方法签名: BuildInsert(table string, data []map[string]interface{}) (string, []interface{}, error)

##### 2.BuildInsertIgnore [insert ignore into ...]
方法签名：
> BuildInsertIgnore(table string, data []map[string]interface{}) (string, []interface{}, error)


##### 3.BuildReplaceInsert [replace into ...]
方法签名：
> BuildReplaceInsert(table string, data []map[string]interface{}) (string, []interface{}, error)

data 是一个映射切片，可一次插入多组数据


```
var data []map[string]interface{}
data = append(data, map[string]interface{}{
    "name": "deen",
    "age":  23,
})
data = append(data, map[string]interface{}{
    "name": "Tony",
    "age":  30,
})
// 普通插入，已存在主键则报错 
cond, vals, err := qb.BuildInsert(table, data)
db.Exec(cond, vals...)

// 如已存在忽略插入 
cond, vals, err := qb.BuildInsertIgnore(table, data)
db.Exec(cond, vals...)

// 替换已存在插入新数据 
cond, vals, err := qb.BuildReplaceInsert(table, data)
db.Exec(cond, vals...)
```

### 四、构建Delete语句
签名方法：
> BuildDelete(table string, where map[string]interface{}) (string, []interface{}, error)

 