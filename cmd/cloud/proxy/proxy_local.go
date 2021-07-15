package proxy

import (
	"fmt"
	"github.com/ory/cli/cmd/cloud/remote"
	"github.com/ory/x/flagx"
	"github.com/spf13/cobra"
)

func NewProxyLocalCmd() *cobra.Command {
	proxyCmd := &cobra.Command{
		Use:   "local [upstream]",
		Short: "Develop an application locally and integrate it with Ory",
		Args:  cobra.ExactArgs(1),
		Long: fmt.Sprintf(`This command starts a reverse proxy which can be deployed in front of your application. This works best on local (your computer) environments, for example when developing a React, NodeJS, Java, PHP app.

To require login before accessing paths in your application, use the --%[1]s flag:

	$ ory proxy local --port 4000 --%[1]s /members --%[1]s /admin \
		http://localhost:3000

%[2]s`, ProtectPathsFlag, jwtHelp),
		/*
		   The --%s values support regular expression templating, meaning that you can use regular expressions within "<>":

		   	$ ory proxy http://localhost:3000 --allow --%s "http://localhost:3000/<(login|dashboard)>" --%s "http://localhost:3000/<([0-9]{3})>"

		   The supported Regular Expression Syntax is RE2 and documented at: https://golang.org/pkg/regexp/
		   To test your Regular Expression, head over to https://regex101.com and select "Golang" on the left.
		*/
		RunE: func(cmd *cobra.Command, args []string) error {
			conf := &config{
				port:                flagx.MustGetInt(cmd, PortFlag),
				protectPathPrefixes: flagx.MustGetStringSlice(cmd, ProtectPathsFlag),
				noCert:              flagx.MustGetBool(cmd, NoCertInstallFlag),
				noOpen:              flagx.MustGetBool(cmd, NoOpenFlag),
				apiEndpoint:         flagx.MustGetString(cmd, remote.FlagAPIEndpoint),
				consoleEndpoint:     flagx.MustGetString(cmd, remote.FlagConsoleAPI),
				isLocal:             true,
				upstream: args[0],
			}

			return run(cmd, conf)
		},
	}

	proxyCmd.Flags().Int(PortFlag, portFromEnv(), "The port the proxy should listen on.")
	proxyCmd.Flags().Bool(NoCertInstallFlag, false, "If set will not try to add the HTTPS certificate to your certificate store.")
	proxyCmd.Flags().StringSlice(ProtectPathsFlag, []string{}, "Require authentication before accessing these paths.")
	proxyCmd.Flags().Bool(NoOpenFlag, false, "Do not open the browser when the proxy starts.")
	remote.RegisterClientFlags(proxyCmd.PersistentFlags())
	return proxyCmd
}
