package pg

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type TgUsersRepo struct {
	conn *pgx.Conn
}

func NewTgUserRepo(conn *pgx.Conn) *TgUsersRepo {
	return &TgUsersRepo{
		conn: conn,
	}
}

func (r *TgUsersRepo) Upsert(ctx context.Context, user domain.TgUser) (domain.TgUser, error) {
	_, err := r.conn.Exec(ctx, `
INSERT INTO tg_users
	(tg_id) VALUES
	(   $1)

ON CONFLICT DO NOTHING
`,
		user.TgId,
	)
	if err != nil {
		return domain.TgUser{}, err
	}

	return user, nil
}

func (r *TgUsersRepo) Get(ctx context.Context, tgId int64) (user domain.TgUser, err error) {
	err = r.conn.QueryRow(ctx, `
SELECT 
	u.tg_id
from tg_users as u
WHERE tg_id = $1
`,
		tgId,
	).Scan(&user.TgId)
	if err != nil {
		return domain.TgUser{}, err
	}

	return user, nil
}
