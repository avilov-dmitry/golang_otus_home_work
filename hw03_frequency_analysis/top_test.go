package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var englishText = `The Dursleys had everything they wanted, but they 
	also had a secret, and their greatest fear was that somebody would discover it.
	They did not think they could bear it if anyone found out about the Potters. Mrs.
	Potter was Mrs. Dursley sister, but they did not met for several years; in fact,
	Mrs. Dursley pretended she did not have a sister, because her sister and her
	good-for-nothing husband were as unDursleyish as it was possible to be.
	The Dursleys shuddered to think what the neighbors would say if the Potters
	arrived in the street. The Dursleys knew that the Potters had a small son, too,
	but they had never even seen him. This boy was another good reason for keeping the
	Potters away; they did not want Dudley mixing with a child like th-at.`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("correct for 1 word", func(t *testing.T) {
		input := "Одно.оооочень.большое,слово!сознаками@препинания"
		expected := []string{"Одно.оооочень.большое,слово!сознаками@препинания"}
		require.Equal(t, expected, Top10(input))
		require.Len(t, Top10(input), 1)
	})

	t.Run("less than 10 words", func(t *testing.T) {
		input := "one two 3 four пять 6 семь eight nine"
		expected := []string{"3", "6", "eight", "four", "nine", "one", "two", "пять", "семь"}
		require.Equal(t, expected, Top10(input))
		require.Len(t, Top10(input), 9)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})

	// t.Run("positive test", func(t *testing.T) {
	// 	expected := []string{
	// 		"the",      // 6
	// 		"they",     // 6
	// 		"a",        // 4
	// 		"did",      // 4
	// 		"not",      // 4
	// 		"had",      // 4
	// 		"was",      // 4
	// 		"but",      // 3
	// 		"for",      // 3
	// 		"Dursleys", // 3
	// 	}
	// 	require.Equal(t, expected, Top10(englishText))
	// })
}
