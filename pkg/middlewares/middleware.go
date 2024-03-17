package middlewares

type Middlewares struct {
	Auth *AuthMiddleware
}

func NewMiddlewares(Auth *AuthMiddleware) *Middlewares {
	return &Middlewares{Auth: Auth}
}
