from argon2 import PasswordHasher
from argon2.profiles import RFC_9106_HIGH_MEMORY, RFC_9106_LOW_MEMORY
from opentelemetry.trace import get_current_span

from .. import typing
from ..settings import PasswordHasherMemoryProfile, settings
from ..trace import decorator_trace


@decorator_trace(name='services.password_hasher.hasher')
async def hasher(*, password: typing.Password) -> str:
    span = get_current_span()
    span.set_attribute('password_hasher_memory_profile', settings.password_hasher_memory_profile)

    if settings.password_hasher_memory_profile == PasswordHasherMemoryProfile.high:
        ph = PasswordHasher(
            time_cost=RFC_9106_HIGH_MEMORY.time_cost,
            memory_cost=RFC_9106_HIGH_MEMORY.memory_cost,
            parallelism=RFC_9106_HIGH_MEMORY.parallelism,
            hash_len=RFC_9106_HIGH_MEMORY.hash_len,
            salt_len=RFC_9106_HIGH_MEMORY.salt_len,
        )
    else:
        ph = PasswordHasher(
            time_cost=RFC_9106_LOW_MEMORY.time_cost,
            memory_cost=RFC_9106_LOW_MEMORY.memory_cost,
            parallelism=RFC_9106_LOW_MEMORY.parallelism,
            hash_len=RFC_9106_LOW_MEMORY.hash_len,
            salt_len=RFC_9106_LOW_MEMORY.salt_len,
        )
    return ph.hash(password=password)