package zifare


/**
	路由的抽象接口
 */
type IRouter interface {
	//在处理conn业务之前
	PreHandle(request IRequest)

	//处理业务的方法
	Handle(request IRequest)

	//处理conn业务之后的方法
	PostHandle(request IRequest)
}
