package intrinsics

// Ref intrinsic
func Ref(ref string) map[string]interface{} {
	return map[string]interface{}{
		"Ref": ref,
	}
}

// Sub intrinsic
func Sub(sub string) map[string]interface{} {
	return map[string]interface{}{
		"Fn::Sub": sub,
	}
}

// GetAtt intrinsic
func GetAtt(vars []string) map[string]interface{} {
	return map[string]interface{}{
		"Fn::GetAtt": vars,
	}
}

// Join intrinsic
func Join(separator string, items []interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Fn::Join": []interface{}{
			separator,
			items,
		},
	}
}
