package main

import (
	"context"

	"github.com/conductorone/baton-sdk/pkg/cli"
)

// config defines the external configuration required for the connector to run.
type config struct {
	cli.BaseConfig `mapstructure:",squash"` // Puts the base config options in the same place as the connector options

	// flags to communicate with the NPM API?
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Email string `mapstructure:"email"`
}

// validateConfig is run after the configuration is loaded, and should return an error if it isn't valid.
func validateConfig(ctx context.Context, cfg *config) error {
	if cfg.Username == nil || cfg.Password == nil || cfg.Email == nil {
		return fmt.Errorf("bad configuration")
	}

	return nil
}

// 2:50 - pass necessary flags
func setCommandFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("username", "", "Username: ($BATON_USERNAME)")
	cmd.PersistentFlags().String("password", "", "Password: ($BATON_PASSWORD)")
	cmd.PersistentFlags().String("email", "", "Username: ($BATON_EMAIL)")
}
