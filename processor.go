package rilcommon

type Processor interface {
	Name() string
	Execute(string, interface{}) (interface{}, error)
}

type ProcessorBuilder interface {
	New(...interface{}) Processor
}
