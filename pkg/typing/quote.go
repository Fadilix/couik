package typing

import (
	"math/rand"
	"os"
	"strings"

	"github.com/fadilix/couik/database"
)

var quotes = []string{
	// Programming quotes
	"The only way to do great work is to love what you do.",
	"Talk is cheap. Show me the code.",
	"Code is like humor. When you have to explain it, it's bad.",
	"The best way to predict the future is to invent it.",
	"First, solve the problem. Then, write the code.",
	"Any fool can write code that a computer can understand. Good programmers write code that humans can understand.",
	"Programming is not about typing, it's about thinking.",
	"Simplicity is the soul of efficiency.",
	"Make it work, make it right, make it fast.",
	"The most disastrous thing that you can ever learn is your first programming language.",
	"Programs must be written for people to read, and only incidentally for machines to execute.",
	"The function of good software is to make the complex appear to be simple.",
	"Debugging is twice as hard as writing the code in the first place.",
	"It's not a bug, it's an undocumented feature.",
	"The best error message is the one that never shows up.",
	"Clean code always looks like it was written by someone who cares.",
	"Truth can only be found in one place: the code.",
	"Before software can be reusable it first has to be usable.",
	"Give someone a program, you frustrate them for a day; teach them how to program, you frustrate them for a lifetime.",

	// Inspirational quotes
	"Success is not final, failure is not fatal: it is the courage to continue that counts.",
	"Stay hungry, stay foolish.",
	"The quick brown fox jumps over the lazy dog.",
	"To be, or not to be, that is the question.",
	"In the beginning God created the heaven and the earth.",
	"Computers are useless. They can only give you answers.",
	"The journey of a thousand miles begins with a single step.",
	"Innovation distinguishes between a leader and a follower.",
	"The future belongs to those who believe in the beauty of their dreams.",
	"It does not matter how slowly you go as long as you do not stop.",
	"Life is what happens when you're busy making other plans.",
	"The only thing we have to fear is fear itself.",
	"In three words I can sum up everything I've learned about life: it goes on.",
	"You miss one hundred percent of the shots you don't take.",
	"Whether you think you can or you think you can't, you're right.",
	"The greatest glory in living lies not in never falling, but in rising every time we fall.",
	"The way to get started is to quit talking and begin doing.",
	"Your time is limited, so don't waste it living someone else's life.",
	"If life were predictable it would cease to be life, and be without flavor.",
	"Spread love everywhere you go. Let no one ever come to you without leaving happier.",

	// Literary excerpts
	"It was the best of times, it was the worst of times, it was the age of wisdom, it was the age of foolishness.",
	"All that is gold does not glitter, not all those who wander are lost.",
	"It is a truth universally acknowledged, that a single man in possession of a good fortune, must be in want of a wife.",
	"Call me Ishmael. Some years ago, never mind how long precisely, having little or no money in my purse.",
	"Happy families are all alike; every unhappy family is unhappy in its own way.",
	"It was a bright cold day in April, and the clocks were striking thirteen.",
	"Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off—then, I account it high time to get to sea as soon as I can.",
}

var (
	wordsRaw = `the be to of and a in that have I it for not on with he as you do at this but his by from they we say her she or an will my one all would there their what so up out if about who get which go me when make can like time no just him know take people into year your good some could them see other than then now look only come its over think also back after use two how our work first well way even new want because any these give day most us is are was were been being had has have having does did doing would should could might must shall will can may need dare ought used able unable must might ought` +
		` able about above actually after again against all almost also although always am among an and another any anyone anything are around as at back be became because become been before being below between both but by came can cannot come could did do does doing done down during each either else enough especially even ever every everyone everything except few find first for found from further get give go goes going gone good got great had has have having he her here hers herself him himself his how however i if in into is it its itself just keep know known last later least less let like likely long look made make many may me might more most mostly much must my myself never new next no not nothing now of off often on once one only or other others our out over own part per perhaps please put quite rather really said same say see seem seemed seeming seems several she should show side since so some something sometime sometimes somewhere still such system take than that the their them themselves then there therefore these they thing think this those though thought three through thus time to together too toward try turn two under until up upon us use used using very want was way we well went were what when where whether which while who whole whose why will with within without work world would written year yet you your yours yourself` +
		` about above across after again against all almost alone along already also although always among amount another answer any anyone anything appear around ask away back bad become been before began begin behind being believe below beside best better between beyond big black blue body book both boy bring build business but buy by call came can cannot car care carry case cause certain change child children city close cold come common company consider contain continue control could country course cut day development did different do does done door down draw during each early east easy eat economic effect either end enough even ever every everyone everything example experience eye face fact fall family far fast father feel feet few field find fine first five follow food for force foreign form former forward found four free friend from front full further game gave general get girl give go going good got government great green ground group grow had half hand happen hard has have head hear heart heavy help here herself high him himself his history hold home hope hot hour house how however human hundred I idea if important in include increase indeed information interest into it its itself job join just keep kind know known land large last late later lead learn least leave left less let life light like likely line list little live local long look lose lot low made main major make man many market matter may me mean men might mind miss money month more morning most mother move much must my myself name nation national natural nature near necessary need never new next night no not note nothing now number of off offer office often old on once one only open or order other others ought our out over own page paper part particular party pass past pay people per perhaps period person personal place plan play point political poor position possible power present president press problem probably produce product program provide public put question quite range rather read ready real really reason receive recent remain report require research result return right room run said same saw say school science second section see seem send sense serve service set several shall she short should show side since sit situation six small so social society some someone something sometimes son soon sort sound south space speak special stand start state station story strong student study success such support sure system take talk tell ten term test than that the their them themselves then there therefore these they thing think third this those though thought thousand three through time to today together too took top toward town travel true try turn under understand unit unite university until up upon us use usually value various very want war water way we week well went were what when where whether which while white who whole whose why wide wife will wish with within without woman women word work world would write year yes yet you young your`
	dictionnary = strings.Fields(wordsRaw)
)

func GetRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}

func GetDictionnary() string {
	rand.Shuffle(len(dictionnary), func(i, j int) {
		dictionnary[i], dictionnary[j] = dictionnary[j], dictionnary[i]
	})
	return strings.Join(dictionnary, " ")
}

// GetQuoteUseCase fetches quotes from database
// and returns the list according to the language and category
func GetQuoteUseCase(lang database.Language, category database.Category) database.Quote {
	quotes := database.GetQuotes(lang, category)
	return quotes[rand.Intn(len(quotes))]
}

// GetQuoteFromFile a reads a file and returns the text
// inside of it
func GetQuoteFromFile(filepath string) (string, error) {
	quote, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(quote)), nil
}
