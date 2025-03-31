// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttachmentsColumns holds the columns for the "attachments" table.
	AttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "file_url", Type: field.TypeString},
		{Name: "task_attachments", Type: field.TypeInt, Nullable: true},
	}
	// AttachmentsTable holds the schema information for the "attachments" table.
	AttachmentsTable = &schema.Table{
		Name:       "attachments",
		Columns:    AttachmentsColumns,
		PrimaryKey: []*schema.Column{AttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attachments_tasks_attachments",
				Columns:    []*schema.Column{AttachmentsColumns[2]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "comment_author", Type: field.TypeInt},
		{Name: "task_comments", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_users_author",
				Columns:    []*schema.Column{CommentsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "comments_tasks_comments",
				Columns:    []*schema.Column{CommentsColumns[4]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
	}
	// NotificationsColumns holds the columns for the "notifications" table.
	NotificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "text", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
	}
	// NotificationsTable holds the schema information for the "notifications" table.
	NotificationsTable = &schema.Table{
		Name:       "notifications",
		Columns:    NotificationsColumns,
		PrimaryKey: []*schema.Column{NotificationsColumns[0]},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// ProjectColumnsColumns holds the columns for the "project_columns" table.
	ProjectColumnsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "order", Type: field.TypeInt},
		{Name: "color", Type: field.TypeString},
		{Name: "project_columns", Type: field.TypeInt, Nullable: true},
	}
	// ProjectColumnsTable holds the schema information for the "project_columns" table.
	ProjectColumnsTable = &schema.Table{
		Name:       "project_columns",
		Columns:    ProjectColumnsColumns,
		PrimaryKey: []*schema.Column{ProjectColumnsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_columns_projects_columns",
				Columns:    []*schema.Column{ProjectColumnsColumns[4]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubtasksColumns holds the columns for the "subtasks" table.
	SubtasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "task_subtasks", Type: field.TypeInt, Nullable: true},
	}
	// SubtasksTable holds the schema information for the "subtasks" table.
	SubtasksTable = &schema.Table{
		Name:       "subtasks",
		Columns:    SubtasksColumns,
		PrimaryKey: []*schema.Column{SubtasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subtasks_tasks_subtasks",
				Columns:    []*schema.Column{SubtasksColumns[3]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "order", Type: field.TypeInt},
		{Name: "type", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "tags", Type: field.TypeJSON},
		{Name: "timer", Type: field.TypeTime},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "is_done", Type: field.TypeBool, Default: false},
		{Name: "project_column_tasks", Type: field.TypeInt, Nullable: true},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_project_columns_tasks",
				Columns:    []*schema.Column{TasksColumns[9]},
				RefColumns: []*schema.Column{ProjectColumnsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserProjectsColumns holds the columns for the "user_projects" table.
	UserProjectsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "project_id", Type: field.TypeInt},
	}
	// UserProjectsTable holds the schema information for the "user_projects" table.
	UserProjectsTable = &schema.Table{
		Name:       "user_projects",
		Columns:    UserProjectsColumns,
		PrimaryKey: []*schema.Column{UserProjectsColumns[0], UserProjectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_projects_user_id",
				Columns:    []*schema.Column{UserProjectsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_projects_project_id",
				Columns:    []*schema.Column{UserProjectsColumns[1]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserTasksColumns holds the columns for the "user_tasks" table.
	UserTasksColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "task_id", Type: field.TypeInt},
	}
	// UserTasksTable holds the schema information for the "user_tasks" table.
	UserTasksTable = &schema.Table{
		Name:       "user_tasks",
		Columns:    UserTasksColumns,
		PrimaryKey: []*schema.Column{UserTasksColumns[0], UserTasksColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_tasks_user_id",
				Columns:    []*schema.Column{UserTasksColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_tasks_task_id",
				Columns:    []*schema.Column{UserTasksColumns[1]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttachmentsTable,
		CommentsTable,
		MessagesTable,
		NotificationsTable,
		ProjectsTable,
		ProjectColumnsTable,
		SubtasksTable,
		TasksTable,
		UsersTable,
		UserProjectsTable,
		UserTasksTable,
	}
)

func init() {
	AttachmentsTable.ForeignKeys[0].RefTable = TasksTable
	CommentsTable.ForeignKeys[0].RefTable = UsersTable
	CommentsTable.ForeignKeys[1].RefTable = TasksTable
	ProjectColumnsTable.ForeignKeys[0].RefTable = ProjectsTable
	SubtasksTable.ForeignKeys[0].RefTable = TasksTable
	TasksTable.ForeignKeys[0].RefTable = ProjectColumnsTable
	UserProjectsTable.ForeignKeys[0].RefTable = UsersTable
	UserProjectsTable.ForeignKeys[1].RefTable = ProjectsTable
	UserTasksTable.ForeignKeys[0].RefTable = UsersTable
	UserTasksTable.ForeignKeys[1].RefTable = TasksTable
}
