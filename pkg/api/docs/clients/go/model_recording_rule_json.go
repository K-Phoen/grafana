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

// RecordingRuleJSON is the external representation of a recording rule
type RecordingRuleJson struct {
	Active            bool                     `json:"active,omitempty"`
	Count             bool                     `json:"count,omitempty"`
	Description       string                   `json:"description,omitempty"`
	DestDataSourceUid string                   `json:"dest_data_source_uid,omitempty"`
	Id                string                   `json:"id,omitempty"`
	Interval          int64                    `json:"interval,omitempty"`
	Name              string                   `json:"name,omitempty"`
	PromName          string                   `json:"prom_name,omitempty"`
	Queries           []map[string]interface{} `json:"queries,omitempty"`
	Range_            int64                    `json:"range,omitempty"`
	TargetRefId       string                   `json:"target_ref_id,omitempty"`
}
