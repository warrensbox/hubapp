[![Build Status](https://travis-ci.org/warrensbox/github-appinstaller.svg?branch=master)](https://travis-ci.org/warrensbox/github-appinstaller)
[![Go Report Card](https://goreportcard.com/badge/github.com/warrensbox/github-appinstaller)](https://goreportcard.com/report/github.com/warrensbox/github-appinstaller)
[![CircleCI](https://circleci.com/gh/warrensbox/github-appinstaller/tree/master.svg?style=shield&circle-token=d74b0de145c45b1d0da97f817363c77350e1a121)](https://circleci.com/gh/warrensbox/github-appinstaller)

# Github App Installer

<img style="text-allign:center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/appinstall/smallerlogo.png" alt="drawing" width="110" height="140"/>


The `appinstall` command line tool lets you install app binaries from github user's repo -release. . 
Once installed, simply select the version you require from the dropdown and start using the github user's binary. 


See installation guide here: [appinstall installation](https://warrensbox.github.io/github-appinstaller/)

## Installation

`appinstall` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/). 


```ruby
brew install warrensbox/tap/appinstall
```

### Linux

Installation for other linux operation systems.

```sh
curl -L https://raw.githubusercontent.com/warrensbox/github-appinstaller/release/install.sh | bash
```

### Install from source

Alternatively, you can install the binary from source [here](https://github.com/warrensbox/github-appinstaller/releases) 

## How to use:
### Use dropdown menu to select version
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/appinstall/appinstall.gif" alt="drawing" style="width: 480px;"/>

1.  You can switch between different versions of github user's app by typing the command `appinstall install user/repo` on your terminal. 
2.  Select the version of binary by using the up and down arrow.
3.  Hit **Enter** to select the desired version.

The most recently selected versions are presented at the top of the dropdown.

### Upgrade current version
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/appinstall/appinstall-v4.gif" alt="drawing" style="width: 480px;"/>

1. You can also upgrade to latest version of the app.
2. For example, `appinstall upgrade user/repo` for version 0.10.5 of appinstall.
3. Hit **Enter** to upgrade.

### Uninstall Installed GitHub app
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/appinstall/appinstall-v4.gif" alt="drawing" style="width: 480px;"/>

1. You can also upgrade to latest version of the app.
2. For example, `appinstall upgrade user/repo` for version 0.10.5 of appinstall.
3. Hit **Enter** to upgrade.

## Additional Info

See how to *upgrade*, *uninstall*, *troubleshoot* here:[More info](https://warrensbox.github.io/github-appinstaller/additional)


## Issues

Please open  *issues* here:  [New Issue](https://github.com/warrensbox/github-appinstaller/issues)
