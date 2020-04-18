package manifest

import (
	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type rmOpts struct {
	target   string
}

func newRmManifestListCommand(dockerCli command.Cli) *cobra.Command {
	opts := rmOpts{}

	cmd := &cobra.Command{
		Use:   "rm MANIFEST_LIST",
		Short: "Delete a manifest list from local storage",
		Args:  cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.target = args[0]
			return runRm(dockerCli, opts)
		},
	}

	return cmd
}

func runRm(dockerCli command.Cli, opts rmOpts) error {

	targetRef, err := normalizeReference(opts.target)
	if err != nil {
		return err
	}

	manifests, err := dockerCli.ManifestStore().GetList(targetRef)
	if err != nil {
		return err
	}

	if len(manifests) == 0 {
		return errors.Errorf("%s not found", targetRef)
	}

	return dockerCli.ManifestStore().Remove(targetRef)
}
