package converter

func BooleanString(b bool) string {
	if b {
		return "true"
	}

	return "false"
}
