# GitHub App Installer

The `appinstall` command line tool lets you install app binaries from github user's releases. 
Once installed, simply select the version you require from the dropdown and start using the downloaded github user's app.  

<hr>

## Installation

`appinstall` is available for MacOS and Linux based operating systems.

### Homebrew

Installation for MacOS is the easiest with Homebrew. [If you do not have homebrew installed, click here](https://brew.sh/){:target="_blank"}. 


```ruby
brew install warrensbox/tap/appinstall
```

### Linux

Installation for Linux operation systems.

```sh
curl -L https://raw.githubusercontent.com/warrensbox/github-appinstaller/release/install.sh | bash
```

### Install from source

Alternatively, you can install the binary from the source [here](https://github.com/warrensbox/github-appinstaller/releases) 

<hr>

## How to use:
### Use dropdown menu to select version
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/aws-find-demo.gif" alt="drawing" style="width: 480px;"/>

1.  You can install and switch between different versions of github user's app by typing the command `appinstall install user/repo` on your terminal. 
2.  Select the version of binary by using the up and down arrow.
3.  Hit **Enter** to install the desired version.

The most recently selected versions are presented at the top of the dropdown.

### Upgrade current version
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/aws-find-demo2.gif" alt="drawing" style="width: 480px;"/>

1. You can also upgrade to latest version of the app.
2. For example, `appinstall upgrade user/repo`  to upgrade to a higher version of the app.
3. Hit **Enter** to upgrade.

### Uninstall GitHub app
<img align="center" src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/aws-find/aws-find-demo3.gif" alt="drawing" style="width: 480px;"/>

1. You can also uninstalled github user's.
2. For example, `appinstall upgrade user/repo` to uninstall to a higher version of the app.
3. Hit **Enter** to uninstall.

<hr>

## Issues

Please open  *issues* here: [New Issue](https://github.com/warrensbox/github-appinstaller/issues){:target="_blank"}

<hr>

See how to *upgrade*, *uninstall*, *troubleshoot* here:
[Additional Info](additional)