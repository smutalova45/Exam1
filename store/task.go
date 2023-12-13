package store

import (
	"fmt"
	"math"
)

type Store struct {
	File []Client `json:"client"`
}
type Client struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Cash      int    `json:"cash"`
	Basket    Basket `json:"basket"`
}
type Basket struct {
	Idbasket string     `json:"id"`
	Product  []Products `json:"products"`
	Total    int        `json:"total"`
}
type Products struct {
	Idproduct string `json:"id"`
	Category  string `json:"category"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
}

func (s *Store) Task1() {
	fmt.Println("-------------------")
	client := s.File
	sumcash := 0
	sumtotal := 0
	for _, v := range client {
		sumcash += v.Cash
		sumtotal += v.Basket.Total

	}
	fmt.Println("Total cash:", sumcash)
	fmt.Println("Total sum:", sumtotal)

}
func (s *Store) Task2() {
	fmt.Println("-------------------")

	client := s.File
	max := 0
	clientmap := map[string]string{}
	clientnamemap := map[string]string{}
	totalmap := map[string]int{}

	for _, v := range client {
		clientnamemap[v.Id] = v.FirstName
		clientmap[v.Id] = v.LastName
		totalmap[v.Id] = v.Basket.Total
	}
	for _, v := range client {
		if totalmap[v.Id] > max {
			max = totalmap[v.Id]
		}
	}
	for v := range totalmap {
		if max == totalmap[v] {
			fmt.Println("eng kop sarf etgan mijoz: ", clientnamemap[v], clientmap[v], max, " sum")
		}
	}

}
func (s *Store) Task3() {
	fmt.Println("-------------------")

	client := s.File
	name := map[string]string{}
	price := map[string]int{}
	max1 := 0

	for _, v := range client {
		for _, v2 := range v.Basket.Product {

			name[v2.Idproduct] = v2.Name
			price[v2.Idproduct] = v2.Price

		}
	}
	// fmt.Println(name)
	// fmt.Println(price)
	for _, v := range client {
		for _, v2 := range v.Basket.Product {
			if price[v2.Idproduct] > max1 {
				max1 = price[v2.Idproduct]
			}
		}
	}
	for v := range price {
		if max1 == price[v] {
			fmt.Println("name:", name[v], "price:", max1)
		}
	}

}
func (s *Store) Task4() {
	fmt.Println("-------------------")
	client := s.File

	sum := 0
	count := 0
	for _, v := range client {
		for _, v2 := range v.Basket.Product {

			sum += v2.Price
			count++

		}
	}

	fmt.Println("avarage sum: ", sum/count)

}
func (s *Store) Task5() {
	fmt.Println("-------------------")

	client := s.File

	clientmap := map[string]string{}
	clientnamemap := map[string]string{}
	minmap := map[string]int{}
	min := math.MaxInt32

	for _, v := range client {
		clientnamemap[v.Id] = v.FirstName
		clientmap[v.Id] = v.LastName
		minmap[v.Id] = v.Basket.Total
	}
	for _, v := range client {
		if minmap[v.Id] < min {
			min = minmap[v.Id]
		}
	}
	for v := range minmap {
		if min == minmap[v] {
			fmt.Println("eng kam sarf etgan mijoz: ", clientnamemap[v], clientmap[v], min, " sum")
		}
	}

}
func (s *Store) Task6() {

	fmt.Println("-------------------")
	client := s.File
	categoryname := map[string]string{}
	quantity := map[string]int{}
	max3 := 0
	for _, v := range client {
		for _, v2 := range v.Basket.Product {
			categoryname[v2.Idproduct] = v2.Category
			quantity[v2.Idproduct] = v2.Quantity
		}
	}

	for _, v := range client {
		for _, v2 := range v.Basket.Product {
			if quantity[v2.Idproduct] > max3 {
				max3 = quantity[v2.Idproduct]
			}
		}
	}
	for i := range quantity {
		if max3 == quantity[i] {
			fmt.Println("category: ", categoryname[i])

		}

	}
}

func (s *Store) Task7() {
	fmt.Println("-------------------")
	client := s.File
	firstname := map[string]string{}
	lastname := map[string]string{}
	nameproduct := map[string]string{}
	quantity := map[string]int{}
	min := math.MaxInt32
	max := 0

	// fmt.Println(firstname)
	// fmt.Println(lastname)

	for _, v := range client {
		for _, v2 := range v.Basket.Product {
			firstname[v2.Idproduct] = v.FirstName
			lastname[v2.Idproduct] = v.LastName
			nameproduct[v2.Idproduct] = v2.Name  //product nomi
			quantity[v2.Idproduct] = v2.Quantity /// product soni

		}
	}

	for _, v := range client {
		for _, v2 := range v.Basket.Product {
			if quantity[v2.Idproduct] > max {
				max = quantity[v2.Idproduct]
			}
			if quantity[v2.Idproduct] < min {
				min = quantity[v2.Idproduct]
			}
		}
	}
	for v := range quantity {
		if max == quantity[v] {
			fmt.Println("client: ", firstname[v], lastname[v], "product:", nameproduct[v], max)
		}
	}
	fmt.Println("***********Min Products*********")

	for v := range quantity {
		count := 0
		if min == quantity[v] {
			count++
			if count == 1 {
				fmt.Println("client:", firstname[v], lastname[v], "product: ", nameproduct[v], min)
				break
			}

		}
	}

}

func (s *Store) Task8() {

	fmt.Println("-------------------")
	client := s.File

	sum := 0
	count := 0
	for _, v := range client {
		for _, v2 := range v.Basket.Product {

			sum += v2.Quantity
			count++

		}
	}

	fmt.Println("avarage quantity : ", sum/count)
}
func (s *Store) Task9() {
	fmt.Println("-------------------")

	client := s.File
	firstName := make(map[string]string)
	lastName := make(map[string]string)
	productCount := make(map[string]int)
	for _, v := range client {
		basket := v.Basket.Product
		firstName[v.Id] = v.FirstName
		lastName[v.Id] = v.LastName
		var total int
		for _, b := range basket {
			total += b.Quantity
		}
		productCount[v.Id] = total
	}
	max := 0
	for value := range productCount {
		if productCount[value] > max {
			max = productCount[value]
		}
	}
	for value := range productCount {
		if productCount[value] == max {
			fmt.Printf("client:%s %s %d buys products ", firstName[value], lastName[value], max)
		}
	}
}
func (s *Store) Task10() {
	fmt.Println("-------------------")

	client := s.File
	products := map[string]int{}
	firstname := map[string]string{}
	productname := map[string]string{}
	for _, v := range client {
		firstname[v.Id] = v.FirstName
		for _, product := range v.Basket.Product {
			products[product.Name] += product.Quantity
			productname[product.Name] = product.Name
		}

	}
	max := 0
	for _, quantity := range products {

		if quantity > max {
			max = quantity

		}
	}
	for key := range products {
		if max == products[key] {

			fmt.Println("most sold product across clients: ", productname[key], max)
		}

	}

}
func (s *Store) Task11() {
	fmt.Println("-------------------")
	client := s.File
	totalmoney := 0
	for _, v := range client {
		totalmoney += v.Cash

	}
	avaragemoney := totalmoney / len(client)

	fmt.Println(" overall avarage money in total:", avaragemoney)
	for _, v := range client {
		if v.Cash > avaragemoney {
			fmt.Printf("client: %s ,%s\n", v.FirstName, v.LastName)
			break
		}

	}

}
func (s *Store) Task12() {
	fmt.Println("-------------------")
	client := s.File

	nameofrevenue := map[string]string{}
	categoryRevenue := make(map[string]int)
	for _, customer := range client {
		for _, product := range customer.Basket.Product {
			categoryRevenue[product.Category] += product.Price * product.Quantity
			nameofrevenue[product.Category] = product.Category
		}
	}
	hieghrevenue := 0
	for _, revenue := range categoryRevenue {

		if revenue > hieghrevenue {
			hieghrevenue = revenue
		}
	}
	for key := range nameofrevenue {

		fmt.Println("highest sold category: ", key, "price is", hieghrevenue)
		break
	}
}

func (s *Store) Task13() { // basket which has the most expensive product
	fmt.Println("-------------------")
	client := s.File
	max := 0
	var maxbasketid string
	for _, v := range client {
		for _, v2 := range v.Basket.Product {
			if v2.Price > max {
				max = v2.Price
				maxbasketid = v.Basket.Idbasket
			}
		}
	}
	fmt.Println("the most expensive price:", max, "in this basket id :", maxbasketid)

}
func (s *Store) Task14() { // har bir mijoz uchun en kop harid qilingan toifani aniqlang
	fmt.Println("-------------------")
	client := s.File
	count := map[string]int{}
	for _, v := range client {
		for _, v2 := range v.Basket.Product {
			count[v2.Category] += v2.Quantity
		}
		category := ""
		qunatity := 0
		for i, v := range count {
			if v > qunatity {
				qunatity = v
				category = i
			}
		}
		fmt.Println("client:", v.FirstName, v.LastName, "category name: ", category, "quantity:", qunatity)
	}

}
func (s *Store) Task15() {
	// har bita product umumiy savdoda qancha sotlgani
	fmt.Println("-------------------")
	client := s.File
	soldcount:=map[string]int{}
	for _,v:=range client{
		for _,v2:=range v.Basket.Product{
			soldcount[v2.Name]+=v2.Quantity
		}
	}

	for name,quantity:=range soldcount{
		fmt.Println("product :",name,"sold:",quantity ,"times sold ")
	}

}
