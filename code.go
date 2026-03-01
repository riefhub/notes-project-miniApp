package main

import "fmt"

type notes struct {
	title, content string
	dates          tanggal
}
type tanggal struct {
	day, month, year string
}

type tabNote [999]notes

var array1, array2 tabNote
var n int
var exit bool

func main() {
	fmt.Println(`
	███████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████
	█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█░░░░░░████████████░░░░░░██████████░░░░░░█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█
	█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀░░████████████░░▄▀░░░░░░░░░░██░░▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█
	█░░▄▀░░░░░░░░░░█░░▄▀░░░░░░▄▀░░█░░▄▀░░░░░░▄▀░░█░░▄▀░░████████████░░▄▀▄▀▄▀▄▀▄▀░░██░░▄▀░░█░░▄▀░░░░░░▄▀░░█░░░░░░▄▀░░░░░░█░░▄▀░░░░░░░░░░█░░▄▀░░░░░░░░░░█
	█░░▄▀░░█████████░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░█░░▄▀░░████████████░░▄▀░░░░░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░█████░░▄▀░░█████░░▄▀░░█████████░░▄▀░░█████████
	█░░▄▀░░█████████░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░█░░▄▀░░████████████░░▄▀░░██░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░█████░░▄▀░░█████░░▄▀░░░░░░░░░░█░░▄▀░░░░░░░░░░█
	█░░▄▀░░█████████░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░█░░▄▀░░████████████░░▄▀░░██░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░█████░░▄▀░░█████░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█
	█░░▄▀░░█████████░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░█░░▄▀░░████████████░░▄▀░░██░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░█████░░▄▀░░█████░░▄▀░░░░░░░░░░█░░░░░░░░░░▄▀░░█
	█░░▄▀░░█████████░░▄▀░░██░░▄▀░░█░░▄▀░░██░░▄▀░░█░░▄▀░░████████████░░▄▀░░██░░▄▀░░░░░░▄▀░░█░░▄▀░░██░░▄▀░░█████░░▄▀░░█████░░▄▀░░█████████████████░░▄▀░░█
	█░░▄▀░░░░░░░░░░█░░▄▀░░░░░░▄▀░░█░░▄▀░░░░░░▄▀░░█░░▄▀░░░░░░░░░░████░░▄▀░░██░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀░░░░░░▄▀░░█████░░▄▀░░█████░░▄▀░░░░░░░░░░█░░░░░░░░░░▄▀░░█
	█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░████░░▄▀░░██░░░░░░░░░░▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█████░░▄▀░░█████░░▄▀▄▀▄▀▄▀▄▀░░█░░▄▀▄▀▄▀▄▀▄▀░░█
	█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█░░░░░░░░░░░░░░████░░░░░░██████████░░░░░░█░░░░░░░░░░░░░░█████░░░░░░█████░░░░░░░░░░░░░░█░░░░░░░░░░░░░░█
	███████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████`)
	start()
}

func start() {
	var x byte

	fmt.Println("						-----------------------------------------------------------")
	fmt.Println("								>> PRESS ENTER TO START <<")
	fmt.Println("						-----------------------------------------------------------")
	fmt.Scanf("%c", &x)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	if x == 13 {
		menu()
	} else {
		fmt.Println(`
		────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
		─██████████████─████████████████───████████████████───██████████████─████████████████──────██████──██████─██████████████─██████──██████─
		─██░░░░░░░░░░██─██░░░░░░░░░░░░██───██░░░░░░░░░░░░██───██░░░░░░░░░░██─██░░░░░░░░░░░░██──────██░░██──██░░██─██░░░░░░░░░░██─██░░██──██░░██─
		─██░░██████████─██░░████████░░██───██░░████████░░██───██░░██████░░██─██░░████████░░██──────██░░██──██░░██─██░░██████░░██─██░░██──██░░██─
		─██░░██─────────██░░██────██░░██───██░░██────██░░██───██░░██──██░░██─██░░██────██░░██──────██░░██──██░░██─██░░██──██░░██─██░░██──██░░██─
		─██░░██████████─██░░████████░░██───██░░████████░░██───██░░██──██░░██─██░░████████░░██──────██░░██████░░██─██░░██──██░░██─██░░██████░░██─
		─██░░░░░░░░░░██─██░░░░░░░░░░░░██───██░░░░░░░░░░░░██───██░░██──██░░██─██░░░░░░░░░░░░██──────██░░░░░░░░░░██─██░░██──██░░██─██░░░░░░░░░░██─
		─██░░██████████─██░░██████░░████───██░░██████░░████───██░░██──██░░██─██░░██████░░████──────██████████░░██─██░░██──██░░██─██████████░░██─
		─██░░██─────────██░░██──██░░██─────██░░██──██░░██─────██░░██──██░░██─██░░██──██░░██────────────────██░░██─██░░██──██░░██─────────██░░██─
		─██░░██████████─██░░██──██░░██████─██░░██──██░░██████─██░░██████░░██─██░░██──██░░██████────────────██░░██─██░░██████░░██─────────██░░██─
		─██░░░░░░░░░░██─██░░██──██░░░░░░██─██░░██──██░░░░░░██─██░░░░░░░░░░██─██░░██──██░░░░░░██────────────██░░██─██░░░░░░░░░░██─────────██░░██─
		─██████████████─██████──██████████─██████──██████████─██████████████─██████──██████████────────────██████─██████████████─────────██████─
		────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────`)
	}
}

func menu() {
	exit = false
	for exit != true {
		var inputmenu int
		fmt.Println("------------------------------------------------------")
		fmt.Println("			MENU")
		fmt.Println("------------------------------------------------------")
		fmt.Println("1. NEW NOTE")
		fmt.Println("2. SAVED NOTES")
		fmt.Println("3. EXIT")
		fmt.Println()
		fmt.Scan(&inputmenu)
		fmt.Println()
		if inputmenu == 1 {
			newnote(&array1, &n)
		} else if inputmenu == 2 {
			savednote(&array1, &n)
		} else if inputmenu == 3 {
			exit = true
		}
	}
}

func newnote(A *tabNote, n *int) {
	var t, c byte
	fmt.Print("Type your note's title here: ")
	fmt.Scanf("%c", &t)
	for t != 13 {
		A[*n].title = A[*n].title + string(t)
		fmt.Scanf("%c", &t)
	}
	fmt.Print("Type your note here: ")
	fmt.Scanf("%c", &c)
	for c != 13 {
		A[*n].content = A[*n].content + string(c)
		fmt.Scanf("%c", &c)
	}
	fmt.Print("Type the desired dates for your saved note (DD MM YYYY): ")
	fmt.Scan(&A[*n].dates.day, &A[*n].dates.month, &A[*n].dates.year)
	*n = *n + 1
	fmt.Println("Your note has been saved :>")
	menu()
}

func savednote(A *tabNote, n *int) {
	var inputview, inputsort1, inputsort2 int
	var inputsearch1, inputsearch2 string
	var x int
	var status, status2 string
	fmt.Println("0. BACK TO MENU	1001. SORT BY	1002. SEARCH BY")
	fmt.Println()
	for i := 0; i <= *n-1; i++ {
		fmt.Print(i+1, ".", A[i].title, " ")
		fmt.Print("(")
		fmt.Printf("%v %v %v", A[i].dates.day, A[i].dates.month, A[i].dates.year)
		fmt.Print(")")
		fmt.Println()
	}
	fmt.Println()
	fmt.Scan(&x)
	fmt.Println()
	if x == 0 {
		menu()
	} else if x == 1001 {
		fmt.Println("1. TITLE")
		fmt.Println("2. DAY")
		fmt.Println("3. MONTH")
		fmt.Println("4. YEAR")
		fmt.Println()
		fmt.Scan(&inputsort1)
		fmt.Println()
		fmt.Println("1. ASCENDING	2.DESCENDING")
		fmt.Scan(&inputsort2)
		if inputsort1 == 1 {
			status = "title"
			if inputsort2 == 1 {
				ascSort(A, status, n)
				savednote(A, n)

			} else if inputsort2 == 2 {
				descSort(A, status, n)
				savednote(A, n)
			}

		} else if inputsort1 == 2 {
			status = "day"
			if inputsort2 == 1 {
				ascSort(A, status, n)
				savednote(A, n)

			} else if inputsort2 == 2 {
				descSort(A, status, n)
				savednote(A, n)
			}

		} else if inputsort1 == 3 {
			status = "month"
			if inputsort2 == 1 {
				ascSort(A, status, n)
				savednote(A, n)

			} else if inputsort2 == 2 {
				descSort(A, status, n)
				savednote(A, n)
			}

		} else if inputsort1 == 4 {
			status = "year"
			if inputsort2 == 1 {
				ascSort(A, status, n)
				savednote(A, n)

			} else if inputsort2 == 2 {
				descSort(A, status, n)
				savednote(A, n)
			}
		}
	} else if x == 1002 {
		fmt.Println("1. TITLE")
		fmt.Println("2. DAY")
		fmt.Println("3. MONTH")
		fmt.Println("4. YEAR")
		fmt.Println()
		fmt.Scan(&inputsearch1)
		fmt.Println()
		fmt.Print("Type your specific description to search: ")
		fmt.Scan(&status2)
		fmt.Println()
		if inputsearch1 == "1" {
			inputsearch2 = "title"
			search(A, &array2, inputsearch2, status2, n)
		} else if inputsearch1 == "2" {
			inputsearch2 = "day"
			search(A, &array2, inputsearch2, status2, n)
		} else if inputsearch1 == "3" {
			inputsearch2 = "month"
			search(A, &array2, inputsearch2, status2, n)
		} else if inputsearch1 == "4" {
			inputsearch2 = "year"
			search(A, &array2, inputsearch2, status2, n)
		}
	}
	if x > 1000 {
		x = 0
	}
	fmt.Println(A[x-1].content)
	fmt.Println()
	fmt.Println("1. EDIT TITLE	2. EDIT NOTE	3. DELETE NOTE	4. BACK TO MENU")
	fmt.Println()
	fmt.Scan(&inputview)
	fmt.Println()
	if inputview == 1 {
		editnote(A, &x)
	} else if inputview == 2 {
		edittitle(A, &x)
	} else if inputview == 3 {
		deletenote(A, &x, n)
	} else if inputview == 4 {
		menu()
	}
}

func edittitle(A *tabNote, n *int) {
	var c byte
	fmt.Print("Type your new note here: ")
	fmt.Scanf("%c", &c)
	for c != 13 {
		A[*n-1].title = A[*n-1].title + string(c)
		fmt.Scanf("%c", &c)
	}
}

func editnote(A *tabNote, n *int) {
	var c byte
	fmt.Print("Type your new note here: ")
	fmt.Scanf("%c", &c)
	for c != 13 {
		A[*n-1].content = A[*n-1].content + string(c)
		fmt.Scanf("%c", &c)
	}
}

func deletenote(A *tabNote, x *int, n *int) {
	for i := *x - 1; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--
}

func ascSort(A *tabNote, status string, n *int) {
	var pass, idx, i int
	var temp notes
	if status == "title" {
		for pass = 1; pass <= *n-1; pass++ {
			idx = pass - 1
			for i = pass; i <= *n-1; i++ {
				if A[idx].title > A[i].title {
					idx = i
				}
			}
			temp = A[idx]
			A[idx] = A[pass-1]
			A[pass-1] = temp
		}
	} else if status == "day" {
		for pass = 1; pass <= *n-1; pass++ {
			idx = pass - 1
			for i = pass; i <= *n-1; i++ {
				if A[idx].dates.day > A[i].dates.day {
					idx = i
				}
			}
			temp = A[idx]
			A[idx] = A[pass-1]
			A[pass-1] = temp
		}
	} else if status == "month" {
		for pass = 1; pass <= *n-1; pass++ {
			idx = pass - 1
			for i = pass; i <= *n-1; i++ {
				if A[idx].dates.month > A[i].dates.month {
					idx = i
				}
			}
			temp = A[idx]
			A[idx] = A[pass-1]
			A[pass-1] = temp
		}
	} else if status == "year" {
		for pass = 1; pass <= *n-1; pass++ {
			idx = pass - 1
			for i = pass; i <= *n-1; i++ {
				if A[idx].dates.year > A[i].dates.year {
					idx = i
				}
			}
			temp = A[idx]
			A[idx] = A[pass-1]
			A[pass-1] = temp
		}
	}
}

func descSort(A *tabNote, status string, n *int) {
	var pass, i int
	var temp notes
	if status == "title" {
		for pass = 1; pass <= *n-1; pass++ {
			temp = A[pass]
			i = pass
			for i > 0 && A[i-1].title < temp.title {
				A[i] = A[i-1]
				i = i - 1
			}
			A[i] = temp
		}
	} else if status == "day" {
		for pass = 1; pass <= *n-1; pass++ {
			temp = A[pass]
			i = pass
			for i > 0 && A[i-1].title < temp.dates.day {
				A[i] = A[i-1]
				i = i - 1
			}
			A[i] = temp
		}
	} else if status == "month" {
		for pass = 1; pass <= *n-1; pass++ {
			temp = A[pass]
			i = pass
			for i > 0 && A[i-1].title < temp.dates.month {
				A[i] = A[i-1]
				i = i - 1
			}
			A[i] = temp
		}
	} else if status == "year" {
		for pass = 1; pass <= *n-1; pass++ {
			temp = A[pass]
			i = pass
			for i > 0 && A[i-1].title < temp.dates.month {
				A[i] = A[i-1]
				i = i - 1
			}
			A[i] = temp
		}
	}
}

func search(A *tabNote, T *tabNote, status string, searchstat string, n *int) {
	var i, x, inputview int
	x = 0
	if status == "title" {
		for i = 0; i <= *n-1; i++ {
			if A[i].title == searchstat {
				fmt.Print(x+1, ".")
				fmt.Println(A[i].title, A[i].dates.day, A[i].dates.month, A[i].dates.year)
				T[x] = A[i]
				x += 1
			}
		}
		fmt.Println()
		fmt.Scan(&x)
		fmt.Println()
		fmt.Println(T[x-1].content)
		fmt.Println()
		fmt.Println("1.	BACK TO MENU")
		fmt.Println()
		fmt.Scan(&inputview)
		fmt.Println()
		if inputview == 1 {
			menu()
		}
	} else if status == "day" {
		for i = 0; i <= *n-1; i++ {
			if A[i].dates.day == searchstat {
				fmt.Print(x+1, ".")
				fmt.Println(A[i].title, A[i].dates.day, A[i].dates.month, A[i].dates.year)
				T[x] = A[i]
				x += 1
			}
		}
		fmt.Println()
		fmt.Scan(&x)
		fmt.Println()
		fmt.Println(T[x-1].content)
		fmt.Println()
		fmt.Println("1. BACK TO MENU")
		fmt.Println()
		fmt.Scan(&inputview)
		fmt.Println()
		if inputview == 1 {
			menu()
		}
	} else if status == "month" {
		for i = 0; i <= *n-1; i++ {
			if A[i].dates.month == searchstat {
				fmt.Print(x+1, ".")
				fmt.Println(A[i].title, A[i].dates.day, A[i].dates.month, A[i].dates.year)
				T[x] = A[i]
				x += 1
			}
		}
		fmt.Println()
		fmt.Scan(&x)
		fmt.Println()
		fmt.Println(T[x-1].content)
		fmt.Println()
		fmt.Println("1. BACK TO MENU")
		fmt.Println()
		fmt.Scan(&inputview)
		fmt.Println()
		if inputview == 1 {
			menu()
		}
	} else if status == "year" {
		for i = 0; i <= *n-1; i++ {
			if A[i].dates.year == searchstat {
				fmt.Print(x+1, ".")
				fmt.Println(A[i].title, A[i].dates.day, A[i].dates.month, A[i].dates.year)
				T[x] = A[i]
				x += 1
			}
		}
		fmt.Println()
		fmt.Scan(&x)
		fmt.Println()
		fmt.Println(T[x-1].content)
		fmt.Println()
		fmt.Println("1. BACK TO MENU")
		fmt.Println()
		fmt.Scan(&inputview)
		fmt.Println()
		if inputview == 1 {
			menu()
		}
	}
}
