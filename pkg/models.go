package pkg

import (
	"fmt"
	"time"
)

const dateLayout = "2006-01-02"

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
	Added *Date  `json:"dateAdded"`
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
		template.DetectionDate = time.Time(*thread.DetectionDate).Format(dateLayout)
	}

	template.Variants = make([]HtmlVariants, len(thread.Variants))
	for i, v := range thread.Variants {
		variant := HtmlVariants{Name: v.Name}
		if v.Added != nil {
			variant.Added = time.Time(*v.Added).Format(dateLayout)
		}

		template.Variants[i] = variant
	}
}

func (d *Date) UnmarshalJSON(data []byte) error {
	// Remove quotes from the JSON string
	data = data[1 : len(data)-1]

	parsedTime, err := time.Parse(dateLayout, string(data))
	if err != nil {
		return fmt.Errorf("error parsing date: %w", err)
	}

	*d = Date(parsedTime)
	return nil
}
