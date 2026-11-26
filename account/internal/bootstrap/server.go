package bootstrap

import (
    "context"
    "fmt"
    "log/slog"
    "net"
    "net/http"
    "os"
    server "github.com/Teexijis/PetAdopt/account/internal/api/account_service_api"
    "github.com/Teexijis/PetAdopt/account/internal/pb/account_api"
    "github.com/go-chi/chi/v5"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    httpSwagger "github.com/swaggo/http-swagger"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func AppRun(api server.AccountServiceAPI) {
    go func() {
        if err := runGRPCServer(api); err != nil {
            panic(fmt.Errorf("ошибка запуска grpc-сервера: %w", err))
        }
    }()

    if err := runGatewayServer(); err != nil {
        panic(fmt.Errorf("ошибка запуска http-сервера: %w", err))
    }
}

func runGRPCServer(api server.AccountServiceAPI) error {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        return err
    }

    s := grpc.NewServer()
    account_api.RegisterAccountServiceServer(s, &api)

    slog.Info("grpc-сервер запущен на порту 50051")
    return s.Serve(lis)
}

func runGatewayServer() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    swaggerPath := os.Getenv("swaggerPath")
    if _, err := os.Stat(swaggerPath); os.IsNotExist(err) {
        panic(fmt.Errorf("swagger файл не найден: %s", swaggerPath))
    }

    r := chi.NewRouter()
    r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, swaggerPath)
    })

    r.Get("/docs/*", httpSwagger.Handler(
        httpSwagger.URL("/swagger.json"),
    ))

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
    err := account_api.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, ":50051", opts)
    if err != nil {
        panic(err)
    }

    r.Mount("/", mux)
    slog.Info("grpc gateway запущен на порту 8080")
    return http.ListenAndServe(":8080", r)
}