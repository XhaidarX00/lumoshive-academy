package model

import "time"

type Drivers struct {
	Driver_id    uint16
	Name         string
	Phone_number string
	Vehicle_type string
	Created_at   time.Time
}

// type User struct {
// 	ID               int64
// 	Username         sql.NullString
// 	Age              sql.NullInt64
// 	Email            sql.NullString
// 	BirthDate        sql.NullTime
// 	RegistrationDate sql.NullTime
// }
