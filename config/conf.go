package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Env `yaml:"env"`
}

type Env struct {
	Title string `yaml:"title"`
	ID    string `yaml:"ID"`
}

func Configure(path string) *Config {
	// 外部からconfの中身を参照できるようにする
	var C Config

	// 設定ファイル名を記載
	viper.SetConfigName("config")

	// ファイルタイプ
	viper.SetConfigType("yml")

	// ファイルパスの設定。クロスプラットフォームで参照できるようにfilepathライブラリを使用
	viper.AddConfigPath(filepath.Join(path))
	viper.AddConfigPath("./config")

	// 環境変数から設定値を上書きできるように設定
	viper.AutomaticEnv()

	// conf読み取り
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file read error")
		fmt.Println(err)
		os.Exit(1)
	}

	// UnmarshalしてCにマッピング
	if err := viper.Unmarshal(&C); err != nil {
		fmt.Println("config file Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}

	return &C
}
