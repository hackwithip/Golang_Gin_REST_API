package configs

import "fmt"

// Define version as a variable
const version = "v1"

// Define routes using the version variable
var (
	SignupRoute = fmt.Sprintf("/api/%s/signup", version)
	LoginRoute  = fmt.Sprintf("/api/%s/login", version)
	VerifyToken = fmt.Sprintf("/api/%s/verify-token", version)
	LogoutRoute = fmt.Sprintf("/api/%s/logout", version)
	
	GetCategoriesRoute = fmt.Sprintf("/api/%s/categories", version)
	CreateCategoriesRoute = fmt.Sprintf("/api/%s/categories", version)
	DeleteCategoryRoute = fmt.Sprintf("/api/%s/categories/:id", version)

	GetAuthorsRoute = fmt.Sprintf("/api/%s/authors", version)
	CreateAuthorRoute = fmt.Sprintf("/api/%s/authors", version)
	DeleteAuthorsRoute = fmt.Sprintf("/api/%s/authors/:id", version)

	GetBlogsRoute = fmt.Sprintf("/api/%s/blogs", version)
	CreateBlogsRoute = fmt.Sprintf("/api/%s/blogs", version)
	UpdateBlogRoute = fmt.Sprintf("/api/%s/blogs/:id", version)
	DeleteBlogRoute = fmt.Sprintf("/api/%s/blogs/:id", version)
)
