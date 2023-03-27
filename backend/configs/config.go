package configs

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	"github.com/KScaesar/jubo-homework/backend/util/database"
	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

type ConfigCenter[T any] interface {
	GetConfig() (*T, error)
}

func NewFileConfigCenter[T any](configPath string) (*FileConfigCenter[T], error) {
	dir := filepath.Dir(configPath)
	name := filepath.Base(configPath)

	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName(name)
	vp.AddConfigPath(dir)

	if err := vp.ReadInConfig(); err != nil {
		return nil, errors.Join3rdPartyWithMsg(errors.ErrSystem, err, "read config file")
	}

	return &FileConfigCenter[T]{
		viper: vp,
	}, nil
}

type FileConfigCenter[T any] struct {
	viper *viper.Viper
}

func (center *FileConfigCenter[T]) GetConfig() (*T, error) {
	cfg := new(T)

	option := func(c *mapstructure.DecoderConfig) { c.TagName = "configs" }
	err := center.viper.Unmarshal(cfg, option)
	if err != nil {
		return nil, errors.Join3rdPartyWithMsg(errors.ErrSystem, err, "unmarshal config file")
	}

	return cfg, nil
}

func NewProjectConfig() (*ProjectConfig, error) {
	workDir, exist := os.LookupEnv("WorkDir")
	if !exist {
		return nil, errors.WrapWithMessage(errors.ErrSystem, "environment variable ${WorkDir} not exist")
	}
	configPath, exist := os.LookupEnv("ConfigPath")
	if !exist {
		return nil, errors.WrapWithMessage(errors.ErrSystem, "environment variable ${ConfigPath} not exist")
	}

	configPath = filepath.Join(workDir, configPath)

	configCenter, err := NewFileConfigCenter[ProjectConfig](configPath)
	if err != nil {
		return nil, err
	}

	config, err := configCenter.GetConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}

type ProjectConfig struct {
	ServerPort string             `configs:"server_port"`
	Pgsql      *database.DbConfig `configs:"pgsql"`
}
