# API Design

The API is JSON REST over HTTP. Health is exposed at `GET /healthz`. Protected resources will use server-side browser sessions in an HttpOnly cookie; session identifiers will never appear in URLs or JSON responses. Validation errors will use a stable JSON shape and authorization will be enforced server-side.

Cookie-authenticated browser routes use the `csvault_session` HttpOnly cookie. Public review responses omit reviewer identity. Request bodies are limited to 1 MiB, reject unknown fields, and must contain exactly one JSON value.

## Implemented resources

- `POST /api/v1/auth/login` accepts `{username,password}` and sets an HttpOnly session cookie.
- `POST /api/v1/auth/logout` revokes the current session.
- `GET /api/v1/auth/me` returns the authenticated user's public account fields.
- `POST /api/v1/auth/activate` accepts `{token,password}` and initializes a pending account.
- `POST /api/v1/auth/password` accepts `{current_password,new_password}` for authenticated password changes.
- `POST /api/v1/auth/reset-request` accepts `{identifier}` and returns a non-enumerating response.
- `POST /api/v1/auth/reset` accepts `{token,password}`.
- `POST /api/v1/admin/accounts` accepts `{username,full_name,email,role}` and creates a pending account.
- `PATCH /api/v1/admin/accounts/status` accepts `{user_id,status}`.
- `GET /api/v1/terms` lists academic terms.
- `GET /api/v1/courses?q=` searches active courses by code or name.
- `GET /api/v1/course-offerings?course_id=` lists offerings, newest terms first.
- `POST /api/v1/my-courses`, `GET /api/v1/my-courses`, and `DELETE /api/v1/my-courses/{id}` manage a Student's saved offerings.
- `POST /api/v1/admin/terms` creates a term; `PUT /api/v1/admin/terms/current` selects the current term.
- `GET /api/v1/reviews?offering_id=&professor_id=` lists published, anonymous reviews with course and term context.
- `GET /api/v1/admin/reviews` and `GET /api/v1/admin/reports` provide moderation queues.

The API is not yet a complete FRS contract. See `FRS_IMPLEMENTATION_STATUS.md` for implemented boundaries and known gaps.
