package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thesunonthesky/GoBackendPackage/errs"
	"github.com/thesunonthesky/GoBackendPackage/model"
	"github.com/thesunonthesky/GoBackendPackage/repository"
)

type UserHandler struct {
	repo repository.UserRepository
}

func handleError(c *gin.Context, err error) {
	if appErr, ok := err.(errs.AppError); ok {
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
	}
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUsersByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		// ใช้ BadRequestError ของเราเอง
		handleError(c, errs.NewBadRequestError("Invalid ID format"))
		return
	}

	users, err := h.repo.GetByID(idInt)
	if err != nil {
		// ส่ง err ให้ฟังก์ชันกลางจัดการ (ซึ่งถ้าหาไม่เจอ มันจะไปดึง status 404 ออกมาเอง)
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUser
	if err := c.ShouldBindJSON(&req); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var errMessage string
			for _, vErr := range validationErrs {
				switch vErr.Field() {
				case "Name":
					errMessage += "กรุณากรอกชื่อ | "
				case "Email":
					if vErr.Tag() == "required" {
						errMessage += "กรุณากรอกอีเมล | "
					} else if vErr.Tag() == "email" {
						errMessage += "รูปแบบอีเมลไม่ถูกต้อง | "
					}
				case "Password":
					if vErr.Tag() == "min" {
						errMessage += "รหัสผ่านต้องมีอย่างน้อย 6 ตัวอักษร | "
					} else {
						errMessage += "กรุณากรอกรหัสผ่าน | "
					}
				}
			}
			handleError(c, errs.NewBadRequestError(errMessage))
			return
		}
		
		handleError(c, errs.NewBadRequestError("รูปแบบข้อมูลไม่ถูกต้อง"))
		return
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.repo.Create(user); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User Created!"})

}
