package textmode

import (
	"log"
	"testing"

	textrank "github.com/DavidBelicza/TextRank"
	"github.com/k0kubun/pp/v3"
)

var rawText = []string{``}

func TestTextRank(t *testing.T) {

	tr := textrank.NewTextRank()
	rule := textrank.NewDefaultRule()
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewDefaultAlgorithm()
	tr.Populate(rawText[0], language, rule)
	tr.Ranking(algorithmDef)
	rankedPhrases := textrank.FindPhrases(tr)
	pp.Print(rankedPhrases)

	sentences := textrank.FindSentencesByRelationWeight(tr, 3)
	// Found sentences
	pp.Println(sentences)

	// Get the most important 10 sentences. Importance by word occurrence.
	sentences = textrank.FindSentencesByWordQtyWeight(tr, 3)
	// Found sentences
	pp.Println(sentences)
}

func TestSummary(t *testing.T) {
	t.Log("Processing HTML to TLDR")
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "niantic", input: "https://techcrunch.com/2022/03/10/pokemon-go-creator-niantic-is-acquiring-webar-development-platform-8th-wall/",
			want: `Niantic, the augmented reality platform behind Pokémon GO, is acquiring WebAR development platform 8th Wall, the company announced on Thursday.`,
		}, {
			name: "moloch", input: "https://slatestarcodex.com/2014/07/30/meditations-on-moloch/",
			want: `From a god’s-eye-view, we can optimize the system to “everyone agrees to stop doing this at once”, but no one within the system is able to effect the transition without great risk to themselves.`,
		}, {
			name: "new-yorker", input: "https://www.newyorker.com/news/letter-from-silicon-valley/the-lonely-work-of-moderating-hacker-news",
			want: `Technologists in Silicon Valley assume familiarity with Hacker News, just as New Yorkers do with the New York Post and the New York Times.`,
		}, {
			name: "hell-world", input: "https://luke.substack.com/p/the-man-who-bowled-a-perfect-game-c37?s=r",
			want: `Technologists in Silicon Valley assume familiarity with Hacker News, just as New Yorkers do with the New York Post and the New York Times.`,
		}, {
			name: "supa-base", input: "https://supabase.com/blog/supabase-realtime-multiplayer-general-availability",
			want: `Technologists in Silicon Valley assume familiarity with Hacker News, just as New Yorkers do with the New York Post and the New York Times.`,
		},
	}
	log.Println("Testing")
	for _, tc := range tests {
		log.Printf("Test: %s", tc.name)
		content, _ := RequestBody("https://earthly-tools.com/text-mode?url=" + tc.input)
		// log.Println("Content:", content)
		section := article{title: "", url: "", content: content}
		summary := ConvertSectionToSummary(section)
		pp.Println(summary)
		t.Errorf("Topic is wrong: %s", summary.Topic)
	}
}
