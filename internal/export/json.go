package export

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/DylanDevelops/tmpo/internal/storage"
)

type ExportEntry struct {
	Project     string  `json:"project"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time,omitempty"`
	Duration    float64 `json:"duration_hours"`
	Description string  `json:"description,omitempty"`
	Milestone   string  `json:"milestone,omitempty"`
}

func ToJson(entries []*storage.TimeEntry, filename string) error {
	var exportEntries []ExportEntry

	for _, entry := range entries {
		export := ExportEntry{
			Project:     entry.ProjectName,
			StartTime:   entry.StartTime.Format("2006-01-02T15:04:05Z07:00"),
			Duration:    entry.Duration().Hours(),
			Description: entry.Description,
		}

		if entry.EndTime != nil {
			export.EndTime = entry.EndTime.Format("2006-01-02T15:04:05Z07:00")
		}

		if entry.MilestoneName != nil {
			export.Milestone = *entry.MilestoneName
		}

		exportEntries = append(exportEntries, export)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create JSON file: %w", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(exportEntries); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}