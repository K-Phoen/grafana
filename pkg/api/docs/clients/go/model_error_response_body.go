/*
 * Grafana HTTP API.
 *
 * The Grafana backend exposes an HTTP API, the same API is used by the frontend to do everything from saving dashboards, creating users and updating data sources.
 *
 * API version: 0.0.1
 * Contact: hello@grafana.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ErrorResponseBody struct {
	// Error An optional detailed description of the actual error. Only included if running in developer mode.
	Error_ string `json:"error,omitempty"`
	// a human readable version of the error
	Message string `json:"message"`
	// Status An optional status to denote the cause of the error.  For example, a 412 Precondition Failed error may include additional information of why that error happened.
	Status string `json:"status,omitempty"`
}
