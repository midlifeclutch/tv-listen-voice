# tv-listen-voice

Generate the value required to listen to a teams comms from a CS2 replay. Namely FACEIT or any other service which still includes comms.

There is also the excellent https://github.com/pythaeusone/Faceit-Demo-Voice-Calculator, I created this project first before being aware of the tool created by pythaeusone.

## Acknowledgements

- [markus-wa/demoinfocs-golang](https://github.com/markus-wa/demoinfocs-golang) All the heavy lifting is done by the work of this project. 
- [How to listen to only your team's comms in a demo *CS2 Tutorial*](https://www.youtube.com/watch?v=8rc-ynAkrXw) by bird. The video turned me on to how to listen to comms in demos.

## Installation

- Download `tv-listen-voice.tgz` from [Releases](https://github.com/midlifeclutch/tv-listen-voice/releases/tag/v1.1.0)
- Extract the file `tv-listen-voice.exe` to a directory of your choosing

*My preferred method is to put small tools in `C:\bin` and [Add it to your path](https://www.howtogeek.com/118594/how-to-edit-your-system-path-for-easy-command-line-access/)*

## Usage

- Download a demo from a FACEIT match or other service e.g. https://www.faceit.com/en/cs2/room/1-5e34c677-eb7e-4795-a6ea-e3e4cf151d66
- Extract the demo to your `csgo` directory `..\Counter-Strike Global Offensive\game\csgo`
- In a CMD prompt or PowerShell window, run `tv-listen-voice.exe` with the extracted path of the demo.

```
PS C:\Users\midlifeclutch\Downloads> .\tv-listen-voice.exe 'D:\Steam\steamapps\common\Counter-Strike Global Offensive\game\csgo\study\1-5e34c677-eb7e-4795-a6ea-e3e4cf151d66-1-1.dem'

┌─────────────────────┐
│ team_kye111         │
├──────┬──────────────┤
│ SPEC │ NAME         │
├──────┼──────────────┤
│   10 │ Summer1nmind │
├──────┼──────────────┤
│    4 │ kye111       │
├──────┼──────────────┤
│   11 │ Fixedinhoo   │
├──────┼──────────────┤
│   13 │ fredybr      │
├──────┼──────────────┤
│    6 │ slimreaper74 │
└──────┴──────────────┘
tv_listen_voice_indices 5672; tv_listen_voice_indices_h 5672

┌─────────────────┐
│ team_otuka888   │
├──────┬──────────┤
│ SPEC │ NAME     │
├──────┼──────────┤
│    5 │ zweih666 │
├──────┼──────────┤
│    7 │ otuka888 │
├──────┼──────────┤
│   14 │ Ava      │
├──────┼──────────┤
│    8 │ baz      │
├──────┼──────────┤
│    9 │ obyj     │
└──────┴──────────┘
tv_listen_voice_indices 8656; tv_listen_voice_indices_h 8656
```

### With `tv-listen-voice` in your PATH

```
PS D:\Steam\steamapps\common\Counter-Strike Global Offensive\game\csgo> tv-listen-voice.exe .\study\1-5e34c677-eb7e-4795-a6ea-e3e4cf151d66-1-1.dem
```
