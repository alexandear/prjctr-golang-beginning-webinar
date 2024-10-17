table "clinics" {
  schema = schema.public
  column "clinic_id" {
    type = serial
    null = false
  }
  column "name" {
    type = varchar(255)
    null = false
  }
  column "address" {
    type = text
    null = false
  }
  column "phone" {
    type = varchar(20)
    null = false
  }
  column "email" {
    type = varchar(100)
    null = false
  }
  primary_key {
    columns = [column.clinic_id]
  }
}
