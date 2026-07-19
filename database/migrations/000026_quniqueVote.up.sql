ALTER TABLE vote
    ADD CONSTRAINT vote_unique_combination
        UNIQUE (
                user_id,
                option_id,
                university_id,
                professor_id,
                major_id,
                lesson_id
            );



ALTER TABLE option
    ADD CONSTRAINT option_unique_combination
        UNIQUE (
                name,
                owner
            );