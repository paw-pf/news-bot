package fetcher

import (
	"context"
	"news-bot/internal/model"

	"time"
)

type ArticleStorage interface {
	Storage(ctx context.Context, article model.Article) error
}

type SourceProvider interface {
	Sources(ctx context.Context) ([]model.Source, error)
}

type Sources interface {
	ID() int64
	Name() string
	FeedCh(ctx context.Context) ([]model.Item, error)
}

type Fetcher struct {
	articles ArticleStorage
	sources  SourceProvider

	fetchInterval  time.Duration
	filterKeywords []string
}

func New(articleStorage ArticleStorage,
	sourceProvider SourceProvider,
	fetchInterval time.Duration,
	filterKeywords []string,
) *Fetcher {
	return &Fetcher{
		articles:       articleStorage,
		sources:        sourceProvider,
		fetchInterval:  fetchInterval,
		filterKeywords: filterKeywords,
	}
}

//func (f *Fetcher) Fetch(ctx context.Context) error {
//	sources, err := f.sources.Sources(ctx)
//	if err != nil {
//		return err
//	}
//
//	var wg sync.WaitGroup
//
//	//for _, source := range sources {
//	//	wg.Add(1)
//	//
//	//	rssSource := source.NewRSSSourceFromModel
//	//
//	//	go func(sources Sources) {
//	//
//	//	}
//	//}
//}
