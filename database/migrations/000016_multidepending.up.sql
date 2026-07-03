CREATE TABLE multi_depending
(
    professor_id  int NOT NULL references professor (id) on delete cascade,
    lesson_id     int NOT NULL references lesson (id) on delete cascade,
    university_id int NOT NULL references university (id) on delete cascade,
    major_id      int NOT NULL references major (id) on delete cascade,
    PRIMARY KEY (lesson_id, professor_id, university_id, major_id)
);
