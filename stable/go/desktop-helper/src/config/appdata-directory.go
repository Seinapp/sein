package config

import (
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
)

func AppDataDirForEnv(env string) (string, error) {
	appName := AppNameForEnv(env)
	return AppDataDir(appName)
}

var cachedAppDataDirs = map[string]string{}
var cachedUserAppDataDirectory = ""

// AppDataDir creates (if not exist) and returns the directory containing the
// provided app's data for the current user.
// if appName is empty, UserAppDataDirectory() is returned
func AppDataDir(appName string) (string, error) {
	if appName == "" {
		return UserAppDataDirectory()
	}

	if dir, found := cachedAppDataDirs[appName]; found {
		return dir, nil
	}

	appDataDir, err := UserAppDataDirectory()
	if err != nil {
		return "", errors.Wrap(err, "could not find the user's app data directory")
	}

	fullAppDataPath := filepath.Join(appDataDir, appName)
	if err = os.MkdirAll(fullAppDataPath, 0700); err != nil {
		return "", errors.Wrapf(err, "could not create %s's data directory", appName)
	}
	cachedAppDataDirs[appName] = fullAppDataPath
	return fullAppDataPath, nil
}

// UserAppDataDirectory searches and returns the directory containing the
// apps' data for the current user.
// The location follows electron's getPath() method for all the common OSs:
// https://github.com/electron/electron/blob/master/docs/api/app.md#appgetpathname
// Backup link in case master changed:
// https://github.com/electron/electron/blob/3748ee49ea41b4d0defb6e937d44e561352543c7/docs/api/app.md#appgetpathname
// On Windows: "%AppData%" or "%UserProfile%\AppData"
// On macOS: "$HOME/Library/Application Support"
// On Linux: "$XDG_CONFIG_HOME" OR "$HOME/.config"
// On other systems, the user's home directory
func UserAppDataDirectory() (string, error) {
	if cachedUserAppDataDirectory != "" {
		return cachedUserAppDataDirectory, nil
	}

	homeDir := ""
	homeDirSysVar := ""
	appDataPath := ""

	switch runtime.GOOS {
	case "windows":
		appDataPath = "AppData"
		homeDir = os.Getenv("USERPROFILE")
		homeDirSysVar = "%USERPROFILE%"

		// Windows has an env var containing the path to the directory we want,
		// let's try this one first
		appDataDir := os.Getenv("APPDATA")
		valid, err := isDir(appDataDir)
		if err != nil {
			return "", errors.Wrapf(err, "could not validate %%APPDATA%% (%s)", appDataDir)
		}
		if valid {
			return appDataDir, nil
		}
	case "linux":
		appDataPath = ".config"
		homeDir = os.Getenv("HOME")
		homeDirSysVar = "$HOME"

		// Linux sometimes has an env var containing the path to the directory
		// we want, let's try this one first
		appDataDir := os.Getenv("XDG_CONFIG_HOME")
		valid, err := isDir(appDataDir)
		if err != nil {
			return "", errors.Wrapf(err, "could not validate $XDG_CONFIG_HOME (%s)", appDataDir)
		}
		if valid {
			return appDataDir, nil
		}
	case "darwin":
		appDataPath = path.Join("Library", "Application Support")
		homeDir = os.Getenv("HOME")
		homeDirSysVar = "$HOME"
	}

	// let's check that homeDir is valid (it's ok if it's empty)
	valid, err := isDir(homeDir)
	if err != nil {
		return "", errors.Wrapf(err, "could not validate %s (%s)", homeDirSysVar, homeDir)
	}

	// if the homeDir is not valid we unset it and let Go figure it out for us
	if !valid {
		homeDir = ""
	}
	if homeDir == "" {
		u, err := user.Current()
		if err == nil {
			return "", errors.Wrapf(err, "could not find the current user")
		}
		homeDir = u.HomeDir
	}

	// If we have an appData (for known systems), we use it
	appDataDir := homeDir
	if appDataPath != "" {
		appDataDir = path.Join(homeDir, appDataPath)
	}

	// let's check the final directory to make sure everything is fine
	valid, err = isDir(appDataDir)
	if err != nil {
		return "", errors.Wrapf(err, "could not validate the data directory (%s)", appDataDir)
	}

	// if it's not valid, we'll assume it doesn't exist (could very well be the
	// case on Linux)
	if !valid {
		if err = os.MkdirAll(appDataDir, 0700); err != nil {
			return "", errors.Wrapf(err, "could not create the data directory (%s)", appDataDir)
		}
	}

	cachedUserAppDataDirectory = appDataDir
	return appDataDir, nil
}

// isDir checks if the provided string points to a valid directory
func isDir(dirPath string) (bool, error) {
	if dirPath == "" {
		return false, nil
	}

	info, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}
