// Copyright 2018 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package manager_test

import (
	"strings"

	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/utils/packaging/commands"
	"github.com/juju/utils/packaging/manager"
	"github.com/juju/utils/proxy"
	gc "gopkg.in/check.v1"
)

var _ = gc.Suite(&StreamsSuite{})

type StreamsSuite struct {
	testing.IsolationSuite
	paccmder commands.PackageCommander
	pacman   manager.PackageManager
}

func (s *StreamsSuite) SetUpSuite(c *gc.C) {
	s.IsolationSuite.SetUpSuite(c)
	s.paccmder = commands.NewStreamsPackageCommander()
	s.pacman = manager.NewStreamsPackageManager()
}

func (s *StreamsSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)
}

func (s *StreamsSuite) TearDownTest(c *gc.C) {
	s.IsolationSuite.TearDownTest(c)
}

func (s *StreamsSuite) TearDownSuite(c *gc.C) {
	s.IsolationSuite.TearDownSuite(c)
}

func (s *StreamsSuite) TestGetProxySettingsEmpty(c *gc.C) {
	cmdChan := s.HookCommandOutput(&manager.CommandOutput, []byte{}, nil)

	out, err := s.pacman.GetProxySettings()
	c.Assert(err, jc.ErrorIsNil)

	cmd := <-cmdChan
	c.Assert(cmd.Args, gc.DeepEquals, strings.Fields(s.paccmder.GetProxyCmd()))
	c.Assert(out, gc.Equals, proxy.Settings{})
}

func (s *StreamsSuite) TestGetProxySettingsConfigured(c *gc.C) {
	const expected = `CommandLine::AsString "Streams-config dump";
Acquire::http::Proxy  "10.0.3.1:3142";
Acquire::https::Proxy "false";
Acquire::ftp::Proxy "none";
Acquire::magic::Proxy "none";`
	cmdChan := s.HookCommandOutput(&manager.CommandOutput, []byte(expected), nil)

	out, err := s.pacman.GetProxySettings()
	c.Assert(err, gc.IsNil)

	cmd := <-cmdChan
	c.Assert(cmd.Args, gc.DeepEquals, strings.Fields(s.paccmder.GetProxyCmd()))

	c.Assert(out, gc.Equals, proxy.Settings{
		Http:  "10.0.3.1:3142",
		Https: "false",
		Ftp:   "none",
	})
}

func (s *StreamsSuite) TestProxySettingsRoundTrip(c *gc.C) {
	initial := proxy.Settings{
		Http:  "some-proxy.local:8080",
		Https: "some-secure-proxy.local:9696",
		Ftp:   "some-ftp-proxy.local:1212",
	}

	expected := s.paccmder.ProxyConfigContents(initial)
	cmdChan := s.HookCommandOutput(&manager.CommandOutput, []byte(expected), nil)

	result, err := s.pacman.GetProxySettings()
	c.Assert(err, gc.IsNil)

	cmd := <-cmdChan
	c.Assert(cmd.Args, gc.DeepEquals, strings.Fields(s.paccmder.GetProxyCmd()))

	c.Assert(result, gc.Equals, initial)
}
