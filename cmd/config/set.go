package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Persist configuration to the config file",
	Long:  `Save the server URL and API key provided as flags to the CLI configuration file (~/.seer-cli.yaml by default).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.ConfigFileUsed() == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			configPath := filepath.Join(home, ".seer-cli.yaml")

			viper.SetConfigFile(configPath)

			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				if err := viper.SafeWriteConfigAs(configPath); err != nil {
					return fmt.Errorf("failed to create config file: %w", err)
				}
			}
		}

		if s, _ := cmd.Root().PersistentFlags().GetString("server"); s != "" {
			viper.Set("server", s)
		}

		if k, _ := cmd.Root().PersistentFlags().GetString("api-key"); k != "" {
			viper.Set("api_key", k)
		}

		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to save configuration: %w", err)
		}

		cmd.Printf("Configuration saved successfully to: %s\n", viper.ConfigFileUsed())
		return nil
	},
}

func init() {
	Cmd.AddCommand(configSetCmd)
}
