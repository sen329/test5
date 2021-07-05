package model

type Rune_color int

const (
	Red Rune_color = iota + 1
	Blue
	Green
	Yellow
	White
)

// func (u *Rune_color) Scan(value interface{}) (err error) {
// 	var skills []string
//     switch value.(type) {
//     case string:
//         err = json.Unmarshal([]byte(src.(string)), &skills)
//     case []byte:
//         err = json.Unmarshal(src.([]byte), &skills)
//     }
//     if err != nil {
//         return
//     }
//     return nil
// 	*u = Rune_color(value.(uint8))
// 	return nil
// }

// func (u Rune_color) Value() (driver.Value, error) {
// 	return int8(u), nil
// }

type Rune struct {
	Rune_id     int    `json:"rune_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rune_color  string `json:"rune_color"`
}
