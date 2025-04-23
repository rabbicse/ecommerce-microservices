package handlers

import (
	"ecommerce-service/internal/api/rest"
	"ecommerce-service/internal/dto"
	"ecommerce-service/internal/repository"
	"ecommerce-service/internal/service"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	// svc UserService
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an istance of user service & inject to handler
	svc := service.UserService{
		Repository: repository.NewUserRepository(rh.DB),
		Auth:       rh.Auth,
	}
	handler := UserHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/users")

	// Public endpoints
	pubRoutes.Post("/register", handler.Register)
	pubRoutes.Post("/login", handler.Login)

	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)

	// Private endpoints
	pvtRoutes.Get("/verify", handler.GetVerificationCode)
	pvtRoutes.Post("/verify", handler.Verify)

	pvtRoutes.Get("/profile", handler.GetProfile)
	pvtRoutes.Post("/profile", handler.CreateProfile)

	pvtRoutes.Get("/cart", handler.GetCart)
	pvtRoutes.Post("/cart", handler.AddToCart)

	pvtRoutes.Post("/order", handler.CreateOrder)
	pvtRoutes.Get("/order", handler.GetOrders)
	pvtRoutes.Get("/order/:id", handler.GetOrder)

	pvtRoutes.Post("/become-seller", handler.BecomeSeller)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	token, err := h.svc.Signup(user)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error on signup",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	user := dto.UserLogin{}
	err := ctx.BodyParser(&user)
	log.Println(err)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide valid inputs",
		})
	}

	token, err := h.svc.Login(user.Email, user.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "error on login",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login",
		"token":   token,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	// create verification code and update to user profile in DB
	code, err := h.svc.GetVerificationCode(user)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to generate verification code",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get verification code",
		"code":    code,
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	user := h.svc.Auth.GetCurrentUser(ctx)

	// request
	var req dto.VerificationCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	err := h.svc.VerifyCode(user.ID, req.Code)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verified successfully",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "profile",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "profile",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "cart",
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "add to cart",
	})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create order",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become seller",
	})
}
