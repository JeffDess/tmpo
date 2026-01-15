package update

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	githubAPIURL    = "https://api.github.com/repos/DylanDevelops/tmpo/releases/latest"
	checkTimeout    = 3 * time.Second
	connectTimeout  = 2 * time.Second
)

type ReleaseInfo struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	HTMLURL string `json:"html_url"`
}

type UpdateInfo struct {
	CurrentVersion string
	LatestVersion  string
	UpdateURL      string
	HasUpdate      bool
}

func IsConnectedToInternet() bool {
	_, err := net.LookupHost("api.github.com")
	return err == nil
}

func GetLatestVersion() (string, error) {
	client := &http.Client{
		Timeout: checkTimeout,
	}

	resp, err := client.Get(githubAPIURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest version: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var release ReleaseInfo
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", fmt.Errorf("failed to parse release info: %w", err)
	}

	return release.TagName, nil
}

func CompareVersions(current, latest string) int {
	current = strings.TrimPrefix(current, "v")
	latest = strings.TrimPrefix(latest, "v")

	// separate out normal version and prerelease
	currentCore, currentPrerelease := splitPrerelease(current)
	latestCore, latestPrerelease := splitPrerelease(latest)

	coreComparison := compareCoreVersions(currentCore, latestCore)

	// return early if core versions differ
	if coreComparison != 0 {
		return coreComparison
	}

	// prerelease is always less than stable
	if currentPrerelease != "" && latestPrerelease == "" {
		return -1
	}

	// stable is always greater than prerelease
	if currentPrerelease == "" && latestPrerelease != "" {
		return 1
	}

	// both are prereleases, compare alphabetically
	if currentPrerelease != "" && latestPrerelease != "" {
		if currentPrerelease < latestPrerelease {
			return -1
		}
		if currentPrerelease > latestPrerelease {
			return 1
		}
	}

	return 0
}

// splitPrerelease separates out version and prerelease tag
func splitPrerelease(version string) (core, prerelease string) {
	if idx := strings.Index(version, "-"); idx != -1 {
		return version[:idx], version[idx+1:]
	}
	return version, ""
}

// compareCoreVersions compares major.minor.patch numerically
func compareCoreVersions(current, latest string) int {
	currentParts := strings.Split(current, ".")
	latestParts := strings.Split(latest, ".")

	maxLen := len(currentParts)
	if len(latestParts) > maxLen {
		maxLen = len(latestParts)
	}

	for i := 0; i < maxLen; i++ {
		var currentVal, latestVal int

		if i < len(currentParts) {
			fmt.Sscanf(currentParts[i], "%d", &currentVal)
		}
		if i < len(latestParts) {
			fmt.Sscanf(latestParts[i], "%d", &latestVal)
		}

		if currentVal < latestVal {
			return -1
		}
		if currentVal > latestVal {
			return 1
		}
	}

	return 0
}

func CheckForUpdate(currentVersion string) (*UpdateInfo, error) {
	info := &UpdateInfo{
		CurrentVersion: currentVersion,
		HasUpdate:      false,
	}

	if !IsConnectedToInternet() {
		return nil, fmt.Errorf("no internet connection")
	}

	latestVersion, err := GetLatestVersion()
	if err != nil {
		return nil, err
	}

	info.LatestVersion = latestVersion
	info.UpdateURL = fmt.Sprintf("https://github.com/DylanDevelops/tmpo/releases/tag/%s", latestVersion)

	comparison := CompareVersions(currentVersion, latestVersion)
	if comparison < 0 {
		info.HasUpdate = true
	}

	return info, nil
}
