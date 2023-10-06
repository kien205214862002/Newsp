package placehttp

import (
	"context"
	"fmt"
	placemodel "go01-airbnb/internal/place/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlaceUseCase interface {
	CreatePlace(context.Context, *placemodel.Place) error
	GetPlaces(context.Context, *common.Paging, *placemodel.Filter) ([]placemodel.Place, error)
	GetPlaceByID(context.Context, int) (*placemodel.Place, error)
	UpdatePlace(context.Context, common.Requester, int, *placemodel.Place) error
	DeletePlace(context.Context, common.Requester, int) error
}

type placeHandler struct {
	placeUC PlaceUseCase
	hasher  *utils.Hasher
}

func NewPlaceHandler(placeUC PlaceUseCase, hasher *utils.Hasher) *placeHandler {
	return &placeHandler{placeUC, hasher}
}

func (hdl *placeHandler) CreatePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet("user").(common.Requester)

		var place placemodel.Place

		if err := c.ShouldBind(&place); err != nil {
			panic(common.ErrBadRequest(err))
		}

		place.OwnerId = requester.GetUserId()

		if err := hdl.placeUC.CreatePlace(c.Request.Context(), &place); err != nil {
			panic(err)
		}

		// Encode id trước trả ra cho client
		place.FakeId = hdl.hasher.Encode(place.Id, common.DBTypePlace)

		c.JSON(http.StatusOK, common.Response(place))
	}
}
//Đoạn mã này là một phần của một ứng dụng sử dụng framework Gin để xử lý HTTP requests và có nhiệm vụ tạo mới một địa điểm (place) dựa trên thông tin được gửi trong request và trả về dữ liệu của địa điểm đã tạo dưới dạng JSON. Hãy cùng tìm hiểu ý nghĩa của từng phần trong mã:

1. `func (hdl *placeHandler) CreatePlace() gin.HandlerFunc`: Đây là một phương thức của struct `placeHandler`. Phương thức này trả về một hàm xử lý HTTP request thuộc kiểu `gin.HandlerFunc`.

2. `return func(c *gin.Context) { ... }`: Phương thức `CreatePlace` trả về một hàm xử lý HTTP request, và đây là định nghĩa của hàm đó. Hàm này nhận một tham số là `c`, là một đối tượng `gin.Context` được truyền từ framework Gin khi có một HTTP request tới endpoint mà phương thức `CreatePlace` này đang xử lý.

3. `requester := c.MustGet("user").(common.Requester)`: Đoạn này trích xuất thông tin về người dùng từ context. Giả sử thông tin người dùng được lưu trữ trong context dưới tên "user", và đối tượng `common.Requester` được sử dụng để đại diện cho người dùng.

4. `var place placemodel.Place`: Đoạn này khai báo một biến `place` kiểu `placemodel.Place`, có vẻ như đây là một đối tượng đại diện cho thông tin của một địa điểm.

5. `if err := c.ShouldBind(&place); err != nil { ... }`: Đoạn này cố gắng ràng buộc dữ liệu từ request body vào biến `place`. Nếu quá trình ràng buộc gặp lỗi (ví dụ: dữ liệu không hợp lệ), thì chương trình sẽ gọi `panic` với lỗi BadRequest được tạo bởi `common.ErrBadRequest(err)`.

6. `place.OwnerId = requester.GetUserId()`: Đoạn này gán `OwnerId` của địa điểm bằng ID của người dùng hiện tại. Giả sử `requester.GetUserId()` trả về ID của người dùng đang thực hiện request.

7. `if err := hdl.placeUC.CreatePlace(c.Request.Context(), &place); err != nil { ... }`: Đoạn này gọi một phương thức `CreatePlace` từ đối tượng `hdl.placeUC` để tạo mới một địa điểm. Tham số truyền vào bao gồm ngữ cảnh của request và địa điểm cần tạo (`&place`). Nếu quá trình tạo địa điểm gặp lỗi, chương trình sẽ gọi `panic` với lỗi được trả về.

8. `place.FakeId = hdl.hasher.Encode(place.Id, common.DBTypePlace)`: Đoạn này thực hiện việc tạo một giá trị `FakeId` bằng cách mã hóa `Id` của địa điểm bằng cách sử dụng đối tượng `hdl.hasher`. Điều này có thể liên quan đến việc tạo một "giả" ID cho địa điểm.

9. `c.JSON(http.StatusOK, common.Response(place))`: Cuối cùng, nếu mọi thứ diễn ra suôn sẻ (không có lỗi), phản hồi HTTP sẽ được gửi về với mã trạng thái 200 OK và một JSON response chứa thông tin về địa điểm (`place`) đã tạo mới.

func (hdl *placeHandler) GetPlaces() gin.HandlerFunc {
	return func(c *gin.Context) {
		// paging
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err))
		}
		paging.Fulfill()

		// filter
		var filter placemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrBadRequest(err))
		}

		data, err := hdl.placeUC.GetPlaces(c.Request.Context(), &paging, &filter)
		if err != nil {
			panic(err)
		}

		// Encode id trước trả ra cho client
		for i, v := range data {
			if data[i].Owner == nil {
				fmt.Println("ERRORRRRRR")
			}
			data[i].FakeId = hdl.hasher.Encode(v.Id, common.DBTypePlace)
			data[i].Owner.FakeId = hdl.hasher.Encode(v.Owner.Id, common.DBTypeUser)
		}

		c.JSON(http.StatusOK, common.ResponseWithPaging(data, paging))
	}
}

func (hdl *placeHandler) GetPlaceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// id, err := strconv.Atoi(c.Param("id"))
		id, err := hdl.hasher.Decode(c.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		data, err := hdl.placeUC.GetPlaceByID(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		data.FakeId = hdl.hasher.Encode(data.Id, common.DBTypePlace)
		data.Owner.FakeId = hdl.hasher.Encode(data.Owner.Id, common.DBTypeUser)
		c.JSON(http.StatusOK, common.Response(data))
	}
}
//Đoạn mã này là một phần của một ứng dụng sử dụng framework Gin để xử lý HTTP requests và có nhiệm vụ lấy thông tin về một địa điểm dựa trên một ID và trả về dữ liệu đó dưới dạng JSON. Hãy cùng tìm hiểu ý nghĩa của nó:

1. `func (hdl *placeHandler) GetPlaceByID() gin.HandlerFunc`: Đây là một phương thức của struct `placeHandler`. Phương thức này trả về một hàm xử lý HTTP request thuộc kiểu `gin.HandlerFunc`.

2. `return func(c *gin.Context) { ... }`: Phương thức `GetPlaceByID` trả về một hàm xử lý HTTP request, và đây là định nghĩa của hàm đó. Hàm này nhận một tham số là `c`, là một đối tượng `gin.Context` được truyền từ framework Gin khi có một HTTP request tới endpoint mà phương thức `GetPlaceByID` này đang xử lý.

3. `id, err := hdl.hasher.Decode(c.Param("id"))`: Đoạn này thực hiện việc giải mã một tham số từ URL được gửi đến endpoint này (thường là một ID), sử dụng một phương thức `Decode` của đối tượng `hdl.hasher`. Kết quả được lưu vào biến `id`, và nếu có lỗi, `err` sẽ được ghi nhận.

4. `data, err := hdl.placeUC.GetPlaceByID(c.Request.Context(), id)`: Đoạn này gọi một phương thức `GetPlaceByID` từ đối tượng `hdl.placeUC` để lấy thông tin về địa điểm dựa trên `id` đã giải mã. Tham số truyền vào bao gồm ngữ cảnh của request và `id`. Kết quả được lưu vào biến `data`, và nếu quá trình lấy dữ liệu gặp lỗi, `err` sẽ được ghi nhận.

5. `data.FakeId = hdl.hasher.Encode(data.Id, common.DBTypePlace)`: Đoạn này thực hiện việc tạo một giá trị `FakeId` bằng cách mã hóa `Id` của địa điểm bằng cách sử dụng đối tượng `hdl.hasher`. Điều này có thể liên quan đến việc tạo một "giả" ID cho địa điểm.

6. `data.Owner.FakeId = hdl.hasher.Encode(data.Owner.Id, common.DBTypeUser)`: Đoạn này thực hiện tương tự như trên, nhưng lần này là với thông tin của chủ sở hữu (Owner) của địa điểm.

7. `c.JSON(http.StatusOK, common.Response(data))`: Cuối cùng, nếu mọi thứ diễn ra suôn sẻ (không có lỗi), phản hồi HTTP sẽ được gửi về với mã trạng thái 200 OK và một JSON response chứa thông tin về địa điểm (`data`) đã lấy từ cơ sở dữ liệu sau khi được xử lý.
func (hdl *placeHandler) UpdatePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy thông tin Requester
		requester := c.MustGet("user").(common.Requester)

		// id, err := strconv.Atoi(c.Param("id"))
		id, err := hdl.hasher.Decode(c.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		var place placemodel.Place

		if err := c.ShouldBind(&place); err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := hdl.placeUC.UpdatePlace(c.Request.Context(), requester, id, &place); err != nil {
			panic(err)
		}

		place.FakeId = hdl.hasher.Encode(place.Id, common.DBTypePlace)
		c.JSON(http.StatusOK, common.Response(place))
	}
}
//Đoạn mã này là một phần của một ứng dụng sử dụng framework Gin để xử lý HTTP requests và dường như liên quan đến việc cập nhật thông tin về một địa điểm. Hãy cùng tìm hiểu ý nghĩa của nó:

1. `func (hdl *placeHandler) UpdatePlace() gin.HandlerFunc`: Đây là một phương thức của struct `placeHandler`. Phương thức này trả về một hàm xử lý HTTP request thuộc kiểu `gin.HandlerFunc`.

2. `return func(c *gin.Context) { ... }`: Phương thức `UpdatePlace` trả về một hàm xử lý HTTP request, và đây là định nghĩa của hàm đó. Hàm này nhận một tham số là `c`, là một đối tượng `gin.Context` được truyền từ framework Gin khi có một HTTP request tới endpoint mà phương thức `UpdatePlace` này đang xử lý.

3. `requester := c.MustGet("user").(common.Requester)`: Đoạn này trích xuất thông tin về người dùng từ context. Giả sử thông tin người dùng được lưu trữ trong context dưới tên "user", và đối tượng `common.Requester` được sử dụng để đại diện cho người dùng.

4. `id, err := hdl.hasher.Decode(c.Param("id"))`: Đoạn này thực hiện việc giải mã một tham số từ URL được gửi đến endpoint này (thường là một ID), sử dụng một phương thức `Decode` của đối tượng `hdl.hasher`. Kết quả được lưu vào biến `id`, và nếu có lỗi, `err` sẽ được ghi nhận.

5. `var place placemodel.Place`: Đoạn này khai báo một biến `place` kiểu `placemodel.Place`, có vẻ như đây là một đối tượng đại diện cho thông tin của một địa điểm.

6. `if err := c.ShouldBind(&place); err != nil { ... }`: Đoạn này cố gắng ràng buộc dữ liệu từ request body vào biến `place`. Nếu quá trình ràng buộc gặp lỗi (ví dụ: dữ liệu không hợp lệ), thì chương trình sẽ gọi `panic` với lỗi BadRequest được tạo bởi `common.ErrBadRequest(err)`.

7. `if err := hdl.placeUC.UpdatePlace(c.Request.Context(), requester, id, &place); err != nil { ... }`: Đoạn này gọi một phương thức `UpdatePlace` từ đối tượng `hdl.placeUC` để cập nhật thông tin của địa điểm. Các tham số truyền vào bao gồm ngữ cảnh của request, thông tin người dùng (`requester`), `id` của địa điểm cần cập nhật, và thông tin cập nhật (`&place`). Nếu quá trình cập nhật gặp lỗi, chương trình sẽ gọi `panic` với lỗi được trả về.

8. `place.FakeId = hdl.hasher.Encode(place.Id, common.DBTypePlace)`: Đoạn này thực hiện việc tạo một giá trị `FakeId` bằng cách mã hóa `Id` của địa điểm bằng cách sử dụng đối tượng `hdl.hasher`. Điều này có thể liên quan đến việc tạo một "giả" ID cho địa điểm.

9. `c.JSON(http.StatusOK, common.Response(place))`: Cuối cùng, nếu mọi thứ diễn ra suôn sẻ (không có lỗi), phản hồi HTTP sẽ được gửi về với mã trạng thái 200 OK và một JSON response chứa thông tin về địa điểm (`place`) sau khi được cập nhật.
func (hdl *placeHandler) DeletePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet("user").(common.Requester)

		// id, err := strconv.Atoi(c.Param("id"))
		id, err := hdl.hasher.Decode(c.Param("id"))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := hdl.placeUC.DeletePlace(c.Request.Context(), requester, id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.Response(true))
	}
}
//Đoạn mã trên là một phần của một ứng dụng sử dụng framework Gin để xử lý HTTP requests. Hãy cùng tìm hiểu ý nghĩa của nó:

1. `func (hdl *placeHandler) DeletePlace() gin.HandlerFunc`: Đây là một phương thức của một struct có tên là `placeHandler`. Phương thức này trả về một hàm xử lý HTTP request thuộc kiểu `gin.HandlerFunc`.

2. `return func(c *gin.Context) { ... }`: Phương thức `DeletePlace` trả về một hàm xử lý HTTP request, và đây là định nghĩa của hàm đó. Hàm này nhận một tham số là `c`, là một đối tượng `gin.Context` được truyền từ framework Gin khi có một HTTP request tới endpoint mà phương thức `DeletePlace` này đang xử lý.

3. `requester := c.MustGet("user").(common.Requester)`: Đoạn này trích xuất thông tin về người dùng từ context. Giả sử rằng thông tin người dùng được lưu trữ trong context dưới tên "user", và đối tượng `common.Requester` được sử dụng để đại diện cho người dùng.

4. `id, err := hdl.hasher.Decode(c.Param("id"))`: Đoạn này thực hiện việc giải mã một tham số từ URL được gửi đến endpoint này (thường là một ID), sử dụng một phương thức `Decode` của đối tượng `hdl.hasher`. Kết quả được lưu vào biến `id`, và nếu có lỗi, `err` sẽ được ghi nhận.

5. `if err != nil { ... }`: Đây là một điều kiện kiểm tra lỗi. Nếu việc giải mã `id` thất bại (tức là `err` khác nil), thì chương trình sẽ gọi `panic` với một lỗi BadRequest được tạo bởi `common.ErrBadRequest(err)`.

6. `if err := hdl.placeUC.DeletePlace(c.Request.Context(), requester, id); err != nil { ... }`: Đoạn này gọi một phương thức `DeletePlace` từ đối tượng `hdl.placeUC` (có vẻ như đây là một đối tượng thực hiện logic kinh doanh liên quan đến địa điểm) và truyền vào ngữ cảnh của request, thông tin về người dùng (`requester`), và `id`. Nếu việc xóa địa điểm không thành công, chương trình sẽ gọi `panic` với lỗi được trả về.

7. `c.JSON(http.StatusOK, common.Response(true))`: Cuối cùng, nếu mọi thứ diễn ra suôn sẻ (không có lỗi), phản hồi HTTP sẽ được gửi về với mã trạng thái 200 OK và một JSON response chứa `true`.

Tóm lại, đoạn mã này đảm bảo xóa một địa điểm dựa trên yêu cầu HTTP được gửi đến và xử lý lỗi nếu có.
