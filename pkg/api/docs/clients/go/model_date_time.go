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

// DateTime is a time but it serializes to ISO8601 format with millis It knows how to read 3 different variations of a RFC3339 date time. Most APIs we encounter want either millisecond or second precision times. This just tries to make it worry-free.
type DateTime struct {
}
