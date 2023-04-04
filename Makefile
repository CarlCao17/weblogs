
MODULE = github.com/weblogs

.PHONY: start
start:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose down

.PHONY: blog
blog:
	go1.20.1 run app/blog/*.go

.PHONY: update_blog_idl
update_blog_idl:
	kitex --thrift-plugin validator -service ${MODULE} ./idl/blog_post.thrift

.PHONY: update_user_idl
update_user_idl:
	kitex --thrift-plugin validator -service ${MODULE} ./idl/user.thrift