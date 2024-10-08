package constants

// route groups
const (
	AuthGroup   = "/auth"
	UserGroup   = "/users"
	PostGroup   = "/posts"
	UploadGroup = "/uploads"
	CartGroup   = "/cart"
)

// auth routes
const (
	RegisterRoute = "/register"
	LoginRoute    = "/login"
	LogoutRoute   = "/logout"
)

// user routes
const (
	GetAllUsersRoute = "/"
	GetUserByIdRoute = "/:userId"
	GetAuthUserRoute = "/authuser"
	CreateUserRoute  = "/"
	UpdateUserRoute  = "/:userId"
	DeleteUserRoute  = "/:userId"
)

// post routes
const (
	GetAllPostsRoute    = "/"
	GetPostsByUserRoute = "/userposts"
	GetPostByIdRoute    = "/:postId"
	CreatePostRoute     = "/"
	UpdatePostRoute     = "/:postId"
	DeletePostRoute     = "/:postId"
)

// upload routes
const (
	UploadFileRoute = "/"
)

// external API routes
const (
	GetCartByUserIdRoute    = "/cartbyuser"
	AddCartByUserIdRoute    = "/addusercart"
	UpdateCartByUserIdRoute = "/updateusercart"
	DeleteCartByIdRoute     = "/deletecart"
)
