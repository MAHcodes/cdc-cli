package types

type Stats struct {
	ChessBlitz struct {
		Best struct {
			Date   int    `json:"date"`
			Game   string `json:"game"`
			Rating int    `json:"rating"`
		} `json:"best"`
		Last struct {
			Date   int `json:"date"`
			Rating int `json:"rating"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Record struct {
			Draw int `json:"draw"`
			Loss int `json:"loss"`
			Win  int `json:"win"`
		} `json:"record"`
	} `json:"chess_blitz"`
	ChessRapid struct {
		Best struct {
			Date   int    `json:"date"`
			Game   string `json:"game"`
			Rating int    `json:"rating"`
		} `json:"best"`
		Last struct {
			Date   int `json:"date"`
			Rating int `json:"rating"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Record struct {
			Draw int `json:"draw"`
			Loss int `json:"loss"`
			Win  int `json:"win"`
		} `json:"record"`
	} `json:"chess_rapid"`
	ChessDaily struct {
		Last struct {
			Date   int `json:"date"`
			Rating int `json:"rating"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Record struct {
			Draw           int `json:"draw"`
			Loss           int `json:"loss"`
			TimePerMove    int `json:"time_per_move"`
			TimeoutPercent int `json:"timeout_percent"`
			Win            int `json:"win"`
		} `json:"record"`
	} `json:"chess_daily"`
	Tactics struct {
		Highest struct {
			Date   int `json:"date"`
			Rating int `json:"rating"`
		} `json:"highest"`
		Lowest struct {
			Date   int `json:"date"`
			Rating int `json:"rating"`
		} `json:"lowest"`
	} `json:"tactics"`
	Fide       int      `json:"fide"`
	PuzzleRush struct{} `json:"puzzle_rush"`
}
