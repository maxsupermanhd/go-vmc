package component

import (
	"io"

	"github.com/maxsupermanhd/go-vmc/v767/nbt/dynbt"
	pk "github.com/maxsupermanhd/go-vmc/v767/net/packet"
)

var _ DataComponent = (*MapDecorations)(nil)

type MapDecorations struct {
	dynbt.Value
}

// ID implements DataComponent.
func (MapDecorations) ID() string {
	return "minecraft:map_decorations"
}

// ReadFrom implements DataComponent.
func (m *MapDecorations) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.NBT(&m.Value).ReadFrom(r)
}

// WriteTo implements DataComponent.
func (m *MapDecorations) WriteTo(w io.Writer) (n int64, err error) {
	return pk.NBT(&m.Value).WriteTo(w)
}
