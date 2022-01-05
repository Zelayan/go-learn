package datastruct

type Person struct {
	ID          int64  `db:"id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Verified    bool   `db:"verified"`
	EmailCode   int32  `db:"email_code"`
	Balance     int64  `db:"balance"`
}

