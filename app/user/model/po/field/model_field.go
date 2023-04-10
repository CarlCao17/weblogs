package field

//go:generate stringer -type=Field -linecomment

type Field int

const (
	// Null Placeholder
	Null Field = iota // null

	UserID    // user_id
	UserName  // user_name
	Password  // password
	UserRole  // user_role
	Email     // email
	FirstName // first_name
	LastName  // last_name
	IsDel     // is_del

	// End Placeholder
	End // end
)
