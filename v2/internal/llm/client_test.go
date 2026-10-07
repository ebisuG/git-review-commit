package llm

import (
	"testing"

	"github.com/mozilla-ai/any-llm-go/providers"
)

func TestReview(t *testing.T) {
	// stubLoader := config.NewStubLoader()
	// llmClient, err := NewReviewer(stubLoader)
	// if err != nil {
	// 	t.Errorf("failed NewReviewer : %v", err)
	// }

	// expectConf, err := stubLoader.Load()
	// if err != nil {
	// 	t.Errorf("failed stubLoader.Load() : %v", err)
	// }
	// if llmClient != expectConf.ApiKey {
	// 	t.Errorf("Failed to pass API_KEY to llmClient.")
	// }
	// if llmClient.ProviderAndModel != expectConf.ProviderAndModel {
	// 	t.Errorf("Failed to pass specification of Provier and Model in llmClient.")
	// }
	client, err := NewLlmClient("dummy-model", providers.Provider{})
	if err != nil {
		t.Errorf("failed to ")
	}
}
