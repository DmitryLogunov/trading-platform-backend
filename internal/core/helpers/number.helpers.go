package helpers

// SafeParseFloat64 : returns null if value null otherwise parses *float32 -> *float64
func SafeParseFloat64(value *float32) *float64 {
	if value == nil {
		return nil
	}

	parsedValue := float64(*value)

	return &parsedValue
}

// SafeParseFloat32 : returns null if value null otherwise parses *float64 -> *float32
func SafeParseFloat32(value *float64) *float32 {
	if value == nil {
		return nil
	}

	parsedValue := float32(*value)

	return &parsedValue
}

// ValueOrNil32 : returns defaultValue if it's not null or value and parses *float64 -> *float32
func ValueOrNil32(value float64, defaultValue *float64) *float32 {
	if defaultValue == nil {
		return SafeParseFloat32(&value)
	}

	return SafeParseFloat32(defaultValue)
}
