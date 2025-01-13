package types

type Position struct {
	wheel_id int
	corn int
	GetOptions Options
	pData *PalenqueData
	cData *ChichenData
}

type SpecificPosition struct {
	wheel_id int
	corn int
	Execute Option
	firstPlayer bool
}

type PalenqueData struct {
	cornTiles int
	woodTiles int
}

type ChichenData struct {
	full bool
}
