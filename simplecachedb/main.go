package main

import "fmt"

var GlobalStore map[string]int

type UserOperation interface {
	Get(key string) int
	Set(key string, val int)
	Delete(key string)
}

type db struct {
	store map[string]int
	next  *db
}
type Database struct {
	DB db
}

type TransactionStack struct {
	top  *db
	size int
}

func (ts *TransactionStack) begin() {
	temp := db{store: make(map[string]int)}
	temp.next = ts.top
	ts.top = &temp
	ts.size++
}

func (ts *TransactionStack) end() error {

	if ts.top == nil {
		return fmt.Errorf("ERROR: No Active Transactions")

	} else {
		node := &db{}
		ts.top = ts.top.next
		node.next = nil
		ts.size--
		return nil
	}
}

func (ts *TransactionStack) commit() {
	ActiveTransaction := ts.Peek()
	if ActiveTransaction != nil {
		for key, value := range ActiveTransaction.store {
			GlobalStore[key] = value
			if ActiveTransaction.next != nil {
				ActiveTransaction.next.store[key] = value
			}
		}
	} else {
		fmt.Println("INFO:Nothing to commit")
	}
}

func (ts *TransactionStack) rollback() error {
	if ts.top == nil {
		return fmt.Errorf("ERROR: No Active Transactions")
	} else {
		for key := range ts.top.store {
			delete(ts.top.store, key)
		}
		return nil
	}
}
func (ts *TransactionStack) Peek() *db {

	return ts.top
}

func (ts *TransactionStack) Get(key string) int {

	ActiveTransaction := ts.Peek()
	if ActiveTransaction == nil {
		if val, ok := GlobalStore[key]; ok {
			return val
		} else {
			return -1
		}
	} else {
		if val, ok := ActiveTransaction.store[key]; ok {
			return val
		} else {
			return -1
		}
	}
}

func (ts *TransactionStack) Set(key string, value int) {
	// Get key:value store from active transaction
	ActiveTransaction := ts.Peek()
	if ActiveTransaction == nil {
		GlobalStore[key] = value
	} else {
		ActiveTransaction.store[key] = value
	}
}

func (T *TransactionStack) Delete(key string) {
	ActiveTransaction := T.Peek()
	if ActiveTransaction == nil {
		delete(GlobalStore, key)
	} else {
		delete(ActiveTransaction.store, key)
	}
	fmt.Println(key, "deleted")
}
