// главное назначение - перетаскивание из одного микрика в другой констант

// В этом файле описаны общие константы для всего Observer
// Они все равно будут перетекать из пакета в пакет,
// А выделять отдельный пакет observer, в котором забить их
// приведет к дикой мешанине логирования/трейса/метрик/grpc/broker в рамках одного пакета

package consts

// Приписка obs. - означает observer
const (
	// Как GRPC так и DBus
	RequestID     = "obs.request_id"
	RequestMethod = "obs.method"
)

// ЭТО КОСТЫЛЬ НА ОБЩИЕ КОНСТАНТЫ ПО КОНТЕКСТУ ВСЕГО OBSERVER
var Consts = []string{
	RequestID,
	RequestMethod,
}
