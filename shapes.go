package main

import (
	"errors"
	"fmt"
	"math"
)

// The 'var' keyword allows us to define values global to the package
var ErrInsufficientFunds = errors.New("Cannot withdraw, insufficient funds")


type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin){
	//fmt.Printf("Address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// This function is implemented the Stringer interface, which lets you define how your type is printed when 
// used with the %s format string in prints.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
