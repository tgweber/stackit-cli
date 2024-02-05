package member

import (
	"github.com/stackitcloud/stackit-cli/internal/cmd/organization/member/add"
	"github.com/stackitcloud/stackit-cli/internal/cmd/organization/member/list"
	"github.com/stackitcloud/stackit-cli/internal/cmd/organization/member/remove"
	"github.com/stackitcloud/stackit-cli/internal/pkg/args"
	"github.com/stackitcloud/stackit-cli/internal/pkg/utils"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member",
		Short: "Provides functionality regarding organization members",
		Long:  "Provides functionality regarding organization members.",
		Args:  args.NoArgs,
		Run:   utils.CmdHelp,
	}
	addSubcommands(cmd)
	return cmd
}

func addSubcommands(cmd *cobra.Command) {
	cmd.AddCommand(add.NewCmd())
	cmd.AddCommand(list.NewCmd())
	cmd.AddCommand(remove.NewCmd())
}
