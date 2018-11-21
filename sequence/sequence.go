package sequence

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

var (
	sequencesMu sync.RWMutex
	sequences   = map[string]Sequence{}
)

type Sequence interface {
	// Open opens the sequence generator.
	Open() (err error)
	// NextSequence generates next sequence integer(unsigned 64bit).
	// If some error happens, err will not be nil and seq will be 0.
	// Else, err will be nil and next valid sequence integer will be in seq.
	NextSequence() (seq uint64, err error)
	// Close closes the sequence generator.
	Close()
}

// GetSequence returns corresponding sequence instance with the specified
// sequenceType. If the specified sequenceType does not register itself, then
// err will be non nil and sequence will be nil. Else, err will be nil and
// sequence will be corresponding sequence instance.
func GetSequence(sequenceType string) (sequence Sequence, err error) {
	sequencesMu.RLock()
	defer sequencesMu.RUnlock()

	if value, ok := sequences[sequenceType]; ok {
		sequence = value
		return sequence, nil
	} else {
		return nil, fmt.Errorf("%v is not registered", sequenceType)
	}
}

// Register makes a sequence generator available by the provided sequenceType.
// If Register is called twice with the same name or if driver is nil, it
// panics.
func MustRegister(sequenceType string, sequence Sequence) {
	log.Printf("Registering sequence %v", sequenceType)
	sequencesMu.Lock()
	defer sequencesMu.Unlock()

	if sequence == nil {
		log.Panic("sequence: Registered sequence is nil")
	}

	if _, dup := sequences[sequenceType]; dup {
		log.Panic("sequence: Register called twice for driver " + sequenceType)
	}

	sequences[sequenceType] = sequence
}

// Sequences returns a sorted list of the types of the registered sequences.
func Sequences() []string {
	sequencesMu.RLock()
	defer sequencesMu.RUnlock()

	var list []string
	for name := range sequences {
		list = append(list, name)
	}

	sort.Strings(list)
	return list
}
