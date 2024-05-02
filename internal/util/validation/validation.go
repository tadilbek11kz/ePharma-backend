package validation

type Validator interface {
	ValidateStruct(model interface{}) error
}
