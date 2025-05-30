package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

func main() {
	f, err := os.Open("1-1910b6c8-6926-4468-9139-8c1c7819d730-1-1.dem")
	if err != nil {
		log.Panic("failed to open demo file")
	}
	defer f.Close()

	p := demoinfocs.NewParser(f)
	defer p.Close()

	p.RegisterEventHandler(func(e events.MatchStart) {
		gState := p.GameState()

		if clanName := gState.TeamCounterTerrorists().ClanName(); clanName != "" {
			ctEnts := []int{}
			for _, player := range gState.TeamCounterTerrorists().Members() {
				ctEnts = append(ctEnts, player.EntityID)
			}

			ctEntsDec := binToDec(ctEnts)
			fmt.Printf("tv_listen_voice_indices %d; tv_listen_voice_indices_h %d // %s\n", ctEntsDec, ctEntsDec, clanName)
		}

		if clanName := gState.TeamTerrorists().ClanName(); clanName != "" {
			tEnts := []int{}
			for _, player := range gState.TeamTerrorists().Members() {
				tEnts = append(tEnts, player.EntityID)
			}

			tEntsDec := binToDec(tEnts)
			fmt.Printf("tv_listen_voice_indices %d; tv_listen_voice_indices_h %d // %s\n", tEntsDec, tEntsDec, clanName)
		}

	})

	err = p.ParseToEnd()
	if err != nil {
		log.Panic("failed to parse demo: ", err)
	}

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
