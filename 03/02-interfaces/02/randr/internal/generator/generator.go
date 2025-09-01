package generator

import (
	"encoding/binary"
	"errors"
	"io"
	"math/rand"
)

type generator struct {
	rnd rand.Source // Генератор случайных чисел. Вообще rand.Rand уже реализует интерфейс io.Reader, но для примера мы реализуем его самостоятельно.
}

// New — обратите внимание, что мы возвращаем generator, присвоенный интерфейсу io.Reader, сама структура generator неэкспортируемая.
// Мы скрыли внутри пакета все детали.
func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

// Read — реализация io.Reader
func (g *generator) Read(buf []byte) (n int, err error) { // error — это тип ошибки, подробнее мы рассмотрим его в следующем разделе.
	if len(buf) == 0 {
		return 0, errors.New("error: zero length buffer")
	}

	for n = 0; n+8 <= len(buf); n += 8 {
		randInt := g.rnd.Int63() // функция возвращает положительное число в пределах от 0 до 2^63
		binary.LittleEndian.PutUint64(buf[n:], uint64(randInt))
	}

	return len(buf), nil
}
