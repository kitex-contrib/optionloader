package interfaces

type translator func(config map[string]interface{}) ([]struct{}, error)

// Loader defines the interface for loading options.
type Loader interface {
	AppendSource(reader struct{}) error
	Load() error
	GetOptions() ([]struct{}, error)
	RegisterTranslator(fieldName string, customedTranslator translator) error
	DeregisterTranslator(fieldName string) error
}
