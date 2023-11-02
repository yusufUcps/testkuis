package routes

import (
	"quiz/configs"
	"quiz/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc controller.UserControllInterface, cfg configs.ProgramConfig) {
	e.POST("/user", uc.Register())
	e.POST("/login", uc.Login())
	e.GET("/my-profile", uc.MyProfile(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/my-profile", uc.UpdateMyProfile(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteQuiz(e *echo.Echo, uq controller.QuizControllInterface, cfg configs.ProgramConfig) {
	e.POST("/quiz", uq.InsertQuiz(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/quiz", uq.GetAllQuiz(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/quiz/:id", uq.GetQuizByID(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/quiz/:id", uq.UpdateQuiz(), echojwt.JWT([]byte(cfg.Secret)))
	e.DELETE("/quiz/:id", uq.DeleteQuiz(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/my-quiz", uq.GetAllMyQuiz(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteQuestion(e *echo.Echo, uq controller.QuestionsControllInterface, cfg configs.ProgramConfig) { 
	e.POST("/question/:id", uq.InsertQuestion(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/questions", uq.GetAllQuestionsQuiz(), echojwt.JWT([]byte(cfg.Secret)))
	e.GET("/question/:id", uq.GetQuetionByID(), echojwt.JWT([]byte(cfg.Secret)))
	e.PUT("/question/:id", uq.UpdateQuestion(), echojwt.JWT([]byte(cfg.Secret)))
	e.POST("/generate", uq.GenerateQuestion(), echojwt.JWT([]byte(cfg.Secret)))
}