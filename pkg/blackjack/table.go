package blackjack

type Table struct {
    ID string
    Dealer Dealer
    Players []Player
}

func (t *Table) StartGame() {

}

func (t *Table) AddPlayer() error {
    return nil
}

func (t *Table) RemovePlayer() error {
    return nil
}

func (t *Table) ClearTable() {

}
