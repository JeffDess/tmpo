package milestones

import "github.com/spf13/cobra"

func MilestoneCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "milestone",
		Short: "Manage milestones",
		Long:  `Manage milestones to group time entries into time-boxed periods (sprints, releases, phases).`,
	}

	cmd.AddCommand(StartCmd())
	cmd.AddCommand(FinishCmd())
	cmd.AddCommand(StatusCmd())
	cmd.AddCommand(ListCmd())

	return cmd
}
