ALTER TABLE passed_lesson_professor_user
    ADD CONSTRAINT passed_lesson_professor_user_unique_combination
        UNIQUE (
                user_id,
                lesson_id
            );