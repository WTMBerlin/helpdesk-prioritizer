package main

type WatsonTone struct {
	ID    string  `json:"tone_id"`
	Name  string  `json:"tone_name"`
	Score float64 `json:"score"`
}

type WatsonToneCategory struct {
	CategoryID   string       `json:"category_id"`
	CategoryName string       `json:"category_name"`
	Tones        []WatsonTone `json:"tones"`
}

type WatsonDocumentTone struct {
	ToneCategories []WatsonToneCategory `json:"tone_categories"`
}

type WatsonToneResponse struct {
	DocumentTone WatsonDocumentTone `json:"document_tone"`
}
