package drummachine

import (
	"context"
	"github.com/IyadAssaf/go-drumachine/pkg/midi"
	"github.com/k0kubun/go-keybind"
)

func (s *Synth) SetupInstrument(ctx context.Context) error {
	bind := keybind.Open()
	defer bind.Close()

	endCh := make(chan error)
	defer close(endCh)

	if err := s.Kit.Render(); err != nil {
		return err
	}
	defer s.Kit.Close()

	go func() {
		for {
			key, err := bind.ReadRune()
			if err != nil {
				//TODO do something better than this
				panic(err)
			}

			note := s.switchKey(ctx, key, endCh)
			if note != nil {
				s.synth.NoteOn(note.Channel, note.Note, note.Velocity)
			}
		}
	}()

	switch <-endCh {
	case nil:
		close(endCh)
		break
	default:
		return <-endCh
	}

	return nil
}

func (s *Synth) switchKey(ctx context.Context, key rune, ch chan error) *midi.NoteMessage {
	switch key {
	case 98:
		return s.Kick()
	case 104:
		return s.HighHatClosed()
	case 106:
		return s.HighHatOpen()
	case 115:
		return s.Snare()
	case 99:
		return s.Clap()
	case 116:
		return s.TomHigh()
	case 121:
		return s.TomLow()
	case keybind.CtrlC:
		//TODO cancel context here
		ch <- nil
	//default:
	//	fmt.Println(key)
	//	if keybind.IsPrintable(key) {
	//		fmt.Printf("%c\n", key)
	//	} else {
	//		fmt.Printf("Ctrl+%c\n", '@'+key)
	//	}
	}
	return nil
}
