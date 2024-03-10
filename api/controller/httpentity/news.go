package httpentity

import (
	"encoding/xml"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"

func (ct *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(ctLayout, v)
	if err != nil {
		return err
	}
	*ct = CustomTime{parse}
	return nil
}

type News struct {
    NewsItems []NewsItem `xml:"NewsletterNewsItems>NewsletterNewsItem" bson:"newsItems" json:"newsItems"`
}

type NewsItem struct {
    ArticleURL        string     `xml:"ArticleURL" bson:"articleURL" json:"articleURL"`
    NewsArticleID     string     `xml:"NewsArticleID" bson:"newsArticleID" json:"newsArticleID"`
    PublishDate       CustomTime `xml:"PublishDate" bson:"publishDate" json:"publishDate"`
    Taxonomies        string     `xml:"Taxonomies" bson:"taxonomies" json:"taxonomies"`
    TeaserText        string     `xml:"TeaserText" bson:"teaserText" json:"teaserText"`
    ThumbnailImageURL string     `xml:"ThumbnailImageURL" bson:"thumbnailImageURL" json:"thumbnailImageURL"`
    Title             string     `xml:"Title" bson:"title" json:"title"`
    OptaMatchId       string     `xml:"OptaMatchId" bson:"optaMatchId" json:"optaMatchId"`
    LastUpdateDate    CustomTime `xml:"LastUpdateDate" bson:"lastUpdateDate" json:"lastUpdateDate"`
    IsPublished       bool       `xml:"IsPublished" bson:"isPublished" json:"isPublished"`
}
