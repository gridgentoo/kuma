package version

import (
	"fmt"

	"github.com/spf13/cobra"

	kuma_version "github.com/Kong/kuma/pkg/version"
)

func NewVersionCmd() *cobra.Command {
	args := struct {
		detailed bool
	}{}
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Long:  `Print version.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			buildInfo := kuma_version.Build

			if args.detailed {
				cmd.Println(fmt.Sprintf("Version:    %s", buildInfo.Version))
				cmd.Println(fmt.Sprintf("Git Tag:    %s", buildInfo.GitTag))
				cmd.Println(fmt.Sprintf("Git Commit: %s", buildInfo.GitCommit))
				cmd.Println(fmt.Sprintf("Build Date: %s", buildInfo.BuildDate))
			} else {
				cmd.Println(buildInfo.Version)
			}

			return nil
		},
	}
	// flags
	cmd.PersistentFlags().BoolVarP(&args.detailed, "detailed", "a", false, "Print detailed version")
	return cmd
}
