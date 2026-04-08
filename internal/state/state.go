package state

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/Sheridanlk/Service/internal/domain/models"
	"gopkg.in/yaml.v3"
)

var ErrStateNotFound = errors.New("state file not found")

type State struct {
	filePath string

	mu   sync.RWMutex
	data models.NodeState
}

func New(path string) *State {
	return &State{
		filePath: path,
		data:     models.NodeState{},
	}
}

func (s *State) Load() error {
	s.mu.RLock()
	defer s.mu.Unlock()

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrStateNotFound
		}
		return fmt.Errorf("failed to read state file: %w", err)
	}

	var st models.NodeState

	if err := yaml.Unmarshal(data, &st); err != nil {
		return fmt.Errorf("failed tp parse state file: %w", err)
	}

	s.data = st

	return nil
}

func (s *State) Save() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := yaml.Marshal(s.data)
	if err != nil {
		return fmt.Errorf("failed to parse state file: %w", err)
	}

	err = os.WriteFile(s.filePath, data, 0600)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}

func (s *State) UpdateTokens(tokens models.AuthTokens) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data.Tokens = tokens

	return s.Save()
}

func (s *State) GetTokens() models.AuthTokens {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.Tokens
}

func (s *State) GetServerSecret() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data.ServerSecret
}
