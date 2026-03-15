package service

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "service",
	Short: "Query Radarr and Sonarr service information",
	Long:  `Retrieve server, quality profile, and root folder details for Radarr and Sonarr.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
