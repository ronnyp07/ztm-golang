package main

import "fmt"

type BeltSize int
type Shipping int
type Payment int

const (
	Small BeltSize = iota
	Medium
	Large
)

const (
	Ground Shipping = iota
	Air
)

const (
	Cash Payment = iota
	ATM
)

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

func (p Payment) String() string {
	return []string{"Cash", "ATM"}[p]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type Paymenter interface {
	Payment() Payment
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
	Paymenter
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Span mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func (s *SpamMail) Payment() Payment {
	return Cash
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
	fmt.Printf("Payment %v type %v\n", item, item.Payment())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func (t *ToxicWaste) Payment() Payment {
	return ATM
}

func main() {
	mail := SpamMail{40000}
	automate(&mail)

	fmt.Println()
	fmt.Println()
	toxiMaterial := ToxicWaste{2000}
	fmt.Printf("Ship %v via %v\n", toxiMaterial, toxiMaterial.Ship())
	fmt.Printf("Payment %v type %v\n", toxiMaterial, toxiMaterial.Payment())
}
