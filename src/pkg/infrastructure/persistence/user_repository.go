package persistence

import (
	"context"

	"github.com/cry999/pm-projects/pkg/domain/model/iam"
	"github.com/cry999/pm-projects/pkg/infrastructure/persistence/models"
	"github.com/rs/xid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// MySQLUserRepository ...
type MySQLUserRepository struct {
}

// NewMySQLUserRepository creates a new MySQLRepository instance
func NewMySQLUserRepository() *MySQLUserRepository {
	boil.DebugMode = true
	return &MySQLUserRepository{}
}

// NextIdentity ...
func (r *MySQLUserRepository) NextIdentity(context.Context) (iam.UserID, error) {
	return iam.NewUserID(xid.New().String())
}

// Save ...
func (r *MySQLUserRepository) Save(ctx context.Context, user *iam.User) (err error) {
	tx, err := TransactionFromContext(ctx)
	if err != nil {
		return
	}
	userInDB := r.Serialize(user)
	cols := boil.Blacklist(models.UserColumns.CreatedAt, models.UserColumns.UpdatedAt)

	exists, err := models.UserExists(ctx, tx, user.ID.String())
	if err != nil {
		return
	}
	if exists {
		_, err = userInDB.Update(ctx, tx, cols)
		return HandleMySQLError(err, "iam.User", user.Email)
	}
	return HandleMySQLError(userInDB.Insert(ctx, tx, cols), "iam.User", user.Email)
}

// Find ...
func (r *MySQLUserRepository) Find(ctx context.Context, userID iam.UserID) (_ *iam.User, err error) {
	tx, err := TransactionFromContext(ctx)
	if err != nil {
		return
	}
	userInDB, err := models.FindUser(ctx, tx, userID.String())
	if err != nil {
		return nil, HandleMySQLError(err, "iam.User", userID.String())
	}
	return r.Deserialize(userInDB)
}

// FindByEmail ...
func (r *MySQLUserRepository) FindByEmail(ctx context.Context, email string) (_ *iam.User, err error) {
	tx, err := TransactionFromContext(ctx)
	if err != nil {
		return
	}
	userInDB, err := models.Users(
		models.UserWhere.Email.EQ(email),
	).One(ctx, tx)
	if err != nil {
		return nil, HandleMySQLError(err, "iam.User", email)
	}
	return r.Deserialize(userInDB)
}

// Serialize ...
func (r *MySQLUserRepository) Serialize(user *iam.User) *models.User {
	return &models.User{
		ID:             user.ID.String(),
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
	}
}

// Deserialize ...
func (r *MySQLUserRepository) Deserialize(userInDB *models.User) (_ *iam.User, err error) {
	userID, err := iam.NewUserID(userInDB.ID)
	if err != nil {
		return
	}
	return &iam.User{
		ID:             userID,
		Email:          userInDB.Email,
		HashedPassword: userInDB.HashedPassword,
	}, nil
}
