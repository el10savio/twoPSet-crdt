package twopset

import (
	"errors"
	"testing"

	"github.com/el10savio/gset-crdt/gset"
	"github.com/stretchr/testify/assert"
)

var (
	twopset TwoPSet
)

func init() {
	twopset = Initialize()
}

// TestList checks the basic functionality of TwoPSet List()
// List() should return all unique values appended to the TwoPSet
func TestList(t *testing.T) {
	twopset, _ = twopset.Addition("xx")

	expectedValue := []string{"xx"}
	actualValue := twopset.List()

	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestList_UpdatedValue checks the functionality of TwoPSet List() when
// multiple values are appended to TwoPSet it should return
// all the unique values appended to the TwoPSet
func TestList_UpdatedValue(t *testing.T) {
	twopset, _ = twopset.Addition("xx")
	twopset, _ = twopset.Addition("yy")
	twopset, _ = twopset.Addition("zz")

	expectedValue := []string{"xx", "yy", "zz"}
	actualValue := twopset.List()

	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestList_RemoveValue checks the functionality of TwoPSet List() when
// multiple values are appended to TwoPSet it should return
// all the unique values appended to the TwoPSet
func TestList_RemoveValue(t *testing.T) {
	twopset, _ = twopset.Addition("xx")
	twopset, _ = twopset.Removal("xx")
	twopset, _ = twopset.Removal("zz")

	expectedValue := []string{}
	actualValue := twopset.List()

	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestList_RemoveEmpty checks the functionality of TwoPSet List() when
// multiple values are appended to TwoPSet it should return
// all the unique values appended to the TwoPSet
func TestList_RemoveEmpty(t *testing.T) {
	twopset, _ = twopset.Removal("zz")

	expectedValue := []string{}
	actualValue := twopset.List()

	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestList_NoValue checks the functionality of TwoPSet List() when
// no values are appended to TwoPSet, it should return
// an empty string slice when the TwoPSet is empty
func TestList_NoValue(t *testing.T) {
	expectedValue := []string{}
	actualValue := twopset.List()

	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestAddition checks the basic functionality of TwoPSet Addition()
// it should return the TwoPSet back when the append is successful
func TestAddition(t *testing.T) {
	expectedValue := TwoPSet{Add: gset.GSet{[]string{"xx"}}, Remove: gset.GSet{[]string{}}}
	actualValue, actualError := twopset.Addition("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestAddition_NoValue checks the functionality of TwoPSet Addition()
// when a nil value is passed to it, it should return
// the an empty string slice back along with an error
func TestAddition_NoValue(t *testing.T) {
	expectedValue := TwoPSet{Add: gset.GSet{[]string{}}, Remove: gset.GSet{[]string{}}}
	expectedError := errors.New("empty value provided")
	actualValue, actualError := twopset.Addition("")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestAddition_Duplicate checks the functionality of TwoPSet Addition()
// when a duplicate value is passed to it, it should return
// only the unique TwoPSet values
func TestAddition_Duplicate(t *testing.T) {
	twopset, _ = twopset.Addition("xx")
	twopset, _ = twopset.Addition("yy")
	twopset, _ = twopset.Addition("xx")

	expectedValue := []string{"xx", "yy"}
	actualValue := twopset.List()

	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestAddition_Duplicate checks the functionality of TwoPSet Addition()
// when a duplicate value is passed to it, it should return
// only the unique TwoPSet values
func TestRemoval_Duplicate(t *testing.T) {
	twopset, _ = twopset.Removal("xx")
	twopset, _ = twopset.Removal("xx")

	expectedValue := []string{}
	actualValue := twopset.List()

	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestClear checks the basic functionality of TwoPSet Clear()
// utility function it clears all the values in a TwoPSet set
func TestClear(t *testing.T) {
	twopset, _ = twopset.Addition("xx1")
	twopset, _ = twopset.Addition("xx2")
	twopset = twopset.Clear()

	expectedValue := []string{}
	actualValue := twopset.List()

	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestClear_EmptyStore checks the functionality of TwoPSet Clear() utility function
// when no values are in it, it clears all the values in a TwoPSet set
func TestClear_EmptyStore(t *testing.T) {
	twopset = twopset.Clear()

	expectedValue := []string{}
	actualValue := twopset.List()

	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestLookup checks the basic functionality of TwoPSet Lookup() function
// it returns a boolean if a value passed is present in the TwoPSet set or not
func TestLookup(t *testing.T) {
	twopset, _ = twopset.Addition("xx")

	expectedValue := true
	actualValue, actualError := twopset.Lookup("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestLookup_NotPresent checks the functionality of TwoPSet Lookup() function
// it returns false if a value passed is not present in the TwoPSet
func TestLookup_NotPresent(t *testing.T) {
	twopset, _ = twopset.Addition("xx")

	expectedValue := false
	actualValue, actualError := twopset.Lookup("yy")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestLookup_NotPresent checks the functionality of TwoPSet Lookup() function
// it returns false if a value passed is not present in the TwoPSet
func TestLookup_Removed(t *testing.T) {
	twopset, _ = twopset.Addition("xx")
	twopset, _ = twopset.Removal("xx")

	expectedValue := false
	actualValue, actualError := twopset.Lookup("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestLookup_EmptySet checks the functionality of TwoPSet Lookup() function
// it returns false if the TwoPSet is empty irrespective of the value passed
func TestLookup_EmptySet(t *testing.T) {
	expectedValue := false
	actualValue, actualError := twopset.Lookup("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestLookup_EmptyLookup checks the functionality of TwoPSet Lookup() function
// it returns an error if the value passed is nil irrespective of the TwoPSet
func TestLookup_EmptyLookup(t *testing.T) {
	expectedValue := false
	expectedError := errors.New("empty value provided")

	actualValue, actualError := twopset.Lookup("")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	twopset = twopset.Clear()
}

// TestMerge checks the basic functionality of the Merge() function on multiple GSets
// it returns all the GSets merged together with unique elements as one single TwoPSet
func TestMerge(t *testing.T) {
	twopset1 := TwoPSet{Add: gset.GSet{[]string{"xx"}}, Remove: gset.GSet{[]string{}}}
	twopset2 := TwoPSet{Add: gset.GSet{[]string{"yy"}}, Remove: gset.GSet{[]string{}}}
	twopset3 := TwoPSet{Add: gset.GSet{[]string{"zz"}}, Remove: gset.GSet{[]string{"xx"}}}

	expectedValue := TwoPSet{Add: gset.GSet{[]string{"xx", "yy", "zz"}}, Remove: gset.GSet{[]string{"xx"}}}
	actualValue := Merge(twopset1, twopset2, twopset3)

	assert.Equal(t, expectedValue, actualValue)

	expectedList := []string{"yy", "zz"}
	actualList := actualValue.List()

	assert.ElementsMatch(t, expectedList, actualList)

	twopset = twopset.Clear()
}

// TestMerge_Empty checks the functionality of the Merge() function on multiple GSets
// when one TwoPSet is empty, it returns an empty TwoPSet followed by an error
func TestMerge_Empty(t *testing.T) {
	twopset1 := TwoPSet{Add: gset.GSet{[]string{"xx"}}, Remove: gset.GSet{[]string{}}}
	twopset2 := TwoPSet{Add: gset.GSet{[]string{}}, Remove: gset.GSet{[]string{}}}
	twopset3 := TwoPSet{Add: gset.GSet{[]string{"zz"}}, Remove: gset.GSet{[]string{"xx"}}}

	expectedValue := TwoPSet{Add: gset.GSet{[]string{"xx", "zz"}}, Remove: gset.GSet{[]string{"xx"}}}
	actualValue := Merge(twopset1, twopset2, twopset3)

	assert.Equal(t, expectedValue, actualValue)

	expectedList := []string{"zz"}
	actualList := actualValue.List()

	assert.ElementsMatch(t, expectedList, actualList)

	twopset = twopset.Clear()
}

// TestMerge_Duplicate checks the functionality of the Merge() function on multiple GSets
// when duplicate values are passed with the TwoPSet it returns all the GSets
// merged together with unique elements as one single TwoPSet
func TestMerge_Duplicate(t *testing.T) {
	twopset1 := TwoPSet{Add: gset.GSet{[]string{"xx"}}, Remove: gset.GSet{[]string{"zz"}}}
	twopset2 := TwoPSet{Add: gset.GSet{[]string{"xx"}}, Remove: gset.GSet{[]string{}}}
	twopset3 := TwoPSet{Add: gset.GSet{[]string{"zz"}}, Remove: gset.GSet{[]string{"zz"}}}

	expectedValue := TwoPSet{Add: gset.GSet{[]string{"xx", "zz"}}, Remove: gset.GSet{[]string{"zz"}}}
	actualValue := Merge(twopset1, twopset2, twopset3)

	assert.Equal(t, expectedValue, actualValue)

	expectedList := []string{"xx"}
	actualList := actualValue.List()

	assert.ElementsMatch(t, expectedList, actualList)

	twopset = twopset.Clear()
}
