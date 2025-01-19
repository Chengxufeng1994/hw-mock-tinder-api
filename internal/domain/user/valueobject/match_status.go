package valueobject

type MatchStatus string

const (
	MatchStatusPending  MatchStatus = "pending"
	MatchStatusAccepted MatchStatus = "accepted"
	MatchStatusRejected MatchStatus = "rejected"
)

func NewMatchStatus(status string) MatchStatus {
	switch status {
	case MatchStatusPending.String():
		return MatchStatusPending
	case MatchStatusAccepted.String():
		return MatchStatusAccepted
	case MatchStatusRejected.String():
		return MatchStatusRejected
	default:
		return ""
	}
}

func (m MatchStatus) String() string {
	return string(m)
}
