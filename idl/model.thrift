namespace go model

enum Status {
    Created // 已创建（编辑态/草稿态）
    Published // 已发布（不可编辑）
    Deleted // 已删除
}

struct Blog {
    1: i64 ID // Blog ID
    2: string Title // Blog title
    3: i64 AuthorID // Blog author
    4: string Content // Blog content
    5: Status Status // Blog status
    6: string Category // Blog category
    7: list<string> Tags // Blog tags
    8: i64 CreatedTime // Blog create time
    9: i64 UpdatedTime // Blog update time
}

enum Role {
    Guest  // 游客
    Normal // 普通用户
    Admin  // 管理员
    Root   // 超级管理员
}

struct User {
    1：i64 ID // User ID
    2: string Name // User name
    3: Role Role // User Role
}