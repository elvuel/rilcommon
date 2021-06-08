package rilcommon

type Trigger interface {
	Emit(interface{}) (interface{}, error)
}
