package task

import (
	"fmt"
	"math"
	"modul/read"
	"sort"
)

// Task 1  function
func Task1(users []read.User) {
	sumCash := 0
	sumTotal := 0
	for _, value := range users {
		sumCash += value.Cash
	}

	for _, value := range users {
		sumTotal += value.Basket.Total
	}
	fmt.Println("1). Barcha customerlarni umumiy mablag'ini va qancha summa harid qilganini hisoblang va ko'rsating")
	fmt.Println("sumCash", sumCash)
	fmt.Println("sumTotal", sumTotal)
}

// Task 2 function
func Task2(users []read.User) {
	var (
		firstname string
		lastname  string
		max       int
		totals    int
	)
	for _, value := range users {
		max = math.MinInt
		if max < value.Basket.Total {
			firstname = value.FirstName
			lastname = value.LastName
			totals = value.Basket.Total
		}
	}
	fmt.Println("2) Eng ko'p pul sariflagan customerni aniqlang va ko'rsating")
	fmt.Println(firstname, lastname, totals)
}

// Task 3  function
func Task3(users []read.User) {
	var (
		max3        int
		productName string
	)
	max3 = math.MinInt
	for _, value := range users {
		for _, products := range value.Basket.Products {
			if products.Price > max3 {
				productName = products.Name
				max3 = products.Price
			}
		}
	}
	fmt.Println("3) To'plamidagi eng qimmat productni aniqlang va ko'rsating")
	fmt.Println(productName, max3)
}

// Task 4 function
func Task4(users []read.User) {
	task4sum := 0
	averageSum := 0
	count := 0
	for _, value := range users {
		for _, products := range value.Basket.Products {
			task4sum += products.Price
			count++
		}
	}
	averageSum = task4sum / count
	fmt.Println("4) Barcha productlar uchun o'rtacha narxni hisoblang va ko'rsating")
	fmt.Println("avragesum", averageSum)
	//fmt.Println(count)
	//fmt.Println(task4sum)
}

// Task 5 function
func Task5(users []read.User) {
	var (
		firstname5 string
		lastname5  string
		min5       = math.MaxInt
		totals     int
	)
	for _, value := range users {
		if min5 > value.Basket.Total {
			firstname5 = value.FirstName
			lastname5 = value.LastName
			totals = value.Basket.Total
			min5 = value.Basket.Total
		}
	}
	fmt.Println("5)  Eng arzon savdo qilgan customerni aniqlang va ko'rsating")
	fmt.Println(firstname5, lastname5, totals)
}

// Task 6 function
func Task6(users []read.User) {
	var max int = -1
	var name string
	for _, value := range users {
		for _, products := range value.Basket.Products {

			if products.Quantity > max {
				max = products.Quantity
				name = products.Name

			}
		}
	}
	fmt.Println("6)  Eng ko'p sotilgan productlar categoriyasini aniqlang va chiqaring")
	fmt.Println("best selling product ", name)
}

// Task 7 function
func Task7(users []read.User) {
	var (
		firstnamemax7  string
		firstnamemin7  string
		lastnamemax7   string
		lastnamemin7   string
		productmaxname string
		productminname string
		productmax7    int
		productmin7    int
		maxSell        int
		minSell        = math.MaxInt
	)

	for _, value := range users {
		for _, value2 := range value.Basket.Products {
			if maxSell < value2.Quantity {
				firstnamemax7 = value.FirstName
				lastnamemax7 = value.LastName
				productmax7 = value2.Quantity
				productmaxname = value2.Name
				maxSell = value2.Quantity
			}
			if minSell > value2.Quantity {
				firstnamemin7 = value.FirstName
				lastnamemin7 = value.LastName
				productmin7 = value2.Quantity
				productminname = value2.Name
				minSell = value2.Quantity
			}
		}
	}
	fmt.Println("7)  Eng kop va eng kam sotilgan proudctlar nomini chiqarish")
	fmt.Println("Max =", productmaxname, productmax7, firstnamemax7, lastnamemax7)
	fmt.Println("Min =", productminname, productmin7, firstnamemin7, lastnamemin7)
}

// Task 8 function
func Task8(users []read.User) {
	var (
		productCountSum int
		count           int
		averageCountSum float32
	)

	for _, value := range users {
		for _, product := range value.Basket.Products {
			productCountSum += product.Quantity
			count++
		}
	}

	if count > 0 {
		averageCountSum = float32(productCountSum) / float32(count)
	}
	fmt.Println("8)   Har bir savdo uchun o'rtacha mahsulot miqdorini hisoblang va ko'rsating")
	fmt.Println(averageCountSum)
	//fmt.Println(count)
	//fmt.Println(productCountSum)
}

// Task 9  function
func Task9(users []read.User) {
	var (
		firstname9 string
		lastname9  string
		maxCount   int = math.MinInt
	)

	for _, value := range users {
		for _, product := range value.Basket.Products {
			if maxCount < product.Quantity {
				firstname9 = value.FirstName
				lastname9 = value.LastName
				maxCount = product.Quantity
			}
		}
	}
	fmt.Println("9) Eng ko'p mahsulot sotib olishgan mijozni aniqlang va ko'rsating")
	fmt.Println(firstname9, lastname9, maxCount)
}

// Task 10 function
func Task10(users []read.User) {
	var (
		firstname10 string
		lastname10  string
		productname string
		max         int = math.MinInt
	)
	for _, value := range users {
		for _, value2 := range value.Basket.Products {
			if max < value2.Quantity {
				max = value2.Quantity
				firstname10 = value.FirstName
				lastname10 = value.LastName
				productname = value2.Name
			}
		}
	}
	fmt.Println("10)  Eng ko'p savdolarda ko'rinadigan mahsulotni aniqlang va ko'rsating")
	fmt.Println(productname, "(", firstname10, lastname10, "-", max, ")")
}

// Task 11 function
func Task11(users []read.User) {
	var (
		count        int = 0
		Sum11        int = 0
		averageSum11 int = 0
		max          int = math.MinInt
		firstname11  string
		lastname11   string
	)
	for _, value := range users {
		Sum11 += value.Basket.Total
		count++
		if max < value.Cash {
			max = value.Cash
			firstname11 = value.FirstName
			lastname11 = value.LastName
		}
	}
	averageSum11 = Sum11 / count
	fmt.Println("11)   O'rtacha savdo mablag'i bo'yicha eng ko'p mablag'aga ega bo'lgan customerni aniqlang va ko'rsating")
	//	fmt.Println("count", count)
	//	fmt.Println("all sum", Sum11)
	fmt.Println("avarage sum", averageSum11)
	fmt.Println(firstname11, lastname11, max)

}

// Task 12 function
func Task12(users []read.User) {
	var (
		maxPrices int = math.MinInt
		catagory  string
		name12    string
	)

	for _, value := range users {
		for _, product := range value.Basket.Products {
			if maxPrices < product.Quantity*product.Price {
				name12 = product.Name
				catagory = product.Category
				maxPrices = product.Quantity * product.Price
			}
		}
	}
	fmt.Println("12)  Eng ko'p umumiy daromad (quantity* price) qilgan toifani aniqlang va koʻrsating")
	fmt.Println(catagory, "(", name12, ")", maxPrices)
}

// Task 13 function

func Task13(users []read.User) {
	fmt.Println("13)   Eng qimmat productni o'z ichiga olgan savdoni aniqlang va ko'rsating")
	for _, user := range users {
		max := math.MinInt
		var found bool
		for _, product := range user.Basket.Products {
			if price := product.Price * product.Quantity; price > max {
				max = price
				found = true
			}
		}
		if found {

			fmt.Println(user.FirstName, user.LastName, max)
		}
	}
}

// Task 14 function
func Task14(users []read.User) {
	fmt.Println("14)   Har bir mijoz uchun eng ko'p xarid qilgan toifani aniqlang va ko'rsating")
	type UserWithMaxPrice struct {
		User     read.User
		MaxPrice int
	}
	var usersWithMaxPrice []UserWithMaxPrice

	for _, user := range users {
		max := math.MinInt
		var found bool
		for _, product := range user.Basket.Products {
			if price := product.Price * product.Quantity; price > max {
				max = price
				found = true
			}
		}
		if found {
			usersWithMaxPrice = append(usersWithMaxPrice, UserWithMaxPrice{User: user, MaxPrice: max})
		}
	}

	sort.Slice(usersWithMaxPrice, func(i, j int) bool {
		return usersWithMaxPrice[i].MaxPrice > usersWithMaxPrice[j].MaxPrice
	})
	for _, userWithMaxPrice := range usersWithMaxPrice {

		fmt.Println(userWithMaxPrice.User.FirstName, userWithMaxPrice.User.LastName, userWithMaxPrice.MaxPrice)
	}
}

// Task 15 function
func Task15(users []read.User) {
	fmt.Println("15)   Har bitta productdan savdo paytida umumiy nechta sotilganini aniqlang va ko'rsating")
	var (
		sum int = 0
	)
	for _, value := range users {
		for _, value2 := range value.Basket.Products {
			sum += value2.Quantity

		}
	}

	fmt.Println("total:", sum)
}
