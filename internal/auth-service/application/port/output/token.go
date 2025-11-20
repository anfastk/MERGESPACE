package output

type TokenService interface {
    GenerateTokens(userID int64, email string) (string, string, error)
}
