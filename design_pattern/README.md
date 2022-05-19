# design_pattern
> 设计模式：https://github.com/sevenelevenlee/go-patterns

> 一切设计模式都是灵活应用struct的组合模式，以及go隐形继承接口的特性
go中的interface就是一些方法装饰, 而struct并不依赖于接口

## 创建类
----

### [建造者模式(Builder Pattern)](./builder_pattern)
> 将一个复杂对象的构建与它的表示分离, 使得同样的构建过程可以创建不同的表示
> 
> 当一个类的构造函数参数个数超过4个，而且这些参数有些是可选的参数，考虑使用构造者模式。

#### [工厂方法模式(Factory Method)](./factory_method_pattern)
> 使一个类的实例化延迟到其子类, 定义一个用于创建对象的接口, 让子类决定将哪一个类实例化

#### [对象池模式(Object Pool)](./object_pool_pattern)
> 根据需求将预测的对象保存到channel中， 用于对象的生成成本大于维持成本

#### [单例模式(Singleton)](./singleton_pattern)
> 单例模式是最简单的设计模式之一, 保证一个类仅有一个实例, 并提供一个全局的访问接口
