package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cluttrdev/deepl-go/deepl"
)

type Resume struct {
	Basics struct {
		Name     string `json:"name"`
		Headline string `json:"headline"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Location string `json:"location"`
		URL      struct {
			Label string `json:"label"`
			Href  string `json:"href"`
		} `json:"url"`
		CustomFields []struct {
			ID    string `json:"id"`
			Icon  string `json:"icon"`
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"customFields"`
		Picture struct {
			URL          string `json:"url"`
			Size         int    `json:"size"`
			AspectRatio  int    `json:"aspectRatio"`
			BorderRadius int    `json:"borderRadius"`
			Effects      struct {
				Hidden    bool `json:"hidden"`
				Border    bool `json:"border"`
				Grayscale bool `json:"grayscale"`
			} `json:"effects"`
		} `json:"picture"`
	} `json:"basics"`
	Sections struct {
		Summary struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Content string `json:"content"`
		} `json:"summary"`
		Awards struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []any  `json:"items"`
		} `json:"awards"`
		Certifications struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []struct {
				ID      string `json:"id"`
				Visible bool   `json:"visible"`
				Name    string `json:"name"`
				Issuer  string `json:"issuer"`
				Date    string `json:"date"`
				Summary string `json:"summary"`
				URL     struct {
					Label string `json:"label"`
					Href  string `json:"href"`
				} `json:"url"`
			} `json:"items"`
		} `json:"certifications"`
		Education struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []struct {
				ID          string `json:"id"`
				Visible     bool   `json:"visible"`
				Institution string `json:"institution"`
				StudyType   string `json:"studyType"`
				Area        string `json:"area"`
				Score       string `json:"score"`
				Date        string `json:"date"`
				Summary     string `json:"summary"`
				URL         struct {
					Label string `json:"label"`
					Href  string `json:"href"`
				} `json:"url"`
			} `json:"items"`
		} `json:"education"`
		Experience struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []struct {
				ID       string `json:"id"`
				Visible  bool   `json:"visible"`
				Company  string `json:"company"`
				Position string `json:"position"`
				Location string `json:"location"`
				Date     string `json:"date"`
				Summary  string `json:"summary"`
				URL      struct {
					Label string `json:"label"`
					Href  string `json:"href"`
				} `json:"url"`
			} `json:"items"`
		} `json:"experience"`
		Volunteer struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []any  `json:"items"`
		} `json:"volunteer"`
		Interests struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []struct {
				ID       string   `json:"id"`
				Visible  bool     `json:"visible"`
				Name     string   `json:"name"`
				Keywords []string `json:"keywords"`
			} `json:"items"`
		} `json:"interests"`
		Languages struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []struct {
				ID          string `json:"id"`
				Visible     bool   `json:"visible"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Level       int    `json:"level"`
			} `json:"items"`
		} `json:"languages"`
		Profiles struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []struct {
				ID       string `json:"id"`
				Visible  bool   `json:"visible"`
				Network  string `json:"network"`
				Username string `json:"username"`
				Icon     string `json:"icon"`
				URL      struct {
					Label string `json:"label"`
					Href  string `json:"href"`
				} `json:"url"`
			} `json:"items"`
		} `json:"profiles"`
		Projects struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []struct {
				ID          string `json:"id"`
				Visible     bool   `json:"visible"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Date        string `json:"date"`
				Summary     string `json:"summary"`
				Keywords    []any  `json:"keywords"`
				URL         struct {
					Label string `json:"label"`
					Href  string `json:"href"`
				} `json:"url"`
			} `json:"items"`
		} `json:"projects"`
		Publications struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []any  `json:"items"`
		} `json:"publications"`
		References struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []any  `json:"items"`
		} `json:"references"`
		Skills struct {
			Name    string `json:"name"`
			Columns int    `json:"columns"`
			Visible bool   `json:"visible"`
			ID      string `json:"id"`
			Items   []struct {
				ID          string `json:"id"`
				Visible     bool   `json:"visible"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Level       int    `json:"level"`
				Keywords    []any  `json:"keywords"`
			} `json:"items"`
		} `json:"skills"`
		Custom struct {
		} `json:"custom"`
	} `json:"sections"`
	Metadata struct {
		Template string       `json:"template"`
		Layout   [][][]string `json:"layout"`
		CSS      struct {
			Value   string `json:"value"`
			Visible bool   `json:"visible"`
		} `json:"css"`
		Page struct {
			Margin  int    `json:"margin"`
			Format  string `json:"format"`
			Options struct {
				BreakLine   bool `json:"breakLine"`
				PageNumbers bool `json:"pageNumbers"`
			} `json:"options"`
		} `json:"page"`
		Theme struct {
			Background string `json:"background"`
			Text       string `json:"text"`
			Primary    string `json:"primary"`
		} `json:"theme"`
		Typography struct {
			Font struct {
				Family   string   `json:"family"`
				Subset   string   `json:"subset"`
				Variants []string `json:"variants"`
				Size     float64  `json:"size"`
			} `json:"font"`
			LineHeight     float64 `json:"lineHeight"`
			HideIcons      bool    `json:"hideIcons"`
			UnderlineLinks bool    `json:"underlineLinks"`
		} `json:"typography"`
		Notes string `json:"notes"`
	} `json:"metadata"`
}

// getResumeFromJSON takes JSON data and returns a resume and an error.
func getResumeFromJSON(data []byte) (Resume, error) {
	rsm := Resume{}
	err := json.Unmarshal(data, &rsm)
	//fmt.Printf("%+v\n", rsm) // Print output
	return rsm, err
}

// toJSON transforms the resume into JSON and returns a slice of bytes for that JSON and an error.
func (rsm Resume) toJSON() ([]byte, error) {
	bs, err := json.Marshal(rsm)

	//  Use below if you want JSON pretty printed
	var prettyJSON bytes.Buffer
	_ = json.Indent(&prettyJSON, bs, "", "    ")

	return prettyJSON.Bytes(), err
}

// Export takes a file name and exports the resume to as a json to the specified filename
func (rsm Resume) Export(fname string) error {
	// Create file
	f, err := os.Create(fname)
	if err != nil {
		// TODO: wrap error
		return err
	}
	defer f.Close()
	// Get data in right format
	data, err := rsm.toJSON()
	if err != nil {
		// TODO: wrap error
		return err
	}
	// Export data
	_, err = f.Write(data)
	if err != nil {
		// TODO: wrap error
		return err
	}
	return nil
}

func GetResume(fname string) (Resume, error) {
	data, err := os.ReadFile(fname)
	if err != nil {
		// TODO: wrap error
		return Resume{}, nil
	}
	rsm, err := getResumeFromJSON(data)
	return rsm, err
}

func (r Resume) Translate(apiKey, targetLang string) (Resume, error) {
	t, err := deepl.NewTranslator(apiKey)
	if err != nil {
		return Resume{}, err
	}

	var translate = func(field *string) {
		translation, err := t.TranslateText([]string{*field}, targetLang)
		if err != nil {
			log.Printf("Unable to translate '%v' with error: '%s'", field, err)
		} else {
			*field = translation[0].Text
		}
	}
	translate(&r.Basics.Headline)
	translate(&r.Sections.Summary.Name)
	translate(&r.Sections.Summary.Content)
	translate(&r.Sections.Awards.Name)
	translate(&r.Sections.Certifications.Name)
	for i := range r.Sections.Certifications.Items {
		translate(&r.Sections.Certifications.Items[i].Name)
		translate(&r.Sections.Certifications.Items[i].Date)
		translate(&r.Sections.Certifications.Items[i].Summary)
	}
	translate(&r.Sections.Education.Name)
	for i := range r.Sections.Education.Items {
		//translate(&r.Sections.Education.Items[i].StudyType)
		translate(&r.Sections.Education.Items[i].Date)
		translate(&r.Sections.Education.Items[i].Summary)
	}
	translate(&r.Sections.Experience.Name)
	for i := range r.Sections.Experience.Items {
		translate(&r.Sections.Experience.Items[i].Position)
		translate(&r.Sections.Experience.Items[i].Date)
		translate(&r.Sections.Experience.Items[i].Summary)
	}
	translate(&r.Sections.Volunteer.Name)
	translate(&r.Sections.Interests.Name)
	for i := range r.Sections.Interests.Items {
		translate(&r.Sections.Interests.Items[i].Name)
		// for j, _ := range r.Sections.Interests.Items[i].Keywords {
		// 	translate(&r.Sections.Interests.Items[i].Keywords[j])
		// }
	}
	translate(&r.Sections.Languages.Name)
	for i := range r.Sections.Languages.Items {
		translate(&r.Sections.Languages.Items[i].Name)
		translate(&r.Sections.Languages.Items[i].Description)
	}
	translate(&r.Sections.Profiles.Name)
	translate(&r.Sections.Projects.Name)
	for i := range r.Sections.Projects.Items {
		translate(&r.Sections.Projects.Items[i].Description)
		translate(&r.Sections.Projects.Items[i].Date)
		translate(&r.Sections.Projects.Items[i].Summary)
	}
	translate(&r.Sections.Skills.Name)
	for i := range r.Sections.Skills.Items {
		translate(&r.Sections.Skills.Items[i].Name)
		translate(&r.Sections.Skills.Items[i].Description)
	}
	return r, nil
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Usage: resume deepl-translator <API_KEY> <REACTIVE_RESUME_JSON_FILENAME> <TARGET_LANG>") 
		// https://developers.deepl.com/docs/api-reference/translate/openapi-spec-for-text-translation
		return
	}

	apiKey := os.Args[1]
	fname := os.Args[2]
	targetLang := os.Args[3]

	rsm, err := GetResume(fname)
	if err != nil {
		log.Fatalf("Error importing JSON: %s", err)
		return
	}

	rsmTranslated, err := rsm.Translate(apiKey, targetLang)
	if err != nil {
		log.Printf("Error translating Resume: %s", err)
	}

	fnameTranslated := fmt.Sprintf("%s %s.json", fname[:len(fname)-5], targetLang)
	err = rsmTranslated.Export(fnameTranslated)
	if err != nil {
		log.Print("Error exporting Resume:", err)
		fmt.Print(rsmTranslated)
		return
	}
	log.Print("Succesfully translated resume and exported to:", fmt.Sprintf("%s %s", fname, targetLang))
}
