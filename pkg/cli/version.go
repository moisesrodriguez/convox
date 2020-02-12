package cli

import (
	rck "github.com/convox/convox/pkg/rack"
	"github.com/convox/convox/sdk"
	"github.com/convox/stdcli"
)

func init() {
	register("version", "display version information", Version, stdcli.CommandOptions{
		Flags:    []stdcli.Flag{flagRack},
		Validate: stdcli.Args(0),
	})
}

func Version(rack sdk.Interface, c *stdcli.Context) error {
	c.Writef("client: <info>%s</info>\n", c.Version())

	r, err := rck.Current(c)
	if err != nil {
		c.Writef("server: <info>none</info>\n")
		return nil
	}

	rc, err := r.Client()
	if err != nil {
		return err
	}

	s, err := rc.SystemGet()
	if err != nil {
		return err
	}

	c.Writef("server: <info>%s</info>\n", s.Version)

	return nil
}
