package timewheel

import (
    "fmt"
    "log"
    "time"
)

type Elem struct {
    time    time.Duration
    done    bool
}

type Timewheel struct {
    size    int
    elem    *Elem
    clockwise   bool
    anticlockwise bool
}
