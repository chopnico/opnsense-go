package opnsense

import (
	"encoding/json"
)

const (
	coreFirmwareBasePath    = "core/firmware/"
	coreFirmwareStatusPath  = coreFirmwareBasePath + "status"
	coreFirmwareInfoPath    = coreFirmwareBasePath + "info"
	coreFirmwareRunningPath = coreFirmwareBasePath + "running"
)

type CoreFirmwareInfo struct {
	ProductName    string `json:"product_name"`
	ProductVersion string `json:"product_version"`
	Package        []struct {
		Name       string `json:"name"`
		Version    string `json:"version"`
		Comment    string `json:"comment"`
		Flatsize   string `json:"flatsize"`
		Locked     string `json:"locked"`
		License    string `json:"license"`
		Repository string `json:"repository"`
		Origin     string `json:"origin"`
		Provided   string `json:"provided"`
		Installed  string `json:"installed"`
		Path       string `json:"path"`
		Configured string `json:"configured"`
	} `json:"package"`
	Plugin []struct {
		Name       string `json:"name"`
		Version    string `json:"version"`
		Comment    string `json:"comment"`
		Flatsize   string `json:"flatsize"`
		Locked     string `json:"locked"`
		License    string `json:"license"`
		Repository string `json:"repository"`
		Origin     string `json:"origin"`
		Provided   string `json:"provided"`
		Installed  string `json:"installed"`
		Path       string `json:"path"`
		Configured string `json:"configured"`
	} `json:"plugin"`
	Changelog []struct {
		Series  string `json:"series"`
		Version string `json:"version"`
		Date    string `json:"date"`
	} `json:"changelog"`
}

type CoreFirmwareStatus struct {
	Connection          string        `json:"connection"`
	DowngradePackages   []interface{} `json:"downgrade_packages"`
	DownloadSize        string        `json:"download_size"`
	LastCheck           string        `json:"last_check"`
	NewPackages         []interface{} `json:"new_packages"`
	OsVersion           string        `json:"os_version"`
	ProductName         string        `json:"product_name"`
	ProductVersion      string        `json:"product_version"`
	ReinstallPackages   []interface{} `json:"reinstall_packages"`
	RemovePackages      []interface{} `json:"remove_packages"`
	Repository          string        `json:"repository"`
	Updates             string        `json:"updates"`
	UpgradeMajorMessage string        `json:"upgrade_major_message"`
	UpgradeMajorVersion string        `json:"upgrade_major_version"`
	UpgradeNeedsReboot  string        `json:"upgrade_needs_reboot"`
	UpgradePackages     []interface{} `json:"upgrade_packages"`
	AllPackages         []interface{} `json:"all_packages"`
	StatusMsg           string        `json:"status_msg"`
	Status              string        `json:"status"`
}

type CoreFirmwareRunning struct {
	Status string `json:"status"`
}

func (api *Api) CoreFirmwareStatus() (*CoreFirmwareStatus, error) {
	b, err := api.Do("GET", coreFirmwareStatusPath, nil)
	if err != nil {
		return nil, err
	}

	status := CoreFirmwareStatus{}

	err = json.Unmarshal(b, &status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (api *Api) CoreFirmwareInfo() (*CoreFirmwareInfo, error) {
	b, err := api.Do("GET", coreFirmwareInfoPath, nil)
	if err != nil {
		return nil, err
	}

	info := CoreFirmwareInfo{}

	err = json.Unmarshal(b, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (api *Api) CoreFirmwareRunning() (*CoreFirmwareRunning, error) {
	b, err := api.Do("GET", coreFirmwareRunningPath, nil)
	if err != nil {
		return nil, err
	}

	running := CoreFirmwareRunning{}

	err = json.Unmarshal(b, &running)
	if err != nil {
		return nil, err
	}

	return &running, nil
}
