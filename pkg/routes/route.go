package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/26thavenue/edutech_go/pkg/handlers"
)

func AuthRoutes() chi.Router{
    ar := chi.NewRouter()

    ar.Post("/register", handlers.RegisterUser)
    ar.Post("/login", handlers.Login)
    ar.Get("/reset-password", hndlers.ResetPassword)
}

func UserRoutes() chi.Router {
    r := chi.NewRouter()
    r.Get("/users", handlers.CreateUser)
    r.Post("/users", handlers.ListUsers)
    r.Get("/users/{id}", handlers.GetOneUser)
    r.Put("/users/{id}", handlers.UpdateUserDetails)
    r.Delete("/users/{id}", handlers.DeleteUser)
    return r
}

func CourseRoutes() chi.Router {
    cr := chi.NewRouter()

    cr.Post("/course", handlers.CreateCourse)
    cr.Get("/course", handlers.ListCourses)
    cr.Get("/course", handlers.ListFreeCourses)
    cr.Get("/ourse", handlers.ListPaidCourses)
    cr.Get("/course/{id}", handlers.GetACourse)
    cr.Put("/course/{id}", handlers.UpdateCourseDetails)
    cr.Delete("/course/{id}", handlers.DeleteCourse)
    return cr
}

func EnrollmentRoutes() chi.Router{
    er := chi.NewRouter()

    er.Get("/{userId}", handlers.GetUserCourses)
    er.Post("/{userId}/{courseId}", handlers.RegisterCourse)
    er.Post("/{userId}/{courseId}", handlers.UnEnrollCourse)

    return er
}

func LessonRoutes() chi.Router {
    lr := chi.NewRouter()

    lr.Post("/{courseId}/lesson", handlers.CreateLesson)
    lr.Get("/{courseId}/lesson", handlers.ListLessons)
    lr.Get("/{courseId}/{Id}", handlers.Get One Le)
    lr.Get("/{id}", handlers.GetACourse)
    lr.Put("/{id}", handlers.UpdateLessonDetails)
    lr.Delete("/{id}", handlers.DeleteLesson)


    return lr
}

func NotificationRoutes() chi.Router{

}

func AssessmentRoutes() chi.Router{

}

func ProgressRoutes() chi.Router{
    
}