package generic

const Extend = "extend"

type Auth struct {
	SessionId string `json:"auth,omitempty"`
}

type Filter map[string]interface{}

type Params struct {
	// Object properties to be returned.
	//
	//Default: extend.
	Output any `json:"output,omitempty"`
	// Return only those results that exactly match the given filter.
	//
	//Accepts an object, where the keys are property names, and the values are either a single value or an array of values to match against.
	//
	//Does not support properties of text data type.
	Filter Filter `json:"filter,omitempty"`
	// Limit the number of records returned.
	Limit int `json:"limit,omitempty"`
	// Return the number of records in the result instead of the actual data.
	CountOutput bool `json:"countOutput,omitempty"`
	// If set to true, return only objects that the user has write permissions to.
	//
	//Default: false.
	Editable bool `json:"editable,omitempty"`
	// Return results that do not match the criteria given in the search parameter.
	ExcludeSearch bool `json:"excludeSearch,omitempty"`
	// Use IDs as keys in the resulting array.
	PreserveKeys bool `json:"preservekeys,omitempty"`
	// Return results that match the given pattern (case-insensitive).
	//
	//Accepts an object, where the keys are property names, and the values are strings to search for. If no additional options are given, this will perform a LIKE "%…%" search.
	//
	//Supports only properties of string and text data type.
	Search Filter `json:"search"`
	// If set to true, return results that match any of the criteria given in the filter or search parameter instead of all of them.
	//
	//Default: false.
	SearchByAny bool `json:"searchByAny,omitempty"`
	// If set to true, enables the use of "*" as a wildcard character in the search parameter.
	//
	//Default: false.
	SearchWildcardsEnabled bool `json:"searchWildcardsEnabled,omitempty"`
	// Sort the result by the given properties. Refer to a specific API get method description for a list of properties that can be used for sorting. Macros are not expanded before sorting.
	//
	//If no value is specified, data will be returned unsorted.
	SortField []string `json:"sortfield,omitempty"`
	// Order of sorting. If an array is passed, each value will be matched to the corresponding property given in the sortfield parameter.
	//
	//Possible values are:
	//ASC - (default) ascending;
	//DESC - descending.
	SortOrder []string `json:"sortorder,omitempty"`
	// The search parameter will compare the beginning of fields, that is, perform a LIKE "…%" search instead.
	//
	//Ignored if searchWildcardsEnabled is set to true.
	StartSearch bool `json:"startSearch"`
}
