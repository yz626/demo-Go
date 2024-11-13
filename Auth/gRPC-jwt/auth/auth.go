package auth

import (
	"context"
	"fmt"
	utils2 "gRPC-jwt/utils"
	"google.golang.org/grpc/metadata"
)

type Token struct {
	token string
}

func NewToken(token string) *Token {
	return &Token{token: token}
}

// GetRequestMetadata 获取认证信息
func (t *Token) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": t.token,
	}, nil
}

// RequireTransportSecurity 是否使用TLS加密传输
func (t *Token) RequireTransportSecurity() bool {
	return false
}

type Auth struct{}

// GetTokenFromContext 从元数据中获取token
func (a Auth) GetTokenFromContext(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	// token的key是authorization
	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", false
	}
	return token[0], true
}

// CheckToken 校验token是否有效
func (a Auth) CheckToken(ctx context.Context) (string, error) {
	token, ok := a.GetTokenFromContext(ctx)
	if !ok {
		return "", fmt.Errorf("get token from context error")
	}

	claims, err := utils2.ParseToken(token)
	if err != nil {
		return "", err
	}

	return claims.Name, nil
}
