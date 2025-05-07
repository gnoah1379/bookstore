# BookStore

### 26/4

=======TODO=======

--tìm hiểu cấu trúc thư mục của dự án
├───cmd             --chứa các thư mục running server
├───configs         --chứa files config của project
├───internal
│   ├───config      --chứa files định nghĩa các config có trong project
│   ├───database    --chứa files connect database
│   ├───handlers    --chứa files xử lý request và reponse  
│   ├───model       --chứa files định nghĩa các model sử dụng trong project
│   ├───repository  --chứa fiels xử lý dữ liệu từ db
│   └───service     --chứa files cung cấp các service
├───migrations      --migration database
└───main.go         --files running project


--tìm hiểu cách hoạt dộng của api register
clients --> handler --> service --> repo --> database 
    ^--------handler-----service-----repo--------|

clients gửi post request đến server 
--> srv.router.POST("/api/v1/user/register", srv.user.Register) --handler/server.go

server.go gọi đến hàm user.Register 
--> func (h *UserHandler) Register(c *gin.Context)              --handler/user.go

user handler gửi request đến service.UserRegistration 
--> func (u *userService) Register(ctx context.Context, userReg UserRegistration) (model.UserInfo, error)      --service/user.go

service gọi đến hàm createUser từ repo 
--> userInfo, err := h.userSvc.Register(c.Request.Context(), request) 
--> func (r *userRepo) CreateUser(ctx context.Context, user *model.User) error 

repo xử lý request đến database 
--> return r.db.WithContext(ctx).Create(user).Error


--tìm hiểu tác dụng và cách hoạt động của từng file
cmd/root.go     (sử dụng cobra)
--> tạo biến rootCmd
--> tạo hàm init() để add serverCmd vào rootCmd
--> tạo hàm Execute() rootCmd.Execute()

cmd/server.go   (sử dụng cobra)
--> tạo biến configPath nhận vào path file config.json từ hàm init()
--> tạo biến serverCmd --> gọi LoadConfig(configPath)
                       --> gọi database.Open(cfg.DB)
                       --> gọi repo, service, handler
                       --> dùng biến srv gọi đến NewServer()
                                --> nếu timeOut chạy hàm Shutdown()
                                --> thành công srv.Start()

configs/config.json
--> viết config cho project

config/config.go
--> định nghĩa các struct liên quan đến config (config(db, server), database, server)
--> tạo hàm LoadConfig() để đọc config từ file config.json và trả về cfg (sử dụng viper)

database.go
--> tạo hàm Open() nhận vào tham số cfg, tạo kết nối db theo config và trả về db (sử dụng gorm)

handlers/reponse.go
--> định nghĩa struct error và success 
--> tạo hàm ResponseError() và ResponseSuccess() để trả về log status (sử dụng gin)

handlers/server.go
--> định nghĩa struct Server
--> tạo hàm NewServer() nhận vào config và cấu hình cho server
--> gọi hàm register cho srv và trả về srv
--> tạo hàm Start(), Shutdown()
--> tạo hàm Register() gửi Post request đến userHandler

handlers/user.go
--> tạo struct UserHandler gọi đến UserService
--> tạo hàm Register() nhận vào request từ api và gửi đến service
--> nếu thành công ReponseSuccess

repository/user.go
--> định nghĩa interface UserRepo{create, getUser}
--> kết nối với db thực hiện hàm create và getUser (sử dụng gorm)

service/password.go 
--> sử dụng thư viện bcypt định nghĩa hai hàm HashPassword() và CheckPasswordHash()
--> trả về password đã mã hóa để đưa vào db và mã hóa password được nhập vào trước khi so sánh với password trong db

service/user.go
--> định nghĩa struct UserRegister cho việc đăng ký
--> định nghĩa interface UserService{register}
--> tạo hàm register nhận vào dữ liệu đăng ký 
--> mã hóa mật khẩu nhập vào, đưa dữ liệu đăng ký vào biến user
--> gọi hàm create từ repo với tham số là user
--> trả về lỗi nếu create thất bại, trả về userInfo nếu thành công

main.go
--> chạy hàm cmd.Execute()


tạo api cho chức năng đăng nhập

### 27/4
tìm hiểu về jwt và ứng dụng jwt vào chức năng đăng nhập
--> jwt(JSONwebtoken) là một phương thức để xác thực giữa server và clients dưới dạng đối tượng JSON
--> 1 jwt được chia làm 3 phần --> Headder{"alg": "HS256", "typ": "JWT"} thuật toán mã hóa, kiểu token
                               --> Payload{name, ...} chứa các clamims
                               --> Signature tạo bằng cách mã hóa headder, payload và một secret key

### 29/4
--> tạo CRUD api cho book
--> tìm hiểu cách sử dụng gorm và đặt tên endpoint cho api

### 30/4
--> test và hoàn thiện chức năng cho api book
--> test và hoàn thiện cho api user

### 1/5
--> tạo CRUD api cho order
--> test và hoàn thiện chức năng cho api order
--> sửa lại lỗi cập nhật user và book
--> thêm chức năng tìm kiếm theo tên cho book
--> sửa lại lỗi xác thực JWT

### 5/5
tạo api cho payment
hoàn thiện lại orderService

### 6/5
tạo api cho chức năng đánh giá sách
-db
review(review_id, book_id, rating, comment, user_comment, count_reply, created/updated_at)
reply-review(review_id, comment, user_comment, created/updated_at)
-service
    get all review    //require admin
    get all reply    //require admin

    create review    //require auth
    get review by book id
    update review    //require auth
    delete review    //require auth

    create reply    //require auth
    get reply review by review id
    update reply    //require auth
    delete reply    //require auth

### 7/5
hoàn thiện và kiểm thử api đánh giá sách
update check role và user id cho các api yêu cầu update và delete
thêm giảm số lượng sách khi thanh toán thành công
cập nhật get all order by user id cho api order
