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
