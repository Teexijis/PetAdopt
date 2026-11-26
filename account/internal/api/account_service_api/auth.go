package account_service_api

import (
    "context"
    "strconv"
    "strings"
    "github.com/Teexijis/PetAdopt/account/internal/services/accountService"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
)


//функция для прверки авторизации чтобы не дублировать
func (a *AccountServiceAPI) GetUserFromToken(ctx context.Context) (int64, string, error) {
    md, _ := metadata.FromIncomingContext(ctx)
    authorization := md.Get("authorization")
    if len(authorization) != 1 {
        return 0, "", status.Error(codes.Unauthenticated, "необходима авторизация")
    }

    parts := strings.Fields(authorization[0])
    if len(parts) != 2 || parts[0] != "Bearer" {
        return 0, "", status.Error(codes.Unauthenticated, "неверный формат токена")
    }

    tokenData, err := accountService.ValidateJWT(parts[1], a.jwtSecret)
    if err != nil {
        return 0, "", status.Error(codes.Unauthenticated, "недействительный токен")
    }

    userID, err := strconv.ParseInt(tokenData.Subject, 10, 64)
    if err != nil || userID <= 0 {
        return 0, "", status.Error(codes.Unauthenticated, "неверный ID пользователя")
    }
    if tokenData.Role != "user" && tokenData.Role != "admin" {
        return 0, "", status.Error(codes.Unauthenticated, "неверная роль пользователя")
    }

    return userID, tokenData.Role, nil
}