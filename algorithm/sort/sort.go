package sort

import (
	"github.com/liyue201/gostl/uitls/comparator"
	"github.com/liyue201/gostl/uitls/iterator"
) 

//sort the container by using quick sort
func Sort(begin, end iterator.SortableIterator, cmp comparator.Comparator) {
	quickSort(begin, end, cmp)
}

func quickSort(begin, end iterator.SortableIterator, cmp comparator.Comparator) {
	//todo
}
