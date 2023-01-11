// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package project

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ory/cli/cmd/cloudx/client"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/ory/x/osx"
)

func NewUpdateNamespaceConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "opl <project-id>",
		Aliases: []string{
			"namespaces-config",
		},
		Args:  cobra.NoArgs,
		Short: "Update the Ory Permission Language file in Ory Network",
		Example: `$ {{ .CommandPath }} ecaaa3cb-0730-4ee8-a6df-9553cdfeef89 \
	--file /path/to/namespace_config.ts

class Example implements Namespace {}
`,
		Long: "Update the Ory Permission Language file in Ory Network. Legacy namespace definitions will be overwritten.",
		RunE: func(cmd *cobra.Command, args []string) error {
			h, err := client.NewCommandHelper(cmd)
			if err != nil {
				return err
			}

			file := flagx.MustGetString(cmd, "file")

			data, err := osx.ReadFileFromAllSources(file)
			if err != nil {
				return err
			}
			patch := fmt.Sprintf(`/services/permission/config/namespaces={"location": "base64://%s"}`,
				base64.StdEncoding.EncodeToString(data))

			project, err := h.GetProject(flagx.MustGetString(cmd, "project"))
			if err != nil {
				return err
			}

			p, err := h.PatchProject(project.Id, nil, nil, []string{patch}, nil)
			if err != nil {
				return cmdx.PrintOpenAPIError(cmd, err)
			}

			cmdx.PrintJSONAble(cmd, outputConfig(p.Project.Services.Permission.Config))

			return h.PrintUpdateProjectWarnings(p)
		},
	}

	cmd.Flags().StringP("file", "f", "",
		"Configuration file (file://namespace_config.ts, https://example.org/namespace_config.ts, ...) to update the Ory Permission Language config")
	client.RegisterYesFlag(cmd.Flags())
	cmdx.RegisterFormatFlags(cmd.Flags())
	client.RegisterProjectFlag(cmd.Flags())

	return cmd
}
