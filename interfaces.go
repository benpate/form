package form

// UnmarshalMaper wraps the UnmarshalMap interface
type UnmarshalMaper interface {

	// UnmarshalMap returns a value in the format map[string]interface
	UnmarshalMap() map[string]any
}
