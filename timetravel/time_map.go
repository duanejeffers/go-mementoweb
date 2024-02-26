package timetravel

import (
	"strings"
	"time"
)

type Memento struct {
	Datetime time.Time `json:"datetime"`
	Uri      string    `json:"uri"`
}

type Mementos struct {
	First Memento   `json:"first"`
	Last  Memento   `json:"last"`
	List  []Memento `json:"list"`
}

type MementoCompliant bool

func (mc *MementoCompliant) UnmarshalJSON(v []byte) error {
	var (
		res  bool
		jStr string
	)
	res = false
	jStr = strings.Trim(string(v), " ")

	if jStr == "yes" {
		res = true
	}

	*mc = MementoCompliant(res)
	return nil
}

type TimeMapUris struct {
	JsonFormat string `json:"json_format"`
	LinkFormat string `json:"link_format"`
}

type TimeMapRefRange struct {
	From             time.Time        `json:"from,omitempty"`
	Until            time.Time        `json:"until,omitempty"`
	Uri              string           `json:"uri"`
	MementoCompliant MementoCompliant `json:"memento_compliant,omitempty"`
	ArchiveId        string           `json:"archive_id,omitempty"`
}

type TimeMapPages struct {
	Prev TimeMapRefRange `json:"prev"`
	Next TimeMapRefRange `json:"next"`
}

type TimeMap struct {
	OriginalUri  string            `json:"original_uri"`
	TimegateUri  string            `json:"timegate_uri"`
	TimemapUri   TimeMapUris       `json:"timemap_uri"`
	Mementos     Mementos          `json:"mementos"`
	Pages        TimeMapPages      `json:"pages,omitempty"`
	TimeMapIndex []TimeMapRefRange `json:"timemap_index,omitempty"`
}
