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

type MetricRequest struct {
	Debug bool `json:"debug,omitempty"`
	// From Start time in epoch timestamps in milliseconds or relative using Grafana time units.
	From string `json:"from"`
	// queries.refId – Specifies an identifier of the query. Is optional and default to “A”. queries.datasourceId – Specifies the data source to be queried. Each query in the request must have an unique datasourceId. queries.maxDataPoints - Species maximum amount of data points that dashboard panel can render. Is optional and default to 100. queries.intervalMs - Specifies the time interval in milliseconds of time series. Is optional and defaults to 1000.
	Queries []Json `json:"queries"`
	// To End time in epoch timestamps in milliseconds or relative using Grafana time units.
	To string `json:"to"`
}
