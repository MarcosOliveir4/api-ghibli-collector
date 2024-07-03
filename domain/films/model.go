package films

import "github.com/google/uuid"

type Film struct {
	ID                     uuid.UUID `json:"id"`
	Title                  string    `json:"title"`
	OriginalTitleRomanised string    `json:"original_title_romanised"`
	Image                  string    `json:"image"`
	Description            string    `json:"description"`
	Director               string    `json:"director"`
	Producer               string    `json:"producer"`
	ReleaseDate            int64     `json:"release_date"`
	RunningTime            int64     `json:"running_time"`
	RtScore                int64     `json:"rt_score"`
}
