package cfg

import (
	"os/user"
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
)

const configFile = "/.slack/config"

// type CurrentUser interface {
// 	Current() (*user.User, error)
// }

// type Parser struct{
// 	CurrentUser CurrentUser
// }

// root configuration
type Config struct {
	Default   string            `yaml:"default"`
	Workspace []ConfigWorkspace `yaml:"workspaces"`
	Status    []ConfigStatus    `yaml:"status"`
	file	  string //absolute config file location
}

// slack workspace as defined by the configuration file
type ConfigWorkspace struct {
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}

// preset user status from config file
type ConfigStatus struct {
	Name       string `yaml:"name"`
	Message    string `yaml:"message"`
	Emoji      string `yaml:"emoji"`
	Expiration int
}

// root type for Slack user profile
type SlackProfileRoot struct {
	Profile SlackProfile `json:"profile"`
}

// key value pairs of user profile settings
type SlackProfile struct {
	Message    string
	Emoji      string
	Expiration int   
}

// returns new Config type
func New() (*Config, error) {
	var cfg = &Config{}

	yaml, err := cfg.GetYaml()
	if err != nil {
		return nil, err
	}
	
	err = cfg.Parse(yaml) 
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) Parse(data []byte) error {
	return yaml.Unmarshal(data, cfg)
}

func (cfg *Config) GetYaml() ([]byte, error) {
	usr, _ := user.Current()
	cfgFile := fmt.Sprintf("%s%s", usr.HomeDir, configFile)
	data, err := ioutil.ReadFile(cfgFile)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (cfg *Config) GetTokenFromConfig(workspaceName string) string {
	var token string

	if workspaceName == "" {
		workspaceName = cfg.Default
	}

	for index, _ := range cfg.Workspace {
		if cfg.Workspace[index].Name == workspaceName {
			token = cfg.Workspace[index].Token
			break
		}
	}

	return token
}

// returns auth token from config file
func (cfg *Config) GetToken(workspaceName string) string {
	return cfg.GetTokenFromConfig(workspaceName)
}

//returns selected status profile from config file
func (cfg *Config) GetStatusProfileFromConfig(profileName string) SlackProfile {
	var status SlackProfile

	for index, _ := range cfg.Status {
		if cfg.Status[index].Name == profileName {
			status.Emoji = cfg.Status[index].Emoji
			status.Message = cfg.Status[index].Message
			status.Expiration = 0
			break
		}
	}
	return status
}