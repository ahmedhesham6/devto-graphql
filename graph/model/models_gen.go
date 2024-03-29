// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Article struct {
	TypeOf                 string        `json:"type_of"`
	ID                     int           `json:"id"`
	Title                  string        `json:"title"`
	Description            string        `json:"description"`
	CoverImage             *string       `json:"cover_image"`
	ReadablePublishDate    string        `json:"readable_publish_date"`
	SocialImage            string        `json:"social_image"`
	TagList                []string      `json:"tag_list"`
	Tags                   string        `json:"tags"`
	Slug                   string        `json:"slug"`
	Path                   string        `json:"path"`
	URL                    string        `json:"url"`
	CanonicalURL           string        `json:"canonical_url"`
	CommentsCount          int           `json:"comments_count"`
	PositiveReactionsCount int           `json:"positive_reactions_count"`
	PublicReactionsCount   int           `json:"public_reactions_count"`
	CreatedAt              string        `json:"created_at"`
	EditedAt               string        `json:"edited_at"`
	CrosspostedAt          *string       `json:"crossposted_at"`
	PublishedAt            *string       `json:"published_at"`
	LastCommentAt          string        `json:"last_comment_at"`
	PublishedTimestamp     string        `json:"published_timestamp"`
	User                   *User         `json:"user"`
	ReadingTimeMinutes     int           `json:"reading_time_minutes"`
	Organization           *Organization `json:"organization"`
	FlareTag               *FlareTag     `json:"flare_tag"`
}

type ArticlesQuery struct {
	Page         *int          `json:"page"`
	PerPage      *int          `json:"per_page"`
	Tag          *string       `json:"tag"`
	Tags         *string       `json:"tags"`
	TagsExclude  *string       `json:"tags_exclude"`
	Username     *string       `json:"username"`
	State        *ArticleState `json:"state"`
	Top          *int          `json:"top"`
	CollectionID *int          `json:"collection_id"`
}

type FlareTag struct {
	Name         string `json:"name"`
	BgColorHex   string `json:"bg_color_hex"`
	TextColorHex string `json:"text_color_hex"`
}

type Organization struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Slug           string `json:"slug"`
	ProfileImage   string `json:"profile_image"`
	ProfileImage90 string `json:"profile_image_90"`
}

type PaginationQuery struct {
	Page    *int `json:"page"`
	PerPage *int `json:"per_page"`
}

type User struct {
	Name            string  `json:"name"`
	Username        string  `json:"username"`
	TwitterUsername *string `json:"twitter_username"`
	GithubUsername  *string `json:"github_username"`
	WebsiteURL      *string `json:"website_url"`
	ProfileImage    string  `json:"profile_image"`
	ProfileImage90  string  `json:"profile_image_90"`
}

type VideoArticle struct {
	TypeOf                 string     `json:"type_of"`
	ID                     int        `json:"id"`
	Path                   string     `json:"path"`
	CloudinaryVideoURL     string     `json:"cloudinary_video_url"`
	Title                  string     `json:"title"`
	UserID                 int        `json:"user_id"`
	VideoDurationInMinutes string     `json:"video_duration_in_minutes"`
	VideoSourceURL         *string    `json:"video_source_url"`
	User                   *VideoUser `json:"user"`
}

type VideoUser struct {
	Name string `json:"name"`
}

type ArticleState string

const (
	ArticleStateFresh  ArticleState = "fresh"
	ArticleStateRising ArticleState = "rising"
	ArticleStateAll    ArticleState = "all"
)

var AllArticleState = []ArticleState{
	ArticleStateFresh,
	ArticleStateRising,
	ArticleStateAll,
}

func (e ArticleState) IsValid() bool {
	switch e {
	case ArticleStateFresh, ArticleStateRising, ArticleStateAll:
		return true
	}
	return false
}

func (e ArticleState) String() string {
	return string(e)
}

func (e *ArticleState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ArticleState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ArticleState", str)
	}
	return nil
}

func (e ArticleState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
