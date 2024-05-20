package freq

import (
	"github.com/tyto-go/internal/distance"
	"github.com/tyto-go/internal/dna"
)

func FrequentWords(Text string, k int) []string {
	word_freq := make(map[string]int)
	T := len(Text)
	max_freq := 0
	for i := 0; i <= T-k; i++ {
		word_freq[Text[i:i+k]]++
		if word_freq[Text[i:i+k]] > max_freq {
			max_freq = word_freq[Text[i:i+k]]
		}
	}
	output := make([]string, 0)
	for word, freq := range word_freq {
		if freq == max_freq {
			output = append(output, word)
		}
	}
	return output
}

func FrequentWordsWithMismatches(Text string, k, d int, reverse bool) []string {
	patterns := make([]string, 0)
	freqMap := make(map[string]int)
	T := len(Text)
	maxFreq := 0
	var reverseText string
	if reverse {
		reverseText = dna.ReverseComplement(Text)
	}
	for i := 0; i <= T-k; i++ {
		neighbors := distance.Neighbors(Text[i:i+k], d)
		if reverse {
			neighbors = append(neighbors, distance.Neighbors(reverseText[i:i+k], d)...)
		}
		for _, neighbor := range neighbors {
			freqMap[neighbor]++
			if freqMap[neighbor] > maxFreq {
				maxFreq = freqMap[neighbor]
			}
		}
	}
	for pattern, freq := range freqMap {
		if freq == maxFreq {
			patterns = append(patterns, pattern)
		}
	}
	return patterns
}
