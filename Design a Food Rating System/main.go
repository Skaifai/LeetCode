package Design_a_Food_Rating_System

import "container/heap"

func main() {

}

type Food struct {
	Name   string
	Rating int
}

type FoodRatings struct {
	RatingsMap     map[string]int
	CuisinesMap    map[string]string
	CuisineHeapMap map[string]*MaxHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	_cuisinesMap := make(map[string]string, len(foods))
	_ratingsMap := make(map[string]int, len(foods))
	_cuisineHeapMap := make(map[string]*MaxHeap, len(foods))

	for i, food := range foods {
		_ratingsMap[food] = ratings[i]
		_cuisinesMap[food] = cuisines[i]

		if _, ok := _cuisineHeapMap[cuisines[i]]; !ok {
			_cuisineHeapMap[cuisines[i]] = &MaxHeap{}
		}

		heap.Push(_cuisineHeapMap[cuisines[i]], Food{foods[i], ratings[i]})
	}

	return FoodRatings{
		RatingsMap:     _ratingsMap,
		CuisinesMap:    _cuisinesMap,
		CuisineHeapMap: _cuisineHeapMap,
	}

}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.RatingsMap[food] = newRating
	cuisine := this.CuisinesMap[food]

	this.CuisineHeapMap[cuisine].ChangeRating(food, newRating)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	highestRatedFood := (*this.CuisineHeapMap[cuisine])[0]

	return highestRatedFood.Name
}

type MaxHeap []Food

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	if h[i].Rating == h[j].Rating {
		return h[i].Name < h[j].Name
	}
	return h[i].Rating > h[j].Rating
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Food))
}

func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h *MaxHeap) ChangeRating(s string, r int) {
	for i := 0; i < len(*h); i++ {
		if (*h)[i].Name == s {
			(*h)[i].Rating = r
			heap.Fix(h, i)
			break
		}
	}
}
