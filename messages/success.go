package messages

func Success(val interface{}) map[string]interface{} {
	var Success = map[string]interface{}{
		"Data":   val,
		"Status": "Success",
	}
	return Success
}
