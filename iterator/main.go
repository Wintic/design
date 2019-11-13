package main

import "fmt"

// 迭代器模式
func main() {
	bookShelf := BookShelf{}
	bookShelf.AppendBook(&Book{"11111"})
	bookShelf.AppendBook(&Book{"22222"})
	bookShelf.AppendBook(&Book{"33333"})
	bookShelf.AppendBook(&Book{"44444"})
	bookShelf.AppendBook(&Book{"55555"})

	// 获取迭代器
	bookShelfIterator := bookShelf.GetIterator()

	// 遍历
	for bookShelfIterator.HasNext() {
		fmt.Println(bookShelfIterator.Next().(*Book).Name)
	}
}

// 迭代器
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Aggregate interface {
	GetIterator() Iterator
}

// 书
type Book struct {
	Name string
}

// 书架
type BookShelf struct {
	book []*Book
	last int
}

// 增加书
func (bs *BookShelf) AppendBook(book *Book) {
	bs.book = append(bs.book, book)
	bs.last++
}

// 获取迭代器
func (bs *BookShelf) GetIterator() Iterator {
	return &BookShelfIterator{
		bookShelf: bs,
		index:     0,
	}
}

// 书架迭代器
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.index < bsi.bookShelf.last {
		return true
	}
	return false
}

func (bsi *BookShelfIterator) Next() interface{} {
	// 防止数组越界
	if bsi.index < bsi.bookShelf.last {
		b := bsi.bookShelf.book[bsi.index]
		bsi.index++
		return b
	}
	return nil
}
