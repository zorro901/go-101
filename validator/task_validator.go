package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-101/model"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

// NewTaskValidator taskValidator構造体のインスタンスを生成するためのコンストラクタを定義
func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),            // タイトル必須
			validation.RuneLength(1, 10).Error("limited max 10 char"), // タイトルは1文字以上10文字以下
		),
	)
}
