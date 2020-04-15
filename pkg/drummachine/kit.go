package drummachine

import (
	eightoheight "github.com/IyadAssaf/go-drummachine/pkg/drummachine/kits/808"
	nineohnine "github.com/IyadAssaf/go-drummachine/pkg/drummachine/kits/909"
	"github.com/IyadAssaf/go-drummachine/pkg/midi"
)

type Kit interface {
	GetSoundPath() (string, error)

	Render() error
	Close() error

	Kick() *midi.NoteMessage
	Snare() *midi.NoteMessage
	Clap() *midi.NoteMessage
	HighHatOpen() *midi.NoteMessage
	HighHatClosed() *midi.NoteMessage
	TomLow() *midi.NoteMessage
	TomHigh() *midi.NoteMessage
}

func GetKit(name string) (Kit, error) {
	switch name {
	case eightoheight.KitName:
		return &eightoheight.EightOhEight{}, nil
	case nineohnine.KitName:
		return &nineohnine.NineOhNine{}, nil

	default:
		return &eightoheight.EightOhEight{}, nil
	}
}
