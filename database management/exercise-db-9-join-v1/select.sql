-- TODO: answer here
SELECT
        reports.id,
        students.fullname,
        students.class,
        students.status,
        reports.study,
        reports.score
    FROM students
    INNER JOIN reports
    ON students.id = reports.student_id
    WHERE students.status = 'active' AND reports.score < 70
    ORDER BY reports.score ASC;