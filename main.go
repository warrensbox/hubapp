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
	"os"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	simplelogger "github.com/mmmorris1975/simple-logger"
	"github.com/warrensbox/github-appinstaller/lib"
	"github.com/warrensbox/github-appinstaller/modal"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const (
	APIURL = "https://api.github.com/repos/%s/releases"
)

var version = "0.1.0\n"

var CLIENT_ID = "xxx"
var CLIENT_SECRET = "xxx"

var (
	debugFlag   *bool
	versionFlag *bool
	helpFlag    *bool
	action      *string
	giturl      *string
	log         *simplelogger.Logger
)

func init() {

	const (
		cmdDesc         = "Install github app binaries on your local machine. Ex: appinstall installwarrensbox/aws-find"
		versionFlagDesc = "Displays the version of appinstall"
		actionArgDesc   = "Provide action needed. Ex: install, update, or uninstall"
		giturlArgDesc   = "Provide giturl in user/repo format. Ex: warrensbox/aws-find"
		debugFlagDesc   = "Provide debug output"
	)

	debugFlag = kingpin.Flag("debug", debugFlagDesc).Short('d').Bool()
	versionFlag = kingpin.Flag("version", versionFlagDesc).Short('v').Bool()
	action = kingpin.Arg("action", actionArgDesc).String()
	giturl = kingpin.Arg("user/repo", giturlArgDesc).String()

	log = simplelogger.StdLogger
	log.SetLevel(simplelogger.INFO)
	log.SetFlags(0)

}

func main() {

	var client modal.Client

	client.ClientID = CLIENT_ID
	client.ClientSecret = CLIENT_SECRET

	kingpin.CommandLine.Interspersed(false)
	kingpin.Parse()

	if *versionFlag {
		fmt.Println(version)
	}

	if *debugFlag {
		log.SetLevel(simplelogger.DEBUG)
	}

	semverRegex := regexp.MustCompile(`^\w+(-)?\w+\/\w+(-)?\w+?$`)
	if semverRegex.MatchString(*giturl) == false && *versionFlag == false {
		log.Infof("Invalid repo format. Must be user/repo. Ex: appinstall install warrensbox/aws-find ")
		os.Exit(1)
	}

	apiURL := fmt.Sprintf(APIURL, *giturl)

	switch *action {
	case "install":
		log.Debug("Action -> install")
		ghlist, assets := lib.GetAppList(apiURL, &client)
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
			log.Info("Prompt failed %v\n", errPrompt)
			os.Exit(1)
		}

		installLocation := lib.Install(*giturl, ghversion, assets)
		lib.AddRecent(ghversion, installLocation)

	case "upgrade":
		log.Debug("Action -> upgrade")

		latestVersion, assets, err := lib.GetAppLatestVersion(apiURL, &client)
		if err != nil {
			log.Error("Could not get the latest version. Try `appinstall install user/repo`")
			os.Exit(1)
		}
		installLocation := lib.Install(*giturl, latestVersion, assets)
		lib.AddRecent(latestVersion, installLocation)

	case "uninstall":
		log.Debug("Action -> uninstall")
		installLocation := lib.Uninstall(*giturl)
		errContent := lib.RemoveContents(installLocation)
		if errContent != nil {
			log.Debug("Unable to remove files. Files might not have existed.")
			os.Exit(0)
		}
		slice := strings.Split(*giturl, "/")
		app := slice[1]
		log.Infof("Uninstalled %s\n", app)
	default:
		if *versionFlag == false {
			fmt.Println("Unknown action. See help. Ex: appinstall --help")
		}
	}
}
