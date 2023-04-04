namespace go github.com.weblogs.blog

include "./base.thrift"
include "./model.thrift"

struct CreatePostRequest {
    1: required string Title (vt.min_size = "1")
    2: required i64 AuthorID (vt.gt = "0")
    3: required string Content (vt.min_size = "0", vt.max_size="10485760") // 博客内容最多包含 10MB
    4: optional string Category
    5: optional list<string> Tags
}
struct CreatePostResponse {
    1: required i64 BlogID
    2: required model.BlogStatus Status
    255: base.BaseResp BaseResp
}

struct UpdatePostRequest {
    1: required i64 BlogID
    2: required i64 AuthorID
    3: optional string Title
    4: optional string Content
    5: optional string Category
    6: optional list<string> Tags
}

struct UpdatePostResponse {
    255: base.BaseResp BaseResp
}

struct DeletePostRequest {
    1: required i64 BlogID
}

struct DeletePostResponse {
    255: base.BaseResp BaseResp
}

service BlogPostService {
    CreatePostResponse CreatePost(1: CreatePostRequest req)
    UpdatePostResponse UpdatePost(1: UpdatePostRequest req)
    DeletePostResponse DeletePost(1: DeletePostRequest req)
    //    GetBlogResponse GetBlog(1: GetBlogRequest)
    //    MGetBlogResponse MGetBlog(1: MGetBlogRequest)
    //    SearchBlogResponse SearchBlog(1: SearchBlogRequest)
}