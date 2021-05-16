/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/5/16 8:07 下午
* Description:
 */
package biz


type Article struct {
	Title string
}

type ArticleRepo interface {
	CreateArticle(*Article) error
	UpdateArticle(*Article) error
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

func (uc *ArticleUsecase) Create(a *Article) error {
	return uc.repo.CreateArticle(a)
}
