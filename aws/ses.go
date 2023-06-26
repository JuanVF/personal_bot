package aws

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// Creates a new SES Session
func getSESSession() *ses.SES {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return ses.New(sess)
}

// Send an email using ses
func SendEmail(tos []*string, templateName *string, templateValues map[string]*string) (*ses.SendTemplatedEmailOutput, error) {
	templateData, err := json.Marshal(templateValues)

	if err != nil {
		return nil, err
	}

	input := &ses.SendTemplatedEmailInput{
		Destination: &ses.Destination{
			ToAddresses: tos,
		},
		Template:     templateName,
		TemplateData: aws.String(string(templateData)),
		Source:       aws.String(config.AWS.SES.Sender),
	}

	session := getSESSession()

	return session.SendTemplatedEmail(input)
}
