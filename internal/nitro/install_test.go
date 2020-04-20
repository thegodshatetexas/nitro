package nitro

import (
	"reflect"
	"testing"
)

func TestInstallCorePackages(t *testing.T) {
	type args struct {
		name string
		php  string
	}
	tests := []struct {
		name    string
		args    args
		want    *Action
		wantErr bool
	}{
		{
			name: "can get commands to install PHP 7.4",
			args: args{
				name: "somename",
				php:  "7.4",
			},
			want: &Action{
				Type:       "exec",
				UseSyscall: false,
				Args:       []string{"exec", "somename", "--", "sudo", "apt-get", "install", "-y", "php7.4", "php7.4-mbstring", "php7.4-cli", "php7.4-curl", "php7.4-fpm", "php7.4-gd", "php7.4-intl", "php7.4-json", "php7.4-mysql", "php7.4-opcache", "php7.4-pgsql", "php7.4-zip", "php7.4-xml", "php7.4-soap", "php-xdebug", "php-imagick", "blackfire-agent", "blackfire-php"},
			},
			wantErr: false,
		},
		{
			name: "can get commands to install PHP 7.3",
			args: args{
				name: "somename",
				php:  "7.3",
			},
			want: &Action{
				Type:       "exec",
				UseSyscall: false,
				Args:       []string{"exec", "somename", "--", "sudo", "apt-get", "install", "-y", "php7.3", "php7.3-mbstring", "php7.3-cli", "php7.3-curl", "php7.3-fpm", "php7.3-gd", "php7.3-intl", "php7.3-json", "php7.3-mysql", "php7.3-opcache", "php7.3-pgsql", "php7.3-zip", "php7.3-xml", "php7.3-soap", "php-xdebug", "php-imagick", "blackfire-agent", "blackfire-php"},
			},
			wantErr: false,
		},
		{
			name: "can get commands to install PHP 7.2",
			args: args{
				name: "somename",
				php:  "7.2",
			},
			want: &Action{
				Type:       "exec",
				UseSyscall: false,
				Args:       []string{"exec", "somename", "--", "sudo", "apt-get", "install", "-y", "php7.2", "php7.2-mbstring", "php7.2-cli", "php7.2-curl", "php7.2-fpm", "php7.2-gd", "php7.2-intl", "php7.2-json", "php7.2-mysql", "php7.2-opcache", "php7.2-pgsql", "php7.2-zip", "php7.2-xml", "php7.2-soap", "php-xdebug", "php-imagick", "blackfire-agent", "blackfire-php"},
			},
			wantErr: false,
		},
		{
			name: "can get commands to install PHP 7.1",
			args: args{
				name: "somename",
				php:  "7.1",
			},
			want: &Action{
				Type:       "exec",
				UseSyscall: false,
				Args:       []string{"exec", "somename", "--", "sudo", "apt-get", "install", "-y", "php7.1", "php7.1-mbstring", "php7.1-cli", "php7.1-curl", "php7.1-fpm", "php7.1-gd", "php7.1-intl", "php7.1-json", "php7.1-mysql", "php7.1-opcache", "php7.1-pgsql", "php7.1-zip", "php7.1-xml", "php7.1-soap", "php-xdebug", "php-imagick", "blackfire-agent", "blackfire-php"},
			},
			wantErr: false,
		},
		{
			name: "can get commands to install PHP 7.0",
			args: args{
				name: "somename",
				php:  "7.0",
			},
			want: &Action{
				Type:       "exec",
				UseSyscall: false,
				Args:       []string{"exec", "somename", "--", "sudo", "apt-get", "install", "-y", "php7.0", "php7.0-mbstring", "php7.0-cli", "php7.0-curl", "php7.0-fpm", "php7.0-gd", "php7.0-intl", "php7.0-json", "php7.0-mysql", "php7.0-opcache", "php7.0-pgsql", "php7.0-zip", "php7.0-xml", "php7.0-soap", "php-xdebug", "php-imagick", "blackfire-agent", "blackfire-php"},
			},
			wantErr: false,
		},
		{
			name: "wrong version of php returns an error",
			args: args{
				name: "somename",
				php:  "7.9",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid name returns an error",
			args: args{
				name: "",
				php:  "7.9",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InstallPackages(tt.args.name, tt.args.php)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstallPackages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstallPackages() got = %v, want %v", got, tt.want)
			}
		})
	}
}
