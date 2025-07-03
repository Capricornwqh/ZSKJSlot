package i18n_repo

import (
	i18n_entity "SlotGameServer/pkgs/dao/i18n/entity"
	"SlotGameServer/utils"

	"github.com/gin-gonic/gin"
)

type emailI18NRepo struct {
	translator *utils.Translator
}

type EmailI18NRepo interface {
	GetEmailVerifyCode(ctx *gin.Context, lang, siteName, verifyCode string) (string, string, error)
}

func NewEmailI18NRepo(translator *utils.Translator) EmailI18NRepo {
	return &emailI18NRepo{
		translator: translator,
	}
}

// 邮箱验证码
func (r *emailI18NRepo) GetEmailVerifyCode(ctx *gin.Context, lang, siteName, verifyCode string) (string, string, error) {
	if r.translator == nil || len(siteName) <= 0 || len(verifyCode) <= 0 {
		return "", "", utils.ErrParameter
	}

	templateData := &i18n_entity.VerifyCodeTemplateData{
		SiteName:   siteName,
		VerifyCode: verifyCode,
	}

	title, err := utils.I18nTranslator.GetLocalize(lang, i18n_entity.EMAILTPL_VERIFYCODE_TITLE, templateData)
	if err != nil {
		return "", "", err
	}

	body, err := utils.I18nTranslator.GetLocalize(lang, i18n_entity.EMAILTPL_VERIFYCODE_BODY, templateData)
	if err != nil {
		return "", "", err
	}

	return title, body, nil
}
