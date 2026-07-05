ALTER TABLE pending_lesson
    ADD CONSTRAINT unique_pending_lesson_name UNIQUE (name),
    ADD CONSTRAINT unique_pending_lesson_english_name UNIQUE (name_english),
    ADD CONSTRAINT unique_pending_lesson_descripton UNIQUE (description),
    ADD CONSTRAINT unique_pending_lesson_english_descripton UNIQUE (description_english);


ALTER TABLE pending_university
    ADD CONSTRAINT unique_pending_university_name UNIQUE (name),
    ADD CONSTRAINT unique_pending_university_english_name UNIQUE (name_english),
    ADD CONSTRAINT unique_pending_university_descripton UNIQUE (description),
    ADD CONSTRAINT unique_pending_university_english_descripton UNIQUE (description_english);




ALTER TABLE pending_major
    ADD CONSTRAINT unique_pending_major_name UNIQUE (name),
    ADD CONSTRAINT unique_pending_major_english_name UNIQUE (name_english),
    ADD CONSTRAINT unique_pending_major_descripton UNIQUE (description),
    ADD CONSTRAINT unique_pending_major_english_descripton UNIQUE (description_english);

ALTER TABLE pending_professor
    ADD CONSTRAINT unique_pending_professor_name UNIQUE (name),
    ADD CONSTRAINT unique_pending_professor_english_name UNIQUE (name_english),
    ADD CONSTRAINT unique_pending_professor_descripton UNIQUE (description),
    ADD CONSTRAINT unique_pending_professor_english_descripton UNIQUE (description_english);