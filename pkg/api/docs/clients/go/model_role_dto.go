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

type RoleDto struct {
	Created     time.Time    `json:"created,omitempty"`
	Delegatable bool         `json:"delegatable,omitempty"`
	Description string       `json:"description,omitempty"`
	DisplayName string       `json:"displayName,omitempty"`
	Group       string       `json:"group,omitempty"`
	Hidden      bool         `json:"hidden,omitempty"`
	Name        string       `json:"name,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
	Uid         string       `json:"uid,omitempty"`
	Updated     time.Time    `json:"updated,omitempty"`
	Version     int64        `json:"version,omitempty"`
}
