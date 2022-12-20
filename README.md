# OAuth2.0이란
[OAuth2.0](https://www.rfc-editor.org/rfc/rfc6749)은 "Open Auhorization"으로 website또는 application이 user를 대신하여 web app에 호스팅된 자원들에 접근할 수 있도록 하기위해 디자인되었다. 이는 2012년에 나온 OAuth 1.0을 대신하며 현재 온라인 인증/인가에 있어 하나의 표준이 되었다. OAuth 2.0은 합의된 접근을 제공하고, user's credential 을 계속해서 공유하지않고도 client app이 user를 대신하여 자원에 행위를 할 수 있는 action들을 제한한다.

비록 web이 OAuth 2.0을 위한 main platform이지만, OAuth 2.0 스펙은 또한 어떻게 다른 client type에 일종의 위임된 접근을 제공할 지에 대해서 자세히 서술한다.

# Principles of OAuth 2.0
OAuth 2.0은 authorization(허가,인가) protocol이지 authentication(인증) protocol이 아니다. 즉, 어떤 자원에 대한 접근 권한(인가)을 받는 방법으로서 디자인되었다는 것인데 가령, 유저 데이터나 remote API들이 있다.

정리하자면 **인가**는 어떤 자원이나 행위에 대한 접근 **권한**에 관한 것이고, **인증**은 누구인가를 확인하는 것이다.

내가 A라는 회사의 임직원임을 확인하는 것은 **authentication(인증)**이 된다. 그러나, 내가 회사 안에 들어가서 비밀 문서들에 접근하는 것은 **authorization(인가)**에 관한 것이다. OAuth2.0은 특정 자원에 대한 **인가**에 대한 것이지, **인증**에 초점을 두지 않았다는 것을 알도록 하자.

OAuth 2.0은 Access Tokens을 사용한다. Access Tokens은 특정 자원에 대한 end user의 접근 권한, 즉 **authorization(인가)**을 의미한다. 다만, OAuth 2.0은 Access Tokens에 대한 어떠한 구체적인 포맷을 정하지 않았다. 그러나, 일부 context에서 JSON Web Token(JWT) 토큰이 자주 사용되고 있다. 이는 token 발행자가 토큰 그 자체에 데이터를 넣을 수 있도록 할 수 있기 때문이다. 또한, 보안상의 이유로 Access Tokens는 expiration date를 가지고 있어야 한다.

# OAuth 2.0 Roles
OAuth 2.0을 구현하는 프레임워크는 다음의 필수적인 OAUth 2.0 system component들을 정의해야 한다.

1. **Resource Owner**: 보호되고 있는 자원들을 소유하고 있는 user또는 system으로 누군가에게 해당 resource에 대한 권한을 부여할 수 있다.
2. **Client**: 보호되고 있는 자원에 대한 접근을 원하는 시스템이다. 자원들에 접근하기 위해서는 client가 적절한 **Access Token**을 가지고 있어야 한다.
3. **Authorization Server**: 이 서버는 Access Token들에 접근하기 위한 **Client**로부터 요청을 받는다. 그리고 **authentication(인증)**이 성공하고, 자원 소유자의 동의가 있을 때 Access Token들을 **Client**에게 발행해준다. **Authorization Server**은 두 가지 endpoints를 열어주는데, 하나는 **Authorization endpoint**으로 인터렉티브하게 인증(authentication)을 처리하고 유저의 동의를 처리한다. 다른 하나는 **Token endpoint**으로 시스템(machine) 간 상호 작용과 관련되어있다.
4. **Resource Server**: user의 자원을 보호하고 client로 부터 접근 요청을 받는 서버이다. 해당 서버는 **Access Token**을 **Client**로 부터 받아 검증하고 허가해주며, 적절한 자원을 client에게 반환해준다.

# OAuth 2.0 Scope
Scopes는 OAuth 2.0에게 굉장히 중요한 개념이다. Scopes는 부여받을 수 있는 어떤 자원에 대한 접근의 이유에 대해 정확하게 구체화할 때 사용된다. 수용할 수 있는 scopre값들 관련된 자원들은 **Resource Server**에 의존한다.

# OAuth 2.0 Access Tokens and Authorization code
OAuth 2.0 인가 서버는 **Resource Owner**가 접근을 인가했다고 해서 바로 직접 **Access Token**을 반환하지 않는다. 대신에 더 나은 보안을 위해서 **Authorization code**가 반환될 수 있는데 이는 **Access Token**으로 교환될 수 있다. 게다가 인가 서버는 **Refresh Token**을 **Access Token**과 함께 발간할 수 있다. **Access Token**과 달리 **Refresh Token**은 대게 **Access Tokens**보다 더 긴 유효 기간을 갖는데, 이는 새로운 **Access Tokens**로 교환되기 위함이다. **Refresh Tokens**가 이러한 특성들을 갖기 때문에 이들은 안전하게 client측에서 저장되야 한다.

# How Does OAuth 2.0 Work?
거정 기본적으로 OAuth2.0을 사용하기 이전에 **Authorization server**로 부터 Access Token을 요구할 때, 그 자신을 인증(authentication)하고 식별하기 위해서 client은 credential을 가지고 있어야 한다. 이는 **Authorization server**로 부터 받을 수 있다. 

OAuth 2.0 접근 요청들은 모두 client들로부터 시작된다. 가령 web, app, smart tv, desktop app 등등이 있다. token을 요구하고 교환되며, 응답하는 구조는 아래의 일반적인 flow를 따른다.

1. client은 인가(authorization) 요청을 **Authorization server(인가 서버)**에 보내면서 **client_id(application 식별)**와 **secret(식별한 application의 비밀번호)**(credential)을 식별을 위해 전송한다. 이는 또한 scope와  Authorization code에 또는 Access Token을 전송할 URI인 endpoint(redirect URI)을 같이 제공한다.
2. 인기 서버는 client를 식별하여 인증해주고, 요청된 scope가 허용가능한 지를 검증한다.
3. Resource Owner는 Authorization server와 상호작용하여 접근 권한을 부여한다.
4. 인가 서버는 client를 Authorization Code 또는 Access Token와 함께 리다이렉트한다. 추가적으로 **Refresh Token**이 주어질 수 있다. 이는 **grant type**(허가 종류)이 무엇이냐에 따라 다르다.
5. **Access Token**과 함께 client는 Resource server로 부터 자원에 대한 접근 요청을 보낸다.

# Grant Types in OAuth 2.0
OAuth 2.0에서 **grants(권한)**은 client가 자원에 대한 접근 인가를 받기위해 수행해야할 일련의 단계들이다. authorization framework은 다른 시나리오를 다루고 있는 몇가지 grant 타입들을 제공한다.

1. **Authorization Code** grant: 인가 서버는 단일 사용 Authorization code를 client에 반환한다. 이는 Access Token과 교환되기 위함이다. 이는 교환(exchange)가 안전하게 서버 사이드에서 이루어질 수 있는 전통적인 web app에게 가장 좋은 선택지이다. **Authroization code**흐름은 SPA와 mobile appe들에 많이 사용된다. 그러나, client secret이 안전하게 저장되지 않기 때문에, 교환 중에 인증은 client_id만 사용하는 것으로 제한된다. 더 나은 차선책은 **Authorization Code with PKCE** grant이다.

2. **Implicit** grant: 단순화된 흐름으로 바로 client에게 Access Token이 주어진다. 인가 서버는 Access Token을 callback URI의 파라미터로서 전송하거나 또는 form post에 대한 응답으로 전송한다. callback URI를 사용하는 첫번째 방식은 잠재적인 token 노출때문에 더 이상 사용하지 않는다.

3. **Authorization Code Grant with Proof Key for Code Exchange(PKCE)**: 인가 흐름은 **Authorization code** grant와 비슷하지만, 추가적인 단계가 있다. 이를 통해 mobile app과 spa에서 더 안전하게 사용할 수 있도록 하였다.

4. **Resource Owner Credential Grant Type**: 이 grant는 client가 먼저 Authorization server로 전송된 자원 소유자의 credential를 획득하도록 요구한다. 이는 client가 완전히 믿을 수 있는 자로 제한된다. 이는 인증 서버(Authorization server)에 대한 어떠한 redirect가 없다는 장점이 있으며, redirect가 적절하지 않은 시스템에서 사용하기가 좋다.

5. **Client Credentials Grant Type**: 비-상호작용적인 application에서 사용된다. 가령, 자동화된 process들이나 msa같은 것들이 있다. 이 경우는 application이 client_id와 암호를 사용하여 그 자체로 인증된다.

6. **Refresh Token Grant**: 이 흐름은 Refresh token으로 새로운 Access Token을 교환하기 위한 것과 관련되어 있다.

# How to use OAuth 2.0 in golang
google ouath를 사용해서 golang으로 oauth 2.0을 사용해보자.

먼저 google platform에서 oauth에 필요한 내용들을 등록해야한다. 위애서 살펴본, `client_id, secret, redirect URL`이 이에 해당한다.

https://console.cloud.google.com/apis/dashboard

마지막에 얻는 `client_id`와 `secret`을 잘 기억하도록 하자.

`.env`파일을 만들어 다음과 같이 기록하도록 하자.
```env
GOOGLE_OAUTH_CLIENT_ID=<client_id>
GOOGLE_OAUTH_SECRET_KEY=<secret_ket>
```

이제 해당 정보를 사용하여 oauth를 구현해보자.

먼저 `oauth2` go 패키지를 다운로드하자, 해당 패키지는 다양한 oauth 인터페이스들을 하나로 통일하여 처리할 수 있도록 해준다. 또한 `google` oauth를 사용할 것이기 때문에 `golang.org/x/oauth2/google` 패키지도 받도록 해주자.

```go
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
```
