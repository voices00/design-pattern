package main

import (
	"fmt"
	"log"
)

//	状态模式是一种行为设计模式，
//	让你能在一个对象的内部状态变化时改变其行为， 使其看上去就像改变了自身所属的类一样。

//	让我们在一台自动售货机上使用状态设计模式。 为简单起见， 让我们假设自动售货机仅会销售一种类型的商品。
//	同时， 依然为了简单起见，我们假设自动售货机可处于 4 种不同的状态中：
//	有商品（hasItem）
//	无商品（noItem）
//	商品已请求 （itemRequested）
//	收到纸币 （hasMoney）
//	同时， 自动售货机也会有不同的操作。 再一次的， 为了简单起见， 我们假设其只会执行 4 种操作：
//	选择商品
//	添加商品
//	插入纸币
//	提供商品
//	当对象可以处于许多不同的状态中时应使用状态设计模式， 同时根据传入请求的不同， 对象需要变更其当前状态。
//
//	在我们的例子中， 自动售货机可以有多种不同的状态， 同时会在这些状态之间持续不断地互相转换。
//	我们假设自动售货机处于 商品已请求状态中。 在 “插入纸币” 的操作发生后， 机器将自动转换至 收到纸币状态。
//
//	根据其当前状态， 机器可就相同请求采取不同的行为。 例如，
//	如果用户想要购买一件商品， 机器将在 有商品状态时继续操作， 而在 无商品状态时拒绝操作。
//
//	自动售货机的代码不会被这一逻辑污染； 所有依赖于状态的代码都存在于各自的状态实现中。

func main() {
	var vendingMachine *VendingMachine
	vendingMachine = newVendingMachine(1, 10)

	var err error
	if err = vendingMachine.requestItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err = vendingMachine.insertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}

	if err = vendingMachine.dispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	if err = vendingMachine.addItem(2); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	if err = vendingMachine.requestItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err = vendingMachine.insertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}

	if err = vendingMachine.dispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}
}

type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}


type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}
func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}


type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}


type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}


type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
