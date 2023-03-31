/*
Copyright 2023 Juan Jose Vargas Fletes

This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/

Under the CC BY-NC license, you are free to:

- Share: copy and redistribute the material in any medium or format
- Adapt: remix, transform, and build upon the material

Under the following terms:

  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

- Non-Commercial: You may not use the material for commercial purposes.

You are free to use this work for personal or non-commercial purposes.
If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.
*/
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
