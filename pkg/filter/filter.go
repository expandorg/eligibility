package filter

type Filter struct {
	ID    uint64 `json:"id"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Filters []Filter

type GroupedFilters map[string]Filters

func (fs *Filters) GroupByType() GroupedFilters {
	grouped := GroupedFilters{}
	for _, f := range *fs {
		grouped[f.Type] = append(grouped[f.Type], f)
	}
	return grouped
}
