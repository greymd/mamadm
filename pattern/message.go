package pattern

// MessagePattern ... ママ活勧誘員が文中で使うパターン
type MessagePattern int

const (
	// GREETING ... 挨拶
	GREETING MessagePattern = iota
	// INTRODUCTION ... 自己紹介 (self introduction)
	INTRODUCTION
	// PROPOSAL ... 商品の提案 (Proposal of product)
	PROPOSAL
	// BENEFIT ... 利益を主張
	BENEFIT
	// CONTACT ... 連絡手段提示
	CONTACT
	// END ... 末筆
	END
)

// MamaStrategy ... ママ活の世界に勧誘するための黄金パターン
var MamaStrategy = [][]MessagePattern{
	{GREETING, PROPOSAL, BENEFIT, CONTACT},
	{GREETING, INTRODUCTION, PROPOSAL, CONTACT},
	{GREETING, INTRODUCTION, PROPOSAL, BENEFIT, CONTACT},
	{PROPOSAL, BENEFIT, CONTACT},
}

var MamaMessages = [][]string{
	GREETING: {
		"お初にお目にかかります{EMOJI}{BR}",
		"急なメッセージ申し訳ございません{EMOJI}{BR}月に{MONEY}万以上稼げる副収入ないですか？{EMOJI}{BR}",
		"フォロワー外のダイレクトメッセージ失礼します{EMOJI}{BR}",
		"フォロワー外の連絡すみません{EMOJI}{BR}",
		"フォロワー外のDMしちゃってごめんなさい{EMOJI}{BR}",
		"いきなり、DMしちゃってごめんね{EMOJI}{BR}",
		"失礼します！{EMOJI}{BR}",
		"失礼します！{EMOJI}{BR}いきなりご連絡送ってごめんなさい{EMOJI}{BR}",
		"急に、連絡しちゃって申し訳ないです{EMOJI}{BR}",
		"ご連絡失礼します{EMOJI}{BR}",
		"連絡しちゃって申し訳ないです{EMOJI}{BR}",
		"初めまして{EMOJI}{BR}",
		"【{MAMA}{KATSU}男子募集中】{BR}{BR}",
		"、しちゃって申し訳ないです{EMOJI}{BR}",
		"DM失礼します{EMOJI}{BR}",
		"急なDM失礼します{EMOJI}{BR}",
		"急に、DMしちゃってすみません{EMOJI}{BR}",
		"", // たまに挨拶を省く無礼なママ活勧誘員
	},
	INTRODUCTION: {
		"綺麗な美女の相手をするだけでお小遣いが得られるまま活の仲介をしているものです{EMOJI}{BR}",
		"気になった人300人にご連絡しています{EMOJI}{BR}",
		"近年TVでも報道されている{MAMA}{KATSU}の仲介をしているものです{EMOJI}{BR}",
		"公式{MAMA}{KATSU}コミュニティ運営しています{EMOJI}{BR}",
		"{MAMA}{KATSU}が大好きな一般人です{EMOJI}{BR}全国各地の男性様を募集中です{EMOJI}{BR}",
		"日国内最高級の{MAMA}{KATSU}グループを運営しています{EMOJI}{BR}",
		"{MAMA}{KATSU}を愛してる20代です{EMOJI}{BR}",
		"副業したい方のケア、支援、仲介やってます{EMOJI}{BR}",
		"セレブな{MAMA}の相手をするだけでお小遣いが得られる{MAMA}{KATSU}の紹介をしているものです{BR}",
		"{MAMA}{KATSU}をおすすめしている{AGE}歳です{EMOJI}{BR}",
		"近年コロナ禍でテレビでも報道している{MAMA}{KATSU}の紹介をしているものです{EMOJI}{BR}",
		"{MAMA}{KATSU}をしたい方の支援をしているものです",
		"あなたのTwitterを拝見した結果、{EMOJI}{BR}{MAMA}{KATSU}男子にふさわしいと思いご連絡させて頂きました{EMOJI}{BR}",
		"あなたのTwitterを拝見した結果、{EMOJI}{BR}{MAMA}{KATSU}男子にふさわしいと思いDM送らせて頂きました{EMOJI}{BR}",
		"{MAMA}{KATSU}の仲介をしてる者です{EMOJI}{BR}",
		"{MAMA}{KATSU}ご案内できます{EMOJI}{BR}",
	},
	PROPOSAL: {
		"大流行中の{MAMA}{KATSU}動興味ないですか？",
		"いきなりなんですが、{MAMA}{KATSU}興味ないですか？{EMOJI}{BR}",
		"大流行中の{MAMA}{KATSU}興味ないですか？{EMOJI}{BR}",
		"重大報告です{EMOJI}{BR}{MAMA}{KATSU}希望{MAMA}さん増加に伴い{EMOJI}{MAMA}{KATSU}男子を招集してます{EMOJI}{BR}",
		"いきなりなんですけど{MAMA}{KATSU}って気になったりしますか？{EMOJI}{BR}",
		"{AGE}歳/{MAMA}さん紹介してます{EMOJI}{BR}仕事にやる気があって目標がある男の子募集してます{EMOJI}{BR}",
		"{MAMA}{KATSU}希望{MAMA}さん急増に伴い{MAMA}{KATSU}男子を招集しています{EMOJI}{BR}",
		"美人{MAMA}と遊んでお小遣いGETしませんか？{EMOJI}{BR}",
		"{MAMA}{KATSU}紹介できるんですけどやってみたい気持ちありませんか？{EMOJI}{BR}",
		"近年コロナ禍でテレビでも報道している{MAMA}{KATSU}の仲介をしているのですが興味あったりしませんか？{EMOJI}{BR}",
		"大流行中の{MAMA}{KATSU}にて{MAMA}{KATSU}男子が不足しています{EMOJI}{BR}是非この機会に{MAMA}{KATSU}初めませんか？{EMOJI}{BR}",
		"月収を安定して稼ぎたくないですか？？{EMOJI}{BR}",
		"夜のdateやベッドの上での遊びの相手をしてくれる男子を探してます{EMOJI}{BR}",
		"最近大流行中の{MAMA}{KATSU}してみませんか？{EMOJI}{BR}",
		"バズり中の{MAMA}{KATSU}興味ないですか？",
	},
	BENEFIT: {
	  "{MAMA}の相手をしちゃうだけでお給料が得られる{EMOJI}{BR}今流行りのアルバイトです{EMOJI}{BR}",
	  "初心者の方向けの徹底的{サポート|アシスト|ケアもしています",
		"毎月で{MONEY}万ぐらい稼げます！{EMOJI}{BR}",
		"お小遣いは{MONEY}万ぐらい稼げます{EMOJI}{BR}",
		"セレブな女性とデートしてお小遣い💸をもらう活動なんですが、興味ありませんか？{EMOJI}{BR}",
		"支援充実{EMOJI}{BR}",
		"期間限定相談受付可能{EMOJI}{BR}",
		"綺麗で欲求不満{MAMA}さんと連絡とりあうだけで{EMOJI}{BR}お小遣いお小遣いもらえちゃいます{EMOJI}{BR}",
		"紹介後のアフターケア。未経験の方向けの徹底的ケアもしているので{EMOJI}{BR}",
		"週給で{MONEY}万ぐらい稼いでる人多いです！{EMOJI}{BR}",
		"他人に縛られてる男性様をお助けします{EMOJI}{BR}セレブ{MAMA}と遊んでお小遣いGETしませんか？{EMOJI}{BR}",
		"ノーリスクで稼げる最強のアルバイトなんですけど、{EMOJI}{BR}月に{MONEY}万ぐらい稼げます{EMOJI}{BR}",
		/**
		  サンプルは少ないが、メンズ不足を嘆いている DM には「稼げる」といったメッセージが見られない傾向を感じる (@greymd)。
		  ブルーオーシャンであることを主張してすでに稼げることを暗に示しているため、省いているのかもしれない。
		  人材不足の旨は、 BENEFIT に統合してしまう。
		*/
		"女の子の募集が殺到しており現状男子の人数が足りていません{EMOJI}{BR}",
		"いけてるメンズいなくて困ってます・・・。{EMOJI}{BR}",
		"男性様少なくて困ってます・・・。{EMOJI}{BR}",
		"男性の人数が一定になった時点で募集は終わるので{EMOJI}{BR}",
	},
	CONTACT: {
		"興味があったらはじめからさいごまでアシストするので固定ツイート見に来てください{EMOJI}{BR}",
		"興味があれば最初から最後までサポートするので固定ツイートご覧ください{EMOJI}{BR}",
		"興味あったら固定ツイートから連絡ください{EMOJI}{BR}",
		"興味あったらツイートから連絡ください{EMOJI}{BR}",
		"{MAMA}{KATSU}男子募集中なので興味ある方は固定ツイート見て下さい{EMOJI}{BR}",
		"固定ツイートからご連絡頂けましたらすぐにはじめれます！{EMOJI}{BR}",
		"固定ツイートから{LINE}頂けましたらすぐにはじめれます！{EMOJI}{BR}",
		"固定ツイートから{LINE}頂けましたら即はじめれます{EMOJI}{BR}",
		"はブログから{LINE}にきてね{EMOJI}{BR}",
		"男子の方の人数が足りていない状態なのでもし気になれば固定ツイートみていただけると嬉しいです{EMOJI}{BR}",
		"{MAMA}{KATSU}をしたい方は固定ツイートから追加してね{EMOJI}{BR}",
		"試したい人は固定ツイートから来てください{EMOJI}{BR}",
		"{MAMA}{KATSU}男子募集中なので興味ある方は固定ツイート見て下さい{EMOJI}{BR}",
		"少しでも気になればsupport致しますので固定ツイート見て下さい{EMOJI}{BR}",
		"興味 ある男子はは固定ツイートから{LINE}にきてください{EMOJI}{BR}", // 誤字は原文ママ（ママ活だけに）
		"もし興味かあれば固定ツイート見に来てください{EMOJI}{BR}",
		"もし興味かあれば固定ツイートご覧ください{EMOJI}{BR}",
		"固定ツイートからご連絡ください{BR}",
		"興味があれば、固定ツイート見てくださいね{BR}",
		"バズり中の為締め切り間近なので{EMOJI}{BR}固定ツイートからお早めに{LINE}ください{EMOJI}{BR}",
		"超人気の為締め切り間近なので{EMOJI}{BR}固定ツイートからお早めに{LINE}ください{EMOJI}{BR}",
		"是非{MAMA}{KATSU}LIFE開始したい男性様は固定ツイート見て下さい",
		"もし気になる方は早めに固定ツイート見てください{EMOJI}{BR}",
		"是非{MAMA}{KATSU}生活STARTしたい男性様は固定ツイート見て下さい",
	},
	END: {
		"早い者勝ちです{EMOJI}{BR}",
		"締め切り間近です{EMOJI}{BR}",
		"身バレ等一切ないので安心してください{EMOJI}{BR}",
		"全国各地応募可能です{EMOJI}{BR}",
	},
}
