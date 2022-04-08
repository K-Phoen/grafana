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

type PrometheusRemoteWriteTargetJson struct {
	DataSourceUid   string `json:"data_source_uid,omitempty"`
	Id              string `json:"id,omitempty"`
	RemoteWritePath string `json:"remote_write_path,omitempty"`
}
