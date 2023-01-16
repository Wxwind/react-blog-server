// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	articleFieldNames          = builder.RawFieldNames(&Article{})
	articleRows                = strings.Join(articleFieldNames, ",")
	articleRowsExpectAutoSet   = strings.Join(stringx.Remove(articleFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), ",")
	articleRowsWithPlaceHolder = strings.Join(stringx.Remove(articleFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), "=?,") + "=?"
)

type (
	articleModel interface {
		Insert(ctx context.Context, data *Article) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Article, error)
		Update(ctx context.Context, data *Article) error
		Delete(ctx context.Context, id int64) error
		Find(ctx context.Context, limitCount int32) ([]*Article, error)
	}

	defaultArticleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Article struct {
		Id          int64          `db:"id"`           // 自增id
		Title       string         `db:"title"`        // 标题
		ImageUrl    string         `db:"image_url"`    // 文章缩略图路径
		Describes   sql.NullString `db:"describes"`    // 文章简述
		PublishTime time.Time      `db:"publish_time"` // 发布时间
		UpdateTime  time.Time      `db:"update_time"`  // 更新时间
		MarkdownUrl string         `db:"markdown_url"` // 文章的md文件路径
	}
)

func newArticleModel(conn sqlx.SqlConn) *defaultArticleModel {
	return &defaultArticleModel{
		conn:  conn,
		table: "`article`",
	}
}

func (m *defaultArticleModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultArticleModel) FindOne(ctx context.Context, id int64) (*Article, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleRows, m.table)
	var resp Article
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultArticleModel) Insert(ctx context.Context, data *Article) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, articleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Title, data.ImageUrl, data.Describes, data.PublishTime, data.MarkdownUrl)
	return ret, err
}

func (m *defaultArticleModel) Update(ctx context.Context, data *Article) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, articleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Title, data.ImageUrl, data.Describes, data.PublishTime, data.MarkdownUrl, data.Id)
	return err
}

func (m *defaultArticleModel) Find(ctx context.Context, limitCount int32) ([]*Article, error) {
	query := fmt.Sprintf("select * from %s limit %d", m.table, limitCount)
	var resp []*Article
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultArticleModel) tableName() string {
	return m.table
}
