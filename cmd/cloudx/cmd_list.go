package cloudx

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/ory/x/cmdx"
)

func NewListCmd(parent *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   fmt.Sprintf("List resources"),
	}

	cmd.AddCommand(NewListProjectsCmd())
	cmd.AddCommand(NewListIdentityCmd(parent))

	RegisterConfigFlag(cmd.PersistentFlags())
	RegisterYesFlag(cmd.PersistentFlags())
	cmdx.RegisterNoiseFlags(cmd.PersistentFlags())
	cmdx.RegisterJSONFormatFlags(cmd.PersistentFlags())
	return cmd
}
