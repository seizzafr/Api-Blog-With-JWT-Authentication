
package routes

import (
    "github.com/gin-gonic/gin"
     "api/controllers"
     auth "api/controllers/auth"
     "api/middleware"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {
        category := api.Group("category",middleware.AuthMiddleware())
        {
            category.GET("/", controllers.GetCategory)        
            category.GET("/:id", controllers.GetCategoryByID) 
            category.POST("/", controllers.CreateCategory)    
            category.PUT("/:id", controllers.UpdateCategory)  
            category.DELETE("/:id", controllers.DeleteCategory) 
        }
         tags := api.Group("tags",middleware.AuthMiddleware())
        {
            tags.GET("/", controllers.GetTags)       
            tags.GET("/:id", controllers.GetTagsByID) 
            tags.POST("/", controllers.CreateTags)     
            tags.PUT("/:id", controllers.UpdateTags)  
            tags.DELETE("/:id", controllers.DeleteTags) 
        }
         posts := api.Group("posts",middleware.AuthMiddleware())
        {
            posts.GET("/", controllers.GetPosts)       
            posts.GET("/:id", controllers.GetPostsByID) 
            posts.POST("/", controllers.CreatePosts)    
            posts.PUT("/:id", controllers.UpdatePosts) 
            posts.DELETE("/:id", controllers.DeletePosts) 
        }
         users := api.Group("users",middleware.AuthMiddleware())
        {
            users.GET("/", controllers.GetUsers)        
            users.GET("/:id", controllers.GetUsersByID) 
            users.PUT("/:id", controllers.UpdateUsers)  
            users.DELETE("/:id", controllers.DeleteUsers) 
        }

         
        PostTag := api.Group("post-tag",middleware.AuthMiddleware())
        {
            PostTag.GET("/", controllers.GetPostTags)       
            PostTag.GET("/:post_id/:tag_id", controllers.GetPostTagByID) 
            PostTag.POST("/", controllers.CreatePostTag)    
            // PostTag.DELETE("/:post_id/:tag_id", controllers.DeletePostTag) // DELETE /api/albums/:id
        }
    }

       api.POST("/login",auth.LoginUser)
       api.POST("/register", controllers.CreateUsers)    
       api.POST("/logout", middleware.AuthMiddleware(),auth.LogoutUser)

    return r
}
