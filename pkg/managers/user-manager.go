package managers

import (
	"ibercs/internal/model"
	"ibercs/internal/repositories"

	"gorm.io/gorm"
)

type UserManager struct {
	repoUsers    *repositories.GenericRepository[model.UserModel]
	repoSessions *repositories.GenericRepository[model.UserSessionModel]
}

func NewUserManager(database *gorm.DB) *UserManager {
	users := repositories.NewGenericRepository[model.UserModel](database)
	sessions := repositories.NewGenericRepository[model.UserSessionModel](database)

	return &UserManager{
		repoUsers:    users,
		repoSessions: sessions,
	}
}

func (m *UserManager) GetByID(id int) (*model.UserModel, error) {
	return m.repoUsers.GetByID(int32(id))
}

func (m *UserManager) GetByFaceitId(faceitId string) (*model.UserModel, error) {
	return m.repoUsers.Get(repositories.Where("faceit_id", faceitId))
}

func (m *UserManager) Create(user *model.UserModel) (*model.UserModel, error) {
	return m.repoUsers.Create(user)
}

func (m *UserManager) Update(user *model.UserModel) (*model.UserModel, error) {
	err := m.repoUsers.Update(user, "id", int32(user.ID))
	if err != nil {
		return nil, err
	}

	return m.GetByID(user.ID)
}

// Sessions
func (m *UserManager) GetSessionByUserId(userId int) (*model.UserSessionModel, error) {
	return m.repoSessions.Get(repositories.Where("user_id", userId))
}

func (m *UserManager) CreateNewSession(userId int) (*model.UserSessionModel, error) {
	session := &model.UserSessionModel{
		UserID:    userId,
		SessionID: model.GenerateSessionId(),
	}

	return m.repoSessions.Create(session)
}

func (m *UserManager) DeleteSession(userId int) error {
	session, err := m.GetSessionByUserId(userId)
	if err != nil {
		return err
	}

	return m.repoSessions.Delete("session_id", session.SessionID)
}

// Streams
func (m *UserManager) GetAllStreams() ([]model.UserModel, error) {
	return m.repoUsers.Find(repositories.Where("twitch != ''"))
}
