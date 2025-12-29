package export

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/DylanDevelops/tmpo/internal/storage"
)

func ToCSV(entries []*storage.TimeEntry, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	header := []string{"Project", "Start Time", "End Time", "Duration (hours)", "Description", "Milestone"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	for _, entry := range entries {
		endTime := ""
		if entry.EndTime != nil {
			endTime = entry.EndTime.Format("2006-01-02 15:04:05")
		}

		milestoneName := ""
		if entry.MilestoneName != nil {
			milestoneName = *entry.MilestoneName
		}

		duration := entry.Duration().Hours()

		record := []string{
			entry.ProjectName,
			entry.StartTime.Format("2006-01-02 15:04:05"),
			endTime,
			fmt.Sprintf("%.2f", duration),
			entry.Description,
			milestoneName,
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}
	}

	return nil
}