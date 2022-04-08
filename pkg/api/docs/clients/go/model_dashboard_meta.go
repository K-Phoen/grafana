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

type DashboardMeta struct {
	CanAdmin              bool      `json:"canAdmin,omitempty"`
	CanDelete             bool      `json:"canDelete,omitempty"`
	CanEdit               bool      `json:"canEdit,omitempty"`
	CanSave               bool      `json:"canSave,omitempty"`
	CanStar               bool      `json:"canStar,omitempty"`
	Created               time.Time `json:"created,omitempty"`
	CreatedBy             string    `json:"createdBy,omitempty"`
	Expires               time.Time `json:"expires,omitempty"`
	FolderId              int64     `json:"folderId,omitempty"`
	FolderTitle           string    `json:"folderTitle,omitempty"`
	FolderUid             string    `json:"folderUid,omitempty"`
	FolderUrl             string    `json:"folderUrl,omitempty"`
	HasAcl                bool      `json:"hasAcl,omitempty"`
	IsFolder              bool      `json:"isFolder,omitempty"`
	IsHome                bool      `json:"isHome,omitempty"`
	IsSnapshot            bool      `json:"isSnapshot,omitempty"`
	IsStarred             bool      `json:"isStarred,omitempty"`
	Provisioned           bool      `json:"provisioned,omitempty"`
	ProvisionedExternalId string    `json:"provisionedExternalId,omitempty"`
	Slug                  string    `json:"slug,omitempty"`
	Type_                 string    `json:"type,omitempty"`
	Updated               time.Time `json:"updated,omitempty"`
	UpdatedBy             string    `json:"updatedBy,omitempty"`
	Url                   string    `json:"url,omitempty"`
	Version               int64     `json:"version,omitempty"`
}
