
ALTER TABLE vote
DROP CONSTRAINT vote_unique_combination;

CREATE UNIQUE INDEX vote_uni_university
    ON vote (user_id, option_id, university_id)
    WHERE university_id IS NOT NULL;

CREATE UNIQUE INDEX vote_uni_professor
    ON vote (user_id, option_id, professor_id)
    WHERE professor_id IS NOT NULL;

CREATE UNIQUE INDEX vote_uni_major
    ON vote (user_id, option_id, major_id)
    WHERE major_id IS NOT NULL;

CREATE UNIQUE INDEX vote_uni_lesson
    ON vote (user_id, option_id, lesson_id)
    WHERE lesson_id IS NOT NULL;