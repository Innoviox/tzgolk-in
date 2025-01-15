package engine

// anything with "Delta" in the name represents change
// anything without is just a setter

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
    ResourcesDelta [4]int
    CornDelta int
    PointsDelta int
    CornTilesDelta int
    WoodTilesDelta int
    FreeWorkersDelta int
    WorkerDeductionDelta int

    LightSide bool
    Buildings []Building
    Monuments []Monument
}

type WorkerDelta struct {
    Available bool
    Wheel_id int
    Position int
}

type CalendarDelta struct {
    WheelDeltas []WheelDelta
    RotationDelta int

    FirstPlayer int
}

type WheelDelta struct {
    Occupied map[int]int
    PositionDeltas []PositionDelta
}

type PositionDelta struct {
    PDataDelta PalenqueDataDelta
    CDataDelta ChichenDataDelta
}

type PalenqueDataDelta struct {
    CornTilesDelta int
    WoodTilesDelta int
}

type ChichenDataDelta struct {
    Full bool
}

type TemplesDelta struct {
    TempleDeltas []TempleDelta
}

type TempleDelta struct {
    PlayerLocationsDelta map[Color]int
}

type ResearchDelta struct {
    LevelsDelta map[Color]LevelsDelta
}

type LevelsDelta map[Science]int