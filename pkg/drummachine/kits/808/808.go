package eightoheight

import (
	"github.com/IyadAssaf/go-drumachine/pkg/midi"
	"github.com/IyadAssaf/go-drumachine/pkg/support"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const title = ` #######    #####    #######  
##     ##  ##   ##  ##     ## 
##     ## ##     ## ##     ## 
 #######  ##     ##  #######  
##     ## ##     ## ##     ## 
##     ##  ##   ##  ##     ## 
 #######    #####    #######  `

type EightOhEight struct {
}

const (
	KitName        = "808"
	velocity uint8 = 127
	channel  uint8 = 1
)

func (s *EightOhEight) GetSoundPath() (string, error) {
	return support.GetSoundFilePath(KitName)
}

func (s *EightOhEight) Close() error {
	ui.Close()
	return nil
}

var titleUI *widgets.Paragraph
func (s *EightOhEight) Render() error {
	if err := ui.Init(); err != nil {
		return err
	}

	titleUI = widgets.NewParagraph()
	titleUI.Text = title
	titleUI.SetRect(0, 0, 150, 9)
	titleUI.TextStyle.Fg = ui.ColorWhite
	titleUI.Border = false

	render(-1)
	return nil
}

func getNote(note uint8) *midi.NoteMessage {
	return &midi.NoteMessage{
		Channel:  channel,
		Velocity: velocity,
		Note:     note,
	}
}

func (s *EightOhEight) Kick() *midi.NoteMessage {
	render(0)
	return getNote(36)
}

func (s *EightOhEight) Snare() *midi.NoteMessage {
	render(1)
	return getNote(40)
}

func (s *EightOhEight) Clap() *midi.NoteMessage {
	render(2)
	return getNote(39)
}

func (s *EightOhEight) HighHatClosed() *midi.NoteMessage {
	render(3)
	return getNote(42)
}

func (s *EightOhEight) HighHatOpen() *midi.NoteMessage {
	render(4)
	return getNote(46)
}

func (s *EightOhEight) TomHigh() *midi.NoteMessage {
	render(5)
	return getNote(50)
}

func (s *EightOhEight) TomLow() *midi.NoteMessage {
	render(6)
	return getNote(45)
}

//TODO allow multiple notes to be played at the same time
func render(column int) {
	grid := widgets.NewTable()
	activeColumns := []string{"", "", "", "", "", "", ""}

	if column != -1 {
		activeColumns[column] = "x"
	}
	grid.Rows = [][]string{
		{"BassDrum (b)", "Snare (s)", "Clap (c)", "HHClosed (h)", "HHOpen (j)", "TomHigh (t)", "TomLow (y)"},
		activeColumns,
	}
	grid.TextStyle = ui.NewStyle(ui.ColorWhite)
	grid.TextAlignment = ui.AlignCenter
	grid.SetRect(0, titleUI.Max.Y, 92, titleUI.Max.Y + 5)

	ui.Render(titleUI, grid)
}
