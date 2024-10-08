package component

import (
	"io"

	"github.com/maxsupermanhd/go-vmc/v767/nbt/dynbt"
	pk "github.com/maxsupermanhd/go-vmc/v767/net/packet"
)

var _ DataComponent = (*CustomData)(nil)

type CustomData struct {
	dynbt.Value
}

// ID implements DataComponent.
func (CustomData) ID() string {
	return "minecraft:custom_data"
}

// ReadFrom implements DataComponent.
func (c *CustomData) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.NBT(c).ReadFrom(r)
}

// WriteTo implements DataComponent.
func (c *CustomData) WriteTo(w io.Writer) (n int64, err error) {
	return pk.NBT(c).WriteTo(w)
}
