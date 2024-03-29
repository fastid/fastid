from datetime import datetime

from sqlalchemy import delete, insert, select

from .. import repositories, typing
from ..trace import decorator_trace
from . import schemes


@decorator_trace(name='repositories.tokens.create')
async def create(
    *,
    token_id: typing.TokenID,
    access_token: str,
    refresh_token: str,
    user_id: typing.UserID,
    expires_at: datetime,
) -> schemes.Tokens | None:
    async with repositories.db.async_session() as session:
        stmt = (
            insert(schemes.Tokens)
            .values(
                token_id=token_id,
                access_token=access_token,
                refresh_token=refresh_token,
                expires_at=expires_at,
                user_id=user_id,
            )
            .returning(schemes.Tokens)
        )

        result = await session.execute(stmt)
        await session.commit()
        return result.scalar()


@decorator_trace(name='repositories.tokens.get_by_id')
async def get_by_id(*, token_id: typing.TokenID) -> schemes.Tokens | None:
    async with repositories.db.async_session() as session:
        stmt = select(schemes.Tokens).where(schemes.Tokens.token_id == token_id)
        result = await session.scalar(stmt)
        await session.commit()
        return result


@decorator_trace(name='repositories.tokens.delete_by_id')
async def delete_by_id(*, token_id: typing.TokenID) -> bool:
    async with repositories.db.async_session() as session:
        stmt = delete(schemes.Tokens).where(schemes.Tokens.token_id == token_id)
        result = await session.execute(stmt)
        await session.commit()

        if result.rowcount:
            return True
        return False
