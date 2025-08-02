package validation

func ErrorMessages() map[string]string {
	return map[string]string{
		"required": "the field is required",
		"email":    "the field must have a valid email address",
		"min":      "should be more than the limit",
		"max":      "should be less than the limit",
	}
}
