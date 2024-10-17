table "patients" {
  schema = schema.public
  column "patient_id" {
    type = serial
    null = false
  }
  column "first_name" {
    type = varchar(50)
    null = false
  }
  column "last_name" {
    type = varchar(50)
    null = false
  }
  column "date_of_birth" {
    type = date
    null = false
  }
  column "gender" {
    type = char(1)
    null = false
  }
  column "email" {
    type = varchar(100)
  }
  column "phone" {
    type = varchar(20)
  }
  column "address" {
    type = text
  }
  primary_key {
    columns = [column.patient_id]
  }
}