package lib

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/mholt/archiver"
	"github.com/warrensbox/hubapp/modal"
)

const (
	recentFile = "RECENT"
)

var (
	homeDir          = os.Getenv("HOME")
	installFile      = "%s"
	installVersion   = "%s_"
	binLocation      = "%s/.local/bin/%s"
	installPath      = "/.%s/"
	installLocation  = "/tmp"
	installedBinPath = "/tmp"
)

// Install : Install the provided version in the argument
func Install(url string, appversion string, assests []modal.Repo) string {

	/* get current user */
	usr, errCurr := user.Current()
	if errCurr != nil {
		log.Fatal(errCurr)
	}

	slice := strings.Split(url, "/")
	app := slice[1]

	installPath = fmt.Sprintf(installPath, app)
	bin := fmt.Sprintf(binLocation, homeDir, app)

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

	/* identify arch and goos, prepare regex to ignore case */
	goarch := runtime.GOARCH
	goarchrgx := "(?i)" + goarch
	goos := runtime.GOOS
	goosrgx := "(?i)" + goos
	urlDownload := ""

	/* some binaries use x86_64 to identify their build instead of amd64 */
	if goarch == "amd64" {
		goarchrgx = goarchrgx + "|x86_64"
	}

	for _, v := range assests {

		/* some github release tags include v in their server tag name */
		/* if "v" is included in the tag name, it is removed" */
		semverRegex := regexp.MustCompile(`\Av\d+(\.\d+){2}\z`)
		version := v.TagName

		if semverRegex.MatchString(v.TagName) {
			trimstr := strings.Trim(v.TagName, "v") /* remove lowercase v */
			version = trimstr
		}

		if version == appversion {
			if len(v.Assets) > 0 {
				for _, b := range v.Assets {

					if b.BrowserDownloadURL != "" {
						matchedOS, _ := regexp.MatchString(goosrgx, b.BrowserDownloadURL)
						matchedARCH, _ := regexp.MatchString(goarchrgx, b.BrowserDownloadURL)
						if matchedOS && matchedARCH {
							urlDownload = b.BrowserDownloadURL
							break
						}
					} else {
						/* no download assets found */
						break
					}

				}
			}
			break
		}
	}

	/* no downloadable binaries with the proper name in the assets */
	if urlDownload == "" {
		log.Fatal(`
		No binaries found matching your computer's operating system or architecture. 
		Please verify user's releases for a valid binary. 
		The release binary should have the operating system and architecture in it's name.
		`)
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
		fmt.Println("File exist")
		/* remove current symlink if exist*/
		symlinkExist := CheckSymlink(installedBinPath)

		if symlinkExist {
			RemoveSymlink(installedBinPath)
		}
		/* set symlink to desired version */
		CreateSymlink(installLocation+installVersion+appversion, installedBinPath)
		fmt.Printf("Switched app to version %s \n", appversion)
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

	ext := filepath.Ext(fileInstalled)

	/* if file is compressed, get extension */
	if ext == ".gz" || ext == ".tar.gz" || ext == ".gip" || ext == ".zip" {

		tmpFile := fmt.Sprintf(installLocation+"files_%s", appversion)

		errArchive := archiver.Unarchive(fileInstalled, tmpFile)
		if errArchive != nil {
			log.Println(errArchive)
		}

		exist := CheckDirHasBin(tmpFile, app)

		if exist {
			RemoveAFile(fileInstalled)
			fileInstalled = fmt.Sprintf(tmpFile+"/%s", app)

		} else {
			log.Fatal("Unable to download and create a symlink to the downloaded binary")
			os.Exit(1)
		}

	}

	/* rename file to app version name - app_x.x.x */
	RenameFile(fileInstalled, installLocation+installVersion+appversion)

	err := os.Chmod(installLocation+installVersion+appversion, 0755)
	if err != nil {
		log.Println(err)
	}

	// /* set symlink to desired version */
	CreateSymlink(installLocation+installVersion+appversion, installedBinPath)
	fmt.Printf("Switched app to version %s \n", appversion)
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

// CreateRecentFile : create a recent file
func CreateRecentFile(requestedVersion string) {
	WriteLines([]string{requestedVersion}, installLocation+recentFile)
}
