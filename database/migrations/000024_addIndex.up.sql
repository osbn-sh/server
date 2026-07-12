CREATE INDEX idx_vote_option_id ON vote (option_id);
CREATE INDEX idx_vote_professor_id ON vote (professor_id) WHERE professor_id IS NOT NULL;
CREATE INDEX idx_vote_university_id ON vote (university_id) WHERE university_id IS NOT NULL;
CREATE INDEX idx_vote_major_id ON vote (major_id) WHERE major_id IS NOT NULL;
CREATE INDEX idx_vote_lesson_id ON vote (lesson_id) WHERE lesson_id IS NOT NULL;