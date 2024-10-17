package patient

type Doctor struct {
	DoctorID       int    `db:"doctor_id"`
	FirstName      string `db:"first_name"`
	LastName       string `db:"last_name"`
	Specialization string `db:"specialization"`
	Email          string `db:"email"`
	Phone          string `db:"phone"`
}
