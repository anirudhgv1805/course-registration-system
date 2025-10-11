package ws

type CourseSeatUpdate struct {
    Type       string `json:"type"`
    CourseID   string `json:"course_id"`
    SeatsLeft  int    `json:"seats_left"`
}
