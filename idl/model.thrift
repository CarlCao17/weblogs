namespace go github.com.weblogs.model

enum BlogStatus {
    NotExist // 不存在
    Created // 已创建（编辑态/草稿态）
    Published // 已发布（不可编辑）
    Deleted // 已删除
}

struct BlogPost {
    1: i64 ID // Blog ID
    2: string Title // Blog title
    3: i64 AuthorID // Blog author
    4: string Content // Blog content
    5: BlogStatus Status // Blog status
    6: string Category // Blog category
    7: list<string> Tags // Blog tags
    8: i64 CreatedTime // Blog create time
    9: i64 UpdatedTime // Blog update time
}

typedef string UserRole
const UserRole ROLE_GUEST  = "guest"  // 游客
const UserRole ROLE_USER  = "user"// 普通用户
const UserRole ROLE_ADMIN = "admin" // 管理员
const UserRole ROLE_ROOT = "root"   // 超级管理员


struct User {
    1: i64 ID // User ID
    2: string UserName // User name
    3: string Password // User password
    4: string Email
    5: string FirstName
    6: string LastName
    7: UserRole Role // User Role
}