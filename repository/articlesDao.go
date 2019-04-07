package repository

import (
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/ozonebg/gofluence/db"
	"github.com/ozonebg/gofluence/interfaces"
	"github.com/ozonebg/gofluence/models"
)

const (
	articlesTableName = "articles"

	// NotFoundArticlesError is error if no article is found.
	NotFoundArticlesError = "no articles found"

	// InvalidDataError is error if input data is invalid.
	InvalidDataError = "invalid data"
)

var daoLogger = logrus.WithField("component", "articles dao")

type articlesDao struct {
}

// NewArticlesDao returns new ArticlesRepository
func NewArticlesDao() interfaces.ArticlesRepository {
	return &articlesDao{}
}

func (a *articlesDao) All() (models.Articles, error) {
	conn := db.CreateDbConnection()
	tx, err := conn.NewSession(nil).Begin()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer tx.RollbackUnlessCommitted()

	var articles models.Articles
	tx.Select("*").From(articlesTableName).Load(&articles)

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	if articles == nil {
		return models.Articles{}, errors.New(NotFoundArticlesError)
	}

	return articles, nil
}

func (a *articlesDao) GetArticle(id int) (*models.Article, error) {
	conn := db.CreateDbConnection()
	tx, err := conn.NewSession(nil).Begin()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	defer tx.RollbackUnlessCommitted()

	var article models.Article
	result, err := tx.Select("*").From(articlesTableName).Where("id = ?", id).Load(&article)
	if err != nil {
		return nil, err
	}

	if result != 1 {
		return nil, errors.New(NotFoundArticlesError)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	daoLogger.WithFields(logrus.Fields{"id": article.ID, "title": article.Title, "content": article.Content}).Info("the article")

	return &article, nil
}

func (a *articlesDao) CreateArticle(article *models.Article) error {
	if article == nil {
		return errors.New(InvalidDataError)
	}

	conn := db.CreateDbConnection()
	tx, err := conn.NewSession(nil).Begin()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer tx.RollbackUnlessCommitted()

	_, err = tx.InsertInto(articlesTableName).Columns("title", "content", "author_id").Record(article).Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *articlesDao) DeleteArticle(id int) error {
	conn := db.CreateDbConnection()
	tx, err := conn.NewSession(nil).Begin()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer tx.RollbackUnlessCommitted()

	_, err = tx.DeleteFrom(articlesTableName).Where("id = ?", id).Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *articlesDao) UpdateArticle(id int, updatedArticle *models.Article) error {
	conn := db.CreateDbConnection()
	tx, err := conn.NewSession(nil).Begin()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer tx.RollbackUnlessCommitted()

	updateMap := getUpdateMapForArticle(updatedArticle)

	_, err = tx.Update(articlesTableName).Where("id = ?", id).SetMap(updateMap).Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

// getUpdateMap returns a new map containing all updated values
func getUpdateMapForArticle(article *models.Article) map[string]interface{} {
	updateMap := make(map[string]interface{}, 3)
	updateMap["title"] = article.Title
	updateMap["content"] = article.Content

	return updateMap
}
