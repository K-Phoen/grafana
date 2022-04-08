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

type PatchPrefsCmd struct {
	// The numerical :id of a favorited dashboard
	HomeDashboardId int64             `json:"homeDashboardId,omitempty"`
	Navbar          *NavbarPreference `json:"navbar,omitempty"`
	Theme           string            `json:"theme,omitempty"`
	Timezone        string            `json:"timezone,omitempty"`
	WeekStart       string            `json:"weekStart,omitempty"`
}
