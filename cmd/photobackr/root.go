package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	conf string

	// rootCmd represents the photobackr command
	rootCmd = &cobra.Command{
		Use:   "photobackr <command> <flags>",
		Short: "Photo backup tool",
		Long:  `photobackdr is a cli tool to backup photos to Synology NAS and Google Photos`,
	}
)

// Execute adds all child commands to photobackr
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if conf != "" {
		viper.SetConfigFile(conf)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".photobackr")
	}

	if err := viper.ReadInConfig(); err == nil {
		log.Info("using config file: %s", viper.ConfigFileUsed())
		// log.Errorf("error using config file: %s", viper.ConfigFileUsed())
	}
}
