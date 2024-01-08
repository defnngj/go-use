
https://github.com/golang/go/wiki/CodeReviewComments#interfaces


Go interfaces generally belong in the package that uses values of the interface type, not the package that implements those values. The implementing package should return concrete (usually pointer or struct) types: that way, new methods can be added to implementations without requiring extensive refactoring.

Go 接口通常属于使用接口类型值的包，而不是实现这些值的包。实现包应返回具体（通常是指针或结构）类型：这样，可以将新方法添加到实现中，而无需进行大量重构。

Do not define interfaces on the implementor side of an API "for mocking"; instead, design the API so that it can be tested using the public API of the real implementation.

不要在 API 的实现器端定义“用于模拟”的接口;相反，请设计 API，以便可以使用实际实现的公共 API 对其进行测试。

Do not define interfaces before they are used: without a realistic example of usage, it is too difficult to see whether an interface is even necessary, let alone what methods it ought to contain.

在使用接口之前不要定义接口：如果没有实际的使用示例，就很难看到接口是否必要，更不用说它应该包含什么方法了。



