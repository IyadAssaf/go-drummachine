package drummachine

import (
	"fmt"
	"github.com/IyadAssaf/fluidsynth"
)

type Synth struct {
	synth  fluidsynth.Synth
	driver fluidsynth.AudioDriver
	Kit
	Sequencer *fluidsynth.Sequencer
}

func (s *Synth) Close() {
	fmt.Println("Closing driver")
	s.driver.Delete()
}

func NewDrumMachine(kitName string) (*Synth, error) {
	settings := fluidsynth.NewSettings()
	synth := fluidsynth.NewSynth(settings)

	driver := fluidsynth.NewAudioDriver(settings, synth)

	seq := fluidsynth.NewSequencer()
	seq.RegisterSynth(synth)

	kit, err := GetKit(kitName)
	if err != nil {
		return nil, err
	}

	soundPath, err := kit.GetSoundPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get soundpath")
	}

	if synth.SFLoad(soundPath, true) != 1 {
		return nil, fmt.Errorf("failed to load soundfile")
	}

	return &Synth{
		synth: synth,
		driver: driver,
		Kit: kit,
		Sequencer: seq,
	}, nil
}