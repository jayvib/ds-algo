package composition

import "fmt"

type Board struct {
	NailsNeeded int
	NailsDriven int
}

// NailDriver represents behavior to drive nails into a board.
type NailDriver interface {
	DriveNail(nailSupply *int,  board *Board)
}

// NailPuller represents behavior to remove nails into a board.
type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}

type NailDrivePuller interface {
	NailDriver
	NailPuller
}

type Mallet struct{}

func (Mallet) DriveNail(nailSupply *int, b *Board) {
	*nailSupply--
	b.NailsDriven++
	fmt.Println("Mallet: pounded nail into the board.")
}

type Crowbar struct{}

func (Crowbar) PullNail(nailSupply *int, b *Board) {
	b.NailsDriven--
	*nailSupply++
	fmt.Println("Crowbar: yanked nail out of the board.")
}

type Contractor struct{}

func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) {
	for b.NailsDriven < b.NailsNeeded {
		d.DriveNail(nailSupply, b)
	}
}

func (Contractor) Unfasten(p NailPuller, nailSupply *int, b *Board) {
	if b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}

func (c Contractor) ProcessBoards(dp NailDrivePuller, nailSupply *int, boards []Board) {
	for i := range boards {
		b := &boards[i]
		fmt.Printf("contractor: examining board #%d: %+v\n", i+1, b)
		switch {
		case b.NailsDriven < b.NailsNeeded:
			c.Fasten(dp, nailSupply, b)
		case b.NailsDriven > b.NailsNeeded:
			c.Unfasten(dp, nailSupply, b)
		}
	}
}

var _  NailDrivePuller = ToolBox{}

// embedding
type ToolBox struct {
	NailPuller
	NailDriver
	int
}



