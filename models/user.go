package models

import "time"

// Represents an individual event
type GitHubEvent struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}

// Represents the actor (user) who triggered the event
type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

// Represents the repository related to the event
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Represents the payload of the event (contains the event details)
type Payload struct {
	Action       string   `json:"action,omitempty"`
	PushID       int      `json:"push_id,omitempty"`
	Size         int      `json:"size,omitempty"`
	DistinctSize int      `json:"distinct_size,omitempty"`
	Ref          string   `json:"ref,omitempty"`
	Head         string   `json:"head,omitempty"`
	Before       string   `json:"before,omitempty"`
	Commits      []Commit `json:"commits,omitempty"`
}

// Represents the commit details in a PushEvent
type Commit struct {
	SHA      string `json:"sha"`
	Author   Author `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	URL      string `json:"url"`
}

// Represents the author of a commit
type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
