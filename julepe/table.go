package julepe

// Table struct contains the pot value,
//   card leftovers from dealing as well
//   as the discard pile
type Table struct {
	Pot       float32
	Leftovers []Card
	Discards  []Card
}

// PotAddAll adds 0.50 from all players to the pot
func (t Table) PotAddAll(p Players) Table {
	for i := range p {
		t.Pot = t.Pot + 0.5
		p[i].AddToPot()
	}

	return t
}
