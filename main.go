package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

var wg sync.WaitGroup
var receivedOrdersCh = make(chan order)
var validOrderCh = make(chan order)
var invalidOrderCh = make(chan invalidOrder)

func main() {

	go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)

	wg.Add(1)
	go func(validOrderCh <-chan order) {
		order := <-validOrderCh
		fmt.Printf("Valid order received: %v\n", order)
		wg.Done()
	}(validOrderCh)

	go func(invalidOrderCh <-chan invalidOrder) {
		order := <-invalidOrderCh
		fmt.Printf("Invalid order received: %v. Issue: %v\n", order.order, order.err)
		wg.Done()
	}(invalidOrderCh)

	wg.Wait()
	//fmt.Println(orders)
}

func validateOrders(in <-chan order, out chan<- order, errCh chan<- invalidOrder) {
	order := <-in
	if order.Quantity <= 0 {
		// error condition
		errCh <- invalidOrder{order: order, err: errors.New("quantity must be greater than zero")}
	} else {
		//success path
		out <- order
	}
}

func receiveOrders(out chan<- order) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		out <- newOrder
		//orders = append(orders, newOrder)
	}
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": -8, "status": 1}`,
}
