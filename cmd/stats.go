package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/DylanDevelops/tmpo/internal/storage"
	"github.com/spf13/cobra"
)

var (
	statsToday bool
	statsWeek bool
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show time tracking statistics",
	Long:  `Display statistics and summaries of your time tracking data.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := storage.Initialize()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			
			os.Exit(1)
		}

		defer db.Close()

		var start, end time.Time
		var periodName string

		if statsToday {
			start = time.Now().Truncate(24 * time.Hour)
			end = start.Add(24 * time.Hour)
			periodName = "Today"
		} else if statsWeek {
			now := time.Now()
			weekday := int(now.Weekday())
			if weekday == 0 {
				weekday = 7
			}

			start = now.AddDate(0, 0, -weekday+1).Truncate(24 * time.Hour)
			end = start.AddDate(0, 0, 7)
			periodName = "This Week"
		} else {
			entries, err := db.GetEntries(0)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)

				os.Exit(1)
			}

			ShowAllTimeStats(entries, db)

			return
		}

		entries, err := db.GetEntriesByDateRange(start, end)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)

			os.Exit(1)
		}

		ShowPeriodStats(entries, periodName)
	},
}

// ShowPeriodStats prints aggregated statistics for a named period to standard
// output. Given a slice of *storage.TimeEntry and a human-readable periodName,
// the function:
//
//  - returns early with a message if entries is empty,
//  - computes and prints the total accumulated time and its hour equivalent,
//  - prints the total number of entries,
//  - aggregates time by project and prints a per-project line with duration and
//    percentage of the total,
//  - attempts to load configuration and, if a positive hourly rate is present,
//    prints an estimated earnings line.
//
// Aggregation is done via a map[string]time.Duration; iteration order is
// therefore non-deterministic. Percentages are computed as projectSeconds /
// totalSeconds * 100, so if the total duration is zero the percentage values
// may be undefined (NaN/Inf). All output is produced using fmt.
func ShowPeriodStats(entries []*storage.TimeEntry, periodName string) {
	if len(entries) == 0 {
		fmt.Printf("No entries for %s.\n", periodName)

		return
	}

	projectStats := make(map[string]time.Duration)
	projectEarnings := make(map[string]float64)
	var totalDuration time.Duration
	var totalEarnings float64
	hasAnyEarnings := false

	for _, entry := range entries {
		duration := entry.Duration()
		projectStats[entry.ProjectName] += duration
		totalDuration += duration

		if entry.HourlyRate != nil {
			earnings := duration.Hours() * *entry.HourlyRate
			projectEarnings[entry.ProjectName] += earnings
			totalEarnings += earnings
			hasAnyEarnings = true
		}
	}

	fmt.Printf("\n[tmpo] Stats for %s\n\n", periodName)
	fmt.Printf("    Total Time: %s (%.2f hours)\n", formatDuration(totalDuration), totalDuration.Hours())
	fmt.Printf("    Total Entries: %d\n", len(entries))

	if hasAnyEarnings {
		fmt.Printf("    Total Estimated Earnings: $%.2f\n", totalEarnings)
	}

	fmt.Println()
	fmt.Println("    By Project:")
	for project, duration := range projectStats {
		percentage := (duration.Seconds() / totalDuration.Seconds()) * 100
		fmt.Printf("        %-20s  %s  (%.1f%%)\n", project, formatDuration(duration), percentage)

		if earnings, ok := projectEarnings[project]; ok && earnings > 0 {
			fmt.Printf("        └─ Estimated Earnings: $%.2f\n", earnings)
		}
	}
}

// ShowAllTimeStats prints aggregated all-time statistics to standard output.
// Given a slice of *storage.TimeEntry and a pointer to the database, the
// function:
//
//  - returns early with a message if entries is empty,
//  - computes and prints the total accumulated time and its hour equivalent,
//  - prints the total number of entries and number of tracked projects,
//  - aggregates time by project and prints a per-project line with duration and
//    percentage of the total.
//
// The function fetches the list of projects from the provided database to
// determine the number of projects tracked. Aggregation is done via a
// map[string]time.Duration; iteration order is therefore non-deterministic.
// If the total duration is zero, percentage values may be undefined. All
// output is produced using fmt.
func ShowAllTimeStats(entries []*storage.TimeEntry, db *storage.Database) {
	if len(entries) == 0 {
		fmt.Println("No entries found.")

		return
	}

	projectStats := make(map[string]time.Duration)
	projectEarnings := make(map[string]float64)
	var totalDuration time.Duration
	var totalEarnings float64
	hasAnyEarnings := false

	for _, entry := range entries {
		duration := entry.Duration()
		projectStats[entry.ProjectName] += duration
		totalDuration += duration

		if entry.HourlyRate != nil {
			earnings := duration.Hours() * *entry.HourlyRate
			projectEarnings[entry.ProjectName] += earnings
			totalEarnings += earnings
			hasAnyEarnings = true
		}
	}

	projects, _ := db.GetAllProjects()

	fmt.Printf("\n[tmpo] All-Time Statistics\n")
	fmt.Printf("    Total Time: %s (%.2f hours)\n", formatDuration(totalDuration), totalDuration.Hours())
	fmt.Printf("    Total Entries: %d\n", len(entries))
	fmt.Printf("    Projects Tracked: %d\n", len(projects))

	if hasAnyEarnings {
		fmt.Printf("    Total Estimated Earnings: $%.2f\n", totalEarnings)
	}

	fmt.Println()
	fmt.Println("    By Project:")
	for project, duration := range projectStats {
		percentage := (duration.Seconds() / totalDuration.Seconds()) * 100
		fmt.Printf("        %-20s  %s  (%.1f%%)\n", project, formatDuration(duration), percentage)

		if earnings, ok := projectEarnings[project]; ok && earnings > 0 {
			fmt.Printf("        └─ Estimated Earnings: $%.2f\n", earnings)
		}
	}
}

func init() {
	rootCmd.AddCommand(statsCmd)

	statsCmd.Flags().BoolVarP(&statsToday, "today", "t", false, "Show today's stats")
	statsCmd.Flags().BoolVarP(&statsWeek, "week", "w", false, "Show this week's stats")
}
