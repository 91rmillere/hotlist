package hotlist

// Entry represents a is a hotlist entry
type Entry struct {
	PlateNumber      string
	MatchingStrategy MatchingStrategy
	Description      string
}

// Result represents a hotlist result.
type Result struct {
	Entry
	IsAlert bool
	IsBest  bool
}

// MatchingStrategy represents a strategy for matching a hotlist entry
type MatchingStrategy int

// MatchingStrategy constants
const (
	MatchingStrategyUndefined MatchingStrategy = iota
	MatchingStrategyLenient
	MatchingStrategyExact
)

// ListType Represents the type of hotlist
type ListType int

// ListType Constants
const (
	ListTypeBlacklist ListType = iota + 1
	ListTypeWhitelist
)

// Whitelist represents a whitelist hotlist. A whitlist used for finding plates that are not
// in the list. A whitelist can only have one entry per plate number
type Whitelist struct {
	Items map[string]*Entry
}

// NewWhitelist returns a pointer to a whitelist
func NewWhitelist() *Whitelist {
	return &Whitelist{
		Items: make(map[string]*Entry),
	}
}

// Search returns a a whitlist hit when passed plate number catidates. The first plate is considered the best match.
// Search will return the hotlist entry, if the result is an alert, and if the match is the best match. If it is an alert
// the matching strategty returned will be MatchingStrategyUndefined and IsBest will be false and the platenumber will be
// the first cantidate
func (w *Whitelist) Search(cantidates ...string) Result {

	for i, plate := range cantidates {
		entry, ok := w.Items[plate]
		if ok {
			return Result{
				Entry:   *entry,
				IsAlert: false,
				// IsBest if matches first cantidate
				IsBest: (i == 0),
			}
		}

	}

	return Result{
		Entry: Entry{
			PlateNumber:      cantidates[0],
			MatchingStrategy: MatchingStrategyUndefined,
			Description:      "",
		},
		IsAlert: true,
		IsBest:  false,
	}

}

// Blacklist represents a blacklist hotlist. A blacklist used for finding plates that are
// in the list. A blacklist can only have one entry per plate number
type Blacklist struct {
	Items map[string]*Entry
}

// NewBlacklist returns a pointer to a new blacklist
func NewBlacklist() *Blacklist {
	return &Blacklist{
		Items: make(map[string]*Entry),
	}
}

// Search returns a whitlist hit when passed plate number catidates. The first plate is considered the best match.
// Search will return the hotlist entry, if the result is an alert, and if the match is the best match. If it is not an alert
// the matching strategty returned will be MatchingStrategyUndefined, IsBest will be false and the description will be an empty string
func (b *Blacklist) Search(cantidates ...string) Result {

	for i, plate := range cantidates {
		entry, ok := b.Items[plate]
		if ok {
			return Result{
				Entry:   *entry,
				IsAlert: true,
				// IsBest if matches first cantidate
				IsBest: (i == 0),
			}
		}
	}

	return Result{
		Entry: Entry{
			PlateNumber: cantidates[0],
		},
		IsAlert: false,
		IsBest:  false,
	}
}
