package cfg

import (
	"os/user"
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
)

const configFile = "/.slack/config"

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
type SlackProfileRoot struct {
	Profile SlackProfile `json:"profile"`
}

// key value pairs of user profile settings
type SlackProfile struct {
	Message    string `json:"status_text"`
	Emoji      string `json:"status_emoji"`
	Expiration int    `json:"status_expiration"`
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

// Parse method on Config
func (cfg *Config) Parse(data []byte) error {
	return yaml.Unmarshal(data, &cfg)
}

// GetYaml method - retrieves yaml from config file
func (cfg *Config) GetYaml() ([]byte, error) {
	usr, _ := user.Current()
	cfgFile := fmt.Sprintf("%s%s", usr.HomeDir, configFile)
	data, err := ioutil.ReadFile(cfgFile)

	if err != nil {
		return nil, err
	}

	return data, nil

}

// Extracts auth token from config
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
func (cfg *Config) GetStatusProfileFromConfig(profileName string) SlackProfileRoot {
	var status SlackProfile

	for index, _ := range cfg.Status {
		if cfg.Status[index].Name == profileName {
			status.Emoji = cfg.Status[index].Emoji
			status.Message = cfg.Status[index].Message
			status.Expiration = 0
			break
		}
	}
	return SlackProfileRoot{status}
}