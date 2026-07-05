ALTER TABLE pending_lesson
    ADD COLUMN action VARCHAR(20) NOT NULL DEFAULT 'create',
    ADD COLUMN target_id INT;

ALTER TABLE pending_major
    ADD COLUMN action VARCHAR(20) NOT NULL DEFAULT 'create',
    ADD COLUMN target_id INT;

ALTER TABLE pending_professor
    ADD COLUMN action VARCHAR(20) NOT NULL DEFAULT 'create',
    ADD COLUMN target_id INT;

ALTER TABLE pending_university
    ADD COLUMN action VARCHAR(20) NOT NULL DEFAULT 'create',
    ADD COLUMN target_id INT;


ALTER TABLE pending_lesson
    ADD CONSTRAINT pending_lesson_action_check
        CHECK (action IN ('create', 'update'));

ALTER TABLE pending_major
    ADD CONSTRAINT pending_major_action_check
        CHECK (action IN ('create', 'update'));

ALTER TABLE pending_professor
    ADD CONSTRAINT pending_professor_action_check
        CHECK (action IN ('create', 'update'));

ALTER TABLE pending_university
    ADD CONSTRAINT pending_university_action_check
        CHECK (action IN ('create', 'update'));