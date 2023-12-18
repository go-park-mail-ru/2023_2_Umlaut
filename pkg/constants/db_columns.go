package constants

var (
	MessageDbField       = "id, dialog_id, sender_id, recipient_id, message_text, is_read, created_at"
	LikeDbField          = "liked_by_user_id, liked_to_user_id"
	UserDbField          = "id, name, mail, password_hash, salt, user_gender, prefer_gender, description, looking, image_paths, education, hobbies, birthday, role, like_counter, online, tags"
	AdminDbField         = "id, mail, password_hash, salt"
	FeedbackDbField      = "id, user_id, rating, liked, need_fix, comment, created_at"
	ComplaintTypeDbFiend = "id, type_name"
	ComplaintDbFiend     = "id, reporter_user_id, reported_user_id, complaint_type_id, complaint_text, created_at"
)
