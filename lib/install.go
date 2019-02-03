package lib

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strings"

	"github.com/warrensbox/github-appinstaller/modal"
)

const (
	recentFile = "RECENT"
)

var (
	installFile      = "%s"
	installVersion   = "%s_"
	binLocation      = "/usr/local/bin/%s"
	installPath      = "/.%s/"
	installLocation  = "/tmp"
	installedBinPath = "/tmp"
)

//Install : Install the provided version in the argument
func Install(url string, appversion string, assests []modal.Repo) string {

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
	/* Create local installation directory if it does not exist */
	CreateDirIfNotExist(installLocation)

	goarch := runtime.GOARCH
	goos := runtime.GOOS
	urlDownload := ""

	for _, v := range assests {

		if v.TagName == appversion {
			if len(v.Assets) > 0 {
				for _, b := range v.Assets {

					matchedOS, _ := regexp.MatchString(goos, b.BrowserDownloadURL)
					matchedARCH, _ := regexp.MatchString(goarch, b.BrowserDownloadURL)
					if matchedOS && matchedARCH {
						urlDownload = b.BrowserDownloadURL
						break
					}
				}
			}
			break
		}
	}

	/* check if selected version already downloaded */
	//fileExist := CheckFileExist(installLocation + installVersion + appversion)
	fileExist := false
	filesExist := GetListOfFile(installLocation)

	for _, file := range filesExist {

		if matchedFile, _ := regexp.MatchString(appversion, file); matchedFile {
			fileExist = true
		}
	}

	/* if selected version already exist, */
	if fileExist {

		/* remove current symlink if exist*/
		symlinkExist := CheckSymlink(installedBinPath)

		if symlinkExist {
			RemoveSymlink(installedBinPath)
		}
		/* set symlink to desired version */
		CreateSymlink(installLocation+installVersion+appversion, installedBinPath)
		fmt.Printf("Switched app to version %q \n", appversion)
		return installLocation
	}

	/* remove current symlink if exist*/
	symlinkExist := CheckSymlink(installedBinPath)

	if symlinkExist {
		RemoveSymlink(installedBinPath)
	}

	/* if selected version already exist, */
	/* proceed to download it from the app release page */
	fileInstalled, _ := DownloadFromURL(installLocation, urlDownload)

	/* rename file to app version name - app_x.x.x */
	RenameFile(fileInstalled, installLocation+installVersion+appversion)

	err := os.Chmod(installLocation+installVersion+appversion, 0755)
	if err != nil {
		log.Println(err)
	}

	// /* set symlink to desired version */
	CreateSymlink(installLocation+installVersion+appversion, installedBinPath)
	fmt.Printf("Switched app to version %q \n", appversion)
	return installLocation
}

// AddRecent : add to recent file
func AddRecent(requestedVersion string, installLocation string) {

	semverRegex := regexp.MustCompile(`\d+(\.\d+){2}\z`)

	fileExist := CheckFileExist(installLocation + recentFile)
	if fileExist {
		lines, errRead := ReadLines(installLocation + recentFile)

		if errRead != nil {
			fmt.Printf("Error: %s\n", errRead)
			return
		}

		for _, line := range lines {
			if !semverRegex.MatchString(line) {
				RemoveFiles(installLocation + recentFile)
				CreateRecentFile(requestedVersion)
				return
			}
		}

		versionExist := VersionExist(requestedVersion, lines)

		if !versionExist {
			if len(lines) >= 3 {
				_, lines = lines[len(lines)-1], lines[:len(lines)-1]

				lines = append([]string{requestedVersion}, lines...)
				WriteLines(lines, installLocation+recentFile)
			} else {
				lines = append([]string{requestedVersion}, lines...)
				WriteLines(lines, installLocation+recentFile)
			}
		}

	} else {
		CreateRecentFile(requestedVersion)
	}
}

// GetRecentVersions : get recent version from file
func GetRecentVersions() ([]string, error) {

	fileExist := CheckFileExist(installLocation + recentFile)
	if fileExist {
		semverRegex := regexp.MustCompile(`\A\d+(\.\d+){2}\z`)

		lines, errRead := ReadLines(installLocation + recentFile)

		if errRead != nil {
			fmt.Printf("Error: %s\n", errRead)
			return nil, errRead
		}

		for _, line := range lines {
			if !semverRegex.MatchString(line) {
				RemoveFiles(installLocation + recentFile)
				return nil, errRead
			}
		}
		return lines, nil
	}
	return nil, nil
}

//CreateRecentFile : create a recent file
func CreateRecentFile(requestedVersion string) {
	WriteLines([]string{requestedVersion}, installLocation+recentFile)
}
