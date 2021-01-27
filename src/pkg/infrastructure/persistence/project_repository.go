package persistence

import (
	"context"

	"github.com/cry999/pm-projects/pkg/domain/model/project"
	"github.com/cry999/pm-projects/pkg/infrastructure/persistence/models"
	"github.com/rs/xid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// MySQLProjectRepository ...
type MySQLProjectRepository struct {
}

// NewMySQLProjectRepository creates a new MySQLProjectRepository instance
func NewMySQLProjectRepository() *MySQLProjectRepository {
	return &MySQLProjectRepository{}
}

// NextIdentity ...
func (r *MySQLProjectRepository) NextIdentity(_ context.Context) (project.ID, error) {
	return project.NewID(xid.New().String())
}

// Save ...
func (r *MySQLProjectRepository) Save(ctx context.Context, prj *project.Project) (err error) {
	defer func() {
		err = HandleMySQLError(err, "project", prj.ID.String())
	}()

	tx, err := TransactionFromContext(ctx)
	if err != nil {
		return
	}

	exists, err := models.ProjectExists(ctx, tx, prj.ID.String())
	if err != nil {
		return
	}
	projectInDB := r.Serialize(prj)
	if exists {
		_, err = projectInDB.Update(ctx, tx, boil.Infer())
	} else {
		err = projectInDB.Insert(ctx, tx, boil.Infer())
	}
	if err != nil {
		return
	}

	// * For simple, at first all plannedProjectTasks are deleted,
	// * and then current all plannedProjectTasks are inserted
	if err = projectInDB.R.PlannedProjectTasks.ReloadAll(ctx, tx); err != nil {
		return
	}
	if _, err = projectInDB.R.PlannedProjectTasks.DeleteAll(ctx, tx); err != nil {
		return
	}
	for _, plannedTaskID := range prj.PlannedTasks {
		if err = projectInDB.AddPlannedProjectTasks(ctx, tx, true, &models.PlannedProjectTask{
			ProjectID:     prj.ID.String(),
			PlannedTaskID: plannedTaskID.String(),
		}); err != nil {
			return
		}
	}

	if err = projectInDB.Reload(ctx, tx); err != nil {
		return
	}
	prj.CreatedAt = projectInDB.CreatedAt
	prj.UpdatedAt = projectInDB.UpdatedAt
	return
}

// FindByID ...
func (r *MySQLProjectRepository) FindByID(ctx context.Context, projectID project.ID) (_ *project.Project, err error) {
	defer func() {
		err = HandleMySQLError(err, "project", projectID.String())
	}()

	tx, err := TransactionFromContext(ctx)
	if err != nil {
		return
	}

	projectInDB, err := models.Projects(
		qm.Load(models.ProjectRels.PlannedProjectTasks),
		models.ProjectWhere.ID.EQ(projectID.String()),
	).One(ctx, tx)
	if err != nil {
		return
	}
	return r.Deserialize(projectInDB)
}

// FindAllRelatedWithUser ...
func (r *MySQLProjectRepository) FindAllRelatedWithUser(ctx context.Context, userID project.UserID) (projects []*project.Project, err error) {
	defer func() {
		err = HandleMySQLError(err, "project", "")
	}()

	tx, err := TransactionFromContext(ctx)
	if err != nil {
		return
	}

	projectsInDB, err := models.Projects(
		qm.Load(models.ProjectRels.PlannedProjectTasks),
		models.ProjectWhere.OwnerID.EQ(userID.String()),
	).All(ctx, tx)
	if err != nil {
		return
	}

	for _, projectInDB := range projectsInDB {
		prj, err := r.Deserialize(projectInDB)
		if err != nil {
			return nil, err
		}
		projects = append(projects, prj)
	}
	return
}

// Serialize ...
func (r *MySQLProjectRepository) Serialize(prj *project.Project) *models.Project {
	prjInDB := &models.Project{
		ID:            prj.ID.String(),
		OwnerID:       prj.OwnerID.String(),
		Name:          prj.Name,
		ElevatorPitch: prj.ElevatorPitch,
		CreatedAt:     prj.CreatedAt,
		UpdatedAt:     prj.UpdatedAt,
	}
	// * a below line is for initializing prjInDB.R.
	prjInDB.AddPlannedProjectTasks(context.Background(), nil, false)
	for _, plannedTaskID := range prj.PlannedTasks {
		prjInDB.R.PlannedProjectTasks = append(prjInDB.R.PlannedProjectTasks, &models.PlannedProjectTask{
			ProjectID:     prj.ID.String(),
			PlannedTaskID: plannedTaskID.String(),
		})
	}
	return prjInDB
}

// Deserialize ...
func (r *MySQLProjectRepository) Deserialize(rec *models.Project) (*project.Project, error) {
	id, err := project.NewID(rec.ID)
	if err != nil {
		return nil, err
	}
	ownerID, err := project.NewUserID(rec.OwnerID)
	if err != nil {
		return nil, err
	}
	prj := &project.Project{
		ID:            id,
		OwnerID:       ownerID,
		Name:          rec.Name,
		ElevatorPitch: rec.ElevatorPitch,
		PlannedTasks:  []project.PlannedTaskID{},
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
	for _, plannedTask := range rec.R.PlannedProjectTasks {
		plannedTaskID, err := project.NewPlannedTaskID(plannedTask.PlannedTaskID)
		if err != nil {
			return nil, err
		}
		prj.PlannedTasks = append(prj.PlannedTasks, plannedTaskID)
	}
	return prj, nil
}
