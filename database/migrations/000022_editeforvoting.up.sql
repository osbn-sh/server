-- حذف فیلدهای target_id و target_type از جدول option
ALTER TABLE option
    DROP COLUMN target_id,
    DROP COLUMN target_type;

-- اضافه کردن فیلدهای مرجع به جدول vote
ALTER TABLE vote
    ADD COLUMN university_id int REFERENCES university (id) on delete cascade,
    ADD COLUMN professor_id  int REFERENCES professor (id) on delete cascade,
    ADD COLUMN major_id      int REFERENCES major (id) on delete cascade,
    ADD COLUMN lesson_id     int REFERENCES lesson (id) on delete cascade;