package blackjack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTable(t *testing.T) {
    table := NewTable()
    if table.ID == "" {
        t.Errorf("Expected id to be added on init")
    }
    if reflect.TypeOf(table) != reflect.TypeOf(&Table{}) {
        fmt.Println(reflect.TypeOf(table))
        t.Errorf("Wrong type iof table")
    }
    if reflect.TypeOf(table.Dealer) != reflect.TypeOf(&Dealer{}) {
        t.Errorf("Expected dealer to be added on init")
    }
}
