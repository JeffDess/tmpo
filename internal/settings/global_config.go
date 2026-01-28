package settings

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/DylanDevelops/tmpo/internal/currency"
	"go.yaml.in/yaml/v3"
)

type GlobalConfig struct {
	Currency   string `yaml:"currency"`
	DateFormat string `yaml:"date_format,omitempty"`
	TimeFormat string `yaml:"time_format,omitempty"`
	Timezone   string `yaml:"timezone,omitempty"`
	ExportPath string `yaml:"export_path,omitempty"`
}

func DefaultGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		Currency:   currency.DefaultCurrency,
		DateFormat: "",
		TimeFormat: "",
		Timezone:   "",
		ExportPath: "",
	}
}

func GetGlobalConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	tmpoDir := filepath.Join(home, ".tmpo")
	if devMode := os.Getenv("TMPO_DEV"); devMode == "1" || devMode == "true" {
		tmpoDir = filepath.Join(home, ".tmpo-dev")
	}

	return filepath.Join(tmpoDir, "config.yaml"), nil
}

func LoadGlobalConfig() (*GlobalConfig, error) {
	configPath, err := GetGlobalConfigPath()
	if err != nil {
		return DefaultGlobalConfig(), nil
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return DefaultGlobalConfig(), nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read global config: %w", err)
	}

	var config GlobalConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse global config at %s: %w (check file syntax)", configPath, err)
	}

	if config.Currency == "" {
		config.Currency = currency.DefaultCurrency
	}

	return &config, nil
}

func (gc *GlobalConfig) Save() error {
	configPath, err := GetGlobalConfigPath()
	if err != nil {
		return err
	}

	tmpoDir := filepath.Dir(configPath)
	if err := os.MkdirAll(tmpoDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := yaml.Marshal(gc)
	if err != nil {
		return fmt.Errorf("failed to marshal global config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write global config: %w", err)
	}

	return nil
}

// GetDisplayTimezone returns the user's configured timezone or local timezone as fallback
func GetDisplayTimezone() *time.Location {
	cfg, err := LoadGlobalConfig()
	if err != nil || cfg.Timezone == "" {
		return time.Local
	}

	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		return time.Local
	}

	return loc
}

// toDisplayTime converts a UTC time to the user's display timezone
func toDisplayTime(t time.Time) time.Time {
	return t.In(GetDisplayTimezone())
}

func FormatTime(t time.Time) string {
	t = toDisplayTime(t)

	cfg, err := LoadGlobalConfig()
	if err != nil || cfg.TimeFormat == "" || cfg.TimeFormat == "Keep current" {
		return t.Format("3:04 PM")
	}

	if cfg.TimeFormat == "24-hour" {
		return t.Format("15:04")
	}

	return t.Format("3:04 PM")
}

func FormatTimePadded(t time.Time) string {
	t = toDisplayTime(t)

	cfg, err := LoadGlobalConfig()
	if err != nil || cfg.TimeFormat == "" || cfg.TimeFormat == "Keep current" {
		return t.Format("03:04 PM")
	}

	if cfg.TimeFormat == "24-hour" {
		return t.Format("15:04")
	}

	return t.Format("03:04 PM")
}

func FormatDate(t time.Time) string {
	t = toDisplayTime(t)

	cfg, err := LoadGlobalConfig()
	if err != nil || cfg.DateFormat == "" || cfg.DateFormat == "Keep current" {
		return t.Format("01/02/2006")
	}

	switch cfg.DateFormat {
	case "MM/DD/YYYY":
		return t.Format("01/02/2006")
	case "DD/MM/YYYY":
		return t.Format("02/01/2006")
	case "YYYY-MM-DD":
		return t.Format("2006-01-02")
	default:
		return t.Format("01/02/2006")
	}
}

func FormatDateDashed(t time.Time) string {
	t = toDisplayTime(t)

	cfg, err := LoadGlobalConfig()
	if err != nil || cfg.DateFormat == "" || cfg.DateFormat == "Keep current" {
		return t.Format("01-02-2006")
	}

	switch cfg.DateFormat {
	case "MM/DD/YYYY":
		return t.Format("01-02-2006")
	case "DD/MM/YYYY":
		return t.Format("02-01-2006")
	case "YYYY-MM-DD":
		return t.Format("2006-01-02")
	default:
		return t.Format("01-02-2006")
	}
}

func FormatDateTime(t time.Time) string {
	return FormatDate(t) + " " + FormatTime(t)
}

func FormatDateTimeDashed(t time.Time) string {
	return FormatDateDashed(t) + " " + FormatTime(t)
}

func FormatDateLong(t time.Time) string {
	t = toDisplayTime(t)

	return t.Format("Mon, Jan 2, 2006")
}

func FormatDateTimeLong(t time.Time) string {
	t = toDisplayTime(t)

	cfg, err := LoadGlobalConfig()
	if err != nil || cfg.TimeFormat == "" || cfg.TimeFormat == "Keep current" {
		return t.Format("Jan 2, 2006 at 3:04 PM")
	}

	if cfg.TimeFormat == "24-hour" {
		return t.Format("Jan 2, 2006 at 15:04")
	}

	return t.Format("Jan 2, 2006 at 3:04 PM")
}
