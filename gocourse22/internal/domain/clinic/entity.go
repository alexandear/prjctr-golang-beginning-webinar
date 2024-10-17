package clinic

type Clinic struct {
	ClinicID int    `db:"clinic_id"` // Унікальний ідентифікатор клініки
	Name     string `db:"name"`      // Назва клініки
	Address  string `db:"address"`   // Адреса клініки
	Phone    string `db:"phone"`     // Телефонний номер клініки
	Email    string `db:"email"`     // Електронна пошта для контакту з клінікою
}
