package nitro

import "github.com/craftcms/nitro/validate"

func Unmount(name, site string) (*Action, error) {
	if err := validate.MachineName(name); err != nil {
		return nil, err
	}

	return &Action{
		Type:       "umount",
		UseSyscall: false,
		Args:       []string{"umount", name + ":/app/sites/" + site},
	}, nil
}

func UnmountDir(name, target string) (*Action, error) {
	if err := validate.MachineName(name); err != nil {
		return nil, err
	}

	return &Action{
		Type:       "umount",
		UseSyscall: false,
		Args:       []string{"umount", name + ":" + target},
	}, nil
}
