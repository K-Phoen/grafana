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

type Config struct {
	Global            *GlobalConfig      `json:"global,omitempty"`
	InhibitRules      []InhibitRule      `json:"inhibit_rules,omitempty"`
	MuteTimeIntervals []MuteTimeInterval `json:"mute_time_intervals,omitempty"`
	Route             *Route             `json:"route,omitempty"`
	Templates         []string           `json:"templates,omitempty"`
}
