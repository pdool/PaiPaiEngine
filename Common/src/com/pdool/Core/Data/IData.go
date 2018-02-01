package Data

const (
	UNKNOW = iota
	INT     // byte,short,int,long int64
	FLOAT   // float,double
	STRING
	OBJECT  // NFIdent

)

type IData interface {
	/** 设置值，类型必须和之前一致*/
	SetInt(value int)
	SetFloat(value float32)
	SetString(value int)
	SetObj(value interface{})

	/**
 * 从other拷贝数据，类型可以不同
 * @param other
 */
	getInt() int64
	getFloat() float32
	getString() string
	getObject() interface{}
}
