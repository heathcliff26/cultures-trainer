[![CI](https://github.com/heathcliff26/cultures-trainer/actions/workflows/ci.yaml/badge.svg?event=push)](https://github.com/heathcliff26/cultures-trainer/actions/workflows/ci.yaml)
[![Coverage Status](https://coveralls.io/repos/github/heathcliff26/cultures-trainer/badge.svg)](https://coveralls.io/github/heathcliff26/cultures-trainer)
[![Editorconfig Check](https://github.com/heathcliff26/cultures-trainer/actions/workflows/editorconfig-check.yaml/badge.svg?event=push)](https://github.com/heathcliff26/cultures-trainer/actions/workflows/editorconfig-check.yaml)
[![Generate go test cover report](https://github.com/heathcliff26/cultures-trainer/actions/workflows/go-testcover-report.yaml/badge.svg)](https://github.com/heathcliff26/cultures-trainer/actions/workflows/go-testcover-report.yaml)
[![Renovate](https://github.com/heathcliff26/cultures-trainer/actions/workflows/renovate.yaml/badge.svg)](https://github.com/heathcliff26/cultures-trainer/actions/workflows/renovate.yaml)

# Cultures Trainer

This Trainer is for the german version of the game. Hence the german resource names. Currently only tested with Northland, should also work with 8th Wonder of the World.

## Installation

### Download binary

1. Download the [latest release](https://github.com/heathcliff26/cultures-trainer/releases/latest)
2. Unpack the archive
3. Install the app for your user by running:
   - You can install it globally by running the script with `sudo`
```bash
./install.sh -i
```

#### Uninstalling

1. Switch to the folder where you have the installation script
2. Uninstall by running:
   - Run as `sudo` if you installed it globally
```bash
./install.sh -u
```
3. Delete the folder.


### Fedora Copr

The app is available as an rpm by using the fedora copr repository [heathcliff26/games](https://copr.fedorainfracloud.org/coprs/heathcliff26/games/).
1. Enable the copr repository
```bash
sudo dnf copr enable heathcliff26/games
```
2. Install the app
```bash
sudo dnf install cultures-trainer
```

## Images

### Main Window

![](images/main-window-dark.png#gh-dark-mode-only)
![](images/main-window-light.png#gh-light-mode-only)

### Setup Window

![](images/setup-window-dark.png#gh-dark-mode-only)
![](images/setup-window-light.png#gh-light-mode-only)
