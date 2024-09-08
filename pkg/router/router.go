package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/akosgarai/projectregister/pkg/controller"
	"github.com/akosgarai/projectregister/pkg/model"
	"github.com/akosgarai/projectregister/pkg/render"
	"github.com/akosgarai/projectregister/pkg/session"
	"github.com/akosgarai/projectregister/pkg/storage"
)

// LoggingResponseWriter is a wrapper for the http.ResponseWriter to store the status code.
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int
}

// NewLoggingResponseWriter is a constructor for the LoggingResponseWriter.
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK, 0}
}

// WriteHeader is a wrapper for the http.ResponseWriter WriteHeader method.
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Write is a wrapper for the http.ResponseWriter Write method.
func (lrw *LoggingResponseWriter) Write(b []byte) (int, error) {
	n, err := lrw.ResponseWriter.Write(b)
	lrw.bytesWritten += n
	return n, err
}

// LoggingMiddleware is a middleware for logging the requests.
func LoggingMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			statuscode := 200
			returnSize := "-"
			defer func() {
				now := time.Now()
				// log the request in apache combined log format
				logger.Printf("%s - - [%s] \"%s %s %s\" %d %s \"%s\" \"%s\"\n",
					r.RemoteAddr,
					now.Format("02/Jan/2006:15:04:05 -0700"),
					r.Method,
					r.RequestURI,
					r.Proto,
					statuscode,
					returnSize,
					r.Header.Get("Referer"),
					r.Header.Get("User-Agent"),
				)
			}()
			loggedResponse := NewLoggingResponseWriter(w)

			next.ServeHTTP(loggedResponse, r)
			// get the status code
			statuscode = loggedResponse.statusCode
			// get the response size
			if loggedResponse.bytesWritten > 0 {
				returnSize = fmt.Sprintf("%d", loggedResponse.bytesWritten)
			}
		})
	}
}

// New creates a new instance of the router gorilla/mux router.
func New(
	repositoryContainer model.RepositoryContainer,
	sessionStore *session.Store,
	csvStorage storage.CSVStorage,
	renderer *render.Renderer,
) *mux.Router {
	r := mux.NewRouter()
	// add logger middleware. The logger default flags has to be empty, because the apache log format is used.
	r.Use(LoggingMiddleware(log.New(renderer.GetLogOutput(), "", 0)))
	// handle the static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(renderer.GetStaticDirectoryPath()))))
	routerController := controller.New(
		repositoryContainer,
		sessionStore,
		csvStorage,
		renderer,
	)
	r.HandleFunc("/health", routerController.HealthController)
	r.HandleFunc("/login", routerController.LoginPageController)
	r.HandleFunc("/auth/login", routerController.LoginActionController).Methods("POST")
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(routerController.AuthMiddleware)
	adminRouter.HandleFunc("/dashboard", routerController.DashboardController)

	adminRouter.HandleFunc("/user/create", routerController.UserCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/user/view/{userId}", routerController.UserViewController)
	adminRouter.HandleFunc("/user/update/{userId}", routerController.UserUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/user/delete/{userId}", routerController.UserDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/user/list", routerController.UserListViewController).Methods("GET", "POST")

	adminRouter.HandleFunc("/role/create", routerController.RoleCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/role/view/{roleId}", routerController.RoleViewController)
	adminRouter.HandleFunc("/role/update/{roleId}", routerController.RoleUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/role/delete/{roleId}", routerController.RoleDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/role/list", routerController.RoleListViewController).Methods("GET", "POST")

	adminRouter.HandleFunc("/client/create", routerController.ClientCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/client/view/{clientId}", routerController.ClientViewController)
	adminRouter.HandleFunc("/client/update/{clientId}", routerController.ClientUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/client/delete/{clientId}", routerController.ClientDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/client/list", routerController.ClientListViewController).Methods("GET", "POST")

	adminRouter.HandleFunc("/project/create", routerController.ProjectCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/project/view/{projectId}", routerController.ProjectViewController)
	adminRouter.HandleFunc("/project/update/{projectId}", routerController.ProjectUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/project/delete/{projectId}", routerController.ProjectDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/project/list", routerController.ProjectListViewController).Methods("GET", "POST")

	adminRouter.HandleFunc("/domain/create", routerController.DomainCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/domain/view/{domainId}", routerController.DomainViewController)
	adminRouter.HandleFunc("/domain/update/{domainId}", routerController.DomainUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/domain/delete/{domainId}", routerController.DomainDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/domain/list", routerController.DomainListViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/domain/check-ssl/{domainId}", routerController.DomainCheckSSLViewController).Methods("GET")

	adminRouter.HandleFunc("/environment/create", routerController.EnvironmentCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/environment/view/{environmentId}", routerController.EnvironmentViewController)
	adminRouter.HandleFunc("/environment/update/{environmentId}", routerController.EnvironmentUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/environment/delete/{environmentId}", routerController.EnvironmentDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/environment/list", routerController.EnvironmentListViewController).Methods("GET", "POST")

	adminRouter.HandleFunc("/runtime/create", routerController.RuntimeCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/runtime/view/{runtimeId}", routerController.RuntimeViewController)
	adminRouter.HandleFunc("/runtime/update/{runtimeId}", routerController.RuntimeUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/runtime/delete/{runtimeId}", routerController.RuntimeDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/runtime/list", routerController.RuntimeListViewController).Methods("GET", "POST")

	adminRouter.HandleFunc("/pool/create", routerController.PoolCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/pool/view/{poolId}", routerController.PoolViewController)
	adminRouter.HandleFunc("/pool/update/{poolId}", routerController.PoolUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/pool/delete/{poolId}", routerController.PoolDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/pool/list", routerController.PoolListViewController).Methods("GET", "POST")

	adminRouter.HandleFunc("/database/create", routerController.DatabaseCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/database/view/{databaseId}", routerController.DatabaseViewController)
	adminRouter.HandleFunc("/database/update/{databaseId}", routerController.DatabaseUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/database/delete/{databaseId}", routerController.DatabaseDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/database/list", routerController.DatabaseListViewController).Methods("GET", "POST")

	adminRouter.HandleFunc("/server/create", routerController.ServerCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/server/view/{serverId}", routerController.ServerViewController)
	adminRouter.HandleFunc("/server/update/{serverId}", routerController.ServerUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/server/delete/{serverId}", routerController.ServerDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/server/list", routerController.ServerListViewController).Methods("GET", "POST")

	adminRouter.HandleFunc("/framework/create", routerController.FrameworkCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/framework/view/{frameworkId}", routerController.FrameworkViewController)
	adminRouter.HandleFunc("/framework/update/{frameworkId}", routerController.FrameworkUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/framework/delete/{frameworkId}", routerController.FrameworkDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/framework/list", routerController.FrameworkListViewController).Methods("GET", "POST")

	adminRouter.HandleFunc("/application/create", routerController.ApplicationCreateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/application/view/{applicationId}", routerController.ApplicationViewController)
	adminRouter.HandleFunc("/application/update/{applicationId}", routerController.ApplicationUpdateViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/application/delete/{applicationId}", routerController.ApplicationDeleteViewController).Methods("POST")
	adminRouter.HandleFunc("/application/list", routerController.ApplicationListViewController).Methods("GET", "POST")
	adminRouter.HandleFunc("/application/import-to-environment/{environmentId}", routerController.ApplicationImportToEnvironmentFormController).Methods("GET", "POST")
	adminRouter.HandleFunc("/application/mapping-to-environment/{environmentId}/{fileId}", routerController.ApplicationMappingToEnvironmentFormController).Methods("GET", "POST")

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(routerController.AuthMiddleware)
	apiRouter.HandleFunc("/user/create", routerController.UserCreateAPIController).Methods("POST")
	apiRouter.HandleFunc("/user/view/{userId}", routerController.UserViewAPIController)
	apiRouter.HandleFunc("/user/update/{userId}", routerController.UserUpdateAPIController).Methods("POST")
	apiRouter.HandleFunc("/user/delete/{userId}", routerController.UserDeleteAPIController).Methods("DELETE")
	apiRouter.HandleFunc("/user/list", routerController.UserListAPIController)

	// Re-define the default NotFound handler, so that the logger middleware can log the 404 status code.
	r.NotFoundHandler = r.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	routerController.CacheTemplates()
	return r
}
