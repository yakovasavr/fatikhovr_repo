/* Даны даты заезда и отъезда каждого гостя. Для каждого гостя дата заезда строго раньше даты отъезда (то есть каждый гость останавливается хотя бы на одну ночь). В пределах одного дня считается, что сначала старые гости выезжают, а затем въезжают новые. Найти максимальное число постояльцев, которые одновременно проживали в гостинице (считаем, что измерение количества постояльцев происходит в конце дня).
sample = [ (1, 2), (1, 3), (2, 4), (2, 3), ]
*/

package main

import "fmt"

type input struct {
	datain	[]int
	dataout	[]int
}

type day struct {
	date	int
	status	int
	next	*day
}

type list struct {
	head	*day
	// length	int
	// tail	*day
}
// 0 is in, 1 is out.

func (hotel *list) findMaxDay() int {
	temp := hotel.head
	counter := 0
	maxday := 0
	maxdate := 0
	for (temp != nil) {
		if (temp.status == 0) {
			counter++
		} else {
			counter--
		}
		if (temp.next == nil || temp.date != temp.next.date) {
			if maxday < counter {
				maxday = counter
				maxdate = temp.date
			}
		}
		temp = temp.next
	}
	return (maxdate)
}

func (hotel *list) printList() {
	temp := hotel.head
	for (temp != nil) {
		fmt.Printf("%d %d\n", temp.date, temp.status)
		temp = temp.next
	}
}

func (hotel *list) addElem(date int, status int) {
	temp := new(day)
	temp.date = date
	temp.status = status
	temp.next = nil
	if (hotel.head == nil) {
		hotel.head = temp
	} else if (hotel.head.date > date) {
		temp.next = hotel.head
		hotel.head = temp
	} else {
		current := hotel.head
		for (current.next != nil) && (current.next.date < date) {
			current = current.next
		}
		if (current.next == nil) {
			current.next = temp
		} else {
			temp.next = current.next
			current.next = temp
		}
	}
}

func (hotel *list) structimport(jan input) {
	for _, x := range jan.datain {
		hotel.addElem(x, 0)
	}
	for _, x := range jan.dataout {
		hotel.addElem(x, 1)
	}
}

func main() {
	hotel := new(list)
	var maxdate int
	var jan input = input{datain: []int{1, 3, 5, 1, 1}, dataout: []int{2, 9, 10, 2, 3}}

	hotel.structimport(jan)
	hotel.printList()
	maxdate = hotel.findMaxDay()
	fmt.Println(maxdate)
}