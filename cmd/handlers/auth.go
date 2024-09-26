package handlers

// import (
//     "net/http"
//     "github.com/labstack/echo/v4"
// )

// // Login handles user login and returns JWT token
// func Login(c echo.Context) error {
//     //username := c.FormValue("username")
//     //password := c.FormValue("password")

//     // در اینجا باید اعتبارسنجی کاربر انجام شود
//     userID := 1  // فرض کنید این ID کاربر است

//     // تولید توکن JWT
//     token, err := "x"//services.GenerateJWT(uint(userID))
//     if err != nil {
//         return c.JSON(http.StatusInternalServerError, map[string]string{
//             "message": "could not generate token",
//         })
//     }

//     return c.JSON(http.StatusOK, map[string]string{
//         "token": token,
//     })
// }
