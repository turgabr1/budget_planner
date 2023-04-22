package repository

type Occurrence string

const (
	weekly   Occurrence = "WEEKLY"
	biWeekly            = "BIWEEKLY"
	monthly             = "MONTHLY"
	yearly              = "YEARLY"
)

func (s Occurrence) String() string {
	return string(s)
}

func OccurrenceList() []Occurrence {
	return []Occurrence{weekly, biWeekly, monthly, yearly}
}
