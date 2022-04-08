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

type PushoverConfig struct {
	Expire       *Duration         `json:"expire,omitempty"`
	Html         bool              `json:"html,omitempty"`
	HttpConfig   *HttpClientConfig `json:"http_config,omitempty"`
	Message      string            `json:"message,omitempty"`
	Priority     string            `json:"priority,omitempty"`
	Retry        *Duration         `json:"retry,omitempty"`
	SendResolved bool              `json:"send_resolved,omitempty"`
	Sound        string            `json:"sound,omitempty"`
	Title        string            `json:"title,omitempty"`
	Token        *Secret           `json:"token,omitempty"`
	Url          string            `json:"url,omitempty"`
	UrlTitle     string            `json:"url_title,omitempty"`
	UserKey      *Secret           `json:"user_key,omitempty"`
}
