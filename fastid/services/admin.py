from .. import repositories, services, typing
from ..exceptions import BadRequestException
from . import models


async def create_user(*, email: typing.Email, password: typing.Password) -> models.Token:
    config = await services.config.get()
    if config.is_setup:
        raise BadRequestException

    user_id = await repositories.users.create(email=email, password=password)
    token = await services.tokens.create(user_id=user_id, audience='internal')
    await services.config.update(key='is_setup', value='1')
    return token