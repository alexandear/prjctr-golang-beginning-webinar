table "schedules" {
  schema = schema.public
  column "schedule_id" {
    type = serial
    null = false
  }
  column "doctor_id" {
    type = int
    null = false
  }
  column "work_day" {
    type = date
    null = false
  }
  column "start_time" {
    type = time
    null = false
  }
  column "end_time" {
    type = time
    null = false
  }
  column "notes" {
    type = text
  }
  primary_key {
    columns = [column.schedule_id]
  }
  foreign_key "fk_doctor_schedule" {
    columns     = [column.doctor_id]
    ref_columns = [table.doctors.column.doctor_id]
  }
}