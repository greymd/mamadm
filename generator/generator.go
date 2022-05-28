package generator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/greymd/mamadm/pattern"
)

// Generate ... メッセージを生成
// taste: 1 or 2 で文体を決定。0 はランダム。
func Generate(taste int) (string, error) {
	rand.Seed(time.Now().UnixNano())
  selectedMessage := selectMessage()
	// メッセージに含まれるタグを変換
	selectedMessage = pattern.ConvertTags(selectedMessage)
  // 末尾の不要な改行を除く
  selectedMessage = strings.TrimSuffix(selectedMessage, "\n")
	// 絵文字の変換
  n := taste
  if n == 0 {
    n = rand.Intn(2)+1
  }
	switch n {
  case 1:
	  selectedMessage = pattern.ConvertEmoji(selectedMessage, true)
  case 2:
	  selectedMessage = pattern.ConvertEmoji(selectedMessage, false)
  default:
	  return "", fmt.Errorf("文体の指定が不正です: %v", n)
  }
	return selectedMessage, nil
}

func selectMessage() string {
	rand.Seed(time.Now().UnixNano())
	selectedMessage := ""

	// 文章の流れ戦略を無作為に選定
	selectedStrategy := pattern.MamaStrategy[rand.Intn(len(pattern.MamaStrategy))]

  // 10 % くらいの確率で末筆を追加する
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	if n <= 10 {
    selectedStrategy = append(selectedStrategy, pattern.END)
  }

	// 重複した表現を避けるためのブラックリストを戦略ごとに用意
	blacklist := map[pattern.MessagePattern]map[int]bool{}
	for i := range pattern.MamaMessages {
		blacklist[pattern.MessagePattern(i)] = make(map[int]bool)
	}

	// アルゴリズム内で表現されたそれぞれの戦略に対応した文言を選定
	for _, s := range selectedStrategy {
		selected := pattern.MamaMessages[s]
		index := 0
		for {
			index = rand.Intn(len(selected))
			if !blacklist[s][index] {
				blacklist[s][index] = true
				selectedMessage += selected[index]
				break
			}
			// 既にすべての表現を使い切っていたら諦める
			if len(blacklist[s]) >= len(selected) {
				blacklist[s][index] = true
				selectedMessage += selected[index]
				break
			}
		}
	}
	return selectedMessage
}
