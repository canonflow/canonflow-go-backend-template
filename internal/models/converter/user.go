package converter

import (
	"canonflow-golang-backend-template/internal/models/domain"
	"canonflow-golang-backend-template/internal/models/web"
)

func ToUserData(data *domain.User) web.UserData {
	return web.UserData{
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
