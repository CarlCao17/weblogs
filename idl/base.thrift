namespace go github.com.weblogs.base

struct BaseResp {
    1: string StatusMessage
    2: i32 StatusCode
    3: optional map<string, string> Extra
}