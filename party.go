package party

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/player"
)

// Party contains the owner of the party and the players in it.
// It also contains the max of players able to enter the party.
type Party struct {
	owner      *player.Player
	players    []*player.Player
	maxPlayers int
}

// New returns a new *Party.
func New(owner *player.Player, maxPlayers int) *Party {
	return &Party{
		owner:      owner,
		maxPlayers: maxPlayers,
	}
}

// Owner retutrns the owner of the party.
func (p *Party) Owner() *player.Player {
	return p.owner
}

// Players returns the players in the party.
func (p *Party) Players() []*player.Player {
	return p.players
}

// MaxPlayers returns the max of players able to join the party.
func (p *Party) MaxPlayers() int {
	return p.maxPlayers
}

// AddPlayer adds the player provided to the party if not full.
func (p *Party) AddPlayer(pl *player.Player) error {
	if len(p.players) >= p.maxPlayers {
		p.players = append(p.players, pl)
	} else {
		return fmt.Errorf("the party is full")
	}
	return nil
}

// RemovePlayer removes the player provided from the party.
func (p *Party) RemovePlayer(pl *player.Player) {
	for n, i := range p.players {
		if i == pl {
			p.players = append(p.players[:n], p.players[n+1:]...)
		}
	}
}

// HasPlayer returns true if the player provided is in the party, and false if not.
func (p *Party) HasPlayer(pl *player.Player) bool {
	for _, i := range p.players {
		if i == pl {
			return true
		}
	}
	return false
}

// SetOwner sets the player provided as the new owner.
func (p *Party) SetOwner(pl *player.Player) error {
	if p.HasPlayer(pl) {
		p.AddPlayer(p.owner)
		p.owner = pl
		p.RemovePlayer(pl)
	} else {
		return fmt.Errorf("player %s is not in that party", pl.Name())
	}
	return nil
}

// SendMessage sends a message to the players and the owner in the party.
func (p *Party) SendMessage(msg ...interface{}) {
	for _, i := range p.players {
		i.Message(msg)
	}
}
