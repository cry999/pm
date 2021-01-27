package persistence

import (
	"context"
	"fmt"
	"os"

	domain "github.com/cry999/pm-projects/pkg/domain/model/task"
	"github.com/cry999/pm-projects/pkg/infrastructure/persistence/models"
	"github.com/rs/xid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MySQLTaskRepository ...
type MySQLTaskRepository struct {
}

// NewMySQLTaskRepository creates a new MySQLTaskRepository instance
func NewMySQLTaskRepository() *MySQLTaskRepository {
	return &MySQLTaskRepository{}
}

// NextIdentity ...
func (r *MySQLTaskRepository) NextIdentity(_ context.Context) (domain.ID, error) {
	return domain.NewID(xid.New().String())
}

// Save ...
func (r *MySQLTaskRepository) Save(ctx context.Context, task *domain.Task) (err error) {
	defer func() {
		err = HandleMySQLError(err, "task.Task", task.ID.String())
	}()

	tx, err := TransactionFromContext(ctx)
	if err != nil {
		return
	}
	taskInDB := r.Serialize(task)
	fmt.Fprintf(os.Stderr, "task:%+v, taskInDB:%+v\n", task, taskInDB)

	exists, err := models.TaskExists(ctx, tx, task.ID.String())
	if err != nil {
		return
	}

	cols := boil.Blacklist(models.TaskColumns.CreatedAt, models.TaskColumns.UpdatedAt)
	if exists {
		_, err = taskInDB.Update(ctx, tx, cols)
	} else {
		err = taskInDB.Insert(ctx, tx, cols)
	}
	if err != nil {
		return
	}
	if err = taskInDB.Reload(ctx, tx); err != nil {
		return
	}
	task.CreatedAt = taskInDB.CreatedAt
	task.UpdatedAt = taskInDB.UpdatedAt
	return
}

// Find ...
func (r *MySQLTaskRepository) Find(ctx context.Context, taskID domain.ID) (_ *domain.Task, err error) {
	defer func() {
		err = HandleMySQLError(err, "task.Task", taskID.String())
	}()

	tx, err := TransactionFromContext(ctx)
	if err != nil {
		return
	}
	taskInDB, err := models.FindTask(ctx, tx, taskID.String())
	if err != nil {
		return
	}
	return r.Deserialize(taskInDB)
}

// FindAllAssociatedWithUser ...
func (r *MySQLTaskRepository) FindAllAssociatedWithUser(ctx context.Context, userID domain.UserID) (tasks []*domain.Task, err error) {
	tx, err := TransactionFromContext(ctx)
	if err != nil {
		return
	}
	tasksInDB, err := models.Tasks(
		models.TaskWhere.OwnerID.EQ(nilIfEmpty(userID.String())),
		qm.Or2(models.TaskWhere.AssigneeID.EQ(nilIfEmpty(userID.String()))),
	).All(ctx, tx)
	if err != nil {
		return
	}
	for _, taskInDB := range tasksInDB {
		task, err := r.Deserialize(taskInDB)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return
}

// Serialize ...
func (r *MySQLTaskRepository) Serialize(task *domain.Task) *models.Task {
	return &models.Task{
		ID:          task.ID.String(),
		Name:        task.Name,
		Description: task.Description,
		OwnerID:     nilIfEmpty(task.OwnerID.String()),
		AssigneeID:  nilIfEmpty(task.AssigneeID.String()),
		ProjectID:   nilIfEmpty(task.ProjectID.String()),
		Status:      string(task.Status),
		Deadline:    null.TimeFromPtr(task.Deadline),
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

// Deserialize ...
func (r *MySQLTaskRepository) Deserialize(taskInDB *models.Task) (_ *domain.Task, err error) {
	taskID, err := domain.NewID(taskInDB.ID)
	if err != nil {
		return
	}
	var (
		ownerID    = domain.UserIDZero
		assigneeID = domain.UserIDZero
		projectID  = domain.ProjectIDZero
	)
	if taskInDB.OwnerID.Valid {
		ownerID, err = domain.NewUserID(taskInDB.OwnerID.String)
		if err != nil {
			return
		}
	}
	if taskInDB.AssigneeID.Valid {
		assigneeID, err = domain.NewUserID(taskInDB.AssigneeID.String)
		if err != nil {
			return
		}
	}
	if taskInDB.ProjectID.Valid {
		projectID, err = domain.NewProjectID(taskInDB.ProjectID.String)
		if err != nil {
			return
		}
	}
	return &domain.Task{
		ID:          taskID,
		Name:        taskInDB.Name,
		Description: taskInDB.Description,
		OwnerID:     ownerID,
		AssigneeID:  assigneeID,
		ProjectID:   projectID,
		Status:      domain.Status(taskInDB.Status),
		Deadline:    taskInDB.Deadline.Ptr(),
		CreatedAt:   taskInDB.CreatedAt,
		UpdatedAt:   taskInDB.UpdatedAt,
	}, nil
}
