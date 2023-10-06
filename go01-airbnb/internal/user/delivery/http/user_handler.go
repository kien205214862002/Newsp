package userhttp

import (
	"context"
	usermodel "go01-airbnb/internal/user/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	Register(context.Context, *usermodel.UserRegister) error
	Login(context.Context, *usermodel.UserLogin) (*utils.Token, error)
}

type userHandler struct {
	userUC UserUsecase
}

func NewUserHandler(userUC UserUsecase) *userHandler {
	return &userHandler{userUC}
}

func (hdl *userHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserRegister

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := hdl.userUC.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{"data": data.Id})
	}
}
Đoạn mã này là một phần của một ứng dụng sử dụng framework Gin để xử lý HTTP requests và có nhiệm vụ đăng ký người dùng mới bằng cách xử lý dữ liệu đăng ký được gửi trong request và trả về ID của người dùng vừa đăng ký dưới dạng JSON. Hãy cùng tìm hiểu ý nghĩa của từng phần trong mã:

1. `func (hdl *userHandler) Register() gin.HandlerFunc`: Đây là một phương thức của struct `userHandler`. Phương thức này trả về một hàm xử lý HTTP request thuộc kiểu `gin.HandlerFunc`.

2. `return func(c *gin.Context) { ... }`: Phương thức `Register` trả về một hàm xử lý HTTP request, và đây là định nghĩa của hàm đó. Hàm này nhận một tham số là `c`, là một đối tượng `gin.Context` được truyền từ framework Gin khi có một HTTP request tới endpoint mà phương thức `Register` này đang xử lý.

3. `var data usermodel.UserRegister`: Đoạn này khai báo một biến `data` kiểu `usermodel.UserRegister`, có vẻ như đây là một cấu trúc dùng để chứa thông tin đăng ký của người dùng.

4. `if err := c.ShouldBind(&data); err != nil { ... }`: Đoạn này cố gắng ràng buộc dữ liệu từ request body vào biến `data`. Nếu quá trình ràng buộc gặp lỗi (ví dụ: dữ liệu không hợp lệ), thì chương trình sẽ gọi `panic` với lỗi BadRequest được tạo bởi `common.ErrBadRequest(err)`.

5. `if err := hdl.userUC.Register(c.Request.Context(), &data); err != nil { ... }`: Đoạn này gọi một phương thức `Register` từ đối tượng `hdl.userUC` để đăng ký người dùng mới. Tham số truyền vào bao gồm ngữ cảnh của request và thông tin đăng ký người dùng (`&data`). Nếu quá trình đăng ký gặp lỗi, chương trình sẽ gọi `panic` với lỗi được trả về.

6. `c.JSON(200, gin.H{"data": data.Id})`: Cuối cùng, nếu mọi thứ diễn ra suôn sẻ (không có lỗi), phản hồi HTTP sẽ được gửi về với mã trạng thái 200 OK và một JSON response chứa ID của người dùng vừa đăng ký (`data.Id`).

func (hdl *userHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials usermodel.UserLogin

		if err := c.ShouldBind(&credentials); err != nil {
			panic(common.ErrBadRequest(err))
		}

		token, err := hdl.userUC.Login(c.Request.Context(), &credentials)
		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{"data": token})
	}
}
