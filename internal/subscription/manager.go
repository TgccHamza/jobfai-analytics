package subscription

import (
	"sync"
)

// EventHandler is a function that processes event data
type EventHandler func(data interface{}) error

// Subscription represents a single subscription to an event
type Subscription struct {
	ID      string
	Handler EventHandler
}

// Manager handles GraphQL subscriptions
type Manager struct {
	mu            sync.RWMutex
	subscriptions map[string]map[string]EventHandler
	nextID        int
}

// NewManager creates a new subscription manager
func NewManager() *Manager {
	return &Manager{
		subscriptions: make(map[string]map[string]EventHandler),
		nextID:        0,
	}
}

// Subscribe registers a handler for a specific event type
// Returns a subscription ID that can be used to unsubscribe
func (m *Manager) Subscribe(eventName string, handler EventHandler) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.subscriptions[eventName]; !exists {
		m.subscriptions[eventName] = make(map[string]EventHandler)
	}

	// Generate a unique subscription ID
	m.nextID++
	subID := string(m.nextID)

	m.subscriptions[eventName][subID] = handler
	return subID
}

// Unsubscribe removes a subscription
func (m *Manager) Unsubscribe(eventName string, subID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if handlers, exists := m.subscriptions[eventName]; exists {
		delete(handlers, subID)
	}
}

// Publish sends data to all subscribers of an event
func (m *Manager) Publish(eventName string, data interface{}) {
	m.mu.RLock()
	handlers := make(map[string]EventHandler)
	if subs, exists := m.subscriptions[eventName]; exists {
		for id, handler := range subs {
			handlers[id] = handler
		}
	}
	m.mu.RUnlock()

	// Execute handlers outside the lock to prevent deadlocks
	for _, handler := range handlers {
		// Ignore errors from individual handlers
		_ = handler(data)
	}
}
