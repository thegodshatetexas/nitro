package nitro

func EnableXdebug(name, version string) []Command {
	return []Command{
		{
			Machine: name,
			Type:    "exec",
			Args:    []string{name, "--", "sudo", "phpenmod", "-v", version, "xdebug"},
		},
	}
}
