package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var errInternalError = errors.New("client quit unexpectedly")

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS play",
	Long:  ``,
}


func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".cobra")          // name of config file (without extension)
	viper.AddConfigPath(os.Getenv("HOME")) // adding home directory as first search path
	viper.AutomaticEnv()                   // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
