package drummachine

import (
	"context"
	"fmt"
	"time"

	"github.com/IyadAssaf/go-drummachine/pkg/midi"
)

func (s *Synth) SetupSequencer(ctx context.Context) error {

	hhClosed := s.HighHatClosed()
	kick := s.Kick()
	snare := s.Snare()

	pattern := &SequencerPattern{
		Beats: 8,
	}
	pattern.Notes = make(map[int][]*midi.NoteMessage)

	for i := 0; i < pattern.Beats; i++ {
		pattern.Notes[i] = make([]*midi.NoteMessage, 0)
	}

	pattern.Notes[0] = []*midi.NoteMessage{hhClosed, kick}
	pattern.Notes[1] = []*midi.NoteMessage{hhClosed}
	pattern.Notes[2] = []*midi.NoteMessage{hhClosed, snare}
	pattern.Notes[3] = []*midi.NoteMessage{hhClosed}
	pattern.Notes[4] = []*midi.NoteMessage{hhClosed, kick}
	pattern.Notes[5] = []*midi.NoteMessage{hhClosed, kick}
	pattern.Notes[6] = []*midi.NoteMessage{hhClosed, snare}
	pattern.Notes[7] = []*midi.NoteMessage{hhClosed}
	if err := s.scheduleSequencerPattern(ctx, pattern); err != nil {
		return err
	}

	return nil
}

type SequencerPattern struct {
	Beats int
	Notes map[int][]*midi.NoteMessage
}

func (s *Synth) scheduleSequencerPattern(ctx context.Context, pattern *SequencerPattern) error {

	tempo := 120

	startTime := time.Now()
	var t time.Duration

	for {
		//elasped := time.Duration(time.Now().Nanosecond()) - time.Duration(startTime)
		elasped := time.Now().Sub(startTime)
		fmt.Println("time elapsed", elasped, "vs", t)

		if elasped > t || t == 0 {
			t = s.applyPattern(ctx, tempo, pattern, t)
		}
	}

	//TODO figure out how to loop a sequence properly
	//t = s.applyPattern(ctx, tempo, pattern, t)

	return nil
}

func (s *Synth) applyPattern(ctx context.Context, tempo int, pattern *SequencerPattern, t time.Duration) time.Duration {

	fmt.Println("applying pattern with time", t)
	tickSize := time.Minute / time.Duration(tempo)
	beatCount := 0
	for i := 0; i < pattern.Beats; i++ {
		beatNumber := beatCount % pattern.Beats
		fmt.Println("beat number", beatNumber, "time", t)

		for _, b := range pattern.Notes[beatNumber] {
			must(s.Sequencer.ScheduleSendNote(b.Channel, b.Note, b.Velocity, t))
		}
		beatCount++
		t += tickSize
	}
	return t
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
