# Security

<!-- Passwords use bcrypt with cost 12. Activation and reset tokens use at least 32 bytes from the operating-system CSPRNG and are represented with URL-safe base64. Browser sessions are stored server-side, use HttpOnly cookies, Secure in production, and SameSite Lax. Session inactivity and absolute lifetimes default to 30 minutes and 24 hours and are configurable.

Review authors remain linked internally through `student_user_id`; public review responses must omit that field. Database uniqueness prevents duplicate reviews per offering. Review status controls aggregation: only `published` rows are eligible. Review submission eligibility, five-minute cooldowns, and ten-per-term limits must be checked in a transaction before insertion.

CSRF protection, rate limits, token persistence, account status transitions, and audit logging will be completed with the remaining authentication and moderation handlers. Secrets remain environment-provided. -->
