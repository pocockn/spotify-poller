package config

import "time"

type (
	// Duration wraps time.Duration.
	Duration struct {
		time.Duration
	}
)

// UnmarshalText satisfies the toml.TextUnmarshaler interface.
func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}