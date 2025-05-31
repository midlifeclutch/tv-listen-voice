# tv-listen-voice

Generate the value required to listen to a teams comms from a CS2 replay. Namely FACEIT or any other service which still includes comms.

## Acknowledgements

- [markus-wa/demoinfocs-golang](https://github.com/markus-wa/demoinfocs-golang) All the heavy lifting is done by the work of this project. 
- [How to listen to only your team's comms in a demo *CS2 Tutorial*](https://www.youtube.com/watch?v=8rc-ynAkrXw) by bird. The video turned me on to how to listen to comms in demos.

## Usage

- Download a demo from a FACEIT match or other service e.g. https://www.faceit.com/en/cs2/room/1-5e34c677-eb7e-4795-a6ea-e3e4cf151d66
- Download `tv-listen-voice.exe` from the [Releases](https://github.com/midlifeclutch/tv-listen-voice/releases) page
- In a CMD prompt or PowerShell window, run `tv-listen-voice.exe` with the extracted path of the demo.

```
PS C:\Users\midlifeclutch\Downloads> .\tv-listen-voice.exe .\1-5e34c677-eb7e-4795-a6ea-e3e4cf151d66-1-1.dem\1-5e34c677-eb7e-4795-a6ea-e3e4cf151d66-1-1.dem
tv_listen_voice_indices 5672; tv_listen_voice_indices_h 5672 // team_kye111
tv_listen_voice_indices 8656; tv_listen_voice_indices_h 8656 // team_otuka888
```
