package cfg

import (
	"flag"
	"reflect"

	"github.com/grafana/dskit/flagext"
	"github.com/pkg/errors"
)

// Source is a generic configuration source. This function may do whatever is
// required to obtain the configuration. It is passed a pointer to the
// destination, which will be something compatible to `json.Unmarshal`. The
// obtained configuration may be written to this object, it may also contain
// data from previous sources.
type Source func(Cloneable) error

// Cloneable is a config which can be cloned into a flagext.Registerer
// Contract: the cloned value must not mutate the original.
type Cloneable interface {
	Clone() flagext.Registerer
}

var (
	ErrNotPointer = errors.New("dst is not a pointer")
)

// Unmarshal merges the values of the various configuration sources and sets them on
// `dst`. The object must be compatible with `json.Unmarshal`.
func Unmarshal(dst Cloneable, sources ...Source) error {
	if len(sources) == 0 {
		panic("No sources supplied to cfg.Unmarshal(). This is most likely a programming issue and should never happen. Check the code!")
	}
	if reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return ErrNotPointer
	}

	for _, source := range sources {
		if err := source(dst); err != nil {
			return err
		}
	}
	return nil
}

// DefaultUnmarshal is a higher level wrapper for Unmarshal that automatically parses flags and a .yaml file
func DefaultUnmarshal(dst Cloneable, args []string, fs *flag.FlagSet) error {
	return Unmarshal(dst,
		Defaults(fs),
		ConfigFileLoader(args, "config.file", true),
		Flags(args, fs),
	)
}
