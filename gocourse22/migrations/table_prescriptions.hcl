table "prescriptions" {
  schema = schema.public
  column "prescription_id" {
    type = serial
    null = false
  }
  column "visit_id" {
    type = int
    null = false
  }
  column "medicine_name" {
    type = varchar(100)
    null = false
  }
  column "dose" {
    type = varchar(50)
  }
  column "frequency" {
    type = varchar(50)
  }
  column "duration" {
    type = varchar(50)
  }
  column "notes" {
    type = text
  }
  primary_key {
    columns = [column.prescription_id]
  }
  foreign_key "fk_visit_prescription" {
    columns     = [column.visit_id]
    ref_columns = [table.visits.column.visit_id]
  }
}