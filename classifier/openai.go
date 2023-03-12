package classifier

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/JuanVF/personal_bot/openai"
)

type OpenAIClassifier struct {
}

// Uses the OpenAI Classifier
func (o OpenAIClassifier) Classify(prompt string) ([]string, error) {
	tags, err := openai.Complete(fmt.Sprintf("payment: %s\ntags", prompt))

	if err != nil {
		return nil, err
	}

	return o.ParseTags(tags.Choices[0].Text), nil
}

// Parse the tags from the GPT response
func (o OpenAIClassifier) ParseTags(data string) []string {
	wordRgx := regexp.MustCompile("((\\w+)[^,])+")
	t := wordRgx.FindAllString(data, 10)
	tags := make([]string, 0)

	for _, tag := range t {
		tags = append(tags, strings.Replace(tag, "\\n", "", 10))
	}

	return tags
}
