package engine

// everything represents a delta
// booleans are ints; positive means true, negative means false

type Delta struct {
    PlayerDeltas map[Color]PlayerDelta
    WorkerDeltas map[int]WorkerDelta

    CalendarDelta CalendarDelta
    TemplesDelta TemplesDelta
    ResearchDelta ResearchDelta

    Monuments []int
    Buildings []int

    CurrPlayer int
    FirstPlayer int
    AccumulatedCorn int
    Age int
    Day int
}

type PlayerDelta struct {
    Resources [4]int
    Corn int
    Points int
    CornTiles int
    WoodTiles int
    FreeWorkers int
    WorkerDeduction int

    LightSide int
    Buildings []int
    Monuments []int
}

type WorkerDelta struct {
    // all args required
    Available int
    Wheel_id int
    Position int
}

type CalendarDelta struct {
    WheelDeltas []WheelDelta
    Rotation int

    FirstPlayer int
}

type WheelDelta struct {
    // todo how should this work?
    Occupied map[int]int
    PositionDeltas []PositionDelta
}

type PositionDelta struct {
    PData PalenqueDataDelta
    CData ChichenDataDelta
}

type PalenqueDataDelta struct {
    CornTiles int
    WoodTiles int
}

type ChichenDataDelta struct {
    Full int
}

type TemplesDelta struct {
    TempleDeltas []TempleDelta
}

type TempleDelta struct {
    PlayerLocations map[Color]int
}

type ResearchDelta struct {
    Levels map[Color]LevelsDelta
}

type LevelsDelta map[Science]int

func Bool(d int, m int) bool {
    return d * m > 0
}