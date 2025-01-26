package pkg

import (
	"fmt"
	"time"
)

type Date time.Time //"2019-04-01"

type Thread struct {
	ThreatName    string     `json:"threatName"`
	Category      string     `json:"category"`
	Size          int        `json:"size"`
	DetectionDate *Date      `json:"detectionDate"`
	Variants      []Variants `json:"variants"`
}

type Variants struct {
	Name  string `json:"name"`
	Added Date   `json:"dateAdded"`
}

type HtmlTemplate struct {
	ThreatName    string
	Category      string
	Size          int
	DetectionDate string
	Variants      []HtmlVariants
}

type HtmlVariants struct {
	Name  string
	Added string
}

func (template *HtmlTemplate) FromAPI(thread Thread) {
	template.ThreatName = thread.ThreatName
	template.Category = thread.Category
	template.Size = thread.Size
	if thread.DetectionDate != nil {
		template.DetectionDate = time.Time(*thread.DetectionDate).Format("2006-01-02")
	}

	template.Variants = make([]HtmlVariants, len(thread.Variants))
	for i, v := range thread.Variants {
		template.Variants[i] = HtmlVariants{
			Name:  v.Name,
			Added: time.Time(v.Added).Format("2006-01-02"),
		}
	}
}

func (cd *Date) UnmarshalJSON(data []byte) error {
	// Remove quotes from the JSON string
	dateStr := string(data)
	dateStr = dateStr[1 : len(dateStr)-1]

	// Define the layout matching the date format
	layout := "2006-01-02"

	// Parse the date string
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return fmt.Errorf("error parsing date: %v", err)
	}

	parsedTime.Format("2006-01-02")

	// Assign the parsed time to the CustomDate
	*cd = Date(parsedTime)
	return nil
}
