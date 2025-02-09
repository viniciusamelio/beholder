package models

import (
	"beholder-api/internal/jet/model"
	"beholder-api/internal/utils"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	ID            int        `json:"id"`
	EnvironmentID int        `json:"environment_uid"`
	UserID        string     `json:"user_id"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Tags          []string   `json:"tags"`

	Env *Environment `json:"env"`
}

func SessionFromDataModel(model model.Sessions) utils.Either[utils.Failure, *Session] {
	ID, err := strconv.Atoi(*model.ID)

	if err != nil {
		return utils.FailureOf[*Session](err, 500)
	}
	EnvironmentID, err := strconv.Atoi(*model.EnvironmentID)
	if err != nil {
		return utils.FailureOf[*Session](err, 500)
	}

	return utils.NewRight[utils.Failure](
		&Session{
			ID:            ID,
			EnvironmentID: EnvironmentID,
			UserID:        *model.UserID,
			CreatedAt:     model.CreatedAt,
			UpdatedAt:     model.UpdatedAt,
			Tags:          strings.Split(*model.Tags, ", "),
		},
	)
}
