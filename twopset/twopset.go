package twopset

import (
	"errors"

	"github.com/el10savio/gset-crdt/gset"
)

// package twopset implements the TwoPSet (2PSet) CRDT data type along with the functionality to
// append, list & lookup values in a TwoPSet. It also provides the functionality to
// merge multiple TwoPSets together and a utility function to clear a TwoPSet used in tests

// TwoPSet is the TwoPSet CRDT data type
type TwoPSet struct {
	// Add ...
	Add gset.GSet `json:"add"`
	// Remove ...
	Remove gset.GSet `json:"remove"`
}

// Initialize returns a new empty TwoPSet
func Initialize() TwoPSet {
	return TwoPSet{
		Add:    gset.Initialize(),
		Remove: gset.Initialize(),
	}
}

// Addition adds a new unique value to the TwoPSet using the
// union operation for each value on the existing TwoPSet
func (twopset TwoPSet) Addition(value string) (TwoPSet, error) {
	// Return an error if the value passed is nil
	if value == "" {
		return twopset, errors.New("empty value provided")
	}

	// Set = Set U value
	twopset.Add.Set, _ = twopset.Add.Append(value)

	// Return the new TwoPSet followed by nil error
	return twopset, nil
}

// Removal ...
func (twopset TwoPSet) Removal(value string) (TwoPSet, error) {
	// Return an error if the value passed is nil
	if value == "" {
		return twopset, errors.New("empty value provided")
	}

	// Set = Set U value
	twopset.Remove.Set, _ = twopset.Remove.Append(value)

	// Return the new TwoPSet followed by nil error
	return twopset, nil
}

// List returns all the elements present in the TwoPSet
func (twopset TwoPSet) List() []string {
	if len(twopset.Remove.Set) == 0 || len(twopset.Add.Set) == 0 {
		return twopset.Add.Set
	}

	resultGSet := twopset.Add

	for _, element := range twopset.Remove.Set {
		resultGSet = Delete(resultGSet, element)
	}

	return resultGSet.Set
}

// Delete ...
func Delete(gset gset.GSet, value string) gset.GSet {
	for index, element := range gset.Set {
		if element == value {
			gset.Set = append(gset.Set[:index], gset.Set[index+1:]...)
			return gset
		}
	}
	return gset
}

// Lookup returns either boolean true/false indicating
// if a given value is present in the TwoPSet or not
func (twopset TwoPSet) Lookup(value string) (bool, error) {
	// Return an error if the value passed is nil
	if value == "" {
		return false, errors.New("empty value provided")
	}

	list := twopset.List()

	// Iterative over the TwoPSet and check if the
	// value is the one we're searching
	// return true if the value exists
	for _, element := range list {
		if element == value {
			return true, nil
		}
	}

	// If the value isn't found after iterating
	// over the entire TwoPSet we return false
	return false, nil
}

// Merge conbines multiple TwoPSets together using Union
// and returns a single merged TwoPSet
func Merge(TwoPSets ...TwoPSet) TwoPSet {
	var twoPSetMerged TwoPSet

	// GSetMerged = GSetMerged U GSetToMergeWith
	for _, twopset := range TwoPSets {
		for _, value := range twopset.Add.Set {
			if value == "" {
				continue
			}
			twoPSetMerged, _ = twoPSetMerged.Addition(value)
		}
		for _, value := range twopset.Remove.Set {
			if value == "" {
				continue
			}
			twoPSetMerged, _ = twoPSetMerged.Removal(value)
		}
	}

	// Return the merged TwoPSet followed by nil error
	return twoPSetMerged
}

// Clear is utility function used only for tests
// to empty the contents of a given TwoPSet
func (twopset TwoPSet) Clear() TwoPSet {
	twopset.Add.Set = []string{}
	twopset.Remove.Set = []string{}
	return twopset
}
