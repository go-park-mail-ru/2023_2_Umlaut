package static

var (
	Host     = Protocol + Adress
	Protocol = "http://"
	Adress   = "localhost"

	MessageDbField       = "id, dialog_id, sender_id, message_text, created_at"
	LikeDbField          = "liked_by_user_id, liked_to_user_id"
	UserDbField          = "id, name, mail, password_hash, salt, user_gender, prefer_gender, description, looking, image_paths, education, hobbies, birthday, online, tags"
	AdminDbField         = "id, mail, password_hash, salt"
	FeedbackDbField      = "id, user_id, rating, liked, need_fix, comment_fix, comment, show, created_at"
	ComplaintTypeDbFiend = "id, type_name"
)
