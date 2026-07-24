package templates

import "embed"

//go:embed web/templates/layouts/*.html web/templates/pages/*.html web/templates/partials/*.html
var TemplatesFS embed.FS
