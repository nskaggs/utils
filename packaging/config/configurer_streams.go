// Copyright 2018 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package config

import (
	"github.com/juju/utils/packaging"
)

// streamsConfigurer is the PackagingConfigurer implementation for apt-based systems.
type streamsConfigurer struct {
	*baseConfigurer
}

// RenderSource is defined on the PackagingConfigurer interface.
func (c *streamsConfigurer) RenderSource(src packaging.PackageSource) (string, error) {
	return src.RenderSourceFile(StreamsSourceTemplate)
}

// RenderPreferences is defined on the PackagingConfigurer interface.
func (c *streamsConfigurer) RenderPreferences(src packaging.PackagePreferences) (string, error) {
	// TODO (aznashwan): research a way of using streams-priorities in the context
	// of single/multiple package pinning and implement it.
	return "", nil
}

// ApplyCloudArchiveTarget is defined on the PackagingConfigurer interface.
func (c *streamsConfigurer) ApplyCloudArchiveTarget(pack string) []string {
	// TODO (aznashwan): implement target application when archive is available.
	return []string{pack}
}
