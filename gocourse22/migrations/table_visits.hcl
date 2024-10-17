table "visits" {
  schema = schema.public
  column "visit_id" {
    type = serial
    null = false
  }
  column "patient_id" {
    type = int
    null = false
  }
  column "doctor_id" {
    type = int
    null = false
  }
  column "visit_date" {
    type = timestamp
    null = false
  }
  column "reason" {
    type = text
  }
  column "diagnosis" {
    type = text
  }
  column "notes" {
    type = text
  }
  primary_key {
    columns = [column.visit_id]
  }
  foreign_key "fk_patient" {
    columns     = [column.patient_id]
    ref_columns = [table.patients.column.patient_id]
  }
  foreign_key "fk_doctor" {
    columns     = [column.doctor_id]
    ref_columns = [table.doctors.column.doctor_id]
  }
}