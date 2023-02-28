package main

import "fmt"

type Ledger struct {
}

func (s *Ledger) makeEntry(accountID, txType string, amount int) {
	fmt.Printf("Make ledger entry for accountID %s with txType %s for amount %d\n", accountID, txType, amount)
}
