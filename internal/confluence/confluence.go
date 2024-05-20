package confluence

import (
    "log"
    "myproject/internal/config"
    "github.com/virtomize/confluence-go-api"
)

// CreatePage creates a Confluence page with the Draw.io diagram
func CreatePage(drawioXML string) error {
    cfg := config.LoadConfig()

    client := confluence.NewAPI(cfg.ConfluenceAPIURL, cfg.ConfluenceUsername, cfg.ConfluenceAPIToken)

    content := confluence.Content{
        Title: "DFD Diagram",
        Type:  "page",
        Space: confluence.Space{Name: "YOUR_SPACE_KEY"},
        Body: confluence.Body{
            Storage: confluence.Storage{
                Value: `<ac:structured-macro ac:name="drawio">
                          <ac:parameter ac:name="diagramName">DFD Diagram</ac:parameter>
                          <ac:parameter ac:name="diagramData">` + drawioXML + `</ac:parameter>
                        </ac:structured-macro>`,
                Representation: "storage",
            },
        },
    }

    newContent, err := client.CreateContent(content)
    if err != nil {
        return err
    }

    log.Printf("Confluence page created: %s\n", newContent.Links.Webui)
    return nil
}