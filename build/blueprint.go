package build

import (
	"fmt"

	"github.com/oleiade/reflections"
)

//methods for types.go -> Blueprint

//Validate function
func (b *Blueprint) Validate() error {
	for name, field := range schemaMap {
		// check if field exists
		ok, err := reflections.HasField(b, name)
		if !ok || err != nil {
			if !ok {
				return fmt.Errorf("Field %s is not a valid field name", name)
			}
			return err
		}

		// check field is populated
		val, err := reflections.GetField(b, name)
		if err != nil {
			return err
		}
		if val == field.NullValue && field.Required {
			return fmt.Errorf("Field %s has no value and is required", name)
		}
		if val == field.NullValue && !field.Required {
			err = reflections.SetField(b, name, field.DefaultValue)
			if err != nil {
				return err
			}
		}

		// check field type
		typ, err := reflections.GetFieldType(b, name)
		if err != nil {
			return err
		}
		if typ != field.ValueType {
			return fmt.Errorf("Found type of %s but expected type %s", typ, field.ValueType)
		}
	}
	return nil
}
