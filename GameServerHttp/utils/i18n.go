package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

type Translator struct {
	bundle       *i18n.Bundle
	mapLocalizer map[string]*i18n.Localizer
}

// translation
var I18nTranslator *Translator

// 初始化国际化
func SetupI18N() {
	entries, err := os.ReadDir(Conf.I18N.Dir)
	if err != nil {
		logrus.Fatal(err)
	}
	// 创建一个新的 Bundle，并指定默认语言
	bundle := i18n.NewBundle(language.Chinese)
	if bundle == nil {
		logrus.Fatal("bundle is nil")
	}
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	I18nTranslator = &Translator{bundle: bundle, mapLocalizer: make(map[string]*i18n.Localizer)}
	for _, file := range entries {
		// ignore directory
		if file.IsDir() {
			continue
		}
		// ignore non-YAML file
		if filepath.Ext(file.Name()) != ".yaml" {
			continue
		}

		// 加载 YAML 文件
		_, err := bundle.LoadMessageFile(filepath.Join(Conf.I18N.Dir, file.Name()))
		if err != nil {
			logrus.Fatal(err)
		}

		// 创建 Localizer，用于处理具体的语言环境
		lang := strings.Replace(file.Name(), filepath.Ext(file.Name()), "", 1)
		localizer := i18n.NewLocalizer(bundle, lang)
		if localizer == nil {
			logrus.Fatal("localizer is nil")
		}
		I18nTranslator.mapLocalizer[lang] = localizer
	}
}

// 获取翻译内容
func (t *Translator) GetLocalize(lang, id string, templateData any) (string, error) {
	local, ok := t.mapLocalizer[lang]
	if !ok {
		local = t.mapLocalizer[Conf.I18N.Default]
	}
	if local == nil {
		return "", ErrParameter
	}

	translation, err := local.Localize(&i18n.LocalizeConfig{MessageID: id, TemplateData: templateData})
	if err != nil {
		local = t.mapLocalizer[Conf.I18N.Default]
		if local != nil {
			translation, err = local.Localize(&i18n.LocalizeConfig{MessageID: id, TemplateData: templateData})
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	return translation, nil
}
