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

import (
	"time"
)

// ActiveSyncStatusDTO holds the information for LDAP background Sync
type ActiveSyncStatusDto struct {
	Enabled  bool        `json:"enabled,omitempty"`
	NextSync time.Time   `json:"nextSync,omitempty"`
	PrevSync *SyncResult `json:"prevSync,omitempty"`
	Schedule string      `json:"schedule,omitempty"`
}
