package common

import (
    "fmt"
    "github.com/hashicorp/terraform-plugin-framework/diag"
    "github.com/pgEdge/terraform-provider-pgedge/client"
)

func HandleProviderError(err error, operation string) diag.Diagnostic {
    if apiErr, ok := err.(*client.APIError); ok {
        if apiErr.StatusCode == 404 {
            return nil
        }
        return diag.NewErrorDiagnostic(
            fmt.Sprintf("API Error during %s", operation),
            fmt.Sprintf("Status code: %d, Message: %s", apiErr.StatusCode, apiErr.Message),
        )
    } else {
        return diag.NewErrorDiagnostic(
            fmt.Sprintf("Error during %s", operation),
            err.Error(),
        )
    }
}