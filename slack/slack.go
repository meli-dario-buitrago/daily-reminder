package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/meli-dario-buitrago/daily-reminder/config"
)

var (
	DangerStyle  = "danger"
	PrimaryStyle = "primary"
	helloMessage = "<!here> :sunny: Buenos días equipo, *¡¡es hora del daily!!*, nuestros facilitadores hoy son:"
)

func NewMarkDownBlock(text string) PlanTextBlock {
	return PlanTextBlock{
		Type: "mrkdwn",
		Text: text,
	}
}

func NewMarkDownSection(text string) HeaderBlock {
	return HeaderBlock{
		Type: "section",
		Text: NewMarkDownBlock(text),
	}
}

func NewFieldsSection(fields []PlanTextBlock) FieldBlock {
	return FieldBlock{
		Type:   "section",
		Fields: fields,
	}
}

func SendSlackMessage() {
	strWeekDay := time.Now().Format("Monday")
	if strings.Contains(strWeekDay, "Sunday") || strings.Contains(strWeekDay, "Saturday") {
		return
	}
	message := WebhookMessage{}
	helloSection := NewMarkDownSection(helloMessage)
	gopherbotsBlock := NewFieldsSection(getTodayGopherbots(strWeekDay))
	linksSection := NewMarkDownSection("Nuestros enlaces...")
	linksBlock := NewFieldsSection(getLinks())

	message.Blocks = append(message.Blocks, helloSection, gopherbotsBlock, linksSection, linksBlock)
	byteResponse, _ := json.MarshalIndent(message, "", "  ")
	sendNotification(byteResponse)
}

func getTodayGopherbots(weekday string) []PlanTextBlock {
	var todayGopherbotsBlock []PlanTextBlock
	todayGopherbotsBlock = append(todayGopherbotsBlock,
		NewMarkDownBlock(fmt.Sprintf("*Presentador:*\n<@%s>", config.GetPresenter(weekday))),
		NewMarkDownBlock(fmt.Sprintf("*Suplente:*\n<@%s>", config.GetAlternate(weekday))),
	)

	return todayGopherbotsBlock
}

func getLinks() []PlanTextBlock {
	var linksBlock []PlanTextBlock

	linksBlock = append(linksBlock,
		NewMarkDownBlock(fmt.Sprintf(":jira: *<%s|Tablero Jira>* :jira:", config.GetJiraUrl())),
		NewMarkDownBlock(fmt.Sprintf(":meet: *<%s|Meet>* :meet:", config.GetMeetUrl())),
	)

	return linksBlock
}

func sendNotification(message []byte) {
	log.Printf("%v", string(message))
	resp, err := http.Post(config.GetWebhookUrl(), "application/json", bytes.NewReader(message))

	if err != nil {
		log.Fatal(err)
	}

	bodyText, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(bodyText))
}
