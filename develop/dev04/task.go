package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
1.Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
2.Выходные данные: ссылка на мапу множеств анаграмм
3.Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
4.Массив должен быть отсортирован по возрастанию.
5.Множества из одного элемента не должны попасть в результат.
6.Все слова должны быть приведены к нижнему регистру.
7.В результате каждое слово должно встречаться только один раз.
*/

func getAnagrams(s []string) map[string][]string {
	grupping := make(map[string][]string)

	for _, text := range s {
		text = strings.ToLower(text)
		key := sortText(text)
		grupping[key] = append(grupping[key], text)
	}
	res := make(map[string][]string, len(grupping))
	for _, words := range grupping {
		if len(words) > 1 {
			res[words[0]] = words
		}
	}
	return res
}

func sortText(s string) string {
	arr := []byte(s)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return string(arr)
}

func main() {
	s := []string{"кулон", "клоун", "газон", "загон"}
	anag := getAnagrams(s)
	fmt.Println(anag)
}
