package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	server  string
	apiKey  string
)

var rootCmd = &cobra.Command{
	Use:   "seer-cli",
	Short: "A CLI to interact with the Seer API",
	Long:  `A command line interface to call endpoints defined in the Seer OpenAPI specification.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip server requirement for config command
		if cmd.Root() != cmd && (cmd.Name() == "config" || cmd.Parent().Name() == "config" || cmd.Name() == "help" || cmd.Name() == "completion") {
			return nil
		}
		// Ensure server is set
		if viper.GetString("server") == "" {
			return fmt.Errorf("server URL is required. Set it via --server flag, SEER_SERVER env var, or in the config file")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior when no subcommands are provided
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.seer-cli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "", "Seer server URL (e.g., http://localhost:5055)")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", "", "Seer API Key")

	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("api_key", rootCmd.PersistentFlags().Lookup("api-key"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".seer-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".seer-cli")
	}

	viper.SetEnvPrefix("SEER")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
