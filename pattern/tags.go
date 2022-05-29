package pattern

import (
	"math/rand"
	"strings"
	"time"
)

// 文章中一種類に統一されるタグ
var uniqTags = map[string][]string{
	// 多様性が求められるママ
	"{MAMA}": {
		"ママ",
		"mama",
		"mama★",
		"MAMA",
	},
	"{KATSU}": {
		"活",
		"カツ",
		"かつ",
	},
	"{BR}": {
		"\n",
	},
	"{AGE}": {
		"20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", // サンプル中の最高齢が 32 だったため 32 にする
	},
	"{MONEY}": { // 稼げるお小遣い
		"10", "20", "100",
	},
	"{LINE}": {
		"LINE",
		"らいん",
		"ライン",
	},
}

// 絵文字。サンプルにあるものをすべて取り出したが全く法則性が見えない。
var emojis = []string{
	"✨",
	"❥",
	"❤",
	"🐀",
	"🛁",
	"🌂",
	"📃",
	"🍄",
	"😅",
	"🙇",
	"🐉",
	"💋",
	"🌎",
	"👏",
	"🏐",
	"🧑",
	"🧔",
	"😕",
	"🦖",
	"💗",
	"🦘",
	"🌙",
	"🥞",
	"👟",
	"🦠",
	"👡",
	"🤣",
	"🌤",
	"💦",
	"🐧",
	"👨",
	"🎩",
	"😪",
	"🌬",
	"🌭",
	"🐮",
	"🎯",
	"🕰",
	"🌴",
	"📷",
	"💸",
	"🤹",
	"🦻",
	"🍃",
	"👑",
	"🌒",
	"🥭",
	"🤐",
	"🍳",
	"🦡",
	"🏰",
	"🌞",
	"💕",
	"🤨",
	"🎺",
	"😃",
	"🦒",
	"😉",
	"😧",
	"🤽",
	"🦓",
	"🦁",
	"🤩",
	"🔬",
	"🙂",
	"🦯",
	"🎍",
	"🛫",
	"😆",
	"🎭",
	"🌻",
	"𓀘",
	"🥕",
	"𓀋",
	"🌫",
	"🙈",
	"🐱",
	"🍤",
	"🦿",
	"⭐️",
	"❤️",
	"☺️",
	"☘️",
	"👩🏻‍❤️‍👨🏻",
	"🧑🏽‍❤️‍🧑🏼",
	"👩🏼‍❤️‍👩🏼",
	"🧑🏻‍❤️‍💋‍🧑🏿",
	"🧔🏾‍♂️",
	"🏋️♂️",
	"🤽‍♀️",
	"🙇‍♀️",
	"🪆,",
	"𓀚,",
	"🪥,",
	"🪶,",
}

// ConvertTags ; message 内にあるタグを置換して結果を返す
func ConvertTags(message string) string {
	rand.Seed(time.Now().UnixNano())
	for tag, pat := range uniqTags {
		content := pat[rand.Intn(len(pat))]
		message = strings.Replace(message, tag, content, -1)
	}
	return message
}

// ConvertTags : message 内にある {EMOJI} を絵文字に置換して結果を返す
// emojiPattern が true の場合、絵文字の前に改行を入れることで、ママ活勧誘DMに時折見られる風情ある文体にする
// emojiPattern が false の場合、比較的一般的な文体になり、連続 3 つまで絵文字を並べる
func ConvertEmoji(message string, emojiPattern bool) string {
	rand.Seed(time.Now().UnixNano())
	tag := "{EMOJI}"
	pat := emojis
	n := strings.Count(message, tag)
	for i := 0; i < n; i++ {
		content := ""
		if emojiPattern {
			content = "\n" + pat[rand.Intn(len(pat))]
		} else {
			content = combineMultiplePatterns(pat, rand.Intn(3)+1)
		}
		// タグを置換
		message = strings.Replace(message, tag, content, 1)
	}
	// なぜか風情あるパターンでは、末尾に絵文字が追加される
	if emojiPattern {
		message = message + pat[rand.Intn(len(pat))]
	}
	return message
}

// combineMultiplePatterns: 複数のパターンをnumber分ランダムにつなげる
func combineMultiplePatterns(patterns []string, number int) string {
	result := ""
	if number <= len(patterns) {
		for i := 0; i < number; i++ {
			index := rand.Intn(len(patterns) - i)
			result += patterns[index]
			patterns[index], patterns[len(patterns)-1-i] = patterns[len(patterns)-1-i], patterns[index]
		}
	} else {
		for i := 0; i < number; i++ {
			result += patterns[rand.Intn(len(patterns))]
		}
	}
	return result
}
