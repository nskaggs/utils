// Copyright 2018 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package commands_test

import (
	"github.com/juju/utils/packaging/commands"
	gc "gopkg.in/check.v1"
)

var _ = gc.Suite(&StreamsSuite{})

type StreamsSuite struct {
	paccmder commands.PackageCommander
}

func (s *StreamsSuite) SetUpSuite(c *gc.C) {
	s.paccmder = commands.NewStreamsPackageCommander()
}
