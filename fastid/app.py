import re
from contextlib import asynccontextmanager

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from fastapi.middleware.trustedhost import TrustedHostMiddleware
from fastapi.responses import ORJSONResponse
from prometheus_fastapi_instrumentator import Instrumentator
from starlette.exceptions import HTTPException
from starlette.exceptions import HTTPException as StarletteHTTPException
from starlette.staticfiles import StaticFiles

from . import __version__, handlers, internal, middlewares, v1
from .exceptions import exc_handlers
from .logger import logger
from .settings import Environment, settings


@asynccontextmanager
async def lifespan(app: FastAPI):
    """
    Lifespan Events
    See - https://fastapi.tiangolo.com/advanced/events/

    :param app: FastAPI

    """
    logger.info('Startup FastID service', extra={'environment': settings.environment.lower()})
    yield


app = FastAPI(
    title=settings.app_name,
    description='Authorization and authentication service',
    default_response_class=ORJSONResponse,
    docs_url='/api/',
    redoc_url=None,
    openapi_url='/api/v1/openapi.json',
    version=__version__,
    debug=True if settings.environment.development == Environment.development else False,
    swagger_ui_parameters={
        'displayRequestDuration': True,
        'persistAuthorization': True,
    },
    lifespan=lifespan,
    license_info={
        'name': 'MIT',
        'url': 'https://github.com/fastid/fastapi/blob/main/LICENSE',
    },
    contact={'name': 'Github', 'url': settings.link_github},
    exception_handlers=exc_handlers,
)

# Prometheus metrics
Instrumentator(excluded_handlers=['/healthcheck/', '/metrics']).instrument(
    app,
    metric_namespace=settings.app_name.lower(),
).expose(app, include_in_schema=False)

# Middlewares
app.add_middleware(
    TrustedHostMiddleware,
    allowed_hosts=settings.trusted_hosts.split(','),
)

if settings.cors_enable:
    app.add_middleware(
        CORSMiddleware,
        allow_origins=settings.cors_allow_origins.split(','),
        allow_credentials=settings.cors_allow_credentials,
        allow_methods=settings.cors_allow_methods.split(','),
        allow_headers=settings.cors_allow_headers.split(','),
        expose_headers=settings.cors_expose_headers.split(','),
    )

# App middleware
app.add_middleware(middlewares.Middleware)


class SPAStaticFiles(StaticFiles):
    async def get_response(self, path: str, scope):
        if re.match(r'^api', path):
            return await super().get_response(path, scope)

        try:
            return await super().get_response(path, scope)
        except (HTTPException, StarletteHTTPException) as ex:
            if ex.status_code == 404:
                return await super().get_response('index.html', scope)


app.include_router(handlers.healthcheck.router)

# Internal API
app.include_router(internal.config.router, prefix='/api/v1/internal', tags=['Internal API'], include_in_schema=True)
app.include_router(internal.users.router, prefix='/api/v1/internal', tags=['Internal API'], include_in_schema=True)
# app.include_router(internal.otp.router, prefix='/api/v1/internal', tags=['Internal API'], include_in_schema=True)

# API
app.include_router(v1.users.router, prefix='/api/v1', tags=['Users'])


app.mount(
    path='/',
    app=SPAStaticFiles(
        directory=f'{settings.base_dir}/fastid/static',
        html=True,
    ),
    name='spa-static-files',
)
