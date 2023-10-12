package demo1

// TableA ...
type TableA struct {
	ID       int
	ValueInt int
}

// TableB ...
type TableB struct {
	ID        int
	ValueBool bool
}

// TableC ...
type TableC struct {
	ID       int
	ValueStr string
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamePrefix = ""

// Prefix 取表名前缀
func Prefix() string {
	return theTableNamePrefix
}

// // PrefixSetter ...
// func PrefixSetter() libgorm.PrefixSetter {
// 	return func(prefix string) {
// 		theTableNamePrefix = prefix
// 	}
// }

////////////////////////////////////////////////////////////////////////////////

// TableName ...
func (t TableA) TableName() string {
	return theTableNamePrefix + "ta"
}

// TableName ...
func (t TableB) TableName() string {
	return theTableNamePrefix + "tb"
}

// TableName ...
func (t TableC) TableName() string {
	return theTableNamePrefix + "tc"
}

////////////////////////////////////////////////////////////////////////////////
