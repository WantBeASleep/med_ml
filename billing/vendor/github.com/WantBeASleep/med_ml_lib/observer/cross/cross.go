// TODO: Проблема: допустим есть несколько модулей, которым требуется пробрасывание кросс-контекста между сервисами
// например логирование и трейсинг, и обоим нужны RequestID.
// Возникает трабл: нужно это объединить в одном месте, что бы каждый модуль observer'а не повторял одну и туже работу,
// которая НЕ ЕГО зона ответственности
// Логирование - логирует, Оно блять не занимается пробрасыванием метаданных для логирования в другом сервисе.
// Написать поверх этого крутой обсервер - я пососу хуй и КАК БЫ я не сделал, прийдется переписывать.
// Поэтому самый менее трудозатратный способ - вкостылить все константы в один контекст, и разбирать их от туда по пакетам

package cross

import (
	"context"
)

type CrossCtxKey struct{}

var CrossKey = CrossCtxKey{}

type CrossValues map[string]any

// вернет значения из контекста
func GetContextAttrs(ctx context.Context) CrossValues {
	ctxCrossValues, ok := ctx.Value(CrossKey).(CrossValues)
	if ctxCrossValues == nil || !ok {
		return CrossValues{}
	}
	return ctxCrossValues
}

// Добавит в контекст кросс-значение
func WithField(ctx context.Context, k string, v any) context.Context {
	values := GetContextAttrs(ctx)
	values[k] = v
	return context.WithValue(ctx, CrossKey, values)
}

// Добавит в контекст кросс-значения
func WithFields(ctx context.Context, values map[string]any) context.Context {
	ctxvalues := GetContextAttrs(ctx)
	for k, v := range values {
		ctxvalues[k] = v
	}
	return context.WithValue(ctx, CrossKey, ctxvalues)
}
