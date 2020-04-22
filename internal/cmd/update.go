package cmd

import (
	"github.com/spf13/cobra"

	"github.com/craftcms/nitro/config"
	"github.com/craftcms/nitro/internal/nitro"
)

var updateCommand = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade"},
	Short:   "Update a machine",
	RunE: func(cmd *cobra.Command, args []string) error {
		machine := "nitro-dev"
		if flagMachineName != "" {
			machine = flagMachineName
		}

		var actions []nitro.Action
		updateAction, err := nitro.Update(machine)
		if err != nil {
			return err
		}
		actions = append(actions, *updateAction)

		upgradeAction, err := nitro.Upgrade(machine)
		if err != nil {
			return err
		}
		actions = append(actions, *upgradeAction)

		return nitro.Run(nitro.NewMultipassRunner("multipass"), actions)
	},
}