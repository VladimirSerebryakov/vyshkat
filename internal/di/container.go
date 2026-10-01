package di

import "reflect"

// Container - минимальный DI-контейнер
// factories - фабрики по типу, instances - кэш уже построенных объектов синглетонов
type Container struct {
	factories map[reflect.Type]func(*Container) any
	instances map[reflect.Type]any
}

// Пустой контейнер
func New() *Container {
	return &Container{
		factories: make(map[reflect.Type]func(*Container) any),
		instances: make(map[reflect.Type]any),
	}
}

// typeOf возвращает reflect.Type для T
func typeOf[T any]() reflect.Type {
	var zero T
	return reflect.TypeOf(&zero).Elem()
}

// Register регистрирует фабрику, строющую объект типа T при первом обращении
func Register[T any](c *Container, factory func(*Container) T) {
	t := typeOf[T]()
	c.factories[t] = func(c *Container) any {
		return factory(c)
	}
}

// RegisterInstance регистрирует значение вместо фабрики, для тестов
func RegisterInstance[T any](c *Container, instance T) {
	t := typeOf[T]()
	c.instances[t] = instance
}

// MustResolve возвращает объект типа T, сперва проверяет кэш (instances) иначее
// строит через фабрику и кэширует результат. Паникует, если для T ничего не зарегистрировано
func MustResolve[T any](c *Container) T {
	t := typeOf[T]()
	if inst, ok := c.instances[t]; ok {
		return inst.(T)
	}

	factory, ok := c.factories[t]
	if !ok {
		panic("di: no factory registered for type " + t.String())
	}
	instance := factory(c)
	c.instances[t] = instance
	return instance.(T)
}
