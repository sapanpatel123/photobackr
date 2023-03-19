package cmd

import (
	"github.com/sapanpatel123/photobackr/internal/synology"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var (
	all    bool
	src    string
	dest   string
	dryRun bool

	backupCmd = &cobra.Command{
		Use:   "backup",
		Short: "Backup a photo directory",
		Long:  "Backup a photo directory to both Synology NAS and Google Photos",
		Run:   runBackup,
	}
)

type synologyConfig struct {
	Username       string `mapstructure:"username"`
	PrivateKeyPath string `mapstructure:"privateKeyPath"`
	Host           string `mapstructure:"host"`
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.PersistentFlags().BoolVar(&all, "all", true, "If --all is used, photos will be backed up in both Synology NAS and Google Photos")
	backupCmd.PersistentFlags().StringVar(&src, "source", "photos", "Source folder to backup")
	backupCmd.PersistentFlags().StringVar(&dest, "destination", "photos2", "Destination folder to copy to")
	backupCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "if --dry-run is used, nothing will be copied")
}

func runBackup(cmd *cobra.Command, args []string) {
	var s synologyConfig

	if err := viper.UnmarshalKey("synology", &s); err != nil {
		log.Errorf("failed to unmarshal synology config: %v", err.Error())
	}

	// all, _ := rootCmd.PersistentFlags().GetBool("all")

	// if all {
	synology.BackupSynology(s.Username, s.PrivateKeyPath, s.Host, src, dest, dryRun)
	// synology.Backup(s.Username, .PrivateKeyPath)
	// }
}
