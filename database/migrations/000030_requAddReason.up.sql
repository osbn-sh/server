ALTER TABLE lesson_co_requisite
    ADD COLUMN IF NOT EXISTS rejection_reason TEXT;

ALTER TABLE lesson_pre_requisite
    ADD COLUMN IF NOT EXISTS rejection_reason TEXT;



