package main

type Sorter interface {
	Sort([]int) []int
}

type Context struct {
	sorter Sorter
}

func (c *Context) SetSorter(sorter Sorter) {
	c.sorter = sorter
}

func (c *Context) ExecuteSort(arr []int) []int {
	return c.sorter.Sort(arr)
}
