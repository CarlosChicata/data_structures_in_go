/*
Purpose:
	implement my version of disjoint-union set.

Version: 1.0

Methods:
	- add [x] : add element in universe if not exists
	- belong [x] : element belong this universe
	- len [x] : count all elements of universe set
	- display [ ] : display all element in universe set
	- union [ ] : merge two set: they are two variants:
		- unionBySize [x] : merge using size
		- unionByRank [ ] : merge using rank
	- parentIn [x] : get parent of set
	- sizeSet [ ] : get size of set by specified element 
	- sizeIn [x] : get number of set in universe.

Internal Methods:
	- preparing [x]: preparing map fields in Disjoint Union
*/

package main

import "fmt"


type DisjointUnion struct {
	universe map[string]string
	size_sets map[string]int
}

/// preparing maps of disjoint-union set
func (D *DisjointUnion) preparing () {
	D.universe = make(map[string]string)
	D.size_sets = make(map[string]int)
}

/// add value of disjoint-union set if it is not exists
func (D *DisjointUnion) add (value string) bool {
	var rptaFlag bool

	if _, ok := D.universe[value]; ok {
		rptaFlag = true
	}else{
		D.universe[value] = value
		D.size_sets[value] = 1
		rptaFlag = true
	}

	return rptaFlag
}

/// length of disjoint-union set
func (D *DisjointUnion) len () int {
	return len(D.universe)
}

/// if value belong this disjoint-union set
func (D *DisjointUnion) belong (value string) bool {
	rptaFlag := false
	if _, ok := D.universe[value]; ok {
		rptaFlag = true
	}
	return rptaFlag
}

/// find parent of value
func (D *DisjointUnion) parentIn (value string) (string, bool) {
	if D.belong(value) == false {
		return value, false
	}
	
	valueMap := D.universe[value]
	
	for value != valueMap {
		value = D.universe[value]
		valueMap = D.universe[valueMap]
	}

	return valueMap, true
}

/// size of subset with element size
func (D *DisjointUnion) sizeSet(value string) int {
	if D.belong(value) {
		parentValue, _ := D.parentIn(value)
		return D.size_sets[parentValue]
	} else { 
		return 0
	}
}

/// union set using size by criterio
func (D *DisjointUnion) unionBySize(value1 string, value2 string) int {
	maxValue, isValid1 := D.parentIn(value1)
	minValue, isValid2 := D.parentIn(value2)

	if isValid1 == false || isValid2 == false {
		return -1
	}

	if D.size_sets[maxValue] < D.size_sets[maxValue] {
		maxValue, minValue = minValue, maxValue
	}

	D.universe[minValue] = D.universe[maxValue]
	D.size_sets[maxValue] = D.size_sets[maxValue] + D.size_sets[minValue]
	return D.size_sets[maxValue]
}


func main(){
	testingDUS := DisjointUnion{}
	testingDUS.preparing()
	testingDUS.add("hola")
	testingDUS.add("que")
	testingDUS.add("tal")
	testingDUS.add("?")
	testingDUS.add("?")
	testingDUS.add("?")
	fmt.Println(testingDUS.unionBySize("hola", "que"))
	fmt.Println(testingDUS.parentIn("hola"))
	fmt.Println(testingDUS.sizeSet("hola"))
	fmt.Println(testingDUS.parentIn("que"))
	fmt.Println(testingDUS.sizeSet("que"))
	fmt.Println(testingDUS.parentIn("tal"))
}
