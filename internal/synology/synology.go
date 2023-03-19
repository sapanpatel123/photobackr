package synology

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/povsister/scp"
	"golang.org/x/crypto/ssh"

	log "github.com/sirupsen/logrus"
)

func BackupSynology(username string, privateKeyPath string, host string, src string, dest string, dryRun bool) error {
	if dryRun {
		fmt.Printf("This is a dry run! Normally, %s would be copied to %s on host %s", src, dest, host)
		return nil
	}

	sshConfig, err := getSSHConfig(username, privateKeyPath)
	if err != nil {
		log.Errorf("unable to get SSH config: %v", err)
		return err
	}

	client, err := scp.NewClient(host, sshConfig, &scp.ClientOption{})
	if err != nil {
		log.Errorf("unable to create new client: %v", err)
		return err
	}
	defer client.Close()

	if !isDir(src) {
		log.Errorf("source file does not exist, exiting: %v", err)
		return err
	}

	if err = client.CopyDirToRemote(src, dest, &scp.DirTransferOption{}); err != nil {
		log.Errorf("unable to copy to remote server: %v", err.Error())
		return err
	}

	return nil
}

func getSSHConfig(username string, privateKeyPath string) (*ssh.ClientConfig, error) {
	privatePEM, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Errorf("unable to read private key: %v", err)
		return nil, err
	}

	sshConfig, err := scp.NewSSHConfigFromPrivateKey(username, privatePEM)
	if err != nil {
		log.Errorf("unable to create ssh config: %v", err)
		return nil, err
	}

	return sshConfig, nil
}

func isDir(src string) bool {
	fileInfo, err := os.Stat(src)
	if err != nil {
		log.Fatalf("unable to retrieve file info: %v", err)
	}

	if fileInfo.IsDir() {
		return true
	}

	return false
}
