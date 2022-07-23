package base

import (
	"errors"

	"github.com/oklookat/toanother/core/datadir"
)

const CONFIG_NAME = "./config.yml"

var ConfigFile = &Config{}

type Config struct {
	YandexMusic YandexMusicSettings `json:"YandexMusic" yaml:"YandexMusic"`
	Spotify     SpotifySettings     `json:"Spotify" yaml:"Spotify"`
}

// create/read config file.
func (c *Config) Init() (err error) {
	// settings.
	isSettingsExists, err := datadir.IsFileExists(CONFIG_NAME)
	if err != nil {
		return err
	}
	if isSettingsExists {
		if err = datadir.GetStructByFile(CONFIG_NAME, true, c); err != nil {
			return err
		}
	} else {
		// set default settings.
		if err = datadir.WriteFileStruct(CONFIG_NAME, true, c); err != nil {
			return err
		}
	}
	ConfigFile = c
	return
}

// write config struct to file.
func (c *Config) WriteToFile() (err error) {
	return datadir.WriteFileStruct(CONFIG_NAME, true, c)
}

type YandexMusicSettings struct {
	Login string `json:"login" yaml:"login"`
}

// write to config file.
func (c *YandexMusicSettings) Apply() (err error) {
	if ConfigFile == nil {
		err = errors.New("config not initialized")
		return
	}
	ConfigFile.YandexMusic = *c
	return ConfigFile.WriteToFile()
}

type SpotifySettings struct {
	ID     string `json:"id" yaml:"id"`
	Secret string `json:"secret" yaml:"secret"`
}

// write to config file.
func (c *SpotifySettings) Apply() (err error) {
	if ConfigFile == nil {
		err = errors.New("config not initialized")
		return
	}
	ConfigFile.Spotify = *c
	return ConfigFile.WriteToFile()
}
