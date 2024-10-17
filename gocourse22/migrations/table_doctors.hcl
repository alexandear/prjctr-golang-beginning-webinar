table "doctors" {
  schema = schema.public
  column "doctor_id" {
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
  column "specialization" {
    type = varchar(100)
    null = false
  }
  column "email" {
    type = varchar(100)
    null = false
    attrs = ["unique"]
  }
  column "phone" {
    type = varchar(20)
  }
  primary_key {
    columns = [column.doctor_id]
  }
}