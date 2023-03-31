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
	"github.com/JuanVF/personal_bot/common"
)

const (
	OPEN_AI int = iota
)

type Classifier struct {
	Model int
}

type ClassifierHandler interface {
	Classify(prompt string) ([]string, error)
}

// Returns the classifier selected
func (c *Classifier) GetClassifier() ClassifierHandler {
	if c.Model == OPEN_AI {
		return OpenAIClassifier{}
	}

	common.GetLogger().Log("Classifier", "Unsupported Classifier")

	return nil
}
