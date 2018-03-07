// Copyright 2018 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package commands

const (
	// the basic command for all streams-get calls:
	streamsget = "curl -sSfw ' binaries from %{url_effective} downloaded: HTTP %{http_code}; time %{time_total}s; size %{size_download} bytes; speed %{speed_download} bytes/s '"

	// the basic format for specifying a proxy option for streams:
	streamsProxySettingFormat = "Acquire::%s::Proxy %q;"

	// disable proxy for a specific host
	streamsNoProxySettingFormat = "Acquire::%s::Proxy::%q \"DIRECT\";"
)

// streamsCmder is the packageCommander instantiation for streams-based systems.
var streamsCmder = packageCommander{
	prereq:                buildCommand(":", "#No action here"),
	update:                buildCommand(":", "#No action here"),
	upgrade:               buildCommand(":", "#No action here"),
	install:               buildCommand(streamsget, "install"),
	remove:                buildCommand(":", "#No action here"),
	purge:                 buildCommand(":", "#No action here"),
	search:                buildCommand(":", "#No action here"),
	isInstalled:           buildCommand(":", "#No action here"),
	listAvailable:         buildCommand(":", "#No action here"),
	listInstalled:         buildCommand(":", "#No action here"),
	addRepository:         buildCommand(":", "#No action here"),
	listRepositories:      buildCommand(":", "#No action here"),
	removeRepository:      buildCommand(":", "#No action here"),
	cleanup:               buildCommand(":", "#No action here"),
	getProxy:              buildCommand(":", "#No action here"),
	proxySettingsFormat:   buildCommand(":", "#No action here"),
	setProxy:              buildCommand(":", "#No action here"),
	noProxySettingsFormat: buildCommand(":", "#No action here"),
	setNoProxy:            buildCommand(":", "#No action here"),
}
