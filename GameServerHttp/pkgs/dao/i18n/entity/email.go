package i18n_entity

const (
	EMAILTPL_CHANGEEMAIL_TITLE = "email_tpl.change_email.title"
	EMAILTPL_CHANGEEMAIL_BODY  = "email_tpl.change_email.body"

	EMAILTPL_NEWANSWER_TITLE = "email_tpl.new_answer.title"
	EMAILTPL_NEWANSWER_BODY  = "email_tpl.new_answer.body"

	EMAILTPL_NEWCOMMENT_TITLE = "email_tpl.new_comment.title"
	EMAILTPL_NEWCOMMENT_BODY  = "email_tpl.new_comment.body"

	EMAILTPL_PASSRESET_TITLE = "email_tpl.pass_reset.title"
	EMAILTPL_PASSRESET_BODY  = "email_tpl.pass_reset.body"

	EMAILTPL_REGISTER_TITLE = "email_tpl.register.title"
	EMAILTPL_REGISTER_BODY  = "email_tpl.register.body"

	EMAILTPL_VERIFYCODE_TITLE = "email_tpl.verify_code.title"
	EMAILTPL_VERIFYCODE_BODY  = "email_tpl.verify_code.body"

	EMAILTPL_TEST_TITLE = "email_tpl.test.title"
	EMAILTPL_TEST_BODY  = "email_tpl.test.body"

	EMAILTPL_INVITEDANSWER_TITLE = "email_tpl.invited_you_to_answer.title"
	EMAILTPL_INVITEDANSWER_BODY  = "email_tpl.invited_you_to_answer.body"

	EMAILTPL_NEWQUESTION_TITLE = "email_tpl.new_question.title"
	EMAILTPL_NEWQUESTION_BODY  = "email_tpl.new_question.body"
)

type RegisterTemplateData struct {
	SiteName    string
	RegisterUrl string
}

type VerifyCodeTemplateData struct {
	SiteName   string
	VerifyCode string
}

type PassResetTemplateData struct {
	SiteName     string
	PassResetUrl string
}

type ChangeEmailTemplateData struct {
	SiteName       string
	ChangeEmailUrl string
}

type TestTemplateData struct {
	SiteName string
}

type NewAnswerTemplateRawData struct {
	AnswerUserDisplayName string
	QuestionTitle         string
	QuestionID            string
	AnswerID              string
	AnswerSummary         string
	UnsubscribeCode       string
}

type NewAnswerTemplateData struct {
	SiteName       string
	DisplayName    string
	QuestionTitle  string
	AnswerUrl      string
	AnswerSummary  string
	UnsubscribeUrl string
}

type NewInviteAnswerTemplateRawData struct {
	InviterDisplayName string
	QuestionTitle      string
	QuestionID         string
	UnsubscribeCode    string
}

type NewInviteAnswerTemplateData struct {
	SiteName       string
	DisplayName    string
	QuestionTitle  string
	InviteUrl      string
	UnsubscribeUrl string
}

type NewCommentTemplateRawData struct {
	CommentUserDisplayName string
	QuestionTitle          string
	QuestionID             string
	AnswerID               string
	CommentID              string
	CommentSummary         string
	UnsubscribeCode        string
}

type NewCommentTemplateData struct {
	SiteName       string
	DisplayName    string
	QuestionTitle  string
	CommentUrl     string
	CommentSummary string
	UnsubscribeUrl string
}

type NewQuestionTemplateRawData struct {
	QuestionAuthorUserID string
	QuestionTitle        string
	QuestionID           string
	UnsubscribeCode      string
	Tags                 []string
	TagIDs               []string
}

type NewQuestionTemplateData struct {
	SiteName       string
	QuestionTitle  string
	QuestionUrl    string
	Tags           string
	UnsubscribeUrl string
}
