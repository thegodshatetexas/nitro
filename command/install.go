package command

import (
	"flag"
	"strings"

	"github.com/craftcms/nitro/scripts"
)

type InstallCommand struct {
	*CoreCommand

	// flags
	flagName            string
	flagPhpVersion      string
	flagDatabaseEngine  string
	flagDatabaseVersion string
}

func (c *InstallCommand) Synopsis() string {
	return "install software on machine"
}

func (c *InstallCommand) Help() string {
	return strings.TrimSpace(`
Usage: nitro install [options]
  This command install software on a virtual machine.
  
  Install software on a virtual machine and override the software version:
      $ nitro install -php-version=7.4 -database-engine=mysql -database-version=5.7
  
  Install software on a virtual machine with the default options:
      $ nitro install

  The default options will be the latest version of PHP (version 7.4) and MySQL 5.7.
`)
}

func (c *InstallCommand) Flags() *flag.FlagSet {
	s := flag.NewFlagSet("install", 127)

	// setup the flags for the command
	s.StringVar(&c.flagName, "name", "", "name of machine")
	s.StringVar(&c.flagPhpVersion, "php-version", "", "version of PHP to install")
	s.StringVar(&c.flagDatabaseEngine, "database-engine", "", "database engine (default: mysql)")
	s.StringVar(&c.flagDatabaseVersion, "database-version", "", "database engine version (default: 5.7)")

	// set defaults from the config file
	if c.Config.IsSet("name") {
		c.flagName = c.Config.GetString("name")
	}
	if c.Config.IsSet("php") {
		c.flagPhpVersion = c.Config.GetString("php")
	}
	if c.Config.IsSet("database.engine") {
		c.flagDatabaseEngine = c.Config.GetString("database.engine")
	}
	if c.Config.IsSet("database.version") {
		c.flagDatabaseVersion = c.Config.GetString("database.version")
	}

	return s
}

func (c *InstallCommand) Run(args []string) int {
	// parse the flags and check for any errors
	if err := c.Flags().Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	// get the php version packages
	packages := scripts.InstallPHP(c.flagPhpVersion)

	if err := c.Runner.Run([]string{"exec", c.flagName, "--", "bash -c", "'sudo apt install -y " + strings.Join(packages, " ") + "'"}); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	return 0
}
