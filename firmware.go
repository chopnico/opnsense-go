package opnsense

import (
	"encoding/json"
)

const (
	firmwareBasePath   = "core/firmware/"
	firmwareStatusPath = firmwareBasePath + "status"
	firmwareInfoPath   = firmwareBasePath + "info"
)

type FirmwareInfo struct {
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

type FirmwareStatus struct {
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

func (api *Api) FirmwareStatus() (*FirmwareStatus, error) {
	b, err := api.Do("GET", api.Host+firmwareStatusPath, nil)
	if err != nil {
		return nil, err
	}

	status := FirmwareStatus{}

	err = json.Unmarshal(b, &status)
	if err != nil {
		return nil, err
	}

	return &status, nil
}

func (api *Api) FirmwareInfo() (*FirmwareInfo, error) {
	b, err := api.Do("GET", api.Host+firmwareInfoPath, nil)
	if err != nil {
		return nil, err
	}

	info := FirmwareInfo{}

	err = json.Unmarshal(b, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
