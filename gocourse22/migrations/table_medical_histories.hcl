table "medical_histories" {
  schema = schema.public
  column "history_id" {
    type = serial
    null = false
  }
  column "patient_id" {
    type = int
    null = false
  }
  column "disease" {
    type = varchar(100)
    null = false
  }
  column "treatment" {
    type = text
  }
  column "start_date" {
    type = date
  }
  column "end_date" {
    type = date
  }
  column "notes" {
    type = text
  }
  primary_key {
    columns = [column.history_id]
  }
  foreign_key "fk_patient_history" {
    columns     = [column.patient_id]
    ref_columns = [table.patients.column.patient_id]
  }
}