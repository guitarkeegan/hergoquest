package control

// for player inputs
import "io"

type Controller struct {
	name string
	io   *io.Reader
}

// TODO: if they pass an empty string?
func (p *Controller) New(name string) *Controller {
	return &Controller{
		name: name,
	}
}

func (p *Controller) Move(from, to [2]int) {}
func (p *Controller) Attack()              {}
func (p *Controller) Defend()              {}
