// Copyright 2018 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package config

import (
	"text/template"
)

const (
	// StreamsSourcesDir is the default directory in which streams sourcefiles are located.
	StreamsSourcesDir = "/etc/streams/repos.d"

	// StreamsKeyfileDir is the default directory for streams repository keys.
	StreamsKeyfileDir = "/etc/pki/rpm-gpg/"
)

// StreamsSourceTemplate is the template specific to a streams source file.
var StreamsSourceTemplate = template.Must(template.New("").Parse(`
[{{.Name}}]
name={{.Name}} (added by Juju)
baseurl={{.URL}}
{{if .Key}}gpgcheck=1
gpgkey=%s{{end}}
enabled=1
`[1:]))
