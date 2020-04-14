package support

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func GetAssetsInstallPath() (string, error) {
	//TODO switch on other places to store sound files
	return "/usr/local/etc/drummachine/sounds/", nil
}

func GetSoundFilePath(kit string) (string, error) {
	installPath, err := GetAssetsInstallPath()
	if err != nil {
		return "", err
	}
	soundFilePath := fmt.Sprintf("%s/%s.sf2", installPath, kit)

	_, err = os.Stat(soundFilePath)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("sound file for kit %s does not exist", kit))
	}

	return soundFilePath, nil
}