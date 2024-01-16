package main

import "math/rand"

func main() {

}

type RandomizedSet struct {
	Indexes map[int]int
	Values  []int
}

func Constructor() RandomizedSet {
	res := RandomizedSet{Indexes: make(map[int]int)}
	return res
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, wasThere := this.Indexes[val]; wasThere {
		return false
	}
	this.Values = append(this.Values, val)
	this.Indexes[val] = len(this.Values) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, wasThere := this.Indexes[val]; !wasThere {
		return false
	}
	valIndex := this.Indexes[val]
	lastElement := this.Values[len(this.Values)-1]
	this.Indexes[lastElement] = valIndex
	this.Values[valIndex] = lastElement
	this.Values = this.Values[:len(this.Values)-1]
	delete(this.Indexes, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	k := rand.Intn(len(this.Values))
	return this.Values[k]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
