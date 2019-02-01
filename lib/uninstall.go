package lib

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

//Uninstall : Install the provided version in the argument
func Uninstall(url string) string {

	/* get current user */
	usr, errCurr := user.Current()
	if errCurr != nil {
		log.Fatal(errCurr)
	}

	slice := strings.Split(url, "/")
	app := slice[1]

	installPath = fmt.Sprintf(installPath, app)
	bin := fmt.Sprintf(binLocation, app)

	installVersion = fmt.Sprintf(installVersion, app)
	installFile = fmt.Sprintf(installFile, app)

	/* set installation location */
	installLocation = usr.HomeDir + installPath

	/* set default binary path for app */
	installedBinPath = bin

	/* find app binary location if app is already installed*/
	cmd := NewCommand(app)
	next := cmd.Find()

	/* overrride installation default binary path if app is already installed */
	/* find the last bin path */
	for path := next(); len(path) > 0; path = next() {
		installedBinPath = path
	}

	/* check if selected version already downloaded */
	//fileExist := CheckFileExist(installLocation + installVersion + appversion)
	filesExist := GetListOfFile(installLocation)

	/* if selected version already exist, */
	if len(filesExist) > 0 {

		symlinkExist := CheckSymlink(installedBinPath)

		if symlinkExist {
			RemoveSymlink(installedBinPath)
		}
		return installLocation
	}

	return installLocation
}

//RemoveContents   remove all files in directory
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
