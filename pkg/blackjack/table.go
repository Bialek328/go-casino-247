package blackjack

type Table struct {
    ID string
    Dealer Dealer
    Players []Player
}

func (t *Table) StartGame() {

}

func (t *Table) AddPlayer() error {

}

func (t *Table) RemovePlayer() error {

}

func (t *Table) ClearTable() {

}
