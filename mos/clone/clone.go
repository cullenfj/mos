package clone

import (
	"context"
	"os"

	"cesanta.com/mos/build"
	"cesanta.com/mos/dev"
	"cesanta.com/mos/version"
	"github.com/cesanta/errors"
	flag "github.com/spf13/pflag"
)

func Clone(ctx context.Context, devConn *dev.DevConn) error {
	var m build.SWModule

	args := flag.Args()
	switch len(args) {
	case 1:
		return errors.Errorf("source location is required")
	case 2:
		m.Location = args[1]
	case 3:
		m.Name = args[2]
		m.Location = args[1]
	default:
		return errors.Errorf("extra arguments")
	}

	d, err := m.PrepareLocalDir(".", os.Stderr, false, /* deleteIfFailed */
		version.GetMosVersion() /* defaultVersion */, 0 /* pullInterval */, 0 /* cloneDepth */)

	// Chdir is needed for the Web UI mode: immediately go into the cloned repo.
	if m.Name != "" {
		os.Rename(d, m.Name)
		os.Chdir(m.Name)
	} else {
		os.Chdir(d)
	}

	return err
}