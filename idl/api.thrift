namespace go github.com.weblogs

include "./base.thrift"
include "./model.thrift"

struct CreateBlogRequest {
    1: required string Title
    2: required i64 AuthorID
    3: required string Content
    4: optional string Category
    5: optional list<string> Tags
}
struct CreateBlogResponse {
    1: required i64 BlogID
    2: required model.Status Status
    255: base.BaseResp BaseResp
}

struct UpdateBlogRequest {
    1: required i64 BlogID
    2: string Title
    3: string Content
    4: string Category
    5: list<string> Tags
}

struct UpdateBlogResponse {
    255: base.BaseResp BaseResp
}

struct DeleteBlogRequest {
    1: required i64 BlogID
}

struct DeleteBlogResponse {
    255: base.BaseResp BaseResp
}

service BlogService {
    CreateBlogResponse CreateBlog(1: CreateBlogRequest)
    UpdateBlogResponse UpdateBlog(1: UpdateBlogRequest)
    DeleteBlogResponse DeleteBlog(1: DeleteBlogRequest)
    GetBlogResponse GetBlog(1: GetBlogRequest)
    MGetBlogResponse MGetBlog(1: MGetBlogRequest)
    SearchBlogResponse SearchBlog(1: SearchBlogRequest)
}