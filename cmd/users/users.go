package users

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "users",
	Short: "Manage users and user settings",
	Long:  `Manage users, their permissions, settings, linked accounts, and push subscriptions.`,
}

func init() {
	Cmd.AddCommand(passwordCmd, settingsCmd, linkedAccountsCmd, pushSubscriptionsCmd)
}
