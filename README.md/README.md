# GoNoter

Easily write down things to remember in a CLI written in Go

## Compilation
### Linux/MacOS/BSD or other
`git clone https://github.com/AlmightyFloppyFish/GoNoter.git && cd GoNoter`  
`go get`  
`go build -o gonoter *.go`  
`sudo mv gonoter /usr/bin/gonoter` (This one might not apply on mac, not sure)  
  
## Usage
Create Note: `gonoter -n "This is a note"`  
View Notes:  `gonoter -v`  
Delete Note: `gonoter -d <NoteID>`  
