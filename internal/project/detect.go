package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/DylanDevelops/tmpo/internal/settings"
)

func DetectProject() (string, error) {
	configPath, err := FindTmporc()
	if err == nil && configPath != "" {
		dir := filepath.Dir(configPath)

		return filepath.Base(dir), nil
	}

	gitName, err := GetGitRepoName()
	if err == nil && gitName != "" {
		return gitName, nil
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	return filepath.Base(cwd), nil
}


func DetectConfiguredProject() (string, error) {
	return DetectConfiguredProjectWithOverride("")
}

// DetectConfiguredProjectWithOverride detects the project (priority: specific project name > .tmporc > git repo > directory name)
func DetectConfiguredProjectWithOverride(explicitProject string) (string, error) {
	// first priority: --project flag
	if explicitProject != "" {
		registry, err := settings.LoadProjects()
		if err != nil {
			return "", fmt.Errorf("failed to load projects registry: %w", err)
		}

		if !registry.Exists(explicitProject) {
			return "", fmt.Errorf("project '%s' not found in global registry", explicitProject)
		}

		return explicitProject, nil
	}

	// second priority: .tmporc configuration
	if cfg, _, err := settings.FindAndLoad(); err == nil && cfg != nil {
		if cfg.ProjectName != "" {
			return cfg.ProjectName, nil
		}
	}

	// third priority: directory-based detection
	return DetectProject()
}

func FindTmporc() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		tmporc := filepath.Join(dir, ".tmporc")
		if _, err := os.Stat(tmporc); err == nil {
			return tmporc, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}

		dir = parent
	}

	return "", nil
}

func GetGitRepoName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	gitRoot := strings.TrimSpace(string(output))

	return filepath.Base(gitRoot), nil
}

func IsInGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	err := cmd.Run()

	return err == nil
}

func GetGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("not in a git repository")
	}

	return strings.TrimSpace(string(output)), nil
}

// GetProjectConfig retrieves project configuration for a given project name.
// Returns hourly rate and export path if configured
func GetProjectConfig(projectName string) (*float64, string, error) {
	// check if global project
	registry, err := settings.LoadProjects()
	if err == nil && registry.Exists(projectName) {
		project, err := registry.GetProject(projectName)
		if err == nil {
			return project.HourlyRate, project.ExportPath, nil
		}
	}

	// fall back to .tmporc
	cfg, _, err := settings.FindAndLoad()
	if err == nil && cfg != nil && cfg.ProjectName == projectName {
		var hourlyRate *float64
		if cfg.HourlyRate > 0 {
			rate := cfg.HourlyRate
			hourlyRate = &rate
		}
		return hourlyRate, cfg.ExportPath, nil
	}

	// no configuration exists
	return nil, "", nil
}
