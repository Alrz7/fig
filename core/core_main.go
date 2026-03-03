package core

type handeler struct {
	dir        string
	path       string
	intData    cfInt
	stringData cfString
	listData   cflist
	// type & etc.
}

// var Handelers = map[string]handeler{}

func CreateNewHandeler(dir, path string) *handeler {
	newHandeler := handeler{dir: dir, path: path, intData: cfInt{}, stringData: cfString{}, listData: cflist{}}
	return &newHandeler
}

func (h *handeler) Wrap() {
}

// we do the initializayions here
// we make a manager to manage every file and it makes handeler for each
// we make a manager to manage the encoding and decoding all the configs for each files
// there should be a Global Get function to search all and a Specified get function for each type
