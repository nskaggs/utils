// Copyright 2018 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package config_test

import (
	"fmt"

	jc "github.com/juju/testing/checkers"
	"github.com/juju/utils/packaging/config"
	gc "gopkg.in/check.v1"
)

var _ = gc.Suite(&StreamsSuite{})

type StreamsSuite struct {
	pacconfer config.PackagingConfigurer
}

func (s *StreamsSuite) SetUpSuite(c *gc.C) {
	s.pacconfer = config.NewStreamsPackagingConfigurer(testedSeriesCentOS)
}

func (s *StreamsSuite) TestDefaultPackages(c *gc.C) {
	c.Assert(s.pacconfer.DefaultPackages(), gc.DeepEquals, config.CentOSDefaultPackages)
}

func (s *StreamsSuite) TestGetPackageNameForSeriesSameSeries(c *gc.C) {
	for _, pack := range testedPackages {
		res, err := s.pacconfer.GetPackageNameForSeries(pack, testedSeriesCentOS)
		c.Assert(err, jc.ErrorIsNil)
		c.Assert(res, gc.Equals, pack)
	}
}

func (s *StreamsSuite) TestGetPackageNameForSeriesErrors(c *gc.C) {
	for _, pack := range testedPackages {
		res, err := s.pacconfer.GetPackageNameForSeries(pack, "some-other-series")
		c.Assert(res, gc.Equals, "")
		c.Assert(err, gc.ErrorMatches, fmt.Sprintf("no equivalent package found for series %s: %s", "some-other-series", pack))
	}
}

func (s *StreamsSuite) TestRenderSource(c *gc.C) {
	expected, err := testedSource.RenderSourceFile(config.StreamsSourceTemplate)
	c.Assert(err, jc.ErrorIsNil)

	res, err := s.pacconfer.RenderSource(testedSource)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(res, gc.Equals, expected)
}
