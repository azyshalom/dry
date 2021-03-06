package appui

import (
	"image"

	termui "github.com/gizak/termui"
	"github.com/moncho/dry/docker"
	"github.com/moncho/dry/docker/formatter"
	drytermui "github.com/moncho/dry/ui/termui"
)

const (
	statusSymbol = string('\u25A3')
)

//ContainerRow is a Grid row showing runtime information about a container
type ContainerRow struct {
	table     drytermui.Table
	container *docker.Container
	Indicator *drytermui.ParColumn
	ID        *drytermui.ParColumn
	Image     *drytermui.ParColumn
	Command   *drytermui.ParColumn
	Status    *drytermui.ParColumn
	Ports     *drytermui.ParColumn
	Names     *drytermui.ParColumn

	drytermui.Row
}

//NewContainerRow creates a new ContainerRow widget
func NewContainerRow(container *docker.Container, table drytermui.Table) *ContainerRow {
	cf := formatter.NewContainerFormatter(container, true)

	row := &ContainerRow{
		container: container,
		Indicator: drytermui.NewThemedParColumn(DryTheme, statusSymbol),
		ID:        drytermui.NewThemedParColumn(DryTheme, cf.ID()),
		Image:     drytermui.NewThemedParColumn(DryTheme, cf.Image()),
		Command:   drytermui.NewThemedParColumn(DryTheme, cf.Command()),
		Status:    drytermui.NewThemedParColumn(DryTheme, cf.Status()),
		Ports:     drytermui.NewThemedParColumn(DryTheme, cf.Ports()),
		Names:     drytermui.NewThemedParColumn(DryTheme, cf.Names()),
	}
	row.Height = 1
	row.Table = table
	//Columns are rendered following the slice order
	row.Columns = []termui.GridBufferer{
		row.Indicator,
		row.ID,
		row.Image,
		row.Command,
		row.Status,
		row.Ports,
		row.Names,
	}
	if !docker.IsContainerRunning(container) {
		row.markAsNotRunning()
	} else {
		row.markAsRunning()
	}

	return row

}

//Highlighted marks this rows as being highlighted
func (row *ContainerRow) Highlighted() {
	row.changeTextColor(
		termui.Attribute(DryTheme.Fg),
		termui.Attribute(DryTheme.CursorLineBg))
}

//NotHighlighted marks this rows as being not highlighted
func (row *ContainerRow) NotHighlighted() {
	var fg termui.Attribute
	if !docker.IsContainerRunning(row.container) {
		fg = inactiveRowColor
	} else {
		fg = termui.Attribute(DryTheme.ListItem)
	}

	row.changeTextColor(
		fg,
		termui.Attribute(DryTheme.Bg))
}

//Buffer returns this Row data as a termui.Buffer
func (row *ContainerRow) Buffer() termui.Buffer {
	buf := termui.NewBuffer()
	//This set the background of the whole row
	buf.Area.Min = image.Point{row.X, row.Y}
	buf.Area.Max = image.Point{row.X + row.Width, row.Y + row.Height}
	buf.Fill(' ', row.ID.TextFgColor, row.ID.TextBgColor)

	for _, col := range row.Columns {
		buf.Merge(col.Buffer())
	}
	return buf
}

func (row *ContainerRow) changeTextColor(fg, bg termui.Attribute) {

	row.ID.TextFgColor = fg
	row.ID.TextBgColor = bg
	row.Image.TextFgColor = fg
	row.Image.TextBgColor = bg
	row.Command.TextFgColor = fg
	row.Command.TextBgColor = bg
	row.Status.TextFgColor = fg
	row.Status.TextBgColor = bg
	row.Ports.TextFgColor = fg
	row.Ports.TextBgColor = bg
	row.Names.TextFgColor = fg
	row.Names.TextBgColor = bg
}

//markAsNotRunning
func (row *ContainerRow) markAsNotRunning() {
	row.Indicator.TextFgColor = NotRunning
	row.ID.TextFgColor = inactiveRowColor
	row.Image.TextFgColor = inactiveRowColor
	row.Command.TextFgColor = inactiveRowColor
	row.Status.TextFgColor = inactiveRowColor
	row.Ports.TextFgColor = inactiveRowColor
	row.Names.TextFgColor = inactiveRowColor
}

//markAsRunning
func (row *ContainerRow) markAsRunning() {
	row.Indicator.TextFgColor = Running

}
