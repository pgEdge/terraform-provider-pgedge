package common

import "github.com/hashicorp/terraform-plugin-framework/types"

// Helper function to convert types.List to []string
func ConvertTFListToStringSlice(list types.List) []string {
	if list.IsNull() || list.IsUnknown() {
		return nil
	}
	var result []string
	for _, elem := range list.Elements() {
		if str, ok := elem.(types.String); ok {
			result = append(result, str.ValueString())
		}
	}
	return result
}
