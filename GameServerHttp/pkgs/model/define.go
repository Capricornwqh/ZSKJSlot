package model

// 默认密码
const DEFAULT_PASSWORD = "hT4w#I2*mX"

// 对象类型
type ObjectType int

const (
	OBJECTTYPE_NOOBJECT ObjectType = 0
	OBJECTTYPE_LANTERN  ObjectType = 1 //心愿
	OBJECTTYPE_ANSWER   ObjectType = 2 //回答
	OBJECTTYPE_COMMENT  ObjectType = 3 //评论
	OBJECTTYPE_TAG      ObjectType = 4 //标签
)

// 排序类型
type SortType int

const (
	SORTTYPE_NOSORT    SortType = 0
	SORTTYPE_TIME_DESC SortType = 1 //时间降序
	SORTTYPE_TIME_ASC  SortType = 2 //时间升序
	SORTTYPE_HOT_SCORE SortType = 3 //热度
	SORTTYPE_NO_ANSWER SortType = 4 //无回答
	SORTTYPE_BRILLIANT SortType = 5 //精彩
)
