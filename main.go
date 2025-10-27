package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var file string = "expenses.txt"

type Expense struct {
	Title  string  // название траты
	Amount float64 // сумма траты
}

func ReadAll() ([]Expense, error) {
	val, err := os.ReadFile(file)

	if err != nil {
		return []Expense{}, errors.New("ERROR")
	}

	s := strings.Split(string(val), "\n")

	// l := []string{}

	r := []Expense{}

	for i := range s {
		if strings.TrimSpace(s[i]) == "" {
			continue
		}
		s1 := strings.Split(s[i], ";")

		a, _ := strconv.ParseFloat(s1[1], 64)

		r = append(r, Expense{Title: s1[0], Amount: a})

	}
	return r, nil
}

func PrintAll() error {
	val, err := ReadAll()

	// fmt.Printf("%.2f", val)

	if err != nil {
		return err
	}

	for i, v := range val {
		fmt.Printf("%v)%s - %.2f\n", i, v.Title, v.Amount)
	}

	return err
}
func Add(title string, amount float64) error {
	val, err := os.ReadFile(file)
	new := []string{}
	if err != nil {
		return errors.New("ERROR")
	}
	if len(val) == 0 {
		os.WriteFile(file, []byte(string(title+";"+fmt.Sprintf("%.2f", amount))), 0644)
	}
	new = append(new, string(val)+"\n"+title+";"+fmt.Sprintf("%.2f", amount))

	new1 := strings.Join(new, "\n")

	os.WriteFile(file, []byte(string(new1)), 0644)
	return nil
}

func Delete(index int) error {
	items, err := ReadAll()

	if err != nil {
		return err
	}

	if index < 0 || index >= len(items) {
		return errors.New("ERROR")
	}

	items = append(items[:index], items[index+1:]...)
	new := []string{}

	for _, v := range items {
		new = append(new, v.Title+";"+fmt.Sprintf("%.2f", v.Amount))
		// fmt.Sprintf("\n%d)%s ; %.2f",i, v.Title, v.Amount)
	}
	new1 := strings.Join(new, "\n")
	os.WriteFile(file, []byte(new1), 0644)
	return nil
}

func Total() (float64, error) {
	val, err := ReadAll()

	if err != nil {
		return 0, errors.New("ERROR")
	}

	sum := 0.

	for _, v := range val {

		sum += v.Amount
	}
	return sum, nil
}

func main() {
	PrintAll()
	Add("Обед в Школе программирования Intocode", 350.0)
	Add("Проезд по Чеченской Республике", 80.0)
	PrintAll()
	// Delete(1)
	PrintAll()
	total, _:= Total()
	fmt.Printf("Сумма расходов: %.2f\n", total)
}
