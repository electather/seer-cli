package issue

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all issues",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		req := apiClient.IssueAPI.IssueGet(ctx)
		if cmd.Flags().Changed("take") {
			v, _ := cmd.Flags().GetFloat32("take")
			req = req.Take(v)
		}
		if cmd.Flags().Changed("skip") {
			v, _ := cmd.Flags().GetFloat32("skip")
			req = req.Skip(v)
		}
		if cmd.Flags().Changed("sort") {
			v, _ := cmd.Flags().GetString("sort")
			req = req.Sort(v)
		}
		if cmd.Flags().Changed("filter") {
			v, _ := cmd.Flags().GetString("filter")
			req = req.Filter(v)
		}
		if cmd.Flags().Changed("requested-by") {
			v, _ := cmd.Flags().GetFloat32("requested-by")
			req = req.RequestedBy(v)
		}
		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "IssueGet")
	},
}

func init() {
	listCmd.Flags().Float32("take", 20, "Number of issues to return")
	listCmd.Flags().Float32("skip", 0, "Number of issues to skip")
	listCmd.Flags().String("sort", "added", "Sort by: added, modified")
	listCmd.Flags().String("filter", "open", "Filter by status: all, open, resolved")
	listCmd.Flags().Float32("requested-by", 0, "Filter by user ID")
	Cmd.AddCommand(listCmd)
}
