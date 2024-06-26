# Wrapped Logger
Пакет поставляет адаптеры к различным библиотекам логирования.
Это позволяет приложению не завязываться на какую-то определенную библиотеку, а при необходимости, безболезненно заменять одну на другую.

Реализуется это за счет того, что ваше приложение завязывается на данный враппер, а не конкретную библиотеку.
Враппер в свою очередь имеет стандартизированный интерфейс логирования, который адаптируется под конкретную библиотеку.

Когда возникнет потребность попробовать какую-то новую библиотеку логирования -- можно воспользоваться готовым адаптером или написать свой, что намного проще, чем переписывать все приложение.

# Документация
- [Russian](README.RU.MD)
- [English](README.MD)

# Имеющиеся адаптеры
- STD Log (https://go.dev/pkg/log)
- Zap (https://github.com/uber-go/zap)
- Nop (для использования в тестах)

# Возможности wdlogger
- zap-адаптер умеет выводить логи в формате **GELF** (в будущем такая возможность появится и у других)

# Статус проекта
Сейчас находится в стадии экспериментальной разработки и используется в моих собственных проектах для удобства, но это не мешает вам при желании предлагать собственные изменения.