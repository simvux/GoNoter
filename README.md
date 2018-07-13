# GoNoter

Easily write down things to remember in a CLI written in Go

## Compilation
### Linux/MacOS/BSD or other
`git clone https://github.com/AlmightyFloppyFish/GoNoter.git && cd GoNoter`  
`go get`  
`go build -o gonoter *.go`  
`sudo mv gonoter /usr/bin/gonoter` (This one might not apply on mac, not sure)  
  
### Windows (untested but should theoretically work)
Install Go and git
go into dir
run `go get`
run `go build -o gonoter.exe *.go`
use gonoter.exe to start

## Usage
Create Note: `gonoter -n "This is a note"`  
View Notes:  `gonoter -v`  
Delete Note: `gonoter -d <NoteID>`  
