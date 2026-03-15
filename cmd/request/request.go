package request

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "request",
	Short: "Manage media requests",
	Long:  `Manage media requests including creating, approving, declining, and retrying requests.`,
}
