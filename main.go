package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

func main() {
	if l := len(os.Args); l < 2 {
		usage()
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Panic("failed to open demo file")
	}
	defer f.Close()

	p := demoinfocs.NewParser(f)
	defer p.Close()

	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true

	p.RegisterEventHandler(func(e events.MatchStart) {
		gState := p.GameState()

		if clanName := gState.TeamCounterTerrorists().ClanName(); clanName != "" {

			ctEnts := []int{}
			ctSpecPlayer := make(map[string]int)
			for _, player := range gState.TeamCounterTerrorists().Members() {
				ctEnts = append(ctEnts, player.EntityID)
				ctSpecPlayer[player.Name] = player.EntityID
			}

			ctEntsDec := binToDec(ctEnts)

			t.SetTitle(fmt.Sprintf(clanName))
			t.AppendHeader(table.Row{"spec", "name"})
			t.SetCaption(fmt.Sprintf("tv_listen_voice_indices %d; tv_listen_voice_indices_h %d", ctEntsDec, ctEntsDec))

			for name, idx := range ctSpecPlayer {
				t.AppendRow(table.Row{idx, name})
			}

			fmt.Println(t.Render())

		}

		t.ResetRows()
		fmt.Println()

		if clanName := gState.TeamTerrorists().ClanName(); clanName != "" {
			tEnts := []int{}
			tSpecPlayer := make(map[string]int)
			for _, player := range gState.TeamTerrorists().Members() {
				tEnts = append(tEnts, player.EntityID)
				tSpecPlayer[player.Name] = player.EntityID
			}

			tEntsDec := binToDec(tEnts)

			t.SetTitle(fmt.Sprintf(clanName))
			t.SetCaption(fmt.Sprintf("tv_listen_voice_indices %d; tv_listen_voice_indices_h %d", tEntsDec, tEntsDec))

			for name, idx := range tSpecPlayer {
				t.AppendRow(table.Row{idx, name})
			}

			fmt.Println(t.Render())

		}

	})

	err = p.ParseToEnd()
	if err != nil {
		log.Panic("failed to parse demo: ", err)
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: tv-listen-voice.exe [demofile]\n")
	os.Exit(1)
}

func binToDec(positions []int) uint64 {
	width := 32
	bits := make([]rune, width)
	for i := 0; i < width; i++ {
		bits[i] = '0'
	}

	for _, pos := range positions {
		if pos >= 0 && pos < width {
			bits[width-pos] = '1'
		}
	}

	decValue, err := strconv.ParseUint(string(bits), 2, 32)
	if err != nil {
		log.Fatalln("binary -> decimal: ", err)
	}

	return decValue
}
