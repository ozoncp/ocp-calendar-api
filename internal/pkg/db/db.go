package db

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"google.golang.org/grpc"
	"log"
)

var dbKey = "db"

func Connect(dsn string) *sql.DB {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatalf("db connection error %v", err)
	}

	return db
}

func NewContext(ctx context.Context, db *sql.DB) context.Context {
	dbCtx := context.WithValue(ctx, &dbKey, db)
	return dbCtx
}

func FromContext(ctx context.Context) *sql.DB {
	client, ok := ctx.Value(&dbKey).(*sql.DB)

	if !ok {
		panic("Client doesn't found in context")
	}

	return client
}

func GetDB(ctx context.Context) *sql.DB {
	return FromContext(ctx)
}

func NewInterceptorWithDB(db *sql.DB) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		return handler(NewContext(ctx, db), req)
	}
}
