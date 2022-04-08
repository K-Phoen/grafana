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

type CreateOrUpdateConfigCmd struct {
	DashboardId        int64             `json:"dashboardId,omitempty"`
	DashboardUid       string            `json:"dashboardUid,omitempty"`
	EnableCsv          bool              `json:"enableCsv,omitempty"`
	EnableDashboardUrl bool              `json:"enableDashboardUrl,omitempty"`
	Message            string            `json:"message,omitempty"`
	Name               string            `json:"name,omitempty"`
	Options            *ReportOptionsDto `json:"options,omitempty"`
	Recipients         string            `json:"recipients,omitempty"`
	ReplyTo            string            `json:"replyTo,omitempty"`
	Schedule           *ScheduleDto      `json:"schedule,omitempty"`
	State              string            `json:"state,omitempty"`
	TemplateVars       interface{}       `json:"templateVars,omitempty"`
}
