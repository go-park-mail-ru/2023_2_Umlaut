package static

var (
	Host     = Protocol + Adress
	Protocol = "http://"
	Adress   = "localhost"

	MessageDbField = "id, dialog_id, sender_id, message_text, created_at"
	LikeDbField    = "liked_by_user_id, liked_to_user_id"
	UserDbField    = "id, name, mail, password_hash, salt, user_gender, prefer_gender, description, looking, image_paths, education, hobbies, birthday, online, tags"
	AdminDbField = "id, name, mail, password_hash, salt"
	FeedbackDbField = "id, user_id, rating, liked, need_fix, comment_fix, comment"
)
type Feedback struct {
	Id         int        `json:"id" db:"id"`
	UserId     int        `json:"user_id" db:"user_id"`
	Rating     *int       `json:"rating" db:"rating"`
	Liked      *string    `json:"liked" db:"liked"`
	NeedFix    *string    `json:"need_fix" db:"need_fix"`
	CommentFix *string    `json:"comment_fix" db:"comment_fix"`
	Comment    *string    `json:"comment" db:"comment"`
	Show       bool       `json:"show" db:"show"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at"`
}
