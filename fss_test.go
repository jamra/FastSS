package fss

import "testing"

func TestBuildFFS(t *testing.T) {
	//Builds the data structure. This may need to be benchmarked and memory measured.
}

func TestHashWord(t *testing.T) {
	//Hash should work similar to Boost v1.36.0
}

func TestSplitWord(t *testing.T) {
	//Large words get split
}

func TestLookup(t *testing.T) {
	//Simply test if looking up a word returns all words within edit distance d
}

func TestSplitQuery(t *testing.T) {
	//Because splitting a query is not the same as splitting a full word, the result is multiple splits
}
