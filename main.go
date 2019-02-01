package main

/*
* Version 0.3.0
* Compatible with Mac OS X ONLY
 */

/*** OPERATION WORKFLOW ***/
/*
* 1- Make GET call to receive release json from github
* 2- Download released app from archive
* 3- Rename the file from `appinstall` to `appinstall_version`
* 4- Read the existing symlink for app (Check if it's a homebrew symlink)
* 6- Remove that symlink (Check if it's a homebrew symlink)
* 7- Create new symlink to binary  `github app`
 */

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/warrensbox/github-appinstaller/lib"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	APIURL = "https://api.github.com/repos/%s/releases"
)

var version = "0.1.0\n"

var (
	versionFlag *bool
	helpFlag    *bool
	action      *string
	giturl      *string
)

func init() {

	const (
		cmdDesc         = "Install github app binaries on your local machine. Ex: appinstall installwarrensbox/aws-find"
		versionFlagDesc = "Displays the version of appinstall"
		actionArgDesc   = "Provide action needed. Ex: install, update, or uninstall"
		giturlArgDesc   = "Provide giturl in user/repo format. Ex: warrensbox/aws-find"
	)

	versionFlag = kingpin.Flag("version", versionFlagDesc).Short('v').Bool()
	action = kingpin.Arg("action", actionArgDesc).Required().String()
	giturl = kingpin.Arg("user/repo", giturlArgDesc).Required().String()

}

func main() {

	kingpin.CommandLine.Interspersed(false)
	kingpin.Parse()
	apiURL := fmt.Sprintf(APIURL, *giturl)

	switch *action {
	case "install":
		fmt.Println("You said install") //remove later
		ghlist, assets := lib.GetAppList(apiURL)
		recentVersions, _ := lib.GetRecentVersions()
		ghlist = append(recentVersions, ghlist...)
		ghlist = lib.RemoveDuplicateVersions(ghlist) //remove duplicate version

		/* prompt user to select version of github app */
		prompt := promptui.Select{
			Label: "Select app version",
			Items: ghlist,
		}

		_, ghversion, errPrompt := prompt.Run()

		if errPrompt != nil {
			log.Printf("Prompt failed %v\n", errPrompt)
			os.Exit(1)
		}

		installLocation := lib.Install(*giturl, ghversion, assets)
		lib.AddRecent(ghversion, installLocation)

	case "update":
		fmt.Println("You said update") //remove later

		latestVersion, assets, err := lib.GetAppLatestVersion(apiURL)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		installLocation := lib.Install(*giturl, latestVersion, assets)
		lib.AddRecent(latestVersion, installLocation)

	case "uninstall":
		fmt.Println("You said uninstall") //remove later
		installLocation := lib.Uninstall(*giturl)
		lib.RemoveContents(installLocation)
	default:
		fmt.Println("Unknown action. See help. Ex: appinstall --help")
	}
}
