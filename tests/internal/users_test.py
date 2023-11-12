import httpx

from fastid import services, typing


async def test_signin(client: httpx.AsyncClient, db_migrations):
    await services.users.create(
        email=typing.Email('user@exmaple.com'),
        password=typing.Password('qazwsx12345'),
    )

    response = await client.post(
        url='/api/v1/internal/signin/',
        json={'email': 'user@exmaple.com', 'password': 'qazwsx12345', 'captcha': '1'},
    )

    assert response.status_code == httpx.codes.CREATED
    assert response.json().get('access_token')
    assert response.json().get('refresh_token')
    assert response.json().get('expires_in')
    assert response.json().get('token_type')

    refresh_token = response.json().get('refresh_token')
    response = await client.post(
        url='/api/v1/internal/refresh_token/',
        json={
            'refresh_token': refresh_token,
        },
    )
    assert response.status_code == httpx.codes.CREATED
    assert response.json().get('access_token')
    assert response.json().get('refresh_token')
    assert response.json().get('expires_in')
    assert response.json().get('token_type')


async def test_signin_not_found_user(client: httpx.AsyncClient, db_migrations):
    response = await client.post(
        url='/api/v1/internal/signin/',
        json={'email': 'user@exmaple.com', 'password': 'qazwsx12345'},
    )
    assert response.status_code == httpx.codes.BAD_REQUEST


async def test_info(client_auth: httpx.AsyncClient):
    response = await client_auth.get(url='/api/v1/internal/info/')
    assert response.status_code == httpx.codes.OK
    assert response.json().get('user_id')
    assert response.json().get('email') == typing.Email('user@exmaple.com')