package middleware

import (
	"context"
	"github.com/golang-jwt/jwt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-streams/pkg/logger"
	"log"
	"os"
)

var (
	AuthUnary  grpc.UnaryServerInterceptor
	AuthStream grpc.StreamServerInterceptor
	JWTKey     string
)

func authFunc(ctx context.Context) (context.Context, error) {
	log.Println("AuthFunc")
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "could not read auth token :%v", err)
	}

	parser := new(jwt.Parser)
	//parsedToken, _, err := parser.ParseUnverified(token, &jwt.StandardClaims{})

	parsedToken, err := parser.ParseWithClaims(
		token,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JWTKey), nil
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "could not parsed auth token :%v", err)
	}

	claims := parsedToken.Claims.(*jwt.StandardClaims)
	user_id := claims.Subject
	return context.WithValue(ctx, "user_id", user_id), nil
}

func init() {
	logger.Log.Info("Initializing auth middleware")
	AuthUnary = grpc_auth.UnaryServerInterceptor(authFunc)
	AuthStream = grpc_auth.StreamServerInterceptor(authFunc)
	JWTKey = os.Getenv("OASIS_KEY")
	if JWTKey == "" {
		log.Fatalln("Failed to load JWT Key from envvar OASIS_KEY")
	}
}

//https://github.com/sukesan1984/snippets/blob/39c0c26766bf2384fa664985aa4f8196e8506351/golang/grpc-go-auth/server/authentication.go
