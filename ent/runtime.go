// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject/ent/attachment"
	"awesomeProject/ent/comment"
	"awesomeProject/ent/notification"
	"awesomeProject/ent/project"
	"awesomeProject/ent/projectcolumn"
	"awesomeProject/ent/schema"
	"awesomeProject/ent/subtask"
	"awesomeProject/ent/task"
	"awesomeProject/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	attachmentFields := schema.Attachment{}.Fields()
	_ = attachmentFields
	// attachmentDescFileURL is the schema descriptor for file_url field.
	attachmentDescFileURL := attachmentFields[0].Descriptor()
	// attachment.FileURLValidator is a validator for the "file_url" field. It is called by the builders before save.
	attachment.FileURLValidator = attachmentDescFileURL.Validators[0].(func(string) error)
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescText is the schema descriptor for text field.
	commentDescText := commentFields[0].Descriptor()
	// comment.TextValidator is a validator for the "text" field. It is called by the builders before save.
	comment.TextValidator = commentDescText.Validators[0].(func(string) error)
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentFields[1].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescTitle is the schema descriptor for title field.
	notificationDescTitle := notificationFields[0].Descriptor()
	// notification.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	notification.TitleValidator = notificationDescTitle.Validators[0].(func(string) error)
	// notificationDescText is the schema descriptor for text field.
	notificationDescText := notificationFields[1].Descriptor()
	// notification.TextValidator is a validator for the "text" field. It is called by the builders before save.
	notification.TextValidator = notificationDescText.Validators[0].(func(string) error)
	// notificationDescType is the schema descriptor for type field.
	notificationDescType := notificationFields[2].Descriptor()
	// notification.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	notification.TypeValidator = notificationDescType.Validators[0].(func(string) error)
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescName is the schema descriptor for name field.
	projectDescName := projectFields[0].Descriptor()
	// project.NameValidator is a validator for the "name" field. It is called by the builders before save.
	project.NameValidator = projectDescName.Validators[0].(func(string) error)
	projectcolumnFields := schema.ProjectColumn{}.Fields()
	_ = projectcolumnFields
	// projectcolumnDescName is the schema descriptor for name field.
	projectcolumnDescName := projectcolumnFields[0].Descriptor()
	// projectcolumn.NameValidator is a validator for the "name" field. It is called by the builders before save.
	projectcolumn.NameValidator = projectcolumnDescName.Validators[0].(func(string) error)
	// projectcolumnDescColor is the schema descriptor for color field.
	projectcolumnDescColor := projectcolumnFields[2].Descriptor()
	// projectcolumn.ColorValidator is a validator for the "color" field. It is called by the builders before save.
	projectcolumn.ColorValidator = projectcolumnDescColor.Validators[0].(func(string) error)
	subtaskFields := schema.Subtask{}.Fields()
	_ = subtaskFields
	// subtaskDescDescription is the schema descriptor for description field.
	subtaskDescDescription := subtaskFields[1].Descriptor()
	// subtask.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	subtask.DescriptionValidator = subtaskDescDescription.Validators[0].(func(string) error)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescName is the schema descriptor for name field.
	taskDescName := taskFields[2].Descriptor()
	// task.NameValidator is a validator for the "name" field. It is called by the builders before save.
	task.NameValidator = taskDescName.Validators[0].(func(string) error)
	// taskDescTimer is the schema descriptor for timer field.
	taskDescTimer := taskFields[5].Descriptor()
	// task.DefaultTimer holds the default value on creation for the timer field.
	task.DefaultTimer = taskDescTimer.Default.(func() time.Time)
	// taskDescIsDone is the schema descriptor for isDone field.
	taskDescIsDone := taskFields[7].Descriptor()
	// task.DefaultIsDone holds the default value on creation for the isDone field.
	task.DefaultIsDone = taskDescIsDone.Default.(bool)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
}
